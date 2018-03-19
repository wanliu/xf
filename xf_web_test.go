package xf

import "testing"

func TestXFUI_Web(t *testing.T) {
	client := NewWebClient("5aaf5fe1", "629e95d1bd0846b289212576a26d3fb9")
	sess := client.NewSession("main", "user_01")

	if buf, err := sess.PostText("你好"); err != nil {
		t.Fatalf("error %s", err)
	} else {
		t.Logf("result: %s\n", buf)
	}
}
