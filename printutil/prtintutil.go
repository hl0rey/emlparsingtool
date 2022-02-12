package printutil

import "fmt"

const (
	INFO  = 0
	GOOD  = 1
	ERROR = 2
)

func Prettyprint(text string, level int) {

	if level == INFO {
		fmt.Println("[*] " + text)
	} else if level == GOOD {
		fmt.Println("[!] " + text)
	} else if level == ERROR {
		fmt.Println("[X] " + text)
	}

}
