
# class ReturnTransform(object):

# 	def trans_retval(self, returns):
# 		def ret_arg(ret):
# 			if ret == "const char *":
# 				return "C.GoString(%s)" % ret
# 			elif ret == "char *":
# 				return "C.GoString(%s)" % ret
# 			else:
# 				return ret
			
# 		return map(ret_arg, returns)
	
# 	def trans_zero(self, arg):
# 		if type(arg) == StringType:
# 			typ = arg
# 		elif type(arg) == TupleType:
# 			typ = arg[1]
# 		else:
# 			typ = arg.__str__()

# 		if typ == 'int' or type == '*int':
# 			val = "0"
# 		elif type == 'string':
# 			val = "nil"
# 		elif typ == 'error':
# 			val = "nil"
# 		else:
# 			val = "nil"
# 		return (val, typ, typ)

# 	def trans_error_val(self, return_type, msg):
# 		def set_msg(ret):
# 			if ret == "error":
# 				return (msg, "error", "error")
		
# 		return map(set_msg, self.return_enum(return_type))

# 	def render(self, rets):
# 		log1('rets:', rets)
# 		returns = self.trans_return(rets)
# 		return "return " + ','.join(returns)
