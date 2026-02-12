package dictclient

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
)

type DictResponse struct {
	Word      string     `json:"word"`
	Phonetic  string     `json:"phonetic,omitempty"`
	Phonetics []Phonetic `json:"phonetics"`
	Origin    string     `json:"origin,omitempty"`
	Meanings  []Meaning  `json:"meanings"`
}

type Phonetic struct {
	Text  string `json:"text"`
	Audio string `json:"audio,omitempty"`
}

type Meaning struct {
	PartOfSpeech string       `json:"partOfSpeech"`
	Definitions  []Definition `json:"definitions"`
}

type Definition struct {
	Definition string   `json:"definition"`
	Example    string   `json:"example,omitempty"`
	Synonyms   []string `json:"synonyms"`
	Antonyms   []string `json:"antonyms"`
}

type Client interface {
	FetchWord(ctx context.Context, word string) ([]DictResponse, error)
}

type DictClient struct {
	httpC       http.Client
	endpointUrl func(word string) string
}

func NewDictClient() Client {
	return &DictClient{
		httpC: http.Client{},
		endpointUrl: func(word string) string {
			return fmt.Sprintf("https://api.dictionaryapi.dev/api/v2/entries/en/%s", word)
		},
	}
}

var (
	ERR_NOT_OK_RES       = errors.New("not ok repsonse from remote api")
	ERR_NOT_FOUND        = errors.New("not found")
	ERR_CANNOT_PARSE_RES = errors.New("cannot parse response from remote api")
)

func (c *DictClient) FetchWord(ctx context.Context, word string) ([]DictResponse, error) {
	var wordEntry []DictResponse

	resp, err := c.httpC.Get(c.endpointUrl(word))
	if err != nil {
		slog.Error(err.Error())
		return wordEntry, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusNotFound:
		return wordEntry, ERR_NOT_FOUND
	default:
		return wordEntry, ERR_NOT_OK_RES
	}

	if err = json.NewDecoder(resp.Body).Decode(&wordEntry); err != nil {
		slog.Error(err.Error())
		return wordEntry, ERR_CANNOT_PARSE_RES
	}

	return wordEntry, nil

}
