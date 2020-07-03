/**
 * @Author: Rellopn
 * @Description:
 * @File:  mqttCfg
 * @Version: 0.1.1
 * @Date: 2020/7/3 10:33
 */
package deviceProtoState

import (
	"github.com/eclipse/paho.mqtt.golang"
)

type MqttProto struct {
	Mc       mqtt.Client
}

func (m *MqttProto) AddClient(mc mqtt.Client) {
	m.Mc = mc
}