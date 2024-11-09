package goleinu

import "errors"

type config struct {
	chunkSize int
	maxInMemorySize int
	bufferSize int
}

type Option interface {
	apply(*config) error
}

type chunkSizeOption int

func (c chunkSizeOption) apply(cfg *config) error {
	if c <= 0 {
		return errors.New("chunk size must be greater than 0")
	}

	cfg.chunkSize = int(c)
	return nil
}

func WithChunkSize(size int) Option {
	return chunkSizeOption(size)
}

type maxInMemorySizeOption int

func (m maxInMemorySizeOption) apply(cfg *config) error {
	if m <= 0 {
		return errors.New("max in memory size must be greater than 0")
	}

	cfg.maxInMemorySize = int(m)
	return nil
}

func WithMaxInMemorySize(size int) Option {
	return maxInMemorySizeOption(size)
}

type bufferSizeOption int

func (b bufferSizeOption) apply(cfg *config) error {
	if b <= 0 {
		return errors.New("buffer size must be greater than 0")
	}

	cfg.bufferSize = int(b)
	return nil
}

func WithBufferSize(size int) Option {
	return bufferSizeOption(size)
}

