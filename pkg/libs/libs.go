package libs

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"os"
)

func HashFileContent(file []byte) string {
	h := sha256.New()
	h.Write(file)
	hash := fmt.Sprintf("%x", h.Sum(nil))
	return hash
}

func WriteToStore(filename string, fileContent []byte, w http.ResponseWriter) {
	out, err := os.Create(filename)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	if err != nil {
		fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
		return
	}

	n2, err := out.Write(fileContent)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n2)

}
