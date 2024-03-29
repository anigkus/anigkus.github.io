/*
Copyright 2022 The https://github.com/anigkus Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
/**
 * https://www.tutorialspoint.com/vuejs/vuejs_environment_setup.htm
 * https://www.tutorialspoint.com/typescript/typescript_variables.htm
 * https://www.electronjs.org/
 */
function app() {
    var message = "Hello, App!";
    console.log(message);
    var _a = [1, [[21, 5], 3]], a = _a[0], _b = _a[1], b = _b[0][0], c = _b[1];
    console.log(a + "," + b + "," + c);
}
//start App
app();
var sy = Symbol("KK");
console.log(sy); // Symbol(KK)
