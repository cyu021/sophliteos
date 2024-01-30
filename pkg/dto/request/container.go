package request

type SecurityRequest struct {
	InsecureRegistry string `json:"insecureRegistry"`
}

type SecurityImageRequest struct {
	InsecureRegistry string `json:"insecureRegistry"`
	Image            string `json:"image"`
}

type ImageRequest struct {
	Image string `json:"image"`
}

type ContainerRequest struct {
	Name string `json:"name"`
}
