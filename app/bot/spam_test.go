package bot

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/umputun/tg-spam/app/bot/mocks"
)

func TestFilter_OnMessage(t *testing.T) {
	mockedHTTPClient := &mocks.HTTPClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			if strings.Contains(req.URL.String(), "101") {
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewBufferString(`{"ok": true, "description": "Is a spammer"}`)),
				}, nil
			}
			return &http.Response{
				StatusCode: 200,
				Body:       io.NopCloser(bytes.NewBufferString(`{"ok": false, "description": "Not a spammer"}`)),
			}, nil
		},
	}

	filter, err := NewSpamFilter(context.Background(), SpamParams{
		SpamSamplesFile:     "testdata/spam-samples.txt", // "win free iPhone\nlottery prize
		HamSamplesFile:      "testdata/ham-samples.txt",
		StopWordsFile:       "testdata/stop-words.txt",
		ExcludedTokensFile:  "testdata/spam-exclude-token.txt",
		SimilarityThreshold: 0.5,
		MinMsgLen:           5,
		Dry:                 false,
		HTTPClient:          mockedHTTPClient,
		SpamMsg:             "this is spam! go to ban",
		MaxAllowedEmoji:     2,
	})
	require.NoError(t, err)

	tests := []struct {
		name     string
		msg      Message
		expected Response
	}{
		{
			"good message",
			Message{From: User{ID: 1, Username: "john", DisplayName: "John"}, Text: "Hello, how are you?", ID: 1},
			Response{},
		},
		{
			"emoji spam",
			Message{From: User{ID: 4, Username: "john", DisplayName: "John"}, Text: "Hello 😁🐶🍕 how are you? ", ID: 4},
			Response{Text: "this is spam! go to ban: \"John\" (4)", Send: true,
				BanInterval: permanentBanDuration, ReplyTo: 4, DeleteReplyTo: true,
				User: User{ID: 4, Username: "john", DisplayName: "John"}},
		},
		{
			"similarity spam",
			Message{From: User{ID: 2, Username: "spammer", DisplayName: "Spammer"}, Text: "Win a free iPhone now!", ID: 2},
			Response{Text: "this is spam! go to ban: \"Spammer\" (2)", Send: true,
				ReplyTo: 2, BanInterval: permanentBanDuration, DeleteReplyTo: true,
				User: User{ID: 2, Username: "spammer", DisplayName: "Spammer"},
			},
		},
		{
			"classifier spam",
			Message{From: User{ID: 2, Username: "spammer", DisplayName: "Spammer"}, Text: "free gift for you", ID: 2},
			Response{Text: "this is spam! go to ban: \"Spammer\" (2)", Send: true,
				ReplyTo: 2, BanInterval: permanentBanDuration, DeleteReplyTo: true,
				User: User{ID: 2, Username: "spammer", DisplayName: "Spammer"},
			},
		},
		{
			"CAS spam",
			Message{From: User{ID: 101, Username: "spammer", DisplayName: "blah"}, Text: "something something", ID: 10},
			Response{Text: "this is spam! go to ban: \"blah\" (101)", Send: true,
				ReplyTo: 10, BanInterval: permanentBanDuration, DeleteReplyTo: true,
				User: User{ID: 101, Username: "spammer", DisplayName: "blah"},
			},
		},
		{
			"stop words spam emoji",
			Message{From: User{ID: 102, Username: "spammer", DisplayName: "blah"}, Text: "something пишите в лс something", ID: 10},
			Response{Text: "this is spam! go to ban: \"blah\" (102)", Send: true,
				ReplyTo: 10, BanInterval: permanentBanDuration, DeleteReplyTo: true,
				User: User{ID: 102, Username: "spammer", DisplayName: "blah"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, filter.OnMessage(test.msg))
		})
	}
}

func TestIsSpam(t *testing.T) {
	spamSamples := strings.NewReader("win free iPhone\nlottery prize")
	filter := SpamFilter{
		tokenizedSpam: []map[string]int{},
		lock:          sync.RWMutex{},
	}

	tests := []struct {
		name      string
		message   string
		threshold float64
		expected  bool
	}{
		{"Not Spam", "Hello, how are you?", 0.5, false},
		{"Exact Match", "Win a free iPhone now!", 0.5, true},
		{"Similar Match", "You won a lottery prize!", 0.3, true},
		{"High Threshold", "You won a lottery prize!", 0.9, false},
		{"Partial Match", "win free", 0.9, false},
		{"Low Threshold", "win free", 0.8, true},
	}

	err := filter.loadSpamSamples(spamSamples)
	require.NoError(t, err)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			filter.SimilarityThreshold = test.threshold // Update threshold for each test case
			assert.Equal(t, test.expected, filter.isSpamSimilarityHigh(test.message))
		})
	}
}

// nolint
func TestTooManyEmojis(t *testing.T) {
	filter := SpamFilter{
		SpamParams: SpamParams{MaxAllowedEmoji: 2},
	}

	tests := []struct {
		name  string
		input string
		count int
		spam  bool
	}{
		{"NoEmoji", "Hello, world!", 0, false},
		{"OneEmoji", "Hi there 👋", 1, false},
		{"TwoEmojis", "Good morning 🌞🌻", 2, false},
		{"Mixed", "👨‍👩‍👧‍👦 Family emoji", 1, false},
		{"EmojiSequences", "🏳️‍🌈 Rainbow flag", 1, false},
		{"TextAfterEmoji", "😊 Have a nice day!", 1, false},
		{"OnlyEmojis", "😁🐶🍕", 3, true},
		{"WithCyrillic", "Привет 🌞 🍕 мир! 👋", 3, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isSpam, count := filter.tooManyEmojis(tt.input, 2)
			assert.Equal(t, tt.count, count)
			assert.Equal(t, tt.spam, isSpam)
		})
	}
}

func TestStopWords(t *testing.T) {
	filter := &SpamFilter{
		stopWords: []string{"в личку", "всем привет"},
	}

	tests := []struct {
		name     string
		message  string
		expected bool
	}{
		{
			name:     "Stop word present",
			message:  "Hello, please send me a message в личку",
			expected: true,
		},
		{
			name:     "Stop word present with emoji",
			message:  "👋Всем привет\nИщу амбициозного человека к се6е в команду\nКто в поисках дополнительного заработка или хочет попробовать себя в новой  сфере деятельности! 👨🏻\u200d💻\nПишите в лс✍️",
			expected: true,
		},
		{
			name:     "No stop word present",
			message:  "Hello, how are you?",
			expected: false,
		},
		{
			name:     "Case insensitive stop word present",
			message:  "Hello, please send me a message В ЛИЧКУ",
			expected: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, filter.hasStopWords(test.message))
		})
	}
}

func TestSpam_isCasSpam(t *testing.T) {

	tests := []struct {
		name           string
		mockResp       string
		mockStatusCode int
		expected       bool
	}{
		{
			name:           "User is not a spammer",
			mockResp:       `{"ok": false, "description": "Not a spammer"}`,
			mockStatusCode: 200,
			expected:       false,
		},
		{
			name:           "User is a spammer",
			mockResp:       `{"ok": true, "description": "Is a spammer"}`,
			mockStatusCode: 200,
			expected:       true,
		},
		{
			name:           "HTTP error",
			mockResp:       "",
			mockStatusCode: 500,
			expected:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockedHTTPClient := &mocks.HTTPClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					return &http.Response{
						StatusCode: tt.mockStatusCode,
						Body:       io.NopCloser(bytes.NewBufferString(tt.mockResp)),
					}, nil
				},
			}

			s := SpamFilter{SpamParams: SpamParams{
				CasAPI:     "http://localhost",
				HTTPClient: mockedHTTPClient,
			}}

			msg := Message{
				From: User{
					ID:          1,
					Username:    "testuser",
					DisplayName: "Test User",
				},
				ID:   1,
				Text: "Hello",
			}

			isSpam := s.isCasSpam(msg.From.ID)
			assert.Equal(t, tt.expected, isSpam)
		})
	}
}

func Test_tokenChan(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{name: "empty", input: "", expected: []string{}},
		{name: "token per line", input: "hello\nworld", expected: []string{"hello", "world"}},
		{name: "token per line", input: "hello 123\nworld", expected: []string{"hello 123", "world"}},
		{name: "token per line with spaces", input: "hello \n world", expected: []string{"hello", "world"}},
		{name: "tokens comma separated", input: "\"hello\",\"world\"\nsomething", expected: []string{"hello", "world", "something"}},
		{name: "tokens comma separated, extra EOL", input: "\"hello\",world\nsomething\n", expected: []string{"hello", "world", "something"}},
		{name: "tokens comma separated, empty tokens", input: "\"hello\",world,\"\"\nsomething\n ", expected: []string{"hello", "world", "something"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := tokenChan(bytes.NewBufferString(tt.input))
			res := []string{}
			for token := range ch {
				res = append(res, token)
			}
			assert.Equal(t, tt.expected, res)
		})
	}
}

func TestSpamFilter_tokenize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]int
	}{
		{name: "empty", input: "", expected: map[string]int{}},
		{name: "no filters or cleanups", input: "hello world", expected: map[string]int{"hello": 1, "world": 1}},
		{name: "with excluded tokens", input: "hello world the she", expected: map[string]int{"hello": 1, "world": 1}},
		{name: "with short tokens", input: "hello world the she a or", expected: map[string]int{"hello": 1, "world": 1}},
		{name: "with repeated tokens", input: "hello world hello world", expected: map[string]int{"hello": 2, "world": 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SpamFilter{
				excludedTokens: []string{"the", "she"},
			}
			assert.Equal(t, tt.expected, s.tokenize(tt.input))
		})
	}
}

func TestSpamFilter_UpdateSpam(t *testing.T) {

	upd := &mocks.SampleUpdater{
		AppendFunc: func(msg string) error {
			return nil
		},
	}

	s := SpamFilter{
		spamSamplesUpd: upd,
		tokenizedSpam:  []map[string]int{},
		spamClassifier: Classifier{},
	}
	s.spamClassifier.Reset()

	err := s.UpdateSpam("some spam message")
	require.NoError(t, err)
	assert.Equal(t, []map[string]int{{"some": 1, "spam": 1, "message": 1}}, s.tokenizedSpam)
	assert.Equal(t, 1, s.spamClassifier.NAllDocument)

	err = s.UpdateSpam("more things")
	require.NoError(t, err)
	assert.Equal(t, []map[string]int{{"some": 1, "spam": 1, "message": 1}, {"more": 1, "things": 1}}, s.tokenizedSpam)
	assert.Equal(t, 2, s.spamClassifier.NAllDocument)
}

func TestSpamFilter_UpdateHam(t *testing.T) {

	upd := &mocks.SampleUpdater{
		AppendFunc: func(msg string) error {
			return nil
		},
	}

	s := SpamFilter{
		hamSamplesUpd:  upd,
		tokenizedSpam:  []map[string]int{},
		spamClassifier: Classifier{},
	}
	s.spamClassifier.Reset()

	err := s.UpdateHam("some spam message")
	require.NoError(t, err)
	assert.Equal(t, []map[string]int{}, s.tokenizedSpam)
	assert.Equal(t, 1, s.spamClassifier.NAllDocument)
	assert.Equal(t, 1, len(upd.AppendCalls()))

	err = s.UpdateHam("more things")
	require.NoError(t, err)
	assert.Equal(t, []map[string]int{}, s.tokenizedSpam)
	assert.Equal(t, 2, s.spamClassifier.NAllDocument)
	assert.Equal(t, 2, len(upd.AppendCalls()))
}

func TestSpamFilter_loadDynFiles(t *testing.T) {
	count := 0
	upd := &mocks.SampleUpdater{
		AppendFunc: func(msg string) error {
			return nil
		},
		ReaderFunc: func() (io.ReadCloser, error) {
			count++
			if count <= 2 { // two calls, for spam, one for ham
				return io.NopCloser(bytes.NewBufferString("spam1 spam2 spam3\nspam4")), nil
			}
			return io.NopCloser(bytes.NewBufferString("ham1\nham2\nham3")), nil
		},
	}

	s := SpamFilter{
		spamSamplesUpd: upd,
		hamSamplesUpd:  upd,
		tokenizedSpam:  []map[string]int{},
		spamClassifier: Classifier{},
	}
	s.spamClassifier.Reset()

	err := s.loadDynFiles()
	require.NoError(t, err)
	assert.Equal(t, []map[string]int{{"spam1": 1, "spam2": 1, "spam3": 1}, {"spam4": 1}}, s.tokenizedSpam)
	assert.Equal(t, 2+3, s.spamClassifier.NAllDocument)
	assert.Equal(t, 0, len(upd.AppendCalls()))
	assert.Equal(t, 2+1, len(upd.ReaderCalls()))
}
