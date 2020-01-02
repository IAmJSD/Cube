package currency

// BuyableRole defines a role which you can buy.
type BuyableRole struct {
	Amount       int
	RoleID       string
	Description  string
	TrialAllowed bool
}

// Currency is the struct which all guild currencies will be based on.
type Currency struct {
	Emoji        *string
	Enabled      bool
	DropsEnabled bool
	DropsImage   *string
	DropsAmount  *int
	RoleShop     []*BuyableRole
}
