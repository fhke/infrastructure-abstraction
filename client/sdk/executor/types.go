package executor

type Executor[I any] interface {
	Exec(input I) error
}
