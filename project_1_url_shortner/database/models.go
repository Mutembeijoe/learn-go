package database

type UrlResource struct {
	ID          int    `json:"id"`
	OriginalUrl string `json:"original_url" binding:"required"`
}
