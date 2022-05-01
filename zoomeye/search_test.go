package zoomeye

import (
	"fmt"
	"net/url"
	"testing"
)

func TestSearchWithPath(t *testing.T) {
	z := NewEnvZoomEyeClient()

	query := "city:beijing%20app:DedeCMS"
	page := 1
	facets := "app,os"

	searchpath := fmt.Sprintf(hostSearchPath, query, page, facets)
	u, err := url.Parse(baseURL + searchpath)
	if err != nil {
		t.Logf("can't parse url in TestSearchWithPath : %v\n", err)
	}

	content, err := z.newRequest("GET", u, nil, nil)
	if err != nil {
		t.Logf("get content error : %v\n", err)
	}

	t.Logf("content ====> %v\n", string(content))
}

func TestSearch(t *testing.T) {
	z := NewEnvZoomEyeClient()

	query := "city:beijing%20app:DedeCMS"
	page := 1
	facets := "app,os"
	hosts, err := z.Search(query, facets, page, 0)

	if err != nil {
		t.Logf("host search err : %v \n", err)
	}

	t.Logf("hosts list result : %v\n", hosts)

	webs, err := z.Search(query, facets, page, 1)
	if err != nil {
		t.Logf("webs search err : %v\n", err)
	}

	t.Logf("webs list result : %v\n", webs)
}
