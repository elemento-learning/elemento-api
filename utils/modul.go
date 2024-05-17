package utils

type ModulRequest struct {
	Title      string
	IsComplete bool
	Subtitle   string
	Image      string
	ImageUrl   string
	Babs       []BabRequest
}

type BabRequest struct {
	Title         string
	Description   string
	Task          string
	ResultStudent string
}
