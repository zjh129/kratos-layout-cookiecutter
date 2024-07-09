# 说明

由于官方kratos命令脚本 创建项目、创建应用仅仅只是将kratos-layout下载，任然存在Greeter等字眼，没有根据项目名动态生成目录或类，因此使用cookiecutter工具生成项目模板。

该项目模板是基于kratos-layout项目模板进行修改的，主要是将Greeter等字眼替换为项目名，同时将项目名、应用名作为目录名。

# 安装cookiecutter工具

### 下载miniconda
https://docs.anaconda.com/miniconda/

根据自己的系统下载对应的miniconda，安装

### 安装cookiecutter
使用超管权限安装cookiecutter
```shell
conda install cookiecutter
```

### 使用cookiecutter
```shell
### 使用gitee（推荐：速度快）
cookiecutter https://gitee.com/zhaojianhui/kratos-layout-cookiecutter.git -f
### 使用github
cookiecutter https://github.com/zjh129/kratos-layout-cookiecutter.git -f
```

# 创建应用和应用后任然需要根据kratos官方命令来生成代码，以下是官方常用命令
### 生成所有proto源码、wire等等
```bash
go generate ./...
```

### 运行程序
```bash
kratos run
```
