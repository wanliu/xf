from __future__ import print_function


level = 0
def set_level(l):
	global level
	level = l

def log(*args):
	if level > 0:
		print(*args)

def log1(*args):
	if level > 1:
		print(*args)

def log2(*args):
	if level > 2:
		print(*args)
