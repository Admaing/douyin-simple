package main

import (
	"douyijn-simple/model"
	"douyijn-simple/routers"
	_ "douyijn-simple/utils"
	"fmt"
	"net/http"
)

func main() {

	model.InitDb()
	server := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: routers.InitRouter(),
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("Server started error", err)
	}
}
