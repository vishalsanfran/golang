package models

type Cat int

const (
	Romance Cat = iota
	Drama
	Action
	Horror
	Documentary
	Comedy
)

type Video struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Duration uint   `json:"duration"`
	Category string `json:"category"`
	Title    string `json:"title"`
}

type VideoInput struct {
	Duration uint   `json:"duration" binding:"required"`
	Category string `json:"category" binding:"required"`
	Title    string `json:"title" binding:"required"`
}
