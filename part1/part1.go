package main

import "net"

func main() {
    listener, err := net.Listen("tcp", ":9000")
    if err != nil {
        panic(err)
    }

    for {
        client, err := listener.Accept()
        if err != nil {
            continue
        }
        handleClient(client)
    }
}

func handleClient(client net.Conn) {
    for {
        buf := make([]byte, 4096)
        numbytes, err := client.Read(buf)
        if numbytes == 0 || err != nil {
            return
        }
        client.Write(buf)
    }
}
