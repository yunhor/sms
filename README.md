## 阿里短消息服务

### 由来

   - 阿里大于合并至阿里云后，摇身一变成了阿里云通信，多了许多新功能，有了新版的api,虽然旧版的仍旧可用。但毕竟不是长久之计。  
   - 官方没有golang sdk
   - 已有好几个go库，但都几乎只实现发送的功能。

### feature
   - golang
   - 现版本基于阿里云协议2017-05-25版
   - [ ] 客户端调用支持命令行与toml参数配置
   - [ ] 短信发送
   - [ ] 验证码接口   
   - [ ] 群发接口
   - [ ] 发送反馈查询接口
   - [ ] 其它业务查询接口
   - [ ] 语音功能接口
   - [ ] 流量功能接口
   - [ ] 物联网功能接口
   - [ ] 私密小号功能接口

### 参考
 
   - [阿里云官网](https://dayu.aliyun.com/?spm=a3142.10677814.0.0.23716ebcliNC2w)

     - [阿里大于api](http://open.taobao.com/doc2/apiDetail.htm?spm=0.0.0.0.bkKKhG&apiId=25450)
     - [云通信短信api](https://help.aliyun.com/document_detail/56189.html?spm=a2c4g.11186623.6.580.bDKh92)

   - [holdno/alidayu](https://github.com/holdno/alidayu)
   - [ltt1987/alidayu](https://github.com/ltt1987/alidayu)
   - [gwpp/alidayu-go](https://github.com/gwpp/alidayu-go)

### 安装

    go get github.com/yunhor/alisms

### 示例

#### 更多例子详见_examples目录