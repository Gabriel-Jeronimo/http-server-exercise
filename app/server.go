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

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	data := make([]byte, 1024)
	_, err = conn.Read(data)

	if err != nil {
		fmt.Println("Error reading data from connection: ", err.Error())
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("Error writing the response into the connection: ", err.Error())
	}

	httpData := strings.Split(string(data), CLFR)

	firstLine := strings.Split(httpData[0], " ")

	_, path, _ := firstLine[0], firstLine[1], firstLine[2]

	if path == "/" {
		conn.Write([]byte(HTTP_VERSION + " " + OK + DOUBLE_CLFR))
	} else if strings.Contains(path, "/echo/") {
		pathRandomString := strings.Split(path, "/echo/")[1]
		pathRandomStringLen := len(pathRandomString)

		conn.Write([]byte(HTTP_VERSION + " " + OK + CLFR))
		conn.Write([]byte("Content-Type: text/plain" + CLFR))
		conn.Write([]byte("Content-Length: " + strconv.Itoa(pathRandomStringLen) + DOUBLE_CLFR))
		conn.Write([]byte(pathRandomString + DOUBLE_CLFR))
	} else {
		conn.Write([]byte(HTTP_VERSION + " " + NOT_FOUND + DOUBLE_CLFR))
	}

}
