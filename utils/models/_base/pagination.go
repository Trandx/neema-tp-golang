package base

type Pagination struct {
	CurrentPage int64 `json:"current_page"`
	Size        int64 `json:"size"`
	TotalItem   int64 `json:"total_item"`
	TotalPages  int64 `json:"total_pages"`
}
