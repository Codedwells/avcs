package cli

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

type Renderer struct {
	screen  tcell.Screen
	colors  []tcell.Style
	cursorX int
	cursorY int
}

func Execute() {
	scn, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}

	scn.Init()

	renderer := &Renderer{}
	renderer.screen = scn

	var reader io.Reader = os.Stdin

	if len(os.Args) > 1 {
		file, err := os.Open(os.Args[1])

		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		reader = file
	}

	err = renderer.ProcessBinaryStream(reader)
	if err != nil && err != io.EOF {
		log.Fatalf("Error processing binary stream: %v", err)
	}

}

func (r *Renderer) ProcessBinaryStream(reader io.Reader) error {
	for {
		// Read  byte
		cmd := make([]byte, 1)
		_, err := reader.Read(cmd)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		// get length byte
		lenByte := make([]byte, 1)
		_, err = reader.Read(lenByte)
		if err != nil {
			return err
		}

		// Read data
		dataLen := int(lenByte[0])
		data := make([]byte, dataLen)
		if dataLen > 0 {
			_, err = io.ReadFull(reader, data)
			if err != nil {
				return err
			}
		}

		err = r.ProcessCommand(cmd[0], data)
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
	}
}

func (r *Renderer) InitializeColors(colorMode byte) {
	switch colorMode {
	case 0x00: // For the monochrome
		r.colors = []tcell.Style{tcell.StyleDefault, tcell.StyleDefault.Background(tcell.ColorBlack)}
	case 0x01: // 16 colors
		r.colors = []tcell.Style{
			tcell.StyleDefault,
			tcell.StyleDefault.Foreground(tcell.ColorBlack),
			tcell.StyleDefault.Foreground(tcell.ColorBlue),
			tcell.StyleDefault.Foreground(tcell.ColorGreen),
			tcell.StyleDefault.Foreground(tcell.ColorDarkCyan),
			tcell.StyleDefault.Foreground(tcell.ColorRed),
			tcell.StyleDefault.Foreground(tcell.ColorDarkMagenta),
			tcell.StyleDefault.Foreground(tcell.ColorYellow),
			tcell.StyleDefault.Foreground(tcell.ColorWhite),
			tcell.StyleDefault.Foreground(tcell.ColorBlack).Bold(true),
			tcell.StyleDefault.Foreground(tcell.ColorBlue).Bold(true),
			tcell.StyleDefault.Foreground(tcell.ColorGreen).Bold(true),
			tcell.StyleDefault.Foreground(tcell.ColorLightCyan).Bold(true),
			tcell.StyleDefault.Foreground(tcell.ColorRed).Bold(true),
			tcell.StyleDefault.Foreground(tcell.ColorDarkMagenta).Bold(true),
			tcell.StyleDefault.Foreground(tcell.ColorYellow).Bold(true),
			tcell.StyleDefault.Foreground(tcell.ColorWhite).Bold(true)}

	case 0x02: // 256 colors
		r.colors = make([]tcell.Style, 256)
		for i := 0; i < 256; i++ {
			r.colors[i] = tcell.StyleDefault.Foreground(tcell.Color(i))
		}
	default:
		r.colors = []tcell.Style{tcell.StyleDefault}
	}
}

func (r *Renderer) ProcessCommand(cmd byte, data []byte) error {
	switch cmd {
	case 0x1: // Do the term screen setup
		if len(data) < 3 {
			return fmt.Errorf("insufficient data for screen setup")
		}

		r.screen.SetSize(int(data[0]), int(data[1]))
		r.InitializeColors(data[2])

	case 0x2: // Draw a character
		if len(data) < 4 {
			return fmt.Errorf("insufficient data for drawing character")
		}
		x, y := int(data[0]), int(data[1])
		colorIdx := int(data[2])
		char := rune(data[3])

		if colorIdx >= len(r.colors) {
			colorIdx = 0 // Default to first color if out of range
		}

		r.screen.SetContent(x, y, char, nil, r.colors[colorIdx])

	case 0x3: // Draw a line
		if len(data) < 6 {
			return fmt.Errorf("insufficient data for drawing line")
		}
		x1, y1 := int(data[0]), int(data[1])
		x2, y2 := int(data[2]), int(data[3])
		colorIdx := int(data[4])
		char := rune(data[5])

		if colorIdx >= len(r.colors) {
			colorIdx = 0 // use default color if out of range
		}

		if x1 == x2 { // Vertical line
			for y := min(y1, y2); y <= max(y1, y2); y++ {
				r.screen.SetContent(x1, y, char, nil, r.colors[colorIdx])
			}
		} else if y1 == y2 { // Horizontal line
			for x := min(x1, x2); x <= max(x1, x2); x++ {
				r.screen.SetContent(x, y1, char, nil, r.colors[colorIdx])
			}
		}

	case 0x4: // Render text to screen
		if len(data) < 3 {
			return fmt.Errorf("insufficient data for rendering text")
		}
		x, y := int(data[0]), int(data[1])
		colorIdx := int(data[2])
		text := string(data[3:])

		if colorIdx >= len(r.colors) {
			colorIdx = 0 // use default color if out of range
		}

		for i, char := range text {
			r.screen.SetContent(x+i, y, char, nil, r.colors[colorIdx])
		}

	case 0x5: // Cursor movement
		if len(data) < 2 {
			return fmt.Errorf("insufficient data for cursor movement")
		}

		r.cursorX = int(data[0])
		r.cursorY = int(data[1])

	case 0x6: // Draw at cursor
		if len(data) < 2 {
			return fmt.Errorf("insufficient data for drawing at cursor")
		}
		char := rune(data[0])
		colorIdx := int(data[1])

		if colorIdx >= len(r.colors) {
			colorIdx = 0 // use default color if out of range
		}
		r.screen.SetContent(r.cursorX, r.cursorY, char, nil, r.colors[colorIdx])

	case 0x7: // Clear screen
		r.screen.Clear()

	case 0xFF: // End of file
		return io.EOF

	default:
		return fmt.Errorf("unknown command: 0x%x", cmd)
	}

	r.screen.Show()
	return nil
}
