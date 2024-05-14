package models

type LogEntry struct {
	Level       string                 `json:"level"`
	Message     string                 `json:"message"`
	ResourceId  string                 `json:"resourceId"`
	Timestamp   string                 `json:"timestamp"`
	TraceId     string                 `json:"traceId"`
	SpanId      string                 `json:"spanId"`
	Commit      string                 `json:"commit"`
	Metadata    Metadata 							 `json:"metadata"`
}

type Metadata struct {
	ParentResourceId string `json:"parentResourceId"`
}