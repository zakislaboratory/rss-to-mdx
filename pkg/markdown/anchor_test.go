package markdown

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestConvertAnchor(t *testing.T) {

	tcs := []struct {
		name    string
		rawHTML string
		want    string
	}{
		{
			name:    "test 1",
			rawHTML: `<a class="link" href="https://soundcloud.com/zakimusicofficial/tracks?utm_source=circadiangrowth.beehiiv.com&utm_medium=newsletter&utm_campaign=melodic-mastery-4-ways-music-can-improve-our-lives" target="_blank" rel="noopener noreferrer nofollow">my own electronic music</a> `,
			want:    `[my own electronic music](https://soundcloud.com/zakimusicofficial/tracks?utm_source=circadiangrowth.beehiiv.com&utm_medium=newsletter&utm_campaign=melodic-mastery-4-ways-music-can-improve-our-lives)`,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {

			doc, err := goquery.NewDocumentFromReader(strings.NewReader(tc.rawHTML))
			if err != nil {
				t.Fatalf("ConvertAnchor(%s) returned error: %v", tc.rawHTML, err)
			}

			a := doc.Find("a").First()

			got := convertAnchor(a)

			if got.Markdown() != tc.want {
				t.Errorf("convertAnchor(%s) = %s; want %s", tc.rawHTML, got.Markdown(), tc.want)
			}

		})
	}

}
