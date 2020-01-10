package gopager_test

import (
	"fmt"
	"testing"

	gopager "github.com/jkbicbic/gopager/gopager"
)

func TestGoPager(t *testing.T) {
	var tests = []struct {
		totalPages, currentPage, pagerLength uint
	}{
		{20, 3, 10},
		{1, 3, 10},
		{40, 200, 20},
		{40, 1, 20},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d,%d", tt.totalPages, tt.currentPage, tt.pagerLength)
		t.Run(testname, func(t *testing.T) {
			p := gopager.New(tt.totalPages, tt.currentPage, tt.pagerLength)
			p.Paginate()
			if p.CurrentPage == 1 && p.HasPrev == true {
				t.Errorf("current page is already %d, but HasPrev returned %v", p.CurrentPage, p.HasPrev)
			}
			if p.CurrentPage != 1 && p.HasPrev == false {
				t.Errorf("current page is %d, but HasPrev returned %v", p.CurrentPage, p.HasPrev)
			}
			if p.CurrentPage == tt.totalPages && p.HasNext == true {
				t.Errorf("current page is already %d, but HasNext returned %v", p.CurrentPage, p.HasPrev)
			}
			if p.CurrentPage != tt.totalPages && p.HasNext == false {
				t.Errorf("current page is %d, but HasNext returned %v", p.CurrentPage, p.HasPrev)
			}
			if p.FirstPage != 1 {
				t.Errorf("First page got %d, want %d", p.FirstPage, 1)
			}
			if p.LastPage != tt.totalPages {
				t.Errorf("Last page got %d, want %d", p.LastPage, tt.totalPages)
			}
			if p.Prev != (p.CurrentPage-1) && p.CurrentPage != 1 {
				t.Errorf("Previous page got %d, want %d", p.Prev, (p.CurrentPage - 1))
			}
			if p.Prev < 1 {
				t.Errorf("Previous page is %d, should not be less than 1", p.Prev)
			}
			if p.PagerStart < 1 {
				t.Errorf("start of pager is %d, should not be less than 1", p.PagerStart)
			}
			if p.Next != (p.CurrentPage+1) && p.CurrentPage != tt.totalPages {
				t.Errorf("Next page got %d, want %d", p.Next, (p.CurrentPage + 1))
			}
			if p.Next > tt.totalPages {
				t.Errorf("Next page should not be greater than %d", tt.totalPages)
			}
		})
	}
}
