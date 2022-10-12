package httpv2

import (
	"fmt"
	"time"
)

type HandleFunc func(c *Context)
type FilterBuilder func(next Filter) Filter
type Filter func(c *Context)

var _ FilterBuilder = MetricFilterBuilder

func MetricFilterBuilder(next Filter) Filter {
	return func(c *Context) {
		start := time.Now().Nanosecond()
		next(c)
		end := time.Now().Nanosecond()
		fmt.Println("用了时间：", end-start)
	}
}
