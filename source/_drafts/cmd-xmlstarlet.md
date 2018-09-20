---
title: Xmlstarlet使用手册
tags:
---
<!-- toc -->
# xmlstarlet 
基于命令行的Xml编辑工具, 常用于脚本化批量处理Xml.
## 如何获得帮助
```sh
$ man xmlstarlet
# 查看具体命令的使用方法
$ xmlstarlet COMMAND --help
```
## 命令 
### 显示结构: el
```
Usage: xml el [<options>] <xml-file>
where
  <xml-file> - input XML document file name (stdin is used if missing)
  <options> is one of:
  -a    - show attributes as well
  -v    - show attributes and their values
  -u    - print out sorted unique lines
  -d<n> - print out sorted unique lines up to depth <n>
```
### 查询: sel
```
Usage: xml sel <global-options> {<template>} [ <xml-file> ... ]
where
  <global-options> - global options for selecting
  <xml-file> - input XML document file name/uri (stdin is used if missing)
  <template> - template for querying XML document with following syntax:
<global-options> are:
  -C or --comp              - display generated XSLT
  -R or --root              - print root element <xsl-select>
  -T or --text              - output is text (default is XML)
  -I or --indent            - indent output
  -D or --xml-decl          - do not omit xml declaration line
  -B or --noblanks          - remove insignificant spaces from XML tree
  -E or --encode <encoding> - output in the given encoding (utf-8, unicode...)
  -N <name>=<value>         - predefine namespaces (name without 'xmlns:')
                              ex: xsql=urn:oracle-xsql
                              Multiple -N options are allowed.
  --net                     - allow fetch DTDs or entities over network
  --help                    - display help
Syntax for templates: -t|--template <options>
where <options>
  -c or --copy-of <xpath>   - print copy of XPATH expression
  -v or --value-of <xpath>  - print value of XPATH expression
  -o or --output <string>   - output string literal
  -n or --nl                - print new line
  -f or --inp-name          - print input file name (or URL)
  -m or --match <xpath>     - match XPATH expression
  -i or --if <test-xpath>   - check condition <xsl:if test="test-xpath">
  -e or --elem <name>       - print out element <xsl:element name="name">
  -a or --attr <name>       - add attribute <xsl:attribute name="name">
  -b or --break             - break nesting
  -s or --sort op xpath     - sort in order (used after -m) where
  op is X:Y:Z, 
      X is A - for order="ascending"
      X is D - for order="descending"
      Y is N - for data-type="numeric"
      Y is T - for data-type="text"
      Z is U - for case-order="upper-first"
      Z is L - for case-order="lower-first"
There can be multiple --match, --copy-of, --value-of, etc options
in a single template. The effect of applying command line templates
can be illustrated with the following XSLT analogue
xml sel -t -c "xpath0" -m "xpath1" -m "xpath2" -v "xpath3" \
        -t -m "xpath4" -c "xpath5"
is equivalent to applying the following XSLT
<?xml version="1.0"?>
<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">
<xsl:template match="/">
  <xsl:call-template name="t1"/>
  <xsl:call-template name="t2"/>
</xsl:template>
<xsl:template name="t1">
  <xsl:copy-of select="xpath0"/>
  <xsl:for-each select="xpath1">
    <xsl:for-each select="xpath2">
      <xsl:value-of select="xpath3"/>
    </xsl:for-each>
  </xsl:for-each>
</xsl:template>
<xsl:template name="t2">
  <xsl:for-each select="xpath4">
    <xsl:copy-of select="xpath5"/>
  </xsl:for-each>
</xsl:template>
</xsl:stylesheet>
XMLStarlet is a command line toolkit to query/edit/check/transform
XML documents (for more information see http://xmlstar.sourceforge.net/)
Current implementation uses libxslt from GNOME codebase as XSLT processor
(see http://xmlsoft.org/ for more details)
```
### 编辑: ed 
```
Usage: xml ed <global-options> {<action>} [ <xml-file-or-uri> ... ]
where
  <global-options>  - global options for editing
  <xml-file-or-uri> - input XML document file name/uri (stdin otherwise)
<global-options> are:
  -P (or --pf)        - preserve original formatting
  -S (or --ps)        - preserve non-significant spaces
  -O (or --omit-decl) - omit XML declaration (<?xml ...?>)
  -L (or --inplace)   - edit file inplace
  -N <name>=<value>   - predefine namespaces (name without 'xmlns:')
                        ex: xsql=urn:oracle-xsql
                        Multiple -N options are allowed.
                        -N options must be last global options.
  --help or -h        - display help
where <action>
  -d or --delete <xpath>
  -i or --insert <xpath> -t (--type) elem|text|attr -n <name> -v (--value) <value>
  -a or --append <xpath> -t (--type) elem|text|attr -n <name> -v (--value) <value>
  -s or --subnode <xpath> -t (--type) elem|text|attr -n <name> -v (--value) <value>
  -m or --move <xpath1> <xpath2>
  -r or --rename <xpath1> -v <new-name>
  -u or --update <xpath> -v (--value) <value>
                         -x (--expr) <xpath> (-x is not implemented yet)
XMLStarlet is a command line toolkit to query/edit/check/transform
XML documents (for more information see http://xmlstar.sourceforge.net/)
```

### 格式化: fo
```
#!sh
# 使用4个空格作为缩进
$ xmlstarlet fo -s 4 demo.xml

# 示例
Introduction
XSel is a command-line program for getting and setting the contents of the X selection. Normally this is only accessible by manually highlighting information and pasting it with the middle mouse button.
To read a file into the X selection:
xsel < file
after which you can paste the file's contents into any X application with the middle mouse button, as though you had highlighted its text. xsel will read in the file contents exactly, whereas manual highlighting invariably breaks lines and transforms tabs into spaces. This is especially handy for copying in large files.
To write the X selection to a file:
xsel > file
after which file will contain exactly the contents of the X selection, without trailing newlines and spaces and crap.
XSel is more than just cat for the X selection.
Append to the X selection:
xsel --append < file
To follow a growing file:
xsel --follow < file
to make the X selection follow standard input as it grows (like tail -f).
Advanced features
XSel also lets you access some of the more esoteric features of the X selection:
Delete the contents of the selection
xsel --delete
Will cause the program in which text is selected to delete that text. This really works, you can try it on xedit to remotely delete text in the editor window.
Manipulate the secondary selection
The X Window System maintains two selections, the usual primary selection and a secondary, which isn't used much ... XSel lets you use the secondary selection, for example:
To get and set the secondary selection:
xsel --secondary < file
xsel --secondary > file
To swap the primary and secondary selections:
xsel --exchange
So for example you can store useful text in the secondary selection and retrieve it later.
Manipulate the clipboard selection
Similarly, X has a clipboard selection. You can use the standard xclipboard program to manage a history of selected text, and you can use xsel to actually get text into that clipboard:
xsel --clipboard < file
Make the selection contents persist in memory
Normally the X selection only exists as long as the program it was selected in is running. Further, some buggy applications tend to forget their selection text after a little while. If you run:
xsel --keep
after selecting some important text, xsel will copy the text into its own memory so you can paste it elsewhere even if the original program exits or crashes.
Naturally all these options have single character equivalents, and xsel --help provides usage information. For complete details, see the xsel man page.
Download
XSel is available in most distributions.
apt-get install xsel
or a similar command will install the xsel package on your system.
Source tarballs
XSel is distributed in source form here (xsel-1.2.0.tar.gz).
Development versions are available via Git
git clone git://github.com/kfish/xsel.git
Debian package info for xsel
stable, testing, unstable.
```
