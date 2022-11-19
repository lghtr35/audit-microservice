package output

type Page struct {
	PageNumber   int32         `json:"page_number"`
	PageSize     int32         `json:"page_size"`
	TotalRecords int64         `json:"total_records"`
	Content      []interface{} `json:"content"`
}
