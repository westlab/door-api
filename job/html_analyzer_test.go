package job

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// wikipedia dictionary file
const (
	UserDicPath = "./userdic.txt"
)

func TestRemoveAnyTagData(t *testing.T) {
	html := `
<html>
<head><title>TITLE</title></head>
<body>
<script type="text/javascript">
function cst(cat, evt, targetUrl) {
if (window._gaq) {
_gaq.push(['_trackSocial', cat, evt, targetUrl]);
} else {
console.log("trackSocial(cat)");
}
};
</script>
BODY
</body>
</html> `

	ohtml1 := RemoveAnyTagData(html, "script")
	text1 := strings.TrimSpace(RemoveHTMLTags(ohtml1))
	assert.Equal(t, "TITLEBODY", text1)

	ohtml2 := RemoveAnyTagData(html, "script", "head")
	text2 := strings.TrimSpace(RemoveHTMLTags(ohtml2))
	assert.Equal(t, "BODY", text2)
}

func TestRemoveHTMLTags(t *testing.T) {
	html := `<html><head><meta charset="utf-8"></head>
	<body><div>WestLab</div></body></html>`
	text := strings.TrimSpace(RemoveHTMLTags(html))
	assert.Equal(t, "WestLab", text)
}

func TestElementAnalytics(t *testing.T) {
	tnz := NewTokenizer(UserDicPath)

	text1 := `桜の花びらが舞う今日とっても茅ヶ崎で
	花びらとってもすごく綺麗でまさに茅ヶ崎で茅ヶ崎`
	words1 := tnz.GetNouns(text1, true)
	assert.Equal(t, []string{"桜", "花びら", "今日", "茅ヶ崎", "花びら", "綺麗", "茅ヶ崎", "茅ヶ崎"}, words1)
	counts := WordCount(words1)
	for _, count := range counts {
		if count.Name == "桜" {
			assert.Equal(t, 1, int(count.Count))
		}
		if count.Name == "花びら" {
			assert.Equal(t, 2, int(count.Count))
		}
		if count.Name == "茅ヶ崎" {
			assert.Equal(t, 3, int(count.Count))
		}
	}

	// use wikipedia dictionary or not
	text2 := "横綱朝青龍"
	words_not_using_wiki := tnz.GetNouns(text2, false)
	assert.Equal(t, []string{"横綱", "朝", "青龍"}, words_not_using_wiki)
	words_using_wiki := tnz.GetNouns(text2, true)
	assert.Equal(t, []string{"横綱", "朝青龍"}, words_using_wiki)
}
