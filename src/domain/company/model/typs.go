package model

type Filterrequest struct {
	CompanyId      string `form:"companyId"`
	FranchiseName  string `json:"franchiseName"`
	CompanyName    string `json:"companyName"`
	OwnerFirstName string `json:"ownerName"`
}
