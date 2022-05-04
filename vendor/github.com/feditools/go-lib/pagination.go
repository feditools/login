package lib

import (
	"net/url"
	"strconv"
)

// GetPaginationFromURL returns a page and item count derived from the url query.
// returns true if count is present in query.
func GetPaginationFromURL(u *url.URL, defaultCount int) (int, int, bool) {
	// get display count
	count := defaultCount
	if qCount, ok := u.Query()["count"]; ok {
		if len(qCount[0]) >= 1 {
			uCount, err := strconv.ParseUint(qCount[0], 10, 64)
			if err == nil {
				count = int(uCount)
			}
		}
	}

	// get display page
	page := 1
	if qPage, ok := u.Query()["page"]; ok {
		if len(qPage[0]) >= 1 {
			uPage, err := strconv.ParseUint(qPage[0], 10, 64)
			if err == nil {
				page = int(uPage)
			}
		}
	}

	return page, count, defaultCount != count
}
