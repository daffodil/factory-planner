
package account


import(
	//"encoding/json"
	//"encoding/json"
)

const (
	Pending = "pending"
	Active = "active"
	InActive = "inactive"
	Error = "error"

)

type Account struct {

	// The primary key
	AccId int ` db:"account_id" json:"account_id"`

	// Given name of the company eg Tesla Mirror Inc
	Company string ` db:"company" json:"company"`

	// The nickname for use inside buisness eg TesMir
	Ticker string

	// The account reference and probably also accounts key eg sage Ref
	AccRef string

	// a list of toop level domains for this client.. ie
	// we dont want to send outside these domain
	// and on input we can sniff out emails from these domains
	// and match to contracts
	// TODO
	Domains []string

	// Flag to indicate account is active
	AccStatus string

	// Flag to indicate account is on hold
	// need this as an alert system
	OnHold int

	// An account has flags for the "type"
	IsClient bool
	IsSupplier bool
	IsSubContracter bool

	// Latest list of notes on this account
	Notes []string


}
