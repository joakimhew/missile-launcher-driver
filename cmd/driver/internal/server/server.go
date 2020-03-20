package server

import (
	"context"

	pb "github.com/joakimhew/missile-launcher-driver/internal/api/driver"
	ct "github.com/joakimhew/missile-launcher-driver/internal/controltransfer"
)

type Server struct {
	Controller *ct.Controller

	pb.UnimplementedDriverServer
}

func (s *Server) Left(ctx context.Context, req *pb.CommandRequest) (*pb.CommandReply, error) {
	return s.sendCommand(ctx, ct.Left)
}

func (s *Server) Up(ctx context.Context, req *pb.CommandRequest) (*pb.CommandReply, error) {
	return s.sendCommand(ctx, ct.Up)
}

func (s *Server) Right(ctx context.Context, req *pb.CommandRequest) (*pb.CommandReply, error) {
	return s.sendCommand(ctx, ct.Right)
}

func (s *Server) Down(ctx context.Context, req *pb.CommandRequest) (*pb.CommandReply, error) {
	return s.sendCommand(ctx, ct.Down)
}

func (s *Server) UpLeft(ctx context.Context, req *pb.CommandRequest) (*pb.CommandReply, error) {
	return s.sendCommand(ctx, ct.UpLeft)
}

func (s *Server) DownLeft(ctx context.Context, req *pb.CommandRequest) (*pb.CommandReply, error) {
	return s.sendCommand(ctx, ct.DownLeft)
}

func (s *Server) UpRight(ctx context.Context, req *pb.CommandRequest) (*pb.CommandReply, error) {
	return s.sendCommand(ctx, ct.UpRight)
}

func (s *Server) DownRight(ctx context.Context, req *pb.CommandRequest) (*pb.CommandReply, error) {
	return s.sendCommand(ctx, ct.DownRight)
}

func (s *Server) Stop(ctx context.Context, req *pb.CommandRequest) (*pb.CommandReply, error) {
	return s.sendCommand(ctx, ct.Stop)
}

func (s *Server) Fire(ctx context.Context, req *pb.CommandRequest) (*pb.CommandReply, error) {
	return s.sendCommand(ctx, ct.Fire)
}

func (s *Server) sendCommand(ctx context.Context, command ct.Command) (*pb.CommandReply, error) {
	d, err := s.Controller.Send(ctx, command)
	if err != nil {
		return nil, err
	}

	return &pb.CommandReply{
		Message: string(d),
	}, nil
}
