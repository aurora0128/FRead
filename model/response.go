package model

//todo 添加单文件的response

type GetArticleListRes struct {
	Success    bool      `json:"success"`
	Data       []Article `json:"data"`
	Pagination struct {
		Page  int `json:"page"`
		Limit int `json:"limit"`
		Total int `json:"total"`
		Pages int `json:"pages"`
	} `json:"pagination"`
}
type GetArticleRes struct {
	Success bool    `json:"success"`
	Data    Article `json:"data"`
}
