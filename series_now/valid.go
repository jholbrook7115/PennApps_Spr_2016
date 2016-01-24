package series_now
/*func validate() {
	for _, value := range stack {
		
	}
}*/
// Example :
// [http://www.tvguide.com/tvshows/teen-mom-og/765609/, http://www.tvguide.com/tvshows/cops/191459/]
// should be
// [ [teen-mom-og: 765609], [cops: 191459/] ]
func basename(s string) string {
	s = s[:len(s) - 2]
	// Preseve everything before last '/'
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[:i]
			break
		}
	}
	// Discard everything passed to tvshows/
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	return s
}
