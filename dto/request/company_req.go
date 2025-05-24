package request

type CreateCompanyRequest struct {
	UserID      	string
	CompanyName 	string
	CompanyAddress 	string
	TotalStaff 		int
}

type ChangeCompanyRequest struct {
	UserId         string
	CompanyName    string
	CompanyAddress string
	TotalStaff 	   int
}
