package paginator

// Pagination models the paginator
type Pagination struct {
	HasPrev                           bool
	HasNext                           bool
	NearLast                          bool
	FirstPage                         int
	LastPage                          int
	Prev                              int
	Next                              int
	Pages                             []int
	totalRows, currentPage, pageLimit int
}

// New builds a new pagination
func New(totalRows, currentPage, pageLimit int) *Pagination {
	return &Pagination{
		totalRows:   totalRows,
		currentPage: currentPage,
		pageLimit:   pageLimit,
	}
}

// Paginate accepts the totalRows, currentPage, and paginatorLimit to create a paginator for a page
func (p *Pagination) Paginate() {
	// new AtPagination struct
	var pageCalc int
	totalPages := p.totalRows / p.pageLimit

	p.FirstPage = 1
	p.LastPage = totalPages

	// check if has prev
	if p.currentPage > p.pageLimit-1 {
		p.HasPrev = true
	}

	// check if has next
	if p.currentPage < totalPages {
		p.HasNext = true
	}

	// for next page add 1
	// to currentPage
	p.Next = p.currentPage + 1
	if p.Next > totalPages {
		p.Next = totalPages
	}

	// for previous page subtract 1
	// to currentPage
	p.Prev = p.currentPage - 1
	if p.Prev <= 0 {
		p.Prev = 1
	}

	// check currentPage and generate pages
	// with different calculations
	if p.currentPage == 1 {
		pageCalc = p.currentPage
	} else {
		pageCalc = p.currentPage - (p.pageLimit / 2)
	}

	// check to make sure
	// it will not overflow
	lastGroup := (totalPages - (p.pageLimit - 1))
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
	for i := pageCalc; i <= (pageCalc + (p.pageLimit - 1)); i++ {
		p.Pages = append(p.Pages, i)
	}

	// check if lastpage is
	// in pages
	if p.Pages[p.pageLimit-1] == totalPages {
		p.NearLast = true
	}
}
