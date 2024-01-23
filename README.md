# design-patterns-go
使用go语言，分别实现23中设计模式，对应的简介、源码、作用和使用场景，可移步博客。

# 设计模式系列
> 设计模式有助于提高软件开发质量、可维护性和可扩展性，设计模式并非一成不变的规则，是需要根据具体的业务去具体的调整和应用，下面给出23中设计模式中简单的例子，以及作用和使用场景。设计模式可以分为创建型、结构型和行为型三种。

使用go语言，每个模式都进行单元测试，简单易懂的学习、了解各个设计模式的案例，从最最简单入手，然后再代入到实现具体业务的时候，应该采用哪种设计模式。


当然，基于各个设计模式的介绍、作用和使用场景，具体的我放到我的博客当中，可以点击下面链接直接进行访问。

### 创建型
- [单例模式](https://boxiaoyang.club/article/208)

  单例模式通常用于控制对唯一实例的访问，以便在系统中的多个部分之间共享该实例。
- [ 工厂方法模式](https://boxiaoyang.club/article/190)

  定一个用于创建对象的接口，但将实例化的过程延迟到子类中，就是由子类去决定实例化哪个对象
- [抽象工厂模式 ](https://boxiaoyang.club/article/191)

  创建相关或依赖对象的家族，而不需要指定它们的具体类
- [ 建造者模式](https://boxiaoyang.club/article/193)

  将一个复杂对象的构建过程和它的表示进行分离，使得同样的构建过程可以创建不同的表示
- [原型模式 ](https://boxiaoyang.club/article/202)

  复制现有对象来创建新对象，而不是通过实例化类。

### 结构型模式
- [ 适配器模式](https://boxiaoyang.club/article/209)

  它允许将一个类的接口转换成客户端期望的另一个接口。
- [装饰器模式 ](https://boxiaoyang.club/article/211)

  它允许向一个现有对象添加新的功能，同时又不改变其结构。
- [代理模式 ](https://boxiaoyang.club/article/210)

  它充当其他对象的接口，以控制对这个对象的访问。
- [外观模式](https://boxiaoyang.club/article/192)

  为复杂子系统提供一个简化的接口，从而使客户端代码更容易使用。
- [桥接模式](https://boxiaoyang.club/article/194)

  通过组合的方式，将抽象类和实现类分离，使得它们可以独立变化而不互相影响。
- [组合模式](https://boxiaoyang.club/article/206)

  将对象组织成树形结构，使得客户端可以一致地处理单个对象和组合对象。
- [享元模式](https://boxiaoyang.club/article/205)

  通过共享对象来减小内存使用或减少计算开销。
### 行为型模式
- [策略模式](https://boxiaoyang.club/article/204)

  允许客户端选择算法的实现方式，使得客户端和算法之间解耦。
- [模板方法模式](https://boxiaoyang.club/article/197)

  它定义了一个实现模板的骨架，将一些步骤推迟到子类。
- [观察者模式](https://boxiaoyang.club/article/212)

  定义了一种一对多的依赖关系，当一个对象的状态发生改变时，所有依赖于它的对象都会得到通知并自动更新。
- [ 迭代器模式](https://boxiaoyang.club/article/196)

  提供一种方法来顺序访问一个聚合对象中的各个元素。
- [责任链模式](https://boxiaoyang.club/article/200)

  通过一系列处理者（处理对象）来逐步处理请求，直到请求被处理为止。
- [命令模式](https://boxiaoyang.club/article/195)

  它将请求封装成一个对象，从而使得可以参数化客户端对象
- [备忘录模式](https://boxiaoyang.club/article/199)

  它允许在不暴露对象内部状态的情况下，捕获并保存对象的内部状态，以便将对象恢复到之前的状态。
- [状态模式](https://boxiaoyang.club/article/203)

  它允许对象在内部状态发生改变时改变其行为。
- [访问者模式](https://boxiaoyang.club/article/198)

  通过在被访问的对象中加入一个对外提供的接受访问者的接口，使得访问者可以访问到对象的内部状态。
- [中介者模式](https://boxiaoyang.club/article/201)

  减少对象之间的直接通信，将复杂的交互关系变得更加松散，从而提高系统的可维护性和可扩展性。
- [解释器模式](https://boxiaoyang.club/article/207)

  定义了一个语言的文法，设计该语言的解释器，用户可以使用特定的解释器来解释给定的语言表达式。