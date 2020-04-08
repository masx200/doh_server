/*
   DNS-over-HTTPS
   Copyright (C) 2017-2018 Star Brilliant <m13253@hotmail.com>

   Permission is hereby granted, free of charge, to any person obtaining a
   copy of this software and associated documentation files (the "Software"),
   to deal in the Software without restriction, including without limitation
   the rights to use, copy, modify, merge, publish, distribute, sublicense,
   and/or sell copies of the Software, and to permit persons to whom the
   Software is furnished to do so, subject to the following conditions:

   The above copyright notice and this permission notice shall be included in
   all copies or substantial portions of the Software.

   THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
   IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
   FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
   AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
   LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
   FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
   DEALINGS IN THE SOFTWARE.
*/

package main

import (
	"flag"
	"fmt"
	"log"
)

var dns_stat map[string]int

func main() {
	confPath := flag.String("c", "/svc/doh-server.conf", "Configuration file")
	verbose := flag.Bool("v", false, "Enable logging")
	flag.Parse()

	conf, err := loadConfig(*confPath)
	if err != nil {
		log.Fatalln(err)
	}

	if *verbose {
		conf.Verbose = true
	}
	dns_stat = make(map[string]int)
	for _, k := range conf.Upstream {
		dns_stat[k] = 0
	}

	name := `
      _
     / \
   _------_
  /        \
 |          |
 |          |
 |     __  __)     DNS
 |    /  \/  \
/\/\ (o   )o  )    over
/c    \__/ --.
\_   _-------'     HTTP2S:
 |  /         \
 | | '\_______)    DOH!
 |  \_____)
 |_____ |
|_____/\/\
/        \   `
	fmt.Println(name)

	server := NewServer(conf)
	_ = server.Start()
}
