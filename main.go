package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"todo/todo"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/mux"
)

func loggingMiddlewareGorillaMux(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func authMiddlewareGorillaMux(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Use authMiddleware")

		tokenString := r.Header.Get("Authorization")

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
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func authMiddlewareGin(c *gin.Context) {
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
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":9090")
}

func mainGin() {
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
	api.Use(authMiddlewareGin)

	api.PUT("/todos", todo.AddTaskGin)
	api.PUT("/todos/:index", todo.MaskDoneGin)
	api.GET("/todos", todo.GetTodoGin)

	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run(":9090")
}

func mainGorillaMux() {
	r := mux.NewRouter()
	r.Use(loggingMiddlewareGorillaMux)

	r.HandleFunc("/auth", func(rw http.ResponseWriter, r *http.Request) {
		mySigningKey := []byte("password")
		claims := &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Minute).Unix(),
			Issuer:    "test",
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := token.SignedString(mySigningKey)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		json.NewEncoder(rw).Encode(map[string]string{
			"token": ss,
		})
	})

	api := r.NewRoute().Subrouter()
	api.Use(authMiddlewareGorillaMux)
	api.HandleFunc("/todos", todo.AddTaskG).Methods(http.MethodPut)
	api.HandleFunc("/todos/{index}", todo.MarkDoneG).Methods(http.MethodPut)
	api.HandleFunc("/todos", todo.GetTodoG).Methods(http.MethodGet)

	err := http.ListenAndServe(":9090", r)
	fmt.Println(err)
}
