package models

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user in the system
type User struct {
	ID           uuid.UUID `json:"id" db:"id"`
	Username     string    `json:"username" db:"username"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"-" db:"password_hash"` // Never expose password hash in JSON
	Avatar       string    `json:"avatar" db:"avatar"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	LastLoginAt  *time.Time `json:"last_login_at,omitempty" db:"last_login_at"`
	IsActive     bool      `json:"is_active" db:"is_active"`
	IsVerified   bool      `json:"is_verified" db:"is_verified"`
	Settings     UserSettings `json:"settings" db:"settings"`
}

// UserSettings contains user-specific settings
type UserSettings struct {
	Theme            string `json:"theme"`             // light, dark, auto
	EditorMode       string `json:"editor_mode"`       // vim, emacs, default
	FontSize         int    `json:"font_size"`         // in pixels
	TabSize          int    `json:"tab_size"`          // spaces per tab
	AutoSave         bool   `json:"auto_save"`
	EmailNotifications bool `json:"email_notifications"`
	Language         string `json:"language"`          // en, es, fr, etc.
}

// NewUser creates a new user instance with default values
func NewUser(username, email, passwordHash string) *User {
	now := time.Now()
	return &User{
		ID:           uuid.New(),
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
		Avatar:       "",
		CreatedAt:    now,
		UpdatedAt:    now,
		LastLoginAt:  nil,
		IsActive:     true,
		IsVerified:   false,
		Settings: UserSettings{
			Theme:              "light",
			EditorMode:         "default",
			FontSize:           14,
			TabSize:            2,
			AutoSave:           true,
			EmailNotifications: true,
			Language:           "en",
		},
	}
}

// UpdateLastLogin updates the last login timestamp
func (u *User) UpdateLastLogin() {
	now := time.Now()
	u.LastLoginAt = &now
	u.UpdatedAt = now
}

// Verify marks the user as verified
func (u *User) Verify() {
	u.IsVerified = true
	u.UpdatedAt = time.Now()
}
