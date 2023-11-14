package playgroud

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMap(t *testing.T) {
	Convey("1", t, func() {
		var targetMap interface{}
		targetMap = make(map[interface{}]interface{})
		subMap0 := make(map[string]interface{})
		subMap1 := make(map[string]interface{})
		subMap2 := make(map[string]interface{})

		targetMap = subMap0
		targetMap = subMap1
		targetMap = subMap2

		t.Logf("targetMap: %+v", targetMap)

	})
}
