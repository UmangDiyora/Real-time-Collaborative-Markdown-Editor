package models

import (
	"time"

	"github.com/google/uuid"
)

// DocumentVersion represents a snapshot of a document at a specific version
type DocumentVersion struct {
	ID         uuid.UUID `json:"id" db:"id"`
	DocumentID uuid.UUID `json:"document_id" db:"document_id"`
	Version    int       `json:"version" db:"version"`
	Content    string    `json:"content" db:"content"`
	AuthorID   uuid.UUID `json:"author_id" db:"author_id"`
	Operations string    `json:"operations,omitempty" db:"operations"` // JSON encoded OT operations
	Checksum   string    `json:"checksum" db:"checksum"`               // For integrity verification
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	Message    string    `json:"message,omitempty" db:"message"`       // Optional commit message
}

// NewDocumentVersion creates a new document version snapshot
func NewDocumentVersion(documentID uuid.UUID, version int, content string, authorID uuid.UUID) *DocumentVersion {
	return &DocumentVersion{
		ID:         uuid.New(),
		DocumentID: documentID,
		Version:    version,
		Content:    content,
		AuthorID:   authorID,
		Operations: "",
		Checksum:   generateChecksum(content),
		CreatedAt:  time.Now(),
		Message:    "",
	}
}

// NewDocumentVersionWithOps creates a new document version with operations
func NewDocumentVersionWithOps(documentID uuid.UUID, version int, content string, authorID uuid.UUID, operations string) *DocumentVersion {
	return &DocumentVersion{
		ID:         uuid.New(),
		DocumentID: documentID,
		Version:    version,
		Content:    content,
		AuthorID:   authorID,
		Operations: operations,
		Checksum:   generateChecksum(content),
		CreatedAt:  time.Now(),
		Message:    "",
	}
}

// SetMessage sets a commit message for the version
func (dv *DocumentVersion) SetMessage(message string) {
	dv.Message = message
}

// VerifyChecksum verifies the integrity of the content
func (dv *DocumentVersion) VerifyChecksum() bool {
	return dv.Checksum == generateChecksum(dv.Content)
}

// generateChecksum generates a simple checksum for content
// In production, use a proper hashing algorithm like SHA256
func generateChecksum(content string) string {
	// TODO: Implement proper checksum using crypto/sha256
	return uuid.NewSHA1(uuid.NameSpaceOID, []byte(content)).String()
}
