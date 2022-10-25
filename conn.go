package redishoneypot

import (
	"io"
	"log"
	"net"
)

func handleConn(conn *net.TCPConn, repo Repository) {
	defer handleConnClose(conn)

	for {
		cmd, err := getCmd(conn)
		if err != nil {
			log.Println(err)
			break
		}

		log.Printf("received command \"%s\" from %s", cmd.ToString(), cmd.IP)
		if err := repo.Save(cmd); err != nil {
			log.Println(err)
			break
		}

		res := makeResStr(cmd)
		if _, err := io.WriteString(conn, res); err != nil {
			log.Println(err)
			break
		}
	}
}

func handleConnClose(conn *net.TCPConn) {
	conn.Close()
	log.Printf("connection from %s closed\n", conn.RemoteAddr().String())
}
