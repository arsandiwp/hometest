package productdto

type ProductResponse struct {
	ID            int    `json:"id"`
	Name          string `json:"name" gorm:"unique"`
	SellingPrice  int    `json:"selling_price"`
	PurchasePrice int    `json:"purchase_price"`
	Stock         int    `json:"stock"`
	Image         string `json:"image"`
}
