package main

type pool chan []byte

func newPool(cap int) pool {
	return make(chan []byte, cap)
}

func (p pool) get() []byte {
	var v []byte

	select {
	case v = <-p:
	default:
		v = make([]byte, 10)
	}

	return v
}

func (p pool) put(b []byte) {
	select {
	case p <- b:
	default:
	}
}
