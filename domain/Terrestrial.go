package domain

/*
Terrestrial :
g = (m/r^2)
*/

type Terre struct {
	Exoplanet
	Type ExoplanetType
}

func NewTerre(p Exoplanet) Terre {
	t := Terre{Exoplanet: p}
	t.Type = Terrestrial
	return t
}

func (t Terre) GravityEstimation() (float64, error) {
	return t.Mass / (t.Radius * t.Radius), nil
}
