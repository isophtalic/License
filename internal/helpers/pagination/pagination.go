package pagination

type Paginate struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalPages int `json:"total_pages"`
}

func NewConfig(page, limit, offset, total_pages int) *Paginate {
	return &Paginate{
		Page:       page,
		Limit:      limit,
		TotalPages: total_pages,
	}
}

func Default() *Paginate {
	return &Paginate{
		Page:  1,
		Limit: 10,
	}
}
