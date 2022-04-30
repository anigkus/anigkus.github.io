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

package regex

import (
	"fmt"
	"regexp"
)

func Main() {
	fmt.Println("regex")

	matchingHHMM()

	matchingHTMLTag()

	//matchEmoji()

	matchHTMLImage()
}

func matchingHHMM() {
	fmt.Println("matchingHHMM")
	string1 := "8:2"
	string2 := "9:9"
	string3 := "12:29"
	string4 := "02:5"
	string5 := "23:59"
	string6 := "55:59"
	string7 := "0:01"
	regex := regexp.MustCompile(`^([0-9]|0[0-9]|1[0-9]|2[0-3]):([0-9]|[0-5][0-9])$`)

	fmt.Printf("Pattern: %v\n", regex.String()) // print pattern
	fmt.Printf("Time: %v\t:%v\n", string1, regex.MatchString(string1))
	fmt.Printf("Time: %v\t:%v\n", string2, regex.MatchString(string2))
	fmt.Printf("Time: %v\t:%v\n", string3, regex.MatchString(string3))
	fmt.Printf("Time: %v\t:%v\n", string4, regex.MatchString(string4))
	fmt.Printf("Time: %v\t:%v\n", string5, regex.MatchString(string5))
	fmt.Printf("Time: %v\t:%v\n", string6, regex.MatchString(string6))
	fmt.Printf("Time: %v\t:%v\n", string7, regex.MatchString(string7))
}

func matchingHTMLTag() {
	fmt.Println("matchingHTMLTag")
	string1 := `<html><body>
			<form name="query" method="post">
			<div id="div1" >matchingHTMLTag div innerText</div>
			<input type="submit" id="su" value="百度一下" class="bg s_btn">
			</form>
			</body></html>`

	regex := regexp.MustCompile(`<div.*?>(.*)</div>`)

	submatchall := regex.FindAllStringSubmatch(string1, -1)
	for _, element := range submatchall {
		fmt.Println(element[1])
	}
}

func matchEmoji() {
	fmt.Println("matchEmoji")
	var emojiRx = regexp.MustCompile(`[🍐]`)
	var str = emojiRx.ReplaceAllString("Thats a nice joke 🚗 🐝 🍎", `[e]`)
	fmt.Println(str)
}

func matchHTMLImage() {
	string1 := `
	<img src="https://cdn4.buysellads.net/uu/1/41334/1550855391-cc_dark.png" alt="ads via Carbon" border="0" height="100" width="130" style="max-width: 130px;">
	<img class="index-logo-srcnew" src="//www.baidu.com/img/flexible/logo/pc/result@2.png" alt="到百度首页" title="到百度首页">
	`
	re := regexp.MustCompile(`<img.*? src=["']([^"']+)["']`)
	submatchall := re.FindAllStringSubmatch(string1, -1)
	for _, element := range submatchall {
		fmt.Println(element[1])
	}
}
