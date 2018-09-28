// Copyright 2018 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"log"

	"github.com/FactomProject/factom"
)

var (
	sflag = flag.Bool("s", false, "encode a secret key")
	pflag = flag.Bool("p", true, "encode a public key")
)

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		usage()
	}

	hexaddr := flag.Args()[0]

	p, err := hex.DecodeString(hexaddr)
	if err != nil {
		log.Fatal(err)
	}

	if *sflag {
		address, err := factom.MakeECAddress(p)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(address)
		fmt.Println(address.SecString())
	} else {
		address := factom.NewECAddress()
		copy(address.Pub[:], p)
		fmt.Println(address)
	}
}

func usage() {
	log.Fatal("Usage: addresser HEXSTRING")
}
