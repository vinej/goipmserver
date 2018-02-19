package services

import (
	"goipmserver/api/parameters"
	"goipmserver/core/authentication"
	"goipmserver/services/models"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"net/http"
	"gopkg.in/mgo.v2/bson"
)

func Login(requestUser *models.User) (int, []byte) {
	authBackend := authentication.InitJWTAuthenticationBackend()

	searchResults, searchError := SearchCollection("users", bson.M{ "user": requestUser.User},0,1)
	if searchError != "" {
		return http.StatusUnauthorized, []byte(searchError)
	}
	var testUser models.User
	err := models.SetStruct(searchResults[0], &testUser)
	if err != nil {
		return http.StatusUnauthorized, []byte(err.Error())
	}

	if authBackend.Authenticate(requestUser, &testUser ) {
		token, err := authBackend.GenerateToken(requestUser.User)
		if err != nil {
			return http.StatusInternalServerError, []byte("")
		} else {
			response, _ := json.Marshal(parameters.TokenAuthentication{token})
			return http.StatusOK, response
		}
	}

	return http.StatusUnauthorized, []byte("")
}

func Register(requestUser *models.User) (int, []byte) {
	authBackend := authentication.InitJWTAuthenticationBackend()
	requestUser.Password = string(authBackend.Register(requestUser))
	InsertCollection("users", requestUser)
	return http.StatusOK, []byte("")
}

func RefreshToken(requestUser *models.User) []byte {
	authBackend := authentication.InitJWTAuthenticationBackend()
	token, err := authBackend.GenerateToken(requestUser.User)
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(parameters.TokenAuthentication{token})
	if err != nil {
		panic(err)
	}
	return response
}

func Logout(req *http.Request) error {
	authBackend := authentication.InitJWTAuthenticationBackend()
	tokenRequest, err := request.ParseFromRequest(req, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		return authBackend.PublicKey, nil
	})
	if err != nil {
		return err
	}
	tokenString := req.Header.Get("Authorization")
	return authBackend.Logout(tokenString, tokenRequest)
}
