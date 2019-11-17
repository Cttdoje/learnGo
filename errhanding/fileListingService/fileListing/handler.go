package fileListing

import (
	"io/ioutil"
	"net/http"
	"os"
)

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	index := len("/list/")
	path := request.URL.Path[index:]
	file, err := os.Open(path)
	if err != nil {
		//http.Error(writer, err.Error(),http.StatusInternalServerError)
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		//panic(err)
		return err
	}

	writer.Write(all)
	return nil
}
