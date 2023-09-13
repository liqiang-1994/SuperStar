package repos

import (
	"SuperStar/internal/config"
	"SuperStar/internal/entity"
	"SuperStar/internal/services"
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/go-redis/redis"
	"github.com/minio/minio-go"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
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
	sms   *sms.Client
}

func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db
}

func NewDBData(db *gorm.DB, rdb *redis.Client, minio *minio.Client, sms *sms.Client, GID *snowflake.Node) (*Data, func(), error) {
	d := &Data{
		db:    db,
		rdb:   rdb,
		minio: minio,
		GID:   GID,
		sms:   sms,
	}
	return d, func() {
	}, nil
}

func NewDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", cfg.DB.Url, cfg.DB.UserName, cfg.DB.Password, cfg.DB.DBName, cfg.DB.Port)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	err = db.AutoMigrate(
		&entity.Circle{},
		&entity.CircleUserRel{},
		&entity.Comment{},
		&entity.TagRel{},
		//&entity.Idiom{},
		//&entity.Poet{},
		//&entity.Poetry{},
		//&entity.Saying{},
		&entity.Tag{},
		&entity.User{},
		&entity.UserRel{},
	)
	if err != nil {
		log.Fatal("Failed to migrate. \n", err)
		os.Exit(1)
	}
	log.Println("🚀 Connected Successfully to the Database")
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
	log.Println("🚀 Connected Successfully to the Redis")
	return client
}

func NewGID() *snowflake.Node {
	node, err := snowflake.NewNode(1)
	fmt.Println("🚀 Generate snowflake.NewNode")
	if err != nil {
		log.Fatal("Failed to Generate snowflake.NewNode.", err)
	}
	return node
}

func NewMinio(cfg *config.Config) *minio.Client {
	minioClient, err := minio.New(cfg.Minio.Url, cfg.Minio.AccessKeyID, cfg.Minio.SecretAccessKey, false)
	if err != nil {
		log.Fatal(err)
	}
	//err = minioClient.MakeBucket(cfg.Minio.BucketName, "xian")
	//if err != nil {
	//	exists, err := minioClient.BucketExists(cfg.Minio.BucketName)
	//	if err == nil && exists {
	//		log.Printf("bucket exists, skip")
	//	} else {
	//		log.Println("bucket err", err)
	//	}
	//}
	//log.Println("Success connect minio")
	return minioClient
}
