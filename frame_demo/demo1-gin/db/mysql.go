package db

import (
	"errors"
	"fmt"
	"sync"

	orm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

type MysqlPoll struct{}

var instance *MysqlPoll
var once sync.Once

var db *orm.DB

var err error

//单例模式
func GetInstance() *MysqlPoll {
	once.Do(func() {
		instance = &MysqlPoll{}
	})
	return instance
}

func (pool *MysqlPoll) InitPoll() bool {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s", viper.GetString("db.username"), viper.GetString("db.password"), viper.GetString("db.host"), viper.GetString("db.name"), viper.GetString("db.charset"))
	// fmt.Println(dsn)
	db, err := orm.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		panic(errors.New("mysql连接失败"))
		return false
	}

	// 连接数配置也可以写入配置，在此读取
	db.DB().SetMaxIdleConns(viper.GetInt("db.MaxIdleConns"))
	db.DB().SetMaxOpenConns(viper.GetInt("db.MaxOpenConns"))

	return true
}
