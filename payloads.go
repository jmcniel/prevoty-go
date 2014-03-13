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

type TrustedQueryResult struct {
	Error               error                            `json:"error,omitempty"`
	ValidQuery          bool                             `json:"valid_query"`
	FoundViolations     bool                             `json:"found_violations"`
	NumSubqueries       int                              `json:"num_subqueries"`
	StatementType       string                           `json:"statement_type"`
	FieldsAccessed      []TrustedQuerySchemaObject       `json:"fields_accessed"`
	FieldsModified      []TrustedQuerySchemaObject       `json:"fields_modfied"`
	FunctionsCalled     []TrustedQueryFunctionCall       `json:"functions_called"`
	StatementViolations []TrustedQueryStatementViolation `json:"statement_violations"`
	FieldViolations     []TrustedQuerySchemaObject       `json:"field_violations"`
	FunctionViolations  []TrustedQueryFunctionCall       `json:"function_violations"`
}

type TrustedQuerySchemaObject struct {
	Database string `json:"database,omitempty"`
	Table    string `json:"table,omitempty"`
	Column   string `json:"column,omitempty"`
}

type TrustedQueryFunctionCall struct {
	Name   string                     `json:"name"`
	Tables []TrustedQuerySchemaObject `json:"tables"`
}

type TrustedQueryStatementViolation struct {
	Type  string                   `json:"type"`
	Table TrustedQuerySchemaObject `json:"table"`
}
