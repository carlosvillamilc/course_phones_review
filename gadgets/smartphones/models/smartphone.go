package models

// Smartphone model structure for smartphones
type Smartphone struct{
	Id 				int64
	Name 			string
	Price 			int
	CountryOrigin 	string
	OperativeSystem string 
}

// CreateSmartphoneCMD
type CreateSmartphoneCMD struct {
	Name			string 	`json:"name"`
	Price		 	int 	`json:"price"`
	CountryOrigin 	string 	`json:"country_origin"`
	OperativeSystem string 	`json:"operative_system"`

}