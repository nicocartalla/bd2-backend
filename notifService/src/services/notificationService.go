package services

import (
	"encoding/json"
	"fmt"
	"notifService/src/database"
	"notifService/src/utils"
	"github.com/streadway/amqp"
	"notifService/src/models"
)

type NotificationService struct{}


func ProcessMessages(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		var notifData models.NotifData
		err := json.Unmarshal(d.Body, &notifData)
		if err != nil {
			utils.ErrorLogger.Println("Failed to unmarshal JSON: %s", err)
			continue
		}
		// Procesar el mensaje (enviar correo y registrar en la tabla de usuarios notificados)
		err = SendEmail(notifData.Data)
		if err != nil {
			utils.ErrorLogger.Println("Failed to send email: %s", err)
			continue
		}
		err = RegisterNotification(notifData.NotifModel)
		if err != nil {
			utils.ErrorLogger.Println("Failed to register notification: %s", err)
			continue
		}
	}
}

func SendEmail(data models.Data) error {

	err := utils.SendEmail(data.Email, data.Subject, data.Body)
	if err != nil {
		utils.ErrorLogger.Println("Error sending email:", err)
		return fmt.Errorf("error sending email: %v", err)
	}
	utils.InfoLogger.Println("Sending email to: ", data.Email)

	return nil
}

func RegisterNotification(notifEvent models.Notification) error {
	query := "INSERT INTO Notifications (document_id, match_id, notification_time, notification_method) VALUES (?, ?, NOW(), ?)"
	_, err := database.InsertDBParams(query, notifEvent.UserId, notifEvent.MatchId, notifEvent.NotificationMethod)

	if err != nil {
		utils.ErrorLogger.Println("Error inserting notification event: ", err)
		return fmt.Errorf("error inserting notification event: %v", err)
	}
	return nil
}