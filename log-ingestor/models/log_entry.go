package models

type LogEntry struct {
	Level       string                 `json:"level"`
	Message     string                 `json:"message"`
	ResourceId  string                 `json:"resourceId"`
	Timestamp   string                 `json:"timestamp"`
	TraceId     string                 `json:"traceId"`
	SpanId      string                 `json:"spanId"`
	Commit      string                 `json:"commit"`
	MetadataID  uint      						 `json:"metadata_id"`
	Metadata    Metadata  						 `json:"metadata" gorm:"foreignKey:MetadataID"`
}