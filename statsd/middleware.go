package statsd

import (
	"strconv"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type Middleware {
    client Client
}

//Statsd func
func (m *Middleware) Middleware() gin.HandlerFunc {
	// Initialize and configure the client and set options if given.
	cc := newConfiguredClient(client)

	return func(c *gin.Context) {
		log.Debug("Starting Statsd middleware")
		start := time.Now()
		c.Next()

        handler := c.HandlerName()
		cc.IncrThroughput(handler)
		cc.IncrStatusCode(c.Writer.Status(), handler)
		cc.IncrSuccess(c.Errors, handler)
		cc.IncrError(c.Errors, handler)
		cc.Timing(start, handler)
	}
}

func join(strs ...string) string { return strings.Join(strs, ".") }
