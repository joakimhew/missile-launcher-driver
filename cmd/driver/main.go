package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"strconv"

	ct "github.com/joakimhew/missile-launcher-driver/internal/controltransfer"

	"github.com/google/gousb"
	pb "github.com/joakimhew/missile-launcher-driver/internal/api/driver"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type server struct {
	Controller ct.Controller
	pb.UnimplementedDriverServer
}


func (s *server) Left(ctx context.Context, req *pb.CommandRequest) (*pb.CommandReply, error) {
	return sendCommand(ctx, ct.Left)
}

func (s *server) Up(ctx context.Context, req *pb.CommandRequest) (*pb.CommandReply, error) {
	return sendCommand(ctx, ct.Up)
}

func (s *server) Right(ctx context.Context, req *pb.CommandRequest) (*pb.CommandReply, error) {
	return sendCommand(ctx, ct.Right)
}

func (s *server) Down(ctx context.Context, req *pb.CommandRequest) (*pb.CommandReply, error) {
	return sendCommand(ctx, ct.Down)
}

func (s *server) UpLeft(ctx context.Context, req *pb.CommandRequest) (*pb.CommandReply, error) {
	return sendCommand(ctx, ct.UpLeft)
}

func (s *server) DownLeft(ctx context.Context, req *pb.CommandRequest) (*pb.CommandReply, error) {
	return sendCommand(ctx, ct.DownLeft)
}

func (s *server) UpRight(ctx context.Context, req *pb.CommandRequest) (*pb.CommandReply, error) {
	return sendCommand(ctx, ct.UpRight)
}

func (s *server) DownRight(ctx context.Context, req *pb.CommandRequest) (*pb.CommandReply, error) {
	return sendCommand(ctx, ct.DownRight)
}

func (s *server) Stop(ctx context.Context, req *pb.CommandRequest) (*pb.CommandReply, error) {
	return sendCommand(ctx, ct.Stop)
}

func (s *server) Fire(ctx context.Context, req *pb.CommandRequest) (*pb.CommandReply, error) {
	return sendCommand(ctx, ct.Fire)
}

func sendCommand(ctx context.Context, command ct.Command) (*pb.CommandReply, error) {
	d, err := ct.Send(ctx, command)
	if err != nil {
		return nil, err
	}

	return &pb.CommandReply{
		Message: string(d),
	}, nil
}

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	var logLevel string
	logLevel = os.Getenv("DRIVER_LOG")

	if logLevel != "" {
		l, err := logrus.ParseLevel(logLevel)
		if err != nil {
			logrus.WithField(
				"DRIVER_LOG",
				logLevel).Errorf("DRIVER_LOG=%v not supported\n", logLevel)
		} else {
			logrus.SetLevel(l)
		}
	}

	ctx := gousb.NewContext()
	defer ctx.Close()

	var libusbDebugLevel int
	level := os.Getenv("LIBUSB_DEBUG_LEVEL")
	libusbDebugLevel, err := strconv.Atoi(level)
	if err != nil {
		libusbDebugLevel = 0
	}

	ctx.Debug(libusbDebugLevel)

	const vendorID = 0x2123
	const productID = 0x1010

	dev, err = ctx.OpenDeviceWithVIDPID(vendorID, productID)
	if dev == nil {
		logrus.WithFields(logrus.Fields{
			"vendorID":  fmt.Sprintf("0x%x", vendorID),
			"productID": fmt.Sprintf("0x%x", productID),
		}).Panic("Could not find device")
	} else if err != nil {
		logrus.Panicf("Could not open a device: %v", err)
	}
	defer dev.Close()

	err = dev.SetAutoDetach(true)
	if err != nil {
		logrus.Panicln(err)
	}

	err = dev.Reset()
	if err != nil {
		logrus.Panicln(err)
	}

	_, done, err := dev.DefaultInterface()
	if err != nil {
		logrus.Panicln(err)
	}

	defer done()

	port := os.Getenv("PORT")

	if port == "" {
		logrus.Panicln("Missing environment variable: PORT")
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		logrus.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	logrus.Debug("Registering gRPC server")
	pb.RegisterDriverServer(s, &server{
		Controller: &ct.Controller{
			Device: dev
		} 
	})

	logrus.WithField("port", port).Debug("Serving gRPC")
	if err := s.Serve(lis); err != nil {
		logrus.Fatalf("failed to serve: %v", err)
	}
}
