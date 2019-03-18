package models

//Registry struct
type Registry struct {
	ID           int    `json:"id,omitempty" db:"id"`
	Name         string `json:"name" db:"name" binding:"required"`
	Email        string `json:"email" db:"email" binding:"required"`
	Cel          string `json:"cel" db:"cel"`
	TypeCustomer int    `json:"type_customer" db:"type_customer" binding:"required"`
	CreatedAt    string `json:"created_at" db:"created_at"`
}
