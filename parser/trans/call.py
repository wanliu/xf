
# class CallTransform(object):

#     def render(self, ctxt):
#         func = self.func
#         return_type = self.return_type
#         args = self.trans_arguments(self.args)

#         ctxt.update({
#             "func": func,
#             "args": args
#         })
#         return "ret := C.{func}({args})".format(**ctxt)
