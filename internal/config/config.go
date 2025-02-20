package config

type Config struct {
    Server struct {
        Port string
    }
    Database struct {
        Driver string
        DSN    string
    }
    Redis struct {
        Addr     string
        Password string
        DB       int
    }
}

func LoadConfig() (*Config, error) {
    // TODO: Implement config loading
    return &Config{}, nil
}
