package prevoty

import "time"

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

type LinkAnalysisResult struct {
	Analysis LinkAnalysis `json:"analysis"`
	Message  string       `json:"message"`
}

type LinkAnalysis struct {
	ScanId           string                      `json:"scan_id"`
	CustomerId       uint                        `json:"customer_id"`
	Url              string                      `json:"url"`
	AnalysisCounters LinkAnalysisCounters        `json:"analysis_counters"`
	ExpectedBaseline LinkAnalysisBaseline        `json:"expected_baseline"`
	DeviationReport  LinkAnalysisDeviationReport `json:"deviation_report"`
	Screenshots      []LinkAnalysisScreenshot    `json:"screenshots"`
	BrowserEvents    LinkAnalysisBrowserEvents   `json:"browser_events"`
	DnsRecords       LinkAnalysisDnsRecords      `json:"dns_records"`
	HtmlDocument     LinkAnalysisHtmlDocument    `json:"html_document"`
	Created          time.Time                   `json:"created"`
}

type LinkAnalysisCounters struct {
	JavaScriptRedirects    uint `json:"javascript_redirects"`
	ClientRedirects        uint `json:"client_redirects"`
	ServerRedirects        uint `json:"server_redirects"`
	Popups                 uint `json:"popups"`
	Downloads              uint `json:"downloads"`
	HtmlDocumentViolations uint `json:"html_document_violations"`
	DnsRecordViolations    uint `json:"dns_record_violations"`
}

type LinkAnalysisBaseline struct {
	BrowserEvents LinkAnalysisBrowserEvents `json:"browser_events"`
	DnsRecords    LinkAnalysisDnsRecords    `json:"dns_records"`
	HtmlDocument  LinkAnalysisHtmlDocument  `json:"html_document"`
}

type LinkAnalysisBrowserEvents []LinkAnalysisBrowserEvent

type LinkAnalysisBrowserEvent struct {
	TabId      int       `json:"tab_id"`
	EventName  string    `json:"event_name"`
	EventValue string    `json:"event_value"`
	Timestamp  time.Time `json:"timestamp"`
	MimeType   string    `json:"mime_type"`
}

type LinkAnalysisDeviation struct {
	Property string `json:"property"`
	Expected string `json:"expected"`
	Actual   string `json:"actual"`
}

type LinkAnalysisDeviationReport struct {
	HtmlDocumentChanged []LinkAnalysisDeviation `json:"html_document_changed"`
	HtmlDocumentRemoved []LinkAnalysisDeviation `json:"html_document_removed"`
	DnsRecordsRemoved   []LinkAnalysisDeviation `json:"dns_records_removed"`
	DnsRecordsAdded     []LinkAnalysisDeviation `json:"dns_records_added"`
	NewBrowserEvents    []LinkAnalysisDeviation `json:"new_browser_events"`
	FailedBrowserEvents []LinkAnalysisDeviation `json:"failed_browser_events"`
}

type LinkAnalysisDnsRecords map[string][]string

type LinkAnalysisHtmlDocument struct {
	Title    string            `json:"title"`
	MetaTags map[string]string `json:"meta_tags"`
}

type LinkAnalysisScreenshot struct {
	TabId         int    `json:"tab_id"`
	ScreenshotUrl string `json:"screenshot_url"`
}
