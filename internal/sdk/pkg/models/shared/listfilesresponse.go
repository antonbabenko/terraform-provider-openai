// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ListFilesResponse struct {
	Data   []OpenAIFile `json:"data"`
	Object string       `json:"object"`
}

func (o *ListFilesResponse) GetData() []OpenAIFile {
	if o == nil {
		return []OpenAIFile{}
	}
	return o.Data
}

func (o *ListFilesResponse) GetObject() string {
	if o == nil {
		return ""
	}
	return o.Object
}
