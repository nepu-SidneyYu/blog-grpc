package model

type Request struct {
	Model       string     `json:"model"`
	Messages    []*Message `json:"messages"`
	ServiceCode string     `json:"serviceCode"`
	MaxTokens   int        `json:"maxTokens"`
	KeyType     string     `json:"keyType"`
	Tag         string     `json:"tag"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Response struct {
	Content string `json:"content"`
	Stop    bool   `json:"stop"`
}
