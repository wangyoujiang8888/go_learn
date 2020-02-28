package go_redis_test

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"time"
)
type MainDB struct {
	Db *gorm.DB
}

func NewMainDB(conf *viper.Viper) *MainDB {
	user := conf.GetString("main_db.user")
	pwd := conf.GetString("main_db.password")
	hostname := conf.GetString("main_db.hostname")
	db := conf.GetString("main_db.db")
	url := fmt.Sprintf(`%s:%s@tcp(%s)/%s?parseTime=true`, user, pwd, hostname, db)
	mainDB, err := gorm.Open("mysql", url)
	if err != nil {
		panic("init master db connect failed" + err.Error())
	}
	mainDB.DB().SetMaxIdleConns(conf.GetInt("db.max_idle_conn"))
	mainDB.DB().SetMaxOpenConns(conf.GetInt("db.max_open_conn"))
	mainDB.DB().SetConnMaxLifetime(time.Hour)
	return &MainDB{
		Db: mainDB,
	}
}

