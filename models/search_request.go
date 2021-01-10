package models

import(

)

type SearchRequest struct {
	Type string `json:"type"`
	QueryID string `json:"id,omitempty"`
	QueryTags string `json:"tags,omitempty"`
}