package requests

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type PaymentRequest struct {
	Amount      float64 `form:"amount" validate:"required,gt=0"`
	Installment int     `form:"installment" validate:"required,min=1,max=12"`
	CardNumber  string  `form:"card_number" validate:"required,len=19"`
	CardHolder  string  `form:"card_holder" validate:"required,min=2"`
	ExpiryDate  string  `form:"expiry_date" validate:"required,len=5"`
	CVV         string  `form:"cvv" validate:"required,len=3"`
}

func ParseAndValidatePaymentRequest(c *fiber.Ctx) (PaymentRequest, error) {
	var req PaymentRequest
	if err := c.BodyParser(&req); err != nil {
		return req, errors.New("Geçersiz istek formatı")
	}
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return req, errors.New("Lütfen formdaki hataları düzeltin")
	}
	return req, nil
}
