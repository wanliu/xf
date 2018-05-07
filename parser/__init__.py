from arg import *
from type import *
from log import *
from constants import *
from parser import *
from wrapper import *
from trans import *


def trans_body(return_type, func, args, source=None):
    wrapper = Wrapper(return_type, func, args, source)

    wrapper.wrap(Transform)
    wrapper.wrap(TypeTransform)
    wrapper.wrap(Pointer2RetTransform)
    wrapper.wrap(StatusTransform)
    # wrapper.wrap(ReturnTransform)

    # if func not in ERROR_RETURN_EXCLUDES:
    #   wrapper.wrap(ErrorTransform)

    # if func in POINTER_RETURN_ARG:
    #   wrapper.wrap(PointerTransform)

    return wrapper.gen()


def trans_c_body(return_type, func, args, source=None):
    wrapper = Wrapper(return_type, func, args, source)

    wrapper.wrap(CTransform)

    return wrapper.gen()


def trans_h_body(return_type, func, args, source=None):
    wrapper = Wrapper(return_type, func, args, source)

    wrapper.wrap(HTransform)

    return wrapper.gen()


def trans_go_body(return_type, func, args, source=None):
    wrapper = Wrapper(return_type, func, args, source)

    wrapper.wrap(WinGoTransform)
    wrapper.wrap(WinGoTypeTransform)
    wrapper.wrap(Pointer2RetTransform)
    wrapper.wrap(WinGoStatusTransform)

    return wrapper.gen()


def trans_error(return_type, func, args):
    if func in STATUS_RETURN:
        return ("error", func, args)
    else:
        return (return_type, func, args)
