<iframe src="../detail-header.html" title="Github of Anigkus" style="height:0px,widht:0px;display:none" id="kusifreamheader"></iframe>

<h1 style="color:#606c71;text-align:center;"  id="h1">常用正则表达式的收藏列表</h1><br/>

[<h1 style="color:#606c71;text-align:center;" >Favorite list of frequently used regular expressions</h1><br/>]:#

![Favorite list of frequently used regular expressions](../assets/images/favorite-list-of-frequently-used-regular-expressions/figure-1.jpeg "Github of Anigkus")

[!Favorite list of frequently used regular expressions(assets/images/favorite-list-of-frequently-used-regular-expressions/figure-1.jpg "Github of Anigkus")]:#

> <br/>&nbsp;&nbsp;&nbsp;&nbsp;各种正则表达式在软件开发工作中非常常见.但有时候需要一些表达式的时候,一时半会也想不出来,所以就干脆就把一些常用的表达方式整理下放到一起，这样后续需要的时候,就干脆直接去复制就行了,避免再次Googles搜索了,也没必要在这方面花更多的时间. 这样才能更好地释放大脑容量,将其用于记住更重要的事情上去.😄<br/>
> <br/>

[> <br/>&nbsp;&nbsp;&nbsp;&nbsp;Various regular expressions are very common in software development. But sometimes when you need some expressions, you can't think of them for a while, so you just put some commonly used expressions together, so that you will need them later. When the time comes, just copy it directly, avoid searching Googles again, and there is no need to spend more time on this. This way, you can better free your brain capacity and use it to remember more important things.😄<br/>]:#
[> <br/>]:#

## 数字表达式 {#id1-h2}
[## Number express {#id1-h2}]:#

[* Digits(zero or more ).]:#
* 数字(0个或多个).
  ```
  ^[0-9]*$
  ```
[* N digits.]:#
* N个数字.
  ```
  ^\d{n}$
  ```
[* At least n digits.]:#
* 最少N个数字.
  ```
  ^\d{n,}$
  ```
[* M to N digits.]:#
* 最少M个和最多N个数字.
  ```
  ^\d{m,n}$
  ```
[* Numbers starting with zero and non-zero.]:#
* 零和非零开头的数字.
  ```
  ^(0|[1-9][0-9]*)$
  ```
[* Non-zero leading number with up to two decimal places.]:#
* 非零开头的最多带两位小数的数字.
  ```
  ^([1-9][0-9]*)+(\.[0-9]{1,2})?$
  ```
[* Positive or negative numbers with 1-2 decimal places.]:#
* 带1-2位小数的正数或负数.
  ```
  ^(\-)?\d+(\.\d{1,2})$
  ```
[* Positive, negative, and decimal numbers.]:#
* 正数、负数、和小数.
  ```
  ^(\-|\+)?\d+(\.\d+)?$
  ```
[* Positive real numbers with two decimal places.]:#
* 有两位小数的正实数.
  ```
  ^[0-9]+(\.[0-9]{2})?$
  ```
[* Positive real numbers with 1 to 3 decimal places.]:#
* 有1~3位小数的正实数.
  ```
  ^[0-9]+(\.[0-9]{1,3})?$
  ```
[* Non-zero Positive integer.]:#
* 非零的正整数.
  ```
  ^[1-9]\d*$
  ^([1-9][0-9]*){1,3}$
  ^\+?[1-9][0-9]*$
  ```
[* Non-zero negative integer.]:#
* 非零的负整数.
  ```
  ^\-[1-9][]0-9"*$ 
  ^-[1-9]\d*$
  ```
[* Non-negative integer.]:#
* 非负整数.
  ```
  ^\d+$ 
  ^[1-9]\d*|0$
  ```
[* Non-positive integer.]:#
* 非正整数.
  ```
  ^-[1-9]\d*|0$
  ^((-\d+)|(0+))$
  ```
[* Non-negative floating point number.]:#
* 非负浮点数.
  ```
  ^\d+(\.\d+)?$
  ^[1-9]\d*\.\d*|0\.\d*[1-9]\d*|0?\.0+|0$
  ```
[* Noe-positive floating point number.]:#
* 非正浮点数.
  ```
  ^((-\d+(\.\d+)?)|(0+(\.0+)?))$
  ^(-([1-9]\d*\.\d*|0\.\d*[1-9]\d*))|0?\.0+|0$
  ```
[* Positive floating point number.]:#
* 正浮点数.
  ```
  ^[1-9]\d*\.\d*|0\.\d*[1-9]\d*$
  ^(([0-9]+\.[0-9]*[1-9][0-9]*)|([0-9]*[1-9][0-9]*\.[0-9]+)|([0-9]*[1-9]- [0-9]*))$
  ```
[* Nagetive floating point number.]:#
* 负浮点数.
  ```
  ^-([1-9]\d*\.\d*|0\.\d*[1-9]\d*)$
  ^(-(([0-9]+\.[0-9]*[1-9][0-9]*)|([0-9]*[1-9][0-9]*\.[0-9]+)|([0-9]*- [1-9][0-9]*)))$
  ```
[* Floating point number.]:#
* 浮点数.
  ```
  ^(-?\d+)(\.\d+)?$ 或 ^-?([1-9]\d*\.\d*|0\.\d*[1-9]\d*|0?\.0+|0)$
  ```

## 字符表达式 {#id2-h2}
[## Character express {#id2-h2}]:#

[* Basic Chinese(zero or more).]:#
* 基础汉字.
  ```
  ^[\u4e00-\u9fa5]{0,}$
  ```
[* English and numbers.]:#
* 英文和数字.
  ```
  ^[A-Za-z0-9]+$
  ^[A-Za-z0-9]{4,40}$ # length 4-40
  ```
[* Any character except newline \n of length 3-20.]:#
* 长度为3-20的所有字符.
  ```
  ^.{3,20}$
  ```
[* A string of 26 english letters.]:#
* 由26个英文字母组成的字符串.
  ```
  ^[A-Za-z]+$
  ```
[* A string of 26 uppercase english letters.]:#
* 由26个大写英文字母组成的字符串.
  ```
  ^[A-Z]+$
  ```
[* A string of 26 lowercase english letters.]:#
* 由26个小写英文字母组成的字符串.
  ```
  ^[a-z]+$
  ```
[* A string of numbers and 26 english letter.]:#
* 由数字和26个英文字母组成的字符串.
  ```
  ^[A-Za-z0-9]+$
  ```
[* A string composed of numbers, 26 english letters or underscores.]:#
* 由数字、26个英文字母或者下划线组成的字符串.
  ```
  ^\w+$
  ^\w{3,20}$ # length 3-20
  ```
[* Basic Chinese, English, numbers including underscore.]:#
* 中文、英文、数字包括下划线.
  ```
  ^[\u4E00-\u9FA5A-Za-z0-9_]+$
  ```
[* Basic Chinese, English, number but not including ubderscores and other symbols.]:#
* 中文、英文、数字但不包括下划线等符号.
  ```
  ^[\u4E00-\u9FA5A-Za-z0-9]+$ 
  ^[\u4E00-\u9FA5A-Za-z0-9]{2,20}$
  ```
[* Characters containing &%',;"=?$ there are not allowed.]:#
* 可以输入含有`&%',;"=?$`等字符
  ```
  [^%&',;=?$]+
  ```
[* Characters containing ~ are not allowed.]:#
* 禁止输入含有~的字符
  ```
  [^~]+
  ```

## 货币表达式 {#id3-h2}
[## Currency expression {#id3-h2}]:#

[* Match numbers that start with non-zero.]:#
* 不是0开头的数字
  ```
  ^[1-9][0-9]*$
  ```
[* Match numbers zero or start with non-zero.]:#
* 匹配数字零或以非零开头
  ```
  ^(0|[1-9][0-9]*)$
  ```
[* Match numbers negative, zero or start with non-zero.]:#
* 匹配负数、零或以非零开头的数字
  ```
  ^(0|-?[1-9][0-9]*)$
  ```
[* Match amount expression.]:#
* 匹配金额表达式.
  ```
  ^[0-9]+(\.[0-9]+)?$
  ```
[* Match amount expression with must have 2 decimal places.]:#
* 匹配金额表达式必须有2位小数.
  ```
  ^[0-9]+(\.[0-9]{2})?$
  ```
[* Match amount expression with must have 1~2 decimal places.]:#
* 匹配金额表达式必须有1~2位小数.
  ```
  ^[0-9]+(\.[0-9]{1,2})?$
  ```
[* Amount as a 3-digit with must have comma-separated expression and have 1~2 decimal places.]:#
* 金额为3位数，必须有逗号分隔的表达式，小数点后1~2位.
  ```
  ^[0-9]{1,3}(,[0-9]{3})*(.[0-9]{1,2})?$ 
  ```
[* Amount as a 3-digit not must comma-separated expression and have 1~2 decimal places,Only all a 3-digit comma-separated or all not comma-separated.]:#
* 1到3个数字,后面跟着任意个逗号+3个数字,逗号成为可选,而不是必须.
  ```
  ^([0-9]+|[0-9]{1,3}(,[0-9]{3})*)(.[0-9]{1,2})?$
  ```
[* Basic Chinese.]:#
* 中文基础汉字的正则表达式.
  ```
  [\u4e00-\u9fa5]
  ```
[* Duble-byte characters expression.]:#
* 双字节字符.
  ```
  [^\x00-\xff]
  ```
[* Blank newline expression.]:#
* 空白行的正则表达式.
  ```
  \n\s*\r
  ```
[* Regular expression for HTML tags.]:#
* HTML标记的正则表达式.
  ```
  <(\S*?)[^>]*>.*?|<.*? />
  ```
[* Regular expression for leading and trailing whitespace characters.]:#
* 首尾空白字符的正则表达式.
  ```
  ^\s*|\s*$
  (^\s*)|(\s*$)
  ```
[* Tencet QQ Account.]:#
* 腾讯QQ号.
  ```
  [1-9][0-9]{4,}
  ```
[* China Postal Code.]:#
* 中国邮政编码.
  ```
  [1-9]\d{5}(?!\d) 
  ```
[* IPv4 regular  expression.]:#
* IPv4地址.
  ```
  ((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})(\.((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})){3}
  ```

## 特殊表达式 {#id4-h2}
[## Special expression {#id4-h2}]:#


[* Matches characters starting with `- ` or `\[- ` on each line.]:#
* 匹配每行以 `- `开头或者以 `[- ` 开头的字符.
  ```
  ^- |^\[- 
  ```
[* Email regular expression.]:#
* Email地址.
  ```
  ^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$
  ```
[* Domain name regular expression.]:#
* 域名.
  ```
  [a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+\.?
  ```
[* Internet URL regular expression.]:#
* http网址.
  ```
  [a-zA-z]+://[^\s]*
  ^http://([\w-]+\.)+[\w-]+(/[\w-./?%&=]*)?$ #Just start with http
  ```
[* China mobile phone regular expression.]:#
* 中国手机号码.
  ```
  ^(13[0-9]|14[5|7]|15[0|1|2|3|4|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])\d{8}$
  ```
[* China mobile phone regular expression for XXX-XXXXXXX.]:#
* 中国手机号码,格式XXX-XXXXXXX、XXXX-XXXXXXXX、XXX-XXXXXXX、XXX-XXXXXXXX、XXXXXXX、XXXXXXXX.
  ```
  ^(\(\d{3,4}-)|\d{3.4}-)?\d{7,8}$
  ```
[* China fixed phone.]:#
* 中国固定电话.
  ```
  \d{3}-\d{8}|\d{4}-\d{7}
  ```
[* Can support including mobile phone number, {3,4} area code - {7,8} direct number + {1,4} extension, {7,8} direct number + {1,4} extension.]:#
* 电话号码正则表达式(支持手机号码，3-4位区号，7-8位直播号码，1－4位分机号).
  ```
  ((\d{11})|^((\d{7,8})|(\d{4}|\d{3})-(\d{7,8})|(\d{4}|\d{3})-(\d{7,8})-(\d{4}|\d{3}|\d{2}|\d{1})|(\d{7,8})-(\d{4}|\d{3}|\d{2}|\d{1}))$)
  ```
[* Chinese ID number, the last digit may by X.]:#
* 身份证号(15位、18位数字)，最后一位是校验位，可能为数字或字符X.
  ```
  (^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)
  ```
[* Start with a letters,allow 5-16 bytes, allow alphanumeric underscore.]:#
* 帐号是否合法(字母开头，允许5-16字节，允许字母数字下划线).
  ```
  ^[a-zA-Z][a-zA-Z0-9_]{4,15}$
  ```
[* Start with a letters, the length is between 6 and 18, and can only contain letters, numbers and underscores.]:#
* 密码(以字母开头，长度在6~18之间，只能包含字母、数字和下划线).
  ```
  ^[a-zA-Z]\w{5,17}$
  ```
[* Strong password,Must contain a combination of uppercase and lowercase letters and numbers, cannot use special characters, and the length is between 8-10.]:#
* 强密码(必须包含大小写字母和数字的组合，不能使用特殊字符，长度在 8-10 之间).
  ```
  ^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])[a-zA-Z0-9]{8,10}$
  ```
[* Strong password,Must contain a combination of uppercase and lowercase letters and numbers,special characters can be used, and the length is between 8-10.]:#
* 强密码(必须包含大小写字母和数字的组合，可以使用特殊字符，长度在8-10之间).
  ```
  ^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{8,10}$
  ```
[* Data format regular expression.]:#
* 日期格式.
  ```
  ^\d{4}-\d{1,2}-\d{1,2}
  ```
[* Month of regular expression.]:#
* 一年的12个月(01～09和1～12).
  ```
  ^(0?[1-9]|1[0-2])$
  ```
[* Every day of regular expression.]:#
* 一个月的31天(01～09和1～31).
  ```
  ^((0?[1-9])|((1|2)[0-9])|30|31)$
  ```

<br>

[back](./)