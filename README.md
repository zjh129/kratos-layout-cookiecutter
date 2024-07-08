# 说明

由于官方kratos 创建项目无法做到足够灵活，创建项目或

# 安装cookiecutter工具

### 下载miniconda
https://docs.anaconda.com/miniconda/

根据自己的系统下载对应的miniconda，安装

### 安装cookiecutter
使用超管权限安装cookiecutter
```shell
conda install cookiecutter
```

### 使用cookiecutter创建项目
```shell
cookiecutter https://gitee.com/zhaojianhui/kratos-layout-cookiecutter.git --directory="project"
```

### 使用cookiecutter创建API应用
```shell
# 进入到项目目录
cd {you_project}
cookiecutter https://gitee.com/zhaojianhui/kratos-layout-cookiecutter.git --directory="appapi"
```