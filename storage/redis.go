package storage

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/ysugimoto/signalingo/connection"
	"golang.org/x/net/websocket"
	"log"
)

type RedisStorage struct {
	Storage
	conn redis.Conn
	set  map[string]*websocket.Conn
	key  string
}

func NewRedisStorage(host string, port int) *RedisStorage {
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		panic(err)
	}

	conn.Do("FLUSHALL")

	return &RedisStorage{
		conn: conn,
		set:  make(map[string]*websocket.Conn),
	}
}

func (r *RedisStorage) Add(key string, conn *connection.Connection) bool {
	if _, ok := r.set[key]; ok {
		log.Printf("Redis add storage error: connection duplicate %s", key)
		return false
	}

	r.set[key] = conn.Conn
	arg := redis.Args{}.Add(key).AddFlat(conn)
	if _, err := r.conn.Do("HMSET", arg...); err != nil {
		log.Printf("Redis storage error: %v\n", err)
		return false
	}

	return true
}

func (r *RedisStorage) Set(key string, conn *connection.Connection) bool {
	if _, ok := r.set[key]; !ok {
		log.Printf("Redis set storage error: connection undefined %s", key)
		return false
	}
	r.set[key] = conn.Conn
	arg := redis.Args{}.Add(key).AddFlat(conn)
	if _, err := r.conn.Do("HMSET", arg...); err != nil {
		log.Printf("Redis storage error: %v\n", err)
		return false
	}

	return true
}

func (r *RedisStorage) Remove(key string) bool {
	if _, err := r.conn.Do("DEL", key); err != nil {
		log.Printf("Redis storage error: %v\n", err)
		return false
	}

	delete(r.set, key)
	return true
}

func (r *RedisStorage) Get(key string) (*connection.Connection, bool) {
	value, err := redis.Values(r.conn.Do("HGETALL", key))
	if err != nil {
		log.Printf("Redis storage error: %v\n", err)
		return nil, false
	}

	var conn connection.Connection
	if err := redis.ScanStruct(value, &conn); err != nil {
		log.Printf("Redis storage error: %v\n", err)
		return nil, false
	}

	if c, ok := r.set[key]; !ok {
		log.Println("Redis storage error: connection map is lost")
		return nil, false
	} else {
		conn.Conn = c
	}

	return &conn, true
}

func (r *RedisStorage) GetAll() (connections []*connection.Connection) {
	for uuid, _ := range r.set {
		if conn, ok := r.Get(uuid); ok {
			connections = append(connections, conn)
		}
	}

	return
}
