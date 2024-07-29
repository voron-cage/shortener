package shortener

type URLRequest struct {
	URL string `json:"url" binding:"required"`
}
