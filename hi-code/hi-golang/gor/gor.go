/*
Copyright 2022 The https://github.com/anigkus Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package gor

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func Main() {

	fmt.Println("Gor...")

	playAndPause()

	goUnordered()

	goOrdered()

	goWaitGroupUnordered()
}

func playAndPause() {
	var wg sync.WaitGroup
	wg.Add(1)
	command := make(chan string)
	go routine(command, &wg)

	time.Sleep(1 * time.Second)
	command <- "Pause"

	time.Sleep(1 * time.Second)
	command <- "Play"

	time.Sleep(1 * time.Second)
	command <- "Stop"

	wg.Wait()
	/*
		1
		2
		3
		4
		Pause
		Play
		5
		6
		7
		8
		Stop
	*/
}

var i int

func work() {
	time.Sleep(250 * time.Millisecond)
	i++
	fmt.Println(i)
}

func routine(command <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	var status = "Play"
	for {
		select {
		case cmd := <-command:
			fmt.Println("x:", cmd)
			switch cmd {
			case "Stop":
				return
			case "Pause":
				status = "Pause"
			default:
				status = "Play"
			}
		default:
			if status == "Play" {
				work()
			}
		}
	}
}

func goUnordered() {
	go requestHttpUnordered("https://www.huoxiaoqiang.com/basic/programming/4333.html")                          //The differeence between statement and expression
	go requestHttpUnordered("https://segmentfault.com/a/1190000018689134")                                       //Anonymous functions and closures
	go requestHttpUnordered("https://medium.com/@kdnotes/how-to-sort-golang-maps-by-value-and-key-eedc1199d944") //Best Blog
	time.Sleep(10 * time.Second)
	/*
		requestHttpUnordered.http.Get:  https://medium.com/@kdnotes/how-to-sort-golang-maps-by-value-and-key-eedc1199d944
		requestHttpUnordered.http.Get:  https://www.huoxiaoqiang.com/basic/programming/4333.html
		requestHttpUnordered.http.Get:  https://segmentfault.com/a/1190000018689134
		requestHttpUnordered.http.Close:  https://medium.com/@kdnotes/how-to-sort-golang-maps-by-value-and-key-eedc1199d944
		requestHttpUnordered.http.Body:  https://medium.com/@kdnotes/how-to-sort-golang-maps-by-value-and-key-eedc1199d944
		requestHttpUnordered.http.Len:  13308
		requestHttpUnordered.http.Close:  https://segmentfault.com/a/1190000018689134
		requestHttpUnordered.http.Body:  https://segmentfault.com/a/1190000018689134
		requestHttpUnordered.http.Len:  42065
		requestHttpUnordered.http.Close:  https://www.huoxiaoqiang.com/basic/programming/4333.html
		requestHttpUnordered.http.Body:  https://www.huoxiaoqiang.com/basic/programming/4333.html
		requestHttpUnordered.http.Len:  49511

	*/
}

func requestHttpUnordered(url string) {
	fmt.Println("requestHttpUnordered.http.Get: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("requestHttpUnordered.http.Close: ", url)
	defer response.Body.Close()

	fmt.Println("requestHttpUnordered.http.Body: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("requestHttpUnordered.http.Len: ", len(body))
}

func goOrdered() {
	len := make(chan int)
	go requestHttpOrdered("https://www.golangprograms.com", len)
	println(<-len)
	go requestHttpOrdered("https://www.w3schools.com/go", len)
	println(<-len)
	go requestHttpOrdered("https://coderwall.com", len)
	time.Sleep(10 * time.Second)
	println(<-len)

	close(len) //Closes the channel
	/*
		requestHttpOrdered.http.Get:  https://www.golangprograms.com
		requestHttpOrdered.http.Close:  https://www.golangprograms.com
		requestHttpOrdered.http.Body:  https://www.golangprograms.com
		requestHttpOrdered.http.Len:  31852
		31852
		requestHttpOrdered.http.Get:  https://www.w3schools.com/go
		requestHttpOrdered.http.Close:  https://www.w3schools.com/go
		requestHttpOrdered.http.Body:  https://www.w3schools.com/go
		requestHttpOrdered.http.Len:  70092
		70092
		requestHttpOrdered.http.Get:  https://coderwall.com
		requestHttpOrdered.http.Close:  https://coderwall.com
		requestHttpOrdered.http.Body:  https://coderwall.com
		requestHttpOrdered.http.Len:  189752
		189752
	*/
}

func requestHttpOrdered(url string, l chan int) {
	fmt.Println("requestHttpOrdered.http.Get: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("requestHttpOrdered.http.Close: ", url)
	defer response.Body.Close()

	fmt.Println("requestHttpOrdered.http.Body: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	len := len(body)
	fmt.Println("requestHttpOrdered.http.Len: ", len)
	l <- len //return values
}

// WaitGroup is used to wait for the program to finish goroutines.
var wg sync.WaitGroup

func goWaitGroupUnordered() {
	//Add a count of three, one for each goroutine
	wg.Add(3)
	fmt.Println("Start goWaitGroup")

	go requestHttpWaitGroupUnordered("https://www.golangprograms.com")
	fmt.Println("x1")
	go requestHttpWaitGroupUnordered("https://stackoverflow.com")
	fmt.Println("x2")
	go requestHttpWaitGroupUnordered("https://coderwall.com")
	fmt.Println("x3")

	// Wait for the goroutines to finish.
	wg.Wait()
	fmt.Println("Ended goWaitGroup")
	/*
		Start goWaitGroup
		requestHttpWaitGroupUnordered.http.Get:  https://coderwall.com
		requestHttpWaitGroupUnordered.http.Get:  https://www.golangprograms.com
		requestHttpWaitGroupUnordered.http.Get:  https://stackoverflow.com
		requestHttpWaitGroupUnordered.http.Close:  https://coderwall.com
		requestHttpWaitGroupUnordered.http.Body:  https://coderwall.com
		requestHttpWaitGroupUnordered.http.Close:  https://stackoverflow.com
		requestHttpWaitGroupUnordered.http.Body:  https://stackoverflow.com
		requestHttpWaitGroupUnordered.http.Len:  189752
		requestHttpWaitGroupUnordered.http.Close:  https://www.golangprograms.com
		requestHttpWaitGroupUnordered.http.Body:  https://www.golangprograms.com
		requestHttpWaitGroupUnordered.http.Len:  31852
		requestHttpWaitGroupUnordered.http.Len:  175259
		Ended goWaitGroup
	*/
}

func requestHttpWaitGroupUnordered(url string) {
	//Schedule the call to WaitGroup's Done to tell goroutine is completed.
	defer wg.Done() //Similar to CountDownLatch in Java
	fmt.Println("requestHttpWaitGroupUnordered.http.Get: ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("requestHttpWaitGroupUnordered.http.Close: ", url)
	defer response.Body.Close()

	fmt.Println("requestHttpWaitGroupUnordered.http.Body: ", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("requestHttpWaitGroupUnordered.http.Len: ", len(body))
}
