package data

type AddProductParams struct {
	Title       string  `json:"title" validate:"required"`
	Price       float64 `json:"price" validate:"required,gte=0"`
	Size        string  `json:"size"`
	Measurement string  `json:"measurement"`
	Brand       string  `json:"brand"`
	Quantity    int     `json:"quantity" validate:"required,gte=0"`
}

type QuanityDto struct {
	Units int `json:"units"`
}
