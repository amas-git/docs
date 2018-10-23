# Nginx Cookbook

```
<section> {
 <directive> <parameters>;
}
```

- 注意每个directive必须以';'结束

```
server {
    listen 80;
}
```

```
events {
	worker_connections 512;
}
http {
 server {
 }
}

http {
 server {
 listen *:80;
 server_name "";
 root /usr/share/nginx/html;
 }
}
```



### 正则表达式: ~regex

## Directives

listen address[:port];
listen port;
listen unix:path;



```nginx
http {
 include /opt/local/etc/nginx/mime.types;
 default_type application/octet-stream;
 sendfile on;
 tcp_nopush on;
 tcp_nodelay on;
 keepalive_timeout 65;
 server_names_hash_max_size 1024;
    
 server {
     listen 80;
     return 444;
 }
 server {
     listen 80;
     server_name www.example.com;
     location / {
     	  try_files $uri $uri/ @mongrel;
     }
     
     location @mongrel {
         proxy_pass http://127.0.0.1:8080;
     }
  }
}
```



### server

```
server {
 listen *:80;
 listen *:443 ssl;
 root /usr/share/nginx/html;
 index maintenance.html index.html;
}
```



### server_name

### location

- location [modifier] uri {...}
  - modifier
    - = 严格相等的比较
    - ~  大小写敏感的正则
    - ~*  大小写不敏感的正则
    - ^~ 
- location @name {…}

```
# first, we enter through the root
location / {
 # then we find a most-specific substring
 # note that this is not a regular expression
 location ^~ /css {
 # here is the regular expression that then gets matched
 location ~* /css/.*\.css$ {
 }
 }
 
}



```

```
location / {
 root /var/www/html;
}
location /foobar/ {
 root /data;
}
```

```
location /gifs/ {
 alias /data/images/;
}
http://example.org/gifs/business_cat.gif -> /data/images/business_cat.gif
```

#### Named Location Blocks

```
# 这个东西暂时没什么用, 有些类似于函数, 后面我们在使用的时候可以引用这里
location @foobar {
 ...
}

# try_files会依次去找 
location / {
  try_files maintenance.html index.html 404.html=404 @foobar;
}

```

### SSL

```
server {
 listen 80;
 listen 443 ssl;
 server_name www.foobar.com;
 ssl_certificate www.foobar.com.crt;
 ssl_certificate_key www.foobar.com.key;
}
```



## 启动和停止服务

```
kill -HUP `cat /var/run/nginx.pid`
kill -QUIT `cat /var/run/nginx.pid`
```

- KILL Halts a stubborn process
- HUP Conguration reload
- USR1 Reopen the log les (useful for log rotation)
- USR2 Upgrade executable on the fly
- WINCH Gracefully shutdown worker processes
- TERM, INT Quick shutdown
- QUIT Graceful shutdown

### default_server: 处理没人管的请求

### user

```
user nobody nogroup;
```



### worker_processes

```
worker_processes 1;
```



### error_log

### pid

### use

### worker_connections

### include



## HTTP client directives

chunked_transfer_encoding

client_body_buffer_size

client_body_in_file_only

client_body_in_single_buffer

client_body_temp_path

client_body_timeout

client_header_buffer_size

client_header_timeout

client_max_body_size

keepalive_disable

keepalive_requests

keepalive_timeout

large_client_header_buffers

msie_padding

msie_refresh

## Sections

### events





## 代理与反向代理

- 代理: Forward Proxy或Proxy
- 反向代理: Reverse Proxy

二者的区别如下:

```
Proxy        : Computer -> [Proxy   ] -> [Internet]
Reverse Proxy: Computer -> [Internet] -> [Reverse Proxy] -> WebApplication
```





## 实战

### 检测配置文件

``` bash
$ nginx -t -c <path-to-nginx.conf>
```



