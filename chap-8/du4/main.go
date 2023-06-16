package main

/*
Cancellation

Sometimes we need to instruct a goroutine to stop what it is
doing, for example, in a web server performing a computation
on behalf of a client that has disconnected.

There is no way for one goroutine to terminate another
directly, since that would leave all its shared variables in
undefined states. In the rocket launch program (§8.7) we
sent a single value on a channel named abort, which the
countdown goroutine interpreted as a request to stop itself.
But what if we need to cancel two goroutines, or an
arbitrary number?

One possibility might be to send as many events on the abort
channel as there are goroutines to cancel. If some of the
goroutines have already terminated themselves, however, our
count will be too large, and our sends will get stuck. On
the other hand, if those goroutines have spawned other
goroutines, our count will be too small, and some goroutines
will remain unaware of the cancellation. In general, it’s
hard to know how many goroutines are working on our behalf
at any given moment. Moreover, when a goroutine receives a
value from the abort channel, it consumes that value so that
other goroutines won’t see it. For cancellation, what we
need is a reliable mechanism to broadcast an event over a
channel so that many goroutines can see it as it occurs and
can later see that it has occurred.

Recall that after a channel has been closed and drained of
all sent values, subsequent receive operations proceed
immediately, yielding zero values. We can exploit this to
create a broadcast mechanism: don’t send a value on the
channel, close it.

We can add cancellation to the du program from the previous
section with a few simple changes. First, we create a
cancellation channel on which no values are ever sent, but
whose closure indicates that it is time for the program to
stop what it is doing. We also define a utility function,
cancelled, that checks or polls the cancellation state at
the instant it is called.
*/

func main() {

}