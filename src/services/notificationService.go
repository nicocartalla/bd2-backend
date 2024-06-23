package services

import (
	"bd2-backend/src/database"
	"bd2-backend/src/models"
	"bd2-backend/src/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"text/template"
	"time"
	"bd2-backend/src/templates"
)

type NotificationService struct{}

func SendEmailToUsersWhoHaveNotMadePredictions() error {

	tmpl, err := template.New("").Parse(templates.EmailInsertPrediction)
 
	if err != nil {
		utils.ErrorLogger.Println("Error parsing template", err)
		return fmt.Errorf("error parsing template: %v", err)
	}
	usersTemplate, err := getUsersWhoHaveNotMadePredictions(8)
	if err != nil {
		return err
	}

	for _, user := range usersTemplate {

		type Data struct {
			Email   string `json:"email"`
			Subject string `json:"subject"`
			Body    string `json:"body"`
		}

		type NotifData struct {
			Data       Data
			NotifModel models.Notification
		}

		var body bytes.Buffer
		err = tmpl.Execute(&body, user)
		if err != nil {
			utils.ErrorLogger.Println("Error executing template", err)
			return fmt.Errorf("error executing template: %v", err)
		}

		data := Data{
			Email:   user.UserEmail,
			Subject: "No te olvides de hacer tu Predicción!",
			Body:    body.String(),
		}

		notifData := NotifData{
			Data:       data,
			NotifModel: models.Notification{UserId: user.UserId, MatchId: user.MatchId, NotificationMethod: "email"},
		}

		jsonData, err := json.Marshal(notifData)
		if err != nil {
			utils.ErrorLogger.Println("Error marshalling data", err)
			return fmt.Errorf("error marshalling data: %v", err)
		}

		body = *bytes.NewBuffer(jsonData)

		errToQueue := utils.AddmessageToQueue(body.Bytes())

		if errToQueue != nil {
			return fmt.Errorf("error adding message to queue: %v", errToQueue)
		}
		utils.InfoLogger.Println("Successfully added message to queue for user: ", user.UserEmail)

	}
	return nil
}

func getUsersWhoHaveNotMadePredictions(nextHours int) ([]models.UsersWhoHaveNotMadePredictionsTemplate, error) {
	query := `
		SELECT u.document_id, gm.match_id, u.email, u.first_name, u.last_name, c.name AS championship_name, tl.name AS local_team, tv.name AS visitor_team, gm.match_date, tl.url_logo AS local_team_logo, tv.url_logo AS visitor_team_logo
		FROM User u
		LEFT JOIN Predictions p ON u.document_id = p.document_id AND p.match_id IN (
			SELECT gm.match_id
			FROM GameMatch gm
			WHERE gm.match_date > NOW() AND gm.match_date <= DATE_ADD(NOW(), INTERVAL ? HOUR)
		)
		JOIN GameMatch gm ON gm.match_id NOT IN (SELECT match_id FROM Predictions WHERE document_id = u.document_id)
		JOIN Championships c ON gm.championship_id = c.championship_id
		JOIN Teams tl ON gm.team_local_id = tl.team_id
		JOIN Teams tv ON gm.team_visitor_id = tv.team_id
		LEFT JOIN Notifications n ON u.document_id = n.document_id AND gm.match_id = n.match_id
		WHERE gm.match_date > NOW() AND gm.match_date <= DATE_ADD(NOW(), INTERVAL ? HOUR)
		AND p.prediction_id IS NULL
		AND n.notification_id IS NULL
		ORDER BY gm.match_date;
	`
	rows, err := database.QueryRowsDBParams(query, nextHours, nextHours)
	if err != nil {
		utils.ErrorLogger.Println("Error querying users to remind prediction:", err)
		return nil, fmt.Errorf("error querying users to remind prediction: %v", err)
	}
	var usersTemplate []models.UsersWhoHaveNotMadePredictionsTemplate
	for rows.Next() {
		var template models.UsersWhoHaveNotMadePredictionsTemplate
		if err := rows.Scan(&template.UserId, &template.MatchId, &template.UserEmail, &template.UserFirstName, &template.UserLastName, &template.ChampionshipName, &template.LocalTeam, &template.VisitorTeam, &template.MatchDate, &template.LocalTeamLogo, &template.VisitorTeamLogo); err != nil {
			utils.ErrorLogger.Println("Error scanning user to remind prediction:", err)
			return nil, fmt.Errorf("error scanning user to remind prediction: %v", err)
		}

		template.MatchDate = parseTime(template.MatchDate)

		usersTemplate = append(usersTemplate, template)
	}
	return usersTemplate, nil
}

func parseTime(dateTimeStr string) string {
	parsedTime, err := time.Parse(time.RFC3339, dateTimeStr)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return ""
	}
	return parsedTime.Format("15:04")
}

func sendNotificationToUsersWhoHaveMadePredictions() error {
	usersTemplate, err := getUsersWhoHaveNotMadePredictions(8)
	if err != nil {
		return err
	}

	for _, data := range usersTemplate {
		errN := insertNotificationEvent(models.Notification{UserId: data.UserId, MatchId: data.MatchId, NotificationMethod: "email"})
		if errN != nil {
			utils.ErrorLogger.Println("Error insert into notification event:", err)
			return fmt.Errorf("error insert into notification event: %v", err)
		}
	}
	return nil
}

func insertNotificationEvent(notifEvent models.Notification) error {
	query := "INSERT INTO Notifications (document_id, match_id, notification_time, notification_method) VALUES (?, ?, NOW(), ?)"
	_, err := database.InsertDBParams(query, notifEvent.UserId, notifEvent.MatchId, notifEvent.NotificationMethod)

	if err != nil {
		utils.ErrorLogger.Println("Error inserting notification event: ", err)
		return fmt.Errorf("error inserting notification event: %v", err)
	}
	return nil
}
