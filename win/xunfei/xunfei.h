// ���� ifdef ���Ǵ���ʹ�� DLL �������򵥵�
// ��ı�׼�������� DLL �е������ļ��������������϶���� XUNFEI_EXPORTS
// ���ű���ġ���ʹ�ô� DLL ��
// �κ�������Ŀ�ϲ�Ӧ����˷��š�������Դ�ļ��а������ļ����κ�������Ŀ���Ὣ
// XUNFEI_API ������Ϊ�Ǵ� DLL ����ģ����� DLL ���ô˺궨���
// ������Ϊ�Ǳ������ġ�
#ifdef XUNFEI_EXPORTS
#define XUNFEI_API __declspec(dllexport)
#else
#define XUNFEI_API __declspec(dllimport)
#endif

extern "C" {
	XUNFEI_API int xfMSPLogin(const char* usr, const char* pwd, const char* params);
}