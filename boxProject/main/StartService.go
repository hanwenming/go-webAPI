package main

import (
	"boxProject/boxProject/service"
	"boxProject/tools"
	"fmt"
	restful "github.com/emicklei/go-restful/v3"
	"log"
	"net/http"
	"os"
)

func main() {

	err := tools.InitMysql()
	if err != nil {
		os.Exit(1)
	}

	container := restful.NewContainer()
	u := service.UserResource{}
	u.RegisterTo(container)
	server := &http.Server{Addr: ":8080", Handler: container}
	defer server.Close()
	ip, err := tools.GetHostIp()
	if err != nil {
		fmt.Println(err)
	}
	log.Println("当前主机IP : ", ip)
	log.Fatal(server.ListenAndServe())
}
