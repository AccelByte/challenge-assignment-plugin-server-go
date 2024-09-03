// Copyright (c) 2024 AccelByte Inc. All Rights Reserved.
// This is licensed software from AccelByte Inc, for limitations
// and restrictions contact your company contract manager.

package server

import (
	pb "challenge-assignment-plugin-server-go/pkg/pb"
	"context"
	"math/rand"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AssignmentServiceServer struct {
	pb.UnimplementedAssignmentFunctionServer
}

func (server *AssignmentServiceServer) Assign(ctx context.Context, request *pb.AssignmentRequest) (*pb.AssignmentResponse, error) {
	response := pb.AssignmentResponse{}
	response.UserId = request.UserId
	response.AssignedGoals = make([]*pb.Goal, 0, 1)

	if len(request.Goals) == 0 {
		return nil, status.Error(codes.InvalidArgument, "no active goals is available to be assigned")
	}

	randomInt := rand.Intn(len(request.Goals))
	goal := request.Goals[randomInt]

	response.AssignedGoals = append(response.AssignedGoals, goal)
	response.Namespace = request.Namespace

	return &response, nil
}

func NewAssignmentServiceServer() *AssignmentServiceServer {
	return &AssignmentServiceServer{}
}
