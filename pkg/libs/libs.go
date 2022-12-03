package libs

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"os"
)

var hash_with_name = map[string]string{}
var name_with_hash = map[string]string{}

func HashFileContent(file []byte) string {
	h := sha256.New()
	h.Write(file)
	hash := fmt.Sprintf("%x", h.Sum(nil))
	return hash
}

func CheckIfFileExists(filename string, hash string) (bool, string) {
	if _, isHash := hash_with_name[hash]; isHash {
		return true, hash_with_name[hash]
	}
	if _, isFile := name_with_hash[filename]; isFile {
		return true, filename
	}
	return false, ""
}

func WriteToStore(
	name string, hash string, fileContent []byte,
	w http.ResponseWriter) {
	filename := "./store-files/" + name
	out, err := os.Create(filename)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	hash_with_name[hash] = name
	name_with_hash[name] = hash

	if err != nil {
		fmt.Fprintf(w, "Unable to create the file for writing. Check your write access privilege")
		return
	}

	_, err = out.Write(fileContent)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
}
