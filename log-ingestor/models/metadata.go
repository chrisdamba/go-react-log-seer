package models

type Metadata struct {
	ID               uint   `json:"id" gorm:"primaryKey"`
	ParentResourceId string `json:"parentResourceId"`
}
