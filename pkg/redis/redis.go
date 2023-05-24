package redis

import (
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	_defaultPoolSize        = 16
	_defaultMinIdleConns    = 20
	_defaultMaxRetries      = 5
	_defaultMinRetryBackoff = 300 * time.Millisecond
	_defaultMaxRetryBackoff = 500 * time.Millisecond
	_defaultReadTimeout     = 5 * time.Second
	_defaultWriteTimeout    = 3 * time.Second
	_defaultDialTimeout     = 5 * time.Second
	_defaultIdleTimeout     = 12 * time.Second
	_defaultPoolTimeout     = 6 * time.Second
)

// Redis -.
type Redis struct {
	password        string
	db              int
	poolSize        int
	minIdleConns    int
	maxRetries      int
	minRetryBackoff time.Duration
	maxRetryBackoff time.Duration
	readTimeout     time.Duration
	writeTimeout    time.Duration
	dialTimeout     time.Duration
	idleTimeout     time.Duration
	poolTimeout     time.Duration
}

func New(addrs []string, opts ...Option) redis.UniversalClient {
	cfg := &Redis{
		poolSize:        _defaultPoolSize,
		minIdleConns:    _defaultMinIdleConns,
		maxRetries:      _defaultMaxRetries,
		minRetryBackoff: _defaultMinRetryBackoff,
		maxRetryBackoff: _defaultMaxRetryBackoff,
		readTimeout:     _defaultReadTimeout,
		writeTimeout:    _defaultWriteTimeout,
		dialTimeout:     _defaultDialTimeout,
		idleTimeout:     _defaultIdleTimeout,
		poolTimeout:     _defaultPoolTimeout,
	}

	// Custom options
	for _, opt := range opts {
		opt(cfg)
	}

	client := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:           addrs,
		DB:              cfg.db,
		Password:        cfg.password,
		PoolSize:        cfg.poolSize,
		MaxRetries:      cfg.maxRetries,
		MinRetryBackoff: cfg.minRetryBackoff,
		MaxRetryBackoff: cfg.maxRetryBackoff,
		DialTimeout:     cfg.dialTimeout,
		ReadTimeout:     cfg.readTimeout,
		WriteTimeout:    cfg.writeTimeout,
		MinIdleConns:    cfg.minIdleConns,
		PoolTimeout:     cfg.poolTimeout,
		IdleTimeout:     cfg.idleTimeout,
	})

	return client
}
