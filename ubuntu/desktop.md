# Ubuntu 安装桌面

安装 pixie cloud 需要使用 chrome


## 安装桌面

```
sudo apt update && sudo apt upgrade 
sudo apt install ubuntu-desktop lightdm
```


```
vim /usr/share/lightdm/lightdm.conf.d/50-ubuntu.conf

greeter-show-manual-login=true
```


## Cleanup

```
sudo apt remove ubuntu-desktop ubuntu-desktop* lightdm
sudo apt autoremove
```

参考：
- https://zhuanlan.zhihu.com/p/130812032