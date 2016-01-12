#!/usr/bin/env python

import random
import sys


def main():
    rng = 10

    if len(sys.argv) > 1:
        rng = int(sys.argv[1])

    strlist = str([random.randint(0,1000) for r in xrange(rng)])
    print(strlist.replace('[', '{').replace(']', '}'))

if __name__ == '__main__':
    main()
