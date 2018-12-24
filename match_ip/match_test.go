package match

import "testing"

func TestRegexpMatch(t *testing.T) {

	ipCheck := RegexpMatch("ip", "127.0.0.1")
	if !ipCheck {
		t.Error(ipCheck)
	}
	t.Log(ipCheck)
}
