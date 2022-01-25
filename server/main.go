package main

import (
	"fmt"
	"github.com/nao-18/test-todo-app/config"
	"log"
)

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)

	log.Println("test")
}
