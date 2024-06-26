package domain

type Gas struct {
	Exoplanet
	Type ExoplanetType
}

/*
Gas Giant :
g = (0.5/r^2)
Terrestrial :
g = (m/r^2)
*/
func NewGas(p Exoplanet) Gas {
	gas := Gas{Exoplanet: p}
	gas.Type = GasGiant
	return gas
}

func (g Gas) GravityEstimation() (float64, error) {
	return 0.5 / (float64(g.Radius) * float64(g.Radius)), nil
}
