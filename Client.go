package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func gestionErreur(err, error) {
	if err != {
		panic(err)
	}
}

const (
	IP = "127.0.0.1" //LocalHost
	PORT = "3569" //Port utilisé
)

func main() {
	//connexion au server
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", IP, PORT))
	gestionErreur(err)

	for {
		//entré utilisateur
		reader := buflo.NewReader(os.Stdin)
		fmt.Print(" Client > ")
		text, err := reader.ReadString('\n')
		gestionErreur(err)

		//envoi du message au server
		conn.Write([]byte(text))

		//écoute les messages emis par le server + retour a la ligne
		message, err := bufio.NewReader(conn).ReadString('\n')
		gestionErreur(err)

		//on affiche le message utilisateur
		fmt.Print(" Server > " + message)
	}
}
