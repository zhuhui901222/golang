package main
import (
	"database/sql"
	"fmt"
	"strings"
	_ "github.com/go-sql-driver/mysql"
)

//github.com/go-sql-driver/mysql

//数据库配置
const (
userName    = "root"
password    = "root"
ip          = "127.0.0.1"
port        = "3306"
dbName      = "test"
)

func main(){
	path := strings.Join([]string{userName, ":", password, "@tcp(",ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	DB, _ := sql.Open("mysql", path)
	//验证连接
	if errConn := DB.Ping(); errConn != nil{
	fmt.Println("open database fail")
	return
	}
	fmt.Println("connnect success")
	defer DB.Close()
}

