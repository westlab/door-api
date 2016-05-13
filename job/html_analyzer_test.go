package job

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveHTMLTags(t *testing.T) {
	html := `<html><head><meta charset="utf-8"></head>
	<body><div>WestLab</div></body></html>`
	text := strings.TrimSpace(RemoveHTMLTags(html))
	assert.Equal(t, "WestLab", text)
}

func TestElementAnalytics(t *testing.T) {
	text := `桜の花びらが舞う今日とっても茅ヶ崎で
	花びらとってもすごく綺麗でまさに茅ヶ崎で茅ヶ崎`
	words := GetNouns(text)
	assert.Equal(t, []string{"桜", "花びら", "今日", "茅ヶ崎", "花びら", "綺麗", "茅ヶ崎", "茅ヶ崎"}, words)
	counts := WordCount(words)
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

	// wikipedia dictionary
	text2 := "横綱朝青龍"
	words_not_using_wiki := DummyGetNouns(text2)
	assert.Equal(t, []string{"横綱", "朝", "青龍"}, words_not_using_wiki)
	words_using_wiki := GetNouns(text2)
	assert.Equal(t, []string{"横綱", "朝青龍"}, words_using_wiki)
}
