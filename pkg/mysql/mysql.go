package mysql

import (
	"fmt"
	"log"
	"time"

	"github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

const (
	_defaultMaxOpenConns    = 20
	_defaultMaxIdleConns    = 20
	_defaultConnAttempts    = 10
	_defaultConnTimeout     = time.Second
	_defaultConnMaxIdleTime = 5 * time.Minute
	_defaultConnMaxLifetime = 5 * time.Minute
)

// Mysql -.
type Mysql struct {
	maxOpenConns    int
	maxIdleConns    int
	connAttempts    int
	connTimeout     time.Duration
	connMaxIdleTime time.Duration
	connMaxLifetime time.Duration

	Builder squirrel.StatementBuilderType
	Pool    *sqlx.DB // DB is a connection pool
}

// New -.
func New(url string, opts ...Option) (*Mysql, error) {
	cfg := &Mysql{
		maxOpenConns:    _defaultMaxOpenConns,
		maxIdleConns:    _defaultMaxIdleConns,
		connAttempts:    _defaultConnAttempts,
		connTimeout:     _defaultConnTimeout,
		connMaxIdleTime: _defaultConnMaxIdleTime,
		connMaxLifetime: _defaultConnMaxLifetime,
	}

	// Custom options
	for _, opt := range opts {
		opt(cfg)
	}

	var (
		err  error
		pool *sqlx.DB
	)
	for cfg.connAttempts > 0 {
		// When we first executed sql.Open("mysql", ds), the DB returned is actually a pool of underlying DB connections.
		// The sql package takes care of maintaining the pool, creating and freeing connections automatically.
		// This DB is also safe to be concurrently accessed by multiple Goroutines.
		pool, err = sqlx.Open("mysql", url)
		if err == nil {
			if err = pool.Ping(); err == nil {
				break
			}
		}

		log.Printf("Mysql is trying to connect, attempts left: %d", cfg.connAttempts)

		time.Sleep(cfg.connTimeout)

		cfg.connAttempts--
	}

	if err != nil {
		return nil, fmt.Errorf("mysql - NewMysql - connAttempts == 0: %w", err)
	}

	cfg.Pool = pool
	cfg.Pool.SetMaxOpenConns(cfg.maxOpenConns)
	cfg.Pool.SetMaxIdleConns(cfg.maxIdleConns)
	cfg.Pool.SetConnMaxIdleTime(cfg.connMaxIdleTime)
	cfg.Pool.SetConnMaxLifetime(cfg.connMaxLifetime)
	cfg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question)

	return cfg, nil
}

// Close -.
func (p *Mysql) Close() {
	if p.Pool != nil {
		_ = p.Pool.Close()
	}
}
