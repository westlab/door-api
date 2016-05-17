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

func TestRemoveHTMLTags(t *testing.T) {
	html := `<html><head><meta charset="utf-8"></head>
	<body><div>WestLab</div></body></html>`
	text := strings.TrimSpace(RemoveHTMLTags(html))
	assert.Equal(t, "WestLab", text)
}

func TestElementAnalytics(t *testing.T) {
	text1 := `桜の花びらが舞う今日とっても茅ヶ崎で
	花びらとってもすごく綺麗でまさに茅ヶ崎で茅ヶ崎`
	tnz1 := NewTokenizer(UserDicPath)
	words1 := tnz1.GetNouns(text1, true)
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

	// wikipedia dictionary
	text2 := "横綱朝青龍"
	tnz2 := NewTokenizer(UserDicPath)
	words_not_using_wiki := tnz2.GetNouns(text2, false)
	assert.Equal(t, []string{"横綱", "朝", "青龍"}, words_not_using_wiki)
	words_using_wiki := tnz2.GetNouns(text2, true)
	assert.Equal(t, []string{"横綱", "朝青龍"}, words_using_wiki)
}
