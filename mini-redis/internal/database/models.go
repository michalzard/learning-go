// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package database

import (
	"time"
)

type Kv struct {
	K         string
	V         string
	CreatedAt time.Time
}
