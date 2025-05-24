package request

type CreateCompanyRequest struct {
	UserID      	string
	CompanyName 	string
	CompanyAddress 	string
	TotalStaff 		int
}

type GetCompanyRequest struct {
	UserId string
}

type ChangeCompanyRequest struct {
	UserId         string
	CompanyName    string
	CompanyAddress string
	TotalStaff 	   int
}

type DeleteCompanyRequest struct {
	UserId string
}
