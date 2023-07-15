package computer

import (
	"time"

	"gorm.io/gorm"
)

type Computer struct {
	Name             string `gorm:"primaryKey" json:"name" form:"name" binding:"required"`
	Ip               string `gorm:"column:ip;not null;" json:"ip" form:"IP" binding:"required"`
	Mac              string `gorm:"column:mac;not null;" json:"mac" form:"MAC" binding:"required"`
	Description      string `gorm:"column:description" json:"description" form:"description"`
	AssignedEmployee string `gorm:"column:assigned_employee" json:"assignedEmployee" form:"assigned_employee"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

func (m *Computer) BeforeCreate(tx *gorm.DB) error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *Computer) BeforeUpdate(tx *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}
