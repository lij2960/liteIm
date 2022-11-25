1. 运行模式配置
    ```
    优先使用环境变量
   设置环境变量的方法：export imRunMode="local"
   当环境变量不存在的时候，使用配置文件 config/config.ini 中的 mode 配置
   
   首次运行，需将 local.ini.example 重命名成 local.ini
   
   支持四种运行环境配置
   local：本地
   dev：开发
   test：测试
   prod：线上
    ```
2. 配置参数的读取
    ```
   读取通用配置参数：Config.Section("").Key("这里填写配置的key").String()
   读取不同环境的配置参数：Config.Section(Env).Key("这里填写配置的key").String()
   ```
3. 启动方法
   ```
   go run cmd/im/main.go
   ```
4. 自测方式
   ```
   unique_id为用户唯一ID，每个用户链接，可以指定一个唯一的用户ID
   socket链接地址：ws://127.0.0.1:8080/ws?unique_id=1
   
   http指定用户推送信息方式：
   post方式推送地址：http://127.0.0.1:8080/pushMsg
   body体参数：
   {
    "to_unique_id":"1",
    "from_unique_id":"2",
    "data":"test"
   }
   
   心跳参数示例：
   请求地址：ws://127.0.0.1:8080/ws?unique_id=1
   body参数：
   {
    "message_type":2,
    "data":{
        "message":"ping"
    }
   }
   
   socket消息发送示例：
   请求地址：ws://127.0.0.1:8080/ws?unique_id=1
   body参数：
   {
    "message_type":3,
    "data":{
        "to_unique_id":"2",
        "from_unique_id":"3",
        "message":"hello"
    }
   }
   ```
