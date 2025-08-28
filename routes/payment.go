package routes

import (
	handlers "zatrano/handlers/payment"

	"github.com/gofiber/fiber/v2"
)

func registerPaymentRoutes(app *fiber.App) {
	paymentGroup := app.Group("/payment")

	paymentHandler := handlers.NewPaymentHandler()

	// Yeni ödeme formu
	paymentGroup.Get("/create", paymentHandler.ShowPaymentForm)
	paymentGroup.Post("/create", paymentHandler.SubmitPayment)

	// Başarı/başarısız sayfaları
	paymentGroup.Get("/success", paymentHandler.ShowSuccess)
	paymentGroup.Get("/error", paymentHandler.ShowError)
}
