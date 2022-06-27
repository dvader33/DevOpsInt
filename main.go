package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var SECRET = []byte("super-secret-auth-key")
var api_key = "2f5ae96c-b558-4c7b-a590-a501ae1c3f6c"

type RequestM struct {
	To            string
	Message       string
	TimeToLifeSec int
	From          string
}

type ResponseM struct {
	Message string
}

func CreateJWT() (string, error) {
	//create Jason Web Token

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenStr, err := token.SignedString(SECRET)

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return tokenStr, nil
}

func ValidateJWT(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	//Validate Jason Web Token
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("not authorized"))
				}
				return SECRET, nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("not authorized: " + err.Error()))
			}

			if token.Valid {
				next(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("not authorized"))
		}
	})
}

func GetJwt(w http.ResponseWriter, r *http.Request) {
	if r.Header["Acceso"] != nil {
		if r.Header["Acceso"][0] != api_key {
			return
		} else {
			token, err := CreateJWT()
			if err != nil {
				return
			}
			fmt.Fprint(w, token)
		}
	}
}

func DevOps(w http.ResponseWriter, r *http.Request) {
	//DevOps function
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "ERROR")
		return
	}
	w.Header().Set("Content-Type", "application/json")

	requestm := &RequestM{}
	err := json.NewDecoder(r.Body).Decode(requestm)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "error", err)
		return
	}
	responsem := &ResponseM{}
	responsem.Message = "Hello " + requestm.To + " your message will be send"

	json.NewEncoder(w).Encode(responsem)
}

func Test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "MASTER")
}

func main() {

	http.Handle("/DevOps", ValidateJWT(DevOps))
	http.HandleFunc("/jwt", GetJwt)
	http.HandleFunc("/", Test)
	http.ListenAndServe(":80", nil)

}
