package dictionary_test

import (
	"context"
	"errors"
	"testing"

	"github.com/firamisu/louis/internal/dictclient"
	"github.com/firamisu/louis/internal/dictionary"
	"github.com/firamisu/louis/internal/domain"
	"github.com/go-playground/assert/v2"
)

type mockClient struct {
	resp []dictclient.DictResponse
	err  error
}

func (m *mockClient) FetchWord(ctx context.Context, word string) ([]dictclient.DictResponse, error) {
	return m.resp, m.err
}

func TestGetWord(t *testing.T) {
	ctx := context.Background()

	dResp := dictclient.DictResponse{
		Word:     "hello",
		Phonetic: "həˈləʊ",
		Origin:   "Old English",
		Meanings: []dictclient.Meaning{
			{
				PartOfSpeech: "interjection",
				Definitions: []dictclient.Definition{
					{
						Definition: "used as a greeting",
						Example:    "Hello! How are you?",
					},
				},
			},
		},
	}

	expectedEntry := domain.Entry{
		Word:     dResp.Word,
		Phonetic: dResp.Phonetic,
		Origin:   dResp.Origin,
		Meanings: []domain.Meaning{
			{
				PartOfSpeech: dResp.Meanings[0].PartOfSpeech,
				Definitions: []domain.Definition{
					{
						Definition: dResp.Meanings[0].Definitions[0].Definition,
						Example:    dResp.Meanings[0].Definitions[0].Example,
					},
				},
			},
		},
	}

	testCases := []struct {
		name      string
		mockResp  []dictclient.DictResponse
		mockErr   error
		wantEntry domain.Entry
		wantErr   error
		inputWord string
	}{
		{
			name:      "success",
			mockResp:  []dictclient.DictResponse{dResp},
			mockErr:   nil,
			wantEntry: expectedEntry,
			wantErr:   nil,
			inputWord: "hello",
		},
		{
			name:      "not found",
			mockResp:  nil,
			mockErr:   dictclient.ERR_NOT_FOUND,
			wantEntry: domain.Entry{},
			wantErr:   dictionary.ERR_NOT_FOUND,
			inputWord: "missingword",
		},
		{
			name:      "empty result",
			mockResp:  []dictclient.DictResponse{},
			mockErr:   nil,
			wantEntry: domain.Entry{},
			wantErr:   dictionary.ERR_EMPTY_RESULT,
			inputWord: "empty",
		},
		{
			name:      "client error preserved",
			mockResp:  nil,
			mockErr:   errors.New("network unreachable"),
			wantEntry: domain.Entry{},
			wantErr:   errors.New("network unreachable"),
			inputWord: "any",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mock := &mockClient{
				resp: tc.mockResp,
				err:  tc.mockErr,
			}

			svc := dictionary.NewService(mock)

			got, err := svc.GetWord(ctx, tc.inputWord)

			if tc.wantErr == nil {
				assert.Equal(t, nil, err)
			} else {
				if err == nil {
					t.Fatalf("expected error %v, got nil", tc.wantErr)
				}
				assert.Equal(t, tc.wantErr.Error(), err.Error())
			}

			assert.Equal(t, tc.wantEntry, got)
		})
	}
}
