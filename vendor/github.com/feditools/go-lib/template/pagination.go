package template

import (
	"fmt"
	"math"
)

// Pagination is a pagination element that can be added to a webpage
type Pagination []PaginationNode

// PaginationNode is an element in a pagination element
type PaginationNode struct {
	Text string
	Icon string
	HRef string

	Active   bool
	Disabled bool
}

// PaginationConfig contains the config to construct pagination
type PaginationConfig struct {
	Count         int    // item count
	DisplayCount  int    // how many items to display per page
	HRef          string // href to add query to
	HRefCount     int    // count to include in the href, if 0 no count is added
	MaxPagination int    // the max number of pages to show
	Page          int    // current page
}

// MakePagination creates a pagination element from the provided parameters
func MakePagination(c *PaginationConfig) Pagination {
	displayItems := c.MaxPagination
	pages := int(math.Ceil(float64(c.Count) / float64(c.DisplayCount)))
	startingNumber := 1

	if pages < displayItems {
		// less than
		displayItems = pages
	} else if c.Page > pages-displayItems/2 {
		// end of the
		startingNumber = pages - displayItems + 1
	} else if c.Page > displayItems/2 {
		// center active
		startingNumber = c.Page - displayItems/2
	}

	var items Pagination

	// previous button
	prevItem := PaginationNode{
		Text: "Previous",
		Icon: "caret-left",
	}
	if c.Page == 1 {
		prevItem.Disabled = true
	} else if c.HRefCount > 0 {
		prevItem.HRef = fmt.Sprintf("%s?page=%d&count=%d", c.HRef, c.Page-1, c.HRefCount)
	} else {
		prevItem.HRef = fmt.Sprintf("%s?page=%d", c.HRef, c.Page-1)
	}
	items = append(items, prevItem)

	// add pages
	for i := 0; i < displayItems; i++ {
		newItem := PaginationNode{
			Text: fmt.Sprintf("%d", startingNumber+i),
		}

		if c.Page == startingNumber+i {
			newItem.Active = true
		} else if c.HRefCount > 0 {
			newItem.HRef = fmt.Sprintf("%s?page=%d&count=%d", c.HRef, startingNumber+i, c.HRefCount)
		} else {
			newItem.HRef = fmt.Sprintf("%s?page=%d", c.HRef, startingNumber+i)
		}

		items = append(items, newItem)
	}

	// next button
	nextItem := PaginationNode{
		Text: "Next",
		Icon: "caret-right",
	}
	if c.Page == pages {
		nextItem.Disabled = true
	} else if c.HRefCount > 0 {
		nextItem.HRef = fmt.Sprintf("%s?page=%d&count=%d", c.HRef, c.Page+1, c.HRefCount)
	} else {
		nextItem.HRef = fmt.Sprintf("%s?page=%d", c.HRef, c.Page+1)
	}
	items = append(items, nextItem)

	return items
}
