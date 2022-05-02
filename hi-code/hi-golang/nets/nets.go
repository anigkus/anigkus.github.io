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

package nets

import (
	"fmt"
	"log"
	"net"
)

func Main() {
	lookUp()
}

/*out:

 */
func lookUp() {
	//
	names, err := net.LookupAddr(`8.8.8.8`)
	if err != nil {
		log.Fatal(err)
	}
	for _, name := range names {
		fmt.Println(name)
		//dns.google.
	}

	//
	cname, err := net.LookupCNAME("m.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cname) //wap.wshifen.com.

	//
	addrs, err := net.LookupHost(`www.google.com`)
	if err != nil {
		log.Fatal(err)
	}
	for _, addr := range addrs {
		fmt.Println(addr)
		// 142.250.186.36
		// 2a00:1450:4001:813::2004
	}
	// ips, err := net.LookupIP("google.com")

	// mxs, err := net.LookupMX("google.com")

	// net.LookupNS("google.com")
	//net.LookupPort("")

	//
	// cname, srvs, err := net.LookupSRV("xmpp-server", "tcp", "golang.org")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("\ncname: %s \n\n", cname)
	// for _, srv := range srvs {
	// 	fmt.Printf("%v:%v:%d:%d\n", srv.Target, srv.Port, srv.Priority, srv.Weight)
	// }

	//
	names, err = net.LookupTXT("google.com")
	if err != nil {
		log.Fatal(err)
	}
	for _, name := range names {
		fmt.Println(name)
		//facebook-domain-verification=22rm551cu4k0ab0bxsw536tlds4h95
		// v=spf1 include:_spf.google.com ~all
		// docusign=05958488-4752-4ef2-95eb-aa7ba8a3bd0e
		// MS=E4A68B9AB2BB9670BCE15412F62916164C0B20BB
		// google-site-verification=TV9-DBe4R80X4v0M4U_bd_J9cpOJM0nikft0jAgjmsQ
		// globalsign-smime-dv=CDYX+XFHUw2wml6/Gb8+59BsH31KzUr6c1l2BPvqKX8=
		// apple-domain-verification=30afIBcvSuDV2PLX
		// docusign=1b0a6754-49b1-4db5-8540-d2c12664b289
		// google-site-verification=wD8N7i1JTNTkezJ49swvWW48f8_9xveREV4oB-0Hf5o
	}

}
