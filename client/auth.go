package client

import "time"

type Auth struct {
	Token  string
	Expiry time.Time
}
