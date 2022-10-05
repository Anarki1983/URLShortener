package bo

import "time"

type CreateShortenURLRequest struct {
	Url       string
	ExpiredAt time.Time
}

type CreateShortenURLResponse struct {
	UrlId string
}
