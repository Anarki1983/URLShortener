package po

import "time"

type CreateShortenURLRequest struct {
	UrlId    string
	Url      string
	Duration time.Duration
}

type CreateShortenURLResponse struct {
	UrlId string
}
