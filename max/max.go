package max

import (
	pb "bidirectional-grpc-go/proto"
	"io"
	"log"
	"os"
	"runtime/debug"
	"sync"
)

var l4g = setLog(os.Stdout, "max")

func setLog(out io.Writer, prefix string) *log.Logger {
	l := log.New(out, prefix, log.Lshortfile)
	l.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	return l
}

type Serve struct {
}

type MaxCount struct {
	max int32
	m   sync.RWMutex
}

var max = new(MaxCount)

func (s *Serve) Max(stream pb.Find_MaxServer) error {
	defer func() {
		if r := recover(); r != nil {
			l4g.Printf("panic recovered in panic %s:%s", r, debug.Stack())
		}
	}()

	l4g.Printf("received max request")
	err := serveRequset(stream)
	l4g.Printf("error %v", err)
	return err
}

func serveRequset(stream pb.Find_MaxServer) error {
	for {
		req, err := stream.Recv()
		if err != nil { // return will close stream from server side
			return err
		}
		if i, ok := max.findMax(req.GetInp()); ok {
			err := stream.Send(&pb.Response{Out: i})
			if err != nil { // return will close stream from server side
				return err
			}
		}
	}
}

func (c *MaxCount) findMax(i int32) (int32, bool) {
	c.m.Lock()
	defer c.m.Unlock()
	if i > c.max {
		c.max = i
		return c.max, true
	}
	return c.max, false
}
