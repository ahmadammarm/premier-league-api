package config

import "time"


type Config struct {
    Port string
    ScrapperURL string
    CacheExpirationTime time.Duration
}

func Configuration() *Config {
    port := "8080"
    scrapperURL := "https://onefootball.com/en/competition/premier-league-9/table"
    cacheExpirationMinutes, _ := time.ParseDuration("5m")

    return &Config{
        Port: port,
        ScrapperURL: scrapperURL,
        CacheExpirationTime: cacheExpirationMinutes,
    }
}