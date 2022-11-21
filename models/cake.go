package models

type Cake struct {
	ID     int    `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	Desc   string `json:"description" db:"description"`
	Rating int    `json:"rating" db:"rating"`
	// IsActive  bool   `json:"is_active" db:"is_active"`
	Image     string `json:"image" db:"image"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"Updated_at" db:"Updated_at"`
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
