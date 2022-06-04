package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"reflect"
	"strings"
	"unsafe"
)

type StoreRedis struct {
	pool *redis.Pool
}

func NewStoreRedis(pool *redis.Pool) IRedis {
	sr := new(StoreRedis)
	sr.pool = pool
	return sr
}

/**
此方法没有使用
*/
func (s *StoreRedis) errorLogic(e error) {
	i := strings.Index(e.Error(), "connect: connection refused")
	if i == -1 {
		fmt.Println(i, "不包含")
	}
}

/**
设置key,value数据
*/
func (s *StoreRedis) Set(key, value string) error {
	_, err := s.pool.Get().Do("SET", key, value)
	defer s.Close()
	if err != nil {
		return err
	}
	return nil
}

/**
设置key,value数据
*/
func (s *StoreRedis) SetBytes(key string, value []byte) error {
	_, err := s.pool.Get().Do("SET", key, value)
	defer s.Close()
	if err != nil {
		return err
	}
	return nil
}

/**
设置key的过期时间
*/
func (s *StoreRedis) SetKvAndExp(key, value string, expire int) error {
	_, err := s.pool.Get().Do("SET", key, value, "EX", expire)
	defer s.Close()
	if err != nil {
		return err
	}
	return nil
}

func (s *StoreRedis) SetKvInt(key string, value int) error {
	_, err := s.pool.Get().Do("SET", key, value)
	defer s.Close()
	if err != nil {
		return err
	}
	return nil
}

/**
根据key获取对应数据
*/
func (s *StoreRedis) Get(key string) string {
	value, err := redis.String(s.pool.Get().Do("GET", key))
	defer s.Close()
	if err != nil {
		fmt.Println("redis get failed:", err)
	}
	return value
}

/**
根据key获取对应数据
*/
func (s *StoreRedis) GetBytes(key string) []byte {
	value, err := s.pool.Get().Do("GET", key)
	defer s.Close()
	if err != nil {
		fmt.Println("redis get failed:", err)

	}

	if value == nil {
		return nil
	}

	return value.([]byte)
}

/**
判断key是否存在
*/
func (s *StoreRedis) IsKeyExists(key string) bool {
	is_key_exit, _ := redis.Bool(s.pool.Get().Do("EXISTS", key))
	defer s.Close()
	return is_key_exit
}

/**
删除key
*/
func (s *StoreRedis) Del(key string) bool {
	is_key_delete, err := redis.Bool(s.pool.Get().Do("DEL", key))
	defer s.Close()
	if err != nil {
		fmt.Println("error:", err)
	}
	return is_key_delete
}

/**
删除key
*/
func (s *StoreRedis) HDel(key string, field string) bool {
	is_key_delete, err := redis.Bool(s.pool.Get().Do("HDEL", key, field))
	defer s.Close()
	if err != nil {
		fmt.Println("error:", err)
	}
	return is_key_delete
}

/**
对象转换成json后进行存储
*/
func (s *StoreRedis) Setnx(key string, value []byte) error {
	_, err := s.pool.Get().Do("SETNX", key, value)
	defer s.Close()
	if err != nil {
		return err
	}
	return nil
}

func (s *StoreRedis) LpushByte(key string, value []byte) error {
	_, err := s.pool.Get().Do("lpush", key, value)
	if err != nil {
		defer s.Close()
		return err
	}
	defer s.Close()
	return nil
}

func (s *StoreRedis) RPopByte(key string) ([]byte, error) {
	v, err := redis.Bytes(s.pool.Get().Do("rpop", key))
	if err != nil {
		defer s.Close()
		return nil, err
	}
	defer s.Close()
	return v, nil
}

func (s *StoreRedis) LPop(key string) (string, error) {
	v, err := s.pool.Get().Do("lpop", key)
	if err != nil {
		defer s.Close()
		return "", err
	}
	defer s.Close()
	if v == nil {
		return "", nil
	}
	vv := BytesToString(v.([]byte))
	return vv, nil
}

func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

func (s *StoreRedis) LLen(key string) (int64, error) {
	v, err := s.pool.Get().Do("llen", key)
	if err != nil {
		defer s.Close()
		return 0, err
	}
	defer s.Close()
	if v == nil {
		return 0, nil
	}
	return v.(int64), nil
}

func (s *StoreRedis) Close() {
	s.pool.Get().Close()
}

/**
Hincryby方法
*/
func (s *StoreRedis) HINCRBY(key, field string) {
	s.pool.Get().Do("HINCRBY", key, field, 1)
	defer s.Close()
}

/**
HGet方法
*/
func (s *StoreRedis) HGet(key, field string) (interface{}, error) {
	value, err := s.pool.Get().Do("HGET", key, field)
	defer s.Close()
	return value, err
}

/**
HGetAll方法
*/
func (s *StoreRedis) HGetAll(key string) ([][]byte, error) {
	inter, err := redis.ByteSlices(s.pool.Get().Do("HGetAll", key))
	defer s.Close()
	return inter, err
}

/**
Hset方法
*/
func (s *StoreRedis) HSet(key string, field string, value string) (interface{}, error) {
	inter, err := s.pool.Get().Do("HSET", key, field, value)
	defer s.Close()
	return inter, err
}

/**
Hset方法
*/
func (s *StoreRedis) HSetByte(key string, field string, value []byte) (interface{}, error) {
	inter, err := s.pool.Get().Do("HSET", key, field, value)
	defer s.Close()
	return inter, err
}

/**
SADD方法
*/
func (s *StoreRedis) SAdd(args []interface{}) (interface{}, error) {
	inter, err := s.pool.Get().Do("SADD", args...)
	defer s.Close()
	return inter, err
}

/**
Scard方法
*/
func (s *StoreRedis) SCard(key string) (interface{}, error) {
	inter, err := s.pool.Get().Do("SCARD", key)
	defer s.Close()
	return inter, err
}

/**
Spop方法
*/
func (s *StoreRedis) SPop(key string) (interface{}, error) {
	inter, err := s.pool.Get().Do("SPOP", key)
	defer s.Close()
	vv := BytesToString(inter.([]byte))
	return vv, err
}

func (s *StoreRedis) RPopLPushByte(key, key1 string) ([]byte, error) {
	v, err := redis.Bytes(s.pool.Get().Do("RPOPLPUSH", key, key1))
	if err != nil {
		defer s.Close()
		return nil, err
	}
	defer s.Close()
	return v, nil
}
