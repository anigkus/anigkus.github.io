
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
# Build the minimum version of the JDK base image 
#
FROM anigkus/glibc:2.34-r0

# Author
LABEL maintainer="anigkus <anigkus@gmail.com>" 

# JDK download , https://www.oracle.com/java/technologies/javase/jdk11-archive-downloads.html
ADD jdk-11.0.13_linux-x64_bin.tar.gz /usr/java/jdk/

# Set C.UTF-8 locale as default and java path
ENV LANG=C.UTF-8

# Dependency Package Info
ENV GLIBC_VERSION=2.34-r0

# Java path
ENV JAVA_VERSION=11.0.13
ENV JAVA_HOME /usr/java/jdk/jdk-11.0.13
ENV PATH ${PATH}:${JAVA_HOME}/bin

# Run container with base path:/opt
WORKDIR /opt