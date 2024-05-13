package testing

import (
	"testing"

	"github.com/robbymaul/golang-postgresql.git/connection"
	"github.com/stretchr/testify/assert"
)

func TestConnection(t *testing.T) {
	conn, err := connection.GetConnection()
	if err != nil {
		t.Fatal(err.Error())
	}

	assert.NotNil(t, conn)
}
