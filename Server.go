package main

import (
	"fmt"
	"net"
)

func gestionErreur(err error){
	if err != nil {
		panic(err)
	}
}

const (
	IP = "127.0.0.1" //LocalHost
	PORT = "3569" //Port utilisé
)

func main() {
	fmt.Println("Lancement du server ...")

	//écoute sur le port 3569
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%s", IP, PORT))
	gestionErreur(err)

	//On accepte les connexions sur le port 3569
	conn, err := ln.Accept()
	if err != nil {
		panic(err)
	}

	//Information sur le ou les client(s) qui se connectent
	fmt.Println(" [+] Un client s'est connecté depuis", conn.RemoteAddr())

	gestionErreur(err)

	for {
		//écoute des messages émis par le clients
		buffer := make([]byte, 4096) // taille maximal du message envoyé par un client
		length, err := conn.Read(buffer)
		message := string(buffer[:length])

		if err != nil {
			fmt.Println(" [!] Le client s'est déconnecté")
		}

		fmt.Println(" Client > ", message)

		conn.Write([]byte(message + "\n"))
	}
}
