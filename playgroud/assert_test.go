package playgroud

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMapAssert(t *testing.T) {

	// 这个测试用例的目的是检查 inter 变量是否可以断言为 map[interface{}]interface{} 和 map[string]interface{} 类型。
	// 从代码中可以看出，inter 变量实际上是一个 map[string]interface{} 类型，因此第二个断言应该成功。而第一个断言将会失败，
	// 因为 Go 语言的类型系统不允许隐式类型转换。
	Convey("1", t, func() {
		var inter interface{}
		inter = map[string]interface{}{}
		_, ok := inter.(map[interface{}]interface{})
		t.Logf("is map[interface{}]interface{}: %+v", ok)

		_, ok = inter.(map[string]interface{})
		t.Logf("is map[string]interface{}: %+v", ok)
	})

	Convey("2", t, func() {
		var inter interface{}
		inter = map[string]interface{}{}
		_, ok := inter.(map[interface{}]interface{})
		t.Logf("is map[interface{}]interface{}: %+v", ok)

		_, ok = inter.(map[string]interface{})
		t.Logf("is map[string]interface{}: %+v", ok)
	})
}
