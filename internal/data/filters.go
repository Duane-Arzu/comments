// Filename: internal/data/filters.go
package data

import (
<<<<<<< HEAD
	"strings"

	"github.com/Duane-Arzu/comments/internal/validator"
)

// The Filters type will contain fields related to pagination
// and eventually the fields related to sorting.
type Filters struct {
	Page         int // Which page number the client wants.
	PageSize     int // How many records per page.
	Sort         string
	SortSafeList []string // allowed sort fields

}

type Metadata struct {
	CurrentPage  int `json:"current_page,omitempty"`
	PageSize     int `json:"page_size,omitempty"`
	FirstPage    int `json:"first_page,omitempty"`
	LastPage     int `json:"last_page,omitempty"`
	TotalRecords int `json:"total_records,omitempty"`
}

// ValidateFilters checks the validity of pagination parameters.
=======
	"github.com/Duane-Arzu/comments/internal/validator"
	_ "github.com/Duane-Arzu/comments/internal/validator"
)

// The Filters type will contain the fields related to pagination
// and eventually the fields related to sorting.
type Filters struct {
	Page     int // which page number does the client want
	PageSize int // how records per page
}

// Next we validate page and PageSize
// We follow the same approach that we used to validate a Comment
>>>>>>> e552bbb7e42555c25294bebb63c793b53c7b49ef
func ValidateFilters(v *validator.Validator, f Filters) {
	v.Check(f.Page > 0, "page", "must be greater than zero")
	v.Check(f.Page <= 500, "page", "must be a maximum of 500")
	v.Check(f.PageSize > 0, "page_size", "must be greater than zero")
	v.Check(f.PageSize <= 100, "page_size", "must be a maximum of 100")
<<<<<<< HEAD
	v.Check(validator.PermittedValue(f.Sort, f.SortSafeList...), "sort",
		"invalid sort value")

}

func (f Filters) sortColumn() string {
	for _, safeValue := range f.SortSafeList {
		if f.Sort == safeValue {
			return strings.TrimPrefix(f.Sort, "-")
		}
	}
	// don't allow the operation to continue
	// if case of SQL injection attack
	panic("unsafe sort parameter: " + f.Sort)
}

func (f Filters) sortDirection() string {
	if strings.HasPrefix(f.Sort, "-") {
		return "DESC"
	}
	return "ASC"
}

// limit returns the number of records per page.
=======
}

// calculate how many records to send back
>>>>>>> e552bbb7e42555c25294bebb63c793b53c7b49ef
func (f Filters) limit() int {
	return f.PageSize
}

<<<<<<< HEAD
// offset calculates the number of records to skip for pagination.
func (f Filters) offset() int {
	return (f.Page - 1) * f.PageSize
}

// calculateMetaData generates pagination metadata.
func calculateMetaData(totalRecords int, currentPage int, pageSize int) Metadata {
	if totalRecords == 0 {
		return Metadata{}
	}

	return Metadata{
		CurrentPage:  currentPage,
		PageSize:     pageSize,
		FirstPage:    1,
		LastPage:     (totalRecords + pageSize - 1) / pageSize,
		TotalRecords: totalRecords,
	}
}
=======
// calculate the offset so that we remember how many records have
// been sent and how many remain to be sent
func (f Filters) offset() int {
	return (f.Page - 1) * f.PageSize
}
>>>>>>> e552bbb7e42555c25294bebb63c793b53c7b49ef
