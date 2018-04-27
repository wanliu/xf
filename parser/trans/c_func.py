# -*- coding: UTF-8 -*-

from __future__ import print_function
from types import *
from ..log import *
from ..constants import *
from ..arg import *
from ..type import *

class CTransform(object):
    def before(self):
        return {
            "before": "",
        }

    def after(self):
        return {
            "after": "",
        }

    def header(self):
        return {
            "func": ' xf' + self.func,
            "struct": "",
            "args": self.trans_define_args(self.args),
            "return_def": 'XUNFEI_API ' + self.trans_return_def(self.rets),
        }

    def kernal(self):
        func = self.func
        rets = self.rets
        args = self.args

        default_return = "return {func}({pass_args})"
        ctxt = {
            "func": func,
            "pass_args": self.trans_arguments(args),
        }

        return {
            # "status": "C.MSP_SUCCESS",
            # "error_msg": "%s failed, error code is: %%d" % func,
            "default_return": default_return.format(**ctxt),
        }

    def run_header(self):
        self.ctxt.update(self.header())

    def run_before(self):
        self.ctxt.update(self.before())

    def run_kernal(self):
        self.ctxt.update(self.kernal())

    def run_after(self):
        self.ctxt.update(self.after())

    def gen(self):
        self.run_header()

        self.run_before()
        self.run_kernal()
        self.run_after()

        return self.render(self.ctxt)

    def super(self, name, *args, **kwargs):
        supe = Transform.__dict__[name].__get__(self, Transform)
        return supe(*args, **kwargs)

    def render(self, ctxt, **kwargs):
        tpl = """
        // {source}
        {return_def} {struct}{func}({args}) {{
            {var_initial}
            {before}
            {kernal}
            {after}
            {free_vars}
            {default_return};
        }}
        """
        return tpl.format(**ctxt)

    def trans_struct(self, struct):
        return struct

    def trans_define_args(self, args):
        return ', '.join(map(self.trans_arg_def, args))

    def trans_arguments(self, args):
        return ', '.join(map(self.trans_arg_call, args))

    # def trans_returns(self, rets):
    #     pass

    def is_pointer(self, typ):
        return typ.strip()[-1] == '*'

    def trans_arg_def(self, arg):
        return "%s %s" % (arg.source, arg.name)

    def trans_typ(self, typ):
        if typ is None:
            return "nil"

        return typ.source

    def trans_arg_call(self, arg):
        log('arg type:', type(arg))
        if arg is None:
            return "nil"

        return arg.call()

    def return_enum(self, returns):
        if returns is None:
            return
        elif type(returns) is StringType:
            returns = [returns]
        elif type(returns) is ListType:
            # returns =
            pass
        elif type(returns) is FunctionType:
            returns = returns()
        else:
            returns = [returns.__str__()]

        for ret in returns:
            yield ret

    """ trans_return_def 转换函数返回定义"""

    def trans_return_def(self, returns):
        log('trans_return_def')
        returns = [ret for ret in self.return_enum(returns)]

        if len(returns) == 0:
            return ''
        elif len(returns) == 1:
            return self.trans_typ(returns[0])
        elif len(returns) > 1:
            return '(' + ', '.join(map(self.trans_typ, returns)) + ')'

    def trans_return(self, returns):
        return map(self.trans_arg_call, self.return_enum(returns))

