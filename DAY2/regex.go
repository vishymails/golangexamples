package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(".com")

	fmt.Println(re.FindString("google.com"))
	fmt.Println(re.FindString("www.ibm.in"))
	fmt.Println(re.FindString("www.st.sg"))
	fmt.Println(re.FindString("www.mega.nz"))

	fmt.Println(re.FindStringIndex("google.com"))
	fmt.Println(re.FindStringIndex("www.ibm.in"))

	re1 := regexp.MustCompile("f([a-z]+)ing([a-z]+)")

	fmt.Println(re1.FindStringSubmatch("kjfkdjfkjfflyinghigh"))

}
