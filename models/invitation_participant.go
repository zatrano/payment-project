package models

type InvitationParticipant struct {
	BaseModel

	// Zorunlu Alanlar
	InvitationID     uint   `gorm:"index;not null"`
	Name             string `gorm:"type:varchar(100);not null"`
	Telephone        string `gorm:"type:varchar(20);not null"`
	ParticipantCount int    `gorm:"not null;default:1"`

	// İlişki Tanımı
	Invitation Invitation `gorm:"foreignKey:InvitationID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (InvitationParticipant) TableName() string {
	return "invitation_participants"
}
