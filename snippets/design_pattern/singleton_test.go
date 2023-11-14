package designpattern

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
单例模式适用于以下场景：

当一个类只需要有一个实例，并且这个实例需要在整个应用程序中共享时。
当一个类需要控制对全局共享资源的访问时，例如配置信息、日志记录、数据库连接池等。
以下是一个使用单例模式的示例：

假设我们正在开发一个 Web 服务器应用程序，该应用程序需要从配置文件中读取一些配置信息。为了避免多次读取配置文件，我们可以使用单例模式创建一个配置管理类，该类只有一个实例，并在整个应用程序中共享。这样，我们可以确保配置信息只被读取一次，并在需要时随时访问。

*/

func TestSingleton(t *testing.T) {

	// 获取单例配置实例
	config1 := GetConfigInstance()
	t.Logf("Config1: %+v", config1)

	// 再次获取单例配置实例
	config2 := GetConfigInstance()
	t.Logf("Config2: %+v", config2)

	// config1 和 config2 指向相同的实例
	assert.Equal(t, config1, config2)
}
