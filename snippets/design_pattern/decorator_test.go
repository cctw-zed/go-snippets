package designpattern

import (
	"fmt"
	"testing"
)

/*
装饰是一种结构设计模式， 允许你通过将对象放入特殊封装对象中来为原对象增加新的行为。

由于目标对象和装饰器遵循同一接口， 因此你可用装饰来对对象进行无限次的封装。 结果对象将获得所有封装器叠加而来的行为。

地铁进站的过程一般情况下只需要买票，检票进站，但是如果你是带行李，就需要进行安全检查，如果是疫情时期，就需要进行疫情防护检查，比如戴口罩、测量体温等，这里买票进站相当于通用进站流程，安检及防疫检查就相当于加强的修饰行为；
*/
func TestDecorator(t *testing.T) {
	xierqiStation := NewSubwayStation("西二旗")
	fmt.Println(EnhanceEnterStationProcess(xierqiStation, false, false).Enter())
	fmt.Println(EnhanceEnterStationProcess(xierqiStation, true, false).Enter())
	fmt.Println(EnhanceEnterStationProcess(xierqiStation, true, true).Enter())
}

// EnhanceEnterStationProcess 根据是否有行李，是否处于疫情，增加进站流程
func EnhanceEnterStationProcess(station Station, hasLuggage bool, hasEpidemic bool) Station {
	if hasLuggage {
		station = NewSecurityCheckDecorator(station)
	}
	if hasEpidemic {
		station = NewEpidemicProtectionDecorator(station)
	}
	return station
}
