package job

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kennygrant/sanitize"
)

// DonwloadHTML downloads web page from given URL
func DonwloadHTML(url string) (html string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Error(err)
		return html, err
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
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
