package answerremover

import "fmt"

type AnnotationProcessor interface {
	Process(trailingSpace string, codeBlock []string, commentVariation *commentvariation) (*string, error)
}
type DefaultProcessor struct {
}
type PanicAnnotationProcessor struct {
}
type ErrorAnnotationProcessor struct {
}
type NopAnnotationProcessor struct {
}
type CodeAnnotationProcessor struct {
	annotation string
}

func (p *DefaultProcessor) Process(trailingSpace string, codeBlock []string, commentVariation *commentvariation) (*string, error) {
	res := trailingSpace
	res += generateTodoComment(commentVariation, "answer here")
	return &res, nil
}

func (p *PanicAnnotationProcessor) Process(trailingSpace string, codeBlock []string, commentVariation *commentvariation) (*string, error) {
	res := trailingSpace
	res += `panic("Not yet implemented")`
	res += fmt.Sprintf(" %s", generateTodoComment(commentVariation, "answer here"))
	return &res, nil
}

func (p *ErrorAnnotationProcessor) Process(trailingSpace string, codeBlock []string, commentVariation *commentvariation) (*string, error) {
	res := trailingSpace
	res += `throw new Error("Method not implemented.");`
	res += fmt.Sprintf(" %s", generateTodoComment(commentVariation, "answer here"))
	return &res, nil
}

func (p *NopAnnotationProcessor) Process(trailingSpace string, codeBlock []string, commentVariation *commentvariation) (*string, error) {
	return nil, nil
}

func (p *CodeAnnotationProcessor) Process(trailingSpace string, codeBlock []string, commentVariation *commentvariation) (*string, error) {
	res := trailingSpace
	res += p.annotation
	res += fmt.Sprintf(" %s", generateTodoComment(commentVariation, "replace this"))
	return &res, nil
}

func generateTodoComment(commentvariation *commentvariation, todoComment string) string {
	switch *commentvariation {
	case doubleslash:
		return fmt.Sprintf(`// TODO: %s`, todoComment)
	case dashdash:
		return fmt.Sprintf(`-- TODO: %s`, todoComment)
	case hash:
		return fmt.Sprintf(`# TODO: %s`, todoComment)
	case slashasterisk:
		return fmt.Sprintf(`/* TODO: %s */`, todoComment)
	case curlyslashasterisk:
		return fmt.Sprintf(`{/* TODO: %s */}`, todoComment)
	case htmlcomment:
		return fmt.Sprintf(`<!-- TODO: %s -->`, todoComment)
	}
	panic("unreachable")
}

func NewProcessor(annotation *string) AnnotationProcessor {
	if annotation == nil {
		return &DefaultProcessor{}
	}
	switch *annotation {
	case "panic":
		return &PanicAnnotationProcessor{}
	case "error":
		return &ErrorAnnotationProcessor{}
	case "nop":
		return &NopAnnotationProcessor{}
	default:
		return &CodeAnnotationProcessor{*annotation}
	}
}
