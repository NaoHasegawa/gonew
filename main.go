package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Create Lightweight DDD")
	os.Mkdir("application", 0755)
	os.Mkdir("domain", 0755)
	os.Mkdir("infrastructure", 0755)
	os.Mkdir("interfaces", 0755)
	os.Mkdir("library", 0755)
	os.OpenFile("main.go", os.O_RDWR|os.O_CREATE, 0755)
}
