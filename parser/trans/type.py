from ..constants import *
from ..log import *
from types import *


class TypeTransform(object):

    def trans_arg_call(self, arg):
        if arg is None:
            return "nil"

        ctyp = TYPE_CONVERTS.get(arg.source, None)
        log1("arg.type", arg.type)
        if ctyp is None:
            return arg.name

        return self.convert_type(ctyp, arg)

    def convert_type(self, ctyp, arg):
        if type(ctyp) is StringType:
            return ctyp.format(arg.name)
        elif type(ctyp) is FunctionType:
            return ctyp(arg)
