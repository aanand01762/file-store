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

func WriteToStore(name string, hash string, fileContent []byte, w http.ResponseWriter) {
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
func ReplaceInStore(name string, hash string, fileContent []byte, w http.ResponseWriter) {
	var returnMsg string

	//if there is a similar file with same name, remove that file
	if _, isFile := name_with_hash[name]; isFile {
		mapped_hash := name_with_hash[name]
		path := "./store-files/" + name
		err := os.Remove(path)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		delete(name_with_hash, name)
		delete(hash_with_name, mapped_hash)
		returnMsg = ("Updated the content of the file '" + name +
			"' with latest value")

		//Else if there is a similar file with same content, remove that file
	} else if _, isHash := hash_with_name[hash]; isHash {
		mapped_name := hash_with_name[hash]
		path := "./store-files/" + mapped_name
		err := os.Remove(path)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}

		delete(hash_with_name, hash)
		delete(name_with_hash, mapped_name)
		returnMsg = ("Changed the file name from: '" + mapped_name +
			"' to new file name: '" + name +
			"' because both files had same content")
	}

	//Now create new file with updated name and content
	new_filename := "./store-files/" + name
	out, err := os.Create(new_filename)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	hash_with_name[hash] = name
	name_with_hash[name] = hash

	if err != nil {
		fmt.Fprintf(
			w, "Unable to create the file for writing while "+
				"replacing. Check your write access privilege")
		return
	}

	_, err = out.Write(fileContent)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintf(w, returnMsg)
}
