package dto

type CommandOperate struct {
	ID      uint   `gorm:"primary_key"`
	Name    string `json:"name" validate:"required"`
	Command string `json:"command" validate:"required"`
}

type CommandInfo struct {
	Name    string `json:"name"`
	Command string `json:"command"`
}
