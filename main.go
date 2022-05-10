package main

import (
	"douyijn-simple/routers"
	"fmt"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: routers.InitRouter(),
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server started error", err)
	}
}
