package zoomeye

import (
	"testing"
)

func TestLogin(t *testing.T) {
	z := NewEnvZoomEyeClient()

	l, err := z.Login()
	if err != nil {
		t.Logf("Can't Get user token, %v\n", err)
	}

	t.Logf("Get User Token information : %v\n", l)
}
