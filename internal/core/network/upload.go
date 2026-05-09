package network

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func UploadFile(path string, url string) error {

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	// progress tracker
	progress := &ProgressTracker{
		Total: info.Size(),
	}

	// duplicate stream into progress tracker
	reader := io.TeeReader(file, progress)

	req, err := http.NewRequest("POST", url, reader)

	if err != nil {
		return err
	}

	req.ContentLength = info.Size()
	req.Header.Set("X-Sambal-Name", info.Name())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("\nUpload completed.")

	return nil
}
