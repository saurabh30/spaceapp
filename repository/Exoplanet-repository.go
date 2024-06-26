package repository

import (
	"fmt"
	"log"
	"spaceapp/domain"
)

type ExoplanetRepository struct {
	Database    map[int]domain.Exoplanet
	ExoplanetId int
}

func NewExoplanetRepository() domain.ExoplanetRepository {
	log.Println("Data Layer Initiated...")
	return &ExoplanetRepository{
		Database:    map[int]domain.Exoplanet{},
		ExoplanetId: 0,
	}
}

func (r *ExoplanetRepository) AddExoplanet(p domain.Exoplanet) int {
	pId := r.ExoplanetId + 1
	r.Database[pId] = p
	r.ExoplanetId = pId
	return pId
}

func (r *ExoplanetRepository) ListExoplanet() ([]domain.Exoplanet, error) {
	list := []domain.Exoplanet{}
	for _, rows := range r.Database {
		list = append(list, rows)
	}
	return list, nil
}

func (r *ExoplanetRepository) GetExoplanetById(id int) (domain.Exoplanet, error) {
	if p, ok := r.Database[id]; ok {
		return p, nil
	}
	return domain.Exoplanet{}, fmt.Errorf("invalid id %d", id)
}

func (r *ExoplanetRepository) UpdateExoplanet(id int, Exoplanet domain.Exoplanet) error {
	if _, ok := r.Database[id]; ok {
		r.Database[id] = Exoplanet
		return nil
	}
	return fmt.Errorf("invalid id %d", id)
}

func (r *ExoplanetRepository) DeleteExoplanet(id int) error {
	if _, ok := r.Database[id]; ok {
		delete(r.Database, id)
		return nil
	}
	return fmt.Errorf("invalid id %d", id)
}
