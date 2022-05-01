package zoomeye

import (
	"testing"
)

func TestDomainSearch(t *testing.T) {
	z := NewEnvZoomEyeClient()

	query := "baidu.com"
	ty := 0
	page := 1

	result, err := z.DomainSearch(query, ty, page)
	if err != nil {
		t.Logf("Can't Get Domain Search result from DomainSearch : %v\n", err)
	}

	t.Logf("domain search result : %v\n", result)
}
