package main

import (
	"log"

	gopager "github.com/jkbicbic/gopager/gopager"
)

func main() {

	totalRows := 40     // total rows of your query
	currentPage := 5    // the current page
	paginatorLimit := 3 // the size of your paginator e.g. [4 5 6] for size 3 [4 5 6 7 8] for size 5

	// builds a new pagination
	p := gopager.New(totalRows, currentPage, paginatorLimit)
	p.Paginate()

	// prints the Pagination struct
	log.Println(p)
}
