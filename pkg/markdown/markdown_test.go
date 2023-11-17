package markdown

import "testing"

func TestConvertHtml(t *testing.T) {

	tcs := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "test 1",
			in:   `<a class="link" href="https://soundcloud.com/zakimusicofficial/tracks?utm_source=circadiangrowth.beehiiv.com&utm_medium=newsletter&utm_campaign=melodic-mastery-4-ways-music-can-improve-our-lives" target="_blank" rel="noopener noreferrer nofollow">my own electronic music</a> `,
			want: `[my own electronic music](https://soundcloud.com/zakimusicofficial/tracks?utm_source=circadiangrowth.beehiiv.com&utm_medium=newsletter&utm_campaign=melodic-mastery-4-ways-music-can-improve-our-lives)`,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {

			got, err := ConvertHtml(tc.in)
			if err != nil {
				t.Fatalf("ConvertHtml(%s) returned error: %v", tc.in, err)
			}

			if got != tc.want {
				t.Errorf("parse(%s) = %s; want %s", tc.in, got, tc.want)
			}
		})
	}

}
