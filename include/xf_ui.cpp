#include "aiui/AIUI.h"
#include "aiui/AIUIType.h"
#include "xf_ui.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

using namespace aiui;

struct _Listener {
	AIUIListener* clListener;
};

struct _Event {
	IAIUIEvent *evt;
	const void* handler;
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

class AIUICallbackListener: virtual public AIUIListener {
	private:
		void (*_callback)(const Event* evt);
		const void * _handler;

	public:

	void onEvent(const IAIUIEvent& event) const {
		if (this->_callback) {
			Event evt;
			evt.evt = (IAIUIEvent*)&event;
			evt.handler = this->_handler;
			this->_callback(&evt);
		}
	}

	void setCallback(const void* handler, void (*callback)(const Event *evt)) {
		this->_callback = callback;
		this->_handler = handler;
	}
};

Listener* createListener(const void* handler, void (*fcn)(const Event*)) {
	AIUICallbackListener* li = new AIUICallbackListener();

	li->setCallback(handler, fcn);
    Listener* listener = new Listener();
	listener->clListener = li;
	return listener;
}

const void* getEventHandler(Event *evt) {
	return evt->handler;
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