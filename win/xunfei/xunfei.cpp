// xunfei.cpp : ���� DLL Ӧ�ó���ĵ���������
//

#include "stdafx.h"
#include "xunfei.h"
#include "xf.h"


XUNFEI_API int xfMSPLogin(const char* usr, const char* pwd, const char* params) {
	return MSPLogin(usr, pwd, params);
}
