package main

import (
	"fmt"

	"github.com/waqarilyas/go_cloud_practice/decrypt"
	"github.com/waqarilyas/go_cloud_practice/encrypt"
)

func main() {
	result := encrypt.Nimbus("gjhdsvghj")
	fmt.Println(decrypt.Nimbus(result))
}
