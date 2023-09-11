// Package storage provides abstractions and utilities for interacting with data storage systems, such as databases.
package storage

import (
	"errors"
	"github.com/gocql/gocql"
	"github.com/mjmoshiri/log-lyfe/gol/internal/models"
	"github.com/mjmoshiri/log-lyfe/gol/internal/pkg/eventer"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
)

// database represents the internal structure of our database, including the active session.
type database struct {
	// TODO: Introduce a custom Session interface for easier mocking.
	session gocqlx.Session
	table   string
}

// DB defines the methods required for server-database interactions.
type DB interface {
	Insert(event *models.Event) error
	Find(filters map[string]string, fetchSize uint) ([]models.Event, error)
	Close()
}

// New initializes a new database instance based on the provided configuration and returns a DB interface.
func New(cfg models.DBConfig) (DB, error) {
	cluster := gocql.NewCluster(cfg.Cluster...)
	cluster.Keyspace = cfg.Keyspace
	cluster.Consistency = gocql.Consistency(cfg.Consistency)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: cfg.Username,
		Password: cfg.Password,
	}
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		return nil, err
	}
	return &database{
		session: session,
		table:   cfg.Table,
	}, nil
}

// Insert adds an event to the database.
func (db *database) Insert(event *models.Event) error {
	event.Bucket = eventer.TimeToBucket(event.Timestamp)
	return db.session.Query(EventTable.Insert()).BindStruct(event).ExecRelease()
}

// Find retrieves events from the database based on provided filters.
// For simplicity, this method uses a fetch size instead of pagination, with a maximum cap of 10,000.
// Note: This method is not optimized for performance. Using 'AllowFiltering' is not recommended in production
// as Cassandra is query-driven, not schema-driven.
func (db *database) Find(filters map[string]string, fetchSize uint) ([]models.Event, error) {
	if fetchSize == 0 {
		fetchSize = 1000
	}
	if fetchSize > 10000 {
		return nil, errors.New("fetch size cannot be greater than 10000")
	}
	q := qb.Select(EventTable.Name()).Limit(fetchSize).AllowFiltering()
	for k, v := range filters {
		q = q.Where(qb.EqLit(k, v))
	}
	n, s := q.ToCql()
	qx := db.session.Query(n, s)
	events := make([]models.Event, 0, fetchSize)
	err := qx.SelectRelease(&events)
	if err != nil {
		return nil, err
	}
	return events, nil
}

// Close terminates the database session.
func (db *database) Close() {
	db.session.Close()
}
