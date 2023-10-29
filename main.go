package main

import (
	"api/src/config"
	"api/src/router"
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func init() {
	key := make([]byte, 64)

	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}

	stringBase64 := base64.StdEncoding.EncodeToString(key)

	fmt.Println(stringBase64)
}

func main() {

	config.Carregar()
	r := router.Generate()

	fmt.Printf("Port %d", config.Port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
