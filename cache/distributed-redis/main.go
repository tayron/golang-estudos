package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

type RedisClient struct {
	conn net.Conn
	rd   *bufio.Reader
}

func NewRedisClient(addr string) (*RedisClient, error) {
	conn, err := net.DialTimeout("tcp", addr, 2*time.Second)
	if err != nil {
		return nil, err
	}

	return &RedisClient{
		conn: conn,
		rd:   bufio.NewReader(conn),
	}, nil
}

func (c *RedisClient) Close() error {
	return c.conn.Close()
}

func (c *RedisClient) command(parts ...string) (string, error) {
	var builder strings.Builder
	fmt.Fprintf(&builder, "*%d\r\n", len(parts))
	for _, part := range parts {
		fmt.Fprintf(&builder, "$%d\r\n%s\r\n", len(part), part)
	}

	if _, err := c.conn.Write([]byte(builder.String())); err != nil {
		return "", err
	}

	line, err := c.rd.ReadString('\n')
	if err != nil {
		return "", err
	}

	line = strings.TrimSpace(line)
	if line == "" {
		return "", nil
	}

	switch line[0] {
	case '+':
		return line[1:], nil
	case '-':
		return "", fmt.Errorf("redis error: %s", line[1:])
	case '$':
		if line == "$-1" {
			return "", nil
		}
		payload, err := c.rd.ReadString('\n')
		if err != nil {
			return "", err
		}
		payload = strings.TrimSuffix(payload, "\n")
		payload = strings.TrimSuffix(payload, "\r")
		if _, err := c.rd.ReadString('\n'); err != nil {
			return "", err
		}
		return payload, nil
	default:
		return line, nil
	}
}

func main() {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "127.0.0.1:6379"
	}

	client, err := NewRedisClient(addr)
	if err != nil {
		fmt.Println("nao foi possivel conectar ao Redis")
		fmt.Println("suba o container com: docker-compose up -d")
		fmt.Println("erro:", err)
		return
	}
	defer client.Close()

	pong, err := client.command("PING")
	if err != nil {
		fmt.Println("erro no ping:", err)
		return
	}
	fmt.Println("PING ->", pong)

	if _, err := client.command("SET", "product:1", "notebook-gamer"); err != nil {
		fmt.Println("erro no SET:", err)
		return
	}

	value, err := client.command("GET", "product:1")
	if err != nil {
		fmt.Println("erro no GET:", err)
		return
	}

	fmt.Println("GET product:1 ->", value)
}
