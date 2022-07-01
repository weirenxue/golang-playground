# Changed Function

This project is an attempt to understand what values will be returned by cobra's `Changed` function.

## What do we want to know?

From the comments and code snippets in the source code of the `Changed` Function

```go
// Changed returns true if the flag was explicitly set during Parse() and false
// otherwise
func (f *FlagSet) Changed(name string) bool {
    flag := f.Lookup(name)
    // If a flag doesn't exist, it wasn't changed....
    if flag == nil {
        return false
    }
    return flag.Changed
}
```

`Changed` Function will return `Changed` field of the flag if the flag exists. So we can look at the comment of the `Changed` field of `Flag` struct

> If the user set the value (or if left to default)

Obviously, this field will be true if the user sets the value of the flag. But we can still do an experiment to give use more confidence.

## Result

```sh
./app
Changed("name"): false
value: John

./app --name=May
Changed("name"): true
value: May

# set the same value as the default
./app --name=John
Changed("name"): true
value: John
```

The return value of the `Changed` function is true no matter what value the user sets (even if it is equal to the default value). If the user does not set, it will return false.
