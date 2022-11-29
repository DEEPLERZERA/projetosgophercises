package urlshort  

import ( //Importando pacotes
	"net/http"

	yaml "gopkg.in/yaml.v2"
)


func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc { //Função que recebe um map que recebe uma string e retorna uma string e um handler e retorna um handler
	return func(w http.ResponseWriter, r *http.Request) { //Retorna uma função que recebe um responseWriter e um request
		path := r.URL.Path //Armazena o path do request na variável path
		if dest, ok := pathsToUrls[path]; ok { //Verifica se o path existe no map
			http.Redirect(w, r, dest, http.StatusFound) //Redireciona para o path
			return //Retorna
		}
		fallback.ServeHTTP(w, r) //Se não existir, chama o fallback
	}
}

func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) { //Função que recebe um slice de bytes e um handler e retorna um handler e um erro
	  
	pathUrls, err := parseYaml(yamlBytes) //Armazena o slice de bytes na variável pathUrls e verifica se tem erro
	if err != nil {
		return nil, err
	}

	
	pathsToUrls := buildMap(pathUrls) //Armazena o map na variável pathsToUrls

	
	return MapHandler(pathsToUrls, fallback), nil //Retorna o mapHandler
}

func buildMap(pathUrls []pathUrl) map[string]string { //Função que recebe um slice de pathUrl e retorna um map que recebe uma string e retorna uma string
	pathsToUrls := make(map[string]string) //Cria um map que recebe uma string e retorna uma string
	for _, pu := range pathUrls { //Percorre o slice de pathUrl
		pathsToUrls[pu.Path] = pu.URL //Adiciona o path e a url no map
	}
	return pathsToUrls //Retorna o map
}

func parseYaml(data []byte) ([]pathUrl, error) { //Função que recebe um slice de bytes e retorna um slice de pathUrl e um erro caso tenha
	var pathUrls []pathUrl //Cria um slice de pathUrl
	err := yaml.Unmarshal(data, &pathUrls) //Deserializa o slice de bytes e armazena no slice de pathUrl e verifica se tem erro
	if err != nil {
		return nil, err
	}
	return pathUrls, nil //Retorna o slice de pathUrl e nil
}

type pathUrl struct { //Cria uma struct que recebe uma string e retorna uma string
	Path string `yaml:"path"` 
	URL  string `yaml:"url"`
}
