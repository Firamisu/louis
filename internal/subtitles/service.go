package subtitles

import (
	"context"
)

type Subtitle struct {
	Name string
}

type Service interface {
	List(ctx context.Context) ([]Subtitle, error)
}

type svc struct {
}

func NewService() Service {
	return &svc{}
}

func (s *svc) List(ctx context.Context) ([]Subtitle, error) {
	return []Subtitle{}, nil
}
