package responses

// TimestampResponse is for responding with a timestamp
type TimestampResponse struct {
	Timestamp int64 `json:"timestamp"`
}

// DateResponse is for responding with a date
type DateResponse struct {
	Date string `json:"date"`
}

// JsonError is for responding with an error
type JsonError struct {
	Message string `json:"message"`
}
