# import copy
# from ..log import *


# class PointerTransform(object):

#     def before(self):
#         nargs = []

#         for i, arg, is_p in self.pointer_enum(self.args):
#             log('pointer enum', i, arg, is_p)
#             if is_p:
#                 nargs.append(arg)

#         log('nargs:', nargs, self.args)

#         def format_arg(arg):
#             typ = self.make_type(arg[1])
#             return "{} {}".format(arg[0], typ)

#         def with_enter(arg):
#             return format_arg(arg) + "\n"

#         if len(nargs) == 1:
#             define = "var " + format_arg(nargs[0])
#         elif len(nargs) > 1:
#             define = '\n'.join(["var (", map(with_enter, nargs), ")"])
#         else:
#             define = ''

#         return {
#             "before": define
#         }

#     def make_type(self, typ):
#         if typ == '*int':
#             return 'int'
#         elif typ == '*char':
#             return 'string'
#         else:
#             return typ

#     def trans_return_def(self, return_type):
#         return_type = copy.copy(return_type)
#         log('return_type:', return_type)

#         if type(return_type) is StringType:
#             return_type = [return_type]
#         elif type(return_type) is ListType:
#             pass
#         else:
#             return_type = []

#         def clean_pointer(arg):
#             if type(arg) is StringType:
#                 return self.make_type(arg)
#             elif type(arg) is TupleType:
#                 return (arg[0], self.make_type(arg[1]))
#             else:
#                 return arg.__str__()

#         pointers = map(clean_pointer, self.pointer_args())
#         return_type = pointers + return_type

#         return self.super('trans_return_def', return_type)

#     def trans_return(self, return_type):
#         return_type = copy.copy(return_type)
#         log('return:', return_type)

#         if type(return_type) is StringType:
#             return_type = [return_type]
#         elif type(return_type) is ListType:
#             pass
#         else:
#             return_type = []

#         pointers = self.pointer_args()
#         return_type = pointers + return_type

#         return self.super('trans_return', return_type)

#     def pointer_args(self):
#         args = copy.copy(self.args)
#         pointers = []

#         for i, arg, is_p in self.pointer_enum(args):
#             if is_p:
#                 pointers.append(arg)

#         return pointers

#     def trans_pointer_return(self, arg):
#         # self.return_type =
#         log("arg:", arg)
#         if arg[1] == '*int':
#             return '&' + arg[0]
#         else:
#             return arg[0]

#     def trans_define_args(self, args):
#         args = copy.copy(args)

#         for i, arg, is_p in self.pointer_enum(args):
#             if is_p:
#                 del args[i]

#         return self.super('trans_define_args', args)

#     def trans_arguments(self, args):
#         args = copy.copy(args)

#         nargs = []
#         for i, arg, is_p in self.pointer_enum(args):
#             if is_p:
#                 nargs.append(self.trans_pointer_return(arg))
#             else:
#                 nargs.append(self.trans_arg(arg))

#         return ', '.join(nargs)

#     def pointer_enum(self, args):
#         cfg = POINTER_RETURN_ARG.get(self.func, None)
#         count = 0
#         if cfg is None:
#             return

#         def fix_idx(idx):
#             return len(args) + idx if idx < 0 else idx

#         idx = cfg["idx"]
#         if type(idx) is IntType:
#             idxs = [fix_idx(idx)]
#         elif type(idx) is ListType:
#             idxs = map(fix_idx, idx)

#         for i, arg in enumerate(args):
#             if i in idxs:
#                 count += 1
#                 yield i, arg, True
#             else:
#                 yield i, arg, False
