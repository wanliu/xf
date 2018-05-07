import re
from types import *
from log import *
from constants import *
from arg import *
from type import *


def split_star(arg):
    if arg is None:
        return []
    arg = arg.strip()
    ws = arg.split(' ')

    for idx, word in enumerate(ws):
        if len(word) == 0:
            continue
        if word[0] == '*':
            ws[idx] = ws[idx][1:]
            ws.insert(idx, '*')
            break
        elif word[-1] == '*':
            ws[idx] = ws[idx][:-1]
            ws.insert(idx + 1, '*')
            break
        elif word.find('*') >= 0:
            pos = word.find('*')
            ws.insert(idx + 1, ws[idx][pos + 1:])
            ws.insert(idx + 1, '*')
            ws[idx] = ws[idx][:pos]
        else:
            pass

    return ws


def trans_c(arg, ignore=True):
    ws = split_star(arg)
    # ws = re.split(r'[\s|\*]', arg)
    for i in xrange(len(ws) - 1, 0, -1):
        key = ' '.join(ws[:i])
        if ignore:
            if TYPES.get(key, None) is None:
                continue
            typ = TYPES[key]
        else:
            typ = key

        if type(typ) is FunctionType:
            return Arg(ws[-1], typ(), key)
        elif type(typ) is StringType:
            return Arg(ws[-1], typ, key)

    return None


def trans_type(arg, ignore=True):
    ws = split_star(arg)
    for i in xrange(len(ws) - 1, -1, -1):
        key = ' '.join(ws[:i + 1])

        if ignore:
            if TYPES.get(key, None) is None:
                continue
            typ = TYPES[key]
        else:
            typ = key

        if type(typ) is FunctionType:
            return Type(typ(), key)
        elif type(typ) is StringType:
            return Type(typ, key)


FuncDef = '(\w+\*?)?\s+(\*?\w+)\((.*?)\)'
ConstFuncDef = '(const\s+\w+\*?)?\s+(\*?\w+)\((.*?)\)'
ConstFuncDef2 = '(const\s+\w+\s+\*?)?\s+(\*?\w+)\((.*?)\)'

fndefs = [ ("match", ConstFuncDef),  ("match", ConstFuncDef2), ("search", FuncDef) ]


def parse(line, ignore=True):
    line = line.replace("MSPAPI ", "")
    # print(line)
    def _trans_c(arg):
        return trans_c(arg, ignore)

    for mth, fndef in fndefs:
        if mth == "match":
            m = re.match(fndef, line)
        else:
            m = re.search(fndef, line)
        # print('M', m.groups())
        if m is not None:
            arg_s = m.group(3)
            func = m.group(2)
            return_type = m.group(1)
            # print(args)
            args = arg_s.split(",")
            log1('func: %s, ret: %s, args: %s' % (func, return_type, arg_s))
            if func[0] == '*':
                func = func[1:]
                return_type = return_type + '*'

            tr_args = map(_trans_c, args)
            tr_args = [x for x in tr_args if x]
            ret = trans_type(return_type, ignore)
            # m = re.search('\s+(\w+)\(', line)
            # m.group(1)

            # (return_type, func, args) = trans_error(return_type, func, tr_args)
            return (ret, func, tr_args)

    raise Exception('invalid parse')
