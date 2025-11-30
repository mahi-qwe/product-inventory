package models

import "github.com/google/uuid"

type VariantOption struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	VariantID uuid.UUID
	Value     string
}
