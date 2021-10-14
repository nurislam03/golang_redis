package repos

import (
	"github.com/kamva/mgm/v3"
	"github.com/nurislam03/golang_redis/data/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// LimitStore ...
type LimitStore struct {
}

// NewLimitStore ...
func NewLimitStore() *LimitStore {
	return &LimitStore{}
}

// Create invoice
func (s *LimitStore) Create(u *models.Limit) error {
	return mgm.Coll(u).Create(u)
}

// Update invoice
func (s *LimitStore) Update(u *models.Limit) error {
	return mgm.Coll(u).Update(u)

}

// GetInvoiceByID ...
func (s *LimitStore) GetLimitByID(id string) (*models.Limit, error) {
	limit := &models.Limit{}
	err := mgm.Coll(limit).FindByID(id, limit)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, err
	}
	return limit, nil
}

// GetInvoiceList ...
func (s *LimitStore) GetLimitList(q map[string]interface{}) ([]models.Limit, error) {
	limitList := []models.Limit{}
	fltr := bson.M(q)

	err := mgm.Coll(&models.Limit{}).SimpleFind(&limitList, fltr)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, err
	}

	return limitList, nil
}
