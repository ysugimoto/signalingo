package signaling

import (
	"encoding/json"
	"fmt"
	"github.com/ysugimoto/signalingo/connection"
	"github.com/ysugimoto/signalingo/env"
	"github.com/ysugimoto/signalingo/operation"
	"github.com/ysugimoto/signalingo/storage"
	"log"
	"strings"
)

type Manager struct {
	storage storage.Storage
}

func NewManager(env env.Env) *Manager {
	var st storage.Storage

	switch env.Storage.Type {
	case "memory":
		st = storage.NewMapStorage()
	case "redis":
		st = storage.NewRedisStorage(env.Redis.Host, env.Redis.Port)
	}
	return &Manager{
		storage: st,
	}
}

func (m *Manager) GetAllUsers() (users []operation.Users) {
	connections := m.storage.GetAll()
	for _, conn := range connections {
		if conn.Closed {
			continue
		}

		fmt.Printf("%v\n", conn)
		users = append(users, operation.Users{
			UUID:   conn.UUID,
			Locked: conn.Locked,
			Extra:  conn.Extra,
		})
	}

	return
}

func (m *Manager) AddConnection(conn *connection.Connection) (ok bool) {
	if ok = m.storage.Add(conn.UUID, conn); ok {
		m.SendConnect(conn.UUID)
	}

	return
}

func (m *Manager) CloseConnection(conn *connection.Connection) (ok bool) {
	if ok = m.storage.Remove(conn.UUID); ok {
		conn.Close()
		m.SendDisconnect(conn.UUID)
	}

	return ok
}

func (m *Manager) Send(message []byte, to string) (err error) {
	if conn, ok := m.storage.Get(to); ok {
		return conn.Send(message)
	}

	return nil
}

func (m *Manager) Broadcast(message []byte) (err error) {
	connections := m.storage.GetAll()
	for _, conn := range connections {
		return conn.Send(message)
	}

	return nil
}

func (m *Manager) BroadcastIgnore(message []byte, id string) (err error) {
	connections := m.storage.GetAll()
	for _, conn := range connections {
		if conn.UUID == id {
			continue
		}
		conn.Send(message)
	}

	return nil
}

func (m *Manager) HandleMessage(msg string) {
	var op operation.Operation
	decoder := json.NewDecoder(strings.NewReader(msg))
	if err := decoder.Decode(&op); err != nil {
		log.Printf("%v\n", err)
	}

	switch op.Type {
	case operation.LOCK:
		if conn, ok := m.storage.Get(op.Sender); ok {
			if conn.Locked {
				m.SendError(op.Sender, operation.LOCK_ERROR, "You have already locked")
			} else {
				conn.Locked = true
				m.storage.Set(op.Sender, conn)
				m.SendLock(op.Sender)
			}
		}
	case operation.UNLOCK:
		if conn, ok := m.storage.Get(op.Sender); ok {
			if !conn.Locked {
				m.SendError(op.Sender, operation.UNLOCK_ERROR, "You have not locked yet")
			} else {
				conn.Locked = false
				m.storage.Set(op.Sender, conn)
				m.SendUnlock(op.Sender)
			}
		}
	case operation.OFFER:
		if conn, ok := m.storage.Get(op.Target); !ok {
			m.SendError(op.Sender, operation.USER_NOTFOUND, "User not found")
		} else {
			if conn.Locked {
				m.SendError(op.Sender, operation.OFFER_ERROR, "Target user is locked")
			} else {
				m.SendOffer(op.Sender, op.Target, op.Sdp, op.Extra)
			}
		}
	case operation.ANSWER:
		if conn, ok := m.storage.Get(op.Target); !ok {
			m.SendError(op.Sender, operation.USER_NOTFOUND, "User not found")
		} else {
			if conn.Locked {
				m.SendError(op.Sender, operation.ANSWER_ERROR, "Target user is locked")
			} else {
				m.SendAnswer(op.Sender, op.Target, op.Sdp, op.Extra)
			}
		}
	case operation.CANDIDATE:
		if _, ok := m.storage.Get(op.Target); !ok {
			m.SendError(op.Sender, operation.USER_NOTFOUND, "User not found")
		} else {
			m.SendCandidate(op.Sender, op.Target, op.Candidate)
		}
	default:
		fmt.Println("Unknown operation")
	}
}

func (m *Manager) SendError(to, signature, message string) {
	if msg, err := operation.NewErrorMessage(signature, message); err != nil {
		log.Printf("%v\n", err)
	} else if err = m.Send(msg, to); err != nil {
		log.Printf("%v\n", err)
	}
}

func (m *Manager) SendConnect(userId string) {
	if msg, err := operation.NewConnectMessage(userId); err != nil {
		log.Printf("%v\n", err)
	} else if err = m.BroadcastIgnore(msg, userId); err != nil {
		log.Printf("%v\n", err)
	}
}

func (m *Manager) SendDisconnect(userId string) {
	if msg, err := operation.NewDisconnectMessage(userId); err != nil {
		log.Printf("%v\n", err)
	} else if err = m.BroadcastIgnore(msg, userId); err != nil {
		log.Printf("%v\n", err)
	}
}

func (m *Manager) SendLock(userId string) {
	if msg, err := operation.NewLockMessage(userId); err != nil {
		log.Printf("%v\n", err)
	} else if err = m.BroadcastIgnore(msg, userId); err != nil {
		log.Printf("%v\n", err)
	}
}
func (m *Manager) SendUnlock(userId string) {
	if msg, err := operation.NewUnlockMessage(userId); err != nil {
		log.Printf("%v\n", err)
	} else if err = m.BroadcastIgnore(msg, userId); err != nil {
		log.Printf("%v\n", err)
	}
}
func (m *Manager) SendOffer(from, to, sdp string, extra map[string]string) {
	if msg, err := operation.NewOfferMessage(from, sdp, extra); err != nil {
		log.Printf("%v\n", err)
	} else if err = m.Send(msg, to); err != nil {
		log.Printf("%v\n", err)
	}
}
func (m *Manager) SendAnswer(from, to, sdp string, extra map[string]string) {
	if msg, err := operation.NewAnswerMessage(from, sdp, extra); err != nil {
		log.Printf("%v\n", err)
	} else if err = m.Send(msg, to); err != nil {
		log.Printf("%v\n", err)
	}
}
func (m *Manager) SendCandidate(from, to, candidate string) {
	if msg, err := operation.NewCandidateMessage(from, candidate); err != nil {
		log.Printf("%v\n", err)
	} else if err = m.BroadcastIgnore(msg, to); err != nil {
		log.Printf("%v\n", err)
	}
}
