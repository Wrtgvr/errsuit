package errsuit

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

func (f ResponseFormat) String() string {
	switch f {
	case ResponseFormatSmart:
		return "smart"
	case ResponseFormatJSON:
		return "json"
	case ResponseFormatXML:
		return "xml"
	case ResponseFormatYAML:
		return "yaml"
	case ResponseFormatPlainText:
		return "plaintext"
	default:
		return "unknown"
	}
}
