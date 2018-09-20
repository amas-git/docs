---
title: 使用apkdiff来发现安装包体积为何变大
tags:
---
# APKDIFF
apkdiff可以帮助你回答以下问题:
    - 是什么导致本次release比上次release大了3M?
    - 某个第三方库升级了, 究竟做了什么改动?
    - 哪些类是特别复杂的?
    - 和上个relase相比, 哪些代码发生了改动?
    - 继承和包含这两种方式对代码体积是否有影响?

## 如何使用
   * 确保aapt在PATH路径下
```
$ java -jar apkdiff.jar
APKDIFF V1.1 @AMAS
apkdiff will diff the two apk files in the following dimensions:
 1. size
 2. manifest changes
 3. dex changes

EXAMPLE:
# check one apk files
$ java -jar apkdiff a.apk
# diff two apk file
$ java -jar apkdiff a.apk b.apk

OTHERS: 
 IF USE PROGUARD
 1. copy mapping file to the same directory of apk file
 2. rename the mapping file to <apk-file-name>.map
 e.g: a.apk -> a.map(this is proguard mapping file)
```

```
$ ls -l
-rw-r----- 1 amas amas 11530381 8月  18 17:01 appy4_70910_gp.apk
-rw-r----- 1 amas amas 14499475 8月  18 17:01 appy4_70910_gp.map <- PROGUARD FILE
-rw-r----- 1 amas amas 12303209 8月  18 17:01 appy4_71100_gp.apk
-rw-r----- 1 amas amas 15638465 8月  18 17:01 appy4_71100_gp.map <- PROGUARD FILE
$ java -jar apkdiff.jar appy4_71100_gp.apk appy4_70910_gp.apk
PKDIFF V1.1 @AMAS
# appy4_71100_gp.apk  VS.  appy4_70910_gp.apk
@size  12303209(+772828)  11530381

# TOP变化
[MOD:407142] /DEX
[NOT:     0] /LIB
[MOD:  2891] /ASSETS
[MOD:163733] /PNG
[MOD: 38449] /XML
[MOD:125017] /OTHER
GID_DEX/TOP:
  [MOD:366696] classes2.dex
  [MOD: 40446] classes.dex
GID_LIB/TOP:
GID_ASSETS/TOP:
  [MOD:  1477] assets/fonts/NR-ICONS.ttf
  [ADD:  1250] assets/playing.json
  [MOD:   152] assets/kfmt.dat
  [MOD:     5] assets/YTPlayerView-iframe-player.html
  [MOD:     4] assets/ytb_video.html
  [MOD:     3] assets/YTPlayerView-iframe-player-locker.html
GID_IMAGE/TOP:
  [ADD: 22329] res/drawable-xhdpi-v4/onews__interest_entertainment.jpg
  [ADD: 11145] res/drawable-xhdpi-v4/onews__interest_sports.jpg
  [ADD:  9576] res/drawable-xhdpi-v4/onews__interest_society.jpg
  [ADD:  9025] res/drawable-xhdpi-v4/onews__interest_politics.jpg
  [ADD:  8847] res/drawable-xhdpi-v4/onews__interest_business.jpg
  [ADD:  8166] res/drawable-xhdpi-v4/onews__interest_technology.jpg
  [ADD:  7573] res/drawable-xhdpi-v4/onews__sdk_update_icon.png
  [ADD:  7481] res/drawable-xhdpi-v4/onews__interest_the_washington_post.jpg
  [ADD:  6933] res/drawable-xhdpi-v4/onews__interest_reuters.jpg
  [ADD:  6851] res/drawable-xhdpi-v4/onews__interest_lifestyle.jpg
  [MOD:  5675] res/drawable-xhdpi-v4/onews__profile_mine_bg.png
  [MOD:  5192] res/drawable-xhdpi-v4/onews__sdk_dialog_real_content.png
  [MOD:  4788] res/drawable-xhdpi-v4/score_guide_top_bg.png
  [ADD:  4636] res/drawable-xhdpi-v4/onews__interest_the_hufffington_post.jpg
  [MOD:  4114] res/drawable-xhdpi-v4/onews__sdk_us_oem_splash_logo.png
  [ADD:  4068] res/drawable-xhdpi-v4/media_notification_logo.png
  [ADD:  2227] res/drawable-xhdpi-v4/media_noti_default.png
  [ADD:  1898] res/drawable-xxxhdpi-v4/design_ic_visibility_off.png
  [ADD:  1844] res/drawable-xhdpi-v4/onews_media_cd_shadow.png
  [ADD:  1806] res/drawable/notification_listen.png
GID_LAYOUT/TOP:
  [ADD:  1166] res/layout-v16/news_album_topic_item.xml
  [ADD:  1129] res/layout/news_album_topic_item.xml
  [ADD:  1073] res/layout-v17/onews__viewstub_splash_interest_guide_2.xml
  [ADD:  1057] res/layout-v16/onews__viewstub_splash_interest_guide_2.xml
  [ADD:  1026] res/layout-v16/audio_player_recycler_item.xml
  [ADD:  1026] res/layout-v16/remote_view_music_player.xml
  [ADD:   997] res/layout/remote_view_music_player.xml
  [ADD:   994] res/layout-v11/notification_template_big_media_narrow_custom.xml
  [ADD:   990] res/layout/onews__viewstub_splash_interest_guide_2.xml
  [ADD:   990] res/layout/audio_player_recycler_item.xml
  [ADD:   967] res/layout-v11/notification_template_big_media_custom.xml
  [ADD:   875] res/layout-v17/exo_playback_control_view.xml
  [ADD:   863] res/layout-v16/news_second_topic_item.xml
  [ADD:   840] res/layout/exo_playback_control_view.xml
  [ADD:   839] res/layout-v16/news_second_topic_no_img_item.xml
  [ADD:   825] res/layout/news_second_topic_item.xml
  [ADD:   815] res/layout-v16/onews__detail_header_title_view.xml
  [ADD:   806] res/layout-v17/onews__update_version_dialog.xml
  [ADD:   803] res/layout/news_second_topic_no_img_item.xml
  [ADD:   780] res/layout/onews__detail_header_title_view.xml
GID_OTHERS/TOP:
  [MOD: 86812] resources.arsc
  [MOD:  7120] META-INF/CERT.SF
  [MOD:  7117] META-INF/MANIFEST.MF
  [ADD:  1787] res/drawable-anydpi-v21/exo_edit_mode_logo.xml
  [ADD:   706] res/drawable-anydpi-v21/design_ic_visibility_off.xml
  [ADD:   533] res/drawable-v21/avd_hide_password_1.xml
  [ADD:   509] res/drawable-v21/avd_show_password_1.xml
  [ADD:   458] res/animator-v21/design_appbar_state_list_animator.xml
  [ADD:   418] res/drawable/audio_seek_bar_bg.xml
  [ADD:   403] res/drawable/mine_center_multi_top_bg.xml
  [ADD:   403] res/drawable/mine_center_multi_top_night_bg.xml
  [ADD:   392] res/drawable/audio_noti_default_bg.xml
  [ADD:   375] res/drawable/onews__detail_audio_btn_bg.xml
  [ADD:   373] res/drawable/night_onews__detail_audio_btn_bg.xml
  [ADD:   369] res/drawable/onews__detail_audio_btn_playing_bg.xml
  [MOD:   367] AndroidManifest.xml
  [ADD:   361] res/drawable-v21/design_password_eye.xml
  [ADD:   349] res/drawable-anydpi-v21/exo_controls_fastforward.xml
  [ADD:   349] res/drawable-anydpi-v21/exo_controls_pause.xml
  [ADD:   348] res/drawable-anydpi-v21/exo_controls_previous.xml


# AndroidManifest变化
[MOD] /manifest/@android:versionCode 71100 -> 70910
[MOD] /manifest/@android:versionName 7.11.0 -> 7.9.10
[ADD] /manifest/application/activity/com.cmcm.onews.ui.AudioPlayerActivity -
[ADD] /manifest/application/activity/com.cmcm.onews.ui.AudioPlayerActivity/@android:configChanges keyboard|navigation|uiMode
[ADD] /manifest/application/activity/com.cmcm.onews.ui.AudioPlayerActivity/@android:launchMode standard
[ADD] /manifest/application/activity/com.cmcm.onews.ui.AudioPlayerActivity/@android:screenOrientation portrait
[ADD] /manifest/application/activity/com.cmcm.onews.ui.AudioPlayerActivity/@android:theme for_day_onews_sdk_audio_player
[ADD] /manifest/application/activity/com.cmcm.onews.ui.AudioPlayerActivity/@android:windowSoftInputMode stateAlwaysHidden|adjustPan
[ADD] /manifest/application/receiver/com.cm.mediaplayer.MineMediaButtonReceiver -
[ADD] /manifest/application/receiver/intent-filter/action/android.intent.action.MEDIA_BUTTON -
[ADD] /manifest/application/receiver/intent-filter/action/com.mobilesrepublic.appy.notification.CLICK_AUDIO -
[ADD] /manifest/application/receiver/intent-filter/action/com.mobilesrepublic.appy.notification.MEDIA -
[ADD] /manifest/application/service/com.cm.mediaplayer.MediaAudioService -
[ADD] /manifest/application/service/com.cm.mediaplayer.MediaAudioService/@android:enabled true
[ADD] /manifest/application/service/com.cm.mediaplayer.MediaAudioService/@android:stopWithTask true
[ADD] /manifest/application/service/com.cmcm.onews.active.KeepLiveService -
[ADD] /manifest/application/service/com.cmcm.onews.active.KeepLiveService/@android:exported true
[ADD] /manifest/application/service/com.cmcm.onews.active.KeepLiveService/@android:permission com.mobilesrepublic.appy.permission.keeplive
[ADD] /manifest/application/service/intent-filter/action/android.intent.action.MEDIA_BUTTON -
[ADD] /manifest/meta-data/android.support.VERSION -
[ADD] /manifest/meta-data/android.support.VERSION/@android:value 25.3.1
[ADD] /manifest/permission/com.mobilesrepublic.appy.permission.keeplive -
[ADD] /manifest/permission/com.mobilesrepublic.appy.permission.keeplive/@android:protectionLevel 0

# 配置变化
[ADD] /anydpi-v21 anydpi-v21

# DEX变化
[ADD] com/google/android/exoplayer2 
  [ADD] com/google/android/exoplayer2/@size 499434
  [ADD] com/google/android/exoplayer2/@class-num 534
  [ADD: 12760] com/google/android/exoplayer2/ExoPlayerImplInternal
  [ADD: 11183] com/google/android/exoplayer2/extractor/mp4/FragmentedMp4Extractor
  [ADD:  9895] com/google/android/exoplayer2/extractor/mkv/MatroskaExtractor$InnerEbmlReaderOutput
  [ADD:  9462] com/google/android/exoplayer2/extractor/mp4/AtomParsers
  [ADD:  7733] com/google/android/exoplayer2/ui/PlaybackControlView
  [ADD:  7656] com/google/android/exoplayer2/video/MediaCodecVideoRenderer
  [ADD:  7472] com/google/android/exoplayer2/audio/AudioTrack
  [ADD:  7320] com/google/android/exoplayer2/mediacodec/MediaCodecRenderer
  [ADD:  6864] com/google/android/exoplayer2/source/hls/HlsSampleStreamWrapper
  [ADD:  6808] com/google/android/exoplayer2/text/ttml/TtmlDecoder
  [ADD:  6548] com/google/android/exoplayer2/text/dvb/DvbParser
  [ADD:  6520] com/google/android/exoplayer2/Format
  [ADD:  6172] com/google/android/exoplayer2/metadata/id3/Id3Decoder
  [ADD:  5993] com/google/android/exoplayer2/extractor/mp4/MetadataUtil
  [ADD:  5788] com/google/android/exoplayer2/util/Util
  [ADD:  5725] com/google/android/exoplayer2/text/cea/Cea708Decoder
  [ADD:  5635] com/google/android/exoplayer2/ui/DefaultTimeBar
  [ADD:  5548] com/google/android/exoplayer2/util/ColorParser
  [ADD:  5541] com/google/android/exoplayer2/extractor/mkv/MatroskaExtractor
  [ADD:  5225] com/google/android/exoplayer2/mediacodec/MediaCodecUtil
  [ADD:  5051] com/google/android/exoplayer2/audio/MediaCodecAudioRenderer
  [ADD:  4808] com/google/android/exoplayer2/ui/SimpleExoPlayerView
  [ADD:  4748] com/google/android/exoplayer2/text/webvtt/WebvttCueParser
  [ADD:  4673] com/google/android/exoplayer2/source/hls/playlist/HlsPlaylistParser
  [ADD:  4632] com/google/android/exoplayer2/trackselection/DefaultTrackSelector
  [ADD:  4208] com/google/android/exoplayer2/text/cea/Cea608Decoder
  [ADD:  4193] com/google/android/exoplayer2/extractor/mp4/Mp4Extractor
  [ADD:  4062] com/google/android/exoplayer2/source/ExtractorMediaPeriod
  [ADD:  3863] com/google/android/exoplayer2/extractor/ts/H264Reader
  [ADD:  3802] com/google/android/exoplayer2/upstream/DefaultHttpDataSource
  [ADD:  3765] com/google/android/exoplayer2/util/ParsableByteArray
  [ADD:  3736] com/google/android/exoplayer2/extractor/ts/H265Reader
  [ADD:  3679] com/google/android/exoplayer2/source/hls/HlsMediaPeriod
  [ADD:  3628] com/google/android/exoplayer2/extractor/mp3/Mp3Extractor
  [ADD:  3546] com/google/android/exoplayer2/audio/Sonic
  [ADD:  3538] com/google/android/exoplayer2/extractor/DefaultTrackOutput
  [ADD:  3538] com/google/android/exoplayer2/ui/SubtitleView
  [ADD:  3379] com/google/android/exoplayer2/SimpleExoPlayer
  [ADD:  3370] com/google/android/exoplayer2/ExoPlayerImpl
  [ADD:  3272] com/google/android/exoplayer2/text/ttml/TtmlNode
  [ADD:  3225] com/google/android/exoplayer2/extractor/mp4/Atom
  [ADD:  3170] com/google/android/exoplayer2/text/cea/Cea708Decoder$CueBuilder
  [ADD:  3054] com/google/android/exoplayer2/extractor/ts/TsExtractor
  [ADD:  3045] com/google/android/exoplayer2/util/NalUnitUtil
  [ADD:  3020] com/google/android/exoplayer2/source/hls/HlsMediaChunk
  [ADD:  2996] com/google/android/exoplayer2/extractor/DefaultTrackOutput$InfoQueue
  [ADD:  2730] com/google/android/exoplayer2/source/hls/playlist/HlsPlaylistTracker$MediaPlaylistBundle
  [ADD:  2479] com/google/android/exoplayer2/text/webvtt/CssParser
  [ADD:  2455] com/google/android/exoplayer2/extractor/ogg/VorbisReader
  [ADD:  2355] com/google/android/exoplayer2/decoder/SimpleDecoder
  [ADD:  2299] com/google/android/exoplayer2/text/TextRenderer
  [ADD:  2217] com/google/android/exoplayer2/extractor/ts/AdtsReader
  [ADD:  2120] com/google/android/exoplayer2/extractor/ts/PsExtractor
  [ADD:  2057] com/google/android/exoplayer2/SimpleExoPlayer$ComponentListener
  [ADD:  2048] com/google/android/exoplayer2/extractor/DefaultExtractorInput
  [ADD:  2028] com/google/android/exoplayer2/text/tx3g/Tx3gDecoder
  [ADD:  2015] com/google/android/exoplayer2/extractor/ts/TsExtractor$PmtReader
  [ADD:  1949] com/google/android/exoplayer2/BaseRenderer
  [ADD:  1940] com/google/android/exoplayer2/extractor/flv/FlvExtractor
  [ADD:  1869] com/google/android/exoplayer2/metadata/scte35/SpliceInsertCommand
  [ADD:  1864] com/google/android/exoplayer2/extractor/ogg/DefaultOggSeeker
  [ADD:  1823] com/google/android/exoplayer2/extractor/ogg/VorbisUtil
  [ADD:  1788] com/google/android/exoplayer2/audio/Ac3Util
  [ADD:  1780] com/google/android/exoplayer2/extractor/ts/PesReader
  [ADD:  1779] com/google/android/exoplayer2/text/cea/Cea608Decoder$CueBuilder
  [ADD:  1730] com/google/android/exoplayer2/extractor/ts/H262Reader
  [ADD:  1726] com/google/android/exoplayer2/audio/SonicAudioProcessor
  [ADD:  1711] com/google/android/exoplayer2/extractor/ogg/StreamReader
  [ADD:  1678] com/google/android/exoplayer2/util/ParsableNalUnitBitArray
  [ADD:  1642] com/google/android/exoplayer2/extractor/mkv/MatroskaExtractor$Track
  [ADD:  1633] com/google/android/exoplayer2/upstream/Loader$LoadTask
  [ADD:  1617] com/google/android/exoplayer2/extractor/MpegAudioHeader
  [ADD:  1568] com/google/android/exoplayer2/source/hls/playlist/HlsPlaylistTracker
  [ADD:  1555] com/google/android/exoplayer2/DefaultRenderersFactory
  [ADD:  1552] com/google/android/exoplayer2/extractor/mkv/DefaultEbmlReader
  [ADD:  1541] com/google/android/exoplayer2/trackselection/MappingTrackSelector
  [ADD:  1499] com/google/android/exoplayer2/extractor/wav/WavExtractor
  [ADD:  1493] com/google/android/exoplayer2/util/ParsableBitArray
  [ADD:  1476] com/google/android/exoplayer2/ui/SubtitlePainter
  [ADD:  1470] com/google/android/exoplayer2/upstream/DefaultAllocator
  [ADD:  1461] com/google/android/exoplayer2/source/hls/HlsMediaSource
  [ADD:  1435] com/google/android/exoplayer2/text/cea/CeaDecoder
  [ADD:  1431] com/google/android/exoplayer2/metadata/MetadataRenderer
  [ADD:  1423] com/google/android/exoplayer2/upstream/DefaultBandwidthMeter
  [ADD:  1402] com/google/android/exoplayer2/video/DummySurface$DummySurfaceThread
  [ADD:  1381] com/google/android/exoplayer2/audio/ResamplingAudioProcessor
  [ADD:  1355] com/google/android/exoplayer2/util/MimeTypes
  [ADD:  1350] com/google/android/exoplayer2/DefaultLoadControl
  [ADD:  1348] com/google/android/exoplayer2/extractor/ts/DefaultTsPayloadReaderFactory
  [ADD:  1344] com/google/android/exoplayer2/extractor/ts/Ac3Reader
  [ADD:  1330] com/google/android/exoplayer2/ExoPlayerImplInternal$MediaPeriodHolder
  [ADD:  1330] com/google/android/exoplayer2/extractor/mp4/Sniffer
  [ADD:  1311] com/google/android/exoplayer2/extractor/ts/MpegAudioReader
  [ADD:  1304] com/google/android/exoplayer2/upstream/ContentDataSource
  [ADD:  1299] com/google/android/exoplayer2/source/AdaptiveMediaSourceEventListener$EventDispatcher
  [ADD:  1297] com/google/android/exoplayer2/mediacodec/MediaCodecInfo
  [ADD:  1290] com/google/android/exoplayer2/source/hls/WebvttExtractor
  [ADD:  1272] com/google/android/exoplayer2/trackselection/BaseTrackSelection
  [ADD:  1270] com/google/android/exoplayer2/ui/PlaybackControlView$ComponentListener
  [ADD:  1264] com/google/android/exoplayer2/audio/ChannelMappingAudioProcessor
  [ADD:  1239] com/google/android/exoplayer2/extractor/ogg/FlacReader
...
```
	
# APK == ZIP
```
$ file nr.zip
nr.apk: Zip archive data, at least v2.0 to extract
```

## ZIP MORE
```
$ zipinfo -t nr.zip
2995 files, 21292528 bytes uncompressed, 11023753 bytes compressed:  48.2%
```

## ZIP MORE MORE
```
$ zipinfo -1 nr.apk
AndroidManifest.xml
META-INF/CERT.RSA
META-INF/CERT.SF
META-INF/MANIFEST.MF
assets/YTPlayerView-iframe-player-locker.html
build-data.properties
classes.dex
classes2.dex
res/anim-v21/design_bottom_sheet_slide_in.xml
res/color/abc_tint_switch_track.xml
res/drawable-hdpi-v4/abc_ab_share_pack_mtrl_alpha.9.png
res/drawable-ldpi-v4/onews__logo_instanews.png
res/drawable-ldrtl-xxhdpi-v17/abc_ic_menu_cut_mtrl_alpha.png
res/drawable-ldrtl-xxxhdpi-v17/abc_ic_menu_copy_mtrl_am_alpha.png
res/drawable-mdpi-v4/abc_ab_share_pack_mtrl_alpha.9.png
res/drawable-tvdpi-v4/ic_plusone_medium_off_client.png
res/drawable-v21/abc_action_bar_item_background_material.xml
res/drawable-v23/abc_control_background_material.xml
res/drawable-xhdpi-v4/abc_ab_share_pack_mtrl_alpha.9.png
res/drawable-xxhdpi-v4/tw__ic_logo_default.png
res/drawable-xxxhdpi-v4/abc_text_select_handle_right_mtrl_light.png
res/drawable/abc_btn_borderless_material.xml
res/layout-sw600dp-v13/design_layout_snackbar.xml
res/layout-v16/album_news_bottom_item.xml
res/layout-v17/abc_dialog_title_material.xml
res/layout-v21/abc_screen_toolbar.xml
res/layout-v22/onews__notification_dialog.xml
res/layout/abc_select_dialog_material.xml
res/layout/brand_video_top_layout_s.xml
res/menu/debug_menu.xml
res/raw/gtm_analytics
res/xml/act_type.xml
resources.arsc
```
    * 这些目录的命名规则是怎样的? 

## ZIP MORE MORE MORE
```
$ zipinfo -m        \
appy4_70910_gp.apk  \ 
resources.arsc      \
'*.dex'             \ 
AndroidManifest.xml \ 
res/raw/notif.mp3   \ 
res/drawable/message_list_night_divider.xml \
assets/js/onews__jquery-2.1.4.min.js        \
'*.so'              \
res/drawable/orion_feedview_bg.png 

-rw-rw-rw-  2.3 unx    65988 b- 83% defN 80-000-00 00:00 AndroidManifest.xml
-rw----     2.4 fat    84349 b- 65% defN 80-000-00 00:00 assets/js/onews__jquery-2.1.4.min.js
-rw----     2.4 fat  8817176 b- 58% defN 80-000-00 00:00 classes.dex
-rw----     2.4 fat  3351756 b- 62% defN 80-000-00 00:00 classes2.dex
-rw----     2.4 fat    99548 b- 49% defN 80-000-00 00:00 lib/armeabi/libkcmutil.so
-rw-rw-rw-  2.3 unx      556 b- 57% defN 80-000-00 00:00 res/drawable/message_list_night_divider.xml
-rw-rw-rw-  2.3 unx      941 b-  0% stor 80-000-00 00:00 res/drawable/orion_feedview_bg.png
-rw-rw-rw-  2.3 unx    43784 b-  0% stor 80-000-00 00:00 res/raw/notif.mp3
-rw-rw-rw-  2.3 unx  2800600 b-  0% stor 80-000-00 00:00 resources.arsc
             +------------------------------------------: 
			     +--------------------------------------:
				        +-------------------------------:
						     +--------------------------: 文件类型
							          |                    b: binary
							          |                    t: text
							          +-----------------: 压缩方法
									                       stor : storing 无压缩
														   re:N : reducing
														         * re:1
                                                                 * re:2
                                                                 * re:3
                                                                 * re:4
														   shrk : shrinking
														         * shrk
														         * i4:2
																 * i8:3
                                                           defN : deflating 
                                                                 * defS : super fast
                                                                 * defF : Fast
	                                                             * defN : Normal
                                                                 * defX : Maximum Compression
                                                           tokn : tokenizing  (never  publicly released)

```
  * 不会被压缩的文件
    * jpg/png/gif/mp3
	* resources.arsc
  * 所有被压缩文件的最后修改时间都被丢弃

### WHY?
aapt只针对特定类型的文件才进行压缩?
```
$ android-7.1.0_r7/frameworks/base/tools/aapt/Package.cpp:
static const char* kExcludeExtension = ".EXCLUDE";

/* these formats are already compressed, or don't compress well */
static const char* kNoCompressExt[] = {
    ".jpg", ".jpeg", ".png", ".gif",
    ".wav", ".mp2", ".mp3", ".ogg", ".aac",
    ".mpg", ".mpeg", ".mid", ".midi", ".smf", ".jet",
    ".rtttl", ".imy", ".xmf", ".mp4", ".m4a",
    ".m4v", ".3gp", ".3gpp", ".3g2", ".3gpp2",
    ".amr", ".awb", ".wma", ".wmv", ".webm", ".mkv"
};
...
```

压缩后的文件大小?
```
-rw-rw-rw-  2.3 unx    65988 b- 83% defN 80-000-00 00:00 AndroidManifest.xml
                                 +---------: 压缩比=减小的体积/原始体积 * 100, 越大越好
$ zipinfo -l  appy4_70910_gp.apk AndroidManifest.xml
-rw-rw-rw-  2.3 unx    65988 b-    11560 defN 80-000-00 00:00 AndroidManifest.xml
	                     +-----------------: 压缩前的长度
						             +-----: 压缩后的长度
```

## 强制对文件进行压缩的后果?

## HOW TO REDUCE SIZE?
假如我们a.apk文件, 在没有源代码的情况下, 是否有可能减小它的体积?


```
$ unzip appy4_70910_gp.apk -d appy4_70910_gp
$ cd appy4_70910_gp
$ zip -r ../a.apk .
$ cd ../
$ zipinfo -t a.apk
3036 files, 21292528 bytes uncompressed, 8925937 bytes compressed:  58.1%
$ zipinfo -t appy4_70910_gp.apk
2995 files, 21292528 bytes uncompressed, 11023753 bytes compressed:  48.2%


-rw-rw-rw-  3.0 unx  2800600 bx 74% defN 79-Dec-31 00:00 resources.arsc
-rw-r--r--  3.0 unx  3351756 bx 61% defN 79-Dec-31 00:00 classes2.dex
-rw-r--r--  3.0 unx    99548 bx 49% defN 79-Dec-31 00:00 lib/armeabi/libkcmutil.so
-rw-r--r--  3.0 unx    84349 tx 65% defN 79-Dec-31 00:00 assets/js/onews__jquery-2.1.4.min.js
-rw-rw-rw-  3.0 unx      941 bx  6% defN 79-Dec-31 00:00 res/drawable/orion_feedview_bg.png
-rw-rw-rw-  3.0 unx      556 bx 57% defN 79-Dec-31 00:00 res/drawable/message_list_night_divider.xml
-rw-rw-rw-  3.0 unx    43784 bx  7% defN 79-Dec-31 00:00 res/raw/notif.mp3
-rw-rw-rw-  3.0 unx    65988 bx 82% defN 79-Dec-31 00:00 AndroidManifest.xml
-rw-r--r--  3.0 unx  8817176 bx 58% defN 79-Dec-31 00:00 classes.dex
```


初步性的胜利? 使用zip重新打包可以减小apk的体积 (11023753 - 8925937 = 2097816) SO EASY???
 * 为什么会变小?
 * 为什么apk里面的文件数会增多?

既然压缩方法是defN,如果换成更强劲的defX会怎样?
```
$ zip -1 -r ../b_9.apk .
$ cd ../
$ zipinfo -t b_9.apk
3035 files, 21292528 bytes uncompressed, 8880679 bytes compressed:  58.3
```
  * 比defN提高了: 8925937 - 8880679 = 45258
  * 看上去优化的空间并不大了

## 可以说我们确实让APK变得更小了么?
在用户下载一个APK的时候一般会采用压缩流的方式, 因此实际上10M的APK, 经过压缩流传输可能仅仅需要9M, 我们来大致模拟一下这个过程:
```
$ zip appy4_70910_gp.apk.zip appy4_70910_gp.apk
$ zip b.apk.zip b.apk
$ zip b_9.apk.zip b_9.apk  
$ ls -alr
-rw-r--r-- 1 amas amas  8879916 8月  17 19:50 appy4_70910_gp.apk.zip
-rw-r--r-- 1 amas amas  8926748 8月  17 19:53 b_9.apk.zip
-rw-r--r-- 1 amas amas  8971605 8月  17 19:52 b.apk.zip
-rw-r--r-- 1 amas amas  9543471 8月  17 19:23 b_9.apk
-rw-r--r-- 1 amas amas  9588729 8月  17 17:38 b.apk
-rw-r--r-- 1 amas amas 10419594 8月  17 19:33 b_1.apk
-rw-r----- 1 amas amas 11530381 8月  17 16:52 appy4_70910_gp.apk
```
  * 结果并不是我们想象的那样, 压缩算法最厉害的b_9.apk经过压缩之后并没有占到便宜, WHY?

## DEX FILE
```
$ unzip a.apk classes.dex
$ file classes.dex
classes.dex: Dalvik dex file version 035
$ dexdump -h classes.dex
Processing 'classes2.dex'...
Opened 'classes2.dex', DEX version '035'
Class #0 header:
class_idx           : 399
access_flags        : 0 (0x0000)
superclass_idx      : 4005
interfaces_off      : 2105068 (0x201eec)
source_file_idx     : 8634
annotations_off     : 2038340 (0x1f1a44)
class_data_off      : 3194140 (0x30bd1c)
static_fields_size  : 0
instance_fields_size: 1
direct_methods_size : 1
virtual_methods_size: 1

Class #0            -
  Class descriptor  : 'Lcom/facebook/share/internal/LikeActionController$10;'
  Access flags      : 0x0000 ()
  Superclass        : 'Ljava/lang/Object;'
  Interfaces        -
    #0              : 'Lcom/facebook/internal/PlatformServiceClient$CompletedListener;'
  Static fields     -
  Instance fields   -
    #0              : (in Lcom/facebook/share/internal/LikeActionController$10;)
      name          : 'this$0'
      type          : 'Lcom/facebook/share/internal/LikeActionController;'
      access        : 0x1010 (FINAL SYNTHETIC)
  Direct methods    -
    #0              : (in Lcom/facebook/share/internal/LikeActionController$10;)
      name          : '<init>'
      type          : '(Lcom/facebook/share/internal/LikeActionController;)V'
      access        : 0x10000 (CONSTRUCTOR)
      code          -
      registers     : 2
      ins           : 2
      outs          : 1
      insns size    : 6 16-bit code units
      catches       : (none)
      positions     : 
        0x0000 line=1139
      locals        : 
        0x0000 - 0x0006 reg=0 this Lcom/facebook/share/internal/LikeActionController$10; 
  Virtual methods   -
    #0              : (in Lcom/facebook/share/internal/LikeActionController$10;)
      name          : 'completed'
      type          : '(Landroid/os/Bundle;)V'
      access        : 0x0001 (PUBLIC)
      code          -
      registers     : 9
      ins           : 2
      outs          : 7
      insns size    : 140 16-bit code units
      catches       : (none)
      positions     : 
        0x0000 line=1143
        0x000b line=1183

```
	
## AndroidManifest.xml
```
$ file AndroidManifest.xml
AndroidManifest.xml: data
```
	* Android Binary XML
	* 如何解析?
```
$ aapt d xmltree appy4_70910_gp.apk AndroidManifest.xml
N: android=http://schemas.android.com/apk/res/android
  E: manifest (line=2)
    A: android:versionCode(0x0101021b)=(type 0x10)0x114fe
    A: android:versionName(0x0101021c)="7.9.10" (Raw: "7.9.10")
    A: android:installLocation(0x010102b7)=(type 0x10)0x0
    A: package="com.mobilesrepublic.appy" (Raw: "com.mobilesrepublic.appy")
    A: platformBuildVersionCode=(type 0x10)0x18 (Raw: "24")
    A: platformBuildVersionName=(type 0x4)0x40e00000 (Raw: "7.0")
    E: uses-sdk (line=8)
      A: android:minSdkVersion(0x0101020c)=(type 0x10)0xf
      A: android:targetSdkVersion(0x01010270)=(type 0x10)0x17
...
    E: application (line=136)
      A: android:theme(0x01010000)=@0x7f0a013c
      A: android:label(0x01010001)=@0x7f0704d2
      A: android:icon(0x01010002)=@0x7f02016c
      A: android:name(0x01010003)="com.mobilesrepublic.appy.Application" (Raw: "com.mobilesrepublic.appy.Application")
      A: android:allowBackup(0x01010280)=(type 0x12)0x0
      A: android:hardwareAccelerated(0x010102d3)=(type 0x12)0xffffffff
      A: android:largeHeap(0x0101035a)=(type 0x12)0xffffffff
      A: android:supportsRtl(0x010103af)=(type 0x12)0xffffffff
      E: activity (line=145)
        A: android:theme(0x01010000)=@0x7f0a0042
        A: android:label(0x01010001)=@0x7f0700f9
        A: android:name(0x01010003)="com.mobilesrepublic.appy.InitActviity" (Raw: "com.mobilesrepublic.appy.InitActviity")
        A: android:configChanges(0x0101001f)=(type 0x11)0x5a4
        A: android:windowSoftInputMode(0x0101022b)=(type 0x11)0x22
        E: intent-filter (line=151)
          E: action (line=152)
            A: android:name(0x01010003)="android.intent.action.MAIN" (Raw: "android.intent.action.MAIN")
          E: category (line=154)
            A: android:name(0x01010003)="android.intent.category.LAUNCHER" (Raw: "android.intent.category.LAUNCHER")
          E: category (line=155)
            A: android:name(0x01010003)="android.intent.category.MULTIWINDOW_LAUNCHER" (Raw: "android.intent.category.MULTIWINDOW_LAUNCHER")

...
```
   * 只有ResourceID没有ResourceName
   * TypedValue

android-7.1.0_r7/frameworks/base/core/res/res/values/attrs_manifest.xml
```
       {@link android.content.Intent#FLAG_ACTIVITY_MULTIPLE_TASK}. -->
    <attr name="launchMode">
        <!-- The default mode, which will usually create a new instance of
             the activity when it is started, though this behavior may change
             with the introduction of other options such as
             {@link android.content.Intent#FLAG_ACTIVITY_NEW_TASK
             Intent.FLAG_ACTIVITY_NEW_TASK}. -->
        <enum name="standard" value="0" />
        <!-- If, when starting the activity, there is already an
            instance of the same activity class in the foreground that is
            interacting with the user, then
            re-use that instance.  This existing instance will receive a call to
            {@link android.app.Activity#onNewIntent Activity.onNewIntent()} with
            the new Intent that is being started. -->
        <enum name="singleTop" value="1" />
        <!-- If, when starting the activity, there is already a task running
            that starts with this activity, then instead of starting a new
            instance the current task is brought to the front.  The existing
            instance will receive a call to {@link android.app.Activity#onNewIntent
            Activity.onNewIntent()}
            with the new Intent that is being started, and with the
            {@link android.content.Intent#FLAG_ACTIVITY_BROUGHT_TO_FRONT
            Intent.FLAG_ACTIVITY_BROUGHT_TO_FRONT} flag set.  This is a superset
            of the singleTop mode, where if there is already an instance
            of the activity being started at the top of the stack, it will
            receive the Intent as described there (without the
            FLAG_ACTIVITY_BROUGHT_TO_FRONT flag set).  See the
            <a href="{@docRoot}guide/topics/fundamentals/tasks-and-back-stack.html">Tasks and Back
            Stack</a> document for more details about tasks.-->
        <enum name="singleTask" value="2" />
        <!-- Only allow one instance of this activity to ever be
            running.  This activity gets a unique task with only itself running
            in it; if it is ever launched again with the same Intent, then that
            task will be brought forward and its
            {@link android.app.Activity#onNewIntent Activity.onNewIntent()}
            method called.  If this
            activity tries to start a new activity, that new activity will be
            launched in a separate task.  See the
            <a href="{@docRoot}guide/topics/fundamentals/tasks-and-back-stack.html">Tasks and Back
            Stack</a> document for more details about tasks.-->
        <enum name="singleInstance" value="3" />
    </attr>

```

