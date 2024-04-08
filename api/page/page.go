package page

import (
	"minevillages/community/database"
	"time"

	"github.com/google/uuid"
)

type PageCreateRequest struct {
	Title      string `json:"title"`
	Permission string `json:"permission"`
}

func (p *PageCreateRequest) PageCreate() database.PageInfo {
	pageId := generateUUID()
	pageInfo := database.PageInfo{
		PageId:     pageId,
		Title:      p.Title,
		Permission: p.Permission,
		CreatedAt:  time.Now(),
		UpdatedAt:  nil,
		DeletedAt:  nil,
	}
	database.CreatePage(pageInfo)
	return pageInfo
}

func generateUUID() string {
	id := uuid.New()
	return id.String()
}
