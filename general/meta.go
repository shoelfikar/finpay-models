package general

type Meta struct {
	TotalData   uint32 `json:"total_data,omitempty"`
	TotalPage   uint32 `json:"total_page,omitempty"`
	CurrentPage uint32 `json:"current_page,omitempty"`
	Limit       uint32 `json:"limit,omitempty"`
}
