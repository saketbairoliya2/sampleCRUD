package dao

// ResponseStruct is generic response struct for any given response
type ResponseStruct struct {
	Result  interface{} `json:"result"`
	Message string      `json:"message,omitempty"`
	Success bool        `json:"success"`
}

type RegistrationForm struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
}
