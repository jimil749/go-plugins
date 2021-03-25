package utils

type PutRequest struct {
	Key   string
	Value []byte
}

type PutResponse struct{}

type GetRequest struct {
	Key string
}

type GetResponse struct {
	Value []byte
}
