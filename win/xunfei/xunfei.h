#ifdef XUNFEI_EXPORTS
#define XUNFEI_API __declspec(dllexport)
#else
#define XUNFEI_API __declspec(dllimport)
#endif

extern "C" {

        XUNFEI_API const void *  xfMSPDownloadData(const char * params, unsigned int * dataLen, int * errorCode);
        


        XUNFEI_API int  xfMSPGetParam(const char * paramName, char * paramValue, unsigned int * valueLen);
        


        XUNFEI_API const char *  xfMSPGetVersion(const char * verName, int * errorCode);
        


        XUNFEI_API int  xfMSPLogin(const char * usr, const char * pwd, const char * params);
        


        XUNFEI_API int  xfMSPLogout();
        


        XUNFEI_API int  xfMSPNlpSchCancel(const char * sessionID, const char * hints);
        


        XUNFEI_API const char *  xfMSPNlpSearch(const char * params, const char * text, unsigned int textLen, int * errorCode, NLPSearchCB callback, void * userData);
        


        XUNFEI_API int  xfMSPRegisterNotify(msp_status_ntf_handler statusCb, void * userData);
        


        XUNFEI_API const char *  xfMSPSearch(const char * params, const char * text, unsigned int * dataLen, int * errorCode);
        


        XUNFEI_API int  xfMSPSetParam(const char * paramName, const char * paramValue);
        


        XUNFEI_API const char *  xfMSPUploadData(const char * dataName, void * data, unsigned int dataLen, const char * params, int * errorCode);
        


        XUNFEI_API int  xfQISEAudioWrite(const char * sessionID, const void * waveData, unsigned int waveLen, int audioStatus, int * epStatus, int * Status);
        


        XUNFEI_API int  xfQISEGetParam(const char * sessionID, const char * paramName, char * paramValue, unsigned int * valueLen);
        


        XUNFEI_API   xfQISEGetResult(const char * sessionID, unsigned int * rsltLen, int * rsltStatus, int * errorCode);
        


        XUNFEI_API const char *  xfQISEResultInfo(const char * sessionID, int * errorCode);
        


        XUNFEI_API const char *  xfQISESessionBegin(const char * params, const char * userModelId, int * errorCode);
        


        XUNFEI_API int  xfQISESessionEnd(const char * sessionID, const char * hints);
        


        XUNFEI_API int  xfQISETextPut(const char * sessionID, const char * textString, unsigned int textLen, const char * params);
        


        XUNFEI_API int  xfQISRAudioWrite(const char * sessionID, const void * waveData, unsigned int waveLen, int audioStatus, int * epStatus, int * recogStatus);
        


        XUNFEI_API int  xfQISRBuildGrammar(const char * grammarType, const char * grammarContent, unsigned int grammarLength, const char * params, GrammarCallBack callback, void * userData);
        


        XUNFEI_API   xfQISRGetBinaryResult(const char * sessionID, unsigned int * rsltLen, int * rsltStatus, int waitTime, int * errorCode);
        


        XUNFEI_API int  xfQISRGetParam(const char * sessionID, const char * paramName, char * paramValue, unsigned int * valueLen);
        


        XUNFEI_API   xfQISRGetResult(const char * sessionID, int * rsltStatus, int waitTime, int * errorCode);
        


        XUNFEI_API const char *  xfQISRSessionBegin(const char * grammarList, const char * params, int * errorCode);
        


        XUNFEI_API int  xfQISRSessionEnd(const char * sessionID, const char * hints);
        


        XUNFEI_API int  xfQISRSetParam(const char * sessionID, const char * paramName, const char * paramValue);
        


        XUNFEI_API int  xfQISRUpdateLexicon(const char * lexiconName, const char * lexiconContent, unsigned int lexiconLength, const char * params, LexiconCallBack callback, void * userData);
        


        XUNFEI_API int  xfQIVWAudioWrite(const char * sessionID, const void * audioData, unsigned int audioLen, int audioStatus);
        


        XUNFEI_API int  xfQIVWRegisterNotify(const char * sessionID, ivw_ntf_handler msgProcCb, void * userData);
        


        XUNFEI_API int  xfQIVWResMerge(const char * srcPath, const char * destPath, const char * params);
        


        XUNFEI_API const char *  xfQIVWSessionBegin(const char * grammarList, const char * params, int * errorCode);
        


        XUNFEI_API int  xfQIVWSessionEnd(const char * sessionID, const char * hints);
        


        XUNFEI_API const void *  xfQTTSAudioGet(const char * sessionID, unsigned int * audioLen, int * synthStatus, int * errorCode);
        


        XUNFEI_API const char *  xfQTTSAudioInfo(const char * sessionID);
        


        XUNFEI_API int  xfQTTSGetParam(const char * sessionID, const char * paramName, char * paramValue, unsigned int * valueLen);
        


        XUNFEI_API const char *  xfQTTSSessionBegin(const char * params, int * errorCode);
        


        XUNFEI_API int  xfQTTSSessionEnd(const char * sessionID, const char * hints);
        


        XUNFEI_API int  xfQTTSSetParam(const char * sessionID, const char * paramName, const char * paramValue);
        


        XUNFEI_API int  xfQTTSTextPut(const char * sessionID, const char * textString, unsigned int textLen, const char * params);
        

}
