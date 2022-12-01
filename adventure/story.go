package adventure

import ( //Importando pacotes
	"encoding/json"
	"io"
)

func JsonStory(r io.Reader) (Story, error) { //Criando função que decodifica o json
	d := json.NewDecoder(r)                  //Criando um decoder
	var story Story                          //Criando uma variável do tipo Story
	if err := d.Decode(&story); err != nil { //Se houver erro na decodificação
		return nil, err //Retorna o erro
	}
	return story, nil //Retorna a história
}

type Story map[string]Chapter //Criando um tipo Story que é um map de string para Chapter

type Chapter struct { //Criando uma struct Chapter
	Title      string   `json:"title"`   //Atributo Title do tipo string
	Paragraphs []string `json:"story"`   //Atributo Paragraphs do tipo []string
	Options    []Option `json:"options"` //Atributo Options do tipo []Option
}

type Option struct { //Criando uma struct Option
	Text    string `json:"text"` //Atributo Text do tipo string
	Chapter string `json:"arc"`  //Atributo Chapter do tipo string
}
