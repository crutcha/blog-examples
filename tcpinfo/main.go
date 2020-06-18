package main

import (
	// "bufio"
	"encoding/json"
	"fmt"
	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
	"io"
	"math/rand"
	"net"
	"net/http"
	"os"
	"sync"
	"syscall"
	"time"
)

const (
	CONCURRENT_DOWNLOADS = 5
	// DOWNLOAD_URL         = "http://ipv4.download.thinkbroadband.com/1GB.zip"
	DOWNLOAD_URL = "http://ipv4.download.thinkbroadband.com/100MB.zip"
)

func myControl(network, address string, c syscall.RawConn) error {
	return c.Control(func(fd uintptr) {
		go func(fd uintptr) {
			var previousState uint8
			for range time.Tick(10 * time.Millisecond) {
				tcpInfo, _ := unix.GetsockoptTCPInfo(int(fd), syscall.SOL_TCP, syscall.TCP_INFO)
				// fmt.Printf("FD: %d %+v\n", int(fd), tcpInfo)
				if tcpInfo.State != previousState {

					fmt.Printf("FD: %d State change %d -> %d\n", int(fd), previousState, tcpInfo.State)
					previousState = tcpInfo.State

					// initial value would be 0
					if previousState == 0 {
						fmt.Printf("FD: %d Done at: %s\n", int(fd), time.Now())
						return
					}
				}
			}
		}(fd)
	})
}

func oldmain() {
	var wg sync.WaitGroup

	for i := 1; i <= CONCURRENT_DOWNLOADS; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			dialer := &net.Dialer{Control: myControl}
			transport := &http.Transport{Dial: dialer.Dial}
			client := &http.Client{Transport: transport}
			resp, clientErr := client.Get(DOWNLOAD_URL)

			if clientErr != nil {
				panic(clientErr)
			}

			defer resp.Body.Close()
			out, outErr := os.Create(fmt.Sprintf("/tmp/tcpinfo-%d", rand.Int()))
			fmt.Println("outErr: ", outErr)
			defer out.Close()

			_, copyErr := io.Copy(out, resp.Body)
			fmt.Println("copyErr: ", copyErr)
			fmt.Println("Copy done at: ", time.Now())

		}()
	}
	wg.Wait()

	// for demo purposes only, since TCP socket cleanup may not be instant
	time.Sleep(30 * time.Second)
	res, _ := netlink.SocketDiagTCPInfo(1)
	fmt.Println(res)
}

func main() {
	res, diagErr := netlink.SocketDiagTCPInfo(unix.AF_INET)
	if diagErr != nil {
		panic(diagErr)
	}
	out, _ := json.Marshal(res)
	fmt.Println(string(out))
	/*
		for _, result := range res {
			if result.TCPInfo != nil {
			    out, _ := json.Marshal(*result)
				fmt.Println(string(out))
			}
		}
	*/
}
