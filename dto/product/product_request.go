package productdto

type CreateProductRequest struct {
	Name          string `json:"name" form:"name"`
	SellingPrice  int    `json:"selling_price" form:"selling_price"`
	PurchasePrice int    `json:"purchase_price" form:"purchase_price"`
	Stock         int    `json:"stock" form:"stock"`
	Image         string `json:"image" form:"image"`
}

type UpdateProductRequest struct {
	Name          string `json:"name" form:"name"`
	SellingPrice  int    `json:"selling_price" form:"selling_price"`
	PurchasePrice int    `json:"purchase_price" form:"purchase_price"`
	Stock         int    `json:"stock" form:"stock"`
	Image         string `json:"image" form:"image"`
}
