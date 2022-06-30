# RUN Functions Execution Bahavior

This project attempts to understand the behavior of cobra as it executes many RUN funcions.

## What is the order of these RUN functions?

According to the comments in the cobra source code. We can know that the order of execution of the RUN functions, such as

- `PersistentPreRun()`
- `PreRun()`
- `Run()`
- `PostRun()`
- `PersistentPostRun()`

But there are many similar functions with the `E` suffix. This project tries to understand the exact order of these run functions.

As the result of this project show. We can set up RUN functions with and without the suffix `E`. Cobra prefers RUN functions with the suffix `E`, so if we set up both functions, only functions with the suffix `E` will be executed.

## What is the meaning of the suffix `E`?

The suffix `E` means that the function will return `Error`. The `Error` will be the return value of the `cobra.Command.Execute` function.

## What happend if a RUN function with an suffix `E` returns a non-nil `Error`?

Suppose `PersistentPreRunE` returns a non-nil `Error`. Then the command will exit and the rest of the RUN functions will be not executed.

This can be experimented with by uncommenting the code that returns the real `Error` in `PersistentPreRunE` or `RunE` function.
