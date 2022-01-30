package main

import (
	"github.com/hucongyang/go-project-lianjia-lianjia_server/internal/routers"
	"net/http"
	"time"
)

func main() {
	router := routers.NewRouter()
	service := &http.Server{
		Addr: ":8080",
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	service.ListenAndServe()
}