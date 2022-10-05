package dto

type RedirectURLRequest struct {
	UrlId string `uri:"url_id" binding:"required"`
}

type RedirectURLResponse struct {
	Url string `json:"url"`
}
