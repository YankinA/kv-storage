package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

type Storage struct {
	data map[string]string
	mu   sync.RWMutex
}

func initStorage() *Storage {
	storage := &Storage{data: make(map[string]string)}
	return storage
}

func (storage *Storage) Set(key, val string) bool {
	storage.mu.Lock()
	defer storage.mu.Unlock()
	storage.data[key] = val
	return true
}

func (storage *Storage) Get(key string) (string, bool) {
	storage.mu.RLock()
	defer storage.mu.RUnlock()

	val, ok := storage.data[key]
	if ok {
		return val, ok
	}
	return "", ok
}

func (storage *Storage) Delete(key string) bool {
	storage.mu.Lock()
	defer storage.mu.Unlock()
	_, ok := storage.data[key]
	if ok {
		delete(storage.data, key)
	}
	return ok
}

func runServer(storage *Storage) {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return

	}
	defer listener.Close()

	fmt.Println("Server is listening on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection err")
			continue
		}
		handleConn(conn, storage)

	}

}

func handleConn(conn net.Conn, storage *Storage) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		commands := strings.Fields(text)

		if len(commands) == 0 {
			continue
		}
		command := strings.ToLower(commands[0])

		switch command {
		case "set":
			if len(commands) < 3 {
				conn.Write([]byte("Use command: set <key> <value>\n"))
				continue
			}

			if ok := storage.Set(commands[1], commands[2]); ok {
				conn.Write([]byte("Ok\n"))
			}
		case "get":
			if len(commands) < 2 {
				conn.Write([]byte("Use command: get <key>\n"))
				continue
			}

			if value, ok := storage.Get(commands[1]); ok {
				conn.Write([]byte(value + "\n"))
			} else {
				conn.Write([]byte("Key not found\n"))
			}

		case "del":
			if len(commands) < 2 {
				conn.Write([]byte("Use command: del <key>\n"))
				continue
			}

			if ok := storage.Delete(commands[1]); ok {
				conn.Write([]byte("Ok\n"))
			} else {
				conn.Write([]byte("Key not found\n"))
			}
		default:
			conn.Write([]byte("Unknown command\n"))
		}

	}

}

func main() {
	storage := initStorage()
	runServer(storage)
}
