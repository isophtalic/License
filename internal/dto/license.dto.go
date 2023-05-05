package dto

type CreateLicenseDTO struct {
	Name         string   `json:"name"`
	Option       string   `json:"option"`
	Description  string   `json:"description"`
	ExpiryDate   string   `json:"expiry_date"`
	Tag          []string `json:"tags"`
	CustomerName string   `json:"customer_name"`
	ProductName  string   `json:"product_name"`
}
type UpdateLicenseDTO struct {
	Name        string   `json:"name"`
	Option      string   `json:"option"`
	Description string   `json:"description"`
	ExpiryDate  string   `json:"expiry_date"`
	Tag         []string `json:"tags"`
}
