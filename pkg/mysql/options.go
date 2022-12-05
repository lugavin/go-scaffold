package mysql

import "time"

// Option -.
type Option func(*Mysql)

// MaxOpenConns -.
func MaxOpenConns(n int) Option {
	return func(c *Mysql) {
		c.maxOpenConns = n
	}
}

// MaxIdleConns -.
func MaxIdleConns(n int) Option {
	return func(c *Mysql) {
		c.maxIdleConns = n
	}
}

// ConnMaxIdleTime -.
func ConnMaxIdleTime(t time.Duration) Option {
	return func(c *Mysql) {
		c.connMaxIdleTime = t
	}
}

// ConnMaxLifetime -.
func ConnMaxLifetime(t time.Duration) Option {
	return func(c *Mysql) {
		c.connMaxLifetime = t
	}
}

// ConnAttempts -.
func ConnAttempts(n int) Option {
	return func(c *Mysql) {
		c.connAttempts = n
	}
}

// ConnTimeout -.
func ConnTimeout(t time.Duration) Option {
	return func(c *Mysql) {
		c.connTimeout = t
	}
}
