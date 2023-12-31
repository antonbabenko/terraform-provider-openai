// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ListModelsResponse struct {
	Data   []Model `json:"data"`
	Object string  `json:"object"`
}

func (o *ListModelsResponse) GetData() []Model {
	if o == nil {
		return []Model{}
	}
	return o.Data
}

func (o *ListModelsResponse) GetObject() string {
	if o == nil {
		return ""
	}
	return o.Object
}
