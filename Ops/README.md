# Essential tools
yum install -y wget lrzsz
yum install -y gcc gcc-c++ make cmake g++ gdbm gdbm-devel 

rpm -ivh QConf-2.0.3-1.el7.x86_64.rpm 

# alias
echo "alias ll='ls -l'" >> ~/.bashrc
echo "alias la='ls -a'" >> ~/.bashrc
echo "alias cdcd='cd /data/codes/dev'" >> ~/.bashrc

hostnamectl set-hostname Uranus // change host name