package utils

type ModulRequest struct {
	Title       string
	IsComplete  bool
	Subtitle    string
	Image       string
	YoutubeLink string
	ImageUrl    string
	Babs        []BabRequest
}

type BabRequest struct {
	Title       string
	Description string
	Task        string
}
