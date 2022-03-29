package answerremover

type Accumulator interface {
	addOutputLine(s string)
	addEnclosedLine(s string)
	getOutputs() []string
	getEnclosedLines() []string
}

type accummulator struct {
	outputs       []string
	enclosedLines []string
}

func (a *accummulator) addOutputLine(s string) {
	a.enclosedLines = make([]string, 0)
	a.outputs = append(a.outputs, s)
}

func (a *accummulator) addEnclosedLine(s string) {
	a.enclosedLines = append(a.enclosedLines, s)
}

func (a *accummulator) getOutputs() []string {
	return a.outputs
}

func (a *accummulator) getEnclosedLines() []string {
	return a.enclosedLines
}

func NewAccumulator() Accumulator {
	return &accummulator{}
}
