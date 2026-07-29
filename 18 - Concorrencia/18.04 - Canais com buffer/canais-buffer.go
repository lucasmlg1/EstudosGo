package main

import "fmt"

func main() {
	//canal com buffer (2) capacidade
	canal := make(chan string, 2)
	// canal com buffer ele só vai bloquear esse trafego de dados quando o canal atingir a capacidade maxima
	//Já o sem buffer, sempre vai bloquear quando tiver esse trafego de dados(Envio e recebimento), são ações bloqueantes

	canal <- "Olá, mundo"
	mensagem := <-canal
	close(canal)
	fmt.Println(mensagem)
}
