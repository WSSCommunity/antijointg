package server

import (
	"context"
	"errors"
	"io"
	"net/http"
	"testing"
	"time"

	tbapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/umputun/tg-spam/app/server/mocks"
)

func TestSpamRest_UnbanURL(t *testing.T) {
	tests := []struct {
		name   string
		url    string
		secret string
		want   string
	}{
		{"empty", "", "", "/unban?user=123&token=d68b50c4f0747630c33bc736bb3087b4c22f19dc645ec63b3bf90760c553e1ae"},
		{"test1", "http://localhost", "secret",
			"http://localhost/unban?user=123&token=71199ea8c011a49df546451e456ad10b0016566a53c4861bf849ec6b2ad2a0b7"},
		{"test2", "http://127.0.0.1:8080", "secret",
			"http://127.0.0.1:8080/unban?user=123&token=71199ea8c011a49df546451e456ad10b0016566a53c4861bf849ec6b2ad2a0b7"},
		{"test3", "http://127.0.0.1:8080", "secret2",
			"http://127.0.0.1:8080/unban?user=123&token=5385a71e8d5b65ea03e3da10175d78028ae59efd58811004e907baf422019b2e"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := SpamWeb{Config: Config{URL: tt.url, Secret: tt.secret}}
			res := srv.UnbanURL(123)
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestSpamRest_Run(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	mockAPI := &mocks.TbAPIMock{
		GetChatFunc: func(config tbapi.ChatInfoConfig) (tbapi.Chat, error) {
			if config.ChatConfig.SuperGroupUsername == "xxx" {
				return tbapi.Chat{}, errors.New("not found")
			}
			if config.ChatConfig.SuperGroupUsername == "@group" {
				return tbapi.Chat{ID: 10}, nil
			}
			return tbapi.Chat{ID: 123}, nil
		},
		RequestFunc: func(c tbapi.Chattable) (*tbapi.APIResponse, error) {
			if c.(tbapi.UnbanChatMemberConfig).UserID == 666 {
				return nil, errors.New("failed")
			}
			return &tbapi.APIResponse{}, nil
		},
	}

	srv, err := NewSpamWeb(mockAPI, Config{
		ListenAddr: ":9900",
		URL:        "http://localhost:9090",
		Secret:     "secret",
		TgGroup:    "group",
	})
	require.NoError(t, err)

	require.Equal(t, 1, len(mockAPI.GetChatCalls()))
	assert.Equal(t, "@group", mockAPI.GetChatCalls()[0].Config.ChatConfig.SuperGroupUsername)

	done := make(chan struct{})
	go func() {
		srv.Run(ctx)
		close(done)
	}()

	time.Sleep(100 * time.Millisecond) // wait for server to start

	t.Run("ping", func(t *testing.T) {
		mockAPI.ResetCalls()
		resp, err := http.Get("http://localhost:9900/ping")
		require.NoError(t, err)
		defer resp.Body.Close()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "text/plain", resp.Header.Get("Content-Type"))
		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		assert.Equal(t, "pong", string(body))
	})

	t.Run("unban forbidden, wrong token", func(t *testing.T) {
		mockAPI.ResetCalls()
		req, err := http.NewRequest("GET", "http://localhost:9900/unban?user=123&token=ssss", http.NoBody)
		require.NoError(t, err)
		client := &http.Client{Timeout: 5 * time.Second}
		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()
		assert.Equal(t, http.StatusForbidden, resp.StatusCode)
		assert.Equal(t, 0, len(mockAPI.RequestCalls()))
		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		t.Logf("body: %s", body)
		assert.Contains(t, string(body), "Error")
		assert.Equal(t, "text/html", resp.Header.Get("Content-Type"))
	})

	t.Run("unban failed, bad id", func(t *testing.T) {
		mockAPI.ResetCalls()
		req, err := http.NewRequest("GET",
			"http://localhost:9900/unban?user=xxx&token=71199ea8c011a49df546451e456ad10b0016566a53c4861bf849ec6b2ad2a0b7", http.NoBody)
		require.NoError(t, err)
		client := &http.Client{Timeout: 5 * time.Second}
		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()
		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.Equal(t, 0, len(mockAPI.RequestCalls()))
		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		t.Logf("body: %s", body)
		assert.Contains(t, string(body), "Error")
		assert.Equal(t, "text/html", resp.Header.Get("Content-Type"))
	})

	t.Run("unban failed, unban request failed", func(t *testing.T) {
		mockAPI.ResetCalls()
		req, err := http.NewRequest("GET",
			"http://localhost:9900/unban?user=666&token=4eeb1bfa92a5c9418e8708953daaba267f86df63281da9480c53206d4cb2be32", http.NoBody)
		require.NoError(t, err)
		client := &http.Client{Timeout: 5 * time.Second}
		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()
		assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
		assert.Equal(t, 1, len(mockAPI.RequestCalls()))
		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		t.Logf("body: %s", body)
		assert.Contains(t, string(body), "Error")
		assert.Equal(t, "text/html", resp.Header.Get("Content-Type"))

	})

	t.Run("unban allowed, matched token", func(t *testing.T) {
		mockAPI.ResetCalls()
		req, err := http.NewRequest("GET",
			"http://localhost:9900/unban?user=123&token=71199ea8c011a49df546451e456ad10b0016566a53c4861bf849ec6b2ad2a0b7", http.NoBody)
		require.NoError(t, err)
		client := &http.Client{Timeout: 5 * time.Second}
		resp, err := client.Do(req)
		require.NoError(t, err)
		defer resp.Body.Close()
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		require.Equal(t, 1, len(mockAPI.RequestCalls()))
		assert.Equal(t, int64(10), mockAPI.RequestCalls()[0].C.(tbapi.UnbanChatMemberConfig).ChatID)
		assert.Equal(t, int64(123), mockAPI.RequestCalls()[0].C.(tbapi.UnbanChatMemberConfig).UserID)
		body, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		t.Logf("body: %s", body)
		assert.Contains(t, string(body), "Success")
		assert.Equal(t, "text/html", resp.Header.Get("Content-Type"))
	})

	<-done
}
