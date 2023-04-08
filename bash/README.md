# 常用脚本

## 查看文件夹

```bash
du -sh * | sort -h
```


## 禁用 IP

```
iptables -A OUTPUT -d 192.168.0.0/16 -j DROP
```