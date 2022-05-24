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
#  
# 
# Debug 1: */1 * * * * sh -xv /Users/frank/STS/github/anigkus.github.io/example/auto.sh >>/tmp/xx.log 2>&1 
# Debug 2: sh -xv auto.sh
# Debug 3: set -x; sh auto.sh
# Debug 4: bash -x auto.sh
# DEbug 5: #git push >>/tmp/xx.log 2>&1 #debug crontab: mac terminal nothing
#


SHELL_FOLDER=$(cd "$(dirname "$0")";pwd)
cd $SHELL_FOLDER
git pull 
date=`date +"%Y-%m-%d" `
gcs=`git log --since='$date 00:00:00' --author="anigkus" --oneline | wc -l`
gc="$gcs" 
#gc=3
fc=`expr 5 - $gc`
if [ $fc -gt 0 ]
then
  while [ $fc -gt 0 ]
  do
     echo "# "`date`"\n  It's auto generated content." >> $SHELL_FOLDER/auto.md
     commit='fix:(example) Auto commit.'`date`
     git commit -m "$commit" $SHELL_FOLDER/auto.md 
     sleep 5
    (( fc-- ))
  done
else
 echo "Nothing..."
fi
git push

