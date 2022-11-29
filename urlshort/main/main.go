package main

import ( //Importando pacotes
	"fmt"
	"net/http"

	"urlshort"
)

func main() { //Função principal
	mux := defaultMux() //Armazenando o mux padrão na variável mux

	pathsToUrls := map[string]string{ //Criando um map que recebe uma string e retorna uma string
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux) //Armazenando o mapHandler na variável mapHandler (: aaa

	yaml := ` #Criando um slice de yaml
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler) //Armazenando o yamlHandler na variável yamlHandler e verificando se tem erro
	if err != nil {
		panic(err)
	}
	fmt.Println("Iniciando o servidor na porta :8080") //Imprime a mensagem
	http.ListenAndServe(":8080", yamlHandler)          //Inicia o servidor na porta 8080
}

func defaultMux() *http.ServeMux { //Função que retorna um mux padrão
	mux := http.NewServeMux()  //Cria um mux
	mux.HandleFunc("/", hello) //Adiciona a função hello no mux
	return mux                 //Retorna o mux
}

func hello(w http.ResponseWriter, r *http.Request) { //Função que imprime a mensagem
	fmt.Fprintln(w, "Olá, mundo!") //Imprime a mensagem
}
