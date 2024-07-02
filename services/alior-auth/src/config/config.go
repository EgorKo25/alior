package config

type Endpoint struct {
	// ListenAddr для сервера
	ListenAddr string
	Port       string
	// Address для клиентского enpoint'a
	Address string
}

type Database struct {
	Address      string
	Port         string
	MaxConn      int32
	User         string
	UserPassword string
	DatabaseName string
}

type Config struct {
	Endpoints *Endpoint `endpoint:"endpoint-service"`
	Database  *Database `database:"database"`
}
