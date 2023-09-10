package storage

import (
	"errors"
	"github.com/gocql/gocql"
	"github.com/mjmoshiri/log-lyfe/gol/internal/models"
	"github.com/mjmoshiri/log-lyfe/gol/internal/utils"
	"github.com/scylladb/gocqlx/v2"
	"github.com/scylladb/gocqlx/v2/qb"
)

// database is the internal representation of our database with the actual session
type database struct {
	// TODO: define a custom Session interface for mocking
	session gocqlx.Session
	table   string
}

// DB is an interface that represents the methods the server needs
type DB interface {
	Insert(event *models.Event) error
	Find(filters map[string]string, fetchSize uint) ([]models.Event, error)
	Close()
}

// New creates a new database instance using the provided configuration and returns it as a DB interface
func New(cfg models.DBConfig) (DB, error) {
	cluster := gocql.NewCluster(cfg.Cluster...)
	cluster.Keyspace = cfg.Keyspace
	cluster.Consistency = cfg.Consistency
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

// Insert inserts an event into the database
func (db *database) Insert(event *models.Event) error {
	event.Bucket = utils.TimeToBucket(event.Timestamp)
	return db.session.Query(EventTable.Insert()).BindStruct(event).ExecRelease()
}

// Find finds events in the database based on the provided filters
// in production, this method would be paginated, but for the sake of simplicity, we'll just use a fetch size
// the fetch size is capped at 10000
// This method is not optimized for performance as
// Allowing Filtering is not recommended in production (Cassandra is Query Driven, not Schema Driven)
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

// Close closes the database session
func (db *database) Close() {
	db.session.Close()
}
