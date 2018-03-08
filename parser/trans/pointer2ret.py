# from ..log import *
from ..type import Type
from ..arg import Arg
import copy
# from ..constants import Next


class Pointer2RetTransform(object):
    PointerConvert = {
        "int": "C.int",
        "uint": "C.uint",
        "int64": "C.int64",
        "uint64": "C.uint64",
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
        for key, val in Pointer2RetTransform.PointerConvert.items():
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
            if Pointer2RetTransform.PointerConvert.get(typ, None) is None:
                return typ
            else:
                return Pointer2RetTransform.PointerConvert[typ]

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
