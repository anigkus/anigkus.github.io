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


#gc=`git log --since='2022-05-21 00:00:00' --author="anigkus" --oneline | wc -l`


gc=3
fc=`expr 5 - $gc`
if [ $fc -gt 0 ]
then
  while [ $fc -gt 0 ]
  do
     echo "# "`date`"\n  It's auto generated content." >> ./auto.md
     commit='fix:(example) Auto commit.'`date`
     SHELL_FOLDER=$(cd "$(dirname "$0")";pwd)
     git add $SHELL_FOLDER/auto.md -m $commit
     sleep 5
    (( fc-- ))
  done
fi