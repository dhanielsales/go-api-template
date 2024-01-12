package entity

type ExecFrom[Result any] func(any) Result
