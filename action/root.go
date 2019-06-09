package action

import (
	"fmt"
	"net/http"
)

type Root interface {
	render()
}

type Home struct {
	Writer   http.ResponseWriter
	Receiver *http.Request
}

func (action Home) render() {
	fmt.Fprintf(action.Writer, "Hi there, I love home")
}
