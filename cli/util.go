package cli

import "os"

func GenerateCommand() {
	file, err := os.Create("stream.bin")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	file.Write([]byte{
		0x1,  // Command: Screen setup
		3,    // Length
		80,   // Screen width
		24,   // Screen height
		0x01, // Color mode: 16 colors
	})

	// Draw some characters
	file.Write([]byte{
		0x2, // Command: Draw character
		4,   // Length
		10,  // x coordinate
		5,   // y coordinate
		1,   // Color index (red)
		'H', // Character 'H'
	})
	file.Write([]byte{
		0x2, // Command: Draw character
		4,   // Length
		11,  // x coordinate
		5,   // y coordinate
		2,   // Color index (green)
		'e', // Character 'e'
	})
	file.Write([]byte{
		0x2, // Command: Draw character
		4,   // Length
		12,  // x coordinate
		5,   // y coordinate
		3,   // Color index (yellow)
		'l', // Character 'l'
	})
	file.Write([]byte{
		0x2, // Command: Draw character
		4,   // Length
		13,  // x coordinate
		5,   // y coordinate
		4,   // Color index (blue)
		'l', // Character 'l'
	})
	file.Write([]byte{
		0x2, // Command: Draw character
		4,   // Length
		14,  // x coordinate
		5,   // y coordinate
		5,   // Color index (magenta)
		'o', // Character 'o'
	})

	// Draw a line
	file.Write([]byte{
		0x3, // Command: Draw line
		6,   // Length
		0,   // x1
		10,  // y1
		20,  // x2
		10,  // y2
		6,   // Color index (cyan)
		'-', // Character '-'
	})

	// Render text
	file.Write([]byte{
		0x4,                                         // Command: Render text
		7,                                           // Length
		30,                                          // x coordinate
		15,                                          // y coordinate
		7,                                           // Color index (white)
		'T', 'e', 's', 't', ' ', 'D', 'e', 'm', 'o', // Text
	})

	// End of file
	file.Write([]byte{0xFF, 0})

}
