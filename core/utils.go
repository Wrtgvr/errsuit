package errsuit

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

func BuildErrorResp(appErr *AppError) ErrorResponse {
	return ErrorResponse{
		ErrMsg: appErr.Message,
		Typ:    appErr.Type,
		Code:   appErr.Code,
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
		b := []byte(fmt.Sprintf("%s: %s", body.Typ, body.ErrMsg))
		ctx.Write(b)
	}
}
