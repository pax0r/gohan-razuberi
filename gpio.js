// Copyright (C) 2015 Bartłomiej Biernacki
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied.
// See the License for the specific language governing permissions and
// limitations under the License.

function gpio_handler(context) {
    var exampleModule = require("razuberi");
    var resource = context.resource;
    var gpio = resource.gpio_pin;

    if (resource.is_up) {
        console.log("GPIO " + gpio + " HIGH");
        exampleModule.SetHigh(gpio);
    } else {
        console.log("GPIO " + gpio + " LOW");
        exampleModule.SetLow(gpio);
    }
}


gohan_register_handler("post_create_in_transaction", gpio_handler);
gohan_register_handler("post_update_in_transaction", gpio_handler);
gohan_register_handler("pre_delete_in_transaction", function(context) {
    var resource = context.resource;
    var exampleModule = require("razuberi");

    exampleModule.SetLow(resource.gpio_pin);
});
