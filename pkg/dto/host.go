package dto

import (
	"time"
)

type HostOperate struct {
	ID       uint   `gorm:"primary_key"`
	Name     string `json:"name"`
	Addr     string `json:"addr" validate:"required"`
	Port     uint   `json:"port" validate:"required,number,max=65535,min=1"`
	User     string `json:"user" validate:"required"`
	Password string `json:"password"`

	Description string `json:"description"`
}

type HostConnTest struct {
	Addr       string `json:"addr" validate:"required"`
	Port       uint   `json:"port" validate:"required,number,max=65535,min=1"`
	User       string `json:"user" validate:"required"`
	Password   string `json:"password"`
	PrivateKey string `json:"privateKey"`
	PassPhrase string `json:"passPhrase"`
}

type HostAddr struct {
	Addr string `json:"addr" validate:"required"`
}

type SearchHostWithPage struct {
	PageInfo
	GroupID uint   `json:"groupID"`
	Info    string `json:"info"`
}

type SearchForTree struct {
	Info string `json:"info"`
}

type ChangeHostGroup struct {
	ID      uint `json:"id" validate:"required"`
	GroupID uint `json:"groupID" validate:"required"`
}

type HostInfo struct {
	ID               uint      `json:"id"`
	CreatedAt        time.Time `json:"createdAt"`
	GroupID          uint      `json:"groupID"`
	GroupBelong      string    `json:"groupBelong"`
	Name             string    `json:"name"`
	Addr             string    `json:"addr"`
	Port             uint      `json:"port"`
	User             string    `json:"user"`
	AuthMode         string    `json:"authMode"`
	Password         string    `json:"password"`
	PrivateKey       string    `json:"privateKey"`
	PassPhrase       string    `json:"passPhrase"`
	RememberPassword bool      `json:"rememberPassword"`

	Description string `json:"description"`
}

type HostTree struct {
	ID       uint        `json:"id"`
	Label    string      `json:"label"`
	Children []TreeChild `json:"children"`
}

type TreeChild struct {
	ID    uint   `json:"id"`
	Label string `json:"label"`
}
