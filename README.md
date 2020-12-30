# envtest
Simple test of getting environment variables with viper (golang)

Nothing much to see here - just a sanity check during some debugging. What happened? [viper](https://github.com/spf13/viper) is how I get environment variables into CLI in support of a [12-factor])https://12factor.net/) approach. Adding cobra subcommands seemed to break this. So I wrote this sanity-checker to isolate the issue. It boiled down to [the order in which the multiple init functions were being called](https://stackoverflow.com/a/24790378). It was more straightfoward to just move the viper calls into the command functions, which is consistent with the environemnt variables [being re-read on each  `viper.Get()`](https://github.com/spf13/viper#working-with-environment-variables).

```
Set these two environment variables, e.g
export ROOTCMD_TEST=bar
export SUBCMD_TEST=pong

------------RESULTS--------------------

main()::viper.Get("test")                   = bar
main()::go func()::viper.Get("test")        = bar
main()::func()::viper.Get("test")           = bar
main():func()::viper.SetEnvPrefix("SUBCMD")
main()::func()::viper.Get("test")           = pong
```
