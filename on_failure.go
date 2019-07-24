package httpapi

type OnFailureFuncMiddleware func(next OnFailureFunc) OnFailureFunc
type OnFailureFuncMiddlewares []OnFailureFuncMiddleware

func (ms OnFailureFuncMiddlewares) OnFailureFn(main OnFailureFunc) OnFailureFunc {
	total := len(ms)
	if total == 0 {
		return main
	}

	head := main
	for i := total - 1; i >= 0; i-- {
		head = ms[i](head)
	}

	return head
}
