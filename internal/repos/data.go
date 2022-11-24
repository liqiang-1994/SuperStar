package repos

import (
	"SuperStar/internal/config"
	"SuperStar/internal/services"
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/go-redis/redis"
	"github.com/minio/minio-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
)

type contextKey struct{}

type Data struct {
	db    *gorm.DB
	rdb   *redis.Client
	minio *minio.Client
	GID   *snowflake.Node
}

func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db
}

func NewDBData(db *gorm.DB, rdb *redis.Client, minio *minio.Client) (*Data, func(), error) {
	node, err := snowflake.NewNode(1)
	fmt.Println("snowflake.NewNode")
	if err != nil {
		fmt.Println(err)
	}
	d := &Data{
		db:    db,
		rdb:   rdb,
		minio: minio,
		GID:   node,
	}
	return d, func() {
	}, nil
}

func NewDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", cfg.DB.Url, cfg.DB.UserName, cfg.DB.Password, cfg.DB.DBName, cfg.DB.Port)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	//db.AutoMigrate(&models.User{})
	return db
}

func NewTransaction(d *Data) services.Transaction {
	return d
}

func (d *Data) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx := context.WithValue(ctx, contextKey{}, tx)
		return fn(ctx)
	})
}

func NewRedis(cfg *config.Config) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Url + ":" + strconv.Itoa(cfg.Redis.Port),
		DB:       cfg.Redis.DB,
		Password: cfg.Redis.Password,
	})
	if err := client.Ping().Err(); err != nil {
		panic("Failed to connect Redis")
	}
	return client
}

func NewMinio(cfg *config.Config) *minio.Client {
	minioClient, err := minio.New(cfg.Minio.Url, cfg.Minio.AccessKeyID, cfg.Minio.SecretAccessKey, false)
	if err != nil {
		log.Fatal(err)
	}
	err = minioClient.MakeBucket(cfg.Minio.BucketName, "xian")
	if err != nil {
		exists, err := minioClient.BucketExists(cfg.Minio.BucketName)
		if err == nil && exists {
			log.Printf("bucket exists, skip")
		} else {
			log.Println("bucket err", err)
		}
	}
	log.Println("Success connect minio")
	return minioClient
}
