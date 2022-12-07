# Benchmark Ticker

This project is used to test the accuracy of the ticker.

## Comment in the source code

> The ticker will adjust the time interval or drop ticks to make up for slow receivers.

We can foresee that consumers will influence the accuracy of the ticker.

## Result

duration | expected number of times per second | actual number of times per second
---      | ---                       | ---
`500ms`  | 2/s                       | 1.999917/s
`20ms`   | 50/s                      | 49.996330/s
`1ms`    | 1000/s                    | 950.591541/s
`500μs`  | 2000/s                    | 970.383533/s
`20μs`   | 50000/s                   | 2252.043507/s
`1μs`    | 1000000/s                 | 362100.173044/s

## Conclusion

We should be careful about using ticker for duration-critical tasks. I suggest that if the duration of ticker is less than `20ms`, then stop using ticker and look for other solutions.
