package handlers

import (
	"backend/config"
	"backend/models"
	"database/sql"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func RegisterUser(c *fiber.Ctx) error {
	var input models.User

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
	}

	input.Password = string(hashedPassword)

	query := `INSERT INTO users (id, email, password, name, team_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var userID string
	err = config.DB.QueryRow(query, input.ID, input.Email, input.Password, input.Name, input.TeamID).Scan(&userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.Status(400).JSON(fiber.Map{"error": "User not created"})
		}
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		ID:    userID,
		Email: input.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "User registered successfully",
		"token":   tokenString,
	})
}
