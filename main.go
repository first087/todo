package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"todo/todo"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func authMiddleware(c *gin.Context) {
	fmt.Println("Use authMiddleware")

	tokenString := c.GetHeader("Authorization")

	tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")

	fmt.Println(tokenString)

	mySigningKey := []byte("password")
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return mySigningKey, nil
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}

func main() {
	r := gin.Default()
	r.GET("/auth", func(c *gin.Context) {
		mySigningKey := []byte("password")
		claims := &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Minute).Unix(),
			Issuer:    "test",
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := token.SignedString(mySigningKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		c.JSON(http.StatusOK, map[string]string{
			"token": ss,
		})
	})

	api := r.Group("")
	api.Use(authMiddleware)

	r.PUT("/todos", todo.AddTask)
	r.PUT("/todos/:index", todo.MaskDone)
	r.GET("/todos", todo.GetTodo)

	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run(":9090")
}
