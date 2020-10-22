package mysql

// Config represents the necessary data to create a mysql connection
type Config struct {
	user     string
	password string
	host     string
	port     string
	database string
}

// NewConfig creates a new Config struct
func NewConfig(user, password, host, port, database string) Config {
	return Config{
		user:     user,
		password: password,
		host:     host,
		port:     port,
		database: database,
	}
}
