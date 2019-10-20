package main

import (
	pb "bidirectional-grpc-go/proto"
	"context"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"time"

	"google.golang.org/grpc"
)

var l4g = setLog(os.Stdout, "client : ")

func setLog(out io.Writer, prefix string) *log.Logger {
	l := log.New(out, prefix, log.Lshortfile)
	l.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	return l
}

func main() {
	err := startClient("127.0.0.1", "9090")
	l4g.Println(err)
}

func startClient(addr, port string) error {
	target := addr + ":" + port
	conn, client, err := getClientTillSuccess(target, 10)
	if err != nil {
		return err
	}
	defer conn.Close() //not checking error if any
	stream, err := client.Max(context.Background())
	if err != nil {
		return err
	}
	return startStream(stream) //infinite loop
}

func startStream(stream pb.Find_MaxClient) error {
	done := make(chan struct{}, 3)
	go sendReq(stream, done)
	go recvMsg(stream, done)
	go doneMsg(stream, done)
	<-done
	return nil
}

const limit = 10

func sendReq(stream pb.Find_MaxClient, done chan struct{}) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := 0; i < limit; i++ {
		n := r.Int31n(limit)
		req := pb.Request{Inp: n}
		l4g.Printf("sending value %d", n)
		if err := stream.Send(&req); err != nil {
			l4g.Printf("error sending req to server %v", err)
			done <- struct{}{}
			return
		}
		time.Sleep(2 * time.Second)
	}
	if err := stream.CloseSend(); err != nil { //close stream when done
		log.Println(err)
	}
	done <- struct{}{}
	return
}

func recvMsg(stream pb.Find_MaxClient, done chan struct{}) {
	for {
		resp, err := stream.Recv()
		if err != nil {
			l4g.Printf("error receiving response from server %v", err)
			done <- struct{}{}
			return
		}
		l4g.Printf("received max val %v", resp.GetOut())
	}
}

func doneMsg(stream pb.Find_MaxClient, done chan struct{}) {
	ctx := stream.Context()
	<-ctx.Done()
	if err := ctx.Err(); err != nil {
		log.Println(err)
	}
	done <- struct{}{}
}

const RETRY_WAIT_TIME_MILLISEC = 1000
const TIME_OUT_SEC = 20

func getClientTillSuccess(target string, count int) (*grpc.ClientConn,
	pb.FindClient, error) {
	var finalErr error
	for i := 0; i < count; i++ {
		l4g.Printf("%d dialing target : %s", i, target)
		conn, client, err := getClientConnAndStream(target)
		if err != nil {
			l4g.Printf("%d opening stream error %v;retrying...", i, err)
			time.Sleep(RETRY_WAIT_TIME_MILLISEC * time.Millisecond)
			finalErr = err
			continue
		} else {
			return conn, client, nil
		}
	}

	return nil, nil, finalErr
}

func getClientConnAndStream(target string) (*grpc.ClientConn,
	pb.FindClient, error) {
	// dial server
	var conn *grpc.ClientConn
	var err error
	conn, err = grpc.Dial(target,
		grpc.WithBlock(),
		grpc.WithTimeout(TIME_OUT_SEC*time.Second),
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(math.MaxInt32)),
	)
	if err != nil {
		return nil, nil, err
	}
	client := pb.NewFindClient(conn)
	return conn, client, nil
}
