package dto

type CompanyDTO struct {
	Company struct {
		Owner struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Contact   struct {
				Email    string `json:"email"`
				Phone    string `json:"phone"`
				Location struct {
					City    string `json:"city"`
					Country string `json:"country"`
					Address string `json:"address"`
					ZipCode string `json:"zip_code"`
				} `json:"location"`
			} `json:"contact"`
		} `json:"owner"`
		Information struct {
			Name      string `json:"name"`
			TaxNumber string `json:"tax_number"`
			Location  struct {
				City    string `json:"city"`
				Country string `json:"country"`
				Address string `json:"address"`
				ZipCode string `json:"zip_code"`
			} `json:"location"`
		} `json:"informacion"`
		Franchises []struct {
			Name     string `json:"name"`
			URL      string `json:"url"`
			Location struct {
				City    string `json:"city"`
				Country string `json:"Country"`
				Address string `json:"Address"`
				ZipCode string `json:"zip_code"`
			} `json:"location"`
		} `json:"franchises"`
	} `json:"company"`
}
