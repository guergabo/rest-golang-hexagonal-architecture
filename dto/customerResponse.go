package dto

// data transfer object - maps to response object
// separating it from the domain object means we can
// change how much information the user gets without
// impacting any of the backend code, we just add omit
// here and nowhere else
type CustomerResponse struct {
	Id          string `json:"customer_id"`
	Name        string `json:"full_name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateOfBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}
