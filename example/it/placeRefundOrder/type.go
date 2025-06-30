package main

import "encoding/json"

type Order struct {
	AutoApproval bool       `json:"autoApproval"`
	Source       string     `json:"source"`
	RefundReason string     `json:"refundReason"`
	Resources    []Resource `json:"resources"`
	CustomInfo   CustomInfo `json:"customInfo"`
}

type CustomInfo struct {
	Phone    string   `json:"phone,omitempty"`
	Identity Identity `json:"identity"`
	Type     int      `json:"type"`
	Email    string   `json:"email,omitempty"`
}

type Identity struct {
	AccountId string `json:"accountId"`
}

type Resource struct {
	ResourceIds []string `json:"resourceIds"`
}

func (o *Order) Marshal() string {
	data, _ := json.Marshal(o)
	return string(data)
}
