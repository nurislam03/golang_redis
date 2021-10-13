package limit

import (
	"github.com/nurislam03/template/data/models"
)

// LimitStore represents limit repository interface
type LimitStore interface {
	Create(lmt *models.Limit) error
	Update(lmt *models.Limit) error
	GetLimitByID(id string) (*models.Limit, error)
	GetLimitList(q map[string]interface{}) ([]models.Limit, error)
}
