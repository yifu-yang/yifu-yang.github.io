---
title: "为什么编写单元测试对不同的人有不同的理解（译）"
date:  2021-07-07T22:35:35+08:00
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

每当我第一次和我的同事讨论TDD（测试驱动开发）时，我首先会试图去区分他/她眼中的的TDD。当我拿到一份新代码时，我也会这么做。我阅读代码并尝试分辨这份代码是用传统派还是模拟派来实现TDD的。



对于我来说，这两种最主要的区别是传统派试图在修改具体的代码实现时，不阻断测试用例的运行（例如重构），而模拟派则并不关心，他们之关系是否是独立的单元。



无论你使用何种风格，我都试图在一份代码库里保持风格的一致，尽管我更偏向于传统派。我确信，片段测试的代码所带来的问题远大于其收益。在我看来，对于独立的代码测试仅仅适用于需要隔离实际的外部调用（例如网络服务）。



有趣的是，最初单元测试之所以被称为单元测试，是因为他们可以以独立的单元运行（独立于其他的测试用例），并不是他们只测试独立的单元。



如果你在重构项目时使用了过多的片段式的单元测试，那么在修改具体的代码实现后，你通常会忙于修复测试用例。这只是让你的测试用例满足于你而不是服务于你。



如果你有一整套良好的测试，那么当你修改了代码实现后，并不需要去这些用例。相反，这些用例只会告诉你代码是否正常工作。定义公共的接口并在你的测试用例里使用这些接口，是实现这个目标的关键。如果你发现一部分代码无法被公共的接口覆盖，那这对你来说也许是个好消息:你应该删掉这些代码了。

原文链接：https://wolfoliver.medium.com/why-writing-unit-tests-means-something-else-to-different-people-de6ef23b213d