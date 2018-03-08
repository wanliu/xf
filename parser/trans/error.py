
# class ErrorTransform(object):

#     def render_code(self, ctxt):
#         tpl = """
#             {call}
#             if ret != {status} {{
#                 {return}
#             }}
#         """
#         returns = self.trans_error_val(
#             self.return_type, "fmt.Errorf(\"{}\")".format(ctxt["error_msg"]))

#         call_render = CallTransform.render.__get__(self, CallTransform)
#         return_render = ReturnTransform.render.__get__(self, ReturnTransform)
#         log('return_type:', returns)

#         ctxt.update({
#             "call": call_render(ctxt),
#             "return": return_render(returns),
#         })
#         return tpl.format(**ctxt)

#     def header(self):
#         self.return_type = self.trans_error(self.return_type)
#         return self.super('header')

#     def trans_error(self, return_type):
#         # if self.return_type == "int"
#         return "error"

#     def kernal(self):
#         func = self.func
#         return_type = self.return_type
#         args = self.args

#         return_render = ReturnTransform.render.__get__(self, ReturnTransform)
#         returns = self.trans_return(self.return_type)
#         kernal = self.render_code({
#             "func": func,
#             "pass_args": self.trans_arguments(args),
#             "status": "C.MSP_SUCCESS",
#             "error_msg": "%s failed, error code is: %%d" % func,
#         })
#         return {
#             "kernal": kernal,
#             "default_return": return_render(returns)
#         }
