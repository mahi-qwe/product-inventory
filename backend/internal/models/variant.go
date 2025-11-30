package models

import "github.com/google/uuid"

type Variant struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	ProductID uuid.UUID
	Name      string

	Options []VariantOption `gorm:"foreignKey:VariantID"`
}
