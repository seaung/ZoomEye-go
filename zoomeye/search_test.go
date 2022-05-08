package zoomeye

import (
	"testing"
)

func TestSearch(t *testing.T) {
	z := NewEnvZoomEyeClient()

	matches, err := z.HostSearch("port:21%20city:beijing", "app,os", 1)
	if err != nil {
		t.Logf("Get Matches error %v\n", err)
	}

	t.Logf("Get Matches : %v\n", matches.Code)

	for _, host := range matches.Matches {
		t.Logf("Jarm : %s\n", host.Jarm)
		t.Logf("ico md5 : %s\n", host.Ico.Md5)
	}
}
