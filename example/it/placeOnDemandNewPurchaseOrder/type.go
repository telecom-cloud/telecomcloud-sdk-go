package main

import "encoding/json"

type Order struct {
	AutoPay    bool        `json:"autoPay"`
	Source     string      `json:"source"`
	Orders     []OrderItem `json:"orders"`
	CustomInfo CustomInfo  `json:"customInfo"`
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

type OrderItem struct {
	Items       []Item `json:"items"`
	CycleType   int    `json:"cycleType"`
	CycleCnt    int    `json:"cycleCnt,omitempty"`
	InstanceCnt int    `json:"instanceCnt"`
}

type Item struct {
	ItemConfig   map[string]interface{} `json:"itemConfig"`
	Master       bool                   `json:"master"`
	ResourceType string                 `json:"resourceType"`
	ServiceTag   string                 `json:"serviceTag"`
	ItemValue    int                    `json:"itemValue"`
}

func (o *Order) Marshal() string {
	data, _ := json.Marshal(o)
	return string(data)
}
