package dictionary

import (
	"context"
	"errors"

	"github.com/firamisu/louis/internal/dictclient"
	"github.com/firamisu/louis/internal/domain"
)

var (
	ERR_EMPTY_RESULT = errors.New("remote api returned empty array")
	ERR_NOT_FOUND    = errors.New("word not found")
)

type Service interface {
	GetWord(ctx context.Context, word string) (domain.Entry, error)
}

type svc struct {
	dictClient dictclient.Client
}

func NewService(
	dictClient dictclient.Client,
) Service {
	return &svc{
		dictClient: dictClient,
	}
}

func (s *svc) GetWord(ctx context.Context, word string) (domain.Entry, error) {
	wEntry, err := s.dictClient.FetchWord(ctx, word)
	if err != nil {
		if errors.Is(err, dictclient.ERR_NOT_FOUND) {
			return domain.Entry{}, ERR_NOT_FOUND
		}
		return domain.Entry{}, err
	}

	if len(wEntry) == 0 {
		return domain.Entry{}, ERR_EMPTY_RESULT
	}

	return domain.FromDictResponse(wEntry[0]), nil
}
