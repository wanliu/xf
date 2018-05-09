# -*- coding: UTF-8 -*-

from __future__ import print_function
from types import *
from ..log import *
from ..constants import *
from ..arg import *
from ..type import *
import copy


class WinGoTransform(object):

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
            "func": ' ' + self.func,
            "struct": "",
            "args": self.trans_define_args(self.args),
            "return_def": ' ' + self.trans_return_def(self.rets),
        }

    def kernal(self):
        func = self.func
        rets = self.rets
        args = self.args

        default_return = "return proc{func}.Call({pass_args})"
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
        func{struct}{func}({args}){return_def} {{
            {var_initial}
            {before}
            {kernal}
            {after}
            {free_vars}
            {default_return}
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
        return arg.define()

    def trans_typ(self, typ):
        if typ is None:
            return "nil"

        return typ.define()

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


class WinGoTypeTransform(object):
    TYPE_CONVERTS = {
        "void *": "uintptr(unsafeString({}))",
        "char *": "uintptr(unsafeString({}))",
        "const char *": "uintptr(unsafeString({}))",
        "unsigned int *": "uintptr(unsafe.Pointer(&{}))",
        "int *": "uintptr(unsafe.Pointer(&{}))",
        "unsigned int64 *": "uintptr(unsafe.Pointer(&{}))",
        "int64 *": "uintptr(unsafe.Pointer(&{}))",
    }

    def trans_arg_call(self, arg):
        if arg is None:
            return "nil"

        ctyp = WinGoTypeTransform.TYPE_CONVERTS.get(arg.source, None)
        log1("arg.type", arg.type)
        if ctyp is None:
            return arg.name

        return self.convert_type(ctyp, arg)

    def convert_type(self, ctyp, arg):
        if type(ctyp) is StringType:
            return ctyp.format(arg.name)
        elif type(ctyp) is FunctionType:
            return ctyp(arg)


class WinGoStatusTransform(object):
    def before(self):
        if len(self.rets) == 0:
            return {}
        elif len(self.rets) == 1 and self.rets[0].type != "int":
            return {}

        log('func', self.func)
        self.ctxt["state_val"] = "r1"
        self.ctxt["status"] = "MSP_SUCCESS"
        self.ctxt["return"] = "return " + \
            ', '.join(self.trans_rets(self.args))

        rets = [x.type for x in self.pointers_enum(self.args)]
        log('args:', [arg.source for arg in self.args])
        ctxt = {
            "func": self.func,
            "pass_args": self.trans_arguments(self.args),
        }

        self.ctxt["call"] = "proc{func}.Call(\n{pass_args})".format(**ctxt)

        log('kernal:', self.ctxt["kernal"])
        self.ctxt.update({
            "kernal": self.render_status(self.ctxt)

        })
        return (self.ctxt, Next)

    def trans_rets(self, rets):
        def call_method(arg):
            log("arg:", arg, arg.type, arg.name, arg.source)
            return arg.name
        return map(call_method, self.pointers_enum(rets))

    def trans_arguments(self, args):
        return ', \n'.join(map(self.trans_syscall_arg_call, args))

    def trans_syscall_arg_call(self, arg):
        return self.trans_arg_call(arg)

    def render_status(self, ctxt):

        tpl = """
            r1, _, _ := {call}
            if {state_val} != {status} {{
                {return}
            }}
        """
        return tpl.format(**ctxt)


class WinGoPointer2RetTransform(object):
    PointerConvert = {
        "int": "int",
        "uint": "uint",
        "int64": "int64",
        "uint64": "uint64",
    }

    def before(self):
        var_initial = self.ctxt["var_initial"]
        var_initial += '\n'.join(["var " + arg.define()
                                  for arg in self.pointers_enum(self.args)])

        self.ctxt.update({
            "var_initial": var_initial
        })
        return self.ctxt

    def trans_define_args(self, args):
        args = copy.copy(args)
        args = filter(lambda x: not self.filter_typ(x), args)
        args = map(self.trans_arg_def, args)
        return ', '.join(args)

    def trans_return_def(self, rets):
        args = self.pointers_enum(self.args)
        converts = {}
        for key, val in WinGoPointer2RetTransform.PointerConvert.items():
            converts[val] = key

        def convert_typ(arg):
            if converts.get(arg.type, None) is None:
                return Type(arg.type, arg.source)
            else:
                return Type(converts[arg.type], arg.source)

        rets += map(convert_typ, args)

        if len(rets) == 0:
            return ''
        elif len(rets) == 1:
            return self.trans_typ(rets[0])
        elif len(rets) > 1:
            return '(' + ', '.join(map(self.trans_typ, rets)) + ')'

    def pointers_enum(self, args):
        args = copy.copy(args)

        def convert_type(typ):
            if WinGoPointer2RetTransform.PointerConvert.get(typ, None) is None:
                return typ
            else:
                return WinGoPointer2RetTransform.PointerConvert[typ]

        def pure_typ(arg):
            return Arg(arg.name, convert_type(arg.type[1:]), arg.source)

        typs = map(pure_typ, filter(lambda x: self.filter_typ(x), args))
        for typ in typs:
            yield typ

    def filter_typ(self, typ):
        return typ.is_pointer()

    def typs_enum(self, typs):
        for i, typ in enumerate(typs):
            yield typ
