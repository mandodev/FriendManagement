package config

//Server : sturct hold model for server confguration
type Server struct {
	Mode            string
	Addr            string
	ShutdownTimeout int
}
