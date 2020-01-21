package gopager_test

import (
	"fmt"
	"testing"

	gopager "github.com/jkbicbic/gopager/gopager"
)

func TestGoPager(t *testing.T) {
	var tests = []struct {
		count, pagelimit, currentPage, pagerLength int
	}{
		{20, 50, 3, 10},
		{1, 9, 3, 10},
		{40, 5, 200, 20},
		{40, 14, 1, 20},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d,%d,%d", tt.count, tt.pagelimit, tt.currentPage, tt.pagerLength)
		t.Run(testname, func(t *testing.T) {
			p := gopager.New(tt.count, tt.pagelimit, tt.currentPage, tt.pagerLength)
			p.Paginate()
			if p.CurrentPage == 1 && p.HasPrev == true {
				t.Errorf("current page is already %d, but HasPrev returned %v", p.CurrentPage, p.HasPrev)
			}
			if p.CurrentPage != 1 && p.HasPrev == false {
				t.Errorf("current page is %d, but HasPrev returned %v", p.CurrentPage, p.HasPrev)
			}
			if p.CurrentPage == p.TotalPages && p.HasNext == true {
				t.Errorf("current page is already %d, but HasNext returned %v", p.CurrentPage, p.HasPrev)
			}
			if p.CurrentPage != p.TotalPages && p.HasNext == false {
				t.Errorf("current page is %d, but HasNext returned %v", p.CurrentPage, p.HasPrev)
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
			if p.Next != (p.CurrentPage+1) && p.CurrentPage != p.TotalPages {
				t.Errorf("Next page got %d, want %d", p.Next, (p.CurrentPage + 1))
			}
			if p.Next > p.TotalPages {
				t.Errorf("Next page should not be greater than %d", p.TotalPages)
			}
		})
	}
}
