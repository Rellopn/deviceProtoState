/**
 * @Author: Rellopn
 * @Description:
 * @File:  dpstate
 * @Version: 0.1.1
 * @Date: 2020/7/3 9:48
 */
package deviceProtoState

import (
	"errors"
	"github.com/go-redis/redis"
	"github.com/patrickmn/go-cache"
	"time"
)

// protocol type
const (
	MQTT_PROTO = iota
)

type DpState struct {
	Opt         Option
	TTl         int64
	RedisClient *redis.Client
	protoType   int
	Cache       *cache.Cache
	MqttProto
}

func NewDpState(option Option) (*DpState, error) {
	opt := &DpState{Opt: option, Cache: cache.New(5*time.Minute, 10*time.Minute)}
	if option.StateDb == REDIS {
		err := opt.RedisState()
		if err != nil {
			return nil, err
		}
	}
	return opt, nil
}

func (d *DpState) UsePorto(proto int) {
	d.protoType = proto
}

func (d *DpState) RedisState() error {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     d.Opt.Host + ":" + d.Opt.Port,
		Password: d.Opt.Pwd, // password set
		DB:       d.Opt.Db}) // use DB

	_, err := redisClient.Ping().Result()
	if err != nil {
		return err
	}
	d.RedisClient = redisClient
	return nil
}

func PubCheckMsg(state *DpState, IMSI string, retained bool, payload interface{},
	randStr string) (string, error) {
	expireSecond := time.Duration(state.TTl) * time.Second
	if token := state.Mc.Publish("/device/"+IMSI, 1, retained, payload);
		token.WaitTimeout(expireSecond) && token.Error() != nil {
		return "", token.Error()
	}
	if err := state.RedisClient.Set(randStr, payload, expireSecond).Err(); err != nil {
		return "", err
	}

	waitChan := make(chan string)
	state.Cache.Set(randStr, waitChan, expireSecond)
	timer := time.NewTimer(expireSecond)
	defer timer.Stop()
	select {
	case pload := <-waitChan:
		state.Cache.Delete(randStr)
		return pload, nil
	case <-timer.C:
		state.Cache.Delete(randStr)
		return "", errors.New("设备:" + IMSI + "超时")
	}
}
func NotifyWait(state *DpState, waitId, payload string) bool {
	waitChan, found := state.Cache.Get(waitId)
	if found {
		waitChan.(chan string) <- payload
	}
	return found
}
