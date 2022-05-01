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
	"log"
	"regexp"
	"strings"
)

func Main() {
	fmt.Println("regex")

	matchingHHMMMatchString()

	matchingHTMLTagFindAllStringSubmatch()

	matchEmojiReplaceAllString()

	matchHTMLImageFindAllStringSubmatch()

	matchreplaceAllString()

	matchFindAllStringParentheses()

	matchFindAllStringSplit()

	mactchFindStringSubmatchExtractFilename()

	mactchFindStringExtractnumbers()
}

/*out:
matchingHHMM
Pattern: ^([0-9]|0[0-9]|1[0-9]|2[0-3]):([0-9]|[0-5][0-9])$
Time: 8:2	:true
Time: 9:9	:true
Time: 12:29	:true
Time: 02:5	:true
Time: 23:59	:true
Time: 55:59	:false
Time: 0:01	:true
*/
func matchingHHMMMatchString() {
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

/*out:
matchingHTMLTag
matchingHTMLTag div innerText
*/
func matchingHTMLTagFindAllStringSubmatch() {
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

/*out:
matchEmoji
Thats a nice joke [x] 😏
*/
func matchEmojiReplaceAllString() {
	fmt.Println("matchEmoji")
	/*
		https://www.runoob.com/charsets/emoji-smiley.html
		1F600~1F609:😀😁😂😃😄😅😆😇😈😉
		1F60A~1F60F:😊😋😌😍😎😏
	*/
	var emojiRx = regexp.MustCompile(`[\x{1F600}-\x{1F608}|[\x{1F60A}-\x{1F60E}]`)
	var str = emojiRx.ReplaceAllString("Thats a nice joke 😀 😏 ", `[x]`)
	fmt.Println(str)
}

/*out:
https://cdn4.buysellads.net/uu/1/41334/1550855391-cc_dark.png
//www.baidu.com/img/flexible/logo/pc/result@2.png
*/
func matchHTMLImageFindAllStringSubmatch() {
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

//https://www.runoob.com/regexp/regexp-syntax.html
/*out:
aAZz0123456789!@#$%^&*()-+=0 : 28 : 0
aAZz0123456789             0 : 28 : 13
*/
func matchreplaceAllString() {

	string1 := "aAZz0123456789!@#$%^&*()-+=0"
	fmt.Println(string1, ":", len(string1), ":", strings.Count(string1, " "))
	regex, err := regexp.Compile(`[^\w]`) //Match non [A-Za-z0-9_]
	if err != nil {
		log.Fatal(err)
	}
	string1 = regex.ReplaceAllString(string1, " ")
	fmt.Println(string1, ":", len(string1), ":", strings.Count(string1, " "))
}

/*out:
Pattern: \((.*?)\)

Text between parentheses:
anigkus
github
microsoft
*/
func matchFindAllStringParentheses() {
	string1 := "This is a (anigkus).((github)).(microsoft).com "

	regex := regexp.MustCompile(`\((.*?)\)`)
	fmt.Printf("Pattern: %v\n", regex.String()) // print pattern

	fmt.Println("\nText between parentheses:")
	submatchall := regex.FindAllString(string1, -1)
	for _, element := range submatchall {
		element = strings.Trim(element, "(")
		element = strings.Trim(element, ")")
		fmt.Println(element)
	}
}

/*out:
Pattern: [A-Z][^A-Z]*
Learning to understand the syntax of the
Golang development language by
Anigkus
*/
func matchFindAllStringSplit() {
	string1 := "Learning to understand the syntax of the Golang development language by Anigkus"

	regex := regexp.MustCompile(`[A-Z][^A-Z]*`)

	fmt.Printf("Pattern: %v\n", regex.String()) // Print Pattern

	submatchall := regex.FindAllString(string1, -1)
	for _, element := range submatchall {
		fmt.Println(element)
	}
}

/*out:
PCtm_d9c8750bed0b3c7d089fa7d55720d6cf
regexp-syntax
*/
func mactchFindStringSubmatchExtractFilename() {
	regex := regexp.MustCompile(`^(.*/)?(?:$|(.+?)(?:(\.[^.]*$)|$))`)

	string1 := `https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png`
	match1 := regex.FindStringSubmatch(string1)
	fmt.Println(match1[2])

	string2 := `https://www.runoob.com/regexp/regexp-syntax.html`
	match2 := regex.FindStringSubmatch(string2)
	fmt.Println(match2[2])
}

/*out:
Pattern: [-]?\d[\d,]*[\.]?[\d{2}]*
String contains any match: true
12
-23
100.00
0.001
100.00001
*/
func mactchFindStringExtractnumbers() {
	string1 := "A 12 b -23 C 100.00 Z 0.001 X:100.00001"

	regex := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	fmt.Printf("Pattern: %v\n", regex.String()) // Print Pattern

	fmt.Printf("String contains any match: %v\n", regex.MatchString(string1)) // True

	submatchall := regex.FindAllString(string1, -1)
	for _, element := range submatchall {
		fmt.Println(element)
	}
}
