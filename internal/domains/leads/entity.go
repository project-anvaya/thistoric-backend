package leads

type Lead struct {
	ID             uint   `json:"id" gorm:"primaryKey"`
	VendorID       uint   `json:"vendor_id" gorm:"index"`
	OrganizerName  string `json:"organizer_name"`
	OrganizerEmail string `json:"organizer_email"`
	Phone          string `json:"phone"`
	Message        string `json:"message"`
	CreatedAt      int64  `json:"created_at" gorm:"autoCreateTime"`
}
