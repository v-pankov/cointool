# cointool
Simple command line tool for working with cryptocurrency coins.

It uses `coinmarketcap` service to retrieve coin information under the hood.

If not specified, `cointool` send requests to `sandbox` environment of `coinmarketcap`.

Examples below features `sandbox` environment of `coinmarketcap`.

**That's why example values doesn't look real**.

Set `--api-key` and `--api-prefix` flags to `pro` environment in order to get real values.

See [Flags](https://github.com/vdrpkv/cointool/tree/main#flags) section for more information.

You can also put flag values to config file or use environment variables instead.

See [Config](https://github.com/vdrpkv/cointool/tree/main#config) section for more information.

# Usecases

## Convert coins
```
cointool convert AMOUNT FROM TO
```

### Example

Convert `10000` `USD` to `BTC`
```
cointool convert 10000 USD BTC
38227.273628733994
```


## Get exchange rate
```
cointool rate FROM TO
```

### Example

Get exchange rate from `USD` to `BTC`
```
cointool rate USD BTC
0.7368213384961149
```

# Flags

## CoinMarketCap API Key
Set `coinmarketcap` API key. Sandbox key is shown below and used by default.
```
--api-key=b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c
```

## CoinMarketCap API Prefix
Set `coinmarketcap` API prefix. `pro` prefix is shown below and requires valid `coinmarketcap` API key to access.
```
--api-key=pro
```

## Timeout
Set maximum time to wait for command to perform. The default timeout is `7` seconds.
```
--timeout=7s
```

## Exchange rate zero value
Set zero value for exchange rate: exchange rate values received from `coinmarketcap` less or equal to zero value are discarded.
```
--erzv=0.0000000000000001
```

# Config

The flags shown in the previous section can be stored in config file or environment variables.

The default config file is present in `configs/` folder. **Copy this file to the directory containig** `cointool` **binary**.

## Config file

`cointool` reads flag values from files `config`, `config.yml` or `config.yaml` residing in the same directory as `cointool` binary if such files exist. The values read from config file are used as default flag values and can be seen in `cointool help` command.

### Example config file contents
```
api:
    key:    b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c
    prefix: sandbox

timeout: 7s

exchange_rate:
    zero_value: 0.0000000000000001

```

## Environment variables

It's also possible to pass flag values via environment variables using the follwoing environment variables mapping:

* `API_KEY` maps to `--api-key` flag.
* `API_PREFIX` maps to `--api-prefix` flag.
* `TIMEOUT` maps to `--timeout` flag.
* `EXCHANGE_RATE_ZERO_VALUE` to `--erzv` flag.

## Config priority

1. Command line flags have most priority and override config and environment variables values.
2. Enviroment variables overrides corresponding config file values.
3. Config file values have least priority among all.

# Build

Just `make`
