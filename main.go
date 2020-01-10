package main

import (
	"log"

	gopager "github.com/jkbicbic/gopager/gopager"
)

func main() {

	totalPages := 3  // total pages is total rows divided by page limit
	currentPage := 1 // the current page
	pagerLength := 1 // the length of your paginator e.g. [4 5 6] for size 3 [4 5 6 7 8] for size 5

	// builds a new pagination
	p := gopager.New(totalPages, currentPage, pagerLength)
	p.Paginate()

	// prints the Pagination struct
	log.Println(p)
}
