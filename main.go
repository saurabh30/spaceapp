package main

import (
	"log"
	"net/http"
	"spaceapp/controller"
	"spaceapp/repository"
	"spaceapp/usecase"
)

func main() {
	http.HandleFunc("/", controller.HealthCheck)
	repo := repository.NewExoplanetRepository()
	usecase := usecase.ExoplanetUsecase{
		Repository: repo,
	}
	ExoplanetController := controller.ExoplanetController{
		ExoplanetUsecase: &usecase,
	}
	http.HandleFunc("/Exoplanet", ExoplanetController.ExoplanetHandler)
	http.HandleFunc("/Exoplanet/Fuel", ExoplanetController.ExoplanetFuelHandler)
	log.Println("starting server")
	err := http.ListenAndServe(":4400", nil)
	if err != nil {
		log.Fatal("server error", err)
	}
	log.Println("server is stopped")
}
