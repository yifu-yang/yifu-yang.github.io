---
title: "电脑里的时钟是怎么工作的（译）"
date: 2023-08-24T22:56:03+08:00
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

即便是重启你的电脑之后，你依然可以得到准确的时间和日期。你想过这是怎么实现的吗？这个答案非常简单；一定有一个独立的电源支持这一个始终运行，即便电脑的主电源关闭了。现在来解释这一切。

RTC-实时时钟

在电脑里有两个时钟，一个是硬件的实时时钟，另一个是软件时钟。实时时钟是使用备用电池的时钟，所以它可以在电脑关机或者低电量的时候继续计时。基本上RTC不是一个物理上的时钟而是一个主板上的集成电路为系统的计时函数和系统时钟服务。实时时钟确保响应所有处理器的请求都是同步的（基本上这是系统时钟的任务，但是系统时钟依赖RTC，因此RTC是间接响应系统中断，时钟，任务计划，以及同步等等）。今天许多公司像飞利浦，ST微电子，德州仪器，制造RTC。已经有对RTC的进一步开发，例如低功耗，提高频率稳定。

系统时钟主要用操作系统内核维护，并且用来设置任务和线程-他们的同步和计划，设置和管理终端，设置时间等。系统时钟上报秒和毫秒当系统启动时的开始点。基本上是系统时钟是一个数字信号发射器。发射一个信号包括高（1）和低（0），因为所有机器和他们的处理器理解二进制语言。

工作原理

正如前面提到的实时时钟是集成电路。它包括一个晶体振荡器。晶体振荡器使用一个压电式晶体来产生信号。晶体包括一个晶体结构（规律的并且重复的原子结构模式）。当在晶体施加电场，他的晶体结构会变形，为了排除这个场它会回反馈到他的晶体结构，然后产生一个非常规律的电信号。

这个晶体的主要特性是频率稳定性在温度，加载和切换电源时产生信号的频率不应该发生变化。晶体振荡器擅长在场中稳定稳定震荡并且成本划算。

通常振荡器频率被设置为32.768KHz，因为这个频率被用在时钟所以设置系统时钟和计时器方便。

实时时钟的电源通常使用锂电池供电。RTC还应该满足最小电流要求，最小的电流应该是100nA，来保持RTC运行一个很长的时间。当主板供电时，这个电池将被充电，然后准备将来使用。

原文链接：https://www.engineersgarage.com/how-clock-in-computer-works/