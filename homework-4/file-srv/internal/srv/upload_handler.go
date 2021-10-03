package srv

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"file-srv/internal/dir"
)

var DefaultDirReader DirReader = &dir.Reader{}

// UploadHandler is a file server requests handler
type UploadHandler struct {
	hostAddr  string
	uploadDir string
	mux       *http.ServeMux
}

// NewUploadHandler makes new UploadHandler instance
// Inputs:
//   hostAddr  - file server host IP address
//   uploadDir - path to the directory where uploaded files will be stored
func NewUploadHandler(hostAddr string, uploadDir string) *UploadHandler {
	handler := &UploadHandler{
		hostAddr:  hostAddr,
		uploadDir: uploadDir,
		mux:       http.NewServeMux(),
	}

	handler.mux.HandleFunc("/upload", handler.UploadRequestHandler)
	handler.mux.HandleFunc("/list", handler.ListRequestHandler)

	return handler
}

// ServeHTTP wrapper
func (h *UploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.ServeHTTP(w, r)
}

// UploadRequestHandler '/upload' request handler
func (h *UploadHandler) UploadRequestHandler(w http.ResponseWriter, r *http.Request) {
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

	filePath := h.uploadDir + "/" + header.Filename

	err = ioutil.WriteFile(filePath, data, 0777)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		return
	}

	fileLink := h.hostAddr + "/" + header.Filename
	fmt.Fprintln(w, fileLink)
}

// ListRequestHandler '/list' request handler
func (h *UploadHandler) ListRequestHandler(w http.ResponseWriter, r *http.Request) {
	files, err := DefaultDirReader.Read(h.uploadDir)
	if err != nil {
		http.Error(w, "Unable to read files list", http.StatusBadRequest)
		return
	}

	ext := r.URL.Query().Get("ext")
	first := true
	fmt.Fprint(w, "[")
	defer fmt.Fprintln(w, "]")

	enc := json.NewEncoder(w)
	for _, file := range files {
		select {
		case <-r.Context().Done():
			return

		default:
			if len(ext) != 0 && !strings.HasSuffix(file.Name, ext) {
				// skip files which have different extention
				continue
			}

			// put comma if needed
			if first {
				first = false
			} else {
				fmt.Fprintf(w, ",")
			}

			// append file name to the output stream
			_ = enc.Encode(file)
			w.(http.Flusher).Flush()
		}
	}
}
