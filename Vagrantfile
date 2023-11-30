# -*- mode: ruby -*-
# vi: set ft=ruby :
Vagrant.configure("2") do |config|
  config.vm.box = "spox/ubuntu-arm"
  config.vm.box_version = "1.0.0"
  # Disable automatic box update checking. If you disable this, then
  # boxes will only be checked for updates when the user runs
  # `vagrant box outdated`. This is not recommended.
  # config.vm.box_check_update = false

  # Create a forwarded port mapping which allows access to a specific port
  # within the machine from a port on the host machine. In the example below,
  # accessing "localhost:8080" will access port 80 on the guest machine.
  # NOTE: This will enable public access to the opened port
  config.vm.network "forwarded_port", guest: 80, host: 8080

  # Create a forwarded port mapping which allows access to a specific port
  # within the machine from a port on the host machine and only allow access
  # via 127.0.0.1 to disable public access
  # config.vm.network "forwarded_port", guest: 80, host: 8080, host_ip: "127.0.0.1"

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
  config.vm.synced_folder ".", "/vagrant_data"

  # Disable the default share of the current code directory. Doing this
  # provides improved isolation between the vagrant box and your host
  # by making sure your Vagrantfile isn't accessable to the vagrant box.
  # If you use this you may want to enable additional shared subfolders as
  # shown above.
  # config.vm.synced_folder ".", "/vagrant", disabled: true

  # Provider-specific configuration so you can fine-tune various
  # backing providers for Vagrant. These expose provider-specific options.
  # Example for VirtualBox:
  #
  # config.vm.provider "virtualbox" do |vb|
  #   # Display the VirtualBox GUI when booting the machine
  #   vb.gui = true
  #
  #   # Customize the amount of memory on the VM:
  #   vb.memory = "1024"
  # end
  #
  # View the documentation for the provider you are using for more
  # information on available options.

  # Enable provisioning with a shell script. Additional provisioners such as
  # Ansible, Chef, Docker, Puppet and Salt are also available. Please see the
  # documentation for more information about their specific syntax and use.
  config.vm.provision "shell", inline: <<-SHELL
    export DEBAIN_FRONTEND=noninteractive

    sudo systemctl stop unattended-upgrades
    sudo apt remove -y -q unattended-upgrades golang-go

    cd /vagrant_data
    sudo apt-get update -qq
    sudo apt-get -qq -y --no-install-recommends install sudo git make upx \
            gcc libc6-dev-amd64-cross \
            gcc-aarch64-linux-gnu \
            libc6-dev-arm64-cross \
            gcc-arm-linux-gnueabihf \
            libc6-dev-armhf-cross

    # Install Nodejs/NPM
    sudo apt-get install -q -y ca-certificates curl gnupg
    sudo mkdir -p /etc/apt/keyrings
    curl -fsSL https://deb.nodesource.com/gpgkey/nodesource-repo.gpg.key | sudo gpg --dearmor -o /etc/apt/keyrings/nodesource.gpg
    NODE_MAJOR=20
    echo "deb [signed-by=/etc/apt/keyrings/nodesource.gpg] https://deb.nodesource.com/node_$NODE_MAJOR.x nodistro main" | sudo tee /etc/apt/sources.list.d/nodesource.list
    sudo apt-get update -qq
    sudo apt-get install nodejs -y -q
    npm install -g yarn

    git submodule init #TODO: fix git submodule not pulling automatically

    # Install go
    whoami
    sudo rm -rf /usr/local/go
    curl -fsSL https://go.dev/dl/go1.21.4.linux-arm64.tar.gz | sudo tar -C /usr/local -xzf - 
    # sudo chmod +x /usr/local/go/bin/go
    export PATH=$PATH:/usr/local/go/bin
    go version

    # Install goreleaser
    echo 'deb [trusted=yes] https://repo.goreleaser.com/apt/ /' | sudo tee /etc/apt/sources.list.d/goreleaser.list
    sudo apt-get update -qq
    sudo apt-get install goreleaser -y -q

    # go install github.com/goreleaser/goreleaser
    # go install github.com/goreleaser/goreleaser@latest
    # make build-ui
    # make build-backend
    # goreleaser release --clean --snapshot

    # go mod tidy


    # export CGO_ENABLED=1
    # export CGO_LDFLAGS=-static
    
    # go build -o ./casa main.go
    # upx --lzma --best casa
  SHELL
end
