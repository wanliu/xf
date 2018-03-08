Next = object()

TYPES = {
    "int": "int",
    "const char": "string",
    "void *": "[]byte",
    "char *": "string",
    "int *": "*int",
    "unsigned int": "uint",
    "unsigned int *": "*uint",
    "unsigned int64": "uint64",
    "unsigned int64 *": "*uint64",
    "const wchar_t*": "*uint16",
    "const char*": "string",
    "const char *": "string",
    "const void *": "string",
    "DownloadStatusCB": "DownloadStatusCB",
    "DownloadResultCB": "DownloadResultCB",
}

STATUS_RETURN = [
    "MSPLogin",
    "MSPLogout"
]

STATUS_CODE = {
}

ERROR_RETURN_EXCLUDES = [
    "*Wchar2Mbytes"
]

POINTER_RETURN_ARG = {
    "MSPUploadData": {
        "idx": -1,
    }
}

TYPE_CONVERTS = {
    "void *": "C.CBytes({})",
    "char *": "C.CString({})",
    "const char *": "C.CString({})",
    "unsigned int *": "&{}",
    "int *": "&{}",
    "unsigned int64 *": "&{}",
    "int64 *": "&{}",
}
