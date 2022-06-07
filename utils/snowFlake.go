package utils

import (
	"sync"
	"time"
)

var (
	machineID     int64 // 机器 id 占10位, 十进制范围是 [ 0, 1023 ]
	sn            int64 // 序列号占 12 位,十进制范围是 [ 0, 4095 ]
	lastTimeStamp int64 // 上次的时间戳(毫秒级), 1秒=1000毫秒, 1毫秒=1000微秒,1微秒=1000纳秒
	mu            sync.Mutex
)

func init() {
	machineID = 101 << 12
	lastTimeStamp = time.Now().UnixNano() / 1000
}

func GetSnowflakeIdProcess() int64 {
	curTimeStamp := time.Now().UnixNano() / 1000
	// 同一毫秒
	if curTimeStamp == lastTimeStamp {
		// 序列号占 12 位,十进制范围是 [ 0, 4095 ]
		if sn > 4095 {
			time.Sleep(time.Microsecond)
			curTimeStamp = time.Now().UnixNano() / 1000
			sn = 0
		}
	} else {
		sn = 0
	}
	sn++

	lastTimeStamp = curTimeStamp
	// 取 64 位的二进制数 0000000000 0000000000 0000000000 0001111111111 1111111111 1111111111  1 ( 这里共 41 个 1 )和时间戳进行并操作
	// 并结果( 右数 )第 42 位必然是 0,  低 41 位也就是时间戳的低 41 位
	rightBinValue := curTimeStamp & 0x1FFFFFFFFFF
	// 机器 id 占用10位空间,序列号占用12位空间,所以左移 22 位; 经过上面的并操作,左移后的第 1 位,必然是 0
	rightBinValue <<= 22
	id := rightBinValue | machineID | sn
	return id
}

func GetSnowflakeId() int64 {
	mu.Lock()
	defer mu.Unlock()
	return GetSnowflakeIdProcess()
}
