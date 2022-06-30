# Lookup Function

This project is an attempt to understand what values will be returned by cobra's `Lookup` function.

## What do we want to know?

From the comments in the source code

> Lookup returns the Flag structure of the named flag, returning nil if none exists.

What does the non-existent mean exactly? Is it not set by the user, or is it not defined by using `Flags` or `PersistentFlags`? Let's do an experiment.

## Result

```sh
./app 
Lookup("the-only-flag"): &pflag.Flag{Name:"the-only-flag", Shorthand:"", Usage:"Your string slice", Value:(*pflag.stringSliceValue)(0xc000118670), DefValue:"[]", Changed:false, NoOptDefVal:"", Deprecated:"", Hidden:false, ShorthandDeprecated:"", Annotations:map[string][]string(nil)}
Lookup("not-exist-flag"): (*pflag.Flag)(nil)
```

If we look up flags that are not defined by `Flags` or `PersistentFlags`, it will return `nil`. It will return a pointer to the flag regardless of whether the user has set the flag in the terminal or not, as long as the flags is defined by `Flags` or `PersistentFlags`.
