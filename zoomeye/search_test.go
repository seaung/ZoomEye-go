package zoomeye

import (
	"testing"
)

func TestSearch(t *testing.T) {
	z := NewEnvZoomEyeClient()

	matches, err := z.Search("port:21%20city:beijing", "app,os", 1, "host")
	if err != nil {
		t.Logf("Get Matches error %v\n", err)
	}

	t.Logf("Get Matches : %v\n", matches)
}
