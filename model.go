package translate

type Request struct {
	Data       []Body   `json:"data"`
	Language   string   `json:"language"`
	ToLanguage []string `json:"toLanguage"`
}

type Body struct {
	Text string `json:"text"`
}

type Response struct {
	DetectedLanguage DetectedLanguage `json:"detectedLanguage"`
	Translations     []Translations   `json:"translations"`
}

type DetectedLanguage struct {
	Language string  `json:"language"`
	Score    float64 `json:"score"`
}

type Translations struct {
	Text string `json:"text"`
	To   string `json:"to"`
}
