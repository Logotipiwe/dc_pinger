package service

type Interruptor struct {
	isInterrupted bool
}

func NewInterruptor() *Interruptor {
	return &Interruptor{}
}

func (i *Interruptor) Resume() {
	i.isInterrupted = false
}

func (i *Interruptor) Interrupt() {
	i.isInterrupted = true
}

func (i *Interruptor) Interrupted() bool {
	return i.isInterrupted
}
