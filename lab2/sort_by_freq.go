package main

type custom []rune

func (s custom) Len() int {
	return len(s)
}
func (s custom) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s custom) Less(i, j int) bool {
	return items[s[i]] > items[s[j]]
}
