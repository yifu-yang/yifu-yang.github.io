---
title: "TCP结束了吗？谷歌的QUIC传输协议加速了互联网旗舰TCP（译）"
date: 2021-07-31T22:39:54+08:00
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

八年后，谷歌计划代替TCP的QUIC，似乎终于要开始了。事实上，IETF发布了以QUIC作为互联网标准在本周早些时候。如果成功，这个协议将代替从974开始就主导互联网传输领域的TCP。



事实上，在线服务和网络浏览器已经实验这项技术多年。然而，现在IETF正式发布了这项标准，全球用户更倾向于全面转换到QUIC。



谷歌首次披露QUIC作为Chrome浏览器的一个实验的附加功能还要追溯到2013年。也就是说，对于在遗留的网络传输协议和基础设施中海量复杂的数据，设备，程序和服务，QUIC给了相当多的时间来开发。然而，随着世界的进步和我们的互联网必须适应管理越来越多的数据，像升级到QUIC，HTTPS安全通信和后量子密码学来保护数据免于未来潜在的量子计算机，也包括升级到IPV6，都已经在积极适应一个快速发展的数据和用户逐渐增长的虚拟世界。



事实上，谷歌已经宣布QUIC承诺降低8%的PC网络搜索等待时间和4%的手机搜索等待时间。相同的是，QUIC也在YouTube表现出PC端的18%和移动设备的15%的缓冲时间的降低。



从核心上说，QUIC与TCP/IP在计算机之间通过网络合作传输数据。TCP总是分割数据到数据包并且确保这些包能够通过网络路由到达目的地。简单地说，这些连接必须必须被设计的能承受和攻击，能恢复因为干扰而传输中丢失的数据包。



另一方面，QUIC使用UDP额外的好处，更快速但是无连接相比于TCP，缺少后者的数据恢复机制。为了实现UDP的速度和TCP的保证数据安全，QUIC采用了比使用TCP更快的自己的恢复方法。



此外，对于之前使用的支持原有HTTP标准的浏览器，QUIC也比TCP能更快的建立加密连接。这种双重优势将不仅提高传输安全性而且将降低浏览器网页加载的时间。



最后，QUIC甚至几乎更容易地在两种网络中传输，就像用户从家中Wi-Fi切换到移动网络。



最近关注到QUIC是在YouTube上看了一个关于RSocket对比gRPC-over-QUIC的演讲看到的。同时QUIC也将作为HTTP/3的传输协议，文中提到的Chrome已经支持，web服务器上，Nginx和Caddy也都有了支持。还是挺有趣的。



原文链接：https://www.tensorbugs.in/2021/06/is-that-tcp-ends-googles-quic.html

