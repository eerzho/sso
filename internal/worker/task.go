package worker

type Task interface {
	Execute() error
}
