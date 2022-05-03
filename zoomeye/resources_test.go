package zoomeye

import (
	"testing"
)

func TestGetResoouceInfo(t *testing.T) {
	z := NewEnvZoomEyeClient()

	info, err := z.GetResourcesInfo()
	if err != nil {
		t.Logf("Can't Get resource info : %v\n", err)
	}

	t.Logf("Get resource ifnormation : %v\n", info)
}
