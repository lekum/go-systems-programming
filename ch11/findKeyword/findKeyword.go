package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
)

type Data struct {
	URL     string
	Keyword string
	Times   int
	Error   error
}

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("Usage %s keyword file", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	var f *os.File
	var keyword string
	filename := ""

	if len(os.Args) == 2 {
		f = os.Stdin
		keyword = os.Args[1]
	} else {
		keyword = os.Args[1]
		filename = os.Args[2]
		fileHandler, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Error opening file %s: %s\n", filename, err)
			os.Exit(1)
		}
		f = fileHandler
	}
	defer f.Close()

	values := make(chan Data, len(os.Args[1:]))

	scanner := bufio.NewScanner(f)
	count := 0

	for scanner.Scan() {
		count++
		go func(URL string) {
			processPage(URL, keyword, values)
		}(scanner.Text())
	}
	monitor(values, count)
}

func monitor(values <-chan Data, count int) {
	for i := 0; i < count; i++ {
		x := <-values
		if x.Error == nil {
			fmt.Printf("\t%s\t", x.Keyword)
			fmt.Printf("\t%d\t in\t%s\n", x.Times, x.URL)
		} else {
			fmt.Printf("\t%s\n", x.Error)
		}
	}
}

func processPage(myUrl, keyword string, out chan<- Data) {
	var err error
	times := 0

	URL, err := url.Parse(myUrl)
	if err != nil {
		out <- Data{URL: myUrl, Keyword: keyword, Times: 0, Error: err}
		return
	}

	c := &http.Client{}
	request, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		out <- Data{URL: myUrl, Keyword: keyword, Times: 0, Error: err}
		return
	}

	httpData, err := c.Do(request)
	if err != nil {
		out <- Data{URL: myUrl, Keyword: keyword, Times: 0, Error: err}
		return
	}

	bodyHTML := ""

	var buffer [1024]byte

	reader := httpData.Body
	for {
		n, err := reader.Read(buffer[0:])
		if err != nil {
			break
		}
		bodyHTML += string(buffer[0:n])
	}

	regExpr := keyword

	r := regexp.MustCompile(regExpr)
	matches := r.FindAllString(bodyHTML, -1)
	times += len(matches)

	newValue := Data{URL: myUrl, Keyword: keyword, Times: times, Error: nil}
	out <- newValue

}
