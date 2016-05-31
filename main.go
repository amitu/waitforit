package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

func copyAndClose(w io.WriteCloser, r io.Reader) {
	io.Copy(w, r)
	if err := w.Close(); err != nil {
		log.Println("Error closing", err)
	}
}

func main() {
	var listen = flag.String("listen", "127.0.0.1:8001", "Listen here.")
	var remote = flag.String("remote", "127.0.0.1:8000", "Forward there.")

	flag.Parse()

	ln, err := net.Listen("tcp", *listen)
	if err != nil {
		log.Panic(err.Error())
	}

	log.Println("Listening on", *listen)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		log.Println("Got connection, proxying.")

		go func(conn net.Conn) {
			var rconn net.Conn
			waited := false
			start := time.Now()
			for {
				rconn, err = net.Dial("tcp", *remote)
				if err != nil {
					// log.Println(err.Error())
					waited = true
					time.Sleep(time.Millisecond * 100)
					// TODO: Check if connection is still open?
					continue
				}
				break
			}

			if waited {
				log.Println("Waited for", time.Since(start))
			}

			go copyAndClose(rconn, conn)
			go copyAndClose(conn, rconn)
		}(conn)
	}
}
