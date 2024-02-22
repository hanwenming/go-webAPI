package tools

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net"
)

var Db *sql.DB

// GetHostIp
//
//	@Description:  获取服务器IP
//	@return string IP
//	@return error 错误信息
func GetHostIp() (string, error) {
	addArr, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	var ipAddr string
	for _, addr := range addArr {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ipAddr = ipNet.IP.String()
			}
		}
	}
	return ipAddr, nil
}

// InitMysql
//
//	@Description: 初始化MySQL
//	@return err
func InitMysql() (err error) {

	// DSN:Data Source Name
	//"user:password@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	dsn := "root:root@tcp(192.168.110.74:3306)/goRestFul"
	//Open  函数只是校验   dsn  的查数是否正确，  并不会连接数据库
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	fmt.Println("连接成功？？？")

	//尝试与数据库进行连接
	err = Db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败", err)
		return
	}
	return
}
