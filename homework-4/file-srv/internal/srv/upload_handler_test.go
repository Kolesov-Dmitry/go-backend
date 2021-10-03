package srv_test

import (
	"encoding/json"
	"file-srv/internal/file"
	"file-srv/internal/srv"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MocDirReader struct{}

func (r *MocDirReader) Read(path string) ([]*file.FileInfo, error) {
	result := make([]*file.FileInfo, 0, 3)

	result = append(result, &file.FileInfo{
		Name: "readme.txt",
		Ext:  ".txt",
		Size: 1234,
	})

	result = append(result, &file.FileInfo{
		Name: "photo.png",
		Ext:  ".png",
		Size: 3213000,
	})

	result = append(result, &file.FileInfo{
		Name: "main.go",
		Ext:  ".go",
		Size: 952,
	})

	return result, nil
}

func TestListRequestHandler_FetchAll(t *testing.T) {
	handler := srv.NewUploadHandler("localhost", "upload")
	srv.DefaultDirReader = &MocDirReader{}

	server := httptest.NewServer(handler)
	defer server.Close()

	resp, err := server.Client().Get(server.URL + "/list")
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", resp.StatusCode, http.StatusOK)
	}

	var files []*file.FileInfo
	_ = json.NewDecoder(resp.Body).Decode(&files)

	if amount := len(files); amount != 3 {
		t.Errorf("handler returned wrong amount of files: got %v want %v", amount, 3)
	}
}

func TestListRequestHandler_Filter(t *testing.T) {
	handler := srv.NewUploadHandler("localhost", "upload")
	srv.DefaultDirReader = &MocDirReader{}

	server := httptest.NewServer(handler)
	defer server.Close()

	resp, err := server.Client().Get(server.URL + "/list?ext=png")
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", resp.StatusCode, http.StatusOK)
	}

	var files []*file.FileInfo
	_ = json.NewDecoder(resp.Body).Decode(&files)

	if amount := len(files); amount != 1 {
		t.Errorf("handler returned wrong amount of files: got %v want %v", amount, 1)
	}

	if name := files[0].Name; name != "photo.png" {
		t.Errorf("handler returned wrong file: got %v want %v", name, "photo.png")
	}
}
