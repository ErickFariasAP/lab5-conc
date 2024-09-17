package main

import (
    "encoding/binary"
    "net"
    "testing"
    "time"
)

func TestRegisterAndSearch(t *testing.T) {
    go main() // Inicia o servidor

    time.Sleep(time.Second) // Aguarda o servidor iniciar

    // Testa a conexão com o servidor de registro
    conn, err := net.Dial("tcp", "localhost:2000")
    if err != nil {
        t.Fatal("Failed to connect to register server:", err)
    }
    defer conn.Close()

    binary.Write(conn, binary.BigEndian, int64(12345)) // Envia o hash

    // Testa a conexão com o servidor de busca
    conn, err = net.Dial("tcp", "localhost:2001")
    if err != nil {
        t.Fatal("Failed to connect to search server:", err)
    }
    defer conn.Close()

    binary.Write(conn, binary.BigEndian, int64(12345)) // Envia o hash para busca

    var ipLen int64
    err = binary.Read(conn, binary.BigEndian, &ipLen)
    if err != nil {
        t.Fatal("Failed to read IP length:", err)
    }

    ipBytes := make([]byte, ipLen)
    _, err = conn.Read(ipBytes)
    if err != nil {
        t.Fatal("Failed to read IP string:", err)
    }

    if len(ipBytes) == 0 {
        t.Fatal("Expected IP string, got empty result")
    }
}
