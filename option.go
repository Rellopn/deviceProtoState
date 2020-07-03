/**
 * @Author: Rellopn
 * @Description:
 * @File:  option
 * @Version: 0.1.1
 * @Date: 2020/7/3 9:50
 */
package deviceProtoState

const (
	REDIS = iota + 1
	MYSQL
)

type redisCfg struct {
	Host string
	Port string
	Pwd  string
	Db   int
}

type Option struct {
	StateDb int
	redisCfg
}

func (o *Option) RedisStateCfg(host, port, pwd string, db ...int) {
	o.Host, o.Port, o.Pwd = host, port, pwd
	if len(db) > 0 {
		o.Db = db[0]
	}
	o.StateDb = REDIS
}
