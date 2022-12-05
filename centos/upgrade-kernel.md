 yum install -y https://www.elrepo.org/elrepo-release-7.el7.elrepo.noarch.rpm
 yum  --disablerepo="*"  --enablerepo="elrepo-kernel"  list  available
 yum  --enablerepo=elrepo-kernel  install  -y  kernel-lt

# 查看启动顺序
 awk -F\' '$1=="menuentry " {print i++ " : " $2}' /boot/grub2/grub.cfg

# 设置启动顺序
 yum install -y grub2-pc
 grub2-set-default 0
 grub2-mkconfig -o /boot/grub2/grub.cfg

 uname -r

# 安装新版本工具包
yum --disablerepo=\* --enablerepo=elrepo-kernel install -y kernel-lt-tools.x86_64

 # 查看已安装内核
rpm -qa | grep kernel