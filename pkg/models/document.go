package models

import (
	"time"

	"github.com/google/uuid"
)

// Document represents a markdown document
type Document struct {
	ID           uuid.UUID `json:"id" db:"id"`
	Title        string    `json:"title" db:"title"`
	Content      string    `json:"content" db:"content"`
	OwnerID      uuid.UUID `json:"owner_id" db:"owner_id"`
	Version      int       `json:"version" db:"version"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	LastEditedBy uuid.UUID `json:"last_edited_by" db:"last_edited_by"`
	Tags         []string  `json:"tags" db:"tags"`
	IsPublic     bool      `json:"is_public" db:"is_public"`
	ShareToken   string    `json:"share_token,omitempty" db:"share_token"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" db:"deleted_at"` // Soft delete
}

// NewDocument creates a new document instance
func NewDocument(title, content string, ownerID uuid.UUID) *Document {
	now := time.Now()
	return &Document{
		ID:           uuid.New(),
		Title:        title,
		Content:      content,
		OwnerID:      ownerID,
		Version:      0,
		CreatedAt:    now,
		UpdatedAt:    now,
		LastEditedBy: ownerID,
		Tags:         []string{},
		IsPublic:     false,
		ShareToken:   "",
		DeletedAt:    nil,
	}
}

// GenerateShareToken generates a unique share token for the document
func (d *Document) GenerateShareToken() {
	d.ShareToken = uuid.New().String()
	d.UpdatedAt = time.Now()
}

// IncrementVersion increments the document version
func (d *Document) IncrementVersion(editorID uuid.UUID) {
	d.Version++
	d.LastEditedBy = editorID
	d.UpdatedAt = time.Now()
}

// SoftDelete marks the document as deleted
func (d *Document) SoftDelete() {
	now := time.Now()
	d.DeletedAt = &now
	d.UpdatedAt = now
}

// IsDeleted checks if the document is soft-deleted
func (d *Document) IsDeleted() bool {
	return d.DeletedAt != nil
}

// MakePublic makes the document publicly accessible
func (d *Document) MakePublic() {
	d.IsPublic = true
	if d.ShareToken == "" {
		d.GenerateShareToken()
	}
	d.UpdatedAt = time.Now()
}

// MakePrivate makes the document private
func (d *Document) MakePrivate() {
	d.IsPublic = false
	d.UpdatedAt = time.Now()
}
