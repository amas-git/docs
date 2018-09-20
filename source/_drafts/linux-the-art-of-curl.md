---
title: Curl脚本艺术(译) 
tags:
---
<!-- toc -->

# The Art of Scripting HTTP Request Using Curl
                 The Art Of Scripting HTTP Requests Using Curl
                 用Curl脚本化HTTP请求的艺术
                =============================================
 
本文档预设读者已经熟悉HTML与网络基本知识。
 
可以用脚本来进行操作，是一个好的计算机系统的基本条件。Unix之所以获得如此大的成功，原因之一就是它的功能能被脚本及各种各样运行自动化命令的工具扩展。
 
随着越来越多应用迁移到网上，"HTTP脚本"的需求愈加旺盛。当前，可以自动地从网络上提取信息，模仿用户行为，提交或者上传数据给服务器都是重要的任务。
Curl是用来进行各种URL操作与传输的命令行工具，但是本篇文档专注于如何用它来进行HTTP请求，为了乐趣等其他目的。这里预设您知道如何调用 'curl --help'或者'curl --manual' 来获取命令的基本信息。
Curl可以用来做任何事情，它可以发送请求，接收数据，发送数据及提取信息。你可能需要用某些脚本语言或者重复的手动调用来把这些事情胶合起来。

## 1. The HTTP Protocol
 
HTTP协议是用来从服务器获取信息的，它是基于TCP/IP的一个非常简单的协议。该协议也用一些方法来从客户端向服务器发送信息，如下文所述。
HTTP是从客户端向服务器发送的纯ASCII文本行，请求特写的操作，然后在真正请求的内容被发回客户端前，服务器回送一些行。
 
Curl可以作为客户端来发送HTTP请求。请求中包含一个方法（如，GET，POST，HEAD等），一些请求头，有时还会有一个请求实体。HTTP服务器响应，包括一个状态行（指示请求的处理结果），响应头及经常会有一个响应实体。“实体”部分是你所请求的数据，如，实际的HTML文件或者图片等。

使用curl的--verbose (-v短选项)会显示curl发给服务器的，以及其他信息文本。
当调试或者理解 curl<->server 间交互的时候，--verbose是最有用的选项。
 
有时--verbose也不够，此时--trace和--trace-ascii提供了更多的细节，因为它们显示了curl收发的所有细节。使用方式：

```
$ curl --trace-ascii debugdump.txt http://www.example.com/
```
## 2. URL
 
URL是用来定位互联网上的特定资源的寻址方法。你肯定知道这个，而且已经看过像`http://curl.haxx.se`或`​https://yourbank.com`这样的URL不下万次。
## 3. GET一个页面
 
用HTTP做的最简单最普遍的请求操作是获取一个页面。URL本身可以是一个页面、一个图片或者一个文件。客户端发送一个GET请求给服务器，并接收到所请求的文档。用如下命令来在终端上显示一个网页的内容：
 
```sh
$ curl http://curl.haxx.se
```
所有HTTP响应都包含一系列响应头，但它们通常被屏敝，用curl的`--include (-i)`选项来使它们同文档的其余部分一同显示。也可以只向服务器请求头，用`--head (-I)`选项（这样，curl只会发送一个HEAD请求）。

## 4. Forms
 
 Forms are the general way a web site can present a HTML page with fields for
 the user to enter data in, and then press some kind of 'OK' or 'submit'
 button to get that data sent to the server. The server then typically uses
 the posted data to decide how to act. Like using the entered words to search
 in a database, or to add the info in a bug track system, display the entered
 address on a map or using the info as a login-prompt verifying that the user
 is allowed to see what it is about to see.
 
 Of course there has to be some kind of program in the server end to receive
 the data you send. You cannot just invent something out of the air.
 
### 4.1 GET
 
  A GET-form uses the method GET, as specified in HTML like:
 
        <form method="GET" action="junk.cgi">
          <input type=text name="birthyear">
          <input type=submit name=press value="OK">
        </form>
 
  In your favorite browser, this form will appear with a text box to fill in
  and a press-button labeled "OK". If you fill in '1905' and press the OK
  button, your browser will then create a new URL to get for you. The URL will
  get "junk.cgi?birthyear=1905&press=OK" appended to the path part of the
  previous URL.
 
  If the original form was seen on the page "www.hotmail.com/when/birth.html",
  the second page you'll get will become
  "www.hotmail.com/when/junk.cgi?birthyear=1905&press=OK".
 
  Most search engines work this way.
 
  To make curl do the GET form post for you, just enter the expected created
  URL:
 
```sh
$ curl "http://www.hotmail.com/when/junk.cgi?birthyear=1905&press=OK"
```

### 4.2 POST
 
  The GET method makes all input field names get displayed in the URL field of
  your browser. That's generally a good thing when you want to be able to
  bookmark that page with your given data, but it is an obvious disadvantage
  if you entered secret information in one of the fields or if there are a
  large amount of fields creating a very long and unreadable URL.
 
  The HTTP protocol then offers the POST method. This way the client sends the
  data separated from the URL and thus you won't see any of it in the URL
  address field.
 
  The form would look very similar to the previous one:
 
        <form method="POST" action="junk.cgi">
          <input type=text name="birthyear">
          <input type=submit name=press value=" OK ">
        </form>
 
  And to use curl to post this form with the same data filled in as before, we
  could do it like:
 
```sh
$ curl --data "birthyear=1905&press=%20OK%20"         http://www.example.com/when.cgi
```

  This kind of POST will use the Content-Type
  application/x-www-form-urlencoded and is the most widely used POST kind.
 
  The data you send to the server MUST already be properly encoded, curl will
  not do that for you. For example, if you want the data to contain a space,
  you need to replace that space with %20 etc. Failing to comply with this
  will most likely cause your data to be received wrongly and messed up.
 
  Recent curl versions can in fact url-encode POST data for you, like this:
 
```sh
$ curl --data-urlencode "name=I am Daniel" http://www.example.com
```

### 4.3 File Upload POST
 
  Back in late 1995 they defined an additional way to post data over HTTP. It
  is documented in the RFC 1867, why this method sometimes is referred to as
  RFC1867-posting.
 
  This method is mainly designed to better support file uploads. A form that
  allows a user to upload a file could be written like this in HTML:
 
    <form method="POST" enctype='multipart/form-data' action="upload.cgi">
      <input type=file name=upload>
      <input type=submit name=press value="OK">
    </form>
 
  This clearly shows that the Content-Type about to be sent is
  multipart/form-data.
 
  To post to a form like this with curl, you enter a command line like:
 
```sh
$ curl --form upload=@localfilename --form press=OK [URL]
```

### 4.4 Hidden Fields
 
  A very common way for HTML based application to pass state information
  between pages is to add hidden fields to the forms. Hidden fields are
  already filled in, they aren't displayed to the user and they get passed
  along just as all the other fields.
 
  A similar example form with one visible field, one hidden field and one
  submit button could look like:
 
    <form method="POST" action="foobar.cgi">
      <input type=text name="birthyear">
      <input type=hidden name="person" value="daniel">
      <input type=submit name="press" value="OK">
    </form>
 
  To post this with curl, you won't have to think about if the fields are
  hidden or not. To curl they're all the same:
 
```sh
$ curl --data "birthyear=1905&press=OK&person=daniel" [URL]
```
 
### 4.5 Figure Out What A POST Looks Like
 
  When you're about fill in a form and send to a server by using curl instead
  of a browser, you're of course very interested in sending a POST exactly the
  way your browser does.
 
  An easy way to get to see this, is to save the HTML page with the form on
  your local disk, modify the 'method' to a GET, and press the submit button
  (you could also change the action URL if you want to).
 
  You will then clearly see the data get appended to the URL, separated with a
  '?'-letter as GET forms are supposed to.

## 5. PUT
 
也许向服务器上传数据的最好方式是使用PUT方法。当然，这需要有人在服务器端放一个程序或者脚本来提取HTTP PUT流。
使用PUT方法来向服务器放一个文件：
```
$ curl --upload-file uploadfile http://www.example.com/receive.cgi
```

## 6. HTTP认证
 
 HTTP Authentication is the ability to tell the server your username and
 password so that it can verify that you're allowed to do the request you're
 doing. The Basic authentication used in HTTP (which is the type curl uses by
 default) is *plain* *text* based, which means it sends username and password
 only slightly obfuscated, but still fully readable by anyone that sniffs on
 the network between you and the remote server.
 
HTTP认证就是告诉服务器你的用户名呵密码，使得服务器可以确认是否处理你的HTTP请求。
 To tell curl to use a user and password for authentication:
使用curl进行认证

```
 $ curl --user name:password http://www.example.com
```

 The site might require a different authentication method (check the headers
 returned by the server), and then --ntlm, --digest, --negotiate or even
 --anyauth might be options that suit you.
服务端可能需要不同的认证方式(看看服务器返回的HTTP报文中的Header), 然后你可以使用一下选项:
 * --ntim
 * --digest
 * --negotiate
 * --anyauth
 
 Sometimes your HTTP access is only available through the use of a HTTP
 proxy. This seems to be especially common at various companies. A HTTP proxy
 may require its own user and password to allow the client to get through to
 the Internet. To specify those with curl, run something like:
如果只能通过HttpProxy访问HTTP服务，HttpProxy也可能需要用户名和密码来确认你的身份。此时可以这样:

```
 $ curl --proxy-user proxyuser:proxypassword curl.haxx.se
```

 If your proxy requires the authentication to be done using the NTLM method,
 use --proxy-ntlm, if it requires Digest use --proxy-digest.
  * 如果HttpProxy需要使用[wiki:Ntlm NTLM]认证方式，则使用`--proxy-ntlm`选项，
  * 如果需要Digest, 则需要制定`--proxy-digest`选项
 
 If you use any one these user+password options but leave out the password
 part, curl will prompt for the password interactively.
只有用户名但没有密码，curl会提示你输入密码。
 
 Do note that when a program is run, its parameters might be possible to see
 when listing the running processes of the system. Thus, other users may be
 able to watch your passwords if you pass them as plain command line
 options. There are ways to circumvent this.

> 注意: 如果使用用户名+密码作为程序运行的参数，当你的程序/脚本运行时，操作系统上的其他用户时可以差看系统中进程的运行状况，比如PID, 执行参数啥的。因此，你的密码可能会被偷窥。

 
 It is worth noting that while this is how HTTP Authentication works, very
 many web sites will not use this concept when they provide logins etc. See
 the Web Login chapter further below for more details on that.

## 7. Referer
 
 A HTTP request may include a 'referer' field (yes it is misspelled), which
 can be used to tell from which URL the client got to this particular
 resource. Some programs/scripts check the referer field of requests to verify
 that this wasn't arriving from an external site or an unknown page. While
 this is a stupid way to check something so easily forged, many scripts still
 do it. Using curl, you can put anything you want in the referer-field and
 thus more easily be able to fool the server into serving your request.
 
HTTP请求可以包含一个`referer`域，用来指示从哪个URL来获取这一特定的资源。有些程序/脚本用请求的referer域来验证请求不是从外部网站或者一个未知页面来的。这样的作法很容易伪造，但很多脚本依然这样做。可以用curl在referer域放上想放的任何东西，更容易来欺骗服务器。
用curl来设置referer域：
 Use curl to set the referer field with:
 
```sh
$ curl --referer http://www.example.come http://www.example.com
```

## 8. 用户代理: User Agent
 
这个字段实际上是客户端(比如: 浏览器)的标识， 服务器端的应用程序会根据这个标识返回给客户端不同的页面。 头脑简单的Web程序员
试图通过这个字段为不同的浏览器返回不同的页面以达到更好的展示效果.
 
有时，你会发现通过curl获取的页面不太像是从浏览器看到的样子。 那么，是该使用UserAgent迷惑Server的时候了.
  
比如: 想伪装成Windows2000上的IE5, 你可以:
```sh
$ curl --user-agent "Mozilla/4.0 (compatible; MSIE 5.01; Windows NT 5.0)" [URL]
```
或者，一台老掉牙的Linux上的Netscape 4.73:
```sh
$ curl --user-agent "Mozilla/4.73 [en] (X11; U; Linux 2.2.15 i686)" [URL]
```

## 9. 重定向 : Redirects
 
 When a resource is requested from a server, the reply from the server may
 include a hint about where the browser should go next to find this page, or a
 new page keeping newly generated output. The header that tells the browser
 to redirect is Location:.
有时，从服务器请求资源，响应里可能返回给浏览器一个提示，告诉它从哪去找这个资源，或者一个新生成的页面。让浏览器去重定向的头是`Location:`。
 
 Curl does not follow Location: headers by default, but will simply display
 such pages in the same manner it display all HTTP replies. It does however
 feature an option that will make it attempt to follow the Location: pointers.
 
Curl默认不会跟着`Location:`头，但会像显示其他HTTP响应那样把页面显示出来。但是，有个选项可以让其尝试把`Location:`提示的资源打开。
 To tell curl to follow a Location:
让curl跟随`Location:`
 
```sh
$ curl --location http://www.example.com
```
 
 If you use curl to POST to a site that immediately redirects you to another
 page, you can safely use --location (-L) and --data/--form together. Curl will
 only use POST in the first request, and then revert to GET in the following
 operations.
 
如果用curl向一个会立即反你重定向到另一个页面的网站发POST，可以放心地将`--location (-L)`跟`--data`/`--form`选项一块使用。Curl会使用POST来发送第一个请求，然后转用GET来进行后续操作。

## 10. Cookies
 
 The way the web browsers do "client side state control" is by using
 cookies. Cookies are just names with associated contents. The cookies are
 sent to the client by the server. The server tells the client for what path
 and host name it wants the cookie sent back, and it also sends an expiration
 date and a few more properties.
 
“客户端状态控制”的方式是采用cookie。Cookie就是个键值对（关联数组），从服务器向客户端发送。服务器告诉客户端cookie要发送回的路径及主机名，同时也发送一个过期时间及其他一些属性。
 When a client communicates with a server with a name and path as previously
 specified in a received cookie, the client sends back the cookies and their
 contents to the server, unless of course they are expired.
 
当客户端与服务器用接收到的cookie里指定的主机名和路径通信时，客户端发回cookie及其内容，当然，过期的就不能发送了。
 Many applications and servers use this method to connect a series of requests
 into a single logical session. To be able to use curl in such occasions, we
 must be able to record and send back cookies the way the web application
 expects them. The same way browsers deal with them.
许多应用与服务器采用这种方式把一系列请求连接成单个逻辑上的会话（session）。在这些情况下使用curl，我们必须如服务器期望的那样记录并向发回cookie，要与浏览器的处理方式一致。
 
 The simplest way to send a few cookies to the server when getting a page with
 curl is to add them on the command line like:
 
用curl来获取页面时，向服务器发回cookie的最简单方法是在命令行上把cookie加上：
```sh
$ curl --cookie "name=Daniel" http://www.example.com
```

 Cookies are sent as common HTTP headers. This is practical as it allows curl
 to record cookies simply by recording headers. Record cookies with curl by
 using the --dump-header (-D) option like:
 
Cookie用普通HTTP头的方式发送。这是一个实用的方式，因为curl只要记录头就可以记录cookie。用`--dump-header (-D)`选项来记录cookie，命令如下：

```sh
$ curl --dump-header headers_and_cookies http://www.example.com
```
 (Take note that the --cookie-jar option described below is a better way to
 store cookies.)
（注意，下面介绍的`--cookie-jar`选项在存储cookie上更优）
 
 Curl has a full blown cookie parsing engine built-in that comes to use if you
 want to reconnect to a server and use cookies that were stored from a
 previous connection (or handicrafted manually to fool the server into
 believing you had a previous connection). To use previously stored cookies,
 you run curl like:
Curl有一个功能完备的内置的解析引擎，可以用来重新连接到服务器，或者使用从旧连接转存的cookie（或者手工打造一个cookie来欺骗服务器以为你已经连接过）。要使用存储的cookie，命令如下：
```sh 
$ curl --cookie stored_cookies_in_file http://www.example.com
```
 Curl's "cookie engine" gets enabled when you use the --cookie option. If you
 only want curl to understand received cookies, use --cookie with a file that
 doesn't exist. Example, if you want to let curl understand cookies from a
 page and follow a location (and thus possibly send back cookies it received),
 you can invoke it like:
 
当使用`--cookie`选项时，curl的“cookie引擎”就打开了。如果只是想让curl能够理解接收到的cookie，使用`--cookie`加一个不存在的文件作为参数。例如，如果你想让curl解析从页面发回的cookie并跟随一个位置（这样，可能发回它接收到的cookie），可以这样用：
```sh
$ curl --cookie nada --location http://www.example.com
```

貌似获取并存储cookie不能使用上面这条命令，而是`-c/--cookie-jar`
```sh
$ curl -c cookie.txt --location http://www.example.com
```

 Curl has the ability to read and write cookie files that use the same file
 format that Netscape and Mozilla do. It is a convenient way to share cookies
 between browsers and automatic scripts. The --cookie (-b) switch
 automatically detects if a given file is such a cookie file and parses it,
 and by using the --cookie-jar (-c) option you'll make curl write a new cookie
 file at the end of an operation:
Curl可以使用Netscap和Mozilla相同格式的文件来读写cookie文件。在浏览器间、自动化脚本间共享cookie是一个很方便的方式。`--cookie (-b)`选项自动检测后面的文件是否是这样的cookie并解析之；另，使用`--cookie-jar (-c)`选项，可以在操作结束后将cookie写入文件。
 
```sh
$ curl --cookie cookies.txt --cookie-jar newcookies.txt         http://www.example.com
```
 
## 11. HTTPS
 
 There are a few ways to do secure HTTP transfers. The by far most common
 protocol for doing this is what is generally known as HTTPS, HTTP over
 SSL. SSL encrypts all the data that is sent and received over the network and
 thus makes it harder for attackers to spy on sensitive information.
 
有好几种方式进行安全的HTTP传输。目前最普遍的协议是HTTPS，基于SSL的HTTP。用SSL对所有收、发的数据进行加密，使攻击者更难窥探敏感信息。
 SSL (or TLS as the latest version of the standard is called) offers a
 truckload of advanced features to allow all those encryptions and key
 infrastructure mechanisms encrypted HTTP requires.
 
SSL（或者最新版的叫TLS），提供了HTTP加密所需的先进的功能和关键的架构机制。
 Curl supports encrypted fetches thanks to the freely available OpenSSL
 libraries. To get a page from a HTTPS server, simply run curl like:
 
借助可以自由使用的OpenSSL库，curl支持加密的取操作。通过HTTPS服务器来获取一个页面，命令如下：
```sh
$ curl https://secure.example.com
```

## 11.1 证书Certificates
 
  In the HTTPS world, you use certificates to validate that you are the one
  you claim to be, as an addition to normal passwords. Curl supports client-
  side certificates. All certificates are locked with a pass phrase, which you
  need to enter before the certificate can be used by curl. The pass phrase
  can be specified on the command line or if not, entered interactively when
  curl queries for it. Use a certificate with curl on a HTTPS server like:
 
在HTTPS世界里，在普通的密码之外，我们使用证书来验证身份。Curl支持客户端证书。所有证书都用一个密码来保护，在curl使用证书之前必须输入密码。密码可以在命令行直接输入或者在交互模式下输入。让curl使用证书与HTTPS服务器交互如下：
```sh
$ curl --cert mycert.pem https://secure.example.com
```
 
  curl also tries to verify that the server is who it claims to be, by
  verifying the server's certificate against a locally stored CA cert
  bundle. Failing the verification will cause curl to deny the connection. You
  must then use --insecure (-k) in case you want to tell curl to ignore that
  the server can't be verified.
 
Curl同时会验证服务器的身份，通过将本地存储的CA证书与服务器的证书进行比较验证。验证失败curl会拒绝连接。如果想要curl忽略服务器验证，必须使用`--insecure (-k)`选项。
  More about server certificate verification and ca cert bundles can be read
  in the SSLCERTS document, available online here:
 
更多服务器证书验证和ca证书的信息，可以在SSLCERTS文档中读：
        http://curl.haxx.se/docs/sslcerts.html
## 12. 定制请求元素Custom Request Elements
 
 Doing fancy stuff, you may need to add or change elements of a single curl
 request.
 
为了做点酷的东西，有时可能需要添加或者改变一个curl请求里的一些元素。
 For example, you can change the POST request to a PROPFIND and send the data
 as "Content-Type: text/xml" (instead of the default Content-Type) like this:
 
比如，你可以把POST方法变为PROPFIND，然后用`Content-Type: text/xml`来发送数据（代替默认的Content-Type）：
```sh
$ curl --data "<xml>" --header "Content-Type: text/xml"  --request PROPFIND url.com
```
 
 You can delete a default header by providing one without content. Like you
 can ruin the request by chopping off the Host: header:
 
可以通过加一个没有值的header来删除一个默认的header。比如你可以用删掉Host:头的方式来干掉一个请求：
```sh
$ curl --header "Host:" http://www.example.com
```
 You can add headers the same way. Your server may want a "Destination:"
 header, and you can add it:
也可以用同样的方式来加header。服务器可能需要一个"Destination:"头，可以这么加上：
 
```sh
$ curl --header "Destination: http://nowhere" http://example.com
```
 
## 13. Web Login
 
 While not strictly just HTTP related, it still cause a lot of people problems
 so here's the executive run-down of how the vast majority of all login forms
 work and how to login to them using curl.
 
虽然跟HTTP不太相关，但登入确实造成不少麻烦。所以这里介绍了绝大多数的登入表单的原理及怎么用curl来实现登入。
 It can also be noted that to do this properly in an automated fashion, you
 will most certainly need to script things and do multiple curl invokes etc.
另外值得注意的是，要正确地自动化处理这些事务，需要把这些事务脚本化以实现多次调用。
 
 First, servers mostly use cookies to track the logged-in status of the
 client, so you will need to capture the cookies you receive in the
 responses. Then, many sites also set a special cookie on the login page (to
 make sure you got there through their login page) so you should make a habit
 of first getting the login-form page to capture the cookies set there.
 
首先，服务器通常使用cookie来记录登入的状态，所以有必要从响应里抓取cookie。然后，许多网站也会对登入页面设置一个特殊的cookie（以保证用户从登入页面进入那里），所以应该养成先获得登入表单页面再从那里抓取cookie。
 Some web-based login systems features various amounts of javascript, and
 sometimes they use such code to set or modify cookie contents. Possibly they
 do that to prevent programmed logins, like this manual describes how to...
 Anyway, if reading the code isn't enough to let you repeat the behavior
 manually, capturing the HTTP requests done by your browers and analyzing the
 sent cookies is usually a working method to work out how to shortcut the
 javascript need.
有些网站登入系统有许多javascript，并使用这些客户端语言来设置或修改cookie内容。这样做通常是为了防止登入机器人的操作，就像本文要做的那样。。。不管怎么样，如果阅读页面源码不足以让你手动实现其行为，那么，一个有效的方法就是把浏览器发的HTTP请求抓下来，分析浏览器发送的cookie。
 
 In the actual <form> tag for the login, lots of sites fill-in random/session
 or otherwise secretly generated hidden tags and you may need to first capture
 the HTML code for the login form and extract all the hidden fields to be able
 to do a proper login POST. Remember that the contents need to be URL encoded
 when sent in a normal POST.
 
在用来登入的实际的`<form>`标签里，很多网站插入了随机的session id或者悄悄生成隐藏的标签，所以有时需要抓取登入用的HTML代码先，提取出所有隐藏的域才能正确发送登入用的POST。需要注意的是，发送POST的时候，内容需要用URL编码方式来编码。

## 14. Debug
 
通常情况下，你会发现用curl与网站交互的结果与用浏览器产生的结果不太一致。
这时应该使curl请求更像浏览器：
  
 * 使用`--trace-ascii`选项来转存请求的详细日志，以便更好分析与理解。
 * 确保使用了必要的cookies（用`--cookie`来读，用`--cookie-jar`来写）
 * 把用户代理设为与最近流行的浏览器的用户代理
 * 像浏览器那样设置referer
 * 如果使用POST方法，确保像浏览器那样按一定的发送所有的域（参见4.5节）
 * LiveHTTPHeader是一个非常好的工具，可以查看所有用Mozilla/Firefox发送及接收的header（甚至在HTTPS协议下也可以）。
 
更底层的方式是用ethereal或tcpdump等工具抓HTTP包，来查看浏览器究竟收发了什么样的header（但在HTTPS下就行不通了）。

## 参考
 * RFC2616 is a must to read if you want in-depth understanding of the HTTP protocol.
 * RFC3986 explains the URL syntax.
 * RFC2109 defines how cookies are supposed to work.
 * RFC1867 defines the HTTP post upload format.
 * http://curl.haxx.se is the home of the cURL project8. User Agent
 * http://www.w3.org/Protocols/HTTP/HTRQ_Headers.html

