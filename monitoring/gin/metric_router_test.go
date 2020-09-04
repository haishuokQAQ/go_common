package gin

import (
	"testing"
)

func TestStartHealthServer(t *testing.T) {
	StartHealthServer(25256)
	ch := make(chan int, 1)
	<-ch
}
