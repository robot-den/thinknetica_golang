package chat

import "sync"

// Chat представляет собой чат-сервис с набором методов
type Chat struct {
	members     map[int]chan string
	mu          *sync.Mutex
	idGenerator int
}

// New создает новый объект чат-сервиса
func New() *Chat {
	c := Chat{
		members:     map[int]chan string{},
		mu:          &sync.Mutex{},
		idGenerator: 0,
	}
	return &c
}

// Subscribe позволяет участнику подписаться на новые сообщения
func (c *Chat) Subscribe(ch chan string) int {
	c.mu.Lock()
	defer c.mu.Unlock()

	memberId := c.idGenerator
	c.members[memberId] = ch
	c.idGenerator++

	return memberId
}

// Unsubscribe позволяет участнику отписаться от получения сообщений
func (c *Chat) Unsubscribe(id int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.members, id)
}

// Broadcast рассылает новое сообщение все участникам чата
func (c *Chat) Broadcast(message string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, member := range c.members {
		member <- message
	}
}
