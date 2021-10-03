package srv

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type uploadHandler struct {
	HostAddr  string
	UploadDir string
}

func (h *uploadHandler) UploadRequestHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusBadRequest)
		return
	}

	filePath := h.UploadDir + "/" + header.Filename

	err = ioutil.WriteFile(filePath, data, 0777)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}

	fileLink := h.HostAddr + "/" + header.Filename
	fmt.Fprintln(w, fileLink)
}

func (h *uploadHandler) ListRequestHandler(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir(h.UploadDir)
	if err != nil {
		http.Error(w, "Unable to read files list", http.StatusBadRequest)
		return
	}

	first := true
	fmt.Fprint(w, "[")
	defer fmt.Fprintln(w, "]")

	for _, file := range files {
		select {
		case <-r.Context().Done():
			return
		default:
			if first {
				first = false
			} else {
				fmt.Fprintf(w, ",")
			}
			fmt.Fprint(w, file.Name())

			w.(http.Flusher).Flush()
		}
	}
}
