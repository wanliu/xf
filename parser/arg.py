from log import *


class Arg(object):

    def __init__(self, name, type, source):
        self.name = name
        self.type = type
        self.source = source
        log2(self.name, self.type, self.source)

    def define(self):
        return "%s %s" % (self.name, self.type)

    def call(self):
        return "%s" % self.name

    def is_pointer(self):
        return self.type.strip()[0] == '*'
