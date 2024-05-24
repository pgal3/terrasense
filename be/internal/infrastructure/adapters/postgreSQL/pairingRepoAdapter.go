package pg_adapter

import (
	"time"

	"github.com/PaoloEG/terrasense/internal/core/domain/entities"
	"github.com/PaoloEG/terrasense/internal/core/domain/errors"
	vo "github.com/PaoloEG/terrasense/internal/core/domain/value-objects"
	pg_mappers "github.com/PaoloEG/terrasense/internal/infrastructure/adapters/postgreSQL/mappers"
	pg_models "github.com/PaoloEG/terrasense/internal/infrastructure/adapters/postgreSQL/models"
	"gorm.io/gorm"
)

type PairingRepoAdapter struct {
	db *gorm.DB
}

func NewPairingRepoAdapter(dbClient *gorm.DB) *PairingRepoAdapter{
	return &PairingRepoAdapter{
		db: dbClient,
	}
}

func (c *PairingRepoAdapter) PairDevice(pairing entities.Pairing) error {
	pairModel := pg_mappers.ToPairingModel(pairing)
	res := c.db.Create(&pairModel)
	if res.Error != nil {
		return &errors.InternalServerError{
			Message:       "Error in saving pairing in DB",
			OriginalError: res.Error.Error(),
		}
	}
	return nil
}

func (c *PairingRepoAdapter) GetPairings(userID string)([]entities.Pairing, error){
	pairings := []pg_models.Pairing{}
	res := c.db.Order("timestamp DESC").Where("user_id = ?", userID).Last(&pairings)
	if res.Error != nil {
		if res.Error == gorm.ErrRecordNotFound {
			return []entities.Pairing{}, &errors.NotFoundError{
				Message: "No pairings for user "+userID,
			}
		}
		return []entities.Pairing{}, &errors.InternalServerError{
			Message:       "Error in running DB query",
			OriginalError: res.Error.Error(),
		}
	}
	telemetry := pg_mappers.ToPairingEntities(pairings)
	return telemetry, nil
}

func (c *PairingRepoAdapter) DeletePair(userID string, chipID int32) error {
	delete := c.db.Delete(&pg_models.Pairing{}, userID, chipID)
	if delete.Error != nil {
		return &errors.InternalServerError{
			Message:       "Error in running DB delete",
			OriginalError: delete.Error.Error(),
		}
	}
	return nil
}

func (c *PairingRepoAdapter) UpdatePairing(userID string, chipID int32, settings vo.PairingSettings) error {
	pairingModel := pg_models.Pairing{
		UserID: userID,
		ChipID: chipID,
	}
	update := c.db.Model(&pairingModel).Updates(map[string]any{
		"notify_me": settings.NotifyMe(), 
		"low_level_threshold": settings.LowLevelThreshold(), 
		"updated_at": time.Now(),
	})
	if update.Error != nil {
		return &errors.InternalServerError{
			Message:       "Error in running DB update",
			OriginalError: update.Error.Error(),
		}
	}
	return nil
}