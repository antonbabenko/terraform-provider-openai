// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type ListFineTunesResponse struct {
	Data   []FineTune `json:"data"`
	Object string     `json:"object"`
}

func (o *ListFineTunesResponse) GetData() []FineTune {
	if o == nil {
		return []FineTune{}
	}
	return o.Data
}

func (o *ListFineTunesResponse) GetObject() string {
	if o == nil {
		return ""
	}
	return o.Object
}
