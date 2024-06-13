package models

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var Db *gorm.DB

func init() {
    Db = Connect()
}

func Connect() *gorm.DB {
    //DB接続
    dsn := "root@tcp(mysql:3306)/first?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

    if err != nil {
        panic("failed to connect database")
    }
    return db
}
