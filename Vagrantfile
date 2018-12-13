# -*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.
Vagrant.configure("2") do |config|
  # The most common configuration options are documented and commented below.
  # For a complete reference, please see the online documentation at
  # https://docs.vagrantup.com.

  # Every Vagrant development environment requires a box. You can search for
  # boxes at https://vagrantcloud.com/search.
  config.vm.box = "ubuntu/xenial64"

  # Disable automatic box update checking. If you disable this, then
  # boxes will only be checked for updates when the user runs
  # `vagrant box outdated`. This is not recommended.
  # config.vm.box_check_update = false

  # Create a forwarded port mapping which allows access to a specific port
  # within the machine from a port on the host machine. In the example below,
  # accessing "localhost:8080" will access port 80 on the guest machine.
  # NOTE: This will enable public access to the opened port
  # config.vm.network "forwarded_port", guest: 80, host: 8080

  # Create a forwarded port mapping which allows access to a specific port
  # within the machine from a port on the host machine and only allow access
  # via 127.0.0.1 to disable public access
  config.vm.network "forwarded_port", guest: 80, host: 8080, host_ip: "127.0.0.1"
  config.vm.network "forwarded_port", guest: 42069, host: 42069, host_ip: "127.0.0.1"

  # Create a private network, which allows host-only access to the machine
  # using a specific IP.
  # config.vm.network "private_network", ip: "192.168.33.10"

  # Create a public network, which generally matched to bridged network.
  # Bridged networks make the machine appear as another physical device on
  # your network.
  # config.vm.network "public_network"

  # Share an additional folder to the guest VM. The first argument is
  # the path on the host to the actual folder. The second argument is
  # the path on the guest to mount the folder. And the optional third
  # argument is a set of non-required options.
  # config.vm.synced_folder "../data", "/vagrant_data"
  config.vm.synced_folder "frontend/", "/home/vagrant/frontend"
  config.vm.synced_folder "backend/", "/home/vagrant/backend"

  # Provider-specific configuration so you can fine-tune various
  # backing providers for Vagrant. These expose provider-specific options.
  # Example for VirtualBox:
  #
  config.vm.provider "virtualbox" do |vb|
      vb.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
  #   # Display the VirtualBox GUI when booting the machine
  #   vb.gui = true
  #
  #   # Customize the amount of memory on the VM:
  #   vb.memory = "1024"
  end

  # run bash commands for more provisioning
  config.vm.provision "shell", inline: <<-SHELL
    echo "apt packages"
    apt-get update -y
    apt-get install -y apache2 httpie gcc
    # check if go is installed because it takes 12 billion years to download
    echo "Checking if go is installed"
    if [ ! -d /usr/local/go ]
      then
        echo "Downloading Go..."
        #download using the actual tar
        curl -s -o out.tar.gz https://dl.google.com/go/go1.11.2.linux-amd64.tar.gz
        tar -C /usr/local -xzf out.tar.gz
        #DELET
        rm out.tar.gz
    fi
    echo "installing nodejs"
    curl -sL https://deb.nodesource.com/setup_11.x | sudo -E bash -
    apt-get install -y nodejs

    #install all npm stuff
    echo "npm install that good stuff"
    cd /home/vagrant/frontend
    npm install
    # you can't run the build script in an arbitrary react app, so yeah
    # build it up
    npm run build
    # copy the files over
    sudo cp -fr /home/vagrant/frontend/build/* /var/www/html
    # apparently vagrant is running at ~ so putting it back to not destroy anything
    cd ~

    
    # compile our binary
    export PATH=$PATH:/usr/local/go/bin
    # if you actually go into the box and try to use go, it won't work for some? reason. So this is a hack to fix that
    echo "export PATH=$PATH:/usr/local/go/bin" >> .bashrc
    go get "github.com/gin-gonic/gin"
    go get "github.com/mattn/go-sqlite3"
    # place binary in /usr/sbin
    go build -o "/usr/sbin/ittsbackend" "/home/vagrant/backend/ittsbackend.go"

    # copy the systemd unit file into its proper location
    sudo cp /home/vagrant/backend/ittsbackend.service /etc/systemd/system/ittsbackend.service

    # start the itts backend service
    sudo systemctl start ittsbackend.service

    echo "Finished provisioning"
  SHELL
end
