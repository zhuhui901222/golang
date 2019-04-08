package main

import (
	"github.com/gohouse/gorose"
	"fmt"
	"reflect"
)

func main() {
	var dbConfig = &gorose.DbConfigSingle{
		Driver:          "mysql", // driver: mysql/sqlite/oracle/mssql/postgres
		EnableQueryLog:  true,    // if enable sql logs
		SetMaxOpenConns: 0,       // connection pool of max Open connections, default zero
		SetMaxIdleConns: 0,       // connection pool of max sleep connections
		Prefix:          "",      // prefix of table
		Dsn:             "root:root@tcp(localhost:3306)/test?charset=utf8", // db dsn
	}

	// 初始化数据库链接, 默认会链接配置中 default 指定的值
	// 也可以在第二个参数中指定对应的数据库链接, 见下边注释的那一行链接示例
	connection, err :=gorose.Open(dbConfig )
	if err != nil {
		fmt.Println("gorose.Open",err)
		return
	}
	// close DB
	defer connection.Close()
	fmt.Println(reflect.TypeOf(connection))
	session := connection.NewSession()

	res,err := session.Table("users").First()
	if err!=nil{
		fmt.Println("db.Table",err)
		return
	}
	fmt.Println(res)
}