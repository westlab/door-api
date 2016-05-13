package job

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ikawaha/kagome/tokenizer"
	"github.com/kennygrant/sanitize"

	"github.com/westlab/door-api/model"
)

const (
	UserDicPath = "./userdic.txt"
)

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

// GetNouns gets nouns from text with wiki dictionary
func GetNouns(text string) (words []string) {
	tnz := tokenizer.New()

	udic, err := tokenizer.NewUserDic(UserDicPath)
	if err != nil {
		log.Println(err)
	}
	tnz.SetUserDic(udic)

	tokens := tnz.Tokenize(text)

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

// GetNouns gets nouns from text without wiki dictionary
func DummyGetNouns(text string) (words []string) {
	tnz := tokenizer.New()
	tokens := tnz.Tokenize(text)

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
