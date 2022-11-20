#!/bin/bash
assert() {
    expected="$1"
    input="$2"

    ./ccompiler-bin "$input" > tmp.s || exit
    gcc -static -o tmp tmp.s
    ./tmp
    actual="$?"

    if [ "$actual" = "$expected" ]; then
        echo "$input => $actual"
    else
        echo "$input => $expected expected, but got $actual"
        exit 1
    fi
}

assert 0 0
assert 42 42
assert 50 '5+20+25-20+20'
assert 20 '5-5-10+30'
assert 20 ' 5 + 15 - 5 + 5  '
assert 10 '10 - 5 + 6 -1 '

echo OK