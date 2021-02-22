package main

import(
	"fmt"
	"strconv"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"
)

type Photo struct{
	Id	int `json:"id"`
	Title	string `json:"title"`
}

type API struct {
	Client  *http.Client
	baseURL string
}

var api API

func init(){
	api = API{ &http.Client{}, "https://jsonplaceholder.typicode.com"}
}

func (api *API) GetData(param string) ([]byte, error){

	response, err := api.Client.Get(api.baseURL+ param);

        if err != nil {
                fmt.Errorf ("Error found with Get request: %v", err)
        }

	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func processAlbumNumber(albumId int) (){

	var photos []Photo

	if (albumId == 0) {
		fmt.Println("Exiting program.")
		return
	}

	responseData, err := api.GetData("/photos?albumId=" + strconv.Itoa(albumId))


	if err != nil {
		fmt.Errorf("Error found with converting response body: %v", err)
	}

	if string(responseData) == "[]" {
		fmt.Println("Album not found or is empty")
		return
	}


	json.Unmarshal(responseData, &photos)

	for _, val := range photos {
		fmt.Printf("\n[%v] %s\n", val.Id, val.Title)
	}

}

func convertAlbumIdInputToPositiveInt(input string) (int64,error){
	albumId, err := strconv.ParseInt(input, 10, 0)

        if err != nil || albumId < 0{
		fmt.Println("Please re-enter a valid positive integer for the photo-album")
		return albumId, errors.New("Invalid Integer")
        }
	return albumId, nil
}

func main() {

	var input string

	fmt.Println("Please type a number for the album which contains the photo ids and titles that you wish to display")
	for {
		fmt.Print("> photo-album ")
		fmt.Scan(&input)

		albumId, err := convertAlbumIdInputToPositiveInt(input)
		if err == nil{
			processAlbumNumber(int(albumId))

			if(albumId == 0){
				break
			}
		}

	}
}
