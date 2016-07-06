## 根据模板生成cloud-config文件(golang)
### 当前进展
1. 已经可以根据模板文件批量生成cloud-config文件
2. 给定mac地址，生成指定的cloud-config文件
3. http服务可以运行，但是还没有与生成cloud-config的程序集成

### TODO：
1. 优化生成cloud-config的程序：支持从github上直接获取模板文件、cluster配置信息
2. http服务集成cloud-config程序

### 今天遇到的问题：
1. yml parser的时候，大小写问题
2. golang的包的概念，目录组织结构等问题。（在go install的时候，浪费了一点时间）
