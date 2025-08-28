package handlers

import (
	"net/http"
	"strconv"

	"zatrano/models"
	"zatrano/pkg/flashmessages"
	"zatrano/pkg/renderer"
	"zatrano/services"

	"github.com/gofiber/fiber/v2"
)

type InvitationParticipantHandler struct {
	participantService services.IInvitationParticipantService
	invitationService  services.IInvitationService // Davet bilgileri için
}

func NewInvitationParticipantHandler() *InvitationParticipantHandler {
	return &InvitationParticipantHandler{
		participantService: services.NewInvitationParticipantService(),
		invitationService:  services.NewInvitationService(), // Davet service'i
	}
}

func (h *InvitationParticipantHandler) ListParticipants(c *fiber.Ctx) error {
	invitationID, err := c.ParamsInt("invitationId")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Geçersiz davet ID")
	}

	// Davet bilgilerini al
	invitation, err := h.invitationService.GetInvitationByID(uint(invitationID))
	if err != nil {
		_ = flashmessages.SetFlashMessage(c, flashmessages.FlashErrorKey, "Davet bulunamadı")
		return c.Redirect("/dashboard/invitations", fiber.StatusSeeOther)
	}

	// Katılımcıları al
	participants, err := h.participantService.GetParticipantsByInvitationID(uint(invitationID))
	if err != nil {
		_ = flashmessages.SetFlashMessage(c, flashmessages.FlashErrorKey, err.Error())
	}

	// İstatistikleri al
	participantCount, _ := h.participantService.GetParticipantCountByInvitation(uint(invitationID))
	totalGuests, _ := h.participantService.GetTotalGuestCountByInvitation(uint(invitationID))

	return renderer.Render(c, "dashboard/invitations/participants", "layouts/dashboard", fiber.Map{
		"Title":            "Davet Katılımcıları - " + invitation.Category.Name,
		"Participants":     participants,
		"Invitation":       invitation, // Tüm davet bilgileri
		"InvitationID":     invitationID,
		"ParticipantCount": participantCount,
		"TotalGuests":      totalGuests,
	}, http.StatusOK)
}

func (h *InvitationParticipantHandler) ShowCreateForm(c *fiber.Ctx) error {
	invitationID, err := c.ParamsInt("invitationId")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Geçersiz davet ID")
	}

	// Davet bilgilerini al
	invitation, err := h.invitationService.GetInvitationByID(uint(invitationID))
	if err != nil {
		_ = flashmessages.SetFlashMessage(c, flashmessages.FlashErrorKey, "Davet bulunamadı")
		return c.Redirect("/dashboard/invitations", fiber.StatusSeeOther)
	}

	return renderer.Render(c, "dashboard/invitation-participants/create", "layouts/dashboard", fiber.Map{
		"Title":        "Yeni Katılımcı Ekle - " + invitation.Category.Name,
		"Invitation":   invitation,
		"InvitationID": invitationID,
	})
}

func (h *InvitationParticipantHandler) CreateParticipant(c *fiber.Ctx) error {
	invitationID, err := c.ParamsInt("invitationId")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Geçersiz davet ID")
	}

	name := c.FormValue("name")
	telephone := c.FormValue("telephone")
	participantCount, _ := strconv.Atoi(c.FormValue("participant_count"))

	if participantCount < 1 {
		participantCount = 1
	}

	participant := &models.InvitationParticipant{
		InvitationID:     uint(invitationID),
		Name:             name,
		Telephone:        telephone,
		ParticipantCount: participantCount,
	}

	if err := h.participantService.CreateOrUpdateParticipant(c.UserContext(), participant); err != nil {
		_ = flashmessages.SetFlashMessage(c, flashmessages.FlashErrorKey, err.Error())
		return c.Redirect("/dashboard/invitations/"+strconv.Itoa(invitationID)+"/participants/create", fiber.StatusSeeOther)
	}

	_ = flashmessages.SetFlashMessage(c, flashmessages.FlashSuccessKey, "Katılımcı başarıyla eklendi/güncellendi.")
	return c.Redirect("/dashboard/invitations/"+strconv.Itoa(invitationID)+"/participants", fiber.StatusFound)
}

func (h *InvitationParticipantHandler) DeleteParticipant(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		if c.Get("Accept") == "application/json" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Geçersiz ID"})
		}
		return c.Status(fiber.StatusBadRequest).SendString("Geçersiz ID")
	}

	if err := h.participantService.DeleteParticipant(c.UserContext(), uint(id)); err != nil {
		if c.Get("Accept") == "application/json" {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		_ = flashmessages.SetFlashMessage(c, flashmessages.FlashErrorKey, err.Error())
	} else {
		if c.Get("Accept") == "application/json" {
			return c.JSON(fiber.Map{"success": true})
		}
		_ = flashmessages.SetFlashMessage(c, flashmessages.FlashSuccessKey, "Katılımcı başarıyla silindi.")
	}

	return c.Redirect(c.Get("Referer"), fiber.StatusFound)
}
