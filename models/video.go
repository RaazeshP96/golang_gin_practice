package models

type Person struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       int8   `json:"age" binding:"gte=1,lte=130"`
	Email     string `json:"email" validate:"requird,email"`
}

type Video struct {
	Title       string `json:"title" form:"title" binding:"min=2 ,max=10"`
	Description string `json:"description form:"title" binding:"min=2, max=20`
	URL         string `json:"url" binding:"required,url"`
	Author      string `json:"author" binding:"required"`
}
