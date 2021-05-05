# Ansible

>
>
>- 如何通过条跳板机访问目的主机？
>- 如何登录宿主机的docker?
>- 如何让部署流程Idenpotent? 使用changed_when
>- 如何通过ansible操作aws?
>- 如何测试编写的ansible role?

## 安装

Ansible由python编写，所以只要装好python, 剩下的事情就靠pip搞定

```sh
$ pip install ansible
$ ansible --version
ansible 2.9.5
  config file = /etc/ansible/ansible.cfg
  configured module search path = ['/home/amas/.ansible/plugins/modules', '/usr/share/ansible/plugins/modules']
  ansible python module location = /usr/lib/python3.8/site-packages/ansible
  executable location = /usr/bin/ansible
  python version = 3.8.1 (default, Jan 22 2020, 06:38:00) [GCC 9.2.0
```



## Hello World

```
$ tree .
.
├── inventory.ini
└── main.yaml

```
inventory.ini:
```ini
[localhost]
127.0.0.1 ansible_connection=local
```

main.yaml:
```yaml
---
# This is a hello test
- hosts: localhost
  gather_facts: false
  
  tasks:
      - name: Get Current Date
        command: date
        register: var_date
        changed_when: false
      - name: Print var_date
        debug:
          msg: "{{ var_date.stdout }}"
```



```sh
$ ansible-playbook main.yaml
PLAY [localhost] ***************************************************************

TASK [Get Current Date] ********************************************************
ok: [localhost]

TASK [Print var_date] **********************************************************
ok: [localhost] => {
    "msg": "Wed 16 Sep 2020 07:02:12 PM CST"
}

PLAY RECAP *********************************************************************
localhost                  : ok=2    changed=0    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0
```



## Inventory

- invetory用于描述机器，以及ansible应该以怎样的方式操作这些机器.
- invetory文件可以是ini格式或yaml
- 你可以在`/etc/ansible/hosts`中配置 

```yaml
all:
  hosts:
    mail.example.com:
  children:
    webservers:
      hosts:
        foo.example.com:
        bar.example.com:
    dbservers:
      hosts:
        one.example.com:
        two.example.com:
        three.example.com:
```

ini文件也可以:

```ini
mail.example.com

[webservers]
foo.example.com
bar.example.com

[dbservers]
one.example.com
two.example.com
three.example.com
```

#### 分组

ansible将主机分组，任何一台机器必属于一组, 默认的两个分组是

- all
- ungrouped

```ini
all:
  hosts:
    mail.example.com:
  children:
    webservers:
      hosts:
        foo.example.com:
        bar.example.com:
    dbservers:
      hosts:
        one.example.com:
        two.example.com:
        three.example.com:
    east:
      hosts:
        foo.example.com:
        one.example.com:
        two.example.com:
    west:
      hosts:
        bar.example.com:
        three.example.com:
    prod:
      hosts:
        foo.example.com:
        one.example.com:
        two.example.com:
    test:
      hosts:
        bar.example.com:
        three.example.com:
```

````yaml
  hosts:
    jumper:
      ansible_port: 5555       # 
      ansible_host: 192.0.2.50 #
      ansible_user:            # 
      ansible_password:
      ansible_ssh_private_key_file:
      ansible_sftp_extra_args:
      ansible_scp_extra_args:
      ansible_ssh_extra_args:
      ansible_ssh_pipelining:
      ansible_docker_extra_args:
      ansible_ssh_executable: (2.2+)
      ansible_become:
      ansible_become_method:
      ansible_become_user:
      ansible_become_password:
      ansible_become_exe:
      ansible_become_flags:
      ansible_shell_type:
      ansible_python_interpreter:
      ansible_*_interpreter:
      ansible_shell_executable:
      ansible_connection: local | docker | ssh | paramiko | smart (默认)
      ansible_ssh_common_args:  '-o ProxyCommand="ssh -W %h:%p -q user@gateway.example.com"' # openssh 5.4+ , forward命令到指定主机，通常我们可以此解决跳板机问题
````



假设本地有vagrant虚拟机: 192.168.50.100, 我们定义`hello_inventory.ini`:

```ini
hello ansible_host=192.168.50.100 ansible_user=vagrant ansible_ssh_private_key_file=~/.vagrant.d/insecure_private_key
```

```bash
$ ansible hello -i hello_inventory.ini -m ping 
hello | SUCCESS => {
    "ansible_facts": {
        "discovered_interpreter_python": "/usr/bin/python"
    },
    "changed": false,
    "ping": "pong"
}
```



Tips:

- 如果机器实在太多， 可以使用DynamicInventory, 可以通过数据库生成json格式的invetory定义文件

- 多Inventories, Ansible可以将这些方法混合起来，需要用-i指定一个目录，这样ansible会将所有inventories合并起来

- 可以将invetories组织到不同的文件中

  

## Playbook

```sh
$ ansible-playbook -i <inventory> <playbook>
```



playbook是ansible是用ansible运行的基本单元，playbook用yaml编写，包含:

	- tasks
	- handlers
	- roles
	- 其他playbook

### Tasks

- cmd,command
- shell
- copy
- apt|yum|pacman
- package
- template

#### block

> block是一组tasks, 可以在block中对tasks出现的错误进行处理

```yaml
---
- hosts: all
    tasks:
    - block:
        - command: /bin/false
        - debug: msg="I will never run as the task above fails"
    rescue:
    	- debug: msg="This will run because the block failed"	
    always:
      - debug: msg="This runs no matter what happens"
```



### Roles

既然有了playbook为什么还要有roles呢？为了更好的模块化，模块是解决复杂性的好办法。

- 在Ansible Galaxy上存储了大量的roles
- 在https://galaxy.ansible.com/上

下载一个role:

```bash
$ ansible-galaxy search helloworld
$ ansible-galaxy install chusiang.helloworld
- downloading role 'helloworld', owned by chusiang
- downloading role from https://github.com/chusiang/helloworld.ansible.role/archive/master.tar.gz
- extracting chusiang.helloworld to /home/amas/.ansible/roles/chusiang.helloworld
- chusiang.helloworld (master) was installed successfully
$ tree /home/amas/.ansible/roles/chusiang.helloworld
/home/amas/.ansible/roles/chusiang.helloworld
├── defaults
│   └── main.yml
├── handlers
│   └── main.yml
├── Makefile
├── meta
│   └── main.yml
├── README.md
├── setup.yml
├── tasks
│   └── main.yml
├── tests
│   ├── inventory
│   └── test.yml
└── vars
    └── main.yml
    
# 自建一个roles
$ ansible-galaxy init hello
```

编写一个playbook, 引用这个role:

```yaml
---
- hosts: all
  roles:
    - chusiang.helloworld
```



### 创建role: ansible-galaxy init 

```bash
$ ansible-galaxy init hello
$ tree hello
hello
├── defaults
│   └── main.yml  # 配置默认变量
├── files         # 其他外部资源
├── handlers      # Handlers
│   └── main.yml 
├── meta          # 这个是给ansible galaxy看的
│   └── main.yml
├── README.md     # 说明书
├── tasks         # Tasks
│   └── main.yml
├── templates     # Jinjia2模板文件，需要拷贝到目标机器配的配置文件模板之类的
├── tests         # 给CI用的，当你修改了role,CI进行自动化测试
│   ├── inventory
│   └── test.yml
└── vars
    └── main.yml
```



根据不同的操作系统执行不同的安装:

```yaml
---
- include: install-debian.yml
  when: ansible_facts['os_family'] == "Debian"
```

但是这个做法比较重量，拿安装软件来说，不同平台的差异可能仅仅是包管理软件和包名有所不同

```yaml
---
- name: Install Apache
  apt: name={{ apache2_package_name }} state=installed
```

```yaml
---
apache2_package_name: httpd
```

```yaml
---
apache2_package_name: apache2
```

```yaml
# 
- name: Copy WordPress DB
  template: src=wp-database.sql dest=/tmp/wp-database.sql
  when: db_exist.rc > 0
```



打印debug信息

```yaml
tasks: print hello world
  - debug: msg="hello world"
  - debug: var=ansible_facts # 打印本机
```



在本机运行:

```yaml
---
  - name: "PLAY LOCALHOST"
    hosts: localhost
    connection: local 
    tasks:
    - name: "just execute a ls -lrt command"
      shell: "ls -lrt"
      register: "output"

    - debug: var=output.stdout_lines
```



## 主机信息收集

> Ansible会在运行时收集目标机器的信息，保存到ansible_facts中
>
> 这些信息可以在playbook中引用



### ansible_facts:

```bash
# 查看目标机器的环境信息
$ ansible -i <inventory> <target> -m setup
```



## 可配置化

这一章节主要讲ansible如何读取各种变量

> - vars
> - hostvars
> - groupvars
> - vars_prompt
> - ansible_facts
> - lookup()
> - ansible-playbook
>   - -e 'key=value
>   - -e <json>
>   - -e @<json-file>

### 向facts中添加信息

## 模块

> ansible的模块分为两组
>
> - ansible-modules-core
> - ansible-modules-extras

```bash
mkdir ansible-module
cd ansible-module
git clone git://github.com/ansible/ansible.git --recursive
source ansible/hacking/env-setup
chmod +x ansible/hacking/test-module
```

## 命令行技巧

```bash
# -m 调用模块
$ ansible web –i /path/to/inventory –m apt –a 'name=nginx state=installed'
# 以root身份运行
-b
--become-user
$ ansible web –i /path/to/inventory -b --become-user root –m apt –a 'name=nginx state=installed'

# 查看目标机器上的软件版本
$ ansible web -i inventory -m shell -a 'dpkg -s nginx | grep Version'

# 查看目标主机的信息
$ ansible web -i inventory -m setup -a 'gather_subset=hardware' 
```

## 安全

### ansible-vault

```zsh
# 加密
$ ansible-vault encrypt roles/mheap.demo/vars/main.yml

# 使用的时候，必须加上--ask-vault-pass
$ ansible-playbook -i inventory playbook.xml --ask-vault-pass

# 如果你执行vagrant provision, 会出现如下错误
ERROR! Attempting to decrypt but no vault secrets found
Ansible failed to complete successfully. Any error output should be
visible above. Please fix these errors and try again.

在Vagrantfile中ansible配置项加上:
ansible.ask_vault_pass = true

# 再次运行的时候，会提示你输入加密密码，然后才能继续

# 想编辑加密文件的内容？
$ ansible-vault edit <file>

# 查看加密文件的明文？
$ ansible-vault view <file>

# 重新设置密码
$ ansible-vault rekey <file>
```

## Ansible上的k8s插件

- k8s
- k8s_info
- k8s_scale
- k8s_exec
- k8s_service
- k8s_log
- geerlingguy.k8s

## 参考

- https://jinja.palletsprojects.com/en/2.10.x/templates/#builtin-filters