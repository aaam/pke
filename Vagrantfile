$enable_serial_logging = false

raise "vagrant-vbguest plugin must be installed" unless Vagrant.has_plugin? "vagrant-vbguest"

Vagrant.configure("2") do |config|
    # Sync time with the local host
    config.vm.provider 'virtualbox' do |vb|
        vb.customize [ "guestproperty", "set", :id, "/VirtualBox/GuestAdd/VBoxService/--timesync-set-threshold", 1000 ]
    end

    # sync build folder
    config.vm.synced_folder '.', '/vagrant', disabled: true
    config.vm.synced_folder 'scripts/vagrant/', '/scripts/', create: true
    config.vm.synced_folder 'build/', '/banzaicloud/', create: true

    $num_instances = 4

    # centos 7 nodes
    (1..$num_instances).each do |n|
        config.vm.define "centos#{n}" do |node|
            node.vm.box = "centos/7"
            node.vm.network "private_network", ip: "192.168.64.#{n+10}"
            node.vm.hostname = "centos#{n}"

            node.vm.provider "virtualbox" do |vb|
                vb.name = "centos#{n}"
                vb.memory = "2048"
                vb.cpus = "2"
                vb.customize ["modifyvm", :id, "--audio", "none"]
                vb.customize ["modifyvm", :id, "--memory", "2048"]
                vb.customize ["modifyvm", :id, "--cpus", "2"]
            end

            node.vm.provision "shell" do |s|
                s.inline = <<-SHELL
                yum update
                yum install -y ntp wget curl vim net-tools socat
                echo 'sync time'
                systemctl start ntpd
                systemctl enable ntpd

                echo 'set host name resolution'
                cat >> /etc/hosts <<EOF
192.168.64.11 centos1
192.168.64.12 centos2
192.168.64.13 centos3
192.168.64.14 centos4
EOF
                cat /etc/hosts

                hostnamectl set-hostname centos#{n}

                SHELL
            end
        end
    end

    # Ubuntu LTS nodes
    (1..$num_instances).each do |n|
        config.vm.define "ubuntu#{n}" do |node|
            node.vm.box = "ubuntu/bionic64"
            node.vm.network "private_network", ip: "192.168.64.#{n+20}"
            node.vm.hostname = "ubuntu#{n}"

            node.vm.provider "virtualbox" do |vb|
                vb.name = "ubuntu#{n}"
                vb.memory = "2048"
                vb.cpus = "2"
                vb.customize ["modifyvm", :id, "--audio", "none"]
                vb.customize ["modifyvm", :id, "--memory", "2048"]
                vb.customize ["modifyvm", :id, "--cpus", "2"]
            end

            node.vm.provision "shell" do |s|
                s.inline = <<-SHELL

                apt-get update
                apt-get install -y ntp wget curl vim net-tools socat
                echo 'sync time'
                systemctl start ntp
                systemctl enable ntp

                echo 'set host name resolution'
                cat >> /etc/hosts <<EOF
192.168.64.21 ubuntu1
192.168.64.22 ubuntu2
192.168.64.23 ubuntu3
192.168.64.24 ubuntu4
EOF
                cat /etc/hosts

                hostnamectl set-hostname ubuntu#{n}

                SHELL
            end
        end
    end
end
