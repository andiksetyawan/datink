package common

type Paging struct {
	Limit      int `json:"limit,omitempty;" query:"limit"`
	Page       int `json:"page,omitempty;" query:"page"`
	TotalPages int `json:"totalPages"`
	TotalItems int `json:"totalItems"`
}

func (p *Paging) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Paging) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Paging) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Paging) GetTotalPage() int {
	if p.TotalItems%p.GetLimit() != 0 {
		return p.TotalItems/p.GetLimit() + 1
	}
	return p.TotalItems / p.GetLimit()
}
