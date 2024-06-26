# Space Voyagers
Space Voyagers is a Golang microservice for supporting the space voyagers who are embarking on a journey to study different exoplanets.

## Installation

Install this using docker


## Usage

Enum used for planet types.

```golang
const (
	GasGiant ExoplanetType = iota + 1
	Terrestrial
)

```

## APIs
```routes
GET|PUT|POST|DELETE :   /Exoplanet
GET: /Exoplanet/Fuel
```

