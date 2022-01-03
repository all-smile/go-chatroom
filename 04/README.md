# 操作步骤
1. 进入 01 目录下
2. 执行 go build -o server.exe .\src\chatroom\server, 在01目录下生成 server.exe
3. .\server.exe 启动服务端
4. 执行 go build -o client.exe .\src\chatroom\client, 在01目录下生成 client.exe
5. .\client.exe 启动客户端

***

> 04 展示了 代码 在 03 基础上，完成 redis 的连接，登录验证逻辑