# http-echo

HTTP Echo is a small go web server that serves the contents it was started with
as an HTML page.

## attribution

This code is a fork of [hashicorp/http-echo](https://github.com/hashicorp/http-echo).

The original code was written by [Seth Vargo](https://github.com/sethvargo).

## license

This code is licensed under the Mozilla Public License v2.0. See [LICENSE](./LICENSE) for more details.

## configuration

The default port is `5678`, but this is configurable via the `-listen` flag:

```bash
> http-echo -listen=:8080 -text="hello world"
```

Then visit http://localhost:8080/ in your browser.
