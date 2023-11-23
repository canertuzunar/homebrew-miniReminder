package services

import "fmt"

var (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Purple = "\033[35m"
	Cyan   = "\033[36m"
	Gray   = "\033[37m"
	White  = "\033[97m"
)

func red(s string) string {
	return fmt.Sprintf("%s%s%s", Red, s, Reset)
}

func green(s string) string {
	return fmt.Sprintf("%s%s%s", Green, s, Reset)
}
func yellow(s string) string {
	return fmt.Sprintf("%s%s%s", Yellow, s, Reset)
}
func blue(s string) string {
	return fmt.Sprintf("%s%s%s", Blue, s, Reset)
}
func purple(s string) string {
	return fmt.Sprintf("%s%s%s", Purple, s, Reset)
}
func cyan(s string) string {
	return fmt.Sprintf("%s%s%s", Cyan, s, Reset)
}
func gray(s string) string {
	return fmt.Sprintf("%s%s%s", Gray, s, Reset)
}
func white(s string) string {
	return fmt.Sprintf("%s%s%s", White, s, Reset)
}
