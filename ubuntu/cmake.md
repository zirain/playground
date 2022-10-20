# Install CMake on ubuntu

```shell

wget https://github.com/Kitware/CMake/releases/download/v3.20.0/cmake-3.20.0.tar.gz
tar -zxvf cmake-3.20.0.tar.gz

cd cmake-3.20.0
sudo ./bootsrap # run `sudo apt-get install libssl-dev` to fix missing ssl
make 
sudo make install

cmake --version
```