package xf

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	MSP_SUCCESS = 0
)

var (
	xunfei       = syscall.MustLoadDLL("xunfei.dll")
	procMSPLogin = xunfei.MustFindProc("xfMSPLogin")
)

// // const void* MSPAPI MSPDownloadData(const char* params, unsigned int* dataLen, int* errorCode);
// func MSPDownloadData(params string) (string, uint, int) {
// 	var dataLen C.uint
// 	var errorCode C.int

// 	r1, r2, err := procMSPDownloadData.Call(
// 		uintptr(unsafe.Pointer(params)),
// 		uintptr(unsafe.Pointer(&dataLen)),
// 		uintptr(unsafe.Pointer(&errorCode)))
// 	if r1 != C.MSP_SUCCESS {
// 		return dataLen, errorCode
// 	}

// 	return procMSPDownloadData.Call(uintptr(unsafe.Pointer(params)),
// 		uintptr(unsafe.Pointer(&dataLen)),
// 		uintptr(unsafe.Pointer(&errorCode)))
// }

// // int MSPAPI MSPGetParam( const char *paramName, char *paramValue, unsigned int *valueLen );
// func MSPGetParam(paramName string, paramValue string) (int, uint) {
// 	var valueLen C.uint

// 	r1, r2, err := procMSPGetParam.Call(
// 		uintptr(unsafe.Pointer(paramName)),
// 		uintptr(unsafe.Pointer(paramValue)),
// 		uintptr(unsafe.Pointer(&valueLen)))
// 	if r1 != C.MSP_SUCCESS {
// 		return valueLen
// 	}

// 	return procMSPGetParam.Call(uintptr(unsafe.Pointer(paramName)),
// 		uintptr(unsafe.Pointer(paramValue)),
// 		uintptr(unsafe.Pointer(&valueLen)))
// }

// // const char* MSPAPI MSPGetVersion(const char *verName, int *errorCode);
// func MSPGetVersion(verName string) (string, int) {
// 	var errorCode C.int

// 	r1, r2, err := procMSPGetVersion.Call(
// 		uintptr(unsafe.Pointer(verName)),
// 		uintptr(unsafe.Pointer(&errorCode)))
// 	if r1 != C.MSP_SUCCESS {
// 		return errorCode
// 	}

// 	return procMSPGetVersion.Call(uintptr(unsafe.Pointer(verName)),
// 		uintptr(unsafe.Pointer(&errorCode)))
// }

// int MSPAPI MSPLogin(const char* usr, const char* pwd, const char* params);
func MSPLogin(params string) error {
	p := []byte(params + "0x00")
	r1, _, _ := procMSPLogin.Call(
		uintptr(0),
		uintptr(0),
		uintptr(unsafe.Pointer(&p[0])))
	if r1 != MSP_SUCCESS {
		return fmt.Errorf("MSPLogin failed, error %d", int(r1))
	}

	return nil
}

// // int MSPAPI MSPLogout();
// func MSPLogout() int {

// 	r1, r2, err := procMSPLogout.Call()
// 	if r1 != C.MSP_SUCCESS {
// 		return
// 	}

// 	return procMSPLogout.Call()
// }

// // int MSPAPI MSPNlpSchCancel(const char *sessionID, const char *hints);
// func MSPNlpSchCancel(sessionID string, hints string) int {

// 	r1, r2, err := procMSPNlpSchCancel.Call(
// 		uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(hints)))
// 	if r1 != C.MSP_SUCCESS {
// 		return
// 	}

// 	return procMSPNlpSchCancel.Call(uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(hints)))
// }

// // const char* MSPAPI MSPNlpSearch(const char* params, const char* text, unsigned int textLen, int *errorCode, NLPSearchCB callback, void *userData);
// func MSPNlpSearch(params string, text string, textLen uint, userData []byte) (string, int) {
// 	var errorCode C.int

// 	r1, r2, err := procMSPNlpSearch.Call(
// 		uintptr(unsafe.Pointer(params)),
// 		uintptr(unsafe.Pointer(text)),
// 		textLen,
// 		uintptr(unsafe.Pointer(&errorCode)),
// 		uintptr(unsafe.Pointer(userData)))
// 	if r1 != C.MSP_SUCCESS {
// 		return errorCode
// 	}

// 	return procMSPNlpSearch.Call(uintptr(unsafe.Pointer(params)),
// 		uintptr(unsafe.Pointer(text)),
// 		textLen,
// 		uintptr(unsafe.Pointer(&errorCode)),
// 		uintptr(unsafe.Pointer(userData)))
// }

// // int MSPAPI MSPRegisterNotify( msp_status_ntf_handler statusCb, void *userData );
// func MSPRegisterNotify(userData []byte) int {

// 	r1, r2, err := procMSPRegisterNotify.Call(
// 		uintptr(unsafe.Pointer(userData)))
// 	if r1 != C.MSP_SUCCESS {
// 		return
// 	}

// 	return procMSPRegisterNotify.Call(uintptr(unsafe.Pointer(userData)))
// }

// // const char* MSPAPI MSPSearch(const char* params, const char* text, unsigned int* dataLen, int* errorCode);
// func MSPSearch(params string, text string) (string, uint, int) {
// 	var dataLen C.uint
// 	var errorCode C.int

// 	r1, r2, err := procMSPSearch.Call(
// 		uintptr(unsafe.Pointer(params)),
// 		uintptr(unsafe.Pointer(text)),
// 		uintptr(unsafe.Pointer(&dataLen)),
// 		uintptr(unsafe.Pointer(&errorCode)))
// 	if r1 != C.MSP_SUCCESS {
// 		return dataLen, errorCode
// 	}

// 	return procMSPSearch.Call(uintptr(unsafe.Pointer(params)),
// 		uintptr(unsafe.Pointer(text)),
// 		uintptr(unsafe.Pointer(&dataLen)),
// 		uintptr(unsafe.Pointer(&errorCode)))
// }

// // int MSPAPI MSPSetParam( const char* paramName, const char* paramValue );
// func MSPSetParam(paramName string, paramValue string) int {

// 	r1, r2, err := procMSPSetParam.Call(
// 		uintptr(unsafe.Pointer(paramName)),
// 		uintptr(unsafe.Pointer(paramValue)))
// 	if r1 != C.MSP_SUCCESS {
// 		return
// 	}

// 	return procMSPSetParam.Call(uintptr(unsafe.Pointer(paramName)),
// 		uintptr(unsafe.Pointer(paramValue)))
// }

// // const char* MSPAPI MSPUploadData(const char* dataName, void* data, unsigned int dataLen, const char* params, int* errorCode);
// func MSPUploadData(dataName string, data []byte, dataLen uint, params string) (string, int) {
// 	var errorCode C.int

// 	r1, r2, err := procMSPUploadData.Call(
// 		uintptr(unsafe.Pointer(dataName)),
// 		uintptr(unsafe.Pointer(data)),
// 		dataLen,
// 		uintptr(unsafe.Pointer(params)),
// 		uintptr(unsafe.Pointer(&errorCode)))
// 	if r1 != C.MSP_SUCCESS {
// 		return errorCode
// 	}

// 	return procMSPUploadData.Call(uintptr(unsafe.Pointer(dataName)),
// 		uintptr(unsafe.Pointer(data)),
// 		dataLen,
// 		uintptr(unsafe.Pointer(params)),
// 		uintptr(unsafe.Pointer(&errorCode)))
// }

// // int MSPAPI QISEAudioWrite(const char* sessionID, const void* waveData, unsigned int waveLen, int audioStatus, int *epStatus, int *Status);
// func QISEAudioWrite(sessionID string, waveData string, waveLen uint, audioStatus int) (int, int, int) {
// 	var epStatus C.int
// 	var Status C.int

// 	r1, r2, err := procQISEAudioWrite.Call(
// 		uintptr(unsafe.Pointer(sessionID)),
// 		waveData,
// 		waveLen,
// 		audioStatus,
// 		uintptr(unsafe.Pointer(&epStatus)),
// 		uintptr(unsafe.Pointer(&Status)))
// 	if r1 != C.MSP_SUCCESS {
// 		return epStatus, Status
// 	}

// 	return procQISEAudioWrite.Call(uintptr(unsafe.Pointer(sessionID)),
// 		waveData,
// 		waveLen,
// 		audioStatus,
// 		uintptr(unsafe.Pointer(&epStatus)),
// 		uintptr(unsafe.Pointer(&Status)))
// }

// // int MSPAPI QISEGetParam(const char* sessionID, const char* paramName, char* paramValue, unsigned int* valueLen);
// func QISEGetParam(sessionID string, paramName string, paramValue string) (int, uint) {
// 	var valueLen C.uint

// 	r1, r2, err := procQISEGetParam.Call(
// 		uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(paramName)),
// 		uintptr(unsafe.Pointer(paramValue)),
// 		uintptr(unsafe.Pointer(&valueLen)))
// 	if r1 != C.MSP_SUCCESS {
// 		return valueLen
// 	}

// 	return procQISEGetParam.Call(uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(paramName)),
// 		uintptr(unsafe.Pointer(paramValue)),
// 		uintptr(unsafe.Pointer(&valueLen)))
// }

// // const char * MSPAPI QISEGetResult(const char* sessionID, unsigned int* rsltLen, int* rsltStatus, int *errorCode);
// func QISEGetResult(sessionID string) (string, uint, int, int) {
// 	var rsltLen C.uint
// 	var rsltStatus C.int
// 	var errorCode C.int

// 	r1, r2, err := procQISEGetResult.Call(
// 		uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(&rsltLen)),
// 		uintptr(unsafe.Pointer(&rsltStatus)),
// 		uintptr(unsafe.Pointer(&errorCode)))
// 	if r1 != C.MSP_SUCCESS {
// 		return rsltLen, rsltStatus, errorCode
// 	}

// 	return procQISEGetResult.Call(uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(&rsltLen)),
// 		uintptr(unsafe.Pointer(&rsltStatus)),
// 		uintptr(unsafe.Pointer(&errorCode)))
// }

// // const char* MSPAPI QISEResultInfo(const char* sessionID, int *errorCode);
// func QISEResultInfo(sessionID string) (string, int) {
// 	var errorCode C.int

// 	r1, r2, err := procQISEResultInfo.Call(
// 		uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(&errorCode)))
// 	if r1 != C.MSP_SUCCESS {
// 		return errorCode
// 	}

// 	return procQISEResultInfo.Call(uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(&errorCode)))
// }

// // const char* MSPAPI QISESessionBegin(const char* params, const char* userModelId, int* errorCode);
// func QISESessionBegin(params string, userModelId string) (string, int) {
// 	var errorCode C.int

// 	r1, r2, err := procQISESessionBegin.Call(
// 		uintptr(unsafe.Pointer(params)),
// 		uintptr(unsafe.Pointer(userModelId)),
// 		uintptr(unsafe.Pointer(&errorCode)))
// 	if r1 != C.MSP_SUCCESS {
// 		return errorCode
// 	}

// 	return procQISESessionBegin.Call(uintptr(unsafe.Pointer(params)),
// 		uintptr(unsafe.Pointer(userModelId)),
// 		uintptr(unsafe.Pointer(&errorCode)))
// }

// // int MSPAPI QISESessionEnd(const char* sessionID, const char* hints);
// func QISESessionEnd(sessionID string, hints string) int {

// 	r1, r2, err := procQISESessionEnd.Call(
// 		uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(hints)))
// 	if r1 != C.MSP_SUCCESS {
// 		return
// 	}

// 	return procQISESessionEnd.Call(uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(hints)))
// }

// // int MSPAPI QISETextPut(const char* sessionID, const char* textString, unsigned int textLen, const char* params);
// func QISETextPut(sessionID string, textString string, textLen uint, params string) int {

// 	r1, r2, err := procQISETextPut.Call(
// 		uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(textString)),
// 		textLen,
// 		uintptr(unsafe.Pointer(params)))
// 	if r1 != C.MSP_SUCCESS {
// 		return
// 	}

// 	return procQISETextPut.Call(uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(textString)),
// 		textLen,
// 		uintptr(unsafe.Pointer(params)))
// }

// // int MSPAPI QISRAudioWrite(const char* sessionID, const void* waveData, unsigned int waveLen, int audioStatus, int *epStatus, int *recogStatus);
// func QISRAudioWrite(sessionID string, waveData string, waveLen uint, audioStatus int) (int, int, int) {
// 	var epStatus C.int
// 	var recogStatus C.int

// 	r1, r2, err := procQISRAudioWrite.Call(
// 		uintptr(unsafe.Pointer(sessionID)),
// 		waveData,
// 		waveLen,
// 		audioStatus,
// 		uintptr(unsafe.Pointer(&epStatus)),
// 		uintptr(unsafe.Pointer(&recogStatus)))
// 	if r1 != C.MSP_SUCCESS {
// 		return epStatus, recogStatus
// 	}

// 	return procQISRAudioWrite.Call(uintptr(unsafe.Pointer(sessionID)),
// 		waveData,
// 		waveLen,
// 		audioStatus,
// 		uintptr(unsafe.Pointer(&epStatus)),
// 		uintptr(unsafe.Pointer(&recogStatus)))
// }

// // int MSPAPI QISRBuildGrammar(const char *grammarType, const char *grammarContent, unsigned int grammarLength, const char *params, GrammarCallBack callback, void *userData);
// func QISRBuildGrammar(grammarType string, grammarContent string, grammarLength uint, params string, userData []byte) int {

// 	r1, r2, err := procQISRBuildGrammar.Call(
// 		uintptr(unsafe.Pointer(grammarType)),
// 		uintptr(unsafe.Pointer(grammarContent)),
// 		grammarLength,
// 		uintptr(unsafe.Pointer(params)),
// 		uintptr(unsafe.Pointer(userData)))
// 	if r1 != C.MSP_SUCCESS {
// 		return
// 	}

// 	return procQISRBuildGrammar.Call(uintptr(unsafe.Pointer(grammarType)),
// 		uintptr(unsafe.Pointer(grammarContent)),
// 		grammarLength,
// 		uintptr(unsafe.Pointer(params)),
// 		uintptr(unsafe.Pointer(userData)))
// }

// // const char * MSPAPI QISRGetBinaryResult(const char* sessionID, unsigned int* rsltLen,int* rsltStatus, int waitTime, int *errorCode);
// func QISRGetBinaryResult(sessionID string, waitTime int) (string, uint, int, int) {
// 	var rsltLen C.uint
// 	var rsltStatus C.int
// 	var errorCode C.int

// 	r1, r2, err := procQISRGetBinaryResult.Call(
// 		uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(&rsltLen)),
// 		uintptr(unsafe.Pointer(&rsltStatus)),
// 		waitTime,
// 		uintptr(unsafe.Pointer(&errorCode)))
// 	if r1 != C.MSP_SUCCESS {
// 		return rsltLen, rsltStatus, errorCode
// 	}

// 	return procQISRGetBinaryResult.Call(uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(&rsltLen)),
// 		uintptr(unsafe.Pointer(&rsltStatus)),
// 		waitTime,
// 		uintptr(unsafe.Pointer(&errorCode)))
// }

// // int MSPAPI QISRGetParam(const char* sessionID, const char* paramName, char* paramValue, unsigned int* valueLen);
// func QISRGetParam(sessionID string, paramName string, paramValue string) (int, uint) {
// 	var valueLen C.uint

// 	r1, r2, err := procQISRGetParam.Call(
// 		uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(paramName)),
// 		uintptr(unsafe.Pointer(paramValue)),
// 		uintptr(unsafe.Pointer(&valueLen)))
// 	if r1 != C.MSP_SUCCESS {
// 		return valueLen
// 	}

// 	return procQISRGetParam.Call(uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(paramName)),
// 		uintptr(unsafe.Pointer(paramValue)),
// 		uintptr(unsafe.Pointer(&valueLen)))
// }

// // const char * MSPAPI QISRGetResult(const char* sessionID, int* rsltStatus, int waitTime, int *errorCode);
// func QISRGetResult(sessionID string, waitTime int) (string, int, int) {
// 	var rsltStatus C.int
// 	var errorCode C.int

// 	r1, r2, err := procQISRGetResult.Call(
// 		uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(&rsltStatus)),
// 		waitTime,
// 		uintptr(unsafe.Pointer(&errorCode)))
// 	if r1 != C.MSP_SUCCESS {
// 		return rsltStatus, errorCode
// 	}

// 	return procQISRGetResult.Call(uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(&rsltStatus)),
// 		waitTime,
// 		uintptr(unsafe.Pointer(&errorCode)))
// }

// // const char* MSPAPI QISRSessionBegin(const char* grammarList, const char* params, int* errorCode);
// func QISRSessionBegin(grammarList string, params string) (string, int) {
// 	var errorCode C.int

// 	r1, r2, err := procQISRSessionBegin.Call(
// 		uintptr(unsafe.Pointer(grammarList)),
// 		uintptr(unsafe.Pointer(params)),
// 		uintptr(unsafe.Pointer(&errorCode)))
// 	if r1 != C.MSP_SUCCESS {
// 		return errorCode
// 	}

// 	return procQISRSessionBegin.Call(uintptr(unsafe.Pointer(grammarList)),
// 		uintptr(unsafe.Pointer(params)),
// 		uintptr(unsafe.Pointer(&errorCode)))
// }

// // int MSPAPI QISRSessionEnd(const char* sessionID, const char* hints);
// func QISRSessionEnd(sessionID string, hints string) int {

// 	r1, r2, err := procQISRSessionEnd.Call(
// 		uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(hints)))
// 	if r1 != C.MSP_SUCCESS {
// 		return
// 	}

// 	return procQISRSessionEnd.Call(uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(hints)))
// }

// // int MSPAPI QISRSetParam(const char* sessionID, const char* paramName, const char* paramValue);
// func QISRSetParam(sessionID string, paramName string, paramValue string) int {

// 	r1, r2, err := procQISRSetParam.Call(
// 		uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(paramName)),
// 		uintptr(unsafe.Pointer(paramValue)))
// 	if r1 != C.MSP_SUCCESS {
// 		return
// 	}

// 	return procQISRSetParam.Call(uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(paramName)),
// 		uintptr(unsafe.Pointer(paramValue)))
// }

// // int MSPAPI QISRUpdateLexicon(const char *lexiconName, const char *lexiconContent, unsigned int lexiconLength, const char *params, LexiconCallBack callback, void *userData);
// func QISRUpdateLexicon(lexiconName string, lexiconContent string, lexiconLength uint, params string, userData []byte) int {

// 	r1, r2, err := procQISRUpdateLexicon.Call(
// 		uintptr(unsafe.Pointer(lexiconName)),
// 		uintptr(unsafe.Pointer(lexiconContent)),
// 		lexiconLength,
// 		uintptr(unsafe.Pointer(params)),
// 		uintptr(unsafe.Pointer(userData)))
// 	if r1 != C.MSP_SUCCESS {
// 		return
// 	}

// 	return procQISRUpdateLexicon.Call(uintptr(unsafe.Pointer(lexiconName)),
// 		uintptr(unsafe.Pointer(lexiconContent)),
// 		lexiconLength,
// 		uintptr(unsafe.Pointer(params)),
// 		uintptr(unsafe.Pointer(userData)))
// }

// // int MSPAPI QIVWAudioWrite(const char *sessionID, const void *audioData, unsigned int audioLen, int audioStatus);
// func QIVWAudioWrite(sessionID string, audioData string, audioLen uint, audioStatus int) int {

// 	r1, r2, err := procQIVWAudioWrite.Call(
// 		uintptr(unsafe.Pointer(sessionID)),
// 		audioData,
// 		audioLen,
// 		audioStatus)
// 	if r1 != C.MSP_SUCCESS {
// 		return
// 	}

// 	return procQIVWAudioWrite.Call(uintptr(unsafe.Pointer(sessionID)),
// 		audioData,
// 		audioLen,
// 		audioStatus)
// }

// // int MSPAPI QIVWRegisterNotify(const char *sessionID, ivw_ntf_handler msgProcCb, void *userData);
// func QIVWRegisterNotify(sessionID string, userData []byte) int {

// 	r1, r2, err := procQIVWRegisterNotify.Call(
// 		uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(userData)))
// 	if r1 != C.MSP_SUCCESS {
// 		return
// 	}

// 	return procQIVWRegisterNotify.Call(uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(userData)))
// }

// // int MSPAPI QIVWResMerge(const char *srcPath, const char *destPath, const char *params);
// func QIVWResMerge(srcPath string, destPath string, params string) int {

// 	r1, r2, err := procQIVWResMerge.Call(
// 		uintptr(unsafe.Pointer(srcPath)),
// 		uintptr(unsafe.Pointer(destPath)),
// 		uintptr(unsafe.Pointer(params)))
// 	if r1 != C.MSP_SUCCESS {
// 		return
// 	}

// 	return procQIVWResMerge.Call(uintptr(unsafe.Pointer(srcPath)),
// 		uintptr(unsafe.Pointer(destPath)),
// 		uintptr(unsafe.Pointer(params)))
// }

// // const char* MSPAPI QIVWSessionBegin(const char *grammarList, const char *params, int *errorCode);
// func QIVWSessionBegin(grammarList string, params string) (string, int) {
// 	var errorCode C.int

// 	r1, r2, err := procQIVWSessionBegin.Call(
// 		uintptr(unsafe.Pointer(grammarList)),
// 		uintptr(unsafe.Pointer(params)),
// 		uintptr(unsafe.Pointer(&errorCode)))
// 	if r1 != C.MSP_SUCCESS {
// 		return errorCode
// 	}

// 	return procQIVWSessionBegin.Call(uintptr(unsafe.Pointer(grammarList)),
// 		uintptr(unsafe.Pointer(params)),
// 		uintptr(unsafe.Pointer(&errorCode)))
// }

// // int MSPAPI QIVWSessionEnd(const char *sessionID, const char *hints);
// func QIVWSessionEnd(sessionID string, hints string) int {

// 	r1, r2, err := procQIVWSessionEnd.Call(
// 		uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(hints)))
// 	if r1 != C.MSP_SUCCESS {
// 		return
// 	}

// 	return procQIVWSessionEnd.Call(uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(hints)))
// }

// // const void* MSPAPI QTTSAudioGet(const char* sessionID, unsigned int* audioLen, int* synthStatus, int* errorCode);
// func QTTSAudioGet(sessionID string) (string, uint, int, int) {
// 	var audioLen C.uint
// 	var synthStatus C.int
// 	var errorCode C.int

// 	r1, r2, err := procQTTSAudioGet.Call(
// 		uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(&audioLen)),
// 		uintptr(unsafe.Pointer(&synthStatus)),
// 		uintptr(unsafe.Pointer(&errorCode)))
// 	if r1 != C.MSP_SUCCESS {
// 		return audioLen, synthStatus, errorCode
// 	}

// 	return procQTTSAudioGet.Call(uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(&audioLen)),
// 		uintptr(unsafe.Pointer(&synthStatus)),
// 		uintptr(unsafe.Pointer(&errorCode)))
// }

// // const char* MSPAPI QTTSAudioInfo(const char* sessionID);
// func QTTSAudioInfo(sessionID string) string {

// 	return procQTTSAudioInfo.Call(uintptr(unsafe.Pointer(sessionID)))
// }

// // int MSPAPI QTTSGetParam(const char* sessionID, const char* paramName, char* paramValue, unsigned int* valueLen);
// func QTTSGetParam(sessionID string, paramName string, paramValue string) (int, uint) {
// 	var valueLen C.uint

// 	r1, r2, err := procQTTSGetParam.Call(
// 		uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(paramName)),
// 		uintptr(unsafe.Pointer(paramValue)),
// 		uintptr(unsafe.Pointer(&valueLen)))
// 	if r1 != C.MSP_SUCCESS {
// 		return valueLen
// 	}

// 	return procQTTSGetParam.Call(uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(paramName)),
// 		uintptr(unsafe.Pointer(paramValue)),
// 		uintptr(unsafe.Pointer(&valueLen)))
// }

// // const char* MSPAPI QTTSSessionBegin(const char* params, int* errorCode);
// func QTTSSessionBegin(params string) (string, int) {
// 	var errorCode C.int

// 	r1, r2, err := procQTTSSessionBegin.Call(
// 		uintptr(unsafe.Pointer(params)),
// 		uintptr(unsafe.Pointer(&errorCode)))
// 	if r1 != C.MSP_SUCCESS {
// 		return errorCode
// 	}

// 	return procQTTSSessionBegin.Call(uintptr(unsafe.Pointer(params)),
// 		uintptr(unsafe.Pointer(&errorCode)))
// }

// // int MSPAPI QTTSSessionEnd(const char* sessionID, const char* hints);
// func QTTSSessionEnd(sessionID string, hints string) int {

// 	r1, r2, err := procQTTSSessionEnd.Call(
// 		uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(hints)))
// 	if r1 != C.MSP_SUCCESS {
// 		return
// 	}

// 	return procQTTSSessionEnd.Call(uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(hints)))
// }

// // int MSPAPI QTTSSetParam(const char *sessionID, const char *paramName, const char *paramValue);
// func QTTSSetParam(sessionID string, paramName string, paramValue string) int {

// 	r1, r2, err := procQTTSSetParam.Call(
// 		uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(paramName)),
// 		uintptr(unsafe.Pointer(paramValue)))
// 	if r1 != C.MSP_SUCCESS {
// 		return
// 	}

// 	return procQTTSSetParam.Call(uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(paramName)),
// 		uintptr(unsafe.Pointer(paramValue)))
// }

// // int MSPAPI QTTSTextPut(const char* sessionID, const char* textString, unsigned int textLen, const char* params);
// func QTTSTextPut(sessionID string, textString string, textLen uint, params string) int {

// 	r1, r2, err := procQTTSTextPut.Call(
// 		uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(textString)),
// 		textLen,
// 		uintptr(unsafe.Pointer(params)))
// 	if r1 != C.MSP_SUCCESS {
// 		return
// 	}

// 	return procQTTSTextPut.Call(uintptr(unsafe.Pointer(sessionID)),
// 		uintptr(unsafe.Pointer(textString)),
// 		textLen,
// 		uintptr(unsafe.Pointer(params)))
// }
