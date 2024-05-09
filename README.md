# DzForex GRPC MicroService

Just a small first project in go, implementinmg a microservice to get parallal exchange rates from http://www.forexalgerie.com/

## Example Usage with [gRPCurl](https://github.com/fullstorydev/grpcurl)

```bash
$ grpcurl --plaintext -d '' localhost:4444 DzForex/GetAvailableCurrencies

{
  "Currencies": [
    "mad",
    "tnd",
    "cad",
    "aed",
    "cad",
    "cn",
    "gbp",
    "usd",
    "chf",
    "chf",
    "sar",
    "aed",
    "mad",
    "eur",
    "usd",
    "tr",
    "tnd",
    "tr",
    "cn",
    "gbp",
    "eur",
    "sar"
  ]
}
```



```bash
$ grpcurl --plaintext -d '{"Currency" : "eur"}' localhost:4444 DzForex/GetRate

{
  "Buy": 238,
  "Sell": 240
}
```
