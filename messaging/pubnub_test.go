package messaging

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testMessage1 = `PRISE EN MAIN - Le Figaro a pu approcher les nouveaux smartphones de Google. Voici nos premières observations. Le premier «smartphone conçu Google». Voilà comment a été présenté mardi le Pixel mardi. Il ne s'agit pas tout à fait de la première`
	testMessage2 = `Everybody copies everybody. It doesn't mean they're "out of ideas" or "in a technological cul-de-sac" - or at least it doesn't necessarily mean that - it does mean they want to make money and keep users.`
)

func BenchmarkEncodeNonASCIIChars(b *testing.B) {
	for i := 0; i < b.N; i++ {
		encodeNonASCIIChars(testMessage1)
		encodeNonASCIIChars(testMessage2)
	}
}

func TestEncodeNonASCIIChars(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{
			input:    testMessage1,
			expected: "PRISE EN MAIN - Le Figaro a pu approcher les nouveaux smartphones de Google. Voici nos premi\\u00e8res observations. Le premier \\u00absmartphone con\\u00e7u Google\\u00bb. Voil\\u00e0 comment a \\u00e9t\\u00e9 pr\\u00e9sent\\u00e9 mardi le Pixel mardi. Il ne s'agit pas tout \\u00e0 fait de la premi\\u00e8re",
		},
		{
			input:    testMessage2,
			expected: testMessage2, // no non-ascii characters here, so the string should be unchanged
		},
		{
			input:    "",
			expected: "",
		},
	}
	for _, tc := range cases {
		assert.Equal(t, encodeNonASCIIChars(tc.input), tc.expected)
	}
}
