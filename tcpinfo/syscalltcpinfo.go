package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"syscall"
	"time"
)

const DOWNLOAD_URL = "http://ipv4.download.thinkbroadband.com/1GB.zip"

func myControl(network, address string, c syscall.RawConn) error {
	return c.Control(func(fd uintptr) {
		go func(fd uintptr) {
			var previousState uint8
			for range time.Tick(10 * time.Millisecond) {
				tcpInfo, _ := unix.GetsockoptTCPInfo(int(fd), syscall.SOL_TCP, syscall.TCP_INFO)
				if tcpInfo.State != previousState {

					fmt.Printf("FD: %d State change %d -> %d\n", int(fd), previousState, tcpInfo.State)
					fmt.Printf("FD: %d %+v\n", int(fd), tcpInfo)
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

func main() {

	dialer := &net.Dialer{Control: myControl}
	transport := &http.Transport{Dial: dialer.Dial}
	client := &http.Client{Transport: transport}
	resp, clientErr := client.Get(DOWNLOAD_URL)

	if clientErr != nil {
		panic(clientErr)
	}

	_, _ = io.Copy(ioutil.Discard, resp.Body)
}
