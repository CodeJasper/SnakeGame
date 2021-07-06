package main

import (
	"context"
	"fmt"
	"net/http"
	"snakeGameApi/config"
)

func main() {
	config.ConfigDB()
	defer config.Connect.Close(context.Background())

	http.HandleFunc("/", index)

	var server = http.ListenAndServe(":8080", nil)
	fmt.Println("Estado del servidor", server)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola mundo")
}
