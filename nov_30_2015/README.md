How to test your program

We have a test driver here, and a sample client that writes random guesses to the driver

If you are on Linux, take the binary (`driver` and `client/solver`).
If not, you might need to re-compile the source for your OS, please contact us for further instructions.

You can also test it by hand, just run your program, then enter data according to the specs


### Example:
If you program is binary, run similar to
```
./driver_linux -c codes.txt ./client/solver
```


If your solver is in ruby:
```
./driver_linux -c codes.txt ruby solver.rb
```

### Examples
We provide the `codes.txt` and `codes.txt.2` file, each come with some example.
You can use that to test or craft your own text file (remember NOT to have the empty line in the file, if you edit it with Vim)

The driver run multiples test in parallel, so the logs (to stderr) might not be in order.
If you need to test for correctness, use `codes.txt.2` or test by hand



