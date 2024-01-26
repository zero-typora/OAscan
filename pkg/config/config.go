package config

import (
	V9_design_save_svg "OAScan/pkg/POC/fanruan/V9-design-save-svg"
	fanraun_mulubianli "OAScan/pkg/POC/fanruan/fanraun-mulubianli"
	v8_get_geofileload "OAScan/pkg/POC/fanruan/v8-get-geofileload"
	eclogy_mysql_config "OAScan/pkg/POC/fanwei/eclogy-mysql-config"
	ecology_HrmCareerApplyPerView_sqljni "OAScan/pkg/POC/fanwei/ecology-HrmCareerApplyPerView-sqljni"
	ecology_SignatureDownLoad "OAScan/pkg/POC/fanwei/ecology-SignatureDownLoad"
	ecology_Userselect "OAScan/pkg/POC/fanwei/ecology-Userselect"
	ecology_VerifyQuickLogin "OAScan/pkg/POC/fanwei/ecology-VerifyQuickLogin"
	ecology_Weaver_fileupload "OAScan/pkg/POC/fanwei/ecology-Weaver-fileupload"
	ecology_Weaver_lnFileDownload "OAScan/pkg/POC/fanwei/ecology-Weaver-lnFileDownload"
	ecology_WorkflowCenterTreeData_sqljni "OAScan/pkg/POC/fanwei/ecology-WorkflowCenterTreeData-sqljni"
	ecology_bashservlet_rce "OAScan/pkg/POC/fanwei/ecology-bashservlet-rce"
	ecology_browser_sqljni "OAScan/pkg/POC/fanwei/ecology-browser-sqljni"
	ecology_filedownload_directory_jqueryFileTree "OAScan/pkg/POC/fanwei/ecology-filedownload-directory-jqueryFileTree"
	ecology_filedownload_directory_traversal "OAScan/pkg/POC/fanwei/ecology-filedownload-directory-traversal"
	ecology_fileupload "OAScan/pkg/POC/fanwei/ecology-fileupload"
	ecology_getdata_sqljni "OAScan/pkg/POC/fanwei/ecology-getdata-sqljni"
	ecology_getsqlData_sqljni "OAScan/pkg/POC/fanwei/ecology-getsqlData-sqljni"
	ecology_group_xml_sqljni "OAScan/pkg/POC/fanwei/ecology-group-xml-sqljni"
	ecology_loginsso_sqljni "OAScan/pkg/POC/fanwei/ecology-loginsso-sqljni"
	ecology_officeServe_uploadfile "OAScan/pkg/POC/fanwei/ecology-officeServe-uploadfile"
	ecology_officeserver_fileload "OAScan/pkg/POC/fanwei/ecology-officeserver-fileload"
	ecology_saveYZJFile "OAScan/pkg/POC/fanwei/ecology-saveYZJFile"
	ecology_uploadFile_office "OAScan/pkg/POC/fanwei/ecology-uploadFile-office"
	ecology_usersdata "OAScan/pkg/POC/fanwei/ecology-usersdata"
	saveYZJFile_filereadlinux "OAScan/pkg/POC/fanwei/saveYZJFile-fileread-linux"
	wanhu_DownloadServlet "OAScan/pkg/POC/wanhu/wanhu-DownloadServlet"
	wanhu_aiframe_sqljni "OAScan/pkg/POC/wanhu/wanhu-aiframe-sqljni"
	wanhu_defaultroot_sqljni "OAScan/pkg/POC/wanhu/wanhu-defaultroot-sqljni"
	wanhu_download_ftp "OAScan/pkg/POC/wanhu/wanhu-download-ftp"
	wanhu_download_http "OAScan/pkg/POC/wanhu/wanhu-download-http"
	wanhu_download_old "OAScan/pkg/POC/wanhu/wanhu-download-old"
	wanhu_ezoffice_documentedit_sql "OAScan/pkg/POC/wanhu/wanhu-ezoffice-documentedit-sql"
	wanhu_fileUpload_controller "OAScan/pkg/POC/wanhu/wanhu-fileUpload-controller"
	wanhu_officeService "OAScan/pkg/POC/wanhu/wanhu-officeService"
	wanhu_smartUpload "OAScan/pkg/POC/wanhu/wanhu-smartUpload"
	wanhu_text2Html "OAScan/pkg/POC/wanhu/wanhu-text2Html"
	"OAScan/pkg/POC/yongyou/ConfigurationNc"
	erpncmlbl "OAScan/pkg/POC/yongyou/ERP-NC-MLBL"
	FileReceiveServlet_Deser "OAScan/pkg/POC/yongyou/FileReceiveServlet-Deser"
	GRP_U8_Proxy_sqljin_xxe "OAScan/pkg/POC/yongyou/GRP-U8-Proxy-sqljin-xxe"
	GRP_U8_U8AppProxy "OAScan/pkg/POC/yongyou/GRP-U8-U8AppProxy"
	GRP_U8_UploadFileData "OAScan/pkg/POC/yongyou/GRP-U8-UploadFileData"
	GRP_U8_historyDataCheck_sqljin "OAScan/pkg/POC/yongyou/GRP-U8-historyDataCheck-sqljin"
	KSOA_ImageUpload "OAScan/pkg/POC/yongyou/KSOA-ImageUpload"
	KSOA_sqljni "OAScan/pkg/POC/yongyou/KSOA-sqljni"
	KSOA_sqljni_PayBill "OAScan/pkg/POC/yongyou/KSOA-sqljni-PayBill"
	KSOA_Imagefield_Sqli "OAScan/pkg/POC/yongyou/KSOA_imagefield_sqli"
	KSOA_sqljni_depti "OAScan/pkg/POC/yongyou/KSOA_sqljni_deptid"
	MessageServlet_Deser "OAScan/pkg/POC/yongyou/MessageServlet-Deser"
	NC_BshServlet "OAScan/pkg/POC/yongyou/NC-BshServlet"
	u8_MxServlet "OAScan/pkg/POC/yongyou/NC-Cloud-MxServlet"
	NC_JiuQiClientReqDispatch "OAScan/pkg/POC/yongyou/NC-JiuQiClientReqDispatch"
	NC_XbrlPersistenceServlet "OAScan/pkg/POC/yongyou/NC-XbrlPersistenceServlet"
	NC6_5_UploadFile "OAScan/pkg/POC/yongyou/NC6.5-UploadFile"
	NCCloud_FS_sqljni "OAScan/pkg/POC/yongyou/NCCloud-FS-sqljni"
	T_CRM_sqljni "OAScan/pkg/POC/yongyou/T-CRM-sqljni"
	T_DownloadProxy_catfile "OAScan/pkg/POC/yongyou/T-DownloadProxy-catfile"
	T_RecoverPassword "OAScan/pkg/POC/yongyou/T-RecoverPassword"
	T_Uploadfile "OAScan/pkg/POC/yongyou/T-Uploadfile"
	U8_ActionHandlerServlet "OAScan/pkg/POC/yongyou/U8-ActionHandlerServlet"
	U8_CacheInvokeServlet "OAScan/pkg/POC/yongyou/U8-CacheInvokeServlet"
	U8_ClientRequestDispatch "OAScan/pkg/POC/yongyou/U8-ClientRequestDispatch"
	u8_FileTransportServlet "OAScan/pkg/POC/yongyou/U8-FileTransportServlet"
	U8_OA_getSessionList "OAScan/pkg/POC/yongyou/U8-OA-getSessionList"
	U8_OA_test_sqjni "OAScan/pkg/POC/yongyou/U8-OA-test-sqjni"
	U8_RegisterServlet "OAScan/pkg/POC/yongyou/U8-RegisterServlet"
	U8_TaskTreeQuery "OAScan/pkg/POC/yongyou/U8-TaskTreeQuery"
	Uapjs_JNDI "OAScan/pkg/POC/yongyou/Uapjs-JNDI"
	UploadServlet_Deser "OAScan/pkg/POC/yongyou/UploadServlet-Deser"
	Youyon_uploadIcon_do_upload "OAScan/pkg/POC/yongyou/Youyon-uploadIcon-do-upload"
	accept_upload "OAScan/pkg/POC/yongyou/accept-upload"
	files_Deser "OAScan/pkg/POC/yongyou/files-Deser"
	fs_dlbypass "OAScan/pkg/POC/yongyou/fs-dlbypass"
	monitorservlet_Desera "OAScan/pkg/POC/yongyou/monitorservlet-Desera"
	nc_cloud_portal_file "OAScan/pkg/POC/yongyou/nc-cloud-portal-file"
	u8_LoginServlet "OAScan/pkg/POC/yongyou/u8-LoginServlet"
	u8_MonitorServlet "OAScan/pkg/POC/yongyou/u8-MonitorServlet"
	u8_ServletCommander "OAScan/pkg/POC/yongyou/u8-ServletCommander"
	u8_TableInputOperServlet "OAScan/pkg/POC/yongyou/u8-TableInputOperServlet"
	uapws_acessBypass "OAScan/pkg/POC/yongyou/uapws-acessBypass"
	uapws_soapFormat_xxe "OAScan/pkg/POC/yongyou/uapws-soapFormat-xxe"
	uapws_wsdl_XXE "OAScan/pkg/POC/yongyou/uapws-wsdl-XXE"
	A6_DownExcelBeanServlet "OAScan/pkg/POC/zhiyuan/A6-DownExcelBeanServlet"
	A6_config_leakage "OAScan/pkg/POC/zhiyuan/A6-config-leakage"
	A6_createMysql_information_leakage "OAScan/pkg/POC/zhiyuan/A6-createMysql-information-leakage"
	A6_getSessionList_session "OAScan/pkg/POC/zhiyuan/A6-getSessionList-session"
	A6_initDataAssess "OAScan/pkg/POC/zhiyuan/A6-initDataAssess"
	A6_setextno_sqljin "OAScan/pkg/POC/zhiyuan/A6-setextno-sqljin"
	A6_test_jsp_sqljni "OAScan/pkg/POC/zhiyuan/A6-test.jsp-sqljni"
	A8_Session_fileupload "OAScan/pkg/POC/zhiyuan/A8-Session-fileupload"
	A8_ajaxfile_upload "OAScan/pkg/POC/zhiyuan/A8-ajaxfile-upload"
	A8_file_load "OAScan/pkg/POC/zhiyuan/A8-file-load"
	A8_getAjaxDataServlet_xxe "OAScan/pkg/POC/zhiyuan/A8-getAjaxDataServlet-xxe"
	A8_htmlofficeservlet_uploadfile "OAScan/pkg/POC/zhiyuan/A8-htmlofficeservlet-uploadfile"
	A8_resetpassword "OAScan/pkg/POC/zhiyuan/A8-resetpassword"
	A8_status_information_leakage "OAScan/pkg/POC/zhiyuan/A8-status-information-leakage"
	A8_webmail_filedown "OAScan/pkg/POC/zhiyuan/A8-webmail-filedown"
	A8_wpsAssistServlet "OAScan/pkg/POC/zhiyuan/A8-wpsAssistServlet"
	A8_wpsAssistServlet_fileread "OAScan/pkg/POC/zhiyuan/A8-wpsAssistServlet-fileread"
	zhiyuan_M1Serve_userTokenService "OAScan/pkg/POC/zhiyuan/zhiyuan-M1Serve-userTokenService"
	zhiyuan_OA_ReportServer "OAScan/pkg/POC/zhiyuan/zhiyuan-OA-ReportServer"
	zhiyuan_m3server_mobile_portal_rce "OAScan/pkg/POC/zhiyuan/zhiyuan-m3server-mobile-portal-rce"
	"fmt"
	"github.com/gookit/color"
	"sync"
)

type WorkExp struct {
	Url    string
	OAType int
}

func (c *WorkExp) ScanRun() {
	switch c.OAType {
	case 1:
		color.Blue.Println("[+] URl: " + c.Url)
		var wg sync.WaitGroup
		wg.Add(46)
		go func() {
			nc_cloud_portal_file.Run(c.Url)
			wg.Done()
		}()
		go func() {
			erpncmlbl.Run(c.Url)
			wg.Done()
		}()
		go func() {
			NC_BshServlet.Run(c.Url)
			wg.Done()
		}()
		go func() {
			NCCloud_FS_sqljni.Run(c.Url)
			wg.Done()
		}()
		go func() {
			NC6_5_UploadFile.Run(c.Url)
			wg.Done()
		}()
		go func() {
			NC_XbrlPersistenceServlet.Run(c.Url)
			wg.Done()
		}()
		go func() {
			U8_OA_getSessionList.Run(c.Url)
			wg.Done()
		}()
		go func() {
			U8_OA_test_sqjni.Run(c.Url)
			wg.Done()
		}()
		go func() {
			GRP_U8_UploadFileData.Run(c.Url)
			wg.Done()
		}()
		go func() {
			GRP_U8_Proxy_sqljin_xxe.Run(c.Url)
			wg.Done()
		}()
		go func() {
			Uapjs_JNDI.Run(c.Url)
			wg.Done()
		}()
		go func() {
			T_CRM_sqljni.Run(c.Url)
			wg.Done()
		}()
		go func() {
			T_Uploadfile.Run(c.Url)
			wg.Done()
		}()
		go func() {
			T_RecoverPassword.Run(c.Url)
			wg.Done()
		}()
		go func() {
			GRP_U8_U8AppProxy.Run(c.Url)
			wg.Done()
		}()
		go func() {
			uapws_acessBypass.Run(c.Url)
			wg.Done()
		}()
		go func() {
			fs_dlbypass.Run(c.Url)
			wg.Done()
		}()
		go func() {
			files_Deser.Run(c.Url)
			wg.Done()
		}()
		go func() {
			T_DownloadProxy_catfile.Run(c.Url)
			wg.Done()
		}()
		go func() {
			KSOA_ImageUpload.Run(c.Url)
			wg.Done()
		}()
		go func() {
			accept_upload.Run(c.Url)
			wg.Done()
		}()
		go func() {
			MessageServlet_Deser.Run(c.Url)
			wg.Done()
		}()
		go func() {
			UploadServlet_Deser.Run(c.Url)
			wg.Done()
		}()
		go func() {
			monitorservlet_Desera.Run(c.Url)
			wg.Done()
		}()
		go func() {
			FileReceiveServlet_Deser.Run(c.Url)
			wg.Done()
		}()
		go func() {
			u8_TableInputOperServlet.Run(c.Url)
			wg.Done()
		}()
		go func() {
			u8_LoginServlet.Run(c.Url)
			wg.Done()
		}()
		go func() {
			u8_FileTransportServlet.Run(c.Url)
			wg.Done()
		}()
		go func() {
			U8_CacheInvokeServlet.Run(c.Url)
			wg.Done()
		}()
		go func() {
			U8_ActionHandlerServlet.Run(c.Url)
			wg.Done()
		}()
		go func() {
			u8_ServletCommander.Run(c.Url)
			wg.Done()
		}()
		go func() {
			u8_MxServlet.Run(c.Url)
			wg.Done()
		}()
		go func() {
			u8_MonitorServlet.Run(c.Url)
			wg.Done()
		}()
		go func() {
			U8_ClientRequestDispatch.Run(c.Url)
			wg.Done()
		}()
		go func() {
			KSOA_sqljni_PayBill.Run(c.Url)
			wg.Done()
		}()
		go func() {
			U8_RegisterServlet.Run(c.Url)
			wg.Done()
		}()

		go func() {
			uapws_wsdl_XXE.Run(c.Url)
			wg.Done()
		}()
		go func() {
			ConfigurationNc.Run(c.Url)
			wg.Done()
		}()
		go func() {
			KSOA_sqljni.Run(c.Url)
			wg.Done()
		}()
		go func() {
			NC_JiuQiClientReqDispatch.Run(c.Url)
			wg.Done()
		}()
		go func() {
			KSOA_Imagefield_Sqli.Run(c.Url)
			wg.Done()
		}()
		go func() {
			KSOA_sqljni_depti.Run(c.Url)
			wg.Done()
		}()
		go func() {
			GRP_U8_historyDataCheck_sqljin.Run(c.Url)
			wg.Done()
		}()
		go func() {
			uapws_soapFormat_xxe.Run(c.Url)
			wg.Done()
		}()
		go func() {
			U8_TaskTreeQuery.Run(c.Url)
			wg.Done()
		}()
		go func() {
			Youyon_uploadIcon_do_upload.Run(c.Url)
			wg.Done()
		}()
		wg.Wait()
	case 2:
		var wg sync.WaitGroup
		wg.Add(19) // 仅此 case 中启动了一个 Goroutine
		go func() {
			A8_htmlofficeservlet_uploadfile.Run(c.Url)
			wg.Done()
		}()
		go func() {
			A8_webmail_filedown.Run(c.Url)
			wg.Done()
		}()
		go func() {
			A8_status_information_leakage.Run(c.Url)
			wg.Done()
		}()
		go func() {
			A6_createMysql_information_leakage.Run(c.Url)
			wg.Done()
		}()
		go func() {
			A6_DownExcelBeanServlet.Run(c.Url)
			wg.Done()
		}()
		go func() {
			A6_initDataAssess.Run(c.Url)
			wg.Done()
		}()
		go func() {
			A6_setextno_sqljin.Run(c.Url)
			wg.Done()
		}()
		go func() {
			A8_Session_fileupload.Run(c.Url)
			wg.Done()
		}()
		go func() {
			A8_ajaxfile_upload.Run(c.Url)
			wg.Done()
		}()
		go func() {
			A6_test_jsp_sqljni.Run(c.Url)
			wg.Done()
		}()
		go func() {
			zhiyuan_OA_ReportServer.Run(c.Url)
			wg.Done()
		}()
		go func() {
			zhiyuan_M1Serve_userTokenService.Run(c.Url)
			wg.Done()
		}()
		go func() {
			A8_wpsAssistServlet.Run(c.Url)
			wg.Done()
		}()
		go func() {
			A6_getSessionList_session.Run(c.Url)
			wg.Done()
		}()
		go func() {
			A6_config_leakage.Run(c.Url)
			wg.Done()
		}()
		go func() {
			A8_wpsAssistServlet_fileread.Run(c.Url)
			wg.Done()
		}()
		go func() {
			A8_getAjaxDataServlet_xxe.Run(c.Url)
			wg.Done()
		}()
		go func() {
			zhiyuan_m3server_mobile_portal_rce.Run(c.Url)
			wg.Done()
		}()
		go func() {
			A8_file_load.Run(c.Url)
			wg.Done()
		}()
		go func() {
			A8_resetpassword.Run(c.Url)
			wg.Done()
		}()
		wg.Wait()
	case 3:
		var wg sync.WaitGroup
		wg.Add(3)
		go func() {
			fanraun_mulubianli.Run(c.Url)
			wg.Done()
		}()
		go func() {
			V9_design_save_svg.Run(c.Url)
			wg.Done()
		}()
		go func() {
			v8_get_geofileload.Run(c.Url)
			wg.Done()
		}()
		wg.Wait()
	case 4:
		var wg sync.WaitGroup
		wg.Add(23)
		go func() {
			ecology_filedownload_directory_traversal.Run(c.Url)
			wg.Done()
		}()
		go func() {
			ecology_filedownload_directory_jqueryFileTree.Run(c.Url)
			wg.Done()
		}()
		go func() {
			ecology_VerifyQuickLogin.Run(c.Url)
			wg.Done()
		}()
		go func() {
			ecology_WorkflowCenterTreeData_sqljni.Run(c.Url)
			wg.Done()
		}()
		go func() {
			ecology_fileupload.Run(c.Url)
			wg.Done()
		}()
		go func() {
			ecology_getdata_sqljni.Run(c.Url)
			wg.Done()
		}()
		go func() {
			ecology_HrmCareerApplyPerView_sqljni.Run(c.Url)
			wg.Done()
		}()
		go func() {
			ecology_loginsso_sqljni.Run(c.Url)
			wg.Done()
		}()
		go func() {
			ecology_browser_sqljni.Run(c.Url)
			wg.Done()
		}()
		go func() {
			ecology_bashservlet_rce.Run(c.Url)
			wg.Done()
		}()
		go func() {
			ecology_Weaver_lnFileDownload.Run(c.Url)
			wg.Done()
		}()
		go func() {
			ecology_Weaver_fileupload.Run(c.Url)
			wg.Done()
		}()
		go func() {
			ecology_SignatureDownLoad.Run(c.Url)
			wg.Done()
		}()
		go func() {
			ecology_uploadFile_office.Run(c.Url)
			wg.Done()
		}()
		go func() {
			ecology_group_xml_sqljni.Run(c.Url)
			wg.Done()
		}()
		go func() {
			ecology_officeserver_fileload.Run(c.Url)
			wg.Done()
		}()
		go func() {
			ecology_officeServe_uploadfile.Run(c.Url)
			wg.Done()
		}()
		go func() {
			ecology_usersdata.Run(c.Url)
			wg.Done()
		}()
		go func() {
			ecology_getsqlData_sqljni.Run(c.Url)
			wg.Done()
		}()
		go func() {
			ecology_Userselect.Run(c.Url)
			wg.Done()
		}()
		go func() {
			eclogy_mysql_config.Run(c.Url)
			wg.Done()
		}()
		go func() {
			ecology_saveYZJFile.Run(c.Url)
			wg.Done()
		}()
		go func() {
			saveYZJFile_filereadlinux.Run(c.Url)
			wg.Done()
		}()
		wg.Wait()
	case 5:
		var wg sync.WaitGroup
		wg.Add(11)
		go func() {
			wanhu_ezoffice_documentedit_sql.Run(c.Url)
			wg.Done()
		}()
		go func() {
			wanhu_DownloadServlet.Run(c.Url)
			wg.Done()
		}()
		go func() {
			wanhu_download_ftp.Run(c.Url)
			wg.Done()
		}()
		go func() {
			wanhu_download_old.Run(c.Url)
			wg.Done()
		}()
		go func() {
			wanhu_download_http.Run(c.Url)
			wg.Done()
		}()
		go func() {
			wanhu_officeService.Run(c.Url)
			wg.Done()
		}()
		go func() {
			wanhu_smartUpload.Run(c.Url)
			wg.Done()
		}()
		go func() {
			wanhu_fileUpload_controller.Run(c.Url)
			wg.Done()
		}()
		go func() {
			wanhu_text2Html.Run(c.Url)
			wg.Done()
		}()
		go func() {
			wanhu_aiframe_sqljni.Run(c.Url)
			wg.Done()
		}()
		go func() {
			wanhu_defaultroot_sqljni.Run(c.Url)
			wg.Done()
		}()
		wg.Wait()
	default:
		fmt.Println("未知的 OA 系统类型")
	}
}
