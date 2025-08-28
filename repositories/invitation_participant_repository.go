package repositories

import (
	"context"

	"zatrano/configs/databaseconfig"
	"zatrano/models"

	"gorm.io/gorm"
)

type IInvitationParticipantRepository interface {
	GetParticipantsByInvitationID(invitationID uint) ([]models.InvitationParticipant, error)
	GetParticipantByInvitationAndTelephone(invitationID uint, telephone string) (*models.InvitationParticipant, error)
	CreateOrUpdateParticipant(ctx context.Context, participant *models.InvitationParticipant) error
	DeleteParticipant(ctx context.Context, id uint) error
	GetParticipantCountByInvitation(invitationID uint) (int64, error)
	GetTotalGuestCountByInvitation(invitationID uint) (int, error)
}

type InvitationParticipantRepository struct {
	db *gorm.DB
}

func NewInvitationParticipantRepository() IInvitationParticipantRepository {
	return &InvitationParticipantRepository{db: databaseconfig.GetDB()}
}

func (r *InvitationParticipantRepository) GetParticipantsByInvitationID(invitationID uint) ([]models.InvitationParticipant, error) {
	var participants []models.InvitationParticipant
	err := r.db.Where("invitation_id = ?", invitationID).Find(&participants).Error
	return participants, err
}

func (r *InvitationParticipantRepository) GetParticipantByInvitationAndTelephone(invitationID uint, telephone string) (*models.InvitationParticipant, error) {
	var participant models.InvitationParticipant
	err := r.db.Where("invitation_id = ? AND telephone = ?", invitationID, telephone).First(&participant).Error
	if err != nil {
		return nil, err
	}
	return &participant, nil
}

func (r *InvitationParticipantRepository) CreateOrUpdateParticipant(ctx context.Context, participant *models.InvitationParticipant) error {
	// Önce aynı telefon numarası ile kayıt var mı kontrol et
	existingParticipant, err := r.GetParticipantByInvitationAndTelephone(participant.InvitationID, participant.Telephone)

	if err == gorm.ErrRecordNotFound {
		// Kayıt yok, yeni oluştur
		return r.db.WithContext(ctx).Create(participant).Error
	} else if err != nil {
		return err
	}

	// Kayıt var, güncelle
	existingParticipant.ParticipantCount = participant.ParticipantCount
	return r.db.WithContext(ctx).Save(existingParticipant).Error
}

func (r *InvitationParticipantRepository) DeleteParticipant(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.InvitationParticipant{}, id).Error
}

func (r *InvitationParticipantRepository) GetParticipantCountByInvitation(invitationID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.InvitationParticipant{}).Where("invitation_id = ?", invitationID).Count(&count).Error
	return count, err
}

func (r *InvitationParticipantRepository) GetTotalGuestCountByInvitation(invitationID uint) (int, error) {
	var totalGuestCount int
	err := r.db.Model(&models.InvitationParticipant{}).
		Where("invitation_id = ?", invitationID).
		Select("COALESCE(SUM(participant_count), 0)").
		Scan(&totalGuestCount).Error
	return totalGuestCount, err
}
