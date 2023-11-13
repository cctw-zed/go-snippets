package designpattern

import (
	"testing"
	"time"
)

/*
原型模式适用于以下场景：

当对象的创建成本较高，而复制现有对象的成本较低时。
当需要创建大量相似对象，且这些对象之间的差异很小或可以动态修改时。
当需要避免使用类层次结构来创建具有许多可配置属性的对象时。
以下是一个使用原型模式的示例：

假设我们正在开发一个游戏，游戏中有许多相似的敌人。这些敌人的属性（如生命值、速度等）可能略有不同，但它们的基本结构和行为是相同的。我们可以使用原型模式来创建新的敌人，而无需每次都实例化一个新对象。
*/
func TestPrototype(t *testing.T) {
	start := time.Now()

	// 通过调用 NewMapData() 创建地图对象，这可能需要花费较长时间
	map1 := &Map{Data: NewMapData()}
	t.Logf("Map1 created in %v\n", time.Since(start))

	start = time.Now()

	// 通过复制现有地图对象来创建新地图对象，这将显著减少创建时间
	map2 := map1.Clone()
	t.Logf("Map2 created in %v\n", time.Since(start))

	t.Logf("Map1: %+v, Map2: %+v", map1, map2)
}
