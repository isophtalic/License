package dto

type PaginationDTO struct {
	PerPage    int    `json:"per_page" query:"perPage"`
	Page       int    `json:"page" query:"page"`
	Sort       string `json:"sort" query:"sort"`
	TotalRows  int64  `json:"total_rows"`
	TotalPages int    `json:"total_pages"`
}

func (p *PaginationDTO) GetPerPage() int {
	if p.PerPage == 0 {
		p.PerPage = 20
	}
	return p.PerPage
}

func (p *PaginationDTO) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *PaginationDTO) GetSort() string {
	return p.Sort
}

func (p *PaginationDTO) GetOffset() int {
	return (p.GetPage() - 1) * p.GetPerPage()
}

func (p *PaginationDTO) GetTotalPages() int {
	return p.TotalPages
}
