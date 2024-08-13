package schemas

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	ProjectName string `gorm:"size:255;not null"`
	Path        string `gorm:"size:255;not null"`
	IssueID     string `gorm:"size:100;not null"`
	KeyID       string `gorm:"size:100;not null"`
}
