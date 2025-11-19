package models

import (
	"time"

	"github.com/google/uuid"
)

// Session represents an active editing session
type Session struct {
	ID             string    `json:"id" db:"id"`
	UserID         uuid.UUID `json:"user_id" db:"user_id"`
	DocumentID     uuid.UUID `json:"document_id" db:"document_id"`
	CursorPosition int       `json:"cursor_position" db:"cursor_position"`
	SelectionStart int       `json:"selection_start" db:"selection_start"`
	SelectionEnd   int       `json:"selection_end" db:"selection_end"`
	Color          string    `json:"color" db:"color"` // User's cursor color
	LastActivity   time.Time `json:"last_activity" db:"last_activity"`
	ConnectionID   string    `json:"connection_id" db:"connection_id"` // WebSocket connection ID
	IsActive       bool      `json:"is_active" db:"is_active"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
}

// NewSession creates a new session instance
func NewSession(userID, documentID uuid.UUID, connectionID string) *Session {
	now := time.Now()
	return &Session{
		ID:             uuid.New().String(),
		UserID:         userID,
		DocumentID:     documentID,
		CursorPosition: 0,
		SelectionStart: 0,
		SelectionEnd:   0,
		Color:          generateUserColor(userID),
		LastActivity:   now,
		ConnectionID:   connectionID,
		IsActive:       true,
		CreatedAt:      now,
	}
}

// UpdateCursor updates the cursor position
func (s *Session) UpdateCursor(position int) {
	s.CursorPosition = position
	s.LastActivity = time.Now()
}

// UpdateSelection updates the text selection
func (s *Session) UpdateSelection(start, end int) {
	s.SelectionStart = start
	s.SelectionEnd = end
	s.LastActivity = time.Now()
}

// UpdateActivity updates the last activity timestamp
func (s *Session) UpdateActivity() {
	s.LastActivity = time.Now()
}

// Deactivate marks the session as inactive
func (s *Session) Deactivate() {
	s.IsActive = false
	s.LastActivity = time.Now()
}

// IsExpired checks if the session has expired (no activity for 5 minutes)
func (s *Session) IsExpired(timeout time.Duration) bool {
	return time.Since(s.LastActivity) > timeout
}

// generateUserColor generates a consistent color for a user based on their ID
func generateUserColor(userID uuid.UUID) string {
	colors := []string{
		"#FF6B6B", // Red
		"#4ECDC4", // Cyan
		"#45B7D1", // Blue
		"#FFA07A", // Light Salmon
		"#98D8C8", // Mint
		"#F7DC6F", // Yellow
		"#BB8FCE", // Purple
		"#85C1E2", // Sky Blue
		"#F8B88B", // Peach
		"#AAB7B8", // Gray
	}

	// Use the first byte of UUID to select a color
	index := int(userID[0]) % len(colors)
	return colors[index]
}
