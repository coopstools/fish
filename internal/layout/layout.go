package layout

import (
	"fmt"
	"os"
)

const (
	RESET_SCREEN = "\033[2J"
	HIGHLIGHT    = "\033[7m"
	RESET        = "\033[0m"
	MOVETO       = "\033[%d;%dH"
)

type mode int

const (
	STD_MODE mode = iota
	SINGLE_SCAN_MODE
	DOUBLE_SCAN_MODE
	ESC_SCAN_MODE
)

type Layout struct {
	program []string

	mode    mode
	subMode mode

	xpos int
	ypos int

	lastXpos int
	lastYpos int

	//lsb_0 - down/right,
	//lsb_1 - reverse (ie 0 - down, 1 - right, 2 - up, 3 - left)
	//lsb_2 - magnitude (0 - one step, 1, two steps)
	//negative means terminate
	direction int

	stack stack
}

func (l *Layout) InitPrint() {

	fmt.Printf("%s%s", RESET_SCREEN, l.homePos())
	fmt.Println()
	for _, line := range l.program {
		fmt.Println(" " + line)
	}
	fmt.Print(l.endPos())
}

func (l *Layout) Print() {
	lastChar := l.program[l.lastYpos][l.lastXpos]
	fmt.Printf("%s%s%c", l.lastPos(), RESET, lastChar)

	curChar := l.program[l.ypos][l.xpos]
	fmt.Printf("%s%s%c", l.curPos(), HIGHLIGHT, curChar)
	fmt.Print(l.endPos())
}

func (l *Layout) curPos() string {
	return fmt.Sprintf(MOVETO, l.ypos+2, l.xpos+2)
}

func (l *Layout) lastPos() string {
	return fmt.Sprintf(MOVETO, l.lastYpos+2, l.lastXpos+2)
}

func (l *Layout) homePos() string {
	return fmt.Sprintf(MOVETO, 1, 1)
}

func (l *Layout) endPos() string {
	return fmt.Sprintf(MOVETO, len(l.program)+3, 2)
}

func (l *Layout) Update() {
	switch l.mode {
	case STD_MODE:
		l.runCommand()
	case SINGLE_SCAN_MODE, DOUBLE_SCAN_MODE, ESC_SCAN_MODE:
		l.scan()
	}
	l.lastXpos = l.xpos
	l.lastYpos = l.ypos
	l.mov()
}

func (l *Layout) runCommand() {
	l.direction &= 0x3 // bring magitude to 1
	c := l.program[l.ypos][l.xpos]
	switch c {
	case ' ':
		break
	case ';':
		l.direction = -1
	case 'v':
		l.direction = 0
	case '>':
		l.direction = 1
	case '^':
		l.direction = 2
	case '<':
		l.direction = 3
	case '#':
		l.direction ^= 0x2
	case '\\':
		l.direction ^= 0x1
	case '/':
		l.direction ^= 0x3
	case '|':
		l.direction ^= (l.direction >> 1)
	case '_':
		l.direction ^= ^(l.direction >> 1)
	case '!':
		l.direction |= 0x4
	case '?':
		if l.stack.pop() == 0 {
			l.direction |= 0x4
		}
	case '.':
		l.lastXpos = l.xpos
		l.lastYpos = l.ypos
		l.ypos, l.xpos = l.stack.pop2()
		return
	case '\'':
		l.mode = SINGLE_SCAN_MODE
	case '"':
		l.mode = DOUBLE_SCAN_MODE
	case '+':
		l.stack.sum()
	case '-':
		l.stack.diff()
	case '*':
		l.stack.prod()
	case ',':
		l.stack.div()
	case '%':
		l.stack.mod()
	case ')':
		l.stack.gt()
	case '(':
		l.stack.lt()
	case ':':
		l.stack.dup()
	case '~':
		l.stack.rm()
	case '$':
		l.stack.swap()
	default:
		if c >= '0' && c <= '9' {
			l.stack.push(int(c - '0'))
			break
		}
		if c >= 'a' && c <= 'f' {
			l.stack.push(int(c-'a') + 10)
			break
		}
		if c >= 'A' && c <= 'F' {
			l.stack.push(int(c-'A') + 10)
			break
		}
		l.stack.push(int(c))
	}
}

func (l *Layout) scan() {
	c := l.program[l.ypos][l.xpos]
	if l.subMode == ESC_SCAN_MODE {
		l.stack.push(int(c))
		l.subMode = STD_MODE
		return
	}
	if c == '\'' && l.mode == SINGLE_SCAN_MODE {
		l.mode = STD_MODE
		return
	}
	if c == '"' && l.mode == DOUBLE_SCAN_MODE {
		l.mode = STD_MODE
		return
	}
	if c == '\\' {
		l.subMode = ESC_SCAN_MODE
		return
	}
	l.stack.push(int(c))
}

func (l *Layout) mov() {
	if l.direction == -1 {
		fmt.Println("Program exited properly")
		os.Exit(0)
	}
	dir := l.direction & 0x3
	mag := ((l.direction & 0x4) >> 2) + 1
	switch dir {
	case 0:
		l.ypos = (l.ypos + mag) % len(l.program)
	case 1:
		l.xpos = (l.xpos + mag) % len(l.program[0])
	case 2:
		l.ypos = (l.ypos - mag + len(l.program)) % len(l.program)
	case 3:
		l.xpos = (l.xpos - mag + len(l.program[0])) % len(l.program[0])
	}
}

func New(programLines []string) *Layout {
	var longestLine int
	for _, line := range programLines {
		if len(line) > longestLine {
			longestLine = len(line)
		}
	}

	for i, line := range programLines {
		programLines[i] = fmt.Sprintf("%-*s", longestLine, line)
	}

	return &Layout{
		program:   programLines,
		mode:      STD_MODE,
		direction: -1,
	}
}
