package gopager

import "math"

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
	TotalPages, pagerLength int
}

// New builds a new pagination
func New(count, pageLimit, currentPage, pagerLength int) *Pagination {
	return &Pagination{
		TotalPages:  int(math.Ceil(float64(count) / float64(pageLimit))),
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
	if p.CurrentPage < p.TotalPages {
		p.HasNext = true
	}

	// for next page add 1
	// to currentPage
	p.Next = p.CurrentPage + 1
	if p.Next > p.TotalPages {
		p.Next = p.TotalPages
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
	if mod == 0 && p.CurrentPage == 1 || mod == p.CurrentPage {
		p.PagerStart = 1
	} else {
		p.PagerStart = p.CurrentPage - (p.pagerLength / 2)
	}

	// build pages
	for i := 1; i <= p.pagerLength; i++ {
		if p.PagerStart <= p.TotalPages {
			p.Pages = append(p.Pages, p.PagerStart)
			p.PagerStart = p.PagerStart + 1
		}
	}

	p.NearLast = false
	// check if lastpage is
	// in pages
	for i := 0; i < len(p.Pages); i++ {
		if p.Pages[i] == p.TotalPages {
			p.NearLast = true
		}
	}
}
