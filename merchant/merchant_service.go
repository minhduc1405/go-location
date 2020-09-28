package merchant

type MerchantService interface {
	Create(model *Merchant) (*Merchant, error)
	Update(model *Merchant) error
	Delete(id string) error
	All() ([]Merchant, error)
	Load(id string) (*Merchant, error)
}
