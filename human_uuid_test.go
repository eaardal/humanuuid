package humanuuid_test

import (
	"github.com/eaardal/humanuuid"
	"testing"
)

func TestIsValid_ReturnsTrueForGeneratedIDs(t *testing.T) {
	for i := 0; i < 10000; i++ {
		uid := humanuuid.New()
		if !humanuuid.IsValid(uid) {
			t.Fatalf("expected %s to be a valid ID", uid)
		}
	}
}
