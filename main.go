package main

import (
	"fmt"
	"github.com/brix-go/brix/cmd/brix"
)

func main() {
	err := brix.Execute()
	if err != nil {
		fmt.Println("execute error: ", err.Error())
	}
}
