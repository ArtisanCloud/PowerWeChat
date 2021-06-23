package decorators

type TerminateResult struct {
	Content interface{}
}

func NewTerminateResult(content interface{}) *TerminateResult {
	return &TerminateResult{
		content,
	}
}
