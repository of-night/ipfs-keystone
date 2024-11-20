package main

import (
	"fmt"
	"github.com/of-night/ipfs-keystone-test"
)

func main() {
	fmt.Println("Calling C function ipfs_keystone from Go...")
	ipfsKeystoneTest.Ipfs_keystone_test(3);
}
