
# Favorite list of frequently used regular expressions {#id1-h1}

## Number expression {#id1-h2}
- Digits(zero or more )
  ```
  ^[0-9]*$
  ```
- N digits
  ```
  ^\d{n}$
  ```
- At least n digits
  ```
  ^\d{n,}$
  ```
- M to N digits
  ```
  ^\d{m,n}$
  ```
- Numbers starting with zero and non-zero
  ```
  ^(0|[1-9][0-9]*)$
  ```
- Non-zero leading number with up to two decimal places
  ```
  ^([1-9][0-9]*)+(\.[0-9]{1,2})?$
  ```
- Positive or negative numbers with 1-2 decimal places
  ```
  ^(\-)?\d+(\.\d{1,2})$
  ```
- Positive, negative, and decimal numbers
  ```
  ^(\-|\+)?\d+(\.\d+)?$
  ```
- Positive real numbers with two decimal places
  ```
  ^[0-9]+(\.[0-9]{2})?$
  ```
- Positive real numbers with 1 to 3 decimal places
  ```
  ^[0-9]+(\.[0-9]{1,3})?$
  ```
- Non-zero Positive integer
  ```
  ^[1-9]\d*$
  ^([1-9][0-9]*){1,3}$
  ^\+?[1-9][0-9]*$
  ```
- Non-zero negative integer
  ```
  ^\-[1-9][]0-9"*$ 
  ^-[1-9]\d*$
  ```
- 非负整数：
- Non-negative integer
  ```
  ^\d+$ 
  ^[1-9]\d*|0$
  ```
- 非正整数：
- Non-positive integer
  ```
  ^-[1-9]\d*|0$
  ^((-\d+)|(0+))$
  ```
- 非负浮点数：
- Non-negative floating point number
  ```
  ^\d+(\.\d+)?$
  ^[1-9]\d*\.\d*|0\.\d*[1-9]\d*|0?\.0+|0$
  ```
- 非正浮点数：
- Noe-positive floating point number
  ```
  ^((-\d+(\.\d+)?)|(0+(\.0+)?))$
  ^(-([1-9]\d*\.\d*|0\.\d*[1-9]\d*))|0?\.0+|0$
  ```
- 正浮点数：
- Positive floating point number
  ```
  ^[1-9]\d*\.\d*|0\.\d*[1-9]\d*$
  ^(([0-9]+\.[0-9]*[1-9][0-9]*)|([0-9]*[1-9][0-9]*\.[0-9]+)|([0-9]*[1-9]- [0-9]*))$
  ```
- 负浮点数：
- Nagetive floating point number
  ```
  ^-([1-9]\d*\.\d*|0\.\d*[1-9]\d*)$
  ^(-(([0-9]+\.[0-9]*[1-9][0-9]*)|([0-9]*[1-9][0-9]*\.[0-9]+)|([0-9]*- [1-9][0-9]*)))$
  ```
- 浮点数：
- Floating point number
  ```
  ^(-?\d+)(\.\d+)?$ 或 ^-?([1-9]\d*\.\d*|0\.\d*[1-9]\d*|0?\.0+|0)$
  ```

## Character expression {#id2-h2}
- 汉字：
- Chinese(zero or more)
  ```
  ^[\u4e00-\u9fa5]{0,}$
  ```
- 英文和数字：
- English and numbers
  ```
  ^[A-Za-z0-9]+$
  ^[A-Za-z0-9]{4,40}$ # length 4-40
  ```
- 长度为3-20的除换行符\n之外的任何字符：
- Any character except newline \n of length 3-20
  ```
  ^.{3,20}$
  ```
- 由26个英文字母组成的字符串：
- A string of 26 english letters
  ```
  ^[A-Za-z]+$
  ```
- 由26个大写英文字母组成的字符串：
- A string of 26 uppercase english letters
  ```
  ^[A-Z]+$
  ```
- 由26个小写英文字母组成的字符串：
- A string of 26 lowercase english letters
  ```
  ^[a-z]+$
  ```
- 由数字和26个英文字母组成的字符串：
- A string of numbers and 26 english letter
  ```
  ^[A-Za-z0-9]+$
  ```
- 由数字、26个英文字母或者下划线组成的字符串：
- A string composed of numbers, 26 english letters or underscores
  ```
  ^\w+$
  ^\w{3,20}$ # length 3-20
  ```
- 中文、英文、数字包括下划线：
- Chinese, English, numbers including underscore
  ```
  ^[\u4E00-\u9FA5A-Za-z0-9_]+$
  ```
- 中文、英文、数字但不包括下划线等符号：
- Chinese, English, number but not including ubderscores and other symbols
  ```
  ^[\u4E00-\u9FA5A-Za-z0-9]+$ 
  ^[\u4E00-\u9FA5A-Za-z0-9]{2,20}$
  ```
- 可以输入含有^%&',;=?$\"等字符：
- Characters containing &%',;"=?$ there are not allowed
  ```
  [^%&',;=?$]+
  ```
- 禁止输入含有~的字符：
- Characters containing ~ are not allowed
  ```
  [^~]+
  ```

## Currency expression {#id3-h2}
- 有四种钱的表示形式我们可以接受:"10000.00" 和 "10,000.00", 和没有 "分" 的 "10000" 和 "10,000"：
- Match numbers that start with non-zero
  ```
  ^[1-9][0-9]*$
  ```
- 这表示任意一个不以0开头的数字,但是,这也意味着一个字符"0"不通过,所以我们采用下面的形式：
- Match numbers zero or start with non-zero
  ```
  ^(0|[1-9][0-9]*)$
  ```
- 一个0或者一个不以0开头的数字.我们还可以允许开头有一个负号：
- Match numbers negative, zero or start with non-zero
  ```
  ^(0|-?[1-9][0-9]*)$
  ```
- 这表示一个0或者一个可能为负的开头不为0的数字.让用户以0开头好了.把负号的也去掉,因为钱总不能是负的吧。下面我们要加的是说明可能的小数部分：
- Match amount expression
  ```
  ^[0-9]+(\.[0-9]+)?$
  ```
- 必须说明的是,小数点后面至少应该有1位数,所以"10."是不通过的,但是 "10" 和 "10.2" 是通过的：
- Match amount expression with must have 2 decimal places
  ```
  ^[0-9]+(\.[0-9]{2})?$
  ```
- 这样我们规定小数点后面必须有两位,如果你认为太苛刻了,可以这样：
- Match amount expression with must have 1~2 decimal places
  ```
  ^[0-9]+(\.[0-9]{1,2})?$
  ```
- 这样就允许用户只写一位小数.下面我们该考虑数字中的逗号了,我们可以这样：
- ```
- ^[0-9]{1,3}(,[0-9]{3})*(.[0-9]{1,2})?$
- ```
- 1到3个数字,后面跟着任意个 逗号+3个数字,逗号成为可选,而不是必须：
  ```
  ^([0-9]+|[0-9]{1,3}(,[0-9]{3})*)(.[0-9]{1,2})?$
  ```
- 备注：这就是最终结果了,别忘了"+"可以用"*"替代如果你觉得空字符串也可以接受的话(奇怪,为什么?)最后,别忘了在用函数时去掉去掉那个反斜杠,一般的错误都在这里
- xml文件：
  ```
  ^([a-zA-Z]+-?)+[a-zA-Z0-9]+\\.[x|X][m|M][l|L]$
  ```
- 中文字符的正则表达式：
  ```
  [\u4e00-\u9fa5]
  ```
- 双字节字符, (包括汉字在内，可以用来计算字符串的长度(一个双字节字符长度计2，ASCII字符计1))
  ```
  [^\x00-\xff]
  ```
- 空白行的正则表达式：(可以用来删除空白行)
  ```
  \n\s*\r
  ```
- HTML标记的正则表达式：
  ```
  <(\S*?)[^>]*>.*?|<.*? />
  ```
- 首尾空白字符的正则表达式： (可以用来删除行首行尾的空白字符(包括空格、制表符、换页符等等)，非常有用的表达式)
  ```
  ^\s*|\s*$或(^\s*)|(\s*$)
  ```
- 腾讯QQ号： (腾讯QQ号从10000开始)
  ```
  [1-9][0-9]{4,}
  ```
- 中国邮政编码：(中国邮政编码为6位数字)
  ```
  [1-9]\d{5}(?!\d) 
  ```
- IPv4地址：
  ```
  ((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})(\.((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})){3}
  ```

## Special expression {#id3-h2}
- Email地址：
  ```
  ^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$
  ```
- 域名：
  ```
  [a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+\.?
  ```
- InternetURL：
  ```
  [a-zA-z]+://[^\s]* 或 ^http://([\w-]+\.)+[\w-]+(/[\w-./?%&=]*)?$
  ```
- 手机号码：
  ```
  ^(13[0-9]|14[5|7]|15[0|1|2|3|4|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])\d{8}$
  ```
- 电话号码("XXX-XXXXXXX"、"XXXX-XXXXXXXX"、"XXX-XXXXXXX"、"XXX-XXXXXXXX"、"XXXXXXX"和"XXXXXXXX)：
  ```
  ^(\(\d{3,4}-)|\d{3.4}-)?\d{7,8}$
  ```
- 国内电话号码(0511-4405222、021-87888822)：
  ```
  \d{3}-\d{8}|\d{4}-\d{7}
  ```
- 电话号码正则表达式（支持手机号码，3-4位区号，7-8位直播号码，1－4位分机号）:
  ```
  ((\d{11})|^((\d{7,8})|(\d{4}|\d{3})-(\d{7,8})|(\d{4}|\d{3})-(\d{7,8})-(\d{4}|\d{3}|\d{2}|\d{1})|(\d{7,8})-(\d{4}|\d{3}|\d{2}|\d{1}))$)
  ```
- 身份证号(15位、18位数字)，最后一位是校验位，可能为数字或字符X：
  ```
  (^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)
  ```
- 帐号是否合法(字母开头，允许5-16字节，允许字母数字下划线)：
  ```
  ^[a-zA-Z][a-zA-Z0-9_]{4,15}$
  ```
- 密码(以字母开头，长度在6~18之间，只能包含字母、数字和下划线)：
  ```
  ^[a-zA-Z]\w{5,17}$
  ```
- 强密码(必须包含大小写字母和数字的组合，不能使用特殊字符，长度在 8-10 之间)：
  ```
  ^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])[a-zA-Z0-9]{8,10}$
  ```
- 强密码(必须包含大小写字母和数字的组合，可以使用特殊字符，长度在8-10之间)：
  ```
  ^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{8,10}$
  ```
- 日期格式：
  ```
  ^\d{4}-\d{1,2}-\d{1,2}
  ```
- 一年的12个月(01～09和1～12)：
  ```
  ^(0?[1-9]|1[0-2])$
  ```
- 一个月的31天(01～09和1～31)：
  ```
  ^((0?[1-9])|((1|2)[0-9])|30|31)$
  ```