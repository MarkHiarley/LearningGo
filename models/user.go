package models

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name" binding:"required"`
	Age  int    `json:"age" binding:"required,min=1"`
}
