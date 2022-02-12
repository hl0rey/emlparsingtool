# emlparsingtool
一个直接从eml文件格式的邮件中提取附件和内容的命令行工具（A command line tool that extracts attachments and content directly from emL file format messages）

# 参数说明

-a    

if you want export attach file.

-ah

print all header.

-all

export attach file and email content

-c    

if you want export mail content.

-f string 

the eml filename.

    如果字段涉及RFC2047编码会自动解码，支持中文编码(GBK)解码并自动转为UTF8

-a


直接把附件导出到当前相同目录

-ah

打印出当前所有头

-all

导出邮件内容和附件到当前目录

-c 只导出邮件内容

-f 被解析的eml文件名


# 引用
对邮件头的解析，基于DusanKasan大佬的代码，修复了go标准库对rfc7578中4.2节标准实现导致的bug，并添加了对中文解码的支持
https://github.com/DusanKasan/parsemail