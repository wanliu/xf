
class Type(object):

    def __init__(self, type, source):
        self.type = type
        self.name = ""
        self.source = source

    def define(self):
        return self.type

    def is_pointer(self):
        return self.type.strip()[0] == '*'
