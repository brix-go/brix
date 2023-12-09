package main

import (
	"fmt"
	"github.com/brix-go/brix/v2/cmd/brix"
)

func main() {
	err := brix.Execute()
	if err != nil {
		fmt.Println("execute error: ", err.Error())
	}
}
