package errsuit

// Error types
const (
	TypeBadRequest   = "bad_request"
	TypeNotFound     = "not_found"
	TypeInternal     = "internal"
	TypeUnauthorized = "unauthorized"
	TypeForbidden    = "forbidden"
	TypeConflict     = "conflict"
)

type ResponseFormat int

// Response formats
const (
	ResponseFormatSmart ResponseFormat = iota
	ResponseFormatJSON
	ResponseFormatXML
	ResponseFormatYAML
	ResponseFormatPlainText
)

// Aliases for response formats
var contentTypeAliases = map[ResponseFormat][]string{
	ResponseFormatJSON:      {"application/json", "text/json", "+json"},
	ResponseFormatXML:       {"application/xml", "text/xml", "+xml"},
	ResponseFormatYAML:      {"application/x-yaml", "text/yaml", "application/yaml", "+yaml"},
	ResponseFormatPlainText: {"text/plain"},
}
