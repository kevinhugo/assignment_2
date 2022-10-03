package models

import (
	"time"

	"gorm.io/gorm"
)

// type Document struct {
// 	gorm.Model
// 	ID               uint   `json:"id" gorm:"primary_key"`
// 	NotificationID       uint   `json:"NotificationID"  binding:"required" gorm:"index:,unique,composite:orderNotification"`
// 	OrderID          string `json:"OrderID" binding:"required" gorm:"index:,unique,composite:orderNotification"`
// 	Notification         Notification
// 	RefferenceNumber string             `json:"refferenceNumber"`
// 	TotalAmount      int                `json:"totalAmount" binding:"required"`
// 	Items            []DocumentUser `gorm:"foreignKey:DocumentID"`
// 	Status           string             `json:"status" gorm:"default:pending"`
// 	SelectedPayment  string             `json:"selectedPayment"`
// 	Token            string             `json:"token"`
// 	TokenAuth        string             `json:"tokenAuth"`
// }

// type DocumentUser struct {
// 	gorm.Model
// 	DocumentID uint
// 	Name          string `json:"name"`
// 	Price         int    `json:"price"`
// }

type Order struct {
	gorm.Model
	ID           uint      `json:"order_id" gorm:"primary_key"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at" gorm:"autoCreateTime"`
	Items        []Item    `gorm:"foreignKey:OrderID"`
}

type Item struct {
	gorm.Model
	ID          *uint  `json:"item_id" gorm:"primary_key"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"order_id"`
	Order       Order  `gorm:"foreignKey:OrderID"`
}
