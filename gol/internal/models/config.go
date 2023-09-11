// Package models defines the primary data structures used throughout the application.
package models

// DBConfig defines the parameters for connecting to a Cassandra cluster.
type DBConfig struct {
	Cluster     []string `yaml:"cluster"`
	Keyspace    string   `yaml:"keyspace"`
	Username    string   `yaml:"username"`
	Password    string   `yaml:"password"`
	Table       string   `yaml:"table"`
	Consistency uint16   `yaml:"consistency"`
}

// ServerConfig holds the server's configuration details.
type ServerConfig struct {
	Port      string `yaml:"port"`
	SecretKey string `yaml:"secretKey"`
	QueryKey  string `yaml:"queryKey"`
}

// SystemInfo encapsulates system details returned by the info endpoint.
type SystemInfo struct {
	Hostname  string `json:"hostname"`
	OS        string `json:"os"`
	Arch      string `json:"arch"`
	CPUs      int    `json:"cpus"`
	GoVersion string `json:"go_version"`
}
