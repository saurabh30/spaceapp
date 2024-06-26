package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"spaceapp/domain"
	"strconv"
)

type ExoplanetController struct {
	ExoplanetUsecase domain.ExoplanetUsecase
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("server is running..."))
}

type Response struct {
	Message   string             `json:"message,omitempty"`
	Exoplanet []domain.Exoplanet `json:"Exoplanet,omitempty"`
	Error     string             `json:"error,omitempty"`
}

func (controller *ExoplanetController) ExoplanetHandler(w http.ResponseWriter, r *http.Request) {
	var res Response

	defer func() {
		json.NewEncoder(w).Encode(res)
	}()

	w.Header().Add("Content-Type", "application/json")

	switch r.Method {
	case http.MethodPost:
		var p domain.Exoplanet
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			res.Error = err.Error()
			return
		}
		id := controller.ExoplanetUsecase.AddExoplanet(p)
		res.Message = fmt.Sprintf("New Exoplanet created with id : %d", id)
		w.WriteHeader(200)
	case http.MethodGet:
		strId := r.URL.Query().Get("id")
		if strId == "" {
			res.Exoplanet, _ = controller.ExoplanetUsecase.ListExoplanet()
			w.WriteHeader(200)
			return
		}
		id, err := strconv.Atoi(strId)
		if err != nil {
			log.Println("Invalid Id")
			res.Error = err.Error()
			w.WriteHeader(400)
			return
		}
		planet, err := controller.ExoplanetUsecase.GetExoplanetById(id)
		res.Exoplanet = append(res.Exoplanet, planet)
		if err != nil {
			res.Error = err.Error()
			w.WriteHeader(400)
			return
		}
		w.WriteHeader(200)
	case http.MethodPut:
		var p domain.Exoplanet
		strId := r.URL.Query().Get("id")
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			res.Error = err.Error()
			return
		}
		id, _ := strconv.Atoi(strId)
		if err != nil {
			log.Println("Invalid Id")
			res.Error = err.Error()
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = controller.ExoplanetUsecase.UpdateExoplanet(id, p)
		if err != nil {
			res.Error = err.Error()
		}
		res.Message = fmt.Sprintf("Exoplanet updated with id : %d", id)
		w.WriteHeader(200)

	case http.MethodDelete:
		strId := r.URL.Query().Get("id")
		if strId == "" {
			res.Error = "Empty Id"
			return
		}
		id, err := strconv.Atoi(strId)
		if err != nil {
			log.Println("Invalid Id")
			res.Error = err.Error()
			w.WriteHeader(400)
			return
		}
		err = controller.ExoplanetUsecase.DeleteExoplanet(id)
		if err != nil {
			res.Error = err.Error()
			w.WriteHeader(400)
			return
		}
		res.Message = fmt.Sprintf("planet deleted with id %d", id)
		w.WriteHeader(200)
	}
}

type FuelEstimationRequest struct {
	Exoplanet domain.Exoplanet `json:"exoplanet,omitempty"`
	CrewSize  int              `json:"crew_size,omitempty"`
}

func (controller *ExoplanetController) ExoplanetFuelHandler(w http.ResponseWriter, r *http.Request) {
	var res Response

	defer func() {
		json.NewEncoder(w).Encode(res)
	}()

	w.Header().Add("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		var p FuelEstimationRequest
		err := json.NewDecoder(r.Body).Decode(&p)
		if err != nil {
			res.Error = err.Error()
			return
		}
		f, err := controller.ExoplanetUsecase.FuelEstimation(p.Exoplanet, p.CrewSize)
		if err != nil {
			res.Error = err.Error()
			w.WriteHeader(400)
			return
		}
		res.Message = fmt.Sprintf("Fuel Estimation : %f", f)
		w.WriteHeader(200)
	}

}
