# 使用 mainline 升级内核


使用以下命令安装：

```console
sudo add-apt-repository ppa:cappelikan/ppa
sudo apt update
sudo apt install mainline
```

list 子命令会显示出来当前系统支持升级的或者已经安装的系统：

```console
mainline --list
```

更新内核支持 [Tetragon](https://github.com/cilium/tetragon):
```console
sudo mainline --install 5.5.19
```
