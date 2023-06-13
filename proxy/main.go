package main

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"io"
	"log"
	"net"
)

type Config struct {
	ServerAddr string `yaml:"server" env:"SERVER_ADDR" env-default:""`
	ServerPort int    `yaml:"port" env:"SERVER_PORT" env-default:"8080"`
	ListenPort int    `yaml:"listen" env:"LISTEN_PORT" env-default:"80"`
}

func main() {

	var cfg Config
	err := cleanenv.ReadEnv(&cfg)

	serverAddr := &cfg.ServerAddr
	serverPort := &cfg.ServerPort
	listenPort := &cfg.ListenPort

	bind := fmt.Sprintf(":%d", *listenPort)
	remote := fmt.Sprintf("%s:%d", *serverAddr, *serverPort)

	fmt.Println(bind, "<=>", remote)

	listener, err := net.Listen("tcp", bind)
	if err != nil {
		fmt.Print(err)
		return
	}
	log.Default().Println("转发服务已启动...")
	for {
		client, err := listener.Accept()
		if err != nil {
			fmt.Print(err)
			continue
		}
		proxy, err := net.Dial("tcp", remote)
		if err != nil {
			fmt.Print(err)
			continue
		}
		go func() {
			defer closeConn(proxy)
			defer closeConn(client)
			//将客户端数据转发到服务器
			_, err := io.Copy(proxy, client)
			if err != nil {
				fmt.Println(err)
				return
			}
		}()
		go func() {
			defer closeConn(proxy)
			defer closeConn(client)
			//将服务端数据转发到客户端
			_, err := io.Copy(client, proxy)
			if err != nil {
				fmt.Println(err)
				return
			}
		}()
	}
}

func closeConn(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
