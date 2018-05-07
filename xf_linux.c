#ifdef __linux__
#include "_cgo_export.h"
#endif

int ivw_cb( const char *sessionID, int msg, int param1, int param2, const void *info, void *userData ) {
	return weakupCallback(sessionID, msg, param1, param2, info, userData);
}

ivw_ntf_handler ivwhandler = ivw_cb;