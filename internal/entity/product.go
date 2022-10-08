package entity

type Product struct {
	Name        string `json:"name"`
	Price       int64  `json:"price"`
	Quantity    int    `json:"quantity"`
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
	Sku         string `json:"sku"`
}
