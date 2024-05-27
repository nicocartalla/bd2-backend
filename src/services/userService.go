package services

import (
    "fmt"
	"bd2-backend/src/database"
	"bd2-backend/src/hashing"
    "bd2-backend/src/models"
	"bd2-backend/src/utils"
    "strconv"
)

type UserService struct {
    User models.User
}

func (u *UserService) checkUserExists() bool {
    query := fmt.Sprintf("SELECT email as Email FROM User WHERE email = '%s'", u.User.Email)
    rows, err := database.QueryDB(query)
    if err != nil {
        utils.ErrorLogger.Println("Error checking if user exists: ", err)
    }
    i := 0
    var emailDB string
    for rows.Next() {
        i++
    }
    return i != 0 && u.User.Email == emailDB
}

func (u *UserService) ValidateLogin() (bool, error) {
    if !u.checkUserExists() {
        return false, fmt.Errorf("user does not exist")
    }
    query := fmt.Sprintf("SELECT user_id, password FROM User WHERE email = '%s'", u.User.Email)
    rows, err := database.QueryDB(query)
    if err != nil {
        utils.ErrorLogger.Println(err.Error())
    }
    i := 0
    var hashFromBD string
    for rows.Next() {
        i++
        err = rows.Scan(&u.User.ID, &hashFromBD)
        if err != nil {
            utils.ErrorLogger.Println(err.Error())
        }
    }
    if i == 0 {
        utils.WarningLogger.Println("Error validate login: ", u.User.Email)
        return false, fmt.Errorf("password error")
    }
    if i == 1 {
        if u.validateHash(hashFromBD) {
            return true, nil
        } else {
            utils.WarningLogger.Println("Error validate login: ", u.User.Email)
            return false, fmt.Errorf("password error")
        }
    }
    utils.ErrorLogger.Println("Multiple users with email: ", u.User.Email)
    return false, fmt.Errorf("multiple users with email: %s", u.User.Email)
}

// func (u *UserService) createAvatar() string {
//     return fmt.Sprintf("https://ui-avatars.com/api/?name=%s+%s?length=2", u.User.Email, u.User.LastName)
// }

func (u *UserService) CreateUser() (int64, error) {
    if u.checkUserExists() {
        return 0, fmt.Errorf("user already exists")
    }
    pswHashed, errHash := hashing.HashPassword(u.User.Password)
    if errHash != nil {
        return 0, fmt.Errorf("error hashing password")
    }
    u.User.Password = pswHashed
    query := fmt.Sprintf("INSERT INTO User (email, first_name, last_name, major, password, role) VALUES ('%s', '%s', '%s', '%s', '%s', '%s')", u.User.Email, u.User.FirstName, u.User.LastName, u.User.Major, u.User.Password, u.User.Role)
    id, err := database.InsertDB(query)
    if err != nil {
        utils.ErrorLogger.Println("Error creating user: ", err)
        return 0, fmt.Errorf("error creating user: %v", err)
    }
    u.User.ID = int(id)

    return id, nil
}

func (u *UserService) GetUser() (models.User, error) {
    query := fmt.Sprintf("SELECT user_id as Id, first_name as firstName, last_name as lastName, email, role FROM User WHERE id = %s", strconv.Itoa(u.User.ID))
    rows, err := database.QueryDB(query)
    if err != nil {
        utils.ErrorLogger.Println("Error getting user: ", err)
        return models.User{}, fmt.Errorf("error getting user: %v", err)
    }
    for rows.Next() {
        err = rows.Scan(&u.User.ID, &u.User.FirstName, &u.User.LastName, &u.User.Email, &u.User.Role)
        if err != nil {
            utils.ErrorLogger.Println("Error getting user: ", err)
            return models.User{}, fmt.Errorf("error getting user: %v", err)
        }
    }
    return u.User, nil
}

func (u *UserService) validateHash(hashFromBD string) bool {
    return hashing.CheckPasswordHash(u.User.Password, hashFromBD)
}
