package structs

type DeliveryDetail struct {
	ID int `json:"id"`
	deliveryID string `json:"deliveryid"`
	CustomerName string `json:"cusotmerName" validate:"required"`
	DistanceFromYou float64 `json:"distanceFromYou"`
	ETA float64 `json:"eta"`
	CustomerMob string `json:"cusotmerMob" validate:"required"`
	PicURL string `json:"picurl"`
	Address string `json:"address" validate:"required"`
	Latitude float64 `json:"latitude" validate:"required,gt=0"`
	Longitude float64 `json:"longitude" validate:"required,gt=0"`
	LongLat string `json:"longlat"`
	BusinessId string `json:"businessId" validate:"required"`
}