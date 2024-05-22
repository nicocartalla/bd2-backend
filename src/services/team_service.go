package services

import (
    "fmt"
	"bd2-backend/src/database"
    "bd2-backend/src/models"
    "bd2-backend/src/utils"
)

type TeamService struct {
    Team models.Team
}

func (t *TeamService) CheckTeamExists(teamName string) bool {
	query := fmt.Sprintf("SELECT name as nameDB FROM Team WHERE name = '%s'", teamName)
	rows, err := database.QueryDB(query)
	if err != nil {
		utils.ErrorLogger.Println("Error checking if team exists: ", err)
		return false
	}
	defer rows.Close()

	var nameDB string
	if rows.Next() {
		err = rows.Scan(&nameDB)
		if err != nil {
			utils.ErrorLogger.Println("Error scanning team name: ", err)
			return false
		}
	}
	return teamName == nameDB
}

func (t *TeamService) GetTeams() ([]models.Team, error) {
	query := "SELECT team_id, name FROM Teams"
	rows, err := database.QueryDB(query)
	if err != nil {
		utils.ErrorLogger.Println("Error getting teams: ", err)
		return nil, fmt.Errorf("error getting teams: %v", err)
	}
	defer rows.Close()

	var teams []models.Team
	for rows.Next() {
		var team models.Team
		err = rows.Scan(&team.ID, &team.Name)
		if err != nil {
			utils.ErrorLogger.Println("Error scanning team: ", err)
			return nil, fmt.Errorf("error scanning team: %v", err)
		}
		teams = append(teams, team)
	}
	return teams, nil
}
/* 
func (u *UserService) checkUserExists() bool {
    query := fmt.Sprintf("SELECT username as usernameDB FROM User WHERE username = '%s'", u.User.Username)
    rows, err := database.QueryDB(query)
    if err != nil {
        ErrorLogger.Println("Error checking if user exists: ", err)
    }
    i := 0
    var usernameDB string
    for rows.Next() {
        i++
        err = rows.Scan(&usernameDB)
    }
    return i != 0 && u.User.Username == usernameDB
}

func (u *UserService) checkActiveUser() bool {
    query := fmt.Sprintf("SELECT username FROM User WHERE username = '%s' AND active = true", u.User.Username)
    rows, err := database.QueryDB(query)
    if err != nil {
        ErrorLogger.Println("Error checking if user exists: ", err)
    }
    i := 0
    var usernameDB string
    for rows.Next() {
        i++
        err = rows.Scan(&usernameDB)
    }
    return i != 0 && u.User.Username == usernameDB
} 
func (u *UserService) ValidateLogin() (bool, error) {
    if !u.checkUserExists() {
        return false, fmt.Errorf("User does not exist")
    }
    if !u.checkActiveUser() {
        return false, fmt.Errorf("User is not active")
    }
    query := fmt.Sprintf("SELECT id,password FROM User WHERE username = '%s'", u.User.Username)
    rows, err := database.QueryDB(query)
    if err != nil {
        ErrorLogger.Println(err.Error())
    }
    i := 0
    var hashFromBD string
    for rows.Next() {
        i++
        err = rows.Scan(&u.User.ID, &hashFromBD)
        if err != nil {
            ErrorLogger.Println(err.Error())
        }
    }
    if i == 0 {
        WarningLogger.Println("Error validate login: ", u.User.Username)
        return false, fmt.Errorf("Password error")
    }
    if i == 1 {
        if u.validateHash(hashFromBD) {
            return true, nil
        } else {
            WarningLogger.Println("Error validate login: ", u.User.Username)
            return false, fmt.Errorf("Password error")
        }
    }
    ErrorLogger.Println("Multiple users with username: ", u.User.Username)
    return false, fmt.Errorf("Multiple users with username: %s", u.User.Username)
}

func (u *UserService) createAvatar() string {
    return fmt.Sprintf("https://ui-avatars.com/api/?name=%s+%s?length=2", u.User.Username, u.User.LastName)
}

func (u *UserService) CreateUser() (int64, error) {
    if u.checkUserExists() {
        return 0, fmt.Errorf("User already exists")
    }
    pswHashed, errHash := hashing.HashPassword(u.User.Password)
    if errHash != nil {
        return 0, fmt.Errorf("Error hashing password")
    }
    u.User.Password = pswHashed
    query := fmt.Sprintf("INSERT INTO User (name, last_name, username, password, email, avatar, active) VALUES ('%s', '%s', '%s', '%s', '%s', '%s',1)", u.User.Name, u.User.LastName, u.User.Username, u.User.Password, u.User.Email, u.createAvatar())
    id, err := database.InsertDB(query)
    if err != nil {
        ErrorLogger.Println("Error creating user: ", err)
        return 0, fmt.Errorf("Error creating user: ", err)
    }
    u.User.ID = int(id)
    _, errBudget := u.createBudget()
    if errBudget != nil {
        ErrorLogger.Println("Error creating budget: ", err)
        return 0, fmt.Errorf("Error creating budget: ", err)
    }
    return id, nil
}

func (u *UserService) GetUser() (models.User, error) {
    query := fmt.Sprintf("SELECT id, name, last_name as lastName, email, avatar, username FROM User WHERE id = %s", strconv.Itoa(u.User.ID))
    rows, err := database.QueryDB(query)
    if err != nil {
        ErrorLogger.Println("Error getting user: ", err)
        return models.User{}, fmt.Errorf("Error getting user: ", err)
    }
    for rows.Next() {
        err = rows.Scan(&u.User.ID, &u.User.Name, &u.User.LastName, &u.User.Email, &u.User.AvatarLink, &u.User.Username)
        if err != nil {
            ErrorLogger.Println("Error getting user: ", err)
            return models.User{}, fmt.Errorf("Error getting user: ", err)
        }
    }
    return u.User, nil
}

func (u *UserService) createBudget() (int64, error) {
    query := fmt.Sprintf("INSERT INTO Budget (user_id, name, amount, start_date, end_date, current_budget) VALUES (%d, 'Budget', 0, now(),DATE_ADD(now(), INTERVAL 1 MONTH),1)", u.User.ID)
    id, err := database.InsertDB(query)
    if err != nil {
        ErrorLogger.Println("Error creating budget: ", err)
        return 0, fmt.Errorf("Error creating budget: ", err)
    }
    return id, nil
}

func (u *UserService) validateHash(hashFromBD string) bool {
    return hashing.CheckPasswordHash(u.User.Password, hashFromBD)
}
*/
