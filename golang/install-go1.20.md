sudo apt-get update  
sudo apt-get -y upgrade
wget  https://go.dev/dl/go1.20.2.linux-amd64.tar.gz
sudo tar -xvf go1.20.2.linux-amd64.tar.gz  
sudo mv go /usr/local  


.bashrc
export GOROOT=/usr/local/go 
export GOPATH=$HOME/Projects/Proj1 
export PATH=$GOPATH/bin:$GOROOT/bin:$PATH 

source ~/.bashrc


done ==> 
$ go version
$ go version go1.20.2 linux/amd64

