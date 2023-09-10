package models

import (
	"github.com/gocql/gocql"
)

// DBConfig represents the configuration for connecting to the Cassandra cluster
type DBConfig struct {
	Cluster     []string          `yaml:"cluster"`
	Keyspace    string            `yaml:"keyspace"`
	Username    string            `yaml:"username"`
	Password    string            `yaml:"password"`
	Table       string            `yaml:"table"`
	Consistency gocql.Consistency `yaml:"consistency"`
}