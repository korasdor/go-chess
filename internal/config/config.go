package config

import (
	"bytes"
	"log"
	"os"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
)

type Config struct {
	Auth     *AuthConfig
	Reddis   *RedisConfig
	Gin      *GinConfig
	Http     *HTTPConfig
	Limiter  *LimiterConfig
	Postgres *PostgresConfig
}

type AuthConfig struct {
	JWT          JWTConfig
	PasswordSalt string
}

type JWTConfig struct {
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
	SigningKey      string
}

type GinConfig struct {
	GinMode string
}

type PostgresConfig struct {
	Host         string
	Port         int
	User         string
	Password     string
	DatabaseName string
	SSLMode      string
}

type RedisConfig struct {
	Addr     string
	Password string
}

type HTTPConfig struct {
	Host               string
	Port               string
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	MaxHeaderMegabytes int
}

type LimiterConfig struct {
	RPS   int
	Burst int
	TTL   time.Duration
}

func NewConfig() (*Config, error) {
	consulAddr := os.Getenv("CONSUL_ADDR")
	if consulAddr == "" {
		consulAddr = "localhost:8500"
	}

	client, err := api.NewClient(&api.Config{
		Address: consulAddr,
		Scheme:  "http",
	})

	api.DefaultConfig()
	if err != nil {
		log.Fatal(err)
	}

	kv := client.KV()
	ext, _, err := kv.Get("auth-service-ext", nil)
	if err != nil {
		log.Fatal(err)
	}

	config, _, err := kv.Get("auth-service-value", nil)
	if err != nil {
		log.Fatal(err)
	}

	viper.SetConfigType(string(ext.Value))
	err = viper.ReadConfig(bytes.NewBuffer(config.Value))
	if err != nil {
		return &Config{}, err
	}

	return &Config{
		Auth: &AuthConfig{
			PasswordSalt: viper.GetString("auth.password_salt"),
			JWT: JWTConfig{
				SigningKey:      viper.GetString("auth.jwt_signing_key"),
				AccessTokenTTL:  time.Duration(viper.GetInt("auth.access_token_ttl")) * time.Second,
				RefreshTokenTTL: time.Duration(viper.GetInt("auth.refresh_token_ttl")) * time.Second,
			},
		},
		Gin: &GinConfig{
			GinMode: viper.GetString("gin.mode"),
		},
		Http: &HTTPConfig{
			Host:               "",
			Port:               viper.GetString("http.port"),
			ReadTimeout:        time.Duration(viper.GetInt("http.read_timeout")) * time.Second,
			WriteTimeout:       time.Duration(viper.GetInt("http.write_timeout")) * time.Second,
			MaxHeaderMegabytes: viper.GetInt("http.max_header_megabytes"),
		},
		Limiter: &LimiterConfig{
			RPS:   viper.GetInt("limiter.rps"),
			Burst: viper.GetInt("limiter.burst"),
			TTL:   time.Duration(viper.GetInt("limiter.ttl")) * time.Second,
		},
		Reddis: &RedisConfig{
			Addr:     viper.GetString("redis.addr"),
			Password: viper.GetString("redis.password"),
		},
		Postgres: &PostgresConfig{
			Host:         viper.GetString("postgres.host"),
			Port:         viper.GetInt("postgres.port"),
			User:         viper.GetString("postgres.user"),
			Password:     viper.GetString("postgres.password"),
			DatabaseName: viper.GetString("postgres.database"),
			SSLMode:      viper.GetString("postgres.sslmode"),
		},
	}, nil
}
