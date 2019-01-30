### Introduction

Chronograph is a group of programs that implement functions found in a chronograph watch.

The programs are:

```
alarm - alarm clock
clock - wall clock
stopwatch - stopwatch
timer - countdown timer
```

### Quick Start

Compile:

```
$ go build alarm.go
$ go build clock.go
$ go build stopwatch.go
$ go build timer.go
```
or if you have GNU make installed:
```
$ make
```

Run:

```
$ clock
$ stopwatch
$ timer 10s echo done
$ alarm <clock_time> echo done
```

### Manual Pages

```
ALARM(1)                         User Commands                        ALARM(1)



NAME
       alarm - alarm clock

SYNOPSIS
       alarm time command

DESCRIPTION
       alarm(1) is an alarm clock. It runs a command at a specified time.

ARGUMENTS
       The first argument specifies the time the alarm clock is set to go off.

       The rest of the arguments are a command to run when the alarm triggers.

EXAMPLES
       An  alarm  set to play an mp3 (with the command aplay alarmbell.mp3) at
       2:30 pm (14:30):

       alarm 2:30pm aplay alarmbell.mp3

       An alarm set to print "hello, world" at 9 pm (21:00):

       alarm 21:00 echo "hello, world"

       Reminder to feed the cat at 10am:

       alarm 10am echo feed the cat

BUGS
       The time is always shown with a 24-hour clock, even when the  alarm  is
       set with a 12-hour clock.

       The  command must be present. Use "echo -n" as the command if you don´t
       want a command to run.

AUTHOR
       Jay Ts (http://jayts.com)

COPYRIGHT
       Copyright 2018 Jay Ts

       Released  under  the  GNU   Public   License,   version   3.0   (GPLv3)
       (http://www.gnu.org/licenses/gpl.html)



Jay Ts                           December 2018                        ALARM(1)
```

```
CLOCK(1)                         User Commands                        CLOCK(1)



NAME
       clock - clock

SYNOPSIS
       clock

DESCRIPTION
       clock(1)  is  a  command-line clock that runs in a virtual terminal. It
       displays the running wall clock time and is accurate to about 1/10 sec‐
       ond.

       Type a Control-C to exit.

ARGUMENTS
       None.

BUGS
       The time is always shown with a 24-hour clock.

AUTHOR
       Jay Ts (http://jayts.com)

COPYRIGHT
       Copyright 2018 Jay Ts

       Released   under   the   GNU   Public   License,  version  3.0  (GPLv3)
       (http://www.gnu.org/licenses/gpl.html)



Jay Ts                           December 2018                        CLOCK(1)
```

```
STOPWATCH(1)                     User Commands                    STOPWATCH(1)



NAME
       stopwatch - stopwatch

SYNOPSIS
       stopwatch [-p]

DESCRIPTION
       stopwatch(1) is a stopwatch that runs in a virtual terminal.

       It is controlled with keys on the keyboard as follows:

       SPACE, p, P

           Pause/restart the stopwatch.

       l, L

           Lap Timer. The current timing is printed, and counting continues on the following line.

       r, R

           Reset. Works only while paused.


       q, Q, e, E, Ctrl-C, Ctrl-D, Enter/Return

           Stop the stopwatch and exit.

ARGUMENTS
       When  the  -p  option  is specified, the stopwatch starts in the paused
       state.

AUTHOR
       Jay Ts (http://jayts.com)

COPYRIGHT
       Copyright 2018 Jay Ts

       Released  under  the  GNU   Public   License,   version   3.0   (GPLv3)
       (http://www.gnu.org/licenses/gpl.html)



Jay Ts                           December 2018                    STOPWATCH(1)
```

```
TIMER(1)                         User Commands                        TIMER(1)



NAME
       timer - countdown timer

SYNOPSIS
       timer duration command

DESCRIPTION
       timer(1)  is  a  countdown  timer.  It runs a command after a specified
       duration of time.

ARGUMENTS
       The first argument specifies the duration.

       The rest of the arguments are a command to run when the  timer  reaches
       0.

EXAMPLES
       After  2  minutes  and  30 seconds, play an mp3 (with the command aplay
       alarmbell.mp3):

       timer 2m30s aplay alarmbell.mp3

       An alarm set to print "hello, world" in 21 minutes:

       alarm 21m echo "hello, world"

BUGS
       The time is always shown with a 24-hour clock, even when  the  duration
       is set with a 12-hour clock.

       The  command must be present. Use "echo -n" as the command if you don´t
       want a command to run.

AUTHOR
       Jay Ts (http://jayts.com)

COPYRIGHT
       Copyright 2018 Jay Ts

       Released  under  the  GNU   Public   License,   version   3.0   (GPLv3)
       (http://www.gnu.org/licenses/gpl.html)



Jay Ts                           December 2018                        TIMER(1)
```
