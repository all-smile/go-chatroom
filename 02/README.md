# 操作步骤
1. 进入 01 目录下
2. 执行 go build -o server.exe .\src\chatroom\server, 在01目录下生成 server.exe
3. .\server.exe 启动服务端
4. 执行 go build -o client.exe .\src\chatroom\client, 在01目录下生成 client.exe
5. .\client.exe 启动客户端

***

> 02 展示了 代码 在 01 基础上，针对服务端的封装
> 面向对象思想，创建结构体，绑定方法，结构体之间相互调用。抽象层的概念，中转站，分发消息类型等