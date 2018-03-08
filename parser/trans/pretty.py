# class PrettyTransform(object):

# 	def render(self, ctxt):
# 		return self.render_after(self.super('render', ctxt))

# 	def render_after(self, code):
# 		lines = code.split('\n')
# 		col = [x for x in lines if x.strip()]
# 		return '\n'.join(col)
