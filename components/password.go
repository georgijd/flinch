package components

import (
	"github.com/gdamore/tcell/v2"
	"github.com/liamg/flinch/core"
)

type passwordInput struct {
	core.Sizer
	content     string
	cursor      int
	style       core.Style
	selected    bool
	keyHandlers []func(key *tcell.EventKey) bool
}

func NewPasswordInput() *passwordInput {
	inp := &passwordInput{
		style: core.StyleInput,
	}
	inp.SetSizeStrategy(core.SizeStrategyMaximumWidth())
	return inp
}

func (n *passwordInput) SetStyle(s core.Style) {
	n.style = s
}

func (n *passwordInput) GetInput() string {
	return n.content
}

func (n *passwordInput) Render(canvas core.Canvas) {

	canvas.Fill(' ', n.style)

	canvas.Set(n.cursor, 0, ' ', n.style.ToggleInvert())

	size := canvas.Size()

	visibleContent := n.content
	clampedCursor := n.cursor

	if n.cursor >= size.W {
		visibleContent = string([]rune(n.content)[n.cursor-size.W:])
		clampedCursor = size.W - 1
	}

	for offset := range []rune(visibleContent) {
		st := n.style
		if offset == clampedCursor {
			st = st.ToggleInvert()
		}
		canvas.Set(offset, 0, '*', st)
	}
}
func (n *passwordInput) MinimumSize() core.Size {
	return core.Size{
		W: 1,
		H: 1,
	}
}

func (n *passwordInput) Deselect() {
	n.selected = false
}

func (n *passwordInput) Select() bool {
	if n.selected {
		return false
	}
	n.selected = true
	return true
}

func (n *passwordInput) OnKeypress(handler func(key *tcell.EventKey) bool) {
	n.keyHandlers = append(n.keyHandlers, handler)
}

func (n *passwordInput) HandleKeypress(key *tcell.EventKey) {

	for _, handler := range n.keyHandlers {
		if handler(key) {
			return
		}
	}

	switch key.Key() {
	case tcell.KeyHome:
		n.cursor = 0
	case tcell.KeyEnd:
		n.cursor = len(n.content)
	case tcell.KeyLeft:
		if n.cursor > 0 {
			n.cursor--
		}
	case tcell.KeyRight:
		if n.cursor < len(n.content) {
			n.cursor++
		}
	case tcell.KeyBackspace, tcell.KeyBackspace2:
		if n.cursor > 0 {
			first := []rune(n.content)[:n.cursor-1]
			second := []rune(n.content)[n.cursor:]
			n.content = string(first) + string(second)
			n.cursor--
		}
	case tcell.KeyDelete:
		if n.cursor < len([]rune(n.content)) {
			first := []rune(n.content)[:n.cursor]
			second := []rune(n.content)[n.cursor+1:]
			n.content = string(first) + string(second)
		}
	case tcell.KeyRune:
		first := []rune(n.content)[:n.cursor]
		second := []rune(n.content)[n.cursor:]
		n.content = string(first) + string(key.Rune()) + string(second)
		n.cursor++
	}
}
