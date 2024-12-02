
# AVCS

Hello world!
Please check on the 13th, I mistakenly clicked apply : )


+--------------+-------------+-------------+-------------+--- ··· ---+----------------+     
| Command Byte | Length Byte | Data Byte 0 | Data Byte 1 |    ···    |  Data Byte n-1 |     
+--------------+-------------+-------------+-------------+--- ··· ---+----------------+


The data format is an array of bytes, containing sections of above form, in succession. Each section begins with a command byte, specifying the type of operation to be performed on the screen, followed by a length byte, and then a sequence of data bytes, which function as arguments to the command, as specified below:

### 0x1 - Screen setup: Defines the dimensions and colour setting of the screen.
The screen must be set up before any other command is sent. Commands are ignored if the screen hasn't been set up.
**Data format:**
Byte 0: Screen Width (in characters)    
Byte 1: Screen Height (in characters)    
Byte 2: Color Mode (0x00 for monochrome, 0x01 for 16 colors, 0x02 for 256 colors)    

### 0x2 - Draw character: Places a character at a given coordinate of the screen.
**Data format:**

Byte 0: x coordinate    
Byte 1: y coordinate    
Byte 2: Color index    
Byte 3: Character to display (ASCII)

### 0x3 - Draw line: Draws a line from one coordinate of the screen to another.
**Data format:**

Byte 0: x1 (starting coordinate)    
Byte 1: y1 (starting coordinate)    
Byte 2: x2 (ending coordinate) Byte 4: y2 (ending coordinate)
Byte 4: Color index    
Byte 5: Character to use (ASCII)    

### 0x4 - Render text: Renders a string starting from a specific position.
**Data format:**

Byte 0: x coordinate    
Byte 1: y coordinate    
Byte 2: Color index    
Byte 3-n: Text data (ASCII characters)

### 0x5 - Cursor movement: Moves cursor to a specific location without drawing on the screen.
**Data format:**

Byte 0: x coordinate    
Byte 1: y coordinate    

### 0x6 - Draw at cursor: Draws a character at the cursor location.
**Data format:**

Byte 0: Character to draw (ASCII)     
Byte 1: Color index

### 0x7 - Clear screen:
**Data format:**
0xFF - End of file: Marks the end of binary stream.

Data format: No additional data.
