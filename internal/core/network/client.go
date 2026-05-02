package network

import (
	"bytes"
	"net/http"
	"os"
)

func UploadFile(path string) {
	file, _ := os.ReadFile(path)

	req, _ := http.NewRequest("POST", "http://localhost:3722/upload", bytes.NewReader(file))
	req.Header.Set("X-Sambal-Name", path)

	http.DefaultClient.Do(req)
}
