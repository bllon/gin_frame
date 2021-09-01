package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dataSourceNameFormat = "%s:%s@tcp(%s:%s)/%s"
	driverName           = "mysql"
)

// Conf 创建数据库连接池所需的配置
type DBConf struct {
	Name        string
	Host        string
	Port        string
	UserName    string
	Password    string
	MaxLifeTime int
	MaxOpenConn int
	MaxIdleConn int
}

// DB 对sql.DB进行装饰, 对常用的操作方法进行封装
type DB struct {
	*sql.DB
}

// NewDB 返回包装了指定配置创建的DB连接池的DB实例
func NewDB(cf *DBConf) (db *DB, err error) {
	dsn := fmt.Sprintf(dataSourceNameFormat,
		cf.UserName,
		cf.Password,
		cf.Host,
		cf.Port,
		cf.Name,
	)

	odb, err := sql.Open(driverName, dsn)
	if err != nil {
		err = fmt.Errorf("sql open: %w", err)
		return
	}

	if err = odb.Ping(); err != nil {
		err = fmt.Errorf("db ping: %w", err)
		return
	}

	odb.SetConnMaxLifetime(time.Duration(cf.MaxLifeTime) * time.Second)
	odb.SetMaxOpenConns(cf.MaxOpenConn)
	odb.SetMaxIdleConns(cf.MaxIdleConn)

	db = &DB{
		DB: odb,
	}

	return
}
