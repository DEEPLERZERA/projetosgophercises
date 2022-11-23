package main

import ( //Importando os pacotes
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() { //Função principal
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'") //Criando uma variável que recebe uma flag para o arquivo csv
	timeLimit := flag.Int("limite", 30, "O tempo limite para o quiz em segundos") //Criando uma variável que recebe uma flag para o tempo limite
	flag.Parse()                                                                                       //Lendo a flag

	file, err := os.Open(*csvFilename) //Abrindo o arquivo csv
	if err != nil {                    //Caso ocorra algum erro
		exit(fmt.Sprintf("Falha ao abrir arquivo csv: %s\n", *csvFilename)) //Imprime a mensagem de erro

	}

	r := csv.NewReader(file)  //Lendo o arquivo csv e armazenando na variável r do tipo csv.Reader
	lines, err := r.ReadAll() //Lendo todas as linhas do arquivo csv e armazenando na variável lines
	if err != nil {           //Caso ocorra algum erro
		exit("Falha ao tentar analisar o arquivo csv providenciado.") //Imprime a mensagem de erro
	}
	problems := parseLines(lines) //Armazenando o slice de problem na variável problems

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second) //Criando um timer com o tempo limite

	correct := 0  //Variável que armazena a quantidade de respostas corretas
	for i, p := range problems { //Percorrendo o slice de problem 
		fmt.Printf("Problema #%d: %s = \n", i+1, p.q) //Imprimindo a pergunta
		answerCh := make(chan string) //Criando um canal que recebe uma string
		go func() { //Criando uma goroutine
			var answer string 							//Criando uma variável que recebe a resposta
			fmt.Scanf("%s\n", &answer)                 //Lendo a resposta
			answerCh <- answer 							//Enviando a resposta para o canal
		}()

		select {
		case <- timer.C:
			fmt.Printf("Você pontuou %d de %d questões!.\n", correct, len(problems)) //Imprime a pontuação
			return 
		case answer := <- answerCh:
			if answer == p.a {                         //Caso a resposta seja igual a resposta correta
				fmt.Println("Você acertou!") //Imprime a mensagem de acerto
				correct++ 				 //Incrementa a variável correct
			} else {
				fmt.Printf("Você errou! A respota correta era %s \n", p.a) //Imprime a mensagem de erro
				}

		}
		
	}
	

	fmt.Printf("Você pontuou %d de %d questões!.\n", correct, len(problems)) //Imprime a pontuação
}

func parseLines(lines [][]string) []problem { //Função que recebe um slice de slice de string e retorna um slice de problem
	ret := make([]problem, len(lines)) //Criando um slice de problem com o tamanho do slice de slice de string
	for i, line := range lines { 	 //Percorrendo o slice de slice de string
		ret[i] = problem{ //Armazenando na posição i do slice de problem
			q: line[0], //A primeira posição do slice de string
			a: strings.TrimSpace(line[1]), //A segunda posição do slice de string
		}
	}
	return ret //Retornando o slice de problem
}

type problem struct { //Criando uma struct problem
	q string //Pergunta
	a string //Resposta
}

func exit(msg string) { //Função que recebe uma string e imprime a mensagem de erro
	fmt.Println(msg) //Imprime a mensagem de erro
	os.Exit(1) 	 //Sai do programa
}
