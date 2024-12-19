package main

import (
	"bufio"
	"os"
	"sync"
	"time"
)

type Record struct {
	ID        string
	Data      map[string]interface{}
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Database struct {
	mu      sync.RWMutex
	records map[string]Record
	indexes map[string]map[interface{}]string
}

// function for new Database Instance

func NewDatabase() *Database {
	return &Database{
		records: make(map[string]Record),
		indexes: make(map[string]map[interface{}]string),
	}
}

type Terminal struct {
	db     *Database
	reader *bufio.Reader
}

func NewTerminal(db *Database) *Terminal {
	return &Terminal{
		db:     db,
		reader: bufio.NewReader(os.Stdin),
	}
}

func main() {

}
