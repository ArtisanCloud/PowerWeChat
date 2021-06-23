package decorators

type FinallyResult struct {
	Content interface{}
}

func NewFinallyResult(content interface{}) *FinallyResult {
	return &FinallyResult{
		content,
	}
}
