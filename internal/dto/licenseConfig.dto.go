package dto

type UpdateAndCreateLicenseConfigDTO struct {
	Name   string `json:"name"`
	Key    string `json:"key"`
	Value  string `json:"value"`
	Status bool   `json:"status"`
}
