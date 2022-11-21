package output

type Page struct {
	PageNumber   int64         `json:"page_number"`
	PageSize     int64         `json:"page_size"`
	TotalRecords int64         `json:"total_records"`
	Content      []interface{} `json:"content"`
	BasicOutput
}
