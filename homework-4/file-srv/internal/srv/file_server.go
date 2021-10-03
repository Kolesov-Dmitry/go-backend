package srv

import (
	"context"
	"net/http"
	"time"
)

type FileServer struct {
	uploadSrv   *http.Server
	downloadSrv *http.Server
}

func NewFileServer() *FileServer {
	uploadSrv := &http.Server{
		Addr:         ":80",
		Handler:      NewUploadHandler("localhost:8080", "upload"),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	downloadSrv := &http.Server{
		Addr:         ":8080",
		Handler:      http.FileServer(http.Dir("upload")),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return &FileServer{
		uploadSrv:   uploadSrv,
		downloadSrv: downloadSrv,
	}
}

func (f *FileServer) Start() {
	go f.downloadSrv.ListenAndServe()
	go f.uploadSrv.ListenAndServe()
}

func (f *FileServer) Shutdown(ctx context.Context) {
	f.uploadSrv.Shutdown(ctx)
	f.downloadSrv.Shutdown(ctx)
}
