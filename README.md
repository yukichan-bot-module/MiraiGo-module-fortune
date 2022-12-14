# MiraiGo-module-fortune

ID: `com.aimerneige.fortune`

Module for [MiraiGo-Template](https://github.com/Logiase/MiraiGo-Template)

## 功能

- 接收到 “求签 <求签内容>” 时，进行一次求签，并返回求签结果。

## 使用方法

在适当位置引用本包

```go
package example

imports (
    // ...

    _ "github.com/yukichan-bot-module/MiraiGo-module-fortune"

    // ...
)

// ...
```

在全局配置文件中写入配置

```yaml
aimerneige:
  fortune:
    blacklist: # 黑名单用户
      - 1781924496
      - 2802340025
    disallowed: # 关闭功能的群
      - 546362685
```

## LICENSE

<a href="https://www.gnu.org/licenses/agpl-3.0.en.html">
<img src="https://www.gnu.org/graphics/agplv3-155x51.png">
</a>

本项目使用 `AGPLv3` 协议开源，您可以在 [GitHub](https://github.com/yukichan-bot-module/MiraiGo-module-fortune) 获取本项目源代码。为了整个社区的良性发展，我们强烈建议您做到以下几点：

- **间接接触（包括但不限于使用 `Http API` 或 跨进程技术）到本项目的软件使用 `AGPLv3` 开源**
- **不鼓励，不支持一切商业使用**
