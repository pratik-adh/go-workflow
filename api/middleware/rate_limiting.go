package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter"
	"github.com/ulule/limiter/drivers/store/memory"
)


var store = memory.NewStore()

func RateLimiter() gin.HandlerFunc {
    // Create a memory store for rate limiting
    store := memory.NewStore()

    rate := limiter.Rate{
        Limit:  4,
        Period: 3 * time.Second,
    }
    
    // Create a rate limiter limiter with the store
    limiter := limiter.New(store, rate)

    return func(c *gin.Context) {
        // Check if the request is allowed by the rate limiter
		context, err := limiter.Get(c, c.FullPath())
		fmt.Println("err in limiter", err)

		// Limit exceeded
		if context.Reached {
            c.JSON(http.StatusTooManyRequests, gin.H{
                "message": "Too many requests, please try again later.",
            })
            c.Abort()
            return
        }
        c.Next()
    }
}
