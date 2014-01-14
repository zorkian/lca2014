package main

import "net"
import "bytes"

type Frotzer interface {
    Frotz([]byte) []byte
}

type Uppercaser struct{}
func (self Uppercaser) Frotz(input []byte) []byte {
    return bytes.ToUpper(input)
}

type Lowercaser struct{}
func (self Lowercaser) Frotz(input []byte) []byte {
    return bytes.ToLower(input)
}

func chatManager(clients *[]net.Conn, input chan []byte, frotz Frotzer) {
    for {
        message := <-input
        for _, client := range *clients {
            client.Write(frotz.Frotz(message))
        }
    }
}

func main() {
    listener, err := net.Listen("tcp", ":9000")
    if err != nil {
        panic(err)
    }

    var clients []net.Conn
    input := make(chan []byte, 10)
    go chatManager(&clients, input, Lowercaser{})
    go chatManager(&clients, input, Uppercaser{})

    for {
        client, err := listener.Accept()
        if err != nil {
            continue
        }
        clients = append(clients, client)
        go handleClient(client, input)
    }
}

func handleClient(client net.Conn, input chan []byte) {
    for {
        buf := make([]byte, 4096)
        numbytes, err := client.Read(buf)
        if numbytes == 0 || err != nil {
            return
        }
        input <- buf
    }
}
