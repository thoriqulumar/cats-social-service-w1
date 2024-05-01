package config

import (
	"context"
	"fmt"
	"strings"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

// Config represents the application configuration
type Config struct {
	DB         DBConfig `env:",prefix=DB_,required"`
	JWTSecret  string   `env:"JWT_SECRET"`
	BcryptSalt int      `env:"BCRYPT_SALT"`
}

type DBConfig struct {
	Name     string `env:"NAME"`
	Port     string `env:"PORT"`
	Host     string `env:"HOST"`
	Username string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
	Params   string `env:"PARAMS"`
}

func Load(ctx context.Context) (*Config, error) {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := envconfig.Process(ctx, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// ConnectionURL returns the connection URL for the database
func (c DBConfig) ConnectionURL() string {
	params := strings.ReplaceAll(c.Params, `"`, "")
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s %s",
		c.Username, c.Password, c.Name, c.Host, c.Port, params)
}
