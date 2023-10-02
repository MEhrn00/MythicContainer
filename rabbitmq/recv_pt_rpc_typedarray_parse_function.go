package rabbitmq

import (
	"encoding/json"
	"fmt"
	agentstructs "github.com/MythicMeta/MythicContainer/agent_structs"
	"github.com/MythicMeta/MythicContainer/logging"
)

func init() {
	agentstructs.AllPayloadData.Get("").AddRPCMethod(agentstructs.RabbitmqRPCMethod{
		RabbitmqRoutingKey:         PT_RPC_COMMAND_TYPEDARRAY_PARSE,
		RabbitmqProcessingFunction: processPtRPCTypedArrayParseMessages,
	})
}

func processPtRPCTypedArrayParseMessages(msg []byte) interface{} {
	incomingMessage := agentstructs.PTRPCTypedArrayParseFunctionMessage{}
	response := agentstructs.PTRPCTypedArrayParseMessageResponse{
		Success: false,
	}
	if err := json.Unmarshal(msg, &incomingMessage); err != nil {
		logging.LogError(err, "Failed to unmarshal JSON into struct")
		response.Error = "Failed to unmarshal JSON message into structs"
		return response
	} else {
		for _, command := range agentstructs.AllPayloadData.Get(incomingMessage.PayloadType).GetCommands() {
			if command.Name == incomingMessage.Command {
				for _, param := range command.CommandParameters {
					if incomingMessage.ParameterName == param.Name {
						if param.TypedArrayParseFunction != nil {
							response.TypedArray = param.TypedArrayParseFunction(incomingMessage)
							response.Success = true
							return response
						} else {
							response.TypedArray = [][]string{}
							response.Error = "Function was nil"
							return response
						}
					}
				}
				response.Error = "Failed to find right parameter for command"
				return response
			}
		}
		response.Error = fmt.Sprintf("Failed to find command %s", incomingMessage.Command)
		return response
	}
}
