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
	totalPages, pagerLength uint
}

// New builds a new pagination
func New(totalRows, currentPage, pagerLength uint) *Pagination {
	return &Pagination{
		totalPages:  totalRows,
		CurrentPage: currentPage,
		pagerLength: pagerLength,
	}
}

// Paginate accepts the totalRows, currentPage, and paginatorLimit to create a paginator for a page
func (p *Pagination) Paginate() {
	// new AtPagination struct
	var pageCalc uint

	p.FirstPage = 1
	p.LastPage = p.totalPages

	// check if current page overflows totalpage
	if p.CurrentPage > p.totalPages {
		p.CurrentPage = p.totalPages
	}

	// check if current page underflows totalpage
	if p.CurrentPage < 1 {
		p.CurrentPage = 1
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
	if p.Next > p.totalPages {
		p.Next = p.totalPages
	}

	// for previous page subtract 1
	// to currentPage
	p.Prev = p.CurrentPage - 1
	if p.Prev <= 0 {
		p.Prev = 1
	}

	// check currentPage and generate pages
	// with different calculations
	if p.CurrentPage == 1 {
		pageCalc = p.CurrentPage
	} else {
		pageCalc = p.CurrentPage - (p.pagerLength / 2)
	}

	// check to make sure
	// it will not overflow
	lastGroup := (p.totalPages - (p.pagerLength - 1))
	if pageCalc >= lastGroup && pageCalc != 1 {
		pageCalc = lastGroup
	}

	// check to make sure
	// pageCalc will not be less
	// than zero
	if pageCalc <= 0 {
		pageCalc = 1
	}

	// build pages
	for i := pageCalc; i <= (pageCalc + (p.pagerLength - 1)); i++ {
		p.Pages = append(p.Pages, i)
	}

	// check if lastpage is
	// in pages
	if p.Pages[p.pagerLength-1] == p.totalPages {
		p.NearLast = true
	}
}
