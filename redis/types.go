package redis

//Key data stored in redis
type Key struct {
	ID          string
	SecretToken string
}

type Request struct {
	Path string
	Body string //@Todo should be rosetta-request type
}

type Response struct {
	Body string //@Todo should be rosetta-response type
}
