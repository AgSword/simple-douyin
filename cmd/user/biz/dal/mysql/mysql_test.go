package mysql

import (
	"context"
	"log"
	"testing"
)

func TestUser(t *testing.T) {
	Init()
	log.Println(DB == nil)
	GetUserByUsername(context.Background(), "123")
}
