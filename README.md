# YouTube puh live
**ffmpeg 直播推流小工具**

循环不间断的推流视频文件到 YouTube 直播。

## 使用说明
先了解一下配置项：

```go
package main

const (
  
    // ffmpegDarwin MacOS 下 ffmpeg 的位置
    ffmpegDarwin = "./ffmpeg/ffmpeg"
    // Linux 同理
    ffmpegLinux = "./ffmpeg/ffmpeg"
    // Windows 同理
    ffmpegWindows = "./ffmpeg/bin/ffmpeg.exe"
)
```

**使用步骤**

1. [下载 `ffmpeg`](http://ffmpeg.org/download.html) ，然后会得到 `ffmpeg` 的执行文件。
2. 在源代码中根据配置项说明配置好自己需要格式转换的后缀，以及 `ffmpwg` 执行文件的相对位置，参考相对位置例如：
   ```
   ├── ffmepg-conv-tool // 本项目目录
   │ ├── ffmpeg         // ffmpeg 文件夹
   │ │ └── ffmpeg       // ffmpeg 执行文件
   │ ├── go.mod         // go mod 文件
   │ ├── main.exe       // Windows 下执行文件
   │ ├── main           // MacOS 或 Linux 下执行文件
   │ ├── main.go        // 全部的源码文件
   │ ├── url.txt        // 放置推流地址的文件
   │ └── videos/test.mp4       // 推流视频文件
   ```
3. 编译本项目（脚本），得到执行文件 `main`。
4. 确保相对位置放好后，执行文件 `main` 和待转换文件应该已在同一个目录下。
5. 然后双击 `main` 执行文件即可。

