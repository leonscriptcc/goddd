package gmysql

import (
	"fmt"
	"github.com/leonscriptcc/goddd/infrastructure/gconfig"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// MysqlDB mysql-gorm
type MysqlDB struct {
	db *gorm.DB
}

// NewMysqlDB 初始化mysql连接
func NewMysqlDB() (MysqlDB, error) {
	// 组装数据源
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local",
		gconfig.Parameters.MysqlConfig.User, gconfig.Parameters.MysqlConfig.Password,
		gconfig.Parameters.MysqlConfig.Host, gconfig.Parameters.MysqlConfig.Port,
		gconfig.Parameters.MysqlConfig.Schema)

	// 打开ODBC连接
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})

	if err != nil {
		return MysqlDB{}, err
	}

	// 维护连接池
	sqlDB, err := db.DB()

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(gconfig.Parameters.MysqlConfig.MaxIdleConns)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(gconfig.Parameters.MysqlConfig.MaxOpenConns)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Duration(gconfig.Parameters.MysqlConfig.ConnMaxLifetime) * time.Minute)

	//err = db.AutoMigrate(Device{}, Terminal{})

	return MysqlDB{db: db}, err
}
