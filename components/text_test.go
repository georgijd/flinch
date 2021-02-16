package components

import (
	"strings"
	"testing"

	"github.com/liamg/flinch/core"

	"github.com/stretchr/testify/assert"
)

func TestTextSizing(t *testing.T) {

	inputs := []string{
		"",
		"hello world",
		"new lines\nare not allowed",
	}

	for _, input := range inputs {
		t.Run(input, func(t *testing.T) {
			text := NewText(input)
			size := text.MinimumSize()
			output := strings.Split(input, "\n")[0]
			assert.Equal(t, len(output), size.W)
			assert.Equal(t, 1, size.H)
		})
	}
}

func TestRendering(t *testing.T) {

	tests := []struct {
		name         string
		input        string
		output       string
		canvasWidth  int
		canvasHeight int
		justify      core.Alignment
	}{
		{
			name:         "left align",
			input:        "hello",
			output:       "hello",
			canvasWidth:  20,
			canvasHeight: 10,
			justify:      core.AlignLeft,
		},
		{
			name:         "right align",
			input:        "hello",
			output:       "hello",
			canvasWidth:  20,
			canvasHeight: 10,
			justify:      core.AlignRight,
		},
		{
			name:         "center align",
			input:        "hello",
			output:       "hello",
			canvasWidth:  20,
			canvasHeight: 10,
			justify:      core.AlignCenter,
		},
		{
			name:         "zero height",
			input:        "hello",
			output:       "",
			canvasWidth:  10,
			canvasHeight: 0,
		},
		{
			name:         "zero width",
			input:        "hello",
			output:       "",
			canvasWidth:  0,
			canvasHeight: 10,
		},
		{
			name:         "zero width/height",
			input:        "hello",
			output:       "",
			canvasWidth:  0,
			canvasHeight: 0,
		},
		{
			name:         "newline in center",
			input:        "hello\nworld",
			output:       "hello",
			canvasWidth:  10,
			canvasHeight: 10,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			text := NewText(test.input).SetAlignment(test.justify)
			canvas := newTestCanvas(0, 0, core.Size{test.canvasWidth, test.canvasHeight})
			text.Render(canvas)

			if test.canvasWidth == 0 || test.canvasHeight == 0 {
				for x := 0; x < test.canvasWidth; x++ {
					for y := 0; y < test.canvasHeight; y++ {
						assert.Equal(t, rune(0), canvas.Get(x, y))
					}
				}
				return
			}

			switch test.justify {
			case core.AlignLeft:
				for x := 0; x < test.canvasWidth; x++ {
					r := canvas.Get(x, 0)
					if x < len([]rune(test.output)) {
						assert.Equal(t, []rune(test.output)[x], r)
					} else {
						assert.Equal(t, rune(0), r)
					}
				}
			case core.AlignRight:
				offset := test.canvasWidth - len([]rune(test.output))
				for x := 0; x < test.canvasWidth; x++ {
					r := canvas.Get(x, 0)
					if x >= offset {
						assert.Equal(t, []rune(test.output)[x-offset], r)
					} else {
						assert.Equal(t, rune(0), r)
					}
				}
			case core.AlignCenter:
				offset := (test.canvasWidth - len([]rune(test.output))) / 2
				for x := 0; x < test.canvasWidth; x++ {
					r := canvas.Get(x, 0)
					if x >= offset && x < test.canvasWidth-offset-1 {
						assert.Equal(t, []rune(test.output)[x-offset], r)
					} else {
						assert.Equal(t, rune(0), r)
					}
				}
			default:
				t.Errorf("invalid alignment: 0x%x", test.justify)
			}

			for x := 0; x < test.canvasWidth; x++ {
				for y := 1; y < test.canvasHeight; y++ {
					assert.Equal(t, rune(0), canvas.Get(x, y))
				}
			}

		})
	}

}
