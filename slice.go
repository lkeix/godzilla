package goleinu

import (
	"os"
)

type Element interface{}

type Slice struct {
	s []Element
	files []*os.File
	chunkSize int
	len int
}

func New[S ~[]E, E Element](len int, cap int, opts ...Option) (*Slice, error){
	cfg := new(config)
	for _, opt := range opts {
		if err := opt.apply(cfg); err != nil {
			return nil, err
		}
	}
	
	s := make([]Element, len, cap)

	f, err := os.CreateTemp("", "tmp")
	if err != nil {
		return nil, err
	}

	return &Slice{
		s: s,
		files: []*os.File{f},
		len: len,
	}, nil
}

func Get[S ~[]E, E any] (s *Slice, i int) (*E, error) {
	return s.s[i].(*E), nil
}

func Append[S ~[]E, E any] (s *Slice, e *E) error {
	if len(s.s) < s.len {
		s.s = append(s.s, e)
		return nil
	}

	s.s = s.s[1:]
	s.s = append(s.s, e)

	s.len++
	return nil
}