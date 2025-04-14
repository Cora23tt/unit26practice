package main

import (
	"banner/internal/core/application"
	"banner/internal/infra/api/rest"
	"net/http"
	"os"
)

func main() {
	err := execute("8080")
	if err != nil {
		print("cant run app:", err)
		os.Exit(1)
	}

}

func execute(port string) error {
	mux := http.NewServeMux()
	app := application.NewApplication()
	router := rest.NewRouter(mux, app).Run()
	srv := http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	return srv.ListenAndServe()
}
