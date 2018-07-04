package main

import (
	"flag"
	"fmt"

	"github.com/ComputePractice2018/AddressBook/backend/utils"
)

func main() {
	name := flag.String("name", "Иван", "имя для преветствия")
	flag.Parse()
	fmt.Println(utils.GetHelloWorldString(*name))

}
