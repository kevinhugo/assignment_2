package repository

import (
	"assignment/app/models"
	"assignment/app/resource"
	"assignment/config"
	"fmt"
)

type OrderRepository interface {
	GetOrderList() ([]models.Order, error, int64)
	AddOrder(Order *models.Order, updateData resource.InputOrder) error
	GetOrderDetailById(id uint, preload bool) (models.Order, error)
	DeleteOrder(id int) error
}

func NewOrderRepository() OrderRepository {
	return &dbConnection{
		connection: config.ConnectDB(),
	}
}

// type Order struct {
// 	gorm.Model
// 	OrderID      uint      `json:"order_id" gorm:"primary_key"`
// 	CustomerName string    `json:"customer_name"`
// 	OrderedAt    time.Time `json:"ordered_at" gorm:"autoCreateTime"`
// 	Items        []Item    `gorm:"foreignKey:OrderID"`
// }

// type Item struct {
// 	gorm.Model
// 	ItemID      uint   `json:"item_id" gorm:"primary_key"`
// 	ItemCode    string `json:"item_code" gorm:"index:document_user_id_index,unique"`
// 	Description string `json:"description" gorm:"index:document_user_id_index,unique"`
// 	Quantity    uint   `json:"quantity"`
// 	OrderID     uint   `json:"order_id"`
// 	Order       Order  `gorm:"foreignKey:OrderID"`
// }

func (db *dbConnection) AddOrder(Order *models.Order, createData resource.InputOrder) error {
	// Removing excess data before proceeding
	fmt.Println("====================================")
	if Order.ID != 0 {
		var existingItem []models.Item
		db.connection.Where("order_id = ?", Order.ID).Find(&existingItem)
		var existingItemCount int = len(existingItem)
		var newItemCount int = len(createData.Items)
		if newItemCount < existingItemCount {
			db.connection.Debug().Unscoped().Model(models.Item{}).Where("order_id = ?", Order.ID).
				Order("id asc").
				Limit(existingItemCount - newItemCount).
				Offset(newItemCount).
				Delete(&models.Item{})
		}
	}
	fmt.Println("====================================")

	Order.CustomerName = createData.CustomerName
	err := db.connection.Save(Order).Error
	if err != nil {
		return err
	}

	for eachIndex, eachItem := range createData.Items {
		var item models.Item
		db.connection.Where("order_id = ?", Order.ID).
			Order("id asc").
			Limit(1).Offset(eachIndex).Find(&item)
		item.OrderID = Order.ID
		item.ItemCode = eachItem.ItemCode
		item.Description = eachItem.Description
		item.Quantity = eachItem.Quantity
		err := db.connection.Save(&item).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *dbConnection) GetOrderList() ([]models.Order, error, int64) {
	var Order []models.Order
	var count int64
	connection := db.connection.Model(&Order).Preload("Items").Find(&Order)
	err := connection.Error
	if err != nil {
		return Order, err, 0
	}
	db.connection.Model(Order).Count(&count)
	return Order, nil, count
}

func (db *dbConnection) GetOrderDetailById(id uint, preload bool) (models.Order, error) {
	var Order models.Order
	connection := db.connection
	fmt.Println("OrderId :", id)
	connection = connection.Where("id = ?", id)
	if preload {
		connection = connection.Preload("Items")
	}
	connection = connection.First(&Order)
	err := connection.Error
	if err != nil {
		return Order, err
	}
	return Order, err
}

func (db *dbConnection) DeleteOrder(id int) error {
	var Order models.Order
	err := db.connection.Delete(&Order, id).Error
	if err != nil {
		return err
	}
	return nil
}
