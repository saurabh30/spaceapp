package usecase

import (
	"fmt"
	"spaceapp/domain"
)

/*
AddExoplanet(p Exoplanet) int
ListExoplanet() ([]Exoplanet, error)
GetExoplanetById(id int) (Exoplanet, error)
UpdateExoplanet(id int, p Exoplanet) error
DeleteExoplanet(id int) error
FuelEstimation()
*/
type ExoplanetUsecase struct {
	Repository domain.ExoplanetRepository
}

func (u *ExoplanetUsecase) AddExoplanet(p domain.Exoplanet) int {
	return u.Repository.AddExoplanet(p)
}

func (u *ExoplanetUsecase) ListExoplanet() ([]domain.Exoplanet, error) {
	return u.Repository.ListExoplanet()
}

func (u *ExoplanetUsecase) GetExoplanetById(id int) (domain.Exoplanet, error) {
	return u.Repository.GetExoplanetById(id)
}

func (u *ExoplanetUsecase) UpdateExoplanet(id int, Exoplanet domain.Exoplanet) error {
	return u.Repository.UpdateExoplanet(id, Exoplanet)
}

func (u *ExoplanetUsecase) DeleteExoplanet(id int) error {
	return u.Repository.DeleteExoplanet(id)
}

/*
f = d / (g^2) * c units
d -> distance of exoplanet from earth
g -> gravity of exoplanet
c -> crew capacity (int)
*/

func (u *ExoplanetUsecase) FuelEstimation(p domain.Exoplanet, crewSize int) (f float64, err error) {
	if p.Radius == 0 {
		return 0, fmt.Errorf("zero radius")
	}

	var gravity float64
	switch p.Type {
	case domain.GasGiant:
		g := domain.NewGas(p)
		gravity, _ = g.GravityEstimation()
	case domain.Terrestrial:
		t := domain.NewTerre(p)
		gravity, _ = t.GravityEstimation()
	default:
		return 0, fmt.Errorf("invalid type")
	}
	return fuel(p.Distance, crewSize, gravity)
}

func fuel(distance, crewSize int, gravity float64) (float64, error) {
	if crewSize == 0 {
		return 0, fmt.Errorf("crew size zero")
	}
	if gravity == 0 {
		return 0, fmt.Errorf("gravity zero")
	}

	return float64(distance) / (gravity * gravity * float64(crewSize)), nil
}
