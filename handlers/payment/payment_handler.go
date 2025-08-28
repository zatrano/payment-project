package handlers

import (
	"net/http"
	"zatrano/models"
	"zatrano/pkg/renderer"
	"zatrano/requests"
	"zatrano/services"

	"github.com/gofiber/fiber/v2"
)

type PaymentHandler struct {
	paymentService services.IPaymentService
}

func NewPaymentHandler() *PaymentHandler {
	return &PaymentHandler{
		paymentService: services.NewPaymentService(),
	}
}

func (h *PaymentHandler) ShowPaymentForm(c *fiber.Ctx) error {
	return renderer.Render(c, "payments/form", "layouts/payment", fiber.Map{"Title": "Ödeme Yap"}, http.StatusOK)
}

func (h *PaymentHandler) SubmitPayment(c *fiber.Ctx) error {
	req, err := requests.ParseAndValidatePaymentRequest(c)
	if err != nil {
		return renderer.Render(c, "payments/form", "layouts/payment", fiber.Map{
			"Title": "Ödeme Yap",
			"Error": err.Error(),
			"Form":  req,
		}, http.StatusBadRequest)
	}
	payment := &models.Payment{
		Amount:      req.Amount,
		Installment: req.Installment,
		CardNumber:  req.CardNumber,
		CardHolder:  req.CardHolder,
		ExpiryDate:  req.ExpiryDate,
		CVV:         req.CVV,
		Status:      "pending",
	}
	err = h.paymentService.CreatePayment(c.UserContext(), payment)
	if err != nil {
		return c.Redirect("/payment/error")
	}
	return c.Redirect("/payment/success")
}

func (h *PaymentHandler) ShowSuccess(c *fiber.Ctx) error {
	return renderer.Render(c, "payments/success", "layouts/payment", fiber.Map{"Title": "Ödeme Başarılı"}, http.StatusOK)
}

func (h *PaymentHandler) ShowError(c *fiber.Ctx) error {
	return renderer.Render(c, "payments/error", "layouts/payment", fiber.Map{"Title": "Ödeme Başarısız"}, http.StatusOK)
}
