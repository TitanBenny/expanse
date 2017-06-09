package main

import (
	"log"

	"github.com/astaxie/beego"
)

func main() {
	go beego.Run()
	log.Println("wx expanse")
}
