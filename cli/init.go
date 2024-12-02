package cli

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

func Execute() {
	scn, err := tcell.NewScreen()
	if err != nil {
		panic(err)
	}

	scn.Init()
	scnWidth, scnHeight := scn.Size()

	fmt.Println(scnWidth, scnWidth)

	scn.SetStyle(tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorRed))

	greet := "Hello, World!"

	for i, r := range greet {
		scn.SetContent(scnWidth/2+i, scnHeight/2, r, nil, tcell.StyleDefault)
	}

	scn.Show()
}
