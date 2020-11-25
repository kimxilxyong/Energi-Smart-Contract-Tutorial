#!/bin/sh
# send character-string to enable UTF-8 mode
if test ".$1" = ".off" ; then
    printf '\033%%@'
else
    printf '\033%%G'
fi
