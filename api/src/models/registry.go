package models

//Registry struct
type Registry struct {
	ID           int     `json:"id,omitempty" db:"id"`
	Name         string  `json:"nombre" db:"name" binding:"required"`
	SecondName   *string `json:"apellido,omitempty" db:"second_name" binding:"required"`
	Email        string  `json:"correo" db:"email" binding:"required"`
	Cel          string  `json:"telefono" db:"cel"`
	TypeCustomer int     `json:"type_customer" db:"type_customer"`
	CreatedAt    string  `json:"created_at,omitempty" db:"created_at"`
}

//Response struct
type Response struct {
	Data Registry `json:"data"`
}

//ResponseRegitries struct
type ResponseRegitries struct {
	Data []Registry `json:"data"`
}
