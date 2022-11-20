# -*- mode: ruby -*-
# vi: set ft=ruby :

class VagrantPlugins::ProviderVirtualBox::Action::Network
    def dhcp_server_matches_config?(dhcp_server, config)
      true
    end
end

$docker = <<-SCRIPT
docker compose -f /vagrant/docker-compose.yaml up
SCRIPT

Vagrant.configure("2") do |config|

    config.vm.define :dockervm do |dockervm|
  
      dockervm.vm.box = "ubuntu/bionic64"
      dockervm.vm.hostname = "dockervm"
  
      dockervm.vm.network "private_network", ip: "192.168.33.2"
  
      dockervm.vm.synced_folder ".", "/vagrant", create: true
  
      dockervm.vm.provider "virtualbox" do |vb|
      dockervm.vagrant.plugins = "vagrant-docker-compose"
  
        vb.memory = "8192"
      end
  
      dockervm.vm.provision :docker_compose
      dockervm.vm.provision "docker" do |d|
        d.post_install_provision "shell", inline: $docker
      end
  
    end
  end