package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("data.md")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// cache1 := ""
	// cache2 := ""

	re, err := regexp.Compile(`[^0-9.]`)
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	wc := ""
	values := []string{}
	// output := []byte{}

	for i, v := range f {
		s := string(v)
		if s == "[" {
			count++
		}
		if count == 2 {
			wc = wc + s

		}
		if len(wc) == 15 {
			str := re.ReplaceAllString(wc, "")
			values = append(values, str)
			count = 0
			wc = ""
			// str1 := string(f[i-10 : i-4])
			fmt.Println("VALUE ==> ", string(f[i-10:i-4]))
			current := string(f[i-10 : i-4])
			strings.Split(current, "")

			// f = append(f[i-10 : i-4], byte("") )
			copy(f[i-10:i-4], "123456789")
			// res := bytes.Replace(f, []byte("replaceme"), v, -1)
			// fmt.Println("INDEX ==> ", i)
			// output = res
		}
		// fmt.Println(string(v))
	}

	ioutil.WriteFile("test.md", f, 0644)
	// fmt.Println(values)

	// output := bytes.Replace(input, []byte("replaceme"), []byte("ok"), -1)

	// if err = ioutil.WriteFile("modified.txt", output, 0666); err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

}

func replaceNumbers(s string) string {
	for i := range s {
		if i == 0 {
			s = "0" + s
		}
		if i == 2 {
			s = s[:i+1] + "0" + s[i+1:]
		}
		if i == 4 {
			s = s[:i+2] + "0" + s[i+2:]
		}
	}
	return s
}
