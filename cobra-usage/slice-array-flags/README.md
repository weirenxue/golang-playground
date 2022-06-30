# StringSlice and StringArray Flags

This project tries to understand the difference between `StringSlice` and `StringArray`.

## What do we want to know?

`StringSlice` and `StringArray` let cli accept string array. So what exactly is the difference between them?

## Result

```sh
./app --slice="a,b" --slice="c" --array="a,b" --array="c"
GetStringSlice: []string{"a", "b", "c"}
GetStringArray: []string{"a,b", "c"}
```

`StringSlice` will split the strings with comma and treat each string as one element, or it can be assigned with each single flag. But `GetStringArray` can only be assigned by each single flag.
