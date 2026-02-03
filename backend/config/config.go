package config

type Config struct {
	ConfigPath  string  `help:"Path to configuration file (json|yaml|toml)" default:"" name:"config" env:"CONFIG"`
	SteamAPIKey string  `help:"Steam API Key" env:"STEAM_API_KEY"`
	LogLevel    string  `help:"Logging level" default:"info" enum:"debug,info,warning,error" env:"LOG_LEVEL"`
	API         API     `embed:"" prefix:"api-"`
	Metrics     Metrics `embed:"" prefix:"metrics-"`
	DB          DB      `embed:"" prefix:"db-"`
}

type API struct {
	ListenAddress string `help:"API server listen address" default:":8889" env:"API_LISTEN_ADDRESS"`
	PublicAddress string `help:"Public API address for OpenAPI page" default:"localhost:8889" env:"API_PUBLIC_ADDRESS"`
	CorsOrigins   string `help:"CORS allowed origins" default:"steaminputdb.com" env:"API_CORS_ORIGINS"`
}

type Metrics struct {
	ListenAddress string `help:"Metrics server listen address" default:":8899" env:"METRICS_LISTEN_ADDRESS"`
}

type DB struct {
	// TODO
}
