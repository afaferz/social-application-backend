package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/afaferz/social-application/models"
	"github.com/afaferz/social-application/utils"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var users map[string][]byte = make(map[string][]byte)
var idxUsers int = 0

//getTokenUserPassword returns a jwt token for a user if the //password is ok
func GetTokenUserPassword(w http.ResponseWriter, r *http.Request) {
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "cannot decode username/password struct", http.StatusBadRequest)
		return
	}
	//here I have a user!
	//Now check if exists
	passwordHash, found := users[u.Username]
	if !found {
		http.Error(w, "Cannot find the username", http.StatusNotFound)
	}
	err = bcrypt.CompareHashAndPassword(passwordHash, []byte(u.Password))
	if err != nil {
		return
	}
	token, err := createToken(u.Username)
	if err != nil {
		http.Error(w, "Cannot create token", http.StatusInternalServerError)
		return
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Cannot decode request", http.StatusBadRequest)
		return
	}
	if _, found := users[u.Username]; found {
		http.Error(w, "User already exists", http.StatusBadRequest)
		return
	}
	//If I'm here-> add user and return a token
	value, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	users[u.Username] = value
	token, err := createToken(u.Username)
	if err != nil {
		http.Error(w, "Cannot create token", http.StatusInternalServerError)
		return
	}
}

func createToken(username string) (string, error) {
	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["username"] = username
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	secret := utils.GetSecret()
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return token, nil
}
