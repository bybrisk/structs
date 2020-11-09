package structs

import "go.mongodb.org/mongo-driver/bson/primitive"

type ConnectToDataBase struct {
	CustomApplyURI string 
	DatabaseName string 
}

type DeliveryDetail struct {
	ID int `json:"id"`
	deliveryID string `json:"deliveryid"`
	CustomerName string `json:"cusotmerName"`
	DistanceFromYou float64 `json:"distanceFromYou"`
	ETA float64 `json:"eta"`
	CustomerMob string `json:"cusotmerMob"`
	PicURL string `json:"picurl"`
	Address string `json:"address"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	LongLat string `json:"longlat"`
	BusinessId string `json:"businessId"`
}

type BusinessAccount struct{
	ID    primitive.ObjectID `json:"-"`
	PicURL string `json: "picurl"`
	UserName string `json: "username"`
	BusinessName string `json: "businessname"`
	Password string `json: "password"`
	City string `json: "city"`
	BusinessPlan string `json: "businessplan"`
	ProfileConfig ProfileConfig `json: "profileConfig"`
	DeliveryPending string `json: "deliveryPending"`
	DeliveryDelivered string `json: "deliveryDelivered"`
	UserID string `json:"userId"`
	ZoneDetailInfo []ZoneInfo `json:"zoneDetailInfo"`
}

type ZoneInfo struct {
	ID    primitive.ObjectID `json:"-"`
	PicURL string `json: "picurl"`
	Name string `json: "name"`
	BusinessUID string `json: "businessUid"`
	DeliveryInZone string `json: "deliveryInZone"`
	UserID string `json:"userId"`
	DeliveryDetail []DeliveryDetail `json: "deliveryDetail"`
	Longitude string `json: "longitude"`
	Latitude string `json: "latitude"`
	Error string `json:"error"`
}

type ProfileConfig struct{
	Zone int64 `json: "zone"`
	MessagePlan int64 `json: "messageplan"`
	Tracking bool `json: "tracking"`
	ZoneID []string `json: "zoneid"`
}

type MapBoxResp struct {
	Code   string `json:"code"`
	Routes []RoutesResp `json:"routes"`
}

type RoutesResp struct {
	Distance float64 `json:"distance"`
	Duration float64 `json:"duration"`
}

type DeliveredDetail struct {
	Longitude string `json:"longitude"`
	Latitude string `json:"latitude"`
	CustomerName string `json:"customerName"`
	CustomerMob string `json:"customerMob"`
	ZoneID string `json:"zoneId"`
	DateOfDelivery string `json:"dateOfDelivery"`
	Address string `json:"address"`
	BusinessUid string `json:"businessUid"`
}

type DeliveredAndAccount struct {
	DeliveredDetails []DeliveredDetail `json:"deliveryDetails"`
	BusinessAccount *BusinessAccount `json:"businessAccount"`
}