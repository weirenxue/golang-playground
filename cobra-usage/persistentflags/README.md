# PersistentFlags

This project is an attempt to understand how to access the flag was defined by `PersistentFlags`.

## What do we want to know?

The difference between the flags were defined by `PersistentFlags` and `Flags` is that the former can be accessed by its all subcommands. `PersistentFlags` and `Flags` are also known as global flags and local flags. But how is it actually accessed?

## Result

### app1

This application shows how to access the persistent flags defined on the root command from root command and subcommands (`sub1`, `sub2`).

```sh
# from the root command
./app1
PersistentFlags().GetString("name") => name: John, err: <nil>
Flags().GetString("name") => name: John, err: <nil>

# from the subcommand sub1
./app1 sub1
PersistentFlags().GetString("name") => name: , err: flag accessed but not defined: name
Flags().GetString("name") => name: John, err: <nil>

# from the subcommand sub2
./app1 sub2
PersistentFlags().GetString("name") => name: , err: flag accessed but not defined: name
Flags().GetString("name") => name: John, err: <nil>
```

If you access persistent flags from where they are defined (here for the root command), you can use both `PersistentFlags` and `Flags` to get the values. However, the subcommand uses only `Flags` to access the persistent flags defined on the parent.

### app2

This application shows what happens if a persistent flag with the same name already defined on the parent command is defined on the subcommand (`sub1`).

```sh
# from the root command
./app2 
PersistentFlags().GetString("name") => name: John, err: <nil>
Flags().GetString("name") => name: John, err: <nil>

# from the subcommand sub1
./app2 sub1
PersistentFlags().GetString("name") => name: May, err: <nil>
Flags().GetString("name") => name: May, err: <nil>

# from the subcommand sub1 of sub1
./app2 sub1 sub1
PersistentFlags().GetString("name") => name: , err: flag accessed but not defined: name
Flags().GetString("name") => name: May, err: <nil>

# from the subcommand sub2
./app2 sub2
PersistentFlags().GetString("name") => name: , err: flag accessed but not defined: name
Flags().GetString("name") => name: John, err: <nil>
```

The subcommand will find the persistent flags on the nearest parent.

### app3

This application shows what happens if a flag is defined on a subcommand (`sub1`) with the same name persistent flag already defined on the parent command.

```sh
# from the root command
./app3
PersistentFlags().GetString("name") => name: John, err: <nil>
Flags().GetString("name") => name: John, err: <nil>

# from the subcommand sub1
./app3 sub1
PersistentFlags().GetString("name") => name: , err: flag accessed but not defined: name
Flags().GetString("name") => name: May, err: <nil>

# from the subcommand sub1 of sub1
./app3 sub1 sub1
PersistentFlags().GetString("name") => name: , err: flag accessed but not defined: name
Flags().GetString("name") => name: John, err: <nil>

# from the subcommand sub2
./app3 sub2
PersistentFlags().GetString("name") => name: , err: flag accessed but not defined: name
Flags().GetString("name") => name: John, err: <nil>
```

`Flags`, also knows as local flags, works only for the command it is defined for. If a flag defined on the subcommand (`sub1`) is the same as a persistent flag of the same name already defined on the parent command, it will override the value of flag that flag (which only works for that subcommand).
