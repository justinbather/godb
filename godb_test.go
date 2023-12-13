package godb

import (
	"testing"
)

func TestCreateDB(t *testing.T) {
	_ = CreateDb("Hello")
	List()
}
