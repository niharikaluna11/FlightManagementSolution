package models

type Routes struct {
	ID uint `gorm:"primaryKey"`

	SourceId      uint `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	DestinationId uint `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	Source      Airport `gorm:"foreignKey:SourceId"`
	Destination Airport `gorm:"foreignKey:DestinationId"`
}
