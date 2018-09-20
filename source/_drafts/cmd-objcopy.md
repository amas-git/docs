---
title: objcopy
tags:
---
# objcopy: 向ELF文件中嵌入数据
怎么往ELF文件中嵌入数据呢？
# nm
```
$ nm a.out
08049558 d _DYNAMIC
08049644 d _GLOBAL_OFFSET_TABLE_
0804846c R _IO_stdin_used
         w _ITM_deregisterTMCloneTable
         w _ITM_registerTMCloneTable
         w _Jv_RegisterClasses
08048548 r __FRAME_END__
08049554 d __JCR_END__
08049554 d __JCR_LIST__
08049660 D __TMC_END__
08049660 B __bss_start
08049658 D __data_start
08048380 t __do_global_dtors_aux
08049550 t __do_global_dtors_aux_fini_array_entry
0804965c D __dso_handle
0804954c t __frame_dummy_init_array_entry
         w __gmon_start__
08049550 t __init_array_end
0804954c t __init_array_start
08048450 T __libc_csu_fini
080483e0 T __libc_csu_init
         U __libc_start_main@@GLIBC_2.0
08048300 T __x86.get_pc_thunk.bx
08049660 D _edata
08049664 B _end
08048454 T _fini
08048468 R _fp_hw
08048274 T _init
080482d0 T _start
08049660 b completed.5975
08049658 W data_start
08048310 t deregister_tm_clones
080483a0 t frame_dummy
080483d0 T main
08048340 t register_tm_clones
# 查看动态链接符号
$ nm -D a.out
0804846c R _IO_stdin_used
         w __gmon_start__
         U __libc_start_main
# 如果目标是动态链接库，没有符号表也是可以查看动态链接符号的，因为动态连接器需要这些信息实现代码的重定位
$ nm -D x.so
# 可以使用strip命令删掉符号表
$ strip a.out
$ nm a.out
nm: b.out: no symbols
# 有时候可能会碰到C++编译器生成的目标文件，这时候直接用nm来查看符号就不太直观了，可以利用c++filter或者-demangle参数来破
$ nm a.out | c++filter
$ nm -demangle a.out
```
# strings: 从二进制文件中提取字符串
```sh
```
# ldd : 查看动态链接库的依赖关系
```sh
$ file /usr/bin/ldd 
/usr/bin/ldd: Bourne-Again shell script, ASCII text executable
```
```
# This is the `ldd' command, which lists what shared libraries are
# used by given dynamically-linked executables.  It works by invoking the
# run-time dynamic linker as a command and setting the environment
# variable LD_TRACE_LOADED_OBJECTS to a non-empty value.
```
```
$ LD_TRACE_LOADED_OBJECTS=1 /bin/ls                                                                                                                                                          ~
        linux-gate.so.1 (0xb77b6000)
        libcap.so.2 => /usr/lib/libcap.so.2 (0xb7798000)
        libacl.so.1 => /usr/lib/libacl.so.1 (0xb778f000)
        libc.so.6 => /usr/lib/libc.so.6 (0xb75e0000)
        libattr.so.1 => /usr/lib/libattr.so.1 (0xb75da000)
        /lib/ld-linux.so.2 (0xb77b7000)
```
# strip
strip默认删掉不必要的段，比方说符号表之类的.
也可以使用-R <section-name>来删除指定的段
```sh
$ strip -R .text a.out 
# 删除代码段
```
# addr2line
地址转行号。
# head
# hexdump
# 以十进制方式打印文件的前四个字节
```sh
$ hexdump -e '10/1 "%01d "' -n 4 target.dat
14 2 5 11 
$ hexdump -e '10/1 "%02d "' -n 4 target.dat
14 02 05 11
```

