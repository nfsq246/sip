# sip

这是使用Go实现的RFC3261-SIP协议。

主要结构体为`Message`，对应一条完整的SDP信息。

### 使用方法

在Go中引用本库：

```go
import "github.com/nfsq246/sip"
```

生成一个消息对象，解析外部收到的字符串：

```go
sipMsg, err := sip.NewMessage(ioReader)
```

根据请求消息，生成一个应答消息对象：

```go
rspMsg = sip.NewResponse(sip.StatusXXX, rspMsg)
```

将消息对象转换为字符串用于传输：

```go
transferString := sipMsg.String()
```

### TODO

- [ ] 完善字段支持。

### 参考资料

* [RFC3261-IETF](https://tools.ietf.org/html/rfc3261)
* [RFC3261-中文版](https://www.docin.com/p-264204745.html)
