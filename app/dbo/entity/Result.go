package entity

import (
	"gorm.io/gorm"
	"time"
)

type Result struct {
	Id            int            `gorm:"column:id;type:int;primaryKey;autoIncrement;not null"`
	ProjectId     int            `gorm:"column:project_id;type:int"`
	ProjectKey    string         `gorm:"column:project_key;type:varchar(255)"`
	Rule          string         `gorm:"column:rule;type:varchar(255)"`
	PrimaryUrl    string         `gorm:"column:primary_url;type:varchar(255)"`
	Path          string         `gorm:"column:path;type:varchar(255)"`
	Title         string         `gorm:"column:title;type:text"`
	Description   string         `gorm:"column:description;type:text"`
	Severity      string         `gorm:"column:severity;type:varchar(255)"`
	LastFoundAt   string         `gorm:"column:last_found_at;type:varchar(255)"`
	StatusResult  int            `gorm:"column:status_result;type:int"`
	References    string         `gorm:"column:references;type:text"`
	ScanType      string         `gorm:"column:scan_type;type:varchar(255)"`
	ScanVersion   int            `gorm:"column:scan_version;type:int"`
	Type          string         `gorm:"column:type;type:varchar(255)"`
	Message       string         `gorm:"column:message;type:text"`
	Resolution    string         `gorm:"column:resolution;type:varchar(255)"`
	CauseMetadata string         `gorm:"column:cause_metadata;type:text"`
	CreatedAt     time.Time      `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP;->"`
	UpdatedAt     time.Time      `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp;default:null;->"`
}

type CauseMetadata struct {
	Provider  string `json:"provider"`
	Service   string `json:"service"`
	StartLine int    `json:"start_line"`
}

func (Result) TableName() string {
	return "results"
}
