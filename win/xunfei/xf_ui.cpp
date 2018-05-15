#include "stdafx.h"
#include "aiui/AIUI.h"
#include "aiui/AIUIType.h"
#include "xf_ui.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <iostream>

using namespace aiui;
using namespace std;

struct _Listener {
	AIUIListener* clListener;
};

struct _Event {
	IAIUIEvent *evt;
};

struct _Agent {
	IAIUIAgent* agent;
};

struct _Message {
	IAIUIMessage* message;
};

struct _Buffer {
	Buffer* buffer;
};
struct _DataBundle {
	IDataBundle *bundle;
};


class TestListener : public IAIUIListener
{
public:
	~TestListener() {}
	void onEvent(const IAIUIEvent& event) const;
};

class AIUICallbackListener: virtual public AIUIListener {
	private:
		void (*_callback)(int idx, const Event* evt);

	public:
	int index;
	~AIUICallbackListener() {}

	void onEvent(const IAIUIEvent& event) const {
		if (this->_callback) {
			Event evt;
			evt.evt = (IAIUIEvent*)&event;
			this->_callback(this->index, &evt);
		}
	}

	void setCallback(void (*callback)(int index, const Event *evt)) {
		this->_callback = callback;
	}
};

Listener* createListener(int idx, void (*fcn)(int idx, const Event*)) {
	AIUICallbackListener* li = new AIUICallbackListener();
	li->index = idx;
	li->setCallback(fcn);
    Listener* listener = new Listener();
	listener->clListener = li;
	return listener;
}

Agent* createAgent(const char *params, const Listener *listener) {
	IAIUIAgent* aiagent = IAIUIAgent::createAgent(params, listener->clListener);
	Agent* agent = new Agent();
	agent->agent = aiagent;
	return agent;
}

void sendMessage(const Agent* agent, const Message* msg) {
	agent->agent->sendMessage(msg->message);
}

Message* buildMessage(int msgType, int arg0, int arg1, const char* params, const void *buf, const int bufLength) {
	Message* msg = new Message();
	Buffer* buffer;
	if (bufLength > 0) {
		buffer = Buffer::alloc(bufLength);
		memcpy(buffer->data(), buf, bufLength);
	} else {
		buffer = 0;
	}

	msg->message = IAIUIMessage::create(msgType, arg0, arg1, params, buffer);
	return msg;
}

void destroyMessage(Message *msg) {
	msg->message->destroy();
}

int getEventType(Event *evt) {
	return evt->evt->getEventType();
}

const char* getEventInfo(Event *evt) {
	return evt->evt->getInfo();
}

int getEventArg1(Event *evt) {
	return evt->evt->getArg1();
}

int getEventArg2(Event *evt) {
	return evt->evt->getArg2();
}

DataBundle* getEventData(Event *evt) {
	DataBundle *bundle = new DataBundle();
	bundle->bundle = evt->evt->getData();
	return bundle;
}

// DataBundle* createDataBundle() {
// 	DataBundle *bundle = new DataBundle();
// 	bundle->bundle = IDataBundle::create();
// 	return bundle;
// }

void destroyDataBundle(DataBundle* bundle) {
	bundle->bundle->destroy();
}

int DataBundleRemove(DataBundle* bundle, const char *key) {
	return bundle->bundle->remove(key);
}

int DataBundlePutInt(DataBundle* bundle, const char *key, int val, int replace = 0) {
	return bundle->bundle->putInt(key, val, replace);
}

int DataBundleGetInt(DataBundle* bundle, const char *key, int defaultVal) {
	return bundle->bundle->getInt(key, defaultVal);
}

int DataBundlePutLong(DataBundle* bundle, const char *key, long val, int replace = 0) {
	return bundle->bundle->putLong(key, val, replace);
}

long DataBundleGetLong(DataBundle* bundle, const char *key, long defaultVal) {
	return bundle->bundle->getLong(key, defaultVal);
}

int DataBundlePutString(DataBundle* bundle, const char *key, const char *val, int replace = 0) {
	return bundle->bundle->putString(key, val, replace);
}

const char* DataBundleGetString(DataBundle* bundle, const char *key, const char* defaultVal) {
	return bundle->bundle->getString(key, defaultVal);
}


int DataBundlePutBinary(DataBundle* bundle, const char *key, MessageBuffer *binary, int replace = 0) {
	return bundle->bundle->putBinary(key, binary->buffer, replace);
}

MessageBuffer* DataBundleGetBinary(DataBundle* bundle, const char *key) {
	MessageBuffer* buffer = new MessageBuffer();
	buffer->buffer = bundle->bundle->getBinary(key);
	return buffer;
}

MessageBuffer* allocBuffer(int len) {
	MessageBuffer* buf = new MessageBuffer();
	buf->buffer = Buffer::alloc(len);
	return buf;
}

int deallocBuffer(MessageBuffer* buf) {
	return Buffer::dealloc(buf->buffer);
}

void* BufferGetData(MessageBuffer* buffer) {
	return buffer->buffer->data();
}

int BufferGetSize(MessageBuffer *buffer) {
	return buffer->buffer->size();
}

int xfuiCreateListener(int idx, void(*fcn)(int idx, const Event*), Listener *listener) {
	AIUICallbackListener* li = new AIUICallbackListener();
	li->index = idx;
	li->setCallback(fcn);
	listener->clListener = li;

	return 0;
}

int xfuiCreateAgent(const char *params, const Listener *listener, Agent* agent) {
	IAIUIAgent* aiagent = IAIUIAgent::createAgent(params, listener->clListener);
	agent->agent = aiagent;
	return 0;
}

int xfuiCreateAgentTest(const char *params, const Listener *listener, Agent* agent) {
	TestListener* li = new(TestListener);
	IAIUIAgent* aiagent = IAIUIAgent::createAgent(params, li);
	agent->agent = aiagent;
	return 0;
}

void xfuiSendMessage(const Agent* agent, const Message* msg) {
	sendMessage(agent, msg);
}

void xfuiBuildMessage(int msgType, int arg0, int arg1, const char* params, const void *buf, const int bufLength, Message* msg) {
	msg->message = buildMessage(msgType, arg0, arg1, params, buf, bufLength)->message;
}

void xfuiDestroyMessage(Message *msg) {
	destroyMessage(msg);
}

int xfuiGetEventType(Event *evt) {
	return getEventType(evt);
}

const char* xfuiGetEventInfo(Event *evt) {
	return getEventInfo(evt);
}

int xfuiGetEventArg1(Event *evt) {
	return getEventArg1(evt);
}

int xfuiGetEventArg2(Event *evt) {
	return getEventArg2(evt);
}

void xfuiGetEventData(Event *evt, DataBundle* data) {
	data->bundle = getEventData(evt)->bundle;
}

void xfuiDestroyDataBundle(DataBundle* bundle) {
	destroyDataBundle(bundle);
}

int xfuiDataBundleRemove(DataBundle* bundle, const char *key) {
	return DataBundleRemove(bundle, key);
}

int xfuiDataBundlePutInt(DataBundle* bundle, const char *key, int val, int replace) {
	return DataBundlePutInt(bundle, key, val, replace);
}

int xfuiDataBundleGetInt(DataBundle* bundle, const char *key, int defaultVal) {
	return DataBundleGetInt(bundle, key, defaultVal);
}

int xfuiDataBundlePutLong(DataBundle* bundle, const char *key, long val, int replace ) {
	return DataBundlePutLong(bundle, key, val, replace);
}

long xfuiDataBundleGetLong(DataBundle* bundle, const char *key, long defaultVal) {
	return DataBundleGetLong(bundle, key, defaultVal);
}

int xfuiDataBundlePutString(DataBundle* bundle, const char *key, const char *val, int replace) {
	return DataBundlePutString(bundle, key, val, replace);
}

const char* xfuiDataBundleGetString(DataBundle* bundle, const char *key, const char* defaultVal) {
	return DataBundleGetString(bundle, key, defaultVal);
}


int xfuiDataBundlePutBinary(DataBundle* bundle, const char *key, MessageBuffer *binary, int replace) {
	return DataBundlePutBinary(bundle, key, binary, replace);
}

void xfuiDataBundleGetBinary(DataBundle* bundle, const char *key, MessageBuffer* buffer) {
	buffer->buffer = bundle->bundle->getBinary(key);
}

void xfuiAllocBuffer(int len, MessageBuffer* buffer) {
	buffer->buffer = allocBuffer(len)->buffer;
}

int xfuiDeallocBuffer(MessageBuffer* buf) {
	return deallocBuffer(buf);
}

void* xfuiBufferGetData(MessageBuffer* buffer) {
	return BufferGetData(buffer);
}

int xfuiBufferGetSize(MessageBuffer *buffer) {
	return BufferGetSize(buffer);
}


/*
AIUI 事件回调接口
*/
void TestListener::onEvent(const IAIUIEvent& event) const
{
	switch (event.getEventType()) {
		/* 状态回调        */
	case AIUIConstant::EVENT_STATE:
	{
		switch (event.getArg1()) {
		case AIUIConstant::STATE_IDLE:
		{
			cout << "EVENT_STATE:" << "IDLE" << endl;
		} break;

		case AIUIConstant::STATE_READY:
		{
			cout << "EVENT_STATE:" << "READY" << endl;
		} break;

		case AIUIConstant::STATE_WORKING:
		{
			cout << "EVENT_STATE:" << "WORKING" << endl;
		} break;
		}
	} break;

	/* 唤醒事件回调        ， 唤醒成功会调用此接口，并且看到唤醒信息               */

	case AIUIConstant::EVENT_WAKEUP:
	{
		cout << "EVENT_WAKEUP:" << event.getInfo() << endl;
	} break;

	/* 休眠事件回调            */
	case AIUIConstant::EVENT_SLEEP:
	{
		cout << "EVENT_SLEEP:arg1=" << event.getArg1() << endl;
	} break;

	/* VAD事件回调，检测到前后端点  */
	case AIUIConstant::EVENT_VAD:
	{
		switch (event.getArg1()) {
		case AIUIConstant::VAD_BOS:
		{
			cout << "EVENT_VAD:" << "BOS" << endl;
		} break;

		case AIUIConstant::VAD_EOS:
		{
			cout << "EVENT_VAD:" << "EOS" << endl;
		} break;

		case AIUIConstant::VAD_VOL:
		{
			//						cout << "EVENT_VAD:" << "VOL" << endl;
		} break;
		}
	} break;

	/*
	最重要的结果事件回调，收到文本和语音语义都会返回此事件，里面有结果信息
	*/
	case AIUIConstant::EVENT_RESULT:
	{
	
		cout << "EVENT_RESULT:" << event.getEventType() << endl;
			
	} break;

	case AIUIConstant::EVENT_ERROR:
	{
		cout << "EVENT_ERROR:" << event.getArg1() << endl;
	} break;
	case AIUIConstant::EVENT_CMD_RETURN:
	{
		cout << "EVENT_CMD_RETURN: CMD " << event.getArg1() << ", Return code:" << event.getArg2() <<
			", Info:" << event.getInfo() << endl;
	} break;
	}
}
