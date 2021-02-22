package main

import(
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
)

//run a test to the api choosing for an albumId that does not exist and seeing if we print out a message to notify the user
func TestProcessAlbumId_EmptyOutput(t *testing.T){
	//create a mock http server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		//return empty array if albumId matches the test value of 100000
		if req.URL.String() == "/photos?albumId=100000"{
			fmt.Fprint(w,`[]`)
		}else{
			fmt.Fprint(w,`{"Some Error?"}`)
		}
	}))

	defer server.Close()

	api = API{server.Client(), server.URL}

	//record the input of the pipe
	r, w, _ := os.Pipe()
	os.Stdout = w

	processAlbumNumber(100000)

	w.Close()
	out, _ := ioutil.ReadAll(r)

	if string(out) !=  "Album not found or is empty\n"{

		t.Errorf("Failed to print 'Album not found or is empty\n' and instead printed out: %s", string(out))

	}


}

//call the api with a valid albumId and verifying the format of the output printed
func TestProcessAlbumNumber_GoodOutput(t *testing.T){
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
                if req.URL.String() == "/photos?albumId=1"{
                        fmt.Fprint(w,`[
  {
    "albumId": 1,
    "id": 1,
    "title": "accusamus beatae ad facilis cum similique qui sunt",
    "url": "https://via.placeholder.com/600/92c952",
    "thumbnailUrl": "https://via.placeholder.com/150/92c952"
  },
  {
    "albumId": 1,
    "id": 2,
    "title": "reprehenderit est deserunt velit ipsam",
    "url": "https://via.placeholder.com/600/771796",
    "thumbnailUrl": "https://via.placeholder.com/150/771796"
  }
]`)

                }else{
                        fmt.Fprint(w,`{"Some Error?"}`)
                }
        }))

        defer server.Close()

        api = API{server.Client(), server.URL}

        r, w, _ := os.Pipe()
        os.Stdout = w

        processAlbumNumber(1)

        w.Close()
        out, _ := ioutil.ReadAll(r)

        if string(out) !="\n[1] accusamus beatae ad facilis cum similique qui sunt\n\n[2] reprehenderit est deserunt velit ipsam\n"{
                t.Errorf("processAlbumNumber(1) printed an unexpected output: %s", string(out))
        }



}

func TestProcessAlbumId_ExitWith0(t *testing.T){
	r, w, _ := os.Pipe()
        os.Stdout = w

        processAlbumNumber(0)

        w.Close()
        out, _ := ioutil.ReadAll(r)

	if string(out) !="Exiting program.\n"{
                t.Errorf("processAlbumNumber(0) printed an unexpected output: %s", string(out))
        }

}

//testing that our function will verify inputs are only integers, and testing that 0 exits the program
func TestConvertAlbumIdInputToPositiveInt(t *testing.T){
	//testing valid input of "1"
	r, w, _:= os.Pipe()
	os.Stdout = w
	returnedInt,err := convertAlbumIdInputToPositiveInt("1")

	w.Close()
	out, _:= ioutil.ReadAll(r)

	if string(out) != ""{
		t.Errorf("convertAlbumIdInputToPositiveInt(\"1\") should not print anything--instead this printed: %s", string(out)) 
	}


	if returnedInt != 1{
		t.Errorf("convertAlbumIdInputToPositiveInt(\"1\") should have returned an int64 of value 1, instead it returned: %v", returnedInt)
	}

	if err!= nil{
		t.Errorf("convertAlbumIdInputToPositiveInt(\"1\") should not have returned an error")
	}

	//testing invalid input "-10" 
	r, w, _ = os.Pipe()
	os.Stdout = w
	returnedInt,err = convertAlbumIdInputToPositiveInt("-10")
	w.Close()
	out, _ = ioutil.ReadAll(r)

	if string(out) != "Please re-enter a valid positive integer for the photo-album\n"{
                t.Errorf("convertAlbumIdInputToPositiveInt(\"-10\") should print a message that valid positive integers are required--instead this printed: %s", string(out))
        }

        if returnedInt != -10{
                t.Errorf("convertAlbumIdInputToPositiveInt(\"-10\") should have returned an int64 of value -10, instead it returned: %v", returnedInt)
        }

        if err == nil{
                t.Errorf("convertAlbumIdInputToPositiveInt(\"-10\") should have returned an error")
        }

	//testing invalid input "asdf"
	r, w, _ = os.Pipe()
	os.Stdout =w

	returnedInt,err = convertAlbumIdInputToPositiveInt("asdf")
        w.Close()
        out, _ = ioutil.ReadAll(r)

        if string(out) != "Please re-enter a valid positive integer for the photo-album\n"{
                t.Errorf("convertAlbumIdInputToPositiveInt(\"asdf\") should print a message that valid positive integers are required--instead this printed: %s", string(out))
        }

        if returnedInt != 0{
                t.Errorf("convertAlbumIdInputToPositiveInt(\"asdf\") should have returned an 0 value, instead it returned: %v", returnedInt)
        }

        if err == nil{
                t.Errorf("convertAlbumIdInputToPositiveInt(\"asdf\") should have returned an error")
        }


}

