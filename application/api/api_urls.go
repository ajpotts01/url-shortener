package api

type urlRequest struct {
	Url string `json:"url"`
}

func (config *ApiConfig) CreateShortenedUrl(url urlRequest) (Url, error) {
	return Url{}, nil
}

func (config *ApiConfig) FetchShortenedUrl(url urlRequest) (Url, error) {
	return Url{}, nil
}
