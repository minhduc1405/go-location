package merchant

type Merchant struct {
	Id   string `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Addr string `json:"addr"`
}
