
# nunu 快速開發專案im

## 可優化項目

- 將Service, Dao接口改為interface
- 使用nunu create all object 指令時, Server的NewAPP方法應同步新增對應的handler, service, dao的依賴注入
- 使用nunu create all object 指令時, wire.go應同步新增handler, service, dao的依賴注入
- 使用nunu create all object 指令時, NewServerHTTP的引數新增對應handler的調用

