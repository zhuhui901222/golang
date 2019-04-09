package pool

import (
	"fmt"
	"github.com/gohouse/gorose"
	"sync"
)

func GetOrm()(*gorose.Connection,error){
	DataBase,err:=GetInstance()
	if err!=nil{
		return nil,err
	}
	return DataBase,nil
}

var mu sync.Mutex
var instance *gorose.Connection

func GetInstance()(*gorose.Connection,error){
	mu.Lock()                    // <--- Unnecessary locking if instance already created
	defer mu.Unlock()
	if instance==nil{
		var err error
		instance,err=NewConn()
		if err != nil{
			return nil,err
		}
	}
	return instance,nil
}

//用于创建连接资源
func NewConn()(*gorose.Connection,error){
	var dbConfig = &gorose.DbConfigSingle{
		Driver:          "mysql", // driver: mysql/sqlite/oracle/mssql/postgres
		EnableQueryLog:  true,    // if enable sql logs
		SetMaxOpenConns: 100,       // connection pool of max Open connections, default zero
		SetMaxIdleConns: 10,       // connection pool of max sleep connections
		Prefix:          "",      // prefix of table
		Dsn:             "root:root@tcp(localhost:3306)/test?charset=utf8", // db dsn
	}
	// 初始化数据库链接, 默认会链接配置中 default 指定的值
	// 也可以在第二个参数中指定对应的数据库链接, 见下边注释的那一行链接示例
	connection,err :=gorose.Open(dbConfig )
	if err != nil {
		return nil,err
	}
	return connection,nil
}


func main() {
	session:=GetOrm()
	res,err := session.Table("users").First()
	if err!=nil{
		fmt.Println("db.Table  err:",err)
		return
	}
	fmt.Println(res)

}