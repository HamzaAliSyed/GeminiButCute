package services

import (
	"backend/internal/database"
	"backend/internal/models"
	"context"
	"errors"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func isValidEmail(email string) bool {
	atIndex := strings.Index(email, "@")
	if atIndex < 1 {
		return false
	}

	dotIndex := strings.Index(email[atIndex:], ".")
	return dotIndex > 1
}

func CreateUser(newUser models.Users) error {
	if !isValidEmail(newUser.Email) {
		return errors.New("invalid email by user")
	}

	userCheckContext, contextCancel := context.WithTimeout(context.Background(), time.Second*10)
	defer contextCancel()

	collections := database.GetCollections()
	_, insertError := collections.Users.InsertOne(userCheckContext, newUser)
	if insertError != nil {
		if mongoWriteException, ok := insertError.(mongo.WriteException); ok {
			for _, writeError := range mongoWriteException.WriteErrors {
				if writeError.Code == 11000 {
					return errors.New("username or email already exists")
				}
			}
		}
		return insertError
	}

	return nil
}
