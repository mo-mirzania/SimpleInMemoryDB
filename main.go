package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		log.Panic()
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	dic := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		input := strings.ToLower(scanner.Text())
		verbs := strings.Fields(input)
		switch verbs[0] {
		case "set":
			if len(verbs) != 3 {
				fmt.Fprintf(conn, "Provide a value\n")
			} else {
				dic[verbs[1]] = verbs[2]
			}
		case "get":
			if len(verbs) != 2 {
				fmt.Fprintf(conn, "Provide value\n")
			} else {
				fmt.Fprintf(conn, "%v\n", dic[verbs[1]])
			}
		case "del":
			if len(verbs) != 2 {
				fmt.Fprintf(conn, "Provide value\n")
			} else {
				delete(dic, verbs[1])
				fmt.Fprintf(conn, "Deleted!\n")
			}
		default:
			fmt.Fprintf(conn, "Verb not found\n")
		}
	}
}
