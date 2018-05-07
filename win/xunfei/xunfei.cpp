#include "stdafx.h"
#include "xunfei.h"
#include "xf.h"


        // const void* MSPAPI MSPDownloadData(const char* params, unsigned int* dataLen, int* errorCode);
        XUNFEI_API const void *  xfMSPDownloadData(const char * params, unsigned int * dataLen, int * errorCode) {
            
            
            
            
            
            return MSPDownloadData(params, dataLen, errorCode);
        }
        


        // int MSPAPI MSPGetParam( const char *paramName, char *paramValue, unsigned int *valueLen );
        XUNFEI_API int  xfMSPGetParam(const char * paramName, char * paramValue, unsigned int * valueLen) {
            
            
            
            
            
            return MSPGetParam(paramName, paramValue, valueLen);
        }
        


        // const char* MSPAPI MSPGetVersion(const char *verName, int *errorCode);
        XUNFEI_API const char *  xfMSPGetVersion(const char * verName, int * errorCode) {
            
            
            
            
            
            return MSPGetVersion(verName, errorCode);
        }
        


        // int MSPAPI MSPLogin(const char* usr, const char* pwd, const char* params);
        XUNFEI_API int  xfMSPLogin(const char * usr, const char * pwd, const char * params) {
            
            
            
            
            
            return MSPLogin(usr, pwd, params);
        }
        


        // int MSPAPI MSPLogout();
        XUNFEI_API int  xfMSPLogout() {
            
            
            
            
            
            return MSPLogout();
        }
        


        // int MSPAPI MSPNlpSchCancel(const char *sessionID, const char *hints);
        XUNFEI_API int  xfMSPNlpSchCancel(const char * sessionID, const char * hints) {
            
            
            
            
            
            return MSPNlpSchCancel(sessionID, hints);
        }
        


        // const char* MSPAPI MSPNlpSearch(const char* params, const char* text, unsigned int textLen, int *errorCode, NLPSearchCB callback, void *userData);
        XUNFEI_API const char *  xfMSPNlpSearch(const char * params, const char * text, unsigned int textLen, int * errorCode, NLPSearchCB callback, void * userData) {
            
            
            
            
            
            return MSPNlpSearch(params, text, textLen, errorCode, callback, userData);
        }
        


        // int MSPAPI MSPRegisterNotify( msp_status_ntf_handler statusCb, void *userData );
        XUNFEI_API int  xfMSPRegisterNotify(msp_status_ntf_handler statusCb, void * userData) {
            
            
            
            
            
            return MSPRegisterNotify(statusCb, userData);
        }
        


        // const char* MSPAPI MSPSearch(const char* params, const char* text, unsigned int* dataLen, int* errorCode);
        XUNFEI_API const char *  xfMSPSearch(const char * params, const char * text, unsigned int * dataLen, int * errorCode) {
            
            
            
            
            
            return MSPSearch(params, text, dataLen, errorCode);
        }
        


        // int MSPAPI MSPSetParam( const char* paramName, const char* paramValue );
        XUNFEI_API int  xfMSPSetParam(const char * paramName, const char * paramValue) {
            
            
            
            
            
            return MSPSetParam(paramName, paramValue);
        }
        


        // const char* MSPAPI MSPUploadData(const char* dataName, void* data, unsigned int dataLen, const char* params, int* errorCode);
        XUNFEI_API const char *  xfMSPUploadData(const char * dataName, void * data, unsigned int dataLen, const char * params, int * errorCode) {
            
            
            
            
            
            return MSPUploadData(dataName, data, dataLen, params, errorCode);
        }
        


        // int MSPAPI QISEAudioWrite(const char* sessionID, const void* waveData, unsigned int waveLen, int audioStatus, int *epStatus, int *Status);
        XUNFEI_API int  xfQISEAudioWrite(const char * sessionID, const void * waveData, unsigned int waveLen, int audioStatus, int * epStatus, int * Status) {
            
            
            
            
            
            return QISEAudioWrite(sessionID, waveData, waveLen, audioStatus, epStatus, Status);
        }
        


        // int MSPAPI QISEGetParam(const char* sessionID, const char* paramName, char* paramValue, unsigned int* valueLen);
        XUNFEI_API int  xfQISEGetParam(const char * sessionID, const char * paramName, char * paramValue, unsigned int * valueLen) {
            
            
            
            
            
            return QISEGetParam(sessionID, paramName, paramValue, valueLen);
        }
        


        // const char * MSPAPI QISEGetResult(const char* sessionID, unsigned int* rsltLen, int* rsltStatus, int *errorCode);
        XUNFEI_API const char *   xfQISEGetResult(const char * sessionID, unsigned int * rsltLen, int * rsltStatus, int * errorCode) {
            
            
            
            
            
            return QISEGetResult(sessionID, rsltLen, rsltStatus, errorCode);
        }
        


        // const char* MSPAPI QISEResultInfo(const char* sessionID, int *errorCode);
        XUNFEI_API const char *  xfQISEResultInfo(const char * sessionID, int * errorCode) {
            
            
            
            
            
            return QISEResultInfo(sessionID, errorCode);
        }
        


        // const char* MSPAPI QISESessionBegin(const char* params, const char* userModelId, int* errorCode);
        XUNFEI_API const char *  xfQISESessionBegin(const char * params, const char * userModelId, int * errorCode) {
            
            
            
            
            
            return QISESessionBegin(params, userModelId, errorCode);
        }
        


        // int MSPAPI QISESessionEnd(const char* sessionID, const char* hints);
        XUNFEI_API int  xfQISESessionEnd(const char * sessionID, const char * hints) {
            
            
            
            
            
            return QISESessionEnd(sessionID, hints);
        }
        


        // int MSPAPI QISETextPut(const char* sessionID, const char* textString, unsigned int textLen, const char* params);
        XUNFEI_API int  xfQISETextPut(const char * sessionID, const char * textString, unsigned int textLen, const char * params) {
            
            
            
            
            
            return QISETextPut(sessionID, textString, textLen, params);
        }
        


        // int MSPAPI QISRAudioWrite(const char* sessionID, const void* waveData, unsigned int waveLen, int audioStatus, int *epStatus, int *recogStatus);
        XUNFEI_API int  xfQISRAudioWrite(const char * sessionID, const void * waveData, unsigned int waveLen, int audioStatus, int * epStatus, int * recogStatus) {
            
            
            
            
            
            return QISRAudioWrite(sessionID, waveData, waveLen, audioStatus, epStatus, recogStatus);
        }
        


        // int MSPAPI QISRBuildGrammar(const char *grammarType, const char *grammarContent, unsigned int grammarLength, const char *params, GrammarCallBack callback, void *userData);
        XUNFEI_API int  xfQISRBuildGrammar(const char * grammarType, const char * grammarContent, unsigned int grammarLength, const char * params, GrammarCallBack callback, void * userData) {
            
            
            
            
            
            return QISRBuildGrammar(grammarType, grammarContent, grammarLength, params, callback, userData);
        }
        


        // const char * MSPAPI QISRGetBinaryResult(const char* sessionID, unsigned int* rsltLen,int* rsltStatus, int waitTime, int *errorCode);
        XUNFEI_API const char *   xfQISRGetBinaryResult(const char * sessionID, unsigned int * rsltLen, int * rsltStatus, int waitTime, int * errorCode) {
            
            
            
            
            
            return QISRGetBinaryResult(sessionID, rsltLen, rsltStatus, waitTime, errorCode);
        }
        


        // int MSPAPI QISRGetParam(const char* sessionID, const char* paramName, char* paramValue, unsigned int* valueLen);
        XUNFEI_API int  xfQISRGetParam(const char * sessionID, const char * paramName, char * paramValue, unsigned int * valueLen) {
            
            
            
            
            
            return QISRGetParam(sessionID, paramName, paramValue, valueLen);
        }
        


        // const char * MSPAPI QISRGetResult(const char* sessionID, int* rsltStatus, int waitTime, int *errorCode);
        XUNFEI_API const char *   xfQISRGetResult(const char * sessionID, int * rsltStatus, int waitTime, int * errorCode) {
            
            
            
            
            
            return QISRGetResult(sessionID, rsltStatus, waitTime, errorCode);
        }
        


        // const char* MSPAPI QISRSessionBegin(const char* grammarList, const char* params, int* errorCode);
        XUNFEI_API const char *  xfQISRSessionBegin(const char * grammarList, const char * params, int * errorCode) {
            
            
            
            
            
            return QISRSessionBegin(grammarList, params, errorCode);
        }
        


        // int MSPAPI QISRSessionEnd(const char* sessionID, const char* hints);
        XUNFEI_API int  xfQISRSessionEnd(const char * sessionID, const char * hints) {
            
            
            
            
            
            return QISRSessionEnd(sessionID, hints);
        }
        


        // int MSPAPI QISRSetParam(const char* sessionID, const char* paramName, const char* paramValue);
        XUNFEI_API int  xfQISRSetParam(const char * sessionID, const char * paramName, const char * paramValue) {
            
            
            
            
            
            return QISRSetParam(sessionID, paramName, paramValue);
        }
        


        // int MSPAPI QISRUpdateLexicon(const char *lexiconName, const char *lexiconContent, unsigned int lexiconLength, const char *params, LexiconCallBack callback, void *userData);
        XUNFEI_API int  xfQISRUpdateLexicon(const char * lexiconName, const char * lexiconContent, unsigned int lexiconLength, const char * params, LexiconCallBack callback, void * userData) {
            
            
            
            
            
            return QISRUpdateLexicon(lexiconName, lexiconContent, lexiconLength, params, callback, userData);
        }
        


        // int MSPAPI QIVWAudioWrite(const char *sessionID, const void *audioData, unsigned int audioLen, int audioStatus);
        XUNFEI_API int  xfQIVWAudioWrite(const char * sessionID, const void * audioData, unsigned int audioLen, int audioStatus) {
            
            
            
            
            
            return QIVWAudioWrite(sessionID, audioData, audioLen, audioStatus);
        }
        


        // int MSPAPI QIVWRegisterNotify(const char *sessionID, ivw_ntf_handler msgProcCb, void *userData);
        XUNFEI_API int  xfQIVWRegisterNotify(const char * sessionID, ivw_ntf_handler msgProcCb, void * userData) {
            
            
            
            
            
            return QIVWRegisterNotify(sessionID, msgProcCb, userData);
        }
        


        // int MSPAPI QIVWResMerge(const char *srcPath, const char *destPath, const char *params);
        XUNFEI_API int  xfQIVWResMerge(const char * srcPath, const char * destPath, const char * params) {
            
            
            
            
            
            return QIVWResMerge(srcPath, destPath, params);
        }
        


        // const char* MSPAPI QIVWSessionBegin(const char *grammarList, const char *params, int *errorCode);
        XUNFEI_API const char *  xfQIVWSessionBegin(const char * grammarList, const char * params, int * errorCode) {
            
            
            
            
            
            return QIVWSessionBegin(grammarList, params, errorCode);
        }
        


        // int MSPAPI QIVWSessionEnd(const char *sessionID, const char *hints);
        XUNFEI_API int  xfQIVWSessionEnd(const char * sessionID, const char * hints) {
            
            
            
            
            
            return QIVWSessionEnd(sessionID, hints);
        }
        


        // const void* MSPAPI QTTSAudioGet(const char* sessionID, unsigned int* audioLen, int* synthStatus, int* errorCode);
        XUNFEI_API const void *  xfQTTSAudioGet(const char * sessionID, unsigned int * audioLen, int * synthStatus, int * errorCode) {
            
            
            
            
            
            return QTTSAudioGet(sessionID, audioLen, synthStatus, errorCode);
        }
        


        // const char* MSPAPI QTTSAudioInfo(const char* sessionID);
        XUNFEI_API const char *  xfQTTSAudioInfo(const char * sessionID) {
            
            
            
            
            
            return QTTSAudioInfo(sessionID);
        }
        


        // int MSPAPI QTTSGetParam(const char* sessionID, const char* paramName, char* paramValue, unsigned int* valueLen);
        XUNFEI_API int  xfQTTSGetParam(const char * sessionID, const char * paramName, char * paramValue, unsigned int * valueLen) {
            
            
            
            
            
            return QTTSGetParam(sessionID, paramName, paramValue, valueLen);
        }
        


        // const char* MSPAPI QTTSSessionBegin(const char* params, int* errorCode);
        XUNFEI_API const char *  xfQTTSSessionBegin(const char * params, int * errorCode) {
            
            
            
            
            
            return QTTSSessionBegin(params, errorCode);
        }
        


        // int MSPAPI QTTSSessionEnd(const char* sessionID, const char* hints);
        XUNFEI_API int  xfQTTSSessionEnd(const char * sessionID, const char * hints) {
            
            
            
            
            
            return QTTSSessionEnd(sessionID, hints);
        }
        


        // int MSPAPI QTTSSetParam(const char *sessionID, const char *paramName, const char *paramValue);
        XUNFEI_API int  xfQTTSSetParam(const char * sessionID, const char * paramName, const char * paramValue) {
            
            
            
            
            
            return QTTSSetParam(sessionID, paramName, paramValue);
        }
        


        // int MSPAPI QTTSTextPut(const char* sessionID, const char* textString, unsigned int textLen, const char* params);
        XUNFEI_API int  xfQTTSTextPut(const char * sessionID, const char * textString, unsigned int textLen, const char * params) {
            
            
            
            
            
            return QTTSTextPut(sessionID, textString, textLen, params);
        }
        

