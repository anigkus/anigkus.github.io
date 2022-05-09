<script>
var pageHeader=document.getElementsByClassName("page-header")[0].innerHTML;
 pageHeader="<center><img style='border-radius: 50% !important;' src='https://avatars.githubusercontent.com/u/88264073?s=400&amp;u=63e618520a5b6aa87636714e69f8228374c4e9b1&amp;v=4' width='200' height='200' alt='@anigkus' title='Github of Anigkus' ></center>"+pageHeader;
document.getElementsByClassName("page-header")[0].innerHTML=pageHeader;
</script>

<h1 style="color:#606c71;text-align:center;" >Favorite list of frequently used regular expressions</h1><br/>

![Favorite list of frequently used regular expressions](assets/images/favorite-list-of-frequently-used-regular-expressions/figure-1.jpg "Github of Anigkus")

> <br/>
> &nbsp;&nbsp;&nbsp;&nbsp; Various regular expressions are very common in software development work. But when some expressions are sometimes needed, I can't think of them all at once, so I just spend some time sorting out some frequently used expressions to avoid I'll spend more time on this later. This way, I can better free my brain and use the capacity for more important things.<br/>
> <br/>

## Number express {#id1-h2}
- Digits(zero or more ).
  ```
  ^[0-9]*$
  ```
- N digits.
  ```
  ^\d{n}$
  ```
- At least n digits.
  ```
  ^\d{n,}$
  ```
- M to N digits.
  ```
  ^\d{m,n}$
  ```
- Numbers starting with zero and non-zero.
  ```
  ^(0|[1-9][0-9]*)$
  ```
- Non-zero leading number with up to two decimal places.
  ```
  ^([1-9][0-9]*)+(\.[0-9]{1,2})?$
  ```
- Positive or negative numbers with 1-2 decimal places.
  ```
  ^(\-)?\d+(\.\d{1,2})$
  ```
- Positive, negative, and decimal numbers.
  ```
  ^(\-|\+)?\d+(\.\d+)?$
  ```
- Positive real numbers with two decimal places.
  ```
  ^[0-9]+(\.[0-9]{2})?$
  ```
- Positive real numbers with 1 to 3 decimal places.
  ```
  ^[0-9]+(\.[0-9]{1,3})?$
  ```
- Non-zero Positive integer.
  ```
  ^[1-9]\d*$
  ^([1-9][0-9]*){1,3}$
  ^\+?[1-9][0-9]*$
  ```
- Non-zero negative integer.
  ```
  ^\-[1-9][]0-9"*$ 
  ^-[1-9]\d*$
  ```
- Non-negative integer.
  ```
  ^\d+$ 
  ^[1-9]\d*|0$
  ```
- Non-positive integer.
  ```
  ^-[1-9]\d*|0$
  ^((-\d+)|(0+))$
  ```
- Non-negative floating point number.
  ```
  ^\d+(\.\d+)?$
  ^[1-9]\d*\.\d*|0\.\d*[1-9]\d*|0?\.0+|0$
  ```
- Noe-positive floating point number.
  ```
  ^((-\d+(\.\d+)?)|(0+(\.0+)?))$
  ^(-([1-9]\d*\.\d*|0\.\d*[1-9]\d*))|0?\.0+|0$
  ```
- Positive floating point number.
  ```
  ^[1-9]\d*\.\d*|0\.\d*[1-9]\d*$
  ^(([0-9]+\.[0-9]*[1-9][0-9]*)|([0-9]*[1-9][0-9]*\.[0-9]+)|([0-9]*[1-9]- [0-9]*))$
  ```
- Nagetive floating point number.
  ```
  ^-([1-9]\d*\.\d*|0\.\d*[1-9]\d*)$
  ^(-(([0-9]+\.[0-9]*[1-9][0-9]*)|([0-9]*[1-9][0-9]*\.[0-9]+)|([0-9]*- [1-9][0-9]*)))$
  ```
- Floating point number.
  ```
  ^(-?\d+)(\.\d+)?$ 或 ^-?([1-9]\d*\.\d*|0\.\d*[1-9]\d*|0?\.0+|0)$
  ```

## Character express {#id2-h2}
- Basic Chinese(zero or more).
  ```
  ^[\u4e00-\u9fa5]{0,}$
  ```
- English and numbers.
  ```
  ^[A-Za-z0-9]+$
  ^[A-Za-z0-9]{4,40}$ # length 4-40
  ```
- Any character except newline \n of length 3-20.
  ```
  ^.{3,20}$
  ```
- A string of 26 english letters.
  ```
  ^[A-Za-z]+$
  ```
- A string of 26 uppercase english letters.
  ```
  ^[A-Z]+$
  ```
- A string of 26 lowercase english letters.
  ```
  ^[a-z]+$
  ```
- A string of numbers and 26 english letter.
  ```
  ^[A-Za-z0-9]+$
  ```
- A string composed of numbers, 26 english letters or underscores.
  ```
  ^\w+$
  ^\w{3,20}$ # length 3-20
  ```
- Basic Chinese, English, numbers including underscore.
  ```
  ^[\u4E00-\u9FA5A-Za-z0-9_]+$
  ```
- Basic Chinese, English, number but not including ubderscores and other symbols.
  ```
  ^[\u4E00-\u9FA5A-Za-z0-9]+$ 
  ^[\u4E00-\u9FA5A-Za-z0-9]{2,20}$
  ```
- Characters containing &%',;"=?$ there are not allowed.
  ```
  [^%&',;=?$]+
  ```
- Characters containing ~ are not allowed.
  ```
  [^~]+
  ```

## Currency expression {#id3-h2}
- Match numbers that start with non-zero.
  ```
  ^[1-9][0-9]*$
  ```
- Match numbers zero or start with non-zero.
  ```
  ^(0|[1-9][0-9]*)$
  ```
- Match numbers negative, zero or start with non-zero.
  ```
  ^(0|-?[1-9][0-9]*)$
  ```
- Match amount expression.
  ```
  ^[0-9]+(\.[0-9]+)?$
  ```
- Match amount expression with must have 2 decimal places.
  ```
  ^[0-9]+(\.[0-9]{2})?$
  ```
- Match amount expression with must have 1~2 decimal places.
  ```
  ^[0-9]+(\.[0-9]{1,2})?$
  ```
- Amount as a 3-digit with must have comma-separated expression and have 1~2 decimal places.
  ```
  ^[0-9]{1,3}(,[0-9]{3})*(.[0-9]{1,2})?$ 
  ```
- Amount as a 3-digit not must comma-separated expression and have 1~2 decimal places,Only all a 3-digit comma-separated or all not comma-separated.
  ```
  ^([0-9]+|[0-9]{1,3}(,[0-9]{3})*)(.[0-9]{1,2})?$
  ```
- Basic Chinese.
  ```
  [\u4e00-\u9fa5]
  ```
- Duble-byte characters expression.
  ```
  [^\x00-\xff]
  ```
- Blank newline expression.
  ```
  \n\s*\r
  ```
- Regular expression for HTML tags.
  ```
  <(\S*?)[^>]*>.*?|<.*? />
  ```
- Regular expression for leading and trailing whitespace characters.
  ```
  ^\s*|\s*$
  (^\s*)|(\s*$)
  ```
- Tencet QQ Account.
  ```
  [1-9][0-9]{4,}
  ```
- China Postal Code.
  ```
  [1-9]\d{5}(?!\d) 
  ```
- IPv4 regular  expression.
  ```
  ((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})(\.((2(5[0-5]|[0-4]\d))|[0-1]?\d{1,2})){3}
  ```

## Special expression {#id4-h2}
- Email regular expression.
  ```
  ^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$
  ```
-  Domain name regular expression.
  ```
  [a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+\.?
  ```
- Internet URL regular expression.
  ```
  [a-zA-z]+://[^\s]*
  ^http://([\w-]+\.)+[\w-]+(/[\w-./?%&=]*)?$ #Just start with http
  ```
- China mobile phone regular expression.
  ```
  ^(13[0-9]|14[5|7]|15[0|1|2|3|4|5|6|7|8|9]|18[0|1|2|3|5|6|7|8|9])\d{8}$
  ```
- China mobile phone regular expression for XXX-XXXXXXX.
  ```
  ^(\(\d{3,4}-)|\d{3.4}-)?\d{7,8}$
  ```
- China fixed phone.
  ```
  \d{3}-\d{8}|\d{4}-\d{7}
  ```
- Can support including mobile phone number, {3,4} area code - {7,8} direct number + {1,4} extension, {7,8} direct number + {1,4} extension
  ```
  ((\d{11})|^((\d{7,8})|(\d{4}|\d{3})-(\d{7,8})|(\d{4}|\d{3})-(\d{7,8})-(\d{4}|\d{3}|\d{2}|\d{1})|(\d{7,8})-(\d{4}|\d{3}|\d{2}|\d{1}))$)
  ```
- Chinese ID number, the last digit may by X.
  ```
  (^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)
  ```
- Start with a letters,allow 5-16 bytes, allow alphanumeric underscore.
  ```
  ^[a-zA-Z][a-zA-Z0-9_]{4,15}$
  ```
- Start with a letters, the length is between 6 and 18, and can only contain letters, numbers and underscores.
  ```
  ^[a-zA-Z]\w{5,17}$
  ```
- Strong password,Must contain a combination of uppercase and lowercase letters and numbers, cannot use special characters, and the length is between 8-10.
  ```
  ^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])[a-zA-Z0-9]{8,10}$
  ```
- Strong password,Must contain a combination of uppercase and lowercase letters and numbers,special characters can be used, and the length is between 8-10.
  ```
  ^(?=.*\d)(?=.*[a-z])(?=.*[A-Z]).{8,10}$
  ```
- Data format regular expression.
  ```
  ^\d{4}-\d{1,2}-\d{1,2}
  ```
- Month of regular expression.
  ```
  ^(0?[1-9]|1[0-2])$
  ```
- Every day of regular expression.
  ```
  ^((0?[1-9])|((1|2)[0-9])|30|31)$
  ```

<br>

[back](./)