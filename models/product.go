package models

type Product struct {
	ID            int    `json:"id"`
	Name          string `json:"name" gorm:"unique"`
	SellingPrice  int    `json:"selling_price" gorm:"type: int"`
	PurchasePrice int    `json:"purchase_price" gorm:"type: int"`
	Stock         int    `json:"stock" gorm:"type: int"`
	Image         string `json:"image" gorm:"type: varchar(255)"`
}
