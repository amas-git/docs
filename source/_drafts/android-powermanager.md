---
title: Android消耗电量是如何计算的
tags:
---
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

