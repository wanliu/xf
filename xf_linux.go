package xf

/*
#cgo CFLAGS:-g -Wall -I./include
#cgo LDFLAGS:-L./libs/x64 -lmsc -lrt -ldl -lpthread
#include "xf.h"

extern ivw_ntf_handler ivwhandler;

*/
import "C"
import (
	"fmt"
	"log"
	"unsafe"
)

const MSP_REC_STATUS_COMPLETE = C.MSP_REC_STATUS_COMPLETE
const MSP_EP_AFTER_SPEECH = C.MSP_EP_AFTER_SPEECH
const MSP_REC_STATUS_SUCCESS = C.MSP_REC_STATUS_SUCCESS
const MSP_AUDIO_SAMPLE_CONTINUE = C.MSP_AUDIO_SAMPLE_CONTINUE
const MSP_EP_LOOKING_FOR_SPEECH = C.MSP_EP_LOOKING_FOR_SPEECH
const MSP_AUDIO_SAMPLE_FIRST = C.MSP_AUDIO_SAMPLE_FIRST
const MSP_AUDIO_SAMPLE_LAST = C.MSP_AUDIO_SAMPLE_LAST
const MSP_TTS_FLAG_DATA_END = C.MSP_TTS_FLAG_DATA_END

func CString(str string) *C.char {
	return C.CString(str + string(0))
}

// const void* MSPAPI MSPDownloadData(const char* params, unsigned int* dataLen, int* errorCode);
func MSPDownloadData(params string) ([]byte, error) {
	var dataLen C.uint
	var errorCode C.int

	ret := C.MSPDownloadData(CString(params), &dataLen, &errorCode)
	if errorCode != C.MSP_SUCCESS {
		return nil, fmt.Errorf("MSPLogin failed, error %d", int(errorCode))
	}

	return C.GoBytes(ret, C.int(dataLen)), nil
}

// int MSPAPI MSPGetParam( const char *paramName, char *paramValue, unsigned int *valueLen );
func MSPGetParam(paramName string) (string, error) {
	var valueLen C.uint
	var paramValue [32]C.char

	ret := C.MSPGetParam(CString(paramName), &paramValue[0], &valueLen)
	if ret != C.MSP_SUCCESS {
		return "", fmt.Errorf("MSPGetParam failed, error %d", int(ret))
	}

	return C.GoStringN(&paramValue[0], C.int(valueLen)), nil
}

// const char* MSPAPI MSPGetVersion(const char *verName, int *errorCode);
func MSPGetVersion(verName string) (string, error) {
	var errorCode C.int

	ret := C.MSPGetVersion(CString(verName), &errorCode)
	if errorCode != C.MSP_SUCCESS {
		return "", fmt.Errorf("MSPGetVersion failed, error %d", int(errorCode))
	}

	return C.GoString(ret), nil
}

// int MSPAPI MSPLogin(const char* usr, const char* pwd, const char* params);
func MSPLogin(params string) error {

	ret := C.MSPLogin(nil, nil, CString(params))
	if ret != C.MSP_SUCCESS {
		return fmt.Errorf("MSPLogin failed, error %d", int(ret))
	}

	return nil
}

// int MSPAPI MSPLogout();
func MSPLogout() error {

	ret := C.MSPLogout()
	if ret != C.MSP_SUCCESS {
		return fmt.Errorf("MSPLogout failed, error %d", int(ret))
	}

	return nil
}

// int MSPAPI MSPNlpSchCancel(const char *sessionID, const char *hints);
func MSPNlpSchCancel(sessionID string, hints string) error {

	ret := C.MSPNlpSchCancel(CString(sessionID), CString(hints))
	if ret != C.MSP_SUCCESS {
		return fmt.Errorf("MSPNlpSchCancel failed, error %d", int(ret))
	}

	return nil
}

// const char* MSPAPI MSPNlpSearch(const char* params, const char* text, unsigned int textLen, int *errorCode, NLPSearchCB callback, void *userData);
func MSPNlpSearch(params string, text string, userData []byte, callback C.NLPSearchCB) (string, error) {
	var errorCode C.int
	var txt = CString(text)
	var textLen = C.strlen(txt)

	ret := C.MSPNlpSearch(CString(params), txt, C.uint(textLen), &errorCode, callback, C.CBytes(userData))
	if errorCode != C.MSP_SUCCESS {
		return "", fmt.Errorf("MSPNlpSearch failed, error %d", int(errorCode))
	}

	return C.GoString(ret), nil
}

// int MSPAPI MSPRegisterNotify( msp_status_ntf_handler statusCb, void *userData );
func MSPRegisterNotify(statusCb C.msp_status_ntf_handler, userData []byte) error {

	ret := C.MSPRegisterNotify(statusCb, C.CBytes(userData))
	if ret != C.MSP_SUCCESS {
		return fmt.Errorf("MSPRegisterNotify failed, error %d", int(ret))
	}

	return nil
}

// const char* MSPAPI MSPSearch(const char* params, const char* text, unsigned int* dataLen, int* errorCode);
func MSPSearch(params string, text string) (string, error) {
	var dataLen C.uint
	var errorCode C.int

	ret := C.MSPSearch(CString(params), CString(text), &dataLen, &errorCode)
	if errorCode != C.MSP_SUCCESS {
		return "", fmt.Errorf("MSPSearch failed, error %d", int(errorCode))
	}

	return C.GoString(ret), nil
}

// int MSPAPI MSPSetParam( const char* paramName, const char* paramValue );
func MSPSetParam(paramName string, paramValue string) error {

	ret := C.MSPSetParam(CString(paramName), CString(paramValue))
	if ret != C.MSP_SUCCESS {
		return fmt.Errorf("MSPSetParam failed, error %d", int(ret))
	}

	return nil
}

// const char* MSPAPI MSPUploadData(const char* dataName, void* data, unsigned int dataLen, const char* params, int* errorCode);
func MSPUploadData(dataName string, data []byte, params string) (string, error) {
	var errorCode C.int

	ret := C.MSPUploadData(CString(dataName), C.CBytes(data), C.uint(len(data)), C.CString(params), &errorCode)
	if errorCode != C.MSP_SUCCESS {
		return "", fmt.Errorf("MSPUploadData failed, error %d", int(errorCode))
	}

	return C.GoString(ret), nil
}

// int MSPAPI QISEAudioWrite(const char* sessionID, const void* waveData, unsigned int waveLen, int audioStatus, int *epStatus, int *Status);
func QISEAudioWrite(sessionID string, waveData []byte, audioStatus int) (int, int, error) {
	var epStatus C.int
	var Status C.int
	var waveLen = C.uint(len(waveData))

	ret := C.QISEAudioWrite(CString(sessionID), unsafe.Pointer(&waveData[0]), waveLen, C.int(audioStatus), &epStatus, &Status)
	if ret != C.MSP_SUCCESS {
		return 0, 0, fmt.Errorf("QISEAudioWrite failed, error %d", int(ret))
	}

	return int(epStatus), int(Status), nil
}

// int MSPAPI QISEGetParam(const char* sessionID, const char* paramName, char* paramValue, unsigned int* valueLen);
func QISEGetParam(sessionID string, paramName string) (string, error) {
	var valueLen C.uint
	var paramValue [64]C.char

	ret := C.QISEGetParam(CString(sessionID), CString(paramName), &paramValue[0], &valueLen)
	if ret != C.MSP_SUCCESS {
		return "", fmt.Errorf("QISEGetParam failed, error %d", int(ret))
	}

	return C.GoStringN(&paramValue[0], C.int(valueLen)), nil
}

// const char * MSPAPI QISEGetResult(const char* sessionID, unsigned int* rsltLen, int* rsltStatus, int *errorCode);
func QISEGetResult(sessionID string) (string, C.int, error) {
	var rsltLen C.uint
	var rsltStatus C.int
	var errorCode C.int

	ret := C.QISEGetResult(CString(sessionID), &rsltLen, &rsltStatus, &errorCode)
	if errorCode != C.MSP_SUCCESS {
		return "", 0, fmt.Errorf("QISEGetResult failed, error %d", int(errorCode))
	}

	return C.GoStringN(ret, C.int(rsltLen)), rsltStatus, nil
}

// const char* MSPAPI QISEResultInfo(const char* sessionID, int *errorCode);
func QISEResultInfo(sessionID string) (string, error) {
	var errorCode C.int

	ret := C.QISEResultInfo(C.CString(sessionID), &errorCode)
	if errorCode != C.MSP_SUCCESS {
		return "", fmt.Errorf("QISEResultInfo failed, error %d", int(errorCode))
	}

	return C.GoString(ret), nil
}

// const char* MSPAPI QISESessionBegin(const char* params, const char* userModelId, int* errorCode);
func QISESessionBegin(params string, userModelId string) (string, error) {
	var errorCode C.int

	ret := C.QISESessionBegin(C.CString(params), nil, &errorCode)
	if errorCode != C.MSP_SUCCESS {
		return "", fmt.Errorf("QISESessionBegin failed, error %d", int(errorCode))
	}

	return C.GoString(ret), nil
}

// int MSPAPI QISESessionEnd(const char* sessionID, const char* hints);
func QISESessionEnd(sessionID string, hints string) error {

	ret := C.QISESessionEnd(C.CString(sessionID), nil)

	if ret != C.MSP_SUCCESS {
		return fmt.Errorf("QISESessionEnd failed, error %d", int(ret))
	}

	return nil
}

// int MSPAPI QISETextPut(const char* sessionID, const char* textString, unsigned int textLen, const char* params);
func QISETextPut(sessionID string, textString string, params string) error {
	var textLen = C.strlen(C.CString(textString))

	ret := C.QISETextPut(C.CString(sessionID), C.CString(textString), C.uint(textLen), CString(params))
	if ret != C.MSP_SUCCESS {
		return fmt.Errorf("QISESessionEnd failed, error %d", int(ret))
	}

	return nil
}

// int MSPAPI QISRAudioWrite(const char* sessionID, const void* waveData, unsigned int waveLen, int audioStatus, int *epStatus, int *recogStatus);
func QISRAudioWrite(_sessionID string, waveData []byte, audioStatus int) (int, int, error) {
	var sessionID = C.CString(_sessionID)
	var epStatus C.int
	var recogStatus C.int
	var waveLen = len(waveData)
	var data unsafe.Pointer
	if waveData == nil {
		data = nil
	} else {
		data = unsafe.Pointer(&waveData[0])
	}

	ret := C.QISRAudioWrite(sessionID, data, C.uint(waveLen), C.int(audioStatus), &epStatus, &recogStatus)
	if ret != C.MSP_SUCCESS {
		return 0, 0, fmt.Errorf("QISRAudioWrite failed, error %d", int(ret))
	}

	return int(epStatus), int(recogStatus), nil
}

// int MSPAPI QISRBuildGrammar(const char *grammarType, const char *grammarContent, unsigned int grammarLength, const char *params, GrammarCallBack callback, void *userData);
func QISRBuildGrammar(grammarType string, grammarContent string, params string, userData []byte, callback C.GrammarCallBack) ([]byte, error) {
	var grammarLength = C.strlen(C.CString(grammarContent))

	ret := C.QISRBuildGrammar(C.CString(grammarType), C.CString(grammarContent), C.uint(grammarLength), C.CString(params), callback, C.CBytes(userData))
	if ret != C.MSP_SUCCESS {
		return nil, fmt.Errorf("QISRBuildGrammar failed, error %d", int(ret))
	}

	return C.GoBytes(unsafe.Pointer(&userData[0]), C.int(len(userData))), nil
}

// const char * MSPAPI QISRGetBinaryResult(const char* sessionID, unsigned int* rsltLen,int* rsltStatus, int waitTime, int *errorCode);
func QISRGetBinaryResult(sessionID string, waitTime int) (string, C.int, error) {
	var rsltLen C.uint
	var rsltStatus C.int
	var errorCode C.int

	ret := C.QISRGetBinaryResult(C.CString(sessionID), &rsltLen, &rsltStatus, C.int(waitTime), &errorCode)
	if errorCode != C.MSP_SUCCESS {
		return "", 0, fmt.Errorf("QISRGetBinaryResult failed, error %d", int(errorCode))
	}

	return C.GoStringN(ret, C.int(rsltLen)), rsltStatus, nil
}

// int MSPAPI QISRGetParam(const char* sessionID, const char* paramName, char* paramValue, unsigned int* valueLen);
func QISRGetParam(sessionID string, paramName string) (string, error) {
	var valueLen C.uint
	var paramValue [32]C.char

	ret := C.QISRGetParam(C.CString(sessionID), C.CString(paramName), &paramValue[0], &valueLen)
	if ret != C.MSP_SUCCESS {
		return "", fmt.Errorf("QISRGetParam failed, error %d", int(ret))
	}

	return C.GoStringN(&paramValue[0], C.int(valueLen)), nil
}

// const char * MSPAPI QISRGetResult(const char* sessionID, int* rsltStatus, int waitTime, int *errorCode);
func QISRGetResult(sessionID string, recStat int, waitTime int) (string, int, error) {
	var rsltStatus = C.int(recStat)
	var errorCode C.int

	ret := C.QISRGetResult(C.CString(sessionID), &rsltStatus, C.int(waitTime), &errorCode)
	// defer C.free(unsafe.Pointer(ret))

	if errorCode != C.MSP_SUCCESS {
		return "", 0, fmt.Errorf("QISRGetResult failed, error %d", int(errorCode))
	}

	return C.GoString(ret), int(rsltStatus), nil
}

// const char* MSPAPI QISRSessionBegin(const char* grammarList, const char* params, int* errorCode);
func QISRSessionBegin(grammarList string, params string) (string, error) {
	var errorCode C.int

	ret := C.QISRSessionBegin(C.CString(grammarList), C.CString(params), &errorCode)
	// defer C.free(unsaferet)

	if errorCode != C.MSP_SUCCESS {
		return "", fmt.Errorf("QISRSessionBegin failed, error %d", int(errorCode))
	}

	return C.GoString(ret), nil
}

// int MSPAPI QISRSessionEnd(const char* sessionID, const char* hints);
func QISRSessionEnd(sessionID string, hints string) error {

	ret := C.QISRSessionEnd(CString(sessionID), nil)
	if ret != C.MSP_SUCCESS {
		return fmt.Errorf("QISRSessionEnd failed, error %d", int(ret))
	}

	return nil
}

// int MSPAPI QISRSetParam(const char* sessionID, const char* paramName, const char* paramValue);
func QISRSetParam(sessionID string, paramName string, paramValue string) error {
	ret := C.QISRSetParam(C.CString(sessionID), C.CString(paramName), C.CString(paramValue))
	if ret != C.MSP_SUCCESS {
		return fmt.Errorf("QISRSetParam failed, error %d", int(ret))
	}

	return nil
}

//* int MSPAPI QISRUpdateLexicon(const char *lexiconName, const char *lexiconContent, unsigned int lexiconLength, const char *params, LexiconCallBack callback, void *userData);
func QISRUpdateLexicon(lexiconName string, lexiconContent string, params string, userData []byte, callback C.LexiconCallBack) error {
	var lexiconLength = C.strlen(C.CString(lexiconContent))

	ret := C.QISRUpdateLexicon(C.CString(lexiconName), C.CString(lexiconContent), C.uint(lexiconLength), C.CString(params), callback, C.CBytes(userData))
	if ret != C.MSP_SUCCESS {
		return fmt.Errorf("QISRUpdateLexicon failed, error %d", int(ret))
	}

	return nil
}

// const void* MSPAPI QTTSAudioGet(const char* sessionID, unsigned int* audioLen, int* synthStatus, int* errorCode);
func QTTSAudioGet(sessionID string) ([]byte, C.int, error) {
	var audioLen C.uint
	var synthStatus C.int
	var errorCode C.int

	ret := C.QTTSAudioGet(C.CString(sessionID), &audioLen, &synthStatus, &errorCode)
	if errorCode != C.MSP_SUCCESS {
		return nil, 0, fmt.Errorf("QTTSAudioGet failed, error %d", int(errorCode))
	}

	return C.GoBytes(ret, C.int(audioLen)), synthStatus, nil
}

// const char* MSPAPI QTTSAudioInfo(const char* sessionID);
func QTTSAudioInfo(sessionID string) string {
	return C.GoString(C.QTTSAudioInfo(C.CString(sessionID)))
}

// int MSPAPI QTTSGetParam(const char* sessionID, const char* paramName, char* paramValue, unsigned int* valueLen);
func QTTSGetParam(sessionID string, paramName string) (string, error) {
	var valueLen C.uint
	var paramValue [32]C.char

	ret := C.QTTSGetParam(C.CString(sessionID), C.CString(paramName), &paramValue[0], &valueLen)
	if ret != C.MSP_SUCCESS {
		return "", fmt.Errorf("QTTSGetParam failed, error %d", int(ret))
	}

	return C.GoStringN(&paramValue[0], C.int(valueLen)), nil
}

// const char* MSPAPI QTTSSessionBegin(const char* params, int* errorCode);
func QTTSSessionBegin(params string) (string, error) {
	var errorCode C.int

	ret := C.QTTSSessionBegin(C.CString(params), &errorCode)
	if errorCode != C.MSP_SUCCESS {
		return "", fmt.Errorf("QTTSSessionBegin failed, error %d", int(errorCode))
	}

	return C.GoString(ret), nil
}

// int MSPAPI QTTSSessionEnd(const char* sessionID, const char* hints);
func QTTSSessionEnd(sessionID string, hints string) error {

	ret := C.QTTSSessionEnd(C.CString(sessionID), C.CString(hints))
	if ret != C.MSP_SUCCESS {
		return fmt.Errorf("QTTSSessionEnd failed, error %d", int(ret))
	}

	return nil
}

// int MSPAPI QTTSSetParam(const char *sessionID, const char *paramName, const char *paramValue);
func QTTSSetParam(sessionID string, paramName string, paramValue string) error {

	ret := C.QTTSSetParam(C.CString(sessionID), C.CString(paramName), C.CString(paramValue))
	if ret != C.MSP_SUCCESS {
		return fmt.Errorf("QTTSSetParam failed, error %d", int(ret))
	}

	return nil
}

// int MSPAPI QTTSTextPut(const char* sessionID, const char* textString, unsigned int textLen, const char* params);
func QTTSTextPut(sessionID string, textString string, params string) error {
	var textLen = C.strlen(C.CString(textString))
	ret := C.QTTSTextPut(C.CString(sessionID), C.CString(textString), C.uint(textLen), C.CString(params))
	if ret != C.MSP_SUCCESS {
		return fmt.Errorf("QTTSTextPut failed, error %d", int(ret))
	}

	return nil
}

// int MSPAPI QIVWAudioWrite(const char *sessionID, const void *audioData, unsigned int audioLen, int audioStatus);
func QIVWAudioWrite(sessionID string, audioData []byte, audioStatus int) error {
	var audioLen = len(audioData)

	ret := C.QIVWAudioWrite(C.CString(sessionID), C.CBytes(audioData), C.uint(audioLen), C.int(audioStatus))
	if ret != C.MSP_SUCCESS {
		return fmt.Errorf("QIVWAudioWrite failed, error %d", int(ret))
	}

	return nil
}

// int MSPAPI QIVWRegisterNotify(const char *sessionID, ivw_ntf_handler msgProcCb, void *userData);
func QIVWRegisterNotify(sessionID string, userData []byte) error {
	cb := C.ivwhandler

	ret := C.QIVWRegisterNotify(C.CString(sessionID), cb, unsafe.Pointer(&userData[0]))
	if ret != C.MSP_SUCCESS {
		return fmt.Errorf("QIVWRegisterNotify failed, error %d", int(ret))
	}

	return nil
}

// int MSPAPI QIVWResMerge(const char *srcPath, const char *destPath, const char *params);
// func QIVWResMerge(srcPath string, destPath string, params string) int {

// 	ret := C.QIVWResMerge(C.CString(srcPath), C.CString(destPath), C.CString(params))
// 	if ret != C.MSP_SUCCESS {
// 		return
// 	}

// 	return C.QIVWResMerge(C.CString(srcPath), C.CString(destPath), C.CString(params))
// }

// typedef int( *ivw_ntf_handler)( const char *sessionID, int msg, int param1, int param2, const void *info, void *userData );

//export weakupCallback
func weakupCallback(sessionID unsafe.Pointer, msg, param1, param2 C.int, info, userData unsafe.Pointer) int {
	sessId := C.GoString((*C.char)(sessionID))
	log.Printf("sessionID %s 唤醒", sessId)
	return 0
}

// const char* MSPAPI QIVWSessionBegin(const char *grammarList, const char *params, int *errorCode);
func QIVWSessionBegin(grammarList string, params string) (string, error) {
	var errorCode C.int

	ret := C.QIVWSessionBegin(C.CString(grammarList), C.CString(params), &errorCode)
	if errorCode != C.MSP_SUCCESS {
		return "", fmt.Errorf("QIVWSessionBegin failed, error %d", int(errorCode))

	}

	return C.GoString(ret), nil
}

// int MSPAPI QIVWSessionEnd(const char *sessionID, const char *hints);
func QIVWSessionEnd(sessionID string, hints string) error {

	ret := C.QIVWSessionEnd(C.CString(sessionID), C.CString(hints))
	if ret != C.MSP_SUCCESS {
		return fmt.Errorf("QIVWSessionEnd failed, error %d", int(ret))
	}

	return nil
}
