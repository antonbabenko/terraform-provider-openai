// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type DeleteModelResponse struct {
	Deleted bool   `json:"deleted"`
	ID      string `json:"id"`
	Object  string `json:"object"`
}

func (o *DeleteModelResponse) GetDeleted() bool {
	if o == nil {
		return false
	}
	return o.Deleted
}

func (o *DeleteModelResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DeleteModelResponse) GetObject() string {
	if o == nil {
		return ""
	}
	return o.Object
}
