package utils

type MagicCardRequest struct {
	NamaMolekul  string           `json:"nama_molekul"`
	UnsurMolekul string           `json:"unsur_molekul"`
	Image        string           `json:"image"`
	ImageUrl     string           `json:"image_url"`
	Description  string           `json:"description"`
	ListSenyawa  []SenyawaRequest `json:"list_senyawa"`
}

type SenyawaRequest struct {
	Judul     string `json:"judul"`
	Unsur     string `json:"unsur"`
	Deskripsi string `json:"deskripsi"`
}
