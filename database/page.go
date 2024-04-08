package database

import (
	"time"
)

type PageInfo struct {
	PageId     string     `json:"pageId" gorm:"column:pageId"`
	Title      string     `json:"title" gorm:"column:title"`
	Permission string     `json:"permission" gorm:"column:permission"`
	CreatedAt  time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt" gorm:"column:deletedAt"`
}

func CreatePage(page PageInfo) {
	db.Table("page").Create(&page)
}
