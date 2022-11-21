package input

type AuditInput struct {
	EventStatus   *bool                  `json:"event_status"`
	EventId       *int64                 `json:"event_id"`
	Action        string                 `json:"action"`
	Description   string                 `json:"description"`
	VariantFields map[string]interface{} `json:"variant_fields"`
}
