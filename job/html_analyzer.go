package job

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kennygrant/sanitize"
)

func DonwloadHTML(url string) (html string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return html, err
	}
	defer resp.Body.Close()
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return html, err
	}
	html = string(contents)
	return html, nil
}

func RemoveHTMLTags(html string) string {
	html = sanitize.HTML(html)
	return html
}
