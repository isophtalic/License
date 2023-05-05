package dto

type UpdateAndAddCustomerDTO struct {
	Email        string   `json:"email"`
	Name         string   `json:"name"`
	Company      string   `json:"company"`
	PhoneNumber  string   `json:"phone_number"`
	Organization string   `json:"organization"`
	BirthDay     string   `json:"birthday"`
	Address      string   `json:"address"`
	Description  string   `json:"description"`
	Tags         []string `json:"tags"`
	Enable       bool     `json:"enable"`
}

type UpdateCustomerDTO struct {
	Email        string   `json:"email"`
	Name         string   `json:"name"`
	Company      string   `json:"company"`
	Phone        string   `json:"phone"`
	Organization string   `json:"organization"`
	BirthDay     string   `json:"birthday"`
	Address      string   `json:"address"`
	Description  string   `json:"description"`
	Tags         []string `json:"tags"`
	Enable       bool     `json:"enable"`
}
