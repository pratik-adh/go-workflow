package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter"
	"github.com/ulule/limiter/drivers/store/memory"
)


var store = memory.NewStore()


func getUserIdentifier(c *gin.Context) string {
    // Replace this with your logic to identify users (e.g., authentication token, IP address, etc.)
    // For the sake of example, we'll use the client's IP address as the identifier.
    return c.ClientIP()
}


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
		context, _ := limiter.Get(c, c.FullPath() +", "+ c.ClientIP())

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
