package prevoty

type KeyInformation struct {
	Maximum   int    `json:"maximum"`
	Used      int    `json:"used"`
	Remaining int    `json:"remaining"`
	Message   string `json:"message"`
}

type TrustedContentResult struct {
	Message    string                    `json:"message"`
	Output     string                    `json:"output"`
	Statistics *TrustedContentStatistics `json:"statistics"`
}

type TrustedContentStatistics struct {
	InvalidAttributes    int `json:"invalid_attributes"`
	InvalidProtocols     int `json:"invalid_protocols"`
	InvalidTags          int `json:"invalid_tags"`
	JavaScriptAttributes int `json:"javascript_attributes"`
	JavaScriptProtocols  int `json:"javascript_protocols"`
	JavaScriptTags       int `json:"javascript_tags"`
	TagsBalanced         int `json:"tags_balanced"`
	Transformations      int `json:"transformations"`
}

type TrustedTokenGenerationResult struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}

type TrustedTokenValidationResult struct {
	Valid   bool   `json:"valid"`
	Message string `json:"message"`
}
