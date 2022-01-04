# 操作步骤
1. 启动 redis 数据库
2. 进入 01 目录下
3. 执行 go build -o server.exe .\src\chatroom\server, 在01目录下生成 server.exe
4. .\server.exe 启动服务端
5. 执行 go build -o client.exe .\src\chatroom\client, 在01目录下生成 client.exe
6. .\client.exe 启动客户端

***

> 06 展示了 代码 在 05 基础上，实现了退出登录，并通知客户端更新在线人员列表