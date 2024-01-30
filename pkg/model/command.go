package model

type Command struct {
	ID      uint   `gorm:"primary_key"`
	Name    string `gorm:"type:varchar(64);unique;not null" json:"name"`
	Command string `gorm:"type:varchar(256);not null" json:"command"`
}
