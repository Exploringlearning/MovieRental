package app
import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestHelloWorld(t *testing.T) {
    assert.Equal(t, "Hello world", hello() )
}