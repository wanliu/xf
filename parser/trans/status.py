from ..constants import Next
from ..log import *


class StatusTransform(object):
    def before(self):
        if len(self.rets) == 0:
            return {}
        elif len(self.rets) == 1 and self.rets[0].type != "int":
            return {}

        log('func', self.func)
        self.ctxt["state_val"] = "ret"
        self.ctxt["status"] = "C.MSP_SUCCESS"
        self.ctxt["return"] = "return " + \
            ', '.join(self.trans_rets(self.args))

        rets = [x.type for x in self.pointers_enum(self.args)]
        log('args:', [arg.source for arg in self.args])
        ctxt = {
            "func": self.func,
            "pass_args": self.trans_arguments(self.args),
        }

        self.ctxt["call"] = "C.{func}({pass_args})".format(**ctxt)

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

    def render_status(self, ctxt):

        tpl = """
            ret := {call}
            if {state_val} != {status} {{
                {return}
            }}
        """
        return tpl.format(**ctxt)
