package response

type AddressResponse struct {
	Line1       string `json:"line1"`
	SubDistrict string `json:"subDistrict"`
	District    string `json:"district"`
	Province    string `json:"province"`
	Postcode    string `json:"postcode"`
}
type UserResponse struct {
	ID        string          `json:"id"`
	FirstName string          `json:"firstName"`
	LastName  string          `json:"lastName"`
	Email     string          `json:"email"`
	Phone     string          `json:"phone"`
	Address   AddressResponse `json:"address"`
}
