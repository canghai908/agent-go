# Zabbix Agent-go

使用 https://github.com/fujiwara/go-zabbix-get 实现的一个支持被动检查的Agent，源码及打包好的Agent的rpm包     
RPM包：https://dl.cactifans.com/tools/agent-go-0.0.1-1.el7.x86_64.rpm  
模版：https://dl.cactifans.com/tools/template_linux_agent_go.tar.gz  
目前一共实现里6个监控指标的被动采集，数据采集使用[github.com/shirou/gopsutil](github.com/shirou/gopsutil) 

| ItemKey      |    说明 | 单位  |
| :-------- | :--------| :--: |
| system.info  | 操作系统信息 |    |
| agent.go.ping | 类似agent.ping检查agent状态 |    |
| cpu.model  | cpu类型|    |
| mem.total  | 总内存 |MB    |
| mem.used  | 已使用内存 | MB   |
| mem.usedper  | 内存使用率 | %   |

### 使用

目前只提供centos7_x64位rpm包，下载并安装RPM包
```
wget https://dl.cactifans.com/tools/agent-go-0.0.1-1.el7.x86_64.rpm
rpm -ivh agent-go-0.0.1-1.el7.x86_64.rpm
```
启动agent
```
systemctl start agent-go
```
查看Agent进程
```
ps -ef|grep agent-go
```
可以看到agent已经启动
```
root      5018     1  1 10:12 ?        00:00:00 /usr/bin/agent-go -c /etc/agent-go/agent-go.ini
```
Agent配置文件位于
```
/etc/agent-go/agent-go.ini
```
配置文件介绍
>Debug 日志级别1-7  
>ListenIP agent监听的ip，默认为0.0.0.0所有ip  
>ListenPort agent监听的port，默认为10049  
>LogFile 日志路径，默认为/var/log/agent-go.log  

Zabbix中导入模版，新建一个主机，并关联导入的模版，主机名可以随意，IP为部署agent的机器ip，端口为10499
![1](https://img.cactifans.com/wp-content/uploads/2019/03/4D5453E7-EC95-48C8-9C91-5DE47BD597AE-1024x430.jpg)
添加并关联模版。
![2](https://img.cactifans.com/wp-content/uploads/2019/03/78E7972B-23C0-44CC-96F3-56A1E4DD7F5F.jpg)
不出意外一会就可以看到主机的Agent状态已变为可用状态
![3](https://img.cactifans.com/wp-content/uploads/2019/03/B5DC034F-2378-4ABC-9A68-8746FB4853A8-1024x209.jpg)
并能在Last data里看到最新的数据
![4](https://www.cactifans.org/wp-content/uploads/2019/03/DE174094-827E-4F59-A4C4-4C7CCD8F5733-1024x420.jpg)