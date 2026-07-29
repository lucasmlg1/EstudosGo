package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("Olá, mundo!", canal)
	for {
		//ERRO DE DEADLOCK, O CANAL AINDA ESTÁ ESPERANDO RECEBER DADOS, MAS NÃO TEM NENHUM LUGAR FORNECENDO ESSE DADO, RETORNA ERRO DE DEADLOCK

		mensagem, aberto := <-canal // enviando um valor
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}
	fmt.Println("Fim de programa")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto // recebendo um valor
		time.Sleep(time.Second)
	}
	//solucao do deadlock
	close(canal)
}
