#ifndef __XF_UI__
#include "xunfei.h"
#ifdef __cplusplus
extern "C" {
#endif

typedef struct _Event Event;
typedef struct  _Listener Listener;
typedef struct _Agent Agent;
typedef struct _Message Message;
typedef struct _Buffer MessageBuffer;
typedef struct _DataBundle DataBundle;

typedef int (*CallbackFcn)(int idx, const Event *evt);
Listener* createListener(int index, void (*fcn)(int idx, const Event*));
Agent* createAgent(const char *params, const Listener* listener);
void sendMessage(const Agent* agent, const Message* msg);
Message* buildMessage(int msgType, int arg0, int arg1, const char* params, const void *buf, const int bufLength);
void destroyMessage(Message *msg);
int getEventType(Event *evt);
const char* getEventInfo(Event *evt);
int getEventArg1(Event *evt);
int getEventArg2(Event *evt);
DataBundle* getEventData(Event *evt);
// DataBundle* createDataBundle();
void destroyDataBundle(DataBundle* bundle);
int DataBundleRemove(DataBundle* bundle, const char *key);
int DataBundlePutInt(DataBundle* bundle, const char *key, int val, int replace);
int DataBundleGetInt(DataBundle* bundle, const char *key, int defaultVal);
int DataBundlePutLong(DataBundle* bundle, const char *key, long val, int replace);
long DataBundleGetLong(DataBundle* bundle, const char *key, long defaultVal);
int DataBundlePutString(DataBundle* bundle, const char *key, const char *val, int replace);
const char* DataBundleGetString(DataBundle* bundle, const char *key, const char* defaultVal);
int DataBundlePutBinary(DataBundle* bundle, const char *key, MessageBuffer *binary, int replace);
MessageBuffer* DataBundleGetBinary(DataBundle* bundle, const char *key);

MessageBuffer* allocBuffer(int len);
int deallocBuffer(MessageBuffer* buf);
void* BufferGetData(MessageBuffer* buffer);
int BufferGetSize(MessageBuffer *buffer);

XUNFEI_API int xfuiCreateListener(int index, void(*fcn)(int idx, const Event*), Listener* listenr);
XUNFEI_API int xfuiCreateAgent(const char *params, const Listener *listener, Agent* agent);
XUNFEI_API void xfuiSendMessage(const Agent* agent, const Message* msg);
XUNFEI_API void xfuiBuildMessage(int msgType, int arg0, int arg1, const char* params, const void *buf, const int bufLength, Message* msg);
XUNFEI_API void xfuiDestroyMessage(Message *msg);

XUNFEI_API int xfuiGetEventType(Event *evt);
XUNFEI_API const char* xfuiGetEventInfo(Event *evt);
XUNFEI_API int xfuiGetEventArg1(Event *evt);
XUNFEI_API int xfuiGetEventArg2(Event *evt);
XUNFEI_API void xfuiGetEventData(Event *evt, DataBundle* data);
XUNFEI_API void xfuiDestroyDataBundle(DataBundle* bundle);
XUNFEI_API int xfuiDataBundleRemove(DataBundle* bundle, const char *key);
XUNFEI_API int xfuiDataBundleRemove(DataBundle* bundle, const char *key);
XUNFEI_API int xfuiDataBundlePutInt(DataBundle* bundle, const char *key, int val, int replace = 0);

XUNFEI_API int xfuiDataBundleGetInt(DataBundle* bundle, const char *key, int defaultVal);
XUNFEI_API int xfuiDataBundlePutLong(DataBundle* bundle, const char *key, long val, int replace = 0);
XUNFEI_API long xfuiDataBundleGetLong(DataBundle* bundle, const char *key, long defaultVal);
XUNFEI_API int xfuiDataBundlePutString(DataBundle* bundle, const char *key, const char *val, int replace = 0);
XUNFEI_API const char* xfuiDataBundleGetString(DataBundle* bundle, const char *key, const char* defaultVal);
XUNFEI_API int xfuiDataBundlePutBinary(DataBundle* bundle, const char *key, MessageBuffer *binary, int replace = 0);
XUNFEI_API void xfuiDataBundleGetBinary(DataBundle* bundle, const char *key, MessageBuffer* buffer);
XUNFEI_API void xfuiAllocBuffer(int len, MessageBuffer* buffer);
XUNFEI_API int xfuiDeallocBuffer(MessageBuffer* buf);
XUNFEI_API void* xfuiBufferGetData(MessageBuffer* buffer);
XUNFEI_API int xfuiBufferGetSize(MessageBuffer *buffer);
XUNFEI_API int xfuiCreateAgentTest(const char *params, const Listener *listener, Agent* agent);


#ifdef __cplusplus    
}


#endif
#endif




