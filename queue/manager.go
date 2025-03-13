package queue

type manager[T any] interface {
	read() T
	moveRead()
	insert(T)
	moveInsert()
	flush() []T
	integrate(queue *Queue[T])
}
