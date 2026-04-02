package receiver

type Reciever interface {
	Notify() error
}

func Notify(reciever Reciever) {
	reciever.Notify()
}
