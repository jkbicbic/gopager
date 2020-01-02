package main

import (
	"log"

	"github.com/jkbicbic/go-paginator/paginator"
)

func main() {

	// total rows of your query
	totalRows := 40
	// the current page
	currentPage := 5
	// the size of your paginator that you will show
	// e.g. [4 5 6] for size 3 [4 5 6 7 8] for size 5
	paginatorLimit := 3

	// builds a new pagination
	p := paginator.New(totalRows, currentPage, paginatorLimit)
	p.Paginate()

	// prints the Pagination struct
	log.Println(p)
}
