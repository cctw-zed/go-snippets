package designpattern

import "time"

type MapData struct {
	TerrainData  []byte
	RoadData     []byte
	BuildingData []byte
}

func NewMapData() *MapData {
	// 模拟从远程服务器加载数据和解析数据的耗时操作
	time.Sleep(2 * time.Second)

	return &MapData{
		TerrainData:  []byte("Terrain data..."),
		RoadData:     []byte("Road data..."),
		BuildingData: []byte("Building data..."),
	}
}

type MapPrototype interface {
	Clone() MapPrototype
}

type Map struct {
	Data *MapData
}

func (m *Map) Clone() MapPrototype {
	return &Map{
		Data: m.Data, // 注意：这里我们复用了相同的 MapData 实例
	}
}
