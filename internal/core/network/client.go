package network

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func UploadFile(path string) {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Upload failed:", err)
		return
	}

	req, _ := http.NewRequest("POST", "http://localhost:3722/upload", bytes.NewReader(file))
	req.Header.Set("X-Sambal-Name", path)

	http.DefaultClient.Do(req)
}
