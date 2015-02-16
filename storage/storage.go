package storage

import (
	"github.com/ysugimoto/signalingo/connection"
)

type Storage interface {
	Add(key string, conn *connection.Connection) bool
	Set(key string, conn *connection.Connection) bool
	Remove(key string) bool
	Get(key string) (*connection.Connection, bool)
	GetAll() []*connection.Connection
	Purge()
}
