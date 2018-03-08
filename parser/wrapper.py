from types import *
from log import *
from constants import *


class Wrapper(object):

    def __init__(self, ret, func, args, source):

        self.ctxt = {
            "var_initial": "",
            "before": "",
            "after": "",
            "free_vars": "",
            "default_return": "",
            "kernal": "",
            "source": source
        }

        self.rets = []
        if ret is not None:
            self.rets.append(ret)

        self.func = func
        self.args = args
        self.call_chains = {}

    def wrap(self, feature):
        for key, val in feature.__dict__.items():
            if key == '__init__':
                continue

            if type(val) is FunctionType:
                log('add_feature_method %s -> %s' % (feature, key))
                # meth = val.__get__(None, feature)
                meth = val.__get__(self, feature)
                self.add_meth_chain(key, meth)
                self.wrap_method(key)

    def add_meth_chain(self, key, meth):
        chain = self.call_chains.get(key, None)
        if chain is None:
            chain = self.call_chains[key] = []

        chain.insert(0, meth)

    def wrap_method(self, key):
        bind_meth = self.__dict__.get(key, None)

        def method(self, *args, **kwargs):
            for call in self.call_chains[key]:
                ret = call(*args, **kwargs)
                if type(ret) == TupleType and ret[-1] == Next:
                    continue
                else:
                    return ret

        if bind_meth is None:
            setattr(self.__class__, key, method)
