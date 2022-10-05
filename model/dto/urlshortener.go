package dto

type CreateShortenURLRequest struct {
	Url       string `json:"url" binding:"required"`
	ExpiredAt string `json:"expireAt" binding:"required"`
}

type CreateShortenURLResponse struct {
	UrlId      string `json:"id"`
	ShortenUrl string `json:"shortenUrl"`
}
