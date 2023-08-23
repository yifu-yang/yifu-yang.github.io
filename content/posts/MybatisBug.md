---
title: "发现Mybatis中的一个bug"
date: 2021-05-13T18:29:13+08:00
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

## 现象

使用```<foreach>```标签进行循环遍历时，如果多于一次使用，在第二次使用foreach时，会出现```index```数值接在上此次```foreach```的···index···的后面，而不是从0开始。
例如

```SQL
<select id="test' resultType="java.util.Map">
 select 
 <foreach index="index" collection="list" seprator=",">
 ${index} as col${index}
 </foreach>
 ,
 <foreach index="index1" collection="list" seprator=",">
 ${index1} as col${index1}
 </foreach>
</select>
```

传入参数是一个长度为5的列表
预期返回的长度是5，实际返回了10

## 存在版本

在目前的验证中，Myabtis 3.2是存在这个问题的，3.4目前的测试没有出现这个问题
