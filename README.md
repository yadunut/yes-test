# Writing my own `yes` comamnd

This mini project was inspired by [A Little Story About the yes Unix Command](https://matthias-endler.de/2017/yes/)

My goal was to recreate the `yes` command and possibly improve its speeds.

## Base Stats

The default `yes` command has a speed of [30.0MiB/s].

`yes | pv -r > /dev/null`
