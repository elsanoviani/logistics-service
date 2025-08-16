package models

import "gorm.io/gorm"

type Shipment struct {
	gorm.Model
	UserID          string `json:"user_id"`
	TrackingNumber  string `json:"tracking_number"`
	SenderName      string `json:"sender_name"`
	SenderAddress   string `json:"sender_address"`
	ReceiverName    string `json:"receiver_name"`
	ReceiverAddress string `json:"receiver_address"`
	ItemDescription string `json:"item_description"`
	Status          string `json:"status"`
}
