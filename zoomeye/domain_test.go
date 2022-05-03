package zoomeye

import (
	"testing"
)

func TestDomainSearch(t *testing.T) {
	z := NewEnvZoomEyeClient()

	domains, err := z.DomainSearch("baidu.com", 0, 1)
	if err != nil {
		t.Logf("Can't Get domain info : %v\n", err)
	}

	t.Logf("Get Domain info : %v\n", domains)
}
