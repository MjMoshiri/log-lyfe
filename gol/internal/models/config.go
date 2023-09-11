package models

// DBConfig represents the configuration for connecting to the Cassandra cluster
type DBConfig struct {
	Cluster     []string `yaml:"cluster"`
	Keyspace    string   `yaml:"keyspace"`
	Username    string   `yaml:"username"`
	Password    string   `yaml:"password"`
	Table       string   `yaml:"table"`
	Consistency uint16   `yaml:"consistency"`
}

// ServerConfig represents the configuration for the server
type ServerConfig struct {
	Port      string `yaml:"port"`
	SecretKey string `yaml:"secretKey"`
	QueryKey  string `yaml:"queryKey"`
}

// SystemInfo represents the system information returned by the info handler
type SystemInfo struct {
	Hostname  string `json:"hostname"`
	OS        string `json:"os"`
	Arch      string `json:"arch"`
	CPUs      int    `json:"cpus"`
	GoVersion string `json:"go_version"`
}
