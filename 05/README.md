# 操作步骤
1. 进入 01 目录下
2. 执行 go build -o server.exe .\src\chatroom\server, 在01目录下生成 server.exe
3. .\server.exe 启动服务端
4. 执行 go build -o client.exe .\src\chatroom\client, 在01目录下生成 client.exe
5. .\client.exe 启动客户端

***

> 05 展示了 代码 在 04 基础上，增加用户注册
> 逻辑大致跟登录一样，只需要改改就可以了，会发现基础的东西写完之后，再写别的逻辑就像是被推着走