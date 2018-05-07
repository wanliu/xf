#!/usr/bin/env python
# -*- coding: UTF-8 -*-

from __future__ import print_function
from ctags import CTags, TagEntry
import sys
import traceback
from optparse import OptionParser

parser = OptionParser()
parser.add_option("-e", "--export", dest="exportfile",
                  help="only export in EXPORT FILE", metavar="FILE")

parser.add_option("-r", "--header", dest="header",
                  help="gen c++ header files", action="store_true", default=False)


parser.add_option("-g", "--golang", dest="golang",
                  help="gen golang file", action="store_true", default=False)

parser.add_option("-d", "--debug", dest="debug",
                  help="open debug log", action="store_true", default=False)

parser.add_option("--d1", "--debug1", dest="debug1",
                  help="open debug1 log", action="store_true", default=False)

parser.add_option("--d2", "--debug2", dest="debug2",
                  help="open debug2 log", action="store_true", default=False)


(options, args) = parser.parse_args()

try:
    tagFile = CTags('TAGS')
except Exception:
    sys.exit(1)

# 在命令行输入
#   ctags --c++-kinds=p --fields=+iaS --extra=+q  -R
#

from parser import *

if options.debug:
    set_level(1)
elif options.debug1:
    set_level(2)
elif options.debug2:
    set_level(3)


def in_exports(entry):
    if options.exportfile is None:
        return True

    with open(options.exportfile, 'r') as exfile:
        for line in exfile:
            if entry['name'] == line.rstrip():
                return True

        return False


def show_func_defs(entry):
    if not in_exports(entry):
        return

    line = entry['pattern'][2:-2]
    if line.find("W(") > 0:
        return

    try:
        log1(line)
        (return_type, func, args) = parse(line)
        body = trans_body(return_type, func, args, line)
        print(body)
        print()
    except Exception as e:
        print('error', e)
        traceback.print_exc()


def show_c_def_impl(entry):
    if not in_exports(entry):
        return

    line = entry['pattern'][2:-2]
    if line.find("W(") > 0:
        return

    try:
        log1(line)
        (return_type, func, args) = parse(line, ignore=False)
        body = trans_c_body(return_type, func, args, line)
        print(body)
        print()
    except Exception as e:
        print('error', e)
        traceback.print_exc()


def show_c_def(entry):
    if not in_exports(entry):
        return

    line = entry['pattern'][2:-2]
    if line.find("W(") > 0:
        return

    try:
        log1(line)
        (return_type, func, args) = parse(line, ignore=False)
        body = trans_h_body(return_type, func, args, line)
        print(body)
        print()
    except Exception as e:
        print('error', e)
        traceback.print_exc()


def show_go_defs(entry):
    if not in_exports(entry):
        return

    line = entry['pattern'][2:-2]
    if line.find("W(") > 0:
        return

    try:
        log1(line)
        (return_type, func, args) = parse(line, ignore=True)
        body = trans_go_body(return_type, func, args, line)
        print(body)
        print()
    except Exception as e:
        print('error', e)
        traceback.print_exc()


def show_go_inits(entry):
    if not in_exports(entry):
        return

    line = entry['pattern'][2:-2]
    if line.find("W(") > 0:
        return

    try:
        log1(line)
        (return_type, func, args) = parse(line, ignore=True)
        print("proc%s = xunfei.NewProc(\"xf%s\")" % (func, func))
    except Exception as e:
        print('error', e)
        traceback.print_exc()


def entities():
    tagFile = CTags('TAGS')
    entry = TagEntry()
    while tagFile.next(entry):
        yield entry


def main():
    gen_inits()

    for entry in entities():
        if options.header:
            show_c_def(entry)
        elif options.golang:
            show_go_defs(entry)
        else:
            show_c_def_impl(entry)


def gen_inits():
    if options.golang:
        print("var (")
        print("xunfei = syscall.NewLazyDLL(\"xunfei.dll\")")

    for entry in entities():
        if options.golang:
            show_go_inits(entry)

    if options.golang:
        print(")")


if __name__ == '__main__' and __package__ is None:
    # __package__ = "gen.xf.wanliu"
    main()
