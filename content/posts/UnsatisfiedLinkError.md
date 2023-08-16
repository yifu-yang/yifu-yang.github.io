---
title: "解决在调试Spring Boot和Https程序时的一个环境问题"
date: 2020-05-18T14:29:11+08:00
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

## 问题描述

在使用Springboot 2.2.4release版本进行开发时，在程序内部配置了https后，出现了在本地调试无法访问接口，报出以下错误：（错误信息借用自链接: [此博客](https://www.coder.work/article/1410970)）

```log
2018-12-03 14:23:46,089 PID=21726 LEVEL=ERROR THREAD=https-openssl-nio-443-exec-2 LOGGER=org.apache.tomcat.util.net.NioEndpoint METHOD=log:175 MESSAGE="java.lang.UnsatisfiedLinkError: org.apache.tomcat.jni.SSL.renegotiatePending(J)I
        at org.apache.tomcat.jni.SSL.renegotiatePending(Native Method) ~[tomcat-embed-core-9.0.13.jar!/:9.0.13]
        at org.apache.tomcat.util.net.openssl.OpenSSLEngine.getHandshakeStatus(OpenSSLEngine.java:1021) ~[tomcat-embed-core-9.0.13.jar!/:9.0.13]
        at org.apache.tomcat.util.net.openssl.OpenSSLEngine.wrap(OpenSSLEngine.java:457) ~[tomcat-embed-core-9.0.13.jar!/:9.0.13]
        at java.base/javax.net.ssl.SSLEngine.wrap(SSLEngine.java:471) ~[na:na]
        at org.apache.tomcat.util.net.SecureNioChannel.handshakeWrap(SecureNioChannel.java:440) ~[tomcat-embed-core-9.0.13.jar!/:9.0.13]
        at org.apache.tomcat.util.net.SecureNioChannel.handshake(SecureNioChannel.java:211) ~[tomcat-embed-core-9.0.13.jar!/:9.0.13]
        at org.apache.tomcat.util.net.NioEndpoint$SocketProcessor.doRun(NioEndpoint.java:1394) ~[tomcat-embed-core-9.0.13.jar!/:9.0.13]
        at org.apache.tomcat.util.net.SocketProcessorBase.run(SocketProcessorBase.java:49) [tomcat-embed-core-9.0.13.jar!/:9.0.13]
        at java.base/java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1135) [na:na]
        at java.base/java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:635) [na:na]
        at org.apache.tomcat.util.threads.TaskThread$WrappingRunnable.run(TaskThread.java:61) [tomcat-embed-core-9.0.13.jar!/:9.0.13]
        at java.base/java.lang.Thread.run(Thread.java:844) [na:na]
"
```

其他信息： 开发环境jdk 1.8.181,操作系统 windows 10

## 问题诊断

 1. 首先我以为是浏览器的问题，测试了其他浏览器，没有生效
 2. 尝试切换了其他端口，没有生效，彻底移除https代码可以解决，基本确定在https的问题上
 3. 通过错误信息定位到了native代码，说明还是环境的问题，参考了 [此博客](https://www.coder.work/article/1410970)，发现和本地已有tomcat有关系，检查版本，本地的为7.0.x，更换版本至9.0.30，问题解决

## 参考文章

 [java - Spring Boot 2.1.1:UnsatisfiedLinkError:org.apache.tomcat.jni.SSL.renegotiatePending](https://www.coder.work/article/1410970)