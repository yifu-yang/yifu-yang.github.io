---
title: "PowerMockito使用中一个需要注意的点"
date: 2021-08-07T22:51:20+08:00
# weight: 1
# aliases: ["/first"]
tags: ["archive"]
author: "blackjack"
# author: ["Me", "You"] # multiple authors
showToc: true
TocOpen: false
draft: false
hidemeta: false
comments: false
description: ""
canonicalURL: "https://canonical.url/to/page"
disableHLJS: true # to disable highlightjs
disableShare: false
disableHLJS: false
hideSummary: false
searchHidden: true
ShowReadingTime: true
ShowBreadCrumbs: true
ShowPostNavLinks: true
ShowWordCount: true
ShowRssButtonInSectionTermList: true
UseHugoToc: true
cover:
    image: "<image path/url>" # image path/url
    alt: "<alt text>" # alt text
    caption: "<text>" # display caption under cover
    relative: false # when using page bundles set this to true
    hidden: true # only hide on current single page
editPost:
    URL: "https://github.com/yifu-yang/yifu-yang.github.io/content"
    Text: "Suggest Changes" # edit text
    appendFilePath: true # to append file path to Edit link
---

## 问题现象

最近在使用PowerMockito进行单元测试时，遇到了一个问题，代码如下：
原方法签名

```Java
ContactInfo getInfoById(long id);
```

测试用例中

```Java
PowerMockito.when(mapper.getInfoById(any())).thenReturn(new ContactInfo());
```

这是要mock一个方法返回一个``ContactInfo``对象，非常简单，但是在运行测试用例的时候，出现了如下错误：

```Java
java.lang.NullPointerException
(和业务相关的已经隐藏，报错的就是上面一行)
 at sun.reflect.NativeMethodAccessorImpl.invoke0(Native Method)
 at sun.reflect.NativeMethodAccessorImpl.invoke(NativeMethodAccessorImpl.java:62)
 sun.reflect.DelegatingMethodAccessorImpl.invoke(DelegatingMethodAccessorImpl.java:43)
 at java.lang.reflect.Method.invoke(Method.java:498)
 at org.junit.runners.model.FrameworkMethod$1.runReflectiveCall(FrameworkMethod.java:59)
 at org.junit.internal.runners.model.ReflectiveCallable.run(ReflectiveCallable.java:12)
 at org.junit.runners.model.FrameworkMethod.invokeExplosively(FrameworkMethod.java:56)
 at org.junit.internal.runners.statements.InvokeMethod.evaluate(InvokeMethod.java:17)
 at org.springframework.test.context.junit4.statements.RunBeforeTestExecutionCallbacks.evaluate(RunBeforeTestExecutionCallbacks.java:74)
 at org.springframework.test.context.junit4.statements.RunAfterTestExecutionCallbacks.evaluate(RunAfterTestExecutionCallbacks.java:84)
 at org.junit.internal.runners.statements.RunBefores.evaluate(RunBefores.java:26)
 at org.springframework.test.context.junit4.statements.RunBeforeTestMethodCallbacks.evaluate(RunBeforeTestMethodCallbacks.java:75)
 at org.springframework.test.context.junit4.statements.RunAfterTestMethodCallbacks.evaluate(RunAfterTestMethodCallbacks.java:86)
```

mapper对象的代理已经生成，但空指针表示有地方获取为空了。

## 排查结果

经过大量排查，发现需要mock的函数参数为``long``类型，使用``any()``做mock的时候会爆出空指针异常，需要替换为基本类型的``anyLong()``。

## 原因

``any()``返回的是一个``null``值,而``anyLong()``返回的是0。
在mock方法时，传入为原始对象要执行拆箱操作，``any()``返回了``null``引发空指针。
在Oracle的官方文档里可以看到自动拆箱的两种触发情况描述如下:[1]
>The Java compiler applies unboxing when an object of a wrapper class is:
>
> 1. Passed as a parameter to a method that expects a value of the corresponding primitive type.
> 2. Assigned to a variable of the corresponding primitive type.

1.作为参数传递给一个期望获得原始类型的方法
2.对一个原始类型的对象赋值的时候
本例中应该是属于第一种，``null``在调用``longValue()``时出现空指针异常

引用
[1]https://docs.oracle.com/javase/tutorial/java/data/autoboxing.html