package main

const (
	Root        = "IoT."
	common      = Root + "Common."
	system      = Root + "System."
	device      = Root + "Device."
	thing       = Root + "Thing."
	asset       = Root + "Asset."
	video       = Root + "Video."
	service     = Root + "Service."
	deviceGroup = Root + "DeviceGroup."
	property    = Root + "Property."
	alarm       = Root + "Alarm."
	task        = Root + "Task."
	token       = Root + "Token."
	dataExport  = Root + "DataExport."
)

const (
	NotExisted  = "NotExisted"
	NameExisted = "NameExisted"
	IdExisted   = "IdExisted"
	ConfExisted = "ConfExisted"
)

const (
	Success       = system + "Success"
	InternalError = system + "InternalError"
	SystemError   = system + "SystemError"
	RpcTimeOut    = system + "TimeOut"

	ParamsTooLong       = common + "ParamsTooLong"
	ResourceNotExisted  = common + "ResourceNotExisted"
	InternalParamsError = common + "InternalParamsError"
	MissingParams       = common + "MissingParams"
)
const (
	//device
	DeviceNotExisted                 = device + NotExisted
	DeviceNameExisted                = device + NameExisted
	DeviceNamePrefixExisted          = device + "NamePrefixExisted"
	DeviceInvalid                    = device + "Invalid"
	ParentDeviceNotExisted           = device + "ParentDevice" + NotExisted
	InvalidParentDevice              = device + "InvalidParentDevice"
	ExistSubDevice                   = device + "ExistSubDevice"
	DeviceStatusSwitchForbidden      = device + "DeviceStatusSwitchForbidden"
	InvalidDeviceStatus              = device + "InvalidDeviceStatus"
	ParentDeviceExistInSubDevices    = device + "ParentDeviceExistInSubDevices"
	DeviceTreeHeightExceed           = device + "DeviceTreeHeightExceed"
	DeviceAlreadyExistParent         = device + "DeviceAlreadyExistParent"
	UpdateDevicePropertyStatusFailed = device + "UpdateDevicePropertyStatusFailed"
	InvalidDeviceAccessProtocol      = device + "InvalidDeviceAccessProtocol"
	InvalidDeviceColumn              = device + "InvalidDeviceColumn"
	InvalidFormattedDeviceName       = device + "InvalidFormattedName"

	DeviceNameTooLong                  = device + "DeviceNameTooLong"
	TaskNamePrefixExisted              = device + "TaskNamePrefixExisted"
	CredentialHasExpired               = device + "CredentialHasExpired"
	RegisterFailedByTaskStatus         = device + "RegisterFailedByTaskStatus"
	TaskHasExpired                     = device + "TaskHasExpired"
	RegisterFailedByExceededTaskMaxNum = device + "RegisterFailedByExceededTaskMaxNum"
	ResourceIdentifierExisted          = device + "ResourceIdentifierExisted"
	DevicePropertyExisted              = device + "Property" + IdExisted

	//thing
	ThingNotExisted           = thing + NotExisted
	ThingNameExisted          = thing + NameExisted
	ThingIdExisted            = thing + IdExisted
	ThingAssociatedWithDevice = thing + "AssociatedWithDevice"
	ThingPropertyExisted      = thing + "Property" + IdExisted      //模型下属性标识符已存在
	ThingEventExisted         = thing + "Event" + IdExisted         //模型下事件标识符已存在
	ThingEventOutputExisted   = thing + "EventOutput" + IdExisted   //模型下事件输出参数标识已存在
	ThingServiceExisted       = thing + "Service" + IdExisted       //模型下服务标识符已存在
	ThingServiceOutputExisted = thing + "ServiceOutput" + IdExisted //模型下服务输出参数标识已存在
	ThingServiceInputExisted  = thing + "ServiceInput" + IdExisted  //模型下服务输入参数标识已存在
	ThingPropertyNotExisted   = thing + "Property" + NotExisted
	ThingExistDevice          = thing + "ThingExistDevice"
	ThingExistDynamicDevice   = thing + "ThingExistDyncmicDevice"

	// task
	TaskExpireTimeInvalid = task + "ExpireTimeInvalid"

	//video
	VideoCreateFail = video + "DeviceCreateFail"
	VideoDeleteFail = video + "DeviceDeleteFail"

	//alarm
	AlarmPolicyCreateFail = alarm + "PolicyCreateFail"

	//asset
	AssetNotExisted           = asset + NotExisted
	AssetNameExisted          = asset + NameExisted
	AssetMissingParams        = asset + "MissingParams"
	AssetAttachedToDevice     = asset + "AssetAttachedToDevice"
	AssetAttachedToDeviceTask = asset + "AssetAttachedToDeviceTask"

	//group
	CanNotDeleteRootDeviceGroup    = deviceGroup + "CanNotDeleteRootDeviceGroup"
	DeviceGroupNotExisted          = deviceGroup + NotExisted
	DeviceGroupNameExisted         = deviceGroup + NameExisted
	DeviceGroupIllegalResourceType = deviceGroup + "IllegalResourceType"

	//token
	TokenNotExisted = token + NotExisted

	//ServiceInfoExisted = service + "ServiceInfo" + IdExisted

	ParamsInvalid             = common + "ParamsInvalid"
	AccountError              = common + "AccountError"
	PitrixError               = common + "PitrixError"
	IAMError                  = common + "IAMError"
	STSError                  = common + "STSError"
	IllegalAccessKey          = common + "IllegalAccessKey"
	AccessKeyNotActive        = common + "AccessKeyNotActive"
	UserNotFound              = common + "UserNotFound"
	UserNotFinishRegistration = common + "UserNotFinishRegistration"
	UserAccessDenied          = common + "UserAccessDenied"
	SuperUserOnly             = common + "SuperUserOnly"
	ResourceNotFound          = common + "ResourceNotFound"

	InvalidPropertyType = property + "InvalidPropertyType"
	InvalidAccessType   = property + "InvalidAccessType"

	// data-export
	DataExportRecordNotExisted            = dataExport + "Record" + NotExisted             // 数据导出记录不存在
	DataExportObjectStorageConfNotExisted = dataExport + "objectStorageConf" + NotExisted  // 对象存储配置不存在
	DataExportObjectStorageConfExisted    = dataExport + "objectStorageConf" + ConfExisted // 对象存储配置已经存在
	DataExportAuthNotExisted              = dataExport + "auth" + NotExisted               // 尚未授权
	DataExportAutoExportInfoNotExisted    = dataExport + "autoExportInfo" + NotExisted     // 数据导出(自动导出)信息不存在
	DataExportInvalidAccessKeyID          = dataExport + "InvalidAccessKeyIDToQingStor"    // 没有 qingStor 资源

	// // data-export
	// DataExportRecordNotExisted  = dataExportRecord + NotExisted
	// ObjectStorageConfNotExisted = objectStorageConf + NotExisted
	// ObjectStorageConfExisted    = objectStorageConf + ConfExisted
	// AuthNotExisted              = auth + NotExisted
	// AutoExportInfoNotExisted    = autoExportInfo + NotExisted
)

var (
	Errors = map[string]string{

		DataExportRecordNotExisted:            "数据导出记录不存在",
		DataExportObjectStorageConfNotExisted: "请授权qingStor对象存储授权",
		DataExportObjectStorageConfExisted:    "已授权qingStor对象存储",
		DataExportAuthNotExisted:              "请授权qingStor对象存储授权",
		DataExportAutoExportInfoNotExisted:    "数据导出(自动导出)信息不存在",
		DataExportInvalidAccessKeyID:          "没有 qingStor 对象存储权限, 请确认已开通并有足够配额!",
	}
)
