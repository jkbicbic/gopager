package gopager

// Pagination models the paginator
type Pagination struct {
	HasPrev                 bool
	HasNext                 bool
	NearLast                bool
	FirstPage               uint
	LastPage                uint
	Prev                    uint
	Next                    uint
	Pages                   []uint
	CurrentPage             uint
	PagerStart              uint
	totalPages, pagerLength uint
}

// New builds a new pagination
func New(totalPages, currentPage, pagerLength uint) *Pagination {
	return &Pagination{
		totalPages:  totalPages,
		CurrentPage: currentPage,
		pagerLength: pagerLength,
	}
}

// Paginate accepts the totalRows, currentPage, and paginatorLimit to create a paginator for a page
func (p *Pagination) Paginate() {
	// new AtPagination struct

	p.FirstPage = 1
	p.LastPage = p.totalPages

	// check if current page overflows totalpage
	if p.CurrentPage > p.totalPages {
		p.CurrentPage = p.totalPages
	}

	// check if has prev
	if p.CurrentPage > 1 {
		p.HasPrev = true
	}

	// check if has next
	if p.CurrentPage < p.totalPages {
		p.HasNext = true
	}

	// for next page add 1
	// to currentPage
	p.Next = p.CurrentPage + 1
	if !p.HasNext {
		p.Next = p.totalPages
	}

	// for previous page subtract 1
	// to currentPage
	p.Prev = p.CurrentPage - 1
	if !p.HasPrev {
		p.Prev = 1
	}

	// check currentPage and generate pages
	// with different calculations
	p.PagerStart = 1
	if p.CurrentPage != 1 {
		p.PagerStart = p.CurrentPage - (p.pagerLength / 2)
	}

	// check to make sure
	// it will not overflow
	lastGroup := (p.totalPages - (p.pagerLength - 1))
	if p.PagerStart >= lastGroup && p.PagerStart != 1 {
		p.PagerStart = lastGroup
	}

	// build pages
	for i := p.PagerStart; i <= (p.PagerStart + (p.pagerLength - 1)); i++ {
		p.Pages = append(p.Pages, i)
	}

	// check if lastpage is
	// in pages
	if p.Pages[p.pagerLength-1] == p.totalPages {
		p.NearLast = true
	}
}
