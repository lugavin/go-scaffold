package redis

import "time"

// Option -.
type Option func(*Redis)

// PoolSize -.
func PoolSize(n int) Option {
	return func(r *Redis) {
		r.poolSize = n
	}
}

// MinIdleConns -.
func MinIdleConns(n int) Option {
	return func(r *Redis) {
		r.minIdleConns = n
	}
}

// MaxRetries -.
func MaxRetries(n int) Option {
	return func(r *Redis) {
		r.maxRetries = n
	}
}

// MinRetryBackoff -.
func MinRetryBackoff(t time.Duration) Option {
	return func(r *Redis) {
		r.minRetryBackoff = t
	}
}

// MaxRetryBackoff -.
func MaxRetryBackoff(t time.Duration) Option {
	return func(r *Redis) {
		r.maxRetryBackoff = t
	}
}

// ReadTimeout -.
func ReadTimeout(t time.Duration) Option {
	return func(r *Redis) {
		r.readTimeout = t
	}
}

// WriteTimeout -.
func WriteTimeout(t time.Duration) Option {
	return func(r *Redis) {
		r.writeTimeout = t
	}
}

// DialTimeout -.
func DialTimeout(t time.Duration) Option {
	return func(r *Redis) {
		r.dialTimeout = t
	}
}

// IdleTimeout -.
func IdleTimeout(t time.Duration) Option {
	return func(r *Redis) {
		r.idleTimeout = t
	}
}

// PoolTimeout -.
func PoolTimeout(t time.Duration) Option {
	return func(r *Redis) {
		r.poolTimeout = t
	}
}
