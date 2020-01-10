package gopager

// Pagination models the paginator
type Pagination struct {
	HasPrev                 bool
	HasNext                 bool
	NearLast                bool
	Prev                    int
	Next                    int
	Pages                   []int
	CurrentPage             int
	PagerStart              int
	totalPages, pagerLength int
}

// New builds a new pagination
func New(totalPages, currentPage, pagerLength int) *Pagination {
	return &Pagination{
		totalPages:  totalPages,
		CurrentPage: currentPage,
		pagerLength: pagerLength,
	}
}

// Paginate ...
func (p *Pagination) Paginate() {

	// check if has prev
	p.HasPrev = false
	if p.CurrentPage > p.pagerLength-1 {
		p.HasPrev = true
	}

	// check if has next
	p.HasNext = false
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
	mod := p.CurrentPage % p.pagerLength
	if mod == 0 || p.CurrentPage == 1 {
		p.PagerStart = p.CurrentPage
	} else if mod == 1 {
		p.PagerStart = p.CurrentPage - mod
	} else {
		p.PagerStart = p.CurrentPage - (mod - 1)
	}

	// build pages
	for i := 1; i <= p.pagerLength; i++ {
		if p.PagerStart <= p.totalPages {
			p.Pages = append(p.Pages, p.PagerStart)
			p.PagerStart = p.PagerStart + 1
		}
	}

	p.NearLast = false
	// check if lastpage is
	// in pages
	for i := 0; i < len(p.Pages); i++ {
		if p.Pages[i] == p.totalPages {
			p.NearLast = true
		}
	}
}
