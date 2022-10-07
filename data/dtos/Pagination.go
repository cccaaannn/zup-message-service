package dtos

type Pagination struct {
	Size          int         `json:"size,omitempty;query:size"`
	Page          int         `json:"page,omitempty;query:page"`
	Sort          string      `json:"sort,omitempty;query:sort"`
	TotalElements int64       `json:"totalElements"`
	TotalPages    int         `json:"totalPages"`
	Content       interface{} `json:"content"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Size == 0 {
		p.Size = 10
	}
	return p.Size
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == "" {
		p.Sort = "Id desc"
	}
	return p.Sort
}
