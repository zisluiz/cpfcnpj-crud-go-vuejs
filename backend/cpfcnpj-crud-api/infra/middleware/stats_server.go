package middleware

import (
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
)

type Stats struct {
	Uptime       time.Time      `json:"uptime"`
	RequestCount uint64         `json:"requestCount"`
	Statuses     map[string]int `json:"responseStatusCount"`
	Mutex        sync.RWMutex   `json:"-"`
}

func NewStats() *Stats {
	return &Stats{
		Uptime:   time.Now(),
		Statuses: map[string]int{},
	}
}

func (s *Stats) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		s.Mutex.Lock()
		defer s.Mutex.Unlock()
		s.RequestCount++
		status := strconv.Itoa(c.Response().Status)
		s.Statuses[status]++
		return nil
	}
}
