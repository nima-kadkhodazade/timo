package main

const (
	Reset  = "\033[0m"
	Yellow = "\033[33m"
	Green  = "\033[32m"
	Red    = "\033[31m"
)

func StatusStyle(status string) (string, string) {
	switch status {
		case "done":
			return Green + "●" + Reset, Green + status + Reset
		case "in-progress":
			return Yellow + "◌" + Reset, Yellow + status + Reset
		case "todo":
			return Red + "○" + Reset, Red + status + Reset

		default:
			return "", status
	}
}