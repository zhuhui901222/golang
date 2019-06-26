package pool

import (
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

var once sync.Once
var instance *gorose.Connection

func GetInstance()(*gorose.Connection,error){
	var err error
	if instance==nil{
		once.Do(func() {
			instance,err=NewConn()
		})
	}
	return instance,err

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


//func main() {
//	session,err:=GetOrm()
//	res,err := session.Table("user").First()
//	if err!=nil{
//		fmt.Println("db.Table  err:",err)
//		return
//	}
//	fmt.Println(res)
//}