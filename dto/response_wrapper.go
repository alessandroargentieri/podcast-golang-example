package dto

import()

type ResponseWrapper struct {
	Success  bool               `json:"success"`
	Entities []interface{}      `json:"entities,omitempty"`
	Error    string             `json:"error,omitempty"`
}
