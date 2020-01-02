# Gopager

a small library that handles pagination

### Installation

Adding the package from github

```
go get github.com/jkbicbic/gopager
```

Importing the package in your project

```GO
import "github.com/jkbicbic/gopager"
```

### Usage

Gopager accepts 3 arguments

```GO
totalRows := 40     // total rows of your query
currentPage := 5    // the current page
paginatorLimit := 3 // the size of your paginator e.g. [4 5 6] for size 3 [4 5 6 7 8] for size 5
```

Build your pagination using `gopager.New()`

```GO
p := gopager.New(totalRows, currentPage, paginatorLimit)  // creates a new instance of pagination
p.Paginate() // builds the pagination
```

You can pass this in a `map[string]interface{}{}` to be used in your template

```GO
data := map[string]interface{}{}
data["paginator"] = p
```

Gopager struct for reference

```GO
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
```

### Contributing

create issues [here](https://github.com/jkbicbic/gopager/issues/new)






