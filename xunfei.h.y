#ifdef XUNFEI_EXPORTS
#define XUNFEI_API __declspec(dllexport)
#else
#define XUNFEI_API __declspec(dllimport)
#endif

extern "C" {
