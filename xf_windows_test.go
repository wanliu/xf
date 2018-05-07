// +build windows

package xf

import "testing"

func TestWindowsXF(t *testing.T) {
	if err := MSPLogin("appid = 58c51121, work_dir = ."); err != nil {
		t.Fatalf("login failed", err)
	}

}
