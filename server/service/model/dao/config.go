package dao

import "time"

type Config struct {
	Host     string
	Port     int
	UserName string
	Password string
	DBName   string
	Charset  string

	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
	MaxOpenConns    int
	MaxIdleConns    int
}
