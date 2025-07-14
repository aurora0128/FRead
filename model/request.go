package model

type AddArticleReq struct {
	URL string `json:"url" binding:"required"`
}
