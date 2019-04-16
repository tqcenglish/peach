---
name: 创建文档仓库
---

## 创建文档仓库

每一个 Peach 文档仓库都包含两部分内容：

- TOC.ini
- 针对每个语言的文档
 
仓库结构大致如下：

```sh
$ tree
.
├── TOC.ini
├── zh-CN
│   ├── advanced
│   │   ├── README.md
│   │   └── ...
│   ├── faqs
│   │   └── README.md
│   ├── howto
│   │   ├── README.md
│   │   ├── ...
│   └── intro
│       ├── README.md
│       ├── ...
└── en-US
│   ├── ...
```
## TOC.ini

在仓库的根目录，您必须创建一个名为 `TOC.ini` 的文件，也就是所谓的 **Table Of Content**。

在这个文件中，您需要使用 [INI](https://en.wikipedia.org/wiki/INI_file) 语法来定义显示哪些目录和文件，以及它们的显示顺序。

下面为 [Peach 文档](http://peachdocs.org) 的 `TOC.ini` 文件：

```ini
-: intro
-: howto
-: advanced
-: faqs

[intro]
-: README
-: installation
-: getting_started
-: roadmap

[howto]
-: README
-: documentation
-: webhook
-: templates
-: static_resources
-: disqus
-: ga

[advanced]
-: README
-: config_cheat_sheet

[faqs]
-: README
```

:speech_balloon: 您可能已经注意到，K-Peach 只**支持一层目录**结构。

在默认分区中，您可以定义显示哪些目录以及它们的显示顺序：

```ini
-: intro
-: howto
-: advanced
-: faqs
```

这些名称必须和目录名称一致。

然后再为每一个目录创建一个分区，顺序在这里是无所谓的：

```ini
[intro]
...
[howto]
...
[advanced]
...
[faqs]
...
```

在每个分区中，您可以定义显示哪些文件以及它们的显示顺序：

```ini
[intro]
-: README
-: installation
-: getting_started
-: roadmap
```

因为文件已经默认使用 Markdown 语法，并且必须以 `.md` 作为扩展名，所以您完全不需要在 `TOC.ini` 文件中说明。

:exclamation: :exclamation: :exclamation:

- 每个分区必须至少包含一个键
- 每个分区的第一个键用于指示该目录的信息
- 这个文件本身不会作为文档单独显示，但是会以目录的形式显示。例如：[简介](../intro)
- 该键的名称是随意的，但一般约定使用 `README`，即使用 `README.md` 作为文件名

## 本地化

在仓库的根目录，您需要为每个支持的语言创建一个名称符合 [Language Culture Name](https://msdn.microsoft.com/en-us/library/ee825488\(v=cs.20\).aspx) 的相应目录。

K-Peach 支持英语（`en-US`）和简体中文（`zh-CN`），所以文档的目录结构为：

```sh
$ tree
.
├── en-US
│   ├── ...
└── zh-CN
│   ├── ...
```

当然，这两个目录拥有完全相同的目录结构和文件名称。

## 文档内容

每个文件都必须在最开头的部分定义自身的信息，然后才是文档的内容：

```
---
name: 简介
---

# K-Peach

K-Peach 是一个简单的知识库系统
...
```

如果您的目录不包含任何文档内容，只是代表分类，则可以省略文档部分：

```
---
name: 高级用法
---
```

## 链接跳转

渲染链接的方式和其它地方大体相同：

- 链接到相同目录的其它页面：`[Webhook](webhook)`.
- 链接到某个目录：`[Introduction](../intro)`.
- 链接到其它目录下的页面：`[Getting Started](../intro/getting_started)`.

## 链接图片

默认情况下，所有的文档页面都会使用 `/docs` 作为 URL 前缀。并且您所有的图片都必须存放于仓库根目录下名为 `images` 的子目录。

然后通过这种语法来链接图片：`![](/docs/images/github_webhook.png)`