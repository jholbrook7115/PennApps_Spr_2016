package series_now

import "testing"

func TestBasename(t *testing.T) {
	s := "http://www.tvguide.com/tvshows/teen-mom-og/765609/"
	str := "teen-mom-og"
	str_got := basename(s)
	if str_got != str {
		t.Errorf("Expected %v, got %v\n", str, str_got)
	}
}

