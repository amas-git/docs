---
title: Imagemagic Play
tags:
---
<!-- toc -->
# ImageMagic
## ImageMagic Convert
### -resize
```sh
# 绝对尺寸
$ convert -resize WxH
# 百分比
$ convert -resize N%
# 最大面积
$ convert -resize N@
```


### -filter
```sh
$ convert -list filter
Bartlett
Blackman
Bohman
Box
Catrom
Cubic
Gaussian
Hamming
Hanning
Hermite
Jinc
Kaiser
Lagrange
Lanczos
Mitchell
Parzen
Point
Quadratic
Sinc
SincFast
Triangle
Tent
Welsh
`+``
*-filter test script:
```sh
#!/bin/zsh
set -A FX $(convert -list filter | xargs)
for filter in $FX
do
    convert xc:  -bordercolor black -border 1 -filter $filter   -resize 3000%      dot_resize_$filter.jpg
    convert xc:  -bordercolor black -border 1 -filter $filter   +distort SRT 30,0  dot_distort_$filter.jpg
done
```

### -resample
重新采样并不会改变图像尺寸,它改变的是图像的分辨率,以便能够在不同dpi的设备上精确显示图像.
只有少数图像格式中包含了Resolution信息:
 * JPEG
 * PNG
 * TIFF


```sh
$ identify -verbose x.jpg  | grep Resolution
  Resolution: 72x72
```
 * Resolution
   * DPI:: Dot Per Inch : 每平方英寸上有多少像素()
   * 比如: 2英寸x3英寸的图像, 在100dpi的设备上, 分辨率即为200x300

### -draw <string>
```
   point           x,y
   line            x0,y0 x1,y1
   rectangle       x0,y0 x1,y1
   roundRectangle  x0,y0 x1,y1 wc,hc
   arc             x0,y0 x1,y1 a0,a1
   ellipse         x0,y0 rx,ry a0,a1
   circle          x0,y0 x1,y1
   polyline        x0,y0  ...  xn,yn
   polygon         x0,y0  ...  xn,yn
   bezier          x0,y0  ...  xn,yn
   path            path specification
   image           operator x0,y0 w,h <filename>
   text            x0,y0 string
   gravity         NorthWest, North, NorthEast, West, Center,
                   East, SouthWest, South, or SouthEast
   rotate     degrees
   translate  dx,dy
   scale      sx,sy
   skewX      degrees
   skewY      degrees
   color  x0,y0 method
   matte  x0,y0 method
```

### -colorspace
```sh
# 将x.png改为灰度图
$ convert -colorspace GRAY x.png y.png
```

### Regular Hexagons
 * [怎样绘制正六边形](http://www.rdwarf.com/lerickson/hex/index.html)

```sh
#!sh
#!/bin/zsh
# 绘制正六边形
# $1: 画布大小
# $2: 输出文件名
zmodload zsh/mathfunc
D=$1
OUT=$2
tg30=$[sqrt(3)/3]
typeset -i A B C
# 斜边
C=$[$D/2]
# 短直角边
A=$[$C/2]
# 长直角边
B=$[$A/$tg30]
CH=$[2*$C]
CW=$[2*$B]
echo "A=$A, B=$B, C=$C"
                         # x,y
P1=0,$A                  # 0,A
P2=0,$[$A+$C]            # 0,A+C
P3=$B,$[2*$C]            # B,2C
P4=$[2*$B],$[$A+$C]      # 2B,A+C
P5=$[2*$B],$A            # 2B,A
P6=$B,0                  # B,0
echo "convert -size  ${CW}x${CH}  xc:none -fill red -draw polygon $P1 $P2 $P3 $P4 $P5 $P6 $OUT"
convert -size ${CW}x${CH} xc:none -fill red -draw "polygon $P1 $P2 $P3 $P4 $P5 $P6" $OUT
```
使用:
```sh
$ hexagon 720 icon.png
```
### 铅笔特效
```sh
$ convert -charcoal 2 from.png to.png 
```

### 绘制带阴影的圆交矩形
```
#!sh
#!/bin/zsh
function main() {
    o_port=(-p 9999)
    o_root=(-r WWW)
    o_log=(-d ZWS.log)
    zparseopts -K -- p:=o_port r:=o_root l:=o_log h=o_help
    if [[ $? != 0 || "$o_help" != "" ]]; then
    fi
    port=$o_port[2]
    root=$o_root[2]
    log=$o_log[2]
    w=72
    h=72
    padding_l=2
    padding_t=2
    padding_r=2
    padding_b=2
    bg=DodgerBlue
    
    button_point_tl=${padding_l},${padding_t}
    button_point_br=$[$w-${padding_r}-1],$[$h-${padding_b}-1]
    
    round=5,5
    convert -size ${w}x${h} xc:none -fill $bg -draw "roundRectangle $button_point_tl $button_point_br $round" ( +clone -background black  -shadow 100x3+0+0 ) +swap -background none -layers merge +repage -resize ${w}x${h} x.png
}
main $*
```
## -shadow 
## 图像变换
### -blur 
 * radius
 * deviation
```sh
$ convert -blur 12x2 s.png t.png
```
### -charcoal
GrayscaleEffect, 素描效果
### -colorize
 * 色彩增强,默认为黑色, 你可以指定一个增量的百分比.
 * 增强颜色默认为黑色, 使用`-fill color`指定增量颜色

```sh
$ convert -colorize 10 s.png t.png
# 可以指定每个通道增强的百分比 
$ convert -colorize 20/20/20 s.png t.png
```
### -implode F
 * F > 0 : 内聚效果(哈哈镜: 中心凹镜)
 * F < 0 : 外爆效果(哈哈镜: 中心凸镜)

### -noise F
类似于SmoothOut

### -noise NoiseType
测试噪声类型:
```sh
#!/bin/zsh
IMG_SRC=$1
IMG_OUT=''
IMG_TYPE=${IMG_SRC##*.}
__NAME__=$0
function help {
    print "Useage:"
    print "    $ $__NAME__ source.png"
    exit $1
}
[ -z $IMG_SRC ] && help 1
NOISE_TYPE=($(convert -list noise | xargs))
for x in $NOISE_TYPE; do
    IMG_OUT=${$(basename xxx.png)%%.*}.$x.$IMG_TYPE
    convert +noise $x $IMG_SRC  $IMG_OUT
done
```

### -paint N
油画效果:: he paint effect simulates an oil painting by replacing the color of a pixel with the most common color in a circular area that is specified by the radius argument to the command. 

### -radial-blur N
 * N : 晃动角度
```
$ command -radial-blur 45 src.png out.png
```

### -raise NxM
 * N: 突起宽度
 * M: 突起高度
 * -raise
 * +raise
### -sepia-tone N%
褐色效果

### -shade NxM
 * N : 光源角度(azimuth, 北:0, 东:90)
 * M : elevation(海拔,高度)
阴影效果

### -sharpen NxM
与-blur相反的操作,可以将色彩边界变得更加明显
 * N: radius
 * M: deviation

### -solarize
曝光,负片效果.

### -spread N
### -swirl N
图像中心扭曲
 * N : 扭曲的角度

### -threshold 
限制各个通道的最大值.

### -unsharp radius x standard deviation + amount + threshold
### -wave amplitude x frequency
### -rotate N
### [+|-]contrast
 * +: 增加对比度
 * -: 降低对比度

### -dither
Dithering reduces the number of colors in an image. The most common example in everyday
use is turning color images into strict black-and-white images for use in newspapers. Dithering
works in amonochrome context by determining the brightness of a given color and then using
the right frequency of black dots per area to imply that brightness. 

### -flip
倒影
### -flop
镜像

### -fill color -tint N%
The tint command-line option
will add the specified percentage of the current fill color to the image. Only the nonpure colors
in the image will be affected; 
 * 纯色不会被影响

### -normalize
Normalization is the process of improving the contrast in an image so that it uses all the available color range. 
### -shear
### -roll +X+Y
滚动图像
### [-+]append 1.png 2.png ... x.png out.png
### -average *.png out.png
### -flatten *.png out.png
### 替换图片中指定的颜色
```sh
# 将x.png中的所有黑色替换为白色
$ convert x.png -fill white -opaque black y.png
# 如果你并不想替换某中颜色，而是将某种颜色区域变为透明， 可以这样
$ convert x.png -transparent white y.png
# 如果你需要将某个颜色之外的其他颜色都变成另外的一种颜色, 可以使用+opaque
$ convert x.png  -fill white +opaque black  y.png
# 有时候你会发现需要模糊替换，某个颜色偏差范围之内的色彩需要被替换
$ convert x.png -fuzz 50% -fill white +opaque black  y.png
```
### 将图片平均分割
```sh
# 垂直均匀分割成4份
$ convert -crop 25%x100% +repage  x.jpg crop_%d.jpg
```

```
For some of its work, ImageMagick uses command-line tools called delegates to encode and
decode the image file in a format that ImageMagick can use.
|| Delegate Name || Desc || URL
|| bzlib || MIFF文件中使用Bzip压缩算法 || http://sources.redhat.com/bzip2/
|| DSP   || DisplayPostScript
|| FlashPIX || FlashPIX格式 || ftp://ftp.imagemagick.org/pub/ImageMagick/delegates/libfpx-1.2.0.9.tar.gz
|| FreeType || TrueType fonts || http://www.freetype.org
|| GhostPCL || PCL page description language || http://www.artifex.com/downloads/
|| Ghostscript || PostScript 或 PDF 文档 || http://www.cs.wisc.edu/~ghost/
|| Graphviz || 
|| JBIG || JBIG无损黑白压缩格式 || http://www.cl.cam.ac.uk/~xml25/jbigkit/
|| JPEG2000 || 下一代JPEG压缩标准 || 
|| LCMS || ICC CMS color management || http://www.littlecms.com/
|| PNG || ||http://www.libpng.org/pub/png/pngcode.html
|| TIFF || || http://www.libtiff.org
|| WMF || || http://sourceforge.net/projects/wvware/
|| zlib || || http://www.gzip.org/zlib/
看看都支持哪些delegate:
```

```sh
$ convert -list delegate
```
## Identify
图片文件信息查看器
```sh
$ identify x.jpg
01.jpg JPEG 650x506 650x506+0+0 8-bit DirectClass 42.5KB 0.000u 0:00.000
```
列出所有颜色:
```sh
$ identify -list color
```

## ImageMagic/ImageGemetry 
图像几何
```
||= 图像几何表示   =||= 说明 =||
||scale%            || Height and width both scaled by specified percentage. ||
||scale-x%xscale-y%	|| Height and width individually scaled by specified percentages. (Only one % symbol needed.) ||
||width             || Width given, height automagically selected to preserve aspect ratio. ||
||xheight           || Height given, width automagically selected to preserve aspect ratio. ||
||widthxheight      || Maximum values of height and width given, aspect ratio preserved. ||
||widthxheight^     || Minimum values of width and height given, aspect ratio preserved.||
||widthxheight!     || Width and height emphatically given, original aspect ratio ignored.||
||widthxheight>     || Change as per widthxheight but only if an image dimension exceeds a specified dimension.||
||widthxheight<     || Change dimensions only if both image dimensions exceed specified dimensions.||
||area@	            || Resize image to have specified area in pixels. Aspect ratio is preserved.||
||{size}{offset}	|| Specifying the offset (default is +0+0). Below, {size} refers to any of the forms above.||
||{size}{+-}x{+-}y  || Horizontal and vertical offsets x and y, specified in pixels. Signs are required for both. Offsets are affected by ‑gravity setting. Offsets are not affected by % or other size operators. ||
```
## Montage
生成照片墙
```
||= 参数      =||= =||
|| -geometry   ||   ||
|| -tile 列x行 ||   ||
||-adaptive-sharpen geometry ||	adaptively sharpen pixels; increase effect near edges
||-adjoin	                 || join images into a single multi-image file
||-affine matrix	         || affine transform matrix
||-alpha	                 || on, activate, off, deactivate, set, opaque, copy", transparent, extract, background, or shape the alpha channel
||-annotate geometry text	 || annotate the image with text
||-authenticate value	     || decrypt image with this password
||-auto-orient	automagically orient image
||-background color	         || 背景色 ||
||-blue-primary point	     || chromaticity blue primary point
||-blur geometry	         || 模糊处理,减少图像的噪声 reduce image noise and reduce detail levels
||-border geometry	surround image with a border of color
||-bordercolor color	border color
||-caption string	assign a caption to an image
||-channel type	apply option to select image channels
||-clone index	clone an image
||-coalesce	merge a sequence of images
||-colors value	preferred number of colors in the image
||-colorspace type	set image colorspace
||-comment string	annotate image with comment
||-compose operator	set image composite operator
||-composite	composite image
||-compress type	image compression type
||-crop geometry	preferred size and location of the cropped image
||-debug events	display copious debugging information
||-define format:option	define one or more image format options
||-density geometry	horizontal and vertical density of the image
||-depth value	image depth
||-display server	get image or font from this X server
||-dispose method	layer disposal method
||-dither method	apply error diffusion to image
||-draw string	annotate the image with a graphic primitive
||-duplicate count,indexes	duplicate an image one or more times
||-endian type	endianness (MSB or LSB) of the image
||-extent geometry	set the image size
||-extract geometry	extract area from image
||-fill color	color to use when filling a graphic primitive
||-filter type	use this filter when resizing an image
||-flatten	flatten a sequence of images
||-flip	flip image in the vertical direction
||-flop	flop image in the horizontal direction
||-font name	render text with this font
||-frame geometry	surround image with an ornamental border
||-gamma value	level of gamma correction
||-geometry geometry	preferred size or location of the image
||-gravity type	horizontal and vertical text placement
||-green-primary point	chromaticity green primary point
||-help	print program options
||-identify	identify the format and characteristics of the image
||-interlace type	type of image interlacing scheme
||-interpolate method	pixel color interpolation method
||-kerning value	the space between two characters
||-label string	assign a label to an image
||-limit type value	pixel cache resource limit
||-log format	format of debugging information
||-mattecolor color	frame color
||-mode type	framing style
||-monitor	monitor progress
||-monochrome	transform image to black and white
||-origin geometry	image origin
||-page geometry	size and location of an image canvas (setting)
||-pointsize value	font point size
||-polaroid angle	simulate a Polaroid picture
||-profile filename	add, delete, or apply an image profile
||-quality value	JPEG/MIFF/PNG compression level
||-quantize colorspace	reduce image colors in this colorspace
||-quiet	suppress all warning messages
||-red-primary point	chromaticity red primary point
||-regard-warnings	pay attention to warning messages.
||-repage geometry	size and location of an image canvas
||-resize geometry	resize the image
||-respect-parentheses	settings remain in effect until parenthesis boundary.
||-rotate degrees	apply Paeth rotation to the image
||-sampling-factor geometry	horizontal and vertical sampling factor
||-scenesrange	image scene range
||-seed value	seed a new sequence of pseudo-random numbers
||-shadow geometry	simulate an image shadow
||-size geometry	width and height of image
||-strip	strip image of all profiles and comments
||-stroke color	graphic primitive stroke color
||-synchronize	synchronize image to storage device
||-taint	mark the image as modified
||-texture filename	name of texture to tile onto the image background
||-tile filename	tile image when filling a graphic primitive
||-tile-offset geometry	set the image tile offset
||-title	decorate the montage image with a title
||-transform	affine transform image
||-transparent color	make this color transparent within the image
||-transpose	flip image in the vertical direction and rotate 90 degrees
||-transparent-color color	transparent color
||-treedepth value	color tree depth
||-trim	trim image edges
||-type type	image type
||-units type	the units of image resolution
||-unsharp geometry	sharpen the image
||-verbose	print detailed information about the image
||-version	print version information
||-view	FlashPix viewing transforms
||-virtual-pixel method	access method for pixels outside the boundaries of the image
||-white-point point	chromaticity white point
```

一个例子:
```sh
$ montage *.png all.png
$ montage *.jpg  -tile 1x -shadow -geometry +1+1 -background black thumb.jpg  
```
