# -f 指定Dockerfile文件,默认为Dockerfile
# -t 镜像名:版本tag
# . 一定必须肯定务必加上
docker build . -t goweb_app



# -d 后台运行
# -p 映射端口:内部端口
# -name 容器名字
# gva-server:1.0为docker build时的-t的参数
#docker run -d -p 8888:8888 --name zinx-v1 zinx:1.0

# -it 以可交互模式运行并进入容器, 使用快捷键Ctrl + p + q即后台运行程序,Ctrl+c为退出容器
# -p 映射端口:内部端口
# -name 容器名字
# gva-server:1.0为docker build时的-t的参数
docker run -p 8888:8888 goweb_app
