package restapi

import "slices"

// Define the Rest API interface. You should not modify the code in this block.

type User struct {
	Name    string
	Owes    map[string]float64
	OwedBy  map[string]float64
	Balance float64
}

type GetUsersRequest struct {
	Users []string
}

type GetUsersResponse struct {
	Users []User
}

type AddUserRequest struct {
	User string
}

type AddUserResponse struct {
	User User
}

type AddIouRequest struct {
	Lender   string
	Borrower string
	Amount   float64
}

type AddIouResponse struct {
	Users []User
}

type RestApi interface {
	GetUsers(GetUsersRequest) GetUsersResponse
	AddUser(AddUserRequest) AddUserResponse
	AddIou(AddIouRequest) AddIouResponse
}

// Your code goes below here. Implement the RestApi interface.

func (u *User) calcBalance() {
	var balance float64 = 0
	for _, v := range u.OwedBy {
		balance += v
	}
	for _, v := range u.Owes {
		balance -= v
	}
	u.Balance = balance
}

func (u User) Compare(other User) int {
	if u.Name < other.Name {
		return -1
	} else if u.Name > other.Name {
		return 1
	} else {
		return 0
	}
}

type Api struct {
	database map[string]User
}

func NewApi(database []User) RestApi {
	usersMap := make(map[string]User)
	for _, u := range database {
		usersMap[u.Name] = u
	}

	return &Api{database: usersMap}
}

func (a *Api) GetUsers(req GetUsersRequest) GetUsersResponse {
	users := make([]User, 0, len(req.Users))
	for _, name := range req.Users {
		users = append(users, a.database[name])
	}
	slices.SortFunc(users, User.Compare)
	return GetUsersResponse{Users: users}
}

func (a *Api) AddUser(req AddUserRequest) AddUserResponse {
	newUser := User{Name: req.User, Owes: make(map[string]float64), OwedBy: make(map[string]float64), Balance: 0}
	a.database[req.User] = newUser
	return AddUserResponse{User: newUser}
}

func (a *Api) AddIou(req AddIouRequest) AddIouResponse {
	lender := a.database[req.Lender]
	borrower := a.database[req.Borrower]
	amount := req.Amount

	// First pay down any debt the lender already has to the borrower.
	if owed := lender.Owes[req.Borrower]; owed > 0 {
		if amount >= owed {
			amount -= owed
			delete(lender.Owes, req.Borrower)
			delete(borrower.OwedBy, req.Lender)
		} else {
			lender.Owes[req.Borrower] = owed - amount
			borrower.OwedBy[req.Lender] = owed - amount
			amount = 0
		}
	}

	// Whatever is left becomes a new loan from lender to borrower.
	if amount > 0 {
		lender.OwedBy[req.Borrower] += amount
		borrower.Owes[req.Lender] += amount
	}

	lender.calcBalance()
	borrower.calcBalance()
	a.database[req.Lender] = lender // write back the updated balances
	a.database[req.Borrower] = borrower

	users := []User{lender, borrower}
	slices.SortFunc(users, User.Compare)
	return AddIouResponse{Users: users}
}
