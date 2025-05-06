package main

import (
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/weather-app/internal/env"
	"github.com/weather-app/internal/repository"
	"github.com/weather-app/internal/store"
	"github.com/weather-app/service"
)

var rdb *redis.Client

func main() {

	cfg := config{
		addr:   env.GetString("ADDR", ":8080"),
		apiURL: env.GetString("WEATHER_API_URL", ""),
		apiKey: env.GetString("WEATHER_API_KEY", ""),
		redisCfg: redisConfig{
			addr: env.GetString("REDIS_ADDR", "localhost:6379"),
			pw:   env.GetString("REDIS_PW", ""),
			db:   env.GetInt("REDIS_DB", 0),
		},
		contextTimeout: env.GetInt("TIME_DURATION", 10),
	}

	rdb = store.NewRedisCache(cfg.redisCfg.addr, cfg.redisCfg.pw, cfg.redisCfg.db)
	// logger.Info("redis cache connection established")
	fmt.Println("redis cache connection established")

	defer rdb.Close()

	expiry := 10 * time.Minute
	timeout := 10 * time.Second
	redisStore := repository.NewWeatherRepo(rdb, expiry)
	weatheService := service.NewWeatherService(redisStore, timeout)


	app := &application{
		config: cfg,
		weatherService: weatheService,
	}

	mux := app.mount()
    if err := app.run(mux); err != nil {
		fmt.Println("err connecting ")
    }
	log.Println(app.run(mux))

}
