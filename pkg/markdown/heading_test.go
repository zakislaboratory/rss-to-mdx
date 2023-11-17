package markdown

import (
	"fmt"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestConvertHeading(t *testing.T) {
	tcs := []struct {
		name    string
		level   int
		rawHTML string
		want    string
	}{
		{
			name:    "test 1",
			level:   2,
			rawHTML: `<h2 level="2">This is a heading</h2>`,
			want:    `## This is a heading` + "\n",
		},
		{
			name:    "test 2",
			level:   3,
			rawHTML: `<h3 level="3">This is a heading</h3>`,
			want:    `### This is a heading` + "\n",
		},
		{
			name:    "test 3",
			level:   4,
			rawHTML: `<h4 level="4">This is a heading</h4>`,
			want:    `#### This is a heading` + "\n",
		},
	}

	for _, tc := range tcs {

		t.Run(tc.name, func(t *testing.T) {

			doc, err := goquery.NewDocumentFromReader(strings.NewReader(tc.rawHTML))
			if err != nil {
				t.Fatalf("ConvertHeading(%s) returned error: %v", tc.rawHTML, err)
			}

			h := doc.Find("h" + fmt.Sprint(tc.level)).First()

			got := NewHeading(h)

			if got.Markdown() != tc.want {
				t.Errorf("convertHeading(%s) = %s; want %s", tc.rawHTML, got.Markdown(), tc.want)
			}

		})
	}

}
