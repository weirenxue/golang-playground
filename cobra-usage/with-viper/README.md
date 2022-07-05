# Cobra with Viper

This project tries to understand how the Cobra works with the Viper.

## Result

### case 1

```sh
go run app.go
in command Run: host=localhost
in command Run: viper.GetString("host")=localhost

go run app.go --host example.com
in command Run: host=example.com
in command Run: viper.GetString("host")=example.com
```

This is a case where the Viper did not get a configuration file. Its `host` configuration will match the Cobra's `host` flag, regardless of whether the user has set the `host` flag or not.

### case 2

```sh
go run app.go --cfg ./config.toml 
in command Run: host=localhost
in command Run: viper.GetString("host")=example.com

go run app.go --cfg ./config.toml --host www.example.com
in command Run: host=www.example.com
in command Run: viper.GetString("host")=www.example.com
```

In this case , the Viper can get a configuration with `host`. If the user does not set the `host` flag, the Viper's `host` value will not change and will remain the same as the value in the configuration file. And if the user sets the `host` flag, the Viper's `host` will match it.
