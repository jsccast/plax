# Recent changes

## New `httpserver` channel

See [`demos/http.yaml`](demos/http.yaml) and
[`demos/httpserver.yaml`](demos/httpserver.yaml) for examples.

## `httpclient` channel request and response improvements

## `@@` and `!!` syntax changes

Now do

1. `{@@filename}` instead of `@@filename` and
1. `{!!javascript!!}` instead of `!!javascript`.


## `HTTPClient` channel request and response

The request message now has lower-case properties instead of
capitalized properties.

The response message has structure supporting status code, body, and
headers.  Sadly this change is not backwards-compatible.


## Schema validation

`pub` and `recv` now support a `schema` property, which gives a URI to
a JSON Schema that's used to validate the message.


## `serialization` properties for `pub` and `recv`

You can specify how `pub` and `recv` serialize messages by giving a
`serialization` property, which should either be `JSON` (the default)
or `string`.


