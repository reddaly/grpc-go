package channelz

/*
Bazel fails to add this dep, probably because the appengine build tag forces our
build to not include to files that typically import this library.

Even if it is not used, importing it ensures gazelle always adds the dependency,
which hopefully will trigger no errors.
*/
import _ "golang.org/x/sys/unix"
