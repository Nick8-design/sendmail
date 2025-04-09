package main

import (
	"log"
	"net/smtp"

	"github.com/gofiber/fiber/v2"
)

type EmailRequest struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}


const (
	smtpHost     = "smtp.gmail.com"
	smtpPort     = "587"
	senderEmail  = "nickeagle888@gmail.com"
	senderPass   = "aoqr pvsd ynkk phwi" 
)

func sendEmail(to, subject, body string) error {
	auth := smtp.PlainAuth("", senderEmail, senderPass, smtpHost)

	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n" +
		body + "\r\n")

	return smtp.SendMail(smtpHost+":"+smtpPort, auth, senderEmail, []string{to}, msg)
}

func main() {
	app := fiber.New()

	app.Post("/send-email", func(c *fiber.Ctx) error {
		req := new(EmailRequest)
		if err := c.BodyParser(req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request",
			})
		}

		err := sendEmail(req.To, req.Subject, req.Body)
		if err != nil {
			log.Println("Email error:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to send email",
			})
		}

		return c.JSON(fiber.Map{
			"message": "Email sent successfully",
		})
	})


	log.Fatal(app.Listen(":8080"))
}


