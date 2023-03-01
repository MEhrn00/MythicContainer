package agentstructs

// PT_TASK_* structs

type PTTaskMessageAllData struct {
	Task            PTTaskMessageTaskData                `json:"task"`
	Callback        PTTaskMessageCallbackData            `json:"callback"`
	BuildParameters []PayloadConfigurationBuildParameter `json:"build_parameters"`
	Commands        []string                             `json:"commands"`
	Payload         PTTaskMessagePayloadData             `json:"payload"`
	C2Profiles      []PayloadConfigurationC2Profile      `json:"c2info"`
	PayloadType     string                               `json:"payload_type"`
	Args            PTTaskMessageArgsData
}

type PTTaskMessageTaskData struct {
	ID                                 int    `json:"id"`
	AgentTaskID                        string `json:"agent_task_id"`
	CommandName                        string `json:"command_name"`
	Params                             string `json:"params"`
	Timestamp                          string `json:"timestamp"`
	CallbackID                         int    `json:"callback_id"`
	Status                             string `json:"status"`
	OriginalParams                     string `json:"original_params"`
	DisplayParams                      string `json:"display_params"`
	Comment                            string `json:"comment"`
	Stdout                             string `json:"stdout"`
	Stderr                             string `json:"stderr"`
	Completed                          bool   `json:"completed"`
	OperatorUsername                   string `json:"operator_username"`
	OpsecPreBlocked                    bool   `json:"opsec_pre_blocked"`
	OpsecPreMessage                    string `json:"opsec_pre_message"`
	OpsecPreBypassed                   bool   `json:"opsec_pre_bypassed"`
	OpsecPreBypassRole                 string `json:"opsec_pre_bypass_role"`
	OpsecPostBlocked                   bool   `json:"opsec_post_blocked"`
	OpsecPostMessage                   string `json:"opsec_post_message"`
	OpsecPostBypassed                  bool   `json:"opsec_post_bypassed"`
	OpsecPostBypassRole                string `json:"opsec_post_bypass_role"`
	ParentTaskID                       int    `json:"parent_task_id"`
	SubtaskCallbackFunction            string `json:"subtask_callback_function"`
	SubtaskCallbackFunctionCompleted   bool   `json:"subtask_callback_function_completed"`
	GroupCallbackFunction              string `json:"group_callback_function"`
	GroupCallbackFunctionCompleted     bool   `json:"group_callback_function_completed"`
	CompletedCallbackFunction          string `json:"completed_callback_function"`
	CompletedCallbackFunctionCompleted bool   `json:"completed_callback_function_completed"`
	SubtaskGroupName                   string `json:"subtask_group_name"`
	TaskingLocation                    string `json:"tasking_location"`
	ParameterGroupName                 string `json:"parameter_group_name"`
	TokenID                            int    `json:"token_id"`
}

type PTTaskMessageCallbackData struct {
	ID                  int    `json:"id"`
	AgentCallbackID     string `json:"agent_callback_id"`
	InitCallback        string `json:"init_callback"`
	LastCheckin         string `json:"last_checkin"`
	User                string `json:"user"`
	Host                string `json:"host"`
	PID                 int    `json:"pid"`
	Ip                  string `json:"ip"`
	ExternalIp          string `json:"external_ip"`
	ProcessName         string `json:"process_name"`
	Description         string `json:"description"`
	OperatorID          int    `json:"operator_id"`
	Active              bool   `json:"active"`
	RegisteredPayloadID int    `json:"registered_payload_id"`
	IntegrityLevel      int    `json:"integrity_level"`
	Locked              bool   `json:"locked"`
	OperationID         int    `json:"operation_id"`
	CryptoType          string `json:"crypto_type"`
	DecKey              []byte `json:"dec_key"`
	EncKey              []byte `json:"enc_key"`
	Os                  string `json:"os"`
	Architecture        string `json:"architecture"`
	Domain              string `json:"domain"`
	ExtraInfo           string `json:"extra_info"`
	SleepInfo           string `json:"sleep_info"`
}

type PTTaskMessagePayloadData struct {
	Os          string `json:"os"`
	UuID        string `json:"uuid"`
	PayloadType string `json:"payload_type"`
}

type PT_TASK_FUNCTION_STATUS = string

type PtTaskFunctionParseArgString func(args *PTTaskMessageArgsData, input string) error
type PtTaskFunctionParseArgDictionary func(args *PTTaskMessageArgsData, input map[string]interface{}) error
type PTTaskMessageArgsData struct {
	args            []CommandParameter
	commandLine     string
	rawCommandLine  string
	taskingLocation string
	manualArgs      *string
}

const (
	PT_TASK_FUNCTION_STATUS_OPSEC_PRE                        PT_TASK_FUNCTION_STATUS = "OPSEC Pre Check Running..."
	PT_TASK_FUNCTION_STATUS_OPSEC_PRE_ERROR                                          = "Error: OPSEC Pre Check"
	PT_TASK_FUNCTION_STATUS_OPSEC_PRE_BLOCKED                                        = "OPSEC Pre Block Hit"
	PT_TASK_FUNCTION_STATUS_PREPROCESSING                                            = "preprocessing"
	PT_TASK_FUNCTION_STATUS_PREPROCESSING_ERROR                                      = "Error: preprocessing"
	PT_TASK_FUNCTION_STATUS_OPSEC_POST                                               = "OPSEC Post Check Running..."
	PT_TASK_FUNCTION_STATUS_OPSEC_POST_ERROR                                         = "Error: OPSEC Post Check"
	PT_TASK_FUNCTION_STATUS_OPSEC_POST_BLOCKED                                       = "OPSEC Post Block Hit"
	PT_TASK_FUNCTION_STATUS_SUBMITTED                                                = "submitted"
	PT_TASK_FUNCTION_STATUS_COMPLETION_FUNCTION                                      = "Completion Function Running..."
	PT_TASK_FUNCTION_STATUS_COMPLETION_FUNCTION_ERROR                                = "Error: Completion Function"
	PT_TASK_FUNCTION_STATUS_SUBTASK_COMPLETED_FUNCTION                               = "SubTask Completion Function Running..."
	PT_TASK_FUNCTION_STATUS_SUBTASK_COMPLETED_FUNCTION_ERROR                         = "Error: Subtask Completion Function"
	PT_TASK_FUNCTION_STATUS_GROUP_COMPLETED_FUNCTION                                 = "Group Completion Function Running..."
	PT_TASK_FUNCTION_STATUS_GROUP_COMPLETED_FUNCTION_ERROR                           = "Error: Group Completion Function"
	PT_TASK_FUNCTION_STATUS_PROCESS_RESPONSE_FUNCTION                                = "Task Processing Response Manually..."
	PT_TASK_FUNCTION_STATUS_COMPLETED                                                = "completed"
)

// Tasking step 1:
// Task message/process before running create_tasking function
//
//	opportunity to run any necessary opsec checks/blocks before the logic in create_tasking runs
//		which can spawn subtasks outside of the opsec checks
type OPSEC_ROLE string

const (
	OPSEC_ROLE_LEAD           OPSEC_ROLE = "lead"
	OPSEC_ROLE_OPERATOR                  = "operator"
	OPSEC_ROLE_OTHER_OPERATOR            = "other_operator"
)

type PtTaskFunctionOPSECPre func(PTTaskMessageAllData) PTTTaskOPSECPreTaskMessageResponse
type PTTTaskOPSECPreTaskMessageResponse struct {
	TaskID             int        `json:"task_id"`
	Success            bool       `json:"success"`
	Error              string     `json:"error"`
	OpsecPreBlocked    bool       `json:"opsec_pre_blocked"`
	OpsecPreMessage    string     `json:"opsec_pre_message"`
	OpsecPreBypassed   *bool      `json:"opsec_pre_bypassed,omitempty"`
	OpsecPreBypassRole OPSEC_ROLE `json:"opsec_pre_bypass_role"`
}

// Tasking step 2:
// Task message/process to run the create_tasking function
//
//	this can start creating subtasks
type PtTaskFunctionCreateTasking func(PTTaskMessageAllData) PTTaskCreateTaskingMessageResponse
type PTTaskCreateTaskingMessageResponse struct {
	TaskID                 int     `json:"task_id"`
	Success                bool    `json:"success"`
	Error                  string  `json:"error"`
	CommandName            *string `json:"command_name,omitempty"`
	TaskStatus             *string `json:"task_status,omitempty"`
	DisplayParams          *string `json:"display_params,omitempty"`
	Stdout                 *string `json:"stdout,omitempty"`
	Stderr                 *string `json:"stderr,omitempty"`
	Completed              *bool   `json:"completed,omitempty"`
	TokenID                *int    `json:"token_id,omitempty"`
	CompletionFunctionName *string `json:"completion_function_name,omitempty"`
	Params                 string  `json:"params"`
	ParameterGroupName     string  `json:"parameter_group_name"`
}

// Tasking step 3:
// Task message/process after running create_tasking but before the task can be picked up by an agent
//
//	this is the time to check any artifacts generated from create_tasking
type PtTaskFunctionOPSECPost func(PTTaskMessageAllData) PTTaskOPSECPostTaskMessageResponse
type PTTaskOPSECPostTaskMessageResponse struct {
	TaskID              int        `json:"task_id"`
	Success             bool       `json:"success"`
	Error               string     `json:"error"`
	OpsecPostBlocked    bool       `json:"opsec_post_blocked"`
	OpsecPostMessage    string     `json:"opsec_post_message"`
	OpsecPostBypassed   *bool      `json:"opsec_post_bypassed,omitempty"`
	OpsecPostBypassRole OPSEC_ROLE `json:"opsec_post_bypass_role"`
}

// Tasking step 4:
// Run this when the specified task completes
type SubtaskGroupName = string

type PTTaskCompletionFunctionMessage struct {
	TaskData               PTTaskMessageAllData  `json:"task"`
	SubtaskData            *PTTaskMessageAllData `json:"subtask,omitempty"`
	SubtaskGroup           *SubtaskGroupName     `json:"subtask_group_name,omitempty"`
	CompletionFunctionName string                `json:"function_name"`
}
type PTTaskCompletionFunction func(PTTaskMessageAllData, *PTTaskMessageAllData, *SubtaskGroupName) PTTaskCompletionFunctionMessageResponse
type PTTaskCompletionFunctionMessageResponse struct {
	TaskID                 int     `json:"task_id"`
	ParentTaskId           int     `json:"parent_task_id"`
	Success                bool    `json:"success"`
	Error                  string  `json:"error"`
	TaskStatus             *string `json:"task_status,omitempty"`
	DisplayParams          *string `json:"display_params,omitempty"`
	Stdout                 *string `json:"stdout,omitempty"`
	Stderr                 *string `json:"stderr,omitempty"`
	Completed              *bool   `json:"completed,omitempty"`
	TokenID                *int    `json:"token_id,omitempty"`
	CompletionFunctionName *string `json:"completion_function_name,omitempty"`
	Params                 *string `json:"params,omitempty"`
	ParameterGroupName     *string `json:"parameter_group_name,omitempty"`
}

// Tasking step 5:
// Task message/process to run for more manual processing of a message's response data
type PtTaskProcessResponseMessage struct {
	TaskData PTTaskMessageAllData `json:"task"`
	Response interface{}          `json:"response"`
}
type PtTaskFunctionProcessResponse func(PtTaskProcessResponseMessage) PTTaskProcessResponseMessageResponse
type PTTaskProcessResponseMessageResponse struct {
	TaskID  int    `json:"task_id"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}
