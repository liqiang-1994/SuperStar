package repos

import (
	"SuperStar/internal/config"
	"SuperStar/internal/services"
	"context"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type contextKey struct{}

type DBData struct {
	db  *gorm.DB
	GID *snowflake.Node
}

func (d *DBData) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db
}

func NewDBData(db *gorm.DB) (*DBData, func(), error) {
	node, err := snowflake.NewNode(1)
	fmt.Println("snowflake.NewNode")
	if err != nil {
		fmt.Println(err)
	}
	d := &DBData{
		db:  db,
		GID: node,
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

func NewTransaction(d *DBData) services.Transaction {
	return d
}

func (d *DBData) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx := context.WithValue(ctx, contextKey{}, tx)
		return fn(ctx)
	})
}
