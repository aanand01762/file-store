package libs

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var hash_with_name = map[string]string{}
var name_with_hash = map[string]string{}
var words_with_count = map[string]int{}
var count_with_word = map[int]string{}

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

func GetHashOfFile(filename string) string {
	if _, ok := name_with_hash[filename]; ok {
		return name_with_hash[string(filename)]
	} else {
		return ""
	}
}

func WriteToStore(name string, hash string, fileContent []byte, w http.ResponseWriter) {
	dir := "./store-files/"
	filename := dir + name
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
	UpdateFreqWithWord(dir, w)
}
func RemoveFile(filename string, hash string, w http.ResponseWriter) {
	dir := "./store-files/"
	path := dir + filename
	err := os.Remove(path)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	delete(name_with_hash, filename)
	delete(hash_with_name, hash)
	UpdateFreqWithWord(dir, w)

}
func ReplaceInStore(name string, hash string, fileContent []byte, w http.ResponseWriter) {
	var returnMsg string

	//if there is a similar file with same name, remove that file
	if _, isFile := name_with_hash[name]; isFile {
		mapped_hash := name_with_hash[name]
		RemoveFile(name, mapped_hash, w)
		returnMsg = ("Updated the content of the file '" + name +
			"' with latest value")

		//Else if there is a similar file with same content, remove that file
	} else if _, isHash := hash_with_name[hash]; isHash {
		mapped_name := hash_with_name[hash]
		RemoveFile(mapped_name, hash, w)
		returnMsg = ("Changed file name from: '" + mapped_name +
			"' to new file name: '" + name +
			"' because both files had same content")
	}

	//Now create new file with updated name and content
	new_filename := "./store-files/" + name
	out, err := os.Create(new_filename)
	if err != nil {
		fmt.Fprintf(
			w, "Unable to create the file for writing while "+
				"replacing. Check your write access privilege")
		return
	}
	hash_with_name[hash] = name
	name_with_hash[name] = hash

	_, err = out.Write(fileContent)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	dir := "./store-files/"
	UpdateFreqWithWord(dir, w)
	fmt.Fprintf(w, returnMsg)
}

func ReadFileWords(filename string, w http.ResponseWriter) {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintln(w, err)
		log.Println("Error: ", err)
		return
	}
	scanner := bufio.NewScanner(file)

	//regex to remove all special characters
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		fmt.Fprintln(w, err)
		log.Fatal(err)
	}

	// Setting Split method to ScanWords.
	scanner.Split(bufio.ScanWords)

	// Scan all words from the file.
	for scanner.Scan() {
		//remove all spaces and special chars
		word := strings.TrimSpace(reg.ReplaceAllString(scanner.Text(), ""))
		if len(word) > 0 {
			words_with_count[strings.ToLower(word)] += 1
		}
	}

	file.Close()

}

func UpdateWordsWithCounts(dir string, w http.ResponseWriter) {
	words_with_count = map[string]int{}
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}

	for _, file := range files {
		ReadFileWords(dir+file.Name(), w)
	}

}

func UpdateFreqWithWord(dir string, w http.ResponseWriter) {
	count_with_word = map[int]string{}
	UpdateWordsWithCounts(dir, w)
	for key, val := range words_with_count {
		count_with_word[val] = key
	}

}
func GetAllWordCount(dir string, w http.ResponseWriter) int {
	words_with_count := GetWordsWithCounts()
	sum := 0
	for _, val := range words_with_count {
		sum += val
	}
	return sum
}

func GetWordsWithCounts() map[string]int {
	return words_with_count
}

func GetFreqwithCounts() map[int]string {
	return count_with_word
}
