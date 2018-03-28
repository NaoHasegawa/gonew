package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Create Lightweight DDD")
	os.Mkdir("application", 0666)
	os.Mkdir("domain", 0666)
	os.Mkdir("infrastructure", 0666)
	os.Mkdir("interfaces", 0666)
	os.Mkdir("library", 0666)
	os.OpenFile("main.go", os.O_RDWR|os.O_CREATE, 0666)
}