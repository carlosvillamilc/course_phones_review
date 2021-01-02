package models

// Buyer model structure
type Buyer struct {
	Id   string
	Name string
	Age  int
}

// CreateBuyerCMD
type CreateBuyerCMD struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//Product model structure
type Product struct {
	Id    string
	Name  string
	Price int
}
