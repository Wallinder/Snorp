package receiver

type Reciever interface {
	Notify(msg string)
}

func Notify(reciever Reciever, msg string) {
	reciever.Notify(msg)
}
