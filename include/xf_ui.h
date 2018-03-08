#ifndef __XF_UI__
#ifdef __cplusplus
extern "C" {
#endif

typedef struct _Event Event;
typedef struct _Listener Listener;
typedef struct _Agent Agent;
typedef struct _Message Message;
typedef struct _Buffer MessageBuffer;
typedef struct _DataBundle DataBundle;

typedef int (*CallbackFcn)(const Event *evt);

Listener* createListener(const void *listener, void (*fcn)(const Event*));
const void* getEventHandler(Event *evt);
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

#ifdef __cplusplus    
}
#endif
#endif




