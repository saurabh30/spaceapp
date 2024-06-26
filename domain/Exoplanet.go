package domain

/*
10 < d < 1000 (light years) : int
0.1 < m < 10 (Earth-Mass unit) : double
0.1 < r < 10 (Earth-radius unit) : double
*/

type ExoplanetType int

const (
	GasGiant ExoplanetType = iota + 1
	Terrestrial
)

type Exoplanet struct {
	Name        string        `json:"name,omitempty"`
	Description string        `json:"description,omitempty"`
	Distance    int           `json:"distance,omitempty"`
	Radius      float64       `json:"radius,omitempty"`
	Mass        float64       `json:"mass,omitempty"` //(will be provided only in case of Terrestrial type of planet)
	Type        ExoplanetType `json:"type,omitempty"` //GasGiant or Terrestrial
}

type ExoplanetRepository interface {
	AddExoplanet(p Exoplanet) int
	ListExoplanet() ([]Exoplanet, error)
	GetExoplanetById(id int) (Exoplanet, error)
	UpdateExoplanet(id int, p Exoplanet) error
	DeleteExoplanet(id int) error
}

type ExoplanetUsecase interface {
	AddExoplanet(p Exoplanet) int
	ListExoplanet() ([]Exoplanet, error)
	GetExoplanetById(id int) (Exoplanet, error)
	UpdateExoplanet(id int, p Exoplanet) error
	DeleteExoplanet(id int) error
	FuelEstimation(p Exoplanet, crewSize int) (float64, error)
}
