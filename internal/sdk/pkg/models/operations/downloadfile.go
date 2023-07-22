// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DownloadFileRequest struct {
	// The ID of the file to use for this request
	FileID string `pathParam:"style=simple,explode=false,name=file_id"`
}

type DownloadFileResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	DownloadFile200ApplicationJSONString *string
}
