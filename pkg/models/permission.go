package models

import (
	"time"

	"github.com/google/uuid"
)

// PermissionLevel represents the access level for a document
type PermissionLevel string

const (
	// PermissionView allows read-only access
	PermissionView PermissionLevel = "VIEW"

	// PermissionEdit allows editing the document
	PermissionEdit PermissionLevel = "EDIT"

	// PermissionAdmin allows full control including sharing and deletion
	PermissionAdmin PermissionLevel = "ADMIN"
)

// Permission represents access control for a document
type Permission struct {
	ID         uuid.UUID       `json:"id" db:"id"`
	DocumentID uuid.UUID       `json:"document_id" db:"document_id"`
	UserID     uuid.UUID       `json:"user_id" db:"user_id"`
	Permission PermissionLevel `json:"permission" db:"permission"`
	GrantedBy  uuid.UUID       `json:"granted_by" db:"granted_by"`
	CreatedAt  time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at" db:"updated_at"`
	ExpiresAt  *time.Time      `json:"expires_at,omitempty" db:"expires_at"` // Optional expiration
}

// NewPermission creates a new permission instance
func NewPermission(documentID, userID, grantedBy uuid.UUID, level PermissionLevel) *Permission {
	now := time.Now()
	return &Permission{
		ID:         uuid.New(),
		DocumentID: documentID,
		UserID:     userID,
		Permission: level,
		GrantedBy:  grantedBy,
		CreatedAt:  now,
		UpdatedAt:  now,
		ExpiresAt:  nil,
	}
}

// SetExpiration sets an expiration time for the permission
func (p *Permission) SetExpiration(expiresAt time.Time) {
	p.ExpiresAt = &expiresAt
	p.UpdatedAt = time.Now()
}

// IsExpired checks if the permission has expired
func (p *Permission) IsExpired() bool {
	if p.ExpiresAt == nil {
		return false
	}
	return time.Now().After(*p.ExpiresAt)
}

// CanView checks if the permission allows viewing
func (p *Permission) CanView() bool {
	return !p.IsExpired() && (p.Permission == PermissionView || p.Permission == PermissionEdit || p.Permission == PermissionAdmin)
}

// CanEdit checks if the permission allows editing
func (p *Permission) CanEdit() bool {
	return !p.IsExpired() && (p.Permission == PermissionEdit || p.Permission == PermissionAdmin)
}

// CanAdmin checks if the permission allows admin actions
func (p *Permission) CanAdmin() bool {
	return !p.IsExpired() && p.Permission == PermissionAdmin
}

// Upgrade upgrades the permission level
func (p *Permission) Upgrade(newLevel PermissionLevel) {
	p.Permission = newLevel
	p.UpdatedAt = time.Now()
}

// Downgrade downgrades the permission level
func (p *Permission) Downgrade(newLevel PermissionLevel) {
	p.Permission = newLevel
	p.UpdatedAt = time.Now()
}

// IsValid checks if the permission level is valid
func (pl PermissionLevel) IsValid() bool {
	return pl == PermissionView || pl == PermissionEdit || pl == PermissionAdmin
}
