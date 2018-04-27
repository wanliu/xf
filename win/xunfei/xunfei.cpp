// xunfei.cpp : 定义 DLL 应用程序的导出函数。
//

#include "stdafx.h"
#include "xunfei.h"
#include "xf.h"


XUNFEI_API int xfMSPLogin(const char* usr, const char* pwd, const char* params) {
	return MSPLogin(usr, pwd, params);
}
