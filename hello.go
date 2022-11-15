package main

import (
	"fmt"
	"strings"

	"github.com/agrawalananya/goTraining/module1"
	"github.com/pborman/uuid"
)

const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

const (
	a = 37
	b = 89
)

func main() {
	uuidWithHyphen := uuid.NewRandom()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	fmt.Println(uuid)
	module1.Something()
	for i := 0; i < len(sample); i++ {
		fmt.Print(sample[i])
	}

	fmt.Println("\n", a, b)

}

func Some() {
	fmt.Println("Hello")
}
