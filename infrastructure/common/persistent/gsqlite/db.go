package gsqlite

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// DB 数据库
type DB struct {
	sqliteDB *gorm.DB
}

// NewDB 工厂
func NewDB(dsn string) (*DB, error) {
	// 日志级别
	//TODO 可以自定义日志实现
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second,  // 慢 SQL 阈值
			LogLevel:                  logger.Error, // 日志级别
			IgnoreRecordNotFoundError: true,         // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,        // 禁用彩色打印
		},
	)
	sqliteDB, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	return &DB{sqliteDB: sqliteDB}, err
}

func (d *DB) GetSqliteDB() *gorm.DB {
	return d.sqliteDB
}
