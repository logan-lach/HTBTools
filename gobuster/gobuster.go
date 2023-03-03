package main 

import (
	"fmt"
	"flag"
	"net/http"
	"strconv"
	"os"
	"bufio"
	"sync"
)



func endpoint(url string, epoint string, client *http.Client, ch chan<-string) {
	fmt.Println(url + "/" + epoint)
	if res, err := client.Get(url + "/" + epoint); err == nil {
		ch <- epoint + ", " + strconv.Itoa(res.StatusCode)
	} else {
		ch <- "Err"
	}
}

// func main() {
// 	client := &http.Client{}
// 	if res,err := client.Get("https://google.com"); err == nil {
// 		fmt.Println(res.StatusCode)
// 	}
// }

func main() {
	urlFlag := flag.String("u", "", "Url to expose endpoints")
	wlFlag := flag.String("w", "", "Path to file of wordlist")
	channel := make(chan string)
	client := &http.Client{}
	var wg sync.WaitGroup

	flag.Parse()

	fmt.Printf("String 1 %s, String 2 %s", *urlFlag, *wlFlag)

	file, err := os.Open(*wlFlag)
	if err != nil {
		fmt.Errorf("%s", "Could not find file")
	}
	scanner := bufio.NewScanner(file)
	fmt.Println("Succesful finding the file")
	defer file.Close()

	for scanner.Scan() {
		fmt.Println("Scanned a line")
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Starting a function call")
			endpoint(*urlFlag, scanner.Text(), client, channel)
		}()
	}

	wg.Wait()
	close(channel)

	fmt.Println(<-channel)
	for v := range channel {
		fmt.Println(v)
	}





	


}