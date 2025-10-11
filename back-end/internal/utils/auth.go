package utils

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserIDFromContext(ctx *gin.Context) (primitive.ObjectID, error) {
	userIDValue, exists := ctx.Get("user_id")
	if !exists {
		err := errors.New("userID not found in context, middleware might be missing")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return primitive.NilObjectID, err
	}

	userID, ok := userIDValue.(primitive.ObjectID)
	if !ok {
		err := errors.New("userID in context is not of type primitive.ObjectID")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error: invalid user context"})
		return primitive.NilObjectID, err
	}

	if userID.IsZero() {
		err := errors.New("userID in context is a zero ObjectID")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: invalid user ID"})
		return primitive.NilObjectID, err
	}

	return userID, nil
}
