# Gopager

[![GoDoc](https://godoc.org/github.com/jkbicbic/gopager?status.svg)](https://godoc.org/github.com/jkbicbic/gopager) [![Go Report Card](https://goreportcard.com/badge/github.com/jkbicbic/gopager)](https://goreportcard.com/report/github.com/jkbicbic/gopager)
[![Maintainability](https://api.codeclimate.com/v1/badges/785ba623d085509dee21/maintainability)](https://codeclimate.com/github/jkbicbic/gopager/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/785ba623d085509dee21/test_coverage)](https://codeclimate.com/github/jkbicbic/gopager/test_coverage)

a small library for your golang project pagination needs.

### Example Pagination

Below is how your pagination could look. Please note that any styles that you see here are not included, you and you alone have the power to design your own pagination this library just provides the calculations.

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

Gopager accepts 4 arguments

```GO
var count int = 40 // totalrows of your query
var pageLimit int = 3 // the number of rows for each page
var currentPage int = 5 // the current page
var pagerLength int = 3 // the length of your paginator e.g. [4 5 6] for size 3 [4 5 6 7 8] for size 5

```

Build your pagination using `gopager.New()`

```GO
p := gopager.New(count, pageLimit, currentPage, pagerLength)  // creates a new instance of pagination
p.Paginate()                                                  // builds the pagination
```

You can pass this in a `map[string]interface{}{}` to be used in your template

```GO
data := map[string]interface{}{}
data["paginator"] = p
```

sample template usage

```HTML
<nav class="pagination" role="navigation" aria-label="pagination">
  <ul class="pagination-list">
    {{if .paginator.HasPrev}}
    <li>
      <a class="pagination-link" href="?page={{.paginator.FirstPage}}"> First </a>
    </li>
    <li>
      <a class="pagination-link" rel="prev" href="?page={{.paginator.Prev}}"> Prev </a>
    </li>
    <li>
      <a class="pagination-link" href="?page={{.paginator.FirstPage}}">{{.paginator.FirstPage}}</a>
    </li>
    {{end}}
    {{$currentPage := .paginator.CurrentPage}}
    {{range $i := .paginator.Pages}}
    {{if eq $i $currentPage}}
    <li>
      <a class="pagination-link is-current">{{$i}}</a>
    </li>
    <li>
      <span class="pagination-ellipsis">&hellip;</span>
    </li>
    {{else}}
    <li>
      <a class="pagination-link" href="?page={{$i}}">{{$i}}</a>
    </li>
    {{end}}
    {{end}}
    {{if not .paginator.NearLast}}
    <li>
      <span class="pagination-ellipsis">&hellip;</span>
    </li>
    <li>
      <a class="pagination-link" href="?page={{.paginator.TotalPages}}">{{.paginator.TotalPages}}</a>
    </li>
    {{end}}
    {{if .paginator.HasNext}}
    <li>
      <a class="pagination-link" rel="next" href="?page={{.paginator.Next}}"> Next </a>
    </li>
    <li>
      <a class="pagination-link" href="?&page={{.paginator.TotalPages}}"> Last </a>
    </li>
    {{end}}
  </ul>
</nav>
```



Gopager struct for reference

```GO
type Pagination struct {
	HasPrev                 bool
	HasNext                 bool
	NearLast                bool
	Firstpage               int
	Prev                    int
	Next                    int
	Pages                   []int
	CurrentPage             int
	PagerStart              int
	TotalPages, pagerLength int
}
```
## Contribute

create issues [here](https://github.com/jkbicbic/gopager/issues/new)






