package test

import "net/http/httptest"

type closeNotifierRecorder struct {
	*httptest.ResponseRecorder
	closed chan bool
}

func NewCloseNotifierRecorder() *closeNotifierRecorder {
	return &closeNotifierRecorder{
		httptest.NewRecorder(),
		make(chan bool, 1),
	}
}

func (c *closeNotifierRecorder) close() {
	c.closed <- true
}

func (c *closeNotifierRecorder) CloseNotify() <-chan bool {
	return c.closed
}
