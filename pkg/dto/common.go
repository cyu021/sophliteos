package dto

type BuildInfo struct {
	Modules   []Module `json:"modules"`
	BuildName string   `json:"buildname"`
	BuildTime string   `json:"buildtime"`
}

type Module struct {
	Name   string `json:"module"`
	Commit string `json:"commit"`
}
