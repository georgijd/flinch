package window

import (
	"testing"

	"github.com/gdamore/tcell/v2"
	"github.com/stretchr/testify/assert"
)

func getTermSize(t *testing.T) (int, int) {
	screen := tcell.NewSimulationScreen("")
	if err := screen.Init(); err != nil {
		t.Fatal(err)
	}
	defer screen.Fini()
	return screen.Size()
}

func TestWindowSize(t *testing.T) {

	w, h := getTermSize(t)

	win, err := New(WindowOptionSimulation())
	if err != nil {
		t.Error(err)
	}

	aw, ah := win.Size()

	assert.Equal(t, w, aw)
	assert.Equal(t, h, ah)
}
