package designpattern

import "testing"

/*
工厂方法模式适用于以下场景：

当一个类无法预知它需要创建哪种类的对象时。
当一个类希望由其子类来指定所创建的对象时。
当将对象创建过程与使用过程解耦有助于提高代码的灵活性和可维护性时。
以下是一个使用工厂方法模式的示例：

假设我们正在开发一个日志记录库，该库支持将日志记录到不同的目标，如文件、数据库或远程服务器等。我们可以使用工厂方法模式来创建一个灵活的日志记录器，该日志记录器可以根据配置轻松切换日志目标。

在这个示例中，我们定义了一个名为 Logger 的接口，它表示可以记录日志的对象。我们还定义了两个具体的日志记录器：FileLogger 和 DatabaseLogger，它们分别将日志记录到文件和数据库。然后，我们创建了一个名为 LoggerFactory 的接口，它包含一个用于创建 Logger 实例的方法。我们为每种日志记录器实现了一个具体的工厂：FileLoggerFactory 和 DatabaseLoggerFactory。要切换日志目标，我们只需更改工厂实例即可。

工厂方法模式使我们能够将对象的创建过程与使用过程分离，从而提高代码的灵活性和可维护性。此外，我们可以轻松地添加新的日志目标，而无需修改现有的日志记录器或工厂代码。
*/
func TestFactory(t *testing.T) {

	var factory LoggerFactory

	// 使用 FileLogger
	factory = &FileLoggerFactory{}
	logger := factory.CreateLogger()
	logger.Log("Hello, World!")

	// 使用 DatabaseLogger
	factory = &DatabaseLoggerFactory{}
	logger = factory.CreateLogger()
	logger.Log("Hello, World!")
}
