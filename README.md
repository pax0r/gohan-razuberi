# Gohan Razuberi

## What it is?

It's an REST service to manage Raspberry PI GPIOs. It is build on top of [Gohan framework](http://gohan.cloudwan.io/).

It aims to provide simple REST API to manage devices connected to RPi via GPIO pins.

It uses [go-rpi](https://github.com/stianeikeland/go-rpio) library for gpios and supports same rpi version as this library.

## How to run it

1. Install go on RasPi (<http://dave.cheney.net/2015/09/04/building-go-1-5-on-the-raspberry-pi>)
2. Install gohan (<http://gohan.cloudwan.io/gohan/installation.html>)
3. Install go-rpi: `go get github.com/stianeikeland/go-rpio`
4. Run: `go run razuberi.go srv --config-file config.yaml`
5. Voila! REST api is running on `http://localhost:9091`.

You can use keystone for authorization and authentication. Default config use fake keystone, which provides no security.

Check gohan docs for more info.

## How to use it

With default config you can use `gohan client` to connect to the server.

1. `source gohan.rc`
2. `gohan client gpio_output create --gpio_pin {pin_no}` to create gpio resource. You can specify optional values like `--name` or `--is_up`.
3. `gohan client gpio_output set --is_up true {pin_uuid or name}` to set HIGH on given GPIO
3. `gohan client gpio_output set --is_up false {pin_uuid or name}` to set LOW on given GPIO

See `schema.yaml` for details about API.

You can also use CURL for direct access to API.

1. Acquire keystone token (for fake keystone this is always `admin_token`)
2. `curl -H X-Auth-Token:{keystone_token} http://localhost:9091/v1.0/gpio_outputs` - list of gpios.

Gohan API support normal actions like POST, PUT, DELETE.

If you don't need auth and don't use `gohan client` you can turn off keystone at all in config.yaml.
`gohan client` requires at least fake keystone.

Default configuration uses sqlite3 but you can use also MySQL. See [gohan docs](http://gohan.cloudwan.io/gohan/) for all posible configuration options

### Example

Full example of lighting up a LED connected to pin 26.

```
# gohan client gpio_output create --name led26 --gpio_pin 26
+-----------+--------------------------------------+
| PROPERTY  |                VALUE                 |
+-----------+--------------------------------------+
| gpio_pin  | 26                                   |
| id        | eb2ebc72-21ba-459f-8ae0-fa0acaa94c79 |
| is_up     | false                                |
| name      | led26                                |
| tenant_id | fc394f2ab2df4114bde39905f800dc57     |
+-----------+--------------------------------------+
# gohan client gpio_output set --is_up true led26
+-----------+--------------------------------------+
| PROPERTY  |                VALUE                 |
+-----------+--------------------------------------+
| gpio_pin  | 26                                   |
| id        | eb2ebc72-21ba-459f-8ae0-fa0acaa94c79 |
| is_up     | true                                 |
| name      | led26                                |
| tenant_id | fc394f2ab2df4114bde39905f800dc57     |
+-----------+--------------------------------------+
# echo 'Led is now lit up'
# gohan client gpio_output set --is_up false led26
+-----------+--------------------------------------+
| PROPERTY  |                VALUE                 |
+-----------+--------------------------------------+
| gpio_pin  | 26                                   |
| id        | eb2ebc72-21ba-459f-8ae0-fa0acaa94c79 |
| is_up     | false                                |
| name      | led26                                |
| tenant_id | fc394f2ab2df4114bde39905f800dc57     |
+-----------+--------------------------------------+
# echo 'And now it is not'
```

## TODO

- Manage also GPIO inputs (for now only outputs are available)
- Some higher level resources for most popular GPIO devices (like RGB Led or Motor).
- PWM support (not supported by `go-rpi` :( )
