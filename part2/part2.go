package main

import "net"

func main() {
    listener, err := net.Listen("tcp", ":9000")
    if err != nil {
        panic(err)
    }

    var clients []net.Conn
    input := make(chan []byte, 10)
    go func() {
        for {
            message := <-input
            for _, client := range clients {
                client.Write(message)
            }
        }
    }()

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
