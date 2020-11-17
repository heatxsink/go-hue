go-hue
======
Wrapper API and cli examples in golang for interacting with lights via philips hue hub HTTP API.

demo
----
[![Alt Demo of go-hue unit tests](http://img.youtube.com/vi/3zMky_9xdJs/0.jpg)](http://www.youtube.com/watch?v=3zMky_9xdJs)

setup
-----
To install "github.com/heatxsink/go-hue" golang module.

	$ go get github.com/heatxsink/go-hue

To run the tests you'll need to set the following environment variables:

	1. HUE_TEST_USERNAME (You can obtain a whitelisted username via examples/discover.go)
	1. HUE_TEST_HOSTNAME (Your hue hub's hostname or IP address)

important note
--------------
As of API `1.31`, Philips/Hue has disabled the `deleteUser` function, and it will always return the error `unauthorized user`.

To delete an app from your bridge, you must login to https://account.meethue.com/apps and manually delete using the website.

bugs and contribution
---------------------
Please feel free to reach out. Issues and PR's are always welcome!


