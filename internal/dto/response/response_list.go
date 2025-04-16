package response

// ListResponse is a generic struct for API responses that return a list of data.
type ListResponse[T any] struct {
	List             []T    `json:"list"`              // The list of items
	TotalCount       int64  `json:"total_count"`       // Total number of items available
	Page             int    `json:"page"`              // Current page number
	PageSize         int    `json:"page_size"`         // Number of items per page
	FiltersAvailable []string `json:"filters_available"` // Filters applied to the list
}
