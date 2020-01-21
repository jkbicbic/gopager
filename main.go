package main

import (
	"log"

	gopager "github.com/jkbicbic/gopager/gopager"
)

func main() {

	count := 40      // total pages is total rows divided by page limit
	pageLimit := 5   // total pages is total rows divided by page limit
	currentPage := 8 // the current page
	pagerLength := 3 // the length of your paginator e.g. [4 5 6] for size 3 [4 5 6 7 8] for size 5

	// builds a new pagination
	p := gopager.New(count, pageLimit, currentPage, pagerLength)
	p.Paginate()

	// prints the Pagination struct
	log.Println(p)
}
