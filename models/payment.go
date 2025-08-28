package models

type Payment struct {
	BaseModel
	Amount      float64 `gorm:"not null"`
	Installment int     `gorm:"not null"`
	CardNumber  string  `gorm:"size:19;not null"`
	CardHolder  string  `gorm:"size:255;not null"`
	ExpiryDate  string  `gorm:"size:5;not null"`
	CVV         string  `gorm:"size:3;not null"`
	Status      string  `gorm:"size:20;not null"`
}

func (Payment) TableName() string {
	return "payments"
}
