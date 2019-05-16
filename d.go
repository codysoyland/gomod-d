package d

import "fmt"

func Greet(from string, greeting string, exclaim string) {
	fmt.Printf("%s from %s%s\n", greeting, from, exclaim)
}
