package goleinu

import (
	"encoding/gob"
	"errors"
	"io"
	"os"
)

type Element interface{}

type Slice struct {
	s []Element
	file *os.File
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
		file: f,
		len: len,
	}, nil
}

func Get[S ~[]E, E any] (s *Slice, i int) (*E, error) {
	if i < 0 || i >= s.len {
		return nil, errors.New("index out of range")
	}

	if i < len(s.s) {
		return s.s[i].(*E), nil
	}

	decoder := gob.NewDecoder(s.file)
	_, err := s.file.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	if i < s.len {
		return s.s[i].(*E), nil
	}

	chunk := make([]*E, 0, s.chunkSize)
	for j := 0; j < s.len; {
		end := min(j+s.chunkSize, s.len)
		if err := decoder.Decode(&chunk); err != nil {
			return nil, err
		}

		if j <= i && i < end {
			return chunk[i-j], nil
		}
		i += s.chunkSize
	}

	return s.s[i].(*E), nil
}

func Append[S ~[]E, E any] (s *Slice, e *E) error {
	if len(s.s) < s.len {
		s.s = append(s.s, e)
		return nil
	}

	s.s = s.s[1:]
	s.s = append(s.s, e)

	encoder := gob.NewEncoder(s.file)
	if err := encoder.Encode(e); err != nil {
		return err
	}

	s.len++
	return nil
}