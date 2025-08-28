package services

import (
	"context"
	"errors"

	"zatrano/configs/logconfig"
	"zatrano/models"
	"zatrano/repositories"

	"go.uber.org/zap"
)

type IInvitationParticipantService interface {
	GetParticipantsByInvitationID(invitationID uint) ([]models.InvitationParticipant, error)
	CreateOrUpdateParticipant(ctx context.Context, participant *models.InvitationParticipant) error
	DeleteParticipant(ctx context.Context, id uint) error
	GetParticipantCountByInvitation(invitationID uint) (int64, error)
	GetTotalGuestCountByInvitation(invitationID uint) (int, error)
}

type InvitationParticipantService struct {
	repo repositories.IInvitationParticipantRepository
}

func NewInvitationParticipantService() IInvitationParticipantService {
	return &InvitationParticipantService{repo: repositories.NewInvitationParticipantRepository()}
}

func (s *InvitationParticipantService) GetParticipantsByInvitationID(invitationID uint) ([]models.InvitationParticipant, error) {
	participants, err := s.repo.GetParticipantsByInvitationID(invitationID)
	if err != nil {
		logconfig.Log.Error("Davet katılımcıları alınamadı", zap.Uint("invitation_id", invitationID), zap.Error(err))
		return nil, errors.New("davet katılımcıları getirilirken bir hata oluştu")
	}
	return participants, nil
}

func (s *InvitationParticipantService) CreateOrUpdateParticipant(ctx context.Context, participant *models.InvitationParticipant) error {
	// Mevcut katılımcıyı kontrol et
	existingParticipant, err := s.repo.GetParticipantByInvitationAndTelephone(participant.InvitationID, participant.Telephone)

	if err != nil && err.Error() != "record not found" {
		logconfig.Log.Error("Katılımcı kontrolü sırasında hata", zap.Error(err))
		return errors.New("katılımcı kontrolü sırasında bir hata oluştu")
	}

	if existingParticipant != nil {
		// Kayıt varsa güncelle
		existingParticipant.ParticipantCount = participant.ParticipantCount
		err = s.repo.CreateOrUpdateParticipant(ctx, participant)
		if err != nil {
			logconfig.Log.Error("Katılımcı güncellenemedi", zap.Error(err))
			return errors.New("katılımcı güncellenirken bir hata oluştu")
		}
	} else {
		// Kayıt yoksa yeni oluştur
		err = s.repo.CreateOrUpdateParticipant(ctx, participant)
		if err != nil {
			logconfig.Log.Error("Katılımcı oluşturulamadı", zap.Error(err))
			return errors.New("katılımcı oluşturulurken bir hata oluştu")
		}
	}

	return nil
}

func (s *InvitationParticipantService) DeleteParticipant(ctx context.Context, id uint) error {
	err := s.repo.DeleteParticipant(ctx, id)
	if err != nil {
		logconfig.Log.Error("Davet katılımcısı silinemedi", zap.Uint("participant_id", id), zap.Error(err))
		return errors.New("davet katılımcısı silinirken bir hata oluştu")
	}
	return nil
}

func (s *InvitationParticipantService) GetParticipantCountByInvitation(invitationID uint) (int64, error) {
	count, err := s.repo.GetParticipantCountByInvitation(invitationID)
	if err != nil {
		logconfig.Log.Error("Davet katılımcısı sayısı alınamadı", zap.Uint("invitation_id", invitationID), zap.Error(err))
		return 0, errors.New("davet katılımcısı sayısı alınırken bir hata oluştu")
	}
	return count, nil
}

func (s *InvitationParticipantService) GetTotalGuestCountByInvitation(invitationID uint) (int, error) {
	totalGuests, err := s.repo.GetTotalGuestCountByInvitation(invitationID)
	if err != nil {
		logconfig.Log.Error("Toplam misafir sayısı alınamadı", zap.Uint("invitation_id", invitationID), zap.Error(err))
		return 0, errors.New("toplam misafir sayısı alınırken bir hata oluştu")
	}
	return totalGuests, nil
}
