package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

type Secret struct {
	APIKey string `json:"apikey"`
}

type MovieResponse struct {
	Title      string   `json:"Title"`
	Year       string   `json:"Year"`
	Rated      string   `json:"Rated"`
	Released   string   `json:"Released"`
	Runtime    string   `json:"Runtime"`
	Genre      string   `json:"Genre"`
	Director   string   `json:"Director"`
	Writer     string   `json:"Writer"`
	Actors     string   `json:"Actors"`
	Plot       string   `json:"Plot"`
	Language   string   `json:"Language"`
	Country    string   `json:"Country"`
	Awards     string   `json:"Awards"`
	Poster     string   `json:"Poster"`
	Ratings    []Rating `json:"Ratings"`
	Metascore  string   `json:"Metascore"`
	IMDBRating string   `json:"imdbRating"`
	IMDBVotes  string   `json:"imdbVotes"`
	IMDBID     string   `json:"imdbID"`
	Type       string   `json:"Type"`
	DVD        string   `json:"DVD"`
	BoxOffice  string   `json:"BoxOffice"`
	Production string   `json:"Production"`
	Website    string   `json:"Website"`
	Response   string   `json:"Response"`
	Error      string   `json:"Error"` // непусто, если Response=="False"
}

type Rating struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

const movieTmpl = `
Title: {{.Title}} ({{.Year}})
Rated: {{.Rated}}, Released: {{.Released}}, Runtime: {{.Runtime}}
Genre: {{.Genre}}
Director: {{.Director}}
Writer: {{.Writer}}
Actors: {{.Actors}}

Plot:
{{.Plot}}

Language: {{.Language}}, Country: {{.Country}}
Awards: {{.Awards}}

Ratings:
{{- range .Ratings }}
  • {{ printf "%-20s" (printf "%s:" .Source) }} {{ .Value }}
{{- end }}

Metascore: {{.Metascore}}, IMDB: {{.IMDBRating}} ({{.IMDBVotes}} votes), ID: {{.IMDBID}}
Type: {{.Type}}, DVD: {{.DVD}}
BoxOffice: {{.BoxOffice}}, Production: {{.Production}}
Website: {{.Website}}
`

var omdbapiURL = "http://www.omdbapi.com/"

func main() {

	var movieName string

	if len(os.Args) == 1 {
		fmt.Println("Ошибка: небходимо передать название фильма в командной строке")
		movieName = "stay"
	} else {
		movieName = os.Args[1]
	}

	data, err := os.ReadFile("conf/apikey.json")
	if err != nil {
		log.Fatal(err)
	}
	var secret Secret
	if err := json.Unmarshal(data, &secret); err != nil {
		log.Fatal(err)
	}

	// проверяем десериализацию
	fmt.Println(secret)
	// -----------------------

	poster := GetPoster(movieName, secret.APIKey)

	if poster == nil {
		fmt.Println("Не удалось получить информацию о фильме")
		return
	}

	t := template.Must(template.New("movie").Parse(movieTmpl))
	t.Execute(os.Stdout, *poster)

	fmt.Print("Нажмите Enter, чтобы выйти...")
	fmt.Scanln() // ждёт ввода до первого пробела или перевода строки
}

func GetPoster(name, apikey string) *MovieResponse {
	url := omdbapiURL + "?t=" + name + "&apikey=" + apikey

	fmt.Println(url)

	resp, err := http.Get(url)

	if err != nil {
		resp.Body.Close()
		log.Fatal(err)
		return nil
	}

	var movie MovieResponse

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil
	}

	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		resp.Body.Close()
		log.Fatal(err)
		return nil
	}

	resp.Body.Close()

	return &movie
}
