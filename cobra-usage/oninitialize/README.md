# OnInitialize Function

This project tries to understand when the function that is an argument to cobra's `OnInitialize` function is executed.

## Result

```sh
./app --cfg ~/
in cobra.OnInitialize, cfg=/home/weirenxue/
in command PersistentPreRun
in command PreRun
in command Run
in command PostRun
in command PersistentPostRun
```

`cobra.OnInitialize` will be executed before the RUNs function of the command. The functions that are its arguments can access global variables to get flag values.
