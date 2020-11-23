package main

import (
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

type rateLimiter struct {
	generalLimiter *limiter.Limiter
	loginLimiter   *limiter.Limiter
}

func main() {
	r := &rateLimiter{}

	rate, err := limiter.NewRateFromFormatted("5-S")
	if err != nil {
		panic(err)
	}
	r.generalLimiter = limiter.New(memory.NewStore(), rate)

	loginRate, err := limiter.NewRateFromFormatted("5-M")
	if err != nil {
		panic(err)
	}
	r.loginLimiter = limiter.New(memory.NewStore(), loginRate)

}
