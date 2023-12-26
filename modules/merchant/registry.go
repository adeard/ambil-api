package merchant

import "gorm.io/gorm"

func MerchantRegistry(db *gorm.DB) Service {
	merchantRepository := NewRepository(db)
	merchantService := NewService(merchantRepository)

	return merchantService
}
