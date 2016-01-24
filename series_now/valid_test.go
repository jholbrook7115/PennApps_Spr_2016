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

func TestValidate(t *testing.T) {
	stack := []string{"http://www.tvguide.com/tvshows/teen-mom-og/765609/",
		"http://www.tvguide.com/tvshows/cops/191459/",
		"http://www.tvguide.com/tvshows/figure-skating/309372/",
		"http://www.tvguide.com/tvshows/mythbusters/191668/"}
	stack_exp := []string{
		"teen-mom-og",
		"cops",
		"figure-skating",
		"mythbusters"}
	stack_got := validate(stack)
	i := 0
	for _, value := range stack_got {
		if value != stack_exp[i] {
			t.Errorf("Expected %v, got %v\n", stack_exp[i], value)
		}
		i++
	}
}
