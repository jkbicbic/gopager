# Gopager

[![GoDoc](https://godoc.org/github.com/jkbicbic/gopager?status.svg)](https://godoc.org/github.com/jkbicbic/gopager) [![Go Report Card](https://goreportcard.com/badge/github.com/jkbicbic/gopager)](https://goreportcard.com/report/github.com/jkbicbic/gopager)
[![Maintainability](https://api.codeclimate.com/v1/badges/785ba623d085509dee21/maintainability)](https://codeclimate.com/github/jkbicbic/gopager/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/785ba623d085509dee21/test_coverage)](https://codeclimate.com/github/jkbicbic/gopager/test_coverage)

a small library that calculates the first, last, prev, next, and slice of pages for your pagination.

### Example Pagination

Below is how your pagination could look.

![alt text](https://github.com/jkbicbic/gopager/blob/master/img/sample.png)

## Installation

Adding the package from github

```
go get github.com/jkbicbic/gopager
```

Importing the package in your project

```GO
import "github.com/jkbicbic/gopager"
```

## Usage

Gopager accepts 3 uint arguments

```GO
var totalPages uint = 40 // total pages is total rows divided by page limit
var currentPage uint = 5 // the current page
var pagerLength uint = 3 // the length of your paginator e.g. [4 5 6] for size 3 [4 5 6 7 8] for size 5

```

Build your pagination using `gopager.New()`

```GO
p := gopager.New(totalPages, currentPage, pagerLength)  // creates a new instance of pagination
p.Paginate()                                                // builds the pagination
```

You can pass this in a `map[string]interface{}{}` to be used in your template

```GO
data := map[string]interface{}{}
data["paginator"] = p
```

sample template usage

```HTML
{{if .paginator.HasPrev}}
<li class="first">
  <a href="?page={{.paginator.FirstPage}}"> First </a>
</li>
<li class="prev">
  <a rel="prev" href="?page={{.paginator.prev}}">{{.paginator.Prev}}</a>
</li>
{{end}}
{{if .paginator.HasPrev}}
<li class="page">
  <a href="?page={{.paginator.FirstPage}}">{{.paginator.FirstPage}}</a>
</li>
{{end}}
{{range $i := p.Pages}}
{{if eq $i $currentPage}}
<li class="page current">
  {{$i}}
</li>
{{else}}
<li class="page">
  <a href="?page={{$i}}">{{$i}}</a>
</li>
{{end}}
{{end}}
{{if not .paginator.NearLast}}
<li class="page">
  <a href="?page={{.paginator.LastPage}}">{{.paginator.LastPage}}</a>
</li>
{{end}}
{{if .paginator.HasNext}}
<li class="next">
  <a rel="next" href="?page={{.paginator.Next}}"> Next </a>
</li>
<li class="last">
  <a href="?&page={{.paginator.LastPage}}"> Last </a>
</li>
{{end}}
```



Gopager struct for reference

```GO
ttype Pagination struct {
	HasPrev                 bool
	HasNext                 bool
	NearLast                bool
	FirstPage               uint
	LastPage                uint
	Prev                    uint
	Next                    uint
	Pages                   []uint
	CurrentPage             uint
	totalPages, pagerLength uint
}
```
## Contribute

create issues [here](https://github.com/jkbicbic/gopager/issues/new)






