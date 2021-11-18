package dto

type SignupReq struct {
	Email           string                   `json:"email"`
	Password        string                   `json:"password"`
	PhoneNumber     string                   `json:"phone_number"`
	ConfirmPassword string                   `json:"confirm_pass"`
	FullName        string                   `json:"full_name"`
	Gender          string                   `json:"gender"`
	DOB             string                   `json:"date_of_birth"`
	HTMLPages       []UserSignUpReqHTMLPages `json:"html_pages"`
}

type UserSignUpReqHTMLPages struct {
	Version int    `json:"version"`
	Page    string `json:"page"`
}
