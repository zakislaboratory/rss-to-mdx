package markdown

import (
	"regexp"
	"testing"
)

func Test(t *testing.T) {

	tcs := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "Single Link",
			in:   `<a class="link" href="https://soundcloud.com/zakimusicofficial/tracks?utm_source=circadiangrowth.beehiiv.com&utm_medium=newsletter&utm_campaign=melodic-mastery-4-ways-music-can-improve-our-lives" target="_blank" rel="noopener noreferrer nofollow">my own electronic music</a>`,
			want: `[my own electronic music](https://soundcloud.com/zakimusicofficial/tracks?utm_source=circadiangrowth.beehiiv.com&utm_medium=newsletter&utm_campaign=melodic-mastery-4-ways-music-can-improve-our-lives)`,
		},
		{
			name: "Nested link",
			in:   `<p class="paragraph" style="text-align:start;">I even produce some of <a class="link" href="https://soundcloud.com/zakimusicofficial/tracks?utm_source=circadiangrowth.beehiiv.com&utm_medium=newsletter&utm_campaign=melodic-mastery-4-ways-music-can-improve-our-lives" target="_blank" rel="noopener noreferrer nofollow">my own electronic music</a> from time to time (shameless self plug).</p>`,
			want: `I even produce some of [my own electronic music](https://soundcloud.com/zakimusicofficial/tracks?utm_source=circadiangrowth.beehiiv.com&utm_medium=newsletter&utm_campaign=melodic-mastery-4-ways-music-can-improve-our-lives) from time to time (shameless self plug).`,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {

			got := NewMarkdownDocument(tc.in)

			content, err := got.Content()
			if err != nil {
				t.Fatalf("got.Content() returned error: %v", err)
			}

			if content != tc.want {
				t.Errorf("[%s] = %s; want %s", tc.name, content, tc.want)
			}
		})
	}

}

func TestPatternMatches(t *testing.T) {

	html := `<p>Hello world</p>
<p>This is a complete match</p>
<p>This is a partial match</p>
	`

	md := NewMarkdownDocument(html)

	// Define the regex pattern for the "complete match" line
	completeMatchPattern := `This is a complete match`
	partialMatchPattern := `partial match`

	// Compile the regex pattern
	re := regexp.MustCompile(completeMatchPattern)
	re2 := regexp.MustCompile(partialMatchPattern)

	// Add the regex pattern to the list of patterns to remove
	md.RemoveMatches(re)
	md.RemoveMatches(re2)

	content, err := md.Content()
	if err != nil {
		t.Fatalf("md.Content() returned error: %v", err)
	}

	// Check that the "complete match" line was removed
	if content != "Hello world" {
		t.Errorf("md.String() = %s; want %s", content, "Hello world")
	}
}
