# eBPF


## Install BCC

https://github.com/iovisor/bcc/blob/master/INSTALL.md

```console
sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 4052245BD4284CDD
echo "deb https://repo.iovisor.org/apt/$(lsb_release -cs) $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/iovisor.list
sudo apt-get update
sudo apt-get install bcc-tools libbcc-examples linux-headers-$(uname -r)
```

switch to python2.7:
```
sudo update-alternatives --install /usr/bin/python python /usr/bin/python2.7  1
sudo update-alternatives --install /usr/bin/python python /usr/bin/python3.6  2

sudo update-alternatives --config python
```


## 参考资料

https://zhuanlan.zhihu.com/p/480811707