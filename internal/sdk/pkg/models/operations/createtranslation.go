// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openai/v2/internal/sdk/pkg/models/shared"
)

type CreateTranslationResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	CreateTranslationResponse *shared.CreateTranslationResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CreateTranslationResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateTranslationResponse) GetCreateTranslationResponse() *shared.CreateTranslationResponse {
	if o == nil {
		return nil
	}
	return o.CreateTranslationResponse
}

func (o *CreateTranslationResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateTranslationResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
