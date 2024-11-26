package main

import (
	"fmt"
	"io"
	"github.com/of-night/ipfs-keystone-test"
)

func main() {
	fmt.Println("Calling C function ipfs_keystone from Go...")
	fmt.Println("test Processing file")
	reader := ipfsKeystoneTest.Ipfs_keystone_test(1, "aestest.txt")
	defer reader.Close()
	buf := make([]byte, 16)
	for {
		n, err := io.ReadFull(io.Reader(&reader), buf)
		fmt.Println("Read is ok")
		if err != nil {
			if err != io.ErrUnexpectedEOF {
				fmt.Println("tt1 ", err)
				err = io.EOF
				fmt.Println("ttt ", err)
				fmt.Println("Read %d bytes: ", n)
				for i := 0; i < n; i++ {
					fmt.Printf("%02X", buf[i])
					if (i+1)%16 == 0 || i == n-1 {
						fmt.Println()
					}
				}
				return
			}
			fmt.Println("t1t ", err)
			fmt.Println("Read %d bytes: ", n)
			for i := 0; i < n; i++ {
				fmt.Printf("%02X", buf[i])
				if (i+1)%16 == 0 || i == n-1 {
					fmt.Println()
				}
			}
			return
		}

		fmt.Println("Read %d bytes: ", n)
		for i := 0; i < n; i++ {
			fmt.Printf("%02X", buf[i])
			if (i+1)%16 == 0 || i == n-1 {
				fmt.Println()
			}
		}
	}
	return
}
