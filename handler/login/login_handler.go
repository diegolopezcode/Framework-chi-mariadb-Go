package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/diegolopezcode/api-crud-complete-chi/configs"
	"github.com/go-chi/jwtauth/v5"
)

type LoginRequest struct {
	Email    string
	Password string
}
type User struct {
	Id       int
	Email    string
	Password string
}

var TokenAuth *jwtauth.JWTAuth

// It takes a request body with an email and password, checks if the email and password are correct,
// and if they are, it returns a JWT token and the user's information
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println(configs.Config("NOMBRE"), configs.Config("NFS_PATH"))
	req := new(LoginRequest)

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return
	}

	if req.Email == "" || req.Password == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Email and password are required",
		})
		return
	}

	if req.Email != configs.Config("USER_EMAIL") || req.Password != configs.Config("USER_PASSWORD") {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Email or password are incorrect",
		})
		return
	}
	user := User{Id: 1, Email: req.Email}
	token, exp, err := createJWTToken(user)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "Error creating token",
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"token": token,
		"exp":   exp,
		"user":  user,
		"name":  configs.Config("NOMBRE"),
	})

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Unix(exp, 0),
	})

}

// It creates a JWT token with the user's ID and email address, and sets the expiration time to 30
// minutes from now
func createJWTToken(user User) (string, int64, error) {

	exp := time.Now().Add(time.Minute * 30).Unix()
	TokenAuth = jwtauth.New("HS256", []byte(configs.Config("JWT_SECRET")), nil)
	_, tokenString, _ := TokenAuth.Encode(
		map[string]interface{}{
			"exp":   exp,
			"id":    user.Id,
			"email": user.Email,
		})
	return tokenString, exp, nil
}
