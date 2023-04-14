package main

import (
	"fmt"
	"log"

	cepv1 "github.com/lucasmmo/brasil-api-sdk-go/service/cep_v1"
)

func main() {
    client := cepv1.NewClient()

    infos, err := client.FindByCepWithMultiplesProviders("03380210")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(infos)
}
