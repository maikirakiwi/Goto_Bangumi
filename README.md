<p align="center">
    <img src="docs/image/icons/light-icon.svg#gh-light-mode-only" width=50%/ alt="">
    <img src="docs/image/icons/dark-icon.svg#gh-dark-mode-only" width=50%/ alt="">
</p>
<p align="center">
    <img title="release date" src="https://img.shields.io/github/release-date/maikirakiwi/Goto_Bangumi" alt="">
    <img title="active issues" src="https://app.deepsource.com/gh/maikirakiwi/Goto_Bangumi.svg/?label=active+issues&show_trend=true&token=3Yw_e7Z12Jeu2AUix9d-PrHU" alt="">
    <img title="go version" src="https://img.shields.io/github/go-mod/go-version/maikirakiwi/Goto_Bangumi/go-main%2Fgo_backend?label=go" alt="">
</p>

<p align="center">
  <a href="https://www.autobangumi.org">官方网站</a> | <a href="https://t.me/autobangumi">TG 群组</a>
</p>

# 项目说明

<p align="center">
    <img title="AutoBangumi" src="docs/image/preview/window.png" alt="" width=75%>
    <img title="Gopher" src="docs/image/preview/gopher.svg" alt="" width=19%>
</p>

本项目是社区基于 [AutoBangumi](https://github.com/EstrellaXD/Auto_Bangumi/) 的 Golang 后端实现。只需要在 [Mikan Project][mikan] 等网站上订阅番剧，就可以全自动追番。
并且整理完成的名称和目录可以直接被 [Plex][plex]、[Jellyfin][plex] 等媒体库软件识别，无需二次刮削。

**兼容目标为 [AutoBangumi 3.1.10](https://github.com/EstrellaXD/Auto_Bangumi/releases/tag/3.1.10)，具体文档 TBA。**

## 手动编译指南
1. 请先安装 [Node.js LTS](https://nodejs.org) 和最新的 [Go](https://go.dev/dl)。
2. `cd AB/webui` 里用 `npm install && npm run build:go`
3. `cd ../../bin` 并用 `./goto_bangumi` 执行编译好的二进制程序。

如在 2. 下显示无 `AB` 目录，请执行 `git submodule update --init --recursive`
如在 3. 下无法执行，则有可能需要将二进制程序设为可执行的。例：`chmod +x ./goto_bangumi` 
执行时可以添加 `dev` 参数来进入开发环境，但请**千万**不要在生产环境下使用。

## FAQ (待翻译）
### - Why not C/C++/Rust?
> Golang, in my opinion, achieves a perfect middle ground between DX and Runtime Performance. This is a hobby project to force me to relearn the semantics of Golang, and so I don't want to make my life harder by trying to be picky about server frameworks or writing high-level abstractions when Go's solutions are more than battle-proven.
  
> It also just happens to be about [as performant as Rust](https://youtu.be/Z0GX2mTUtfo) in a Server context.

## AutoBangumi 功能说明

- 简易单次配置就能持续使用
- 无需介入的 `RSS` 解析器，解析番组信息并且自动生成下载规则。
- 番剧文件整理:

    ```
    Bangumi
    ├── bangumi_A_title
    │   ├── Season 1
    │   │   ├── A S01E01.mp4
    │   │   ├── A S01E02.mp4
    │   │   ├── A S01E03.mp4
    │   │   └── A S01E04.mp4
    │   └── Season 2
    │       ├── A S02E01.mp4
    │       ├── A S02E02.mp4
    │       ├── A S02E03.mp4
    │       └── A S02E04.mp4
    ├── bangumi_B_title
    │   └─── Season 1
    ```

- 全自动重命名，重命名后 99% 以上的番剧可以直接被媒体库软件直接刮削

    ```
  [Lilith-Raws] Kakkou no Iinazuke - 07 [Baha][WEB-DL][1080p][AVC AAC][CHT][MP4].mp4 
  >>
   Kakkou no Iinazuke S01E07.mp4
  ```

- 自定义重命名，可以根据上级文件夹对所有子文件重命名。
- 季中追番可以补全当季遗漏的所有剧集
- 高度可自定义的功能选项，可以针对不同媒体库软件微调
- 支持多种 RSS 站点，支持聚合 RSS 的解析。
- 无需维护完全无感使用
- 内置 TDMB 解析器，可以直接生成完整的 TMDB 格式的文件以及番剧信息。

## 相关群组

- 前端 Bug 反馈群：[Telegram](https://t.me/+yNisOnDGaX5jMTM9)

***计划开发的功能：***

- TBA

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=maikirakiwi/Goto_Bangumi&type=Date)](https://star-history.com/#maikirakiwi/Goto_Bangumi)

## 贡献

欢迎提供 ISSUE 或者 PR。

贡献者名单请见：

<a href="https://github.com/maikirakiwi/Goto_Bangumi/graphs/contributors"><img src="https://contrib.rocks/image?repo=maikirakiwi/Goto_Bangumi"></a>


## Licenses

AutoBangumi & GotoBangumi are licensed under the [MIT License](https://github.com/EstrellaXD/Auto_Bangumi/blob/main/LICENSE)

[mikan]: https://mikanani.me
[plex]: https://plex.tv
[jellyfin]: https://jellyfin.org
