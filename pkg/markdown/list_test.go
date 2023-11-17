package markdown

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestConvertUnorderedList(t *testing.T) {

	tcs := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "test 1",
			in:   `<ul><li>item 1</li><li>item 2</li></ul>`,
			want: "* item 1\n* item 2\n",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {

			doc, err := goquery.NewDocumentFromReader(strings.NewReader(tc.in))
			if err != nil {
				t.Fatalf("ConvertUnorderedList(%s) returned error: %v", tc.in, err)
			}

			got := convertUnorderedList(doc.Find("ul"))

			if got != tc.want {
				t.Errorf("convertUnorderedList(%s) = %s; want %s", tc.in, got, tc.want)
			}
		})
	}

}

func TestConvertOrderedList(t *testing.T) {

	tcs := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "test 1",
			in:   `<ol><li>item 1</li><li>item 2</li></ol>`,
			want: "1. item 1\n2. item 2\n",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {

			doc, err := goquery.NewDocumentFromReader(strings.NewReader(tc.in))
			if err != nil {
				t.Fatalf("ConvertOrderedList(%s) returned error: %v", tc.in, err)
			}

			got := convertOrderedList(doc.Find("ol"))

			if got != tc.want {
				t.Errorf("convertOrderedList(%s) = %s; want %s", tc.in, got, tc.want)
			}
		})
	}

}
