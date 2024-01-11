package banner

import "gorm.io/gorm"

func BannerRegistry(db *gorm.DB) Service {
	bannerRepository := NewRepository(db)
	bannerService := NewService(bannerRepository)

	return bannerService
}
