package fiber

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFiberServer(t *testing.T) {
	app := GetFiber()
	go func() {
		time.Sleep(1000 * time.Millisecond)
		assert.NoError(t, app.Shutdown())
	}()
	if err := app.Listen(":8000"); err != nil {
		t.Fatal(err)
	}
}
