package accounts

type Address struct {



	AddressId int `json:"location_id" gorm:"column:location_id; primary_key:yes"`
	AccountId int `json:"account_id" `
	LocActive bool `json:"loc_active" sql:"type:int(2);not null;default:0" `
	Location string `json:"location" sql:"type:varchar(100);default:''" `
	Address string `json:"address" sql:"type:varchar(100);default:''" `
	Postcode string `json:"postcode" sql:"type:varchar(100);default:''" `
	LocSearch string `json:"loc_search" sql:"type:varchar(100);default:''" `

	Tel string `json:"tel" sql:"type:varchar(100);default:''" `
	Fax string `json:"fax" sql:"type:varchar(100);default:''" `
	IsHq bool `json:"is_hq" sql:"type:int(2);not null;default:0" `
	IsAddress bool `json:"is_address" sql:"type:int(2);not null;default:0" `
	IsBulling bool `json:"is_billing" sql:"type:int(2);not null;default:0" `

}
func (me Address) TableName() string {
	return "addresses"
}

