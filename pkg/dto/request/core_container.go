package request

type CoreSecurityRequest struct {
	Ip               string `json:"ip"`
	InsecureRegistry string `json:"insecureRegistry"`
}

type CoreSecurityImageRequest struct {
	InsecureRegistry string `json:"insecureRegistry"`
	Image            string `json:"image"`
}

type CoreImageRequest struct {
	Ip    string `json:"ip"`
	Image string `json:"image"`
}

type CoreContainerRequest struct {
	Ip   string `json:"ip"`
	Name string `json:"name"`
}

type CoreContainerOperationRequest struct {
	Ip        string `json:"ip"`
	Name      string `json:"name"`
	Operation string `json:"operation"`
}
