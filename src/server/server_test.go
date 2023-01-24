package server

import (
	"context"
	"testing"

	protos "github.com/datastx/IceBridge/src/protos/pbs"
	"github.com/stretchr/testify/require"
)

func TestHealthCheck(t *testing.T) {
	s := &healthServer{}

	req := &protos.HealthCheckRequest{}
	res, err := s.Check(context.Background(), req)
	require.NoError(t, err)
	require.Equal(t, protos.HealthCheckResponse_SERVING, res.Status)
}

// func TestHealthCheckError(t *testing.T) {
// 	s := &healthServer{}

// 	req := &protos.HealthCheckRequest{}
// 	_, err := s.Check(context.Background(), req)
// 	require.Error(t, err)
// 	st, ok := status.FromError(err)
// 	require.True(t, ok)
// 	require.Equal(t, codes.Internal, st.Code())
// 	require.Equal(t, "Internal Error", st.Message())
// }
