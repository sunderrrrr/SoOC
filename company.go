package sooc

type Company struct {
	CompanyId   int    `json:"comp_id"`
	CompanyName string `json:"comp_name"`
	CompAdrr    string `json:"comp_addr"`
	LogoUrl     string `json:"logo_photo"`
}
