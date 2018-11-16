package csanitize_test

import (
	"testing"

	"github.com/yookoala/poeditcheck/lib/csanitize"
)

func TestGetTrims(t *testing.T) {
	left, right := csanitize.GetTrims("\n\t    \r\nHello \n\n World\n\t\r    \n", "\r\n\t ")
	if want, have := "\n\t    \r\n", left; want != have {
		t.Errorf("expected %#v, got %#v", want, have)
	}
	if want, have := "\n\t\r    \n", right; want != have {
		t.Errorf("expected %#v, got %#v", want, have)
	}
}

func TestGetPlaceholder(t *testing.T) {
	var results []string

	// non-positional placeholder
	results = csanitize.GetPlaceholder("abc %s def %d ghi %f hello")
	if want, have := 3, len(results); want != have {
		t.Errorf("expected %#v, got %#v", want, have)
		t.FailNow()
	}
	if want, have := "%s", results[0]; want != have {
		t.Errorf("expected %#v, got %#v", want, have)
	}
	if want, have := "%d", results[1]; want != have {
		t.Errorf("expected %#v, got %#v", want, have)
	}
	if want, have := "%f", results[2]; want != have {
		t.Errorf("expected %#v, got %#v", want, have)
	}

	// positional placeholder
	results = csanitize.GetPlaceholder("abc %1$s def %2$d ghi %30$f hello")
	if want, have := 3, len(results); want != have {
		t.Errorf("expected %#v, got %#v", want, have)
		t.FailNow()
	}
	if want, have := "%1$s", results[0]; want != have {
		t.Errorf("expected %#v, got %#v", want, have)
	}
	if want, have := "%2$d", results[1]; want != have {
		t.Errorf("expected %#v, got %#v", want, have)
	}
	if want, have := "%30$f", results[2]; want != have {
		t.Errorf("expected %#v, got %#v", want, have)
	}
}
