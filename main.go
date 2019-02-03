package main

import (
    "fmt"
    "io/ioutil"
	"net/http"
	"bufio"
	"os"
	"strings"
)

/*Intenté crear una estructura a la que convertir el resultado de la petición, pero no obtuve los resultados
esperados*/
/*type Results struct{
	id int `json:"id"`
	name string `json:"name"`
	description string `json:"description"`
	modified string `json:"modified"`
}

type Data struct{
	offset int `json:"offset"`
	limit int `json:"limit"`
	total int `json:"total"`
	count int `json:"count"`
	results *Results `json:"results"`
}

type SuperHero struct{
	code int `json:"code"`
	status string `json:"status"`
	data *Data `json:"data"`
	etag string `json:"etag"`
	copyright string `json:"copyright"`
	attributionText string `json:"attributionText"`
	attributionHTML string `json:"attributionHTML"`
}*/

func main() {
	opciones := 
	`
	Bienvenido al consumo de la API de Marvel
	[ 1 ] Buscar por nombre
	[ 2 ] Listar
	Selecciona una opción
	`
	fmt.Print(opciones)

	reader := bufio.NewReader(os.Stdin)

	entrada, _ := reader.ReadString('\n')
	eleccion := strings.TrimRight(entrada, "\r\n")
	

	switch eleccion{
		case "1":
			fmt.Println("Ingresa el nombre del superheroe: ")
			nombre := bufio.NewReader(os.Stdin)
			input, _ := nombre.ReadString('\n')
			ingreso := strings.TrimRight(input, "\r\n")
			getByName(ingreso)
		case "2":
			fmt.Println("Listado de Superheroes")
			getContent()
		default:
			fmt.Println("El parámetro ingresado no corresponde a una opción del menú")
	}
}

func getContent(){
	publicKey := "38b5f450e7f7bdb2be5a5d21dbcf36d7"
	ts := "9"
	hash:= "ca02fac9e223fcbf9501302657212b91"
	URL:= fmt.Sprintf("http://gateway.marvel.com/v1/public/characters?orderBy=name&ts=%s&apikey=%s&hash=%s", ts, publicKey, hash)

	response, err := http.Get(URL)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))

		/*
		La manera en la que intenté convertirlo fue así:
		var superhero SuperHero
		json.Unmarshal(data, &superhero)
		*/
	}

}

func getByName(nombre string){
	publicKey := "38b5f450e7f7bdb2be5a5d21dbcf36d7"
	ts := "9"
	hash:= "ca02fac9e223fcbf9501302657212b91"

	URL:= fmt.Sprintf("http://gateway.marvel.com/v1/public/characters?name=%s&ts=%s&apikey=%s&hash=%s",nombre, ts, publicKey, hash)

	response, err := http.Get(URL)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}