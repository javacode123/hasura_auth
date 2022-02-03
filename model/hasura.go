package model

// HasuraUserRsp 为结构化 Data
//type Response struct {
//	Errors     []*errors.QueryError   `json:"errors,omitempty"`
//	Data       json.RawMessage        `json:"data,omitempty"`
//	Extensions map[string]interface{} `json:"extensions,omitempty"`
//}
type HasuraUserRsp struct {
	User struct {
		Id    string   `json:"id"`
		Name  string   `json:"name"`
		Roles []string `json:"roles"`
	} `json:"user"`
}
