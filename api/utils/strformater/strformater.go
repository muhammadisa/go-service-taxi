package strformater

// ToQueryString convert string for where like query
func ToQueryString(str string) string {
	return "%" + str + "%"
}
