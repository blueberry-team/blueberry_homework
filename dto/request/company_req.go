package request

type CreateCompanyRequest struct {
	CompanyName 	string
	CompanyAddress 	string
	TotalStaff 		int
}

type ChangeCompanyRequest struct {
	CompanyName    string
	CompanyAddress string
	TotalStaff 	   int
}
