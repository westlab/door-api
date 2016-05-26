package job

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/ikawaha/kagome/tokenizer"
	"github.com/kennygrant/sanitize"

	"github.com/westlab/door-api/common"
	"github.com/westlab/door-api/context"
	"github.com/westlab/door-api/model"
)

// wikipedia dictionary file
const (
	userDicPath = "./job/userdic.txt" // TODO: should be changed
)

// HTMLAnalyzer is analyzer for HTML page words
type HTMLAnalyzer struct {
	tokenizer *Tokenizer
}

// NewHTMLAnalyzer creates HTMLAnalyzer
func NewHTMLAnalyzer() *HTMLAnalyzer {
	tokenizer := NewTokenizer(userDicPath)
	return &HTMLAnalyzer{&tokenizer}
}

// Manage manages HTMLAnalyzer
func (htmlAnalyzer *HTMLAnalyzer) Manage(b *model.Browsing) {
	url := b.URL
	hashed := common.GetMD5Hash(url)
	cxt := context.GetContext()
	ok := common.IsFileExist(hashed, cxt.GetConf().WordsPath)
	if ok { // If the file is existed, do nothing
		return
	}

	html, err := DonwloadHTML(url)
	if err != nil {
		return // else to write?
	}
	html = RemoveAnyTagData(html, "script") // remove scirpt tag data
	sanitizedText := RemoveHTMLTags(html)
	wordList := htmlAnalyzer.tokenizer.GetNouns(sanitizedText, true)
	SaveWordsText(hashed, wordList)

	// save words to Word table
	counts := WordCount(wordList)
	var words []*model.Word
	for _, count := range counts {
		w := model.NewWord(count.Name, count.Count)
		words = append(words, w)
	}
	err = model.WordBulkInsert(words)
	if err != nil {
		log.Println("WordBulkInsert: ", err)
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

// RemoveAnyTagData removes tag's data from selected tag name
func RemoveAnyTagData(html string, tags ...string) string {
	for _, tag := range tags {
		re, _ := regexp.Compile("\\<" + tag + "[\\S\\s]+?\\</" + tag + "\\>")
		html = re.ReplaceAllString(html, "")
	}
	return html
}

// RemoveHTMLTags removes html tags from text
func RemoveHTMLTags(html string) string {
	html = sanitize.HTML(html)
	return html
}

// SaveWordsText saves text file from string slice
func SaveWordsText(fileName string, strList []string) bool {
	cxt := context.GetContext()
	content := []byte(strings.Join(strList, "\n")) // should use another split char like ","?
	fpath := filepath.Join(cxt.GetConf().WordsPath, fileName)
	err := ioutil.WriteFile(fpath, content, os.ModePerm) // conf.WordsPath is the dir to save words files
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
