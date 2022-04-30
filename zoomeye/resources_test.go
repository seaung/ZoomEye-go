package zoomeye

import (
	"testing"
)

func TestGetResource(t *testing.T) {
	z := NewEnvZoomEyeClient()

	result := z.GetResourcesInfo()

	t.Logf("resource info : %v\n", result)
}
