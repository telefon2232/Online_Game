package main
import (
	"fmt"
	"net"
)


func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_,err := conn.WriteToUDP([]byte("From server: Hello I got your message "), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}


func main() {
	p := make([]byte, 148)
	addr := net.UDPAddr{
		Port: 1234,
		IP: net.ParseIP("127.0.0.1"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	for {
		_,remoteaddr,err := ser.ReadFromUDP(p)
		for i:=0;i<4;i++{
			fmt.Println(p[i])
		}
		fmt.Printf("Read a message from %v %s \n", remoteaddr, len(p))
		if err !=  nil {
			fmt.Printf("Some error  %v", err)
			continue
		}
		//go sendResponse(ser, remoteaddr)
	}
}