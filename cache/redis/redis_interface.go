package redis

type IRedis interface {
	Set(key, value string) error
	Get(key string) string
	LpushByte(key string, value []byte) error
	Close()
	RPopByte(key string) ([]byte, error)
	LLen(key string) (int64, error)
	RPopLPushByte(key, key1 string) ([]byte, error)
}
