package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	HTTP_VERSION = "HTTP/1.1"
	OK           = "200 OK"
	NOT_FOUND    = "404 NOT FOUND"
	CLFR         = "\r\n"
	DOUBLE_CLFR  = CLFR + CLFR
)

func HandleRequests(conn net.Conn) {
	data := make([]byte, 1024)
	_, err := conn.Read(data)

	if err != nil {
		fmt.Println("Error reading data from connection: ", err.Error())
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("Error writing the response into the connection: ", err.Error())
	}

	req := strings.Split(string(data), CLFR)

	startLine := strings.Split(req[0], " ")
	userAgent := req[2]

	_, path, _ := startLine[0], startLine[1], startLine[2]

	if path == "/" {
		conn.Write([]byte(HTTP_VERSION + " " + OK + DOUBLE_CLFR))
	} else if strings.Contains(path, "/echo/") {
		pathRandomString := strings.Split(path, "/echo/")[1]

		conn.Write([]byte(HTTP_VERSION + " " + OK + CLFR))
		conn.Write([]byte("Content-Type: text/plain" + CLFR))
		conn.Write([]byte("Content-Length: " + strconv.Itoa(len(pathRandomString)) + DOUBLE_CLFR))
		conn.Write([]byte(pathRandomString + DOUBLE_CLFR))
	} else if path == "/user-agent" {
		userAgentValue := strings.Split(userAgent, ": ")[1]

		conn.Write([]byte(HTTP_VERSION + " " + OK + CLFR))
		conn.Write([]byte("Content-Type: text/plain" + CLFR))
		conn.Write([]byte("Content-Length: " + strconv.Itoa(len(userAgentValue)) + DOUBLE_CLFR))
		conn.Write([]byte(userAgentValue + DOUBLE_CLFR))
	} else {
		conn.Write([]byte(HTTP_VERSION + " " + NOT_FOUND + DOUBLE_CLFR))
	}

	defer conn.Close()
}
func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()

		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go HandleRequests(conn)
	}
}
