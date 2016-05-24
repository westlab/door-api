package job

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ikawaha/kagome/tokenizer"
	"github.com/kennygrant/sanitize"

	"github.com/westlab/door-api/model"
)

// wikipedia dictionary file
const (
	userDicPath = "./job/userdic.txt" // TODO: should be changed
)

// HTMLAnalyzer is analyzer for HTML page words
type HTMLAnalyzer struct {
	savedURLs map[string]string
	tokenizer *Tokenizer
}

// NewHTMLAnalyzer creates HTMLAnalyzer
func NewHTMLAnalyzer() *HTMLAnalyzer {
	savedURLs := make(map[string]string)
	tokenizer := NewTokenizer(userDicPath)
	return &HTMLAnalyzer{savedURLs, &tokenizer}
}

// Manage manages HTMLAnalyzer
func (htmlAnalyzer *HTMLAnalyzer) Manage(b *model.Browsing) {
	url := b.URL
	hashed := GetMD5Hash(url)
	_, ok := htmlAnalyzer.savedURLs[hashed]
	// If the URL is existed, do nothing
	if ok {
		return
	}

	html, err := DonwloadHTML(url)
	if err != nil {
		return // else to write?
	}
	wordList := htmlAnalyzer.tokenizer.GetNouns(RemoveHTMLTags(html), true)
	if ok := SaveWordsText(hashed, wordList); ok {
		htmlAnalyzer.savedURLs[hashed] = url
	}
}

// Tokenizer is a tokenizer and user dictionary model in kagome/tokenizer
type Tokenizer struct {
	tnz  tokenizer.Tokenizer
	udic tokenizer.UserDic
}

// NewTokenizer initializes Tokenizer
func NewTokenizer(path string) Tokenizer {
	newTnz := tokenizer.New()
	newUdic, err := tokenizer.NewUserDic(path)
	if err != nil {
		log.Println("new user dic: unexpected error")
	}
	return Tokenizer{tnz: newTnz, udic: newUdic}
}

// GetMD5Hash gets hash string from a string
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// DonwloadHTML downloads web page from given URL
func DonwloadHTML(url string) (html string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return html, err
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return html, err
	}
	html = string(contents)
	return html, nil
}

// RemoveHTMLTags removes html tags from text
func RemoveHTMLTags(html string) string {
	html = sanitize.HTML(html)
	return html
}

// SaveWordsText saves text file from string slice
func SaveWordsText(fileName string, strList []string) bool {
	content := []byte(strings.Join(strList, "\n"))          // should use another split char like ","?
	err := ioutil.WriteFile(fileName, content, os.ModePerm) // TODO: confirm save place
	if err != nil {
		return false
	}
	return true
}

// GetNouns gets nouns from text
func (t Tokenizer) GetNouns(text string, useUdic bool) (words []string) {
	if useUdic {
		t.tnz.SetUserDic(t.udic)
	}

	tokens := t.tnz.Tokenize(text)

	for _, token := range tokens {
		if token.Class == tokenizer.DUMMY {
			continue
		}
		if token.Pos() == "名詞" {
			words = append(words, token.Surface)
		}
	}

	return words
}

// WordCount counts words from words slice
func WordCount(words []string) []model.Count {
	wordmap := make(map[string]int64)
	for _, word := range words {
		if _, ok := wordmap[word]; ok {
			wordmap[word]++
		} else {
			wordmap[word] = 1
		}
	}

	// convert wordmap to counts
	counts := make([]model.Count, len(wordmap), len(wordmap))
	c := 0
	for key, value := range wordmap {
		counts[c] = model.Count{key, value}
		c++
	}

	return counts
}
