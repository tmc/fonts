package fonts

import (
	"fmt"

	"testing"
)

func TestListingAndLoading(t *testing.T) {
	fonts, err := List()
	if err != nil {
		t.Error("List() failed:", err)
	}
	if len(fonts) == 0 {
		t.Error("No fonts found!")
	} else {
		for _, f := range fonts {
			fmt.Println(f.Name)
		}
		_, err := Load(fonts[0].Name)
		if err != nil {
			t.Error("Error loading", fonts[0].Name, "-", err)
		}
	}
}

func TestBadPaths(t *testing.T) {
	_, e := Load("invalid font name")
	if e == nil {
		t.Error("expected error")
	}
}
