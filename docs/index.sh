#!/bin/bash

# Copyright 2022 The https://github.com/anigkus Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at#

#     http://www.apache.org/licenses/LICENSE-2.0#

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# 
# Generate Homepage Content
#

ul="";
assets_prefix=""
pwd=`pwd`

# Check Args
if [ "$1" != "" ];then
  cd "$pwd/zh-cn/"
  assets_prefix="../"
fi

# Parse HTML
for filename in `ls -al | grep ".md" | grep -Ev "index.|anigkus-|draft-" | awk '{print $9}'`;  
do          
  title_name=`cat \`pwd\`"/$filename" | grep "<h1" | head -1 | sed 's:<h1[^>]*>\([^<]*\)</h1><br/>:\1:;q'`;   ul+="- [$title_name](./$filename)<br/>\n" 
done

index_content="<script>\nvar pageHeader=document.getElementsByClassName(\"page-header\")[0].innerHTML;\npageHeader=\"<center><img style='border-radius: 50% \!important;' src='https://avatars.githubusercontent.com/u/88264073?s=400&amp;u=63e618520a5b6aa87636714e69f8228374c4e9b1&amp;v=4' width='200' height='200' alt='@anigkus' title='Github of Anigkus' ></center>\"+pageHeader;\ndocument.getElementsByClassName(\"page-header\")[0].innerHTML=pageHeader;\n</script>\n\n![Anigkus github article template title]("$assets_prefix"assets/images/figure-1.jpeg \"Github of Anigkus\")
<br/>\n\n$ul";

# Input file
if [ "$assets_prefix" == "../" ]; then 
  echo $index_content > "$pwd/zh-cn/index.md";
else 
  echo $index_content > "$pwd/index.md";
fi