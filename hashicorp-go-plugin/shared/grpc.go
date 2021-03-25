package shared

import (
	"github.com/jimil749/hashicorp-go-plugin/proto"
	"golang.org/x/net/context"
)

// GRPCClient is an implementation of KV that talks over RPC
type GRPCClient struct {
	client proto.KVClient
}

func (m *GRPCClient) Put(key string, value []byte) error {
	_, err := m.client.Put(context.Background(), &proto.PutRequest{
		Key:   key,
		Value: value,
	})
	return err
}

func (m *GRPCClient) Get(key string) ([]byte, error) {
	resp, err := m.client.Get(context.Background(), &proto.GetRequest{
		Key: key,
	})
	if err != nil {
		return nil, err
	}

	return resp.Value, nil
}

// GRPCServer is the gRPC Server that the client talks to
type GRPCServer struct {
	// Real implementation
	Impl KV
	proto.UnimplementedKVServer
}

func (m *GRPCServer) Put(ctx context.Context, req *proto.PutRequest) (*proto.PutResponse, error) {
	return &proto.PutResponse{}, m.Impl.Put(req.Key, req.Value)
}

func (m *GRPCServer) Get(ctx context.Context, req *proto.GetRequest) (*proto.GetResponse, error) {
	v, err := m.Impl.Get(req.Key)
	return &proto.GetResponse{Value: v}, err
}
