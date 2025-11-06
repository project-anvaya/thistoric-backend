package vendors

type Vendor struct {
	ID            uint    `json:"id" gorm:"primaryKey"`
	OwnerUserID   *string `json:"owner_user_id" gorm:"type:uuid"`
	Name          string  `json:"name" gorm:"index"`
	Category      string  `json:"category" gorm:"index"`
	Location      string  `json:"location" gorm:"index"`
	Description   string  `json:"description"`
	Phone         string  `json:"phone"`
	Website       string  `json:"website"`
	PortfolioKeys string  `json:"-" gorm:"type:text"` // future: R2 keys
}
