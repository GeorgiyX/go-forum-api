package main

import (
	"fmt"
	"github.com/caarlos0/env/v6"
)

type Urls struct {
	User string
	Root string
}

type Settings struct {
	Release    bool   `env:"RELEASE_FLAG" envDefault:"false"`
	APIPort    uint   `env:"PG_DEFAULT_USER" envDefault:"8080"`
	DBHost     string `env:"PG_HOST" envDefault:"localhost"`
	DBPort     uint   `env:"PG_PORT" envDefault:"5432"`
	DBName     string `env:"PG_DB_NAME" envDefault:"forum"`
	DBUser     string `env:"PG_FORUM_USER" envDefault:"forum_user"`
	DBPassword string `env:"PG_PASSWORD" envDefault:"forum_user_password"`
	Urls       Urls
}

func GetUrls() Urls {
	return Urls{
		User: "/user",
		Root: "/api",
	}
}

func LoadSettings() *Settings {
	settings := Settings{}
	settings.Urls = GetUrls()
	if err := env.Parse(settings); err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", settings)
	return &settings
}
