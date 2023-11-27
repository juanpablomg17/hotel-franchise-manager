package model

type Filterrequest struct {
	FranchiseName string `json:"franchise_name"`
	CompanyName   string `json:"company_name"`
	CompanyId     string `json:"company_id"`
	OwnerName     string `json:"owner_name"`
}
