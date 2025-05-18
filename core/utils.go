package errsuit

import (
	"encoding/json"
	"encoding/xml"
	"strings"

	"gopkg.in/yaml.v3"
)

func BuildErrorResp(appErr *AppError) ErrorResponse {
	msg := appErr.Message
	if appErr.Type == TypeInternal {
		msg = "internal server error"
	}
	return ErrorResponse{
		msg,
		appErr.Type,
	}
}

func determineFormat(ctx Context, respFormat ResponseFormat) ResponseFormat {
	if respFormat != ResponseFormatSmart {
		return respFormat
	}

	accept := ctx.GetHeader("Accept")
	accept = strings.ToLower(accept)

	for format, aliases := range contentTypeAliases {
		for _, alias := range aliases {
			if strings.Contains(accept, alias) {
				return format
			}
		}
	}

	// fallback
	return ResponseFormatJSON
}

func WriteError(ctx Context, err *AppError, format ResponseFormat) {
	actualFormat := determineFormat(ctx, format)
	body := BuildErrorResp(err)

	ctx.SetStatus(err.Code)

	switch actualFormat {
	case ResponseFormatJSON:
		ctx.SetHeader("Content-Type", "application/json")
		b, _ := json.Marshal(body)
		ctx.Write(b)
	case ResponseFormatYAML:
		ctx.SetHeader("Content-Type", "application/x-yaml")
		b, _ := yaml.Marshal(body)
		ctx.Write(b)
	case ResponseFormatXML:
		ctx.SetHeader("Content-Type", "application/xml")
		b, _ := xml.Marshal(body)
		ctx.Write(b)
	case ResponseFormatPlainText:
		ctx.SetHeader("Content-Type", "text/plain")
		b, _ := json.Marshal(body)
		ctx.Write(b)
	}
}
