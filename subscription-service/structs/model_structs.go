package structs

type User struct {
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Tickers   []string `json:"tickers"`
	Email     string   `json:"email"`
	Id        int      `json:"id"`
	Phone     int      `json:"phone"`
}

func NewUser(fname string, lname string, tickers []string, email string, phone int)*User{
	return &User{FirstName: fname, LastName: lname, 
				Tickers: tickers, Email: email, Phone: phone}
}