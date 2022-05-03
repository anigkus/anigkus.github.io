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

package cryptography

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Main() {
	fmt.Println("Cryptography...")

	cryptoHash()
}

func cryptoHash() {

	fmt.Println("\n----------------Small Message----------------")
	message := []byte(RandStringRunes(10))

	fmt.Printf("Md5: %x\n\n", md5.Sum(message))
	fmt.Printf("Sha1: %x\n\n", sha1.Sum(message))
	fmt.Printf("Sha256: %x\n\n", sha256.Sum256(message))
	fmt.Printf("Sha512: %x\n\n", sha512.Sum512(message))

	fmt.Println("\n----------------Large Message----------------")
	message = []byte(RandStringRunes(100))

	fmt.Printf("Md5: %x\n\n", md5.Sum(message))
	fmt.Printf("Sha1: %x\n\n", sha1.Sum(message))
	fmt.Printf("Sha256: %x\n\n", sha256.Sum256(message))
	fmt.Printf("Sha512: %x\n\n", sha512.Sum512(message))
}
