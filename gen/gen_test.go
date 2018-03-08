package gen

import (
	"testing"

	"github.com/kr/pretty"
)

func TestGen(t *testing.T) {
	tags, err := ParseTags("../TAGS")
	if err != nil {
		t.Fatalf("parse tags failed, error: %s", err)
	} else {
		t.Logf("tags:\n% #v", pretty.Formatter(tags))
	}

	for _, tag := range tags {
		t.Log(tag.Source())
	}
}
