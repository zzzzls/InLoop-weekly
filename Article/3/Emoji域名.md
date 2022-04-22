# InLoop 周刊（第 3 期）：Emoji 域名

Tags: 周刊  
创建时间: April 10, 2022 11:50 AM  
周刊: Yes  
归档: Yes  

# Emoji 域名

Emoji，大家在聊天中应该经常用到吧，合理使用可以减少文字的枯燥乏味，也能准确表达自己的内心活动和想法。但你是否知道，浏览器中也可以使用 emoji 了呢？

### **什么是 Yat**

`y.at` 是一个域名，和 baidu.com 一样，将它输入进浏览器，它将跳转至 y.at 的首页。而如果在浏览器地址栏输入 `y.at/🐸🌹🎋🐴` ，那么浏览器将跳转至我的主页。甚至，在 Opera 浏览器中直接输入 `🐸🌹🎋🐴` 这四个 emoji 后，就可以直达。

后跟的 emoji 符号越少，也就越稀有价格自然更高，以我注册的 `🐸🌹🎋🐴` 为例，仅需要 `$5`，但一些热门 emoji ，可能需要上百 USD

![https://img.lspider.cn/notion/202204101104712.png?imageMogr2/format/webp](https://img.lspider.cn/notion/202204101104712.png?imageMogr2/format/webp)

### **为什么不是 🐸🌹🎋🐴.com**

上述所谓的 "emoji 域名" 再怎么说也只是个子目录域名，如果 y.at 关停或者跑路，那么其下所有的 emoji 域名全部都没了。存在真正 emoji 域名吗？

- ♨️.com
- ♨️.net
- ☮️.com
- ♂️.com

如上四个域名应该是世界上最早注册的四个 emoji 域名。不过，这类域名也存在一些问题：

- **不好注册**：支持 emoji 的域名注册商是有限的
- **浏览器支持不友好**：大部分浏览器会将 emoji 域名显示为一串编码 ♨️.net 『xn--j6h.net』，而 yat 域名则不存在这个问题

### **买完之后可以做什么**

购买完成之后，可以自定义一个个性化的 y.at 主页，如下所示，或者做一个重定向，跳转到指定的网页上

![https://img.lspider.cn/notion/202204101124854.png?imageMogr2/format/webp](https://img.lspider.cn/notion/202204101124854.png?imageMogr2/format/webp)

### **白嫖一个**

如果你仅想体验一下，它的实现原理也很简单，借助 Flask 模块，仅需 10 行代码我们就可以实现上述效果。没错，这就是智商税🥲

![https://img.lspider.cn/notion/202204101145265.png?imageMogr2/format/webp](https://img.lspider.cn/notion/202204101145265.png?imageMogr2/format/webp)

# 一言

一瞬，一刹，一弹指，一须弥

不是形容时间的，是形容快乐的

# 工具

1. [Screenshotapi](https://www.screenshotapi.net/)
   
    调用 api 获取指定 url 的网站截图，免费版每月 100 次限额
    
2. [readme.so](http://readme.so/)
   
    在线生成 Readme 文件，内置多套模板
    
3. [Thunder Client](https://www.thunderclient.com/)
   
    一款 VSCode 插件，可以作为轻量级的 Postman 替代方案
    
    ![Untitled](image/Untitled.png)
    
1. [Qwerty Learner](https://qwerty.kaiyi.cool/)
   
    一个通过打字记忆英语单词的网站，内置多套词库（涵盖从小学到考研词汇）
    
2. [SimpleLogin](https://simplelogin.io/)
   
    在线的邮箱别名服务，提供真实邮箱别名，可以接收、发送邮件，并转发至你的邮箱
    
3. [Duck DNS](https://www.duckdns.org/)
   
    一个免费的动态域名服务，允许注册 [duckdns.org](http://duckdns.org/) 的子域名，可以随时更新指向 ip
    

# 文摘

1. [一条豆瓣动态引发的『疫情文学』复兴](https://www.douban.com/note/827837123/?_dtcc=1&_i=965137992YVUQx,965139492YVUQx)
   
    3月18日，豆瓣网友@K 模仿卡夫卡经典作品《变形记》的开头，发了一条动态：
    
    > **一天早晨，格里高尔.萨姆沙从不安的睡梦中醒来，发现自己小区被封了**
    > 
    
    这句话无异于吹响了创作大赛的号角，众网友纷纷转发评论，掀起一波『疫情文学』高潮
    
    - 多年以后，面对社区防疫人员，奥雷里亚诺•布恩迪亚上校将会回想起父亲带他去商城却被封禁隔离的那个遥远的下午
    - 深蓝的天空中挂着一轮金黄的圆月，下面是新封的小区，都排着一望无际的长队，其间有一个二十来岁的青年，头戴防护罩，手捏一柄棉签，向一张嘴尽力的刺去
    - 小区里有两栋楼，一栋封了，另一栋也封了
    - 天下大势，封久必核，核久必封