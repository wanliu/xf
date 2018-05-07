package main

import (
	"log"

	"github.com/wanliu/xf"
)

func main() {
	if err := xf.MSPLogin("appid = 599cf4a7, work_dir = ."); err != nil {
		log.Fatalf("login failed %s", err)
	}

}
