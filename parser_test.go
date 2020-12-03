package main

import (
	"fmt"
	"testing"
)

func TestOptionStatements(t *testing.T) {
	input := `
"erpc_shim"
@output_dir("erpc_shim")
@types_header("rpc_system_header.h")

program rpc
@group("system")

interface rpc_system {
    rpc_system_version() -> string
    rpc_system_ack(uint8 c) -> uint8
    rpc_wifi_send_ssl_data(uint32 ssl_client, in binary data, uint16 len) -> int32
}

import "rpc_ble_common.erpc"

/** @brief Event type to inform app*/
enum RPC_T_SERVICE_CALLBACK_TYPE
{
    RPC_SERVICE_CALLBACK_TYPE_INDIFICATION_NOTIFICATION = 1,    /**< CCCD update event */
    RPC_SERVICE_CALLBACK_TYPE_READ_CHAR_VALUE = 2,              /**< client read event */
    RPC_SERVICE_CALLBACK_TYPE_WRITE_CHAR_VALUE = 3,             /**< client write event */
};

enum RPC_T_APP_RESULT
{
    RPC_APP_RESULT_SUCCESS                    = 0x0000,
    RPC_APP_RESULT_PENDING                    = 0x0D01,
    RPC_APP_RESULT_ACCEPT                     = 0x0D03,
    RPC_APP_RESULT_REJECT                     = 0x0D04,
    RPC_APP_RESULT_NOT_RELEASE                = 0x0D05,
    RPC_APP_RESULT_PREP_QUEUE_FULL            = 0x0409,
    RPC_APP_RESULT_INVALID_OFFSET             = 0x0407,
    RPC_APP_RESULT_INVALID_VALUE_SIZE         = 0x040D,
    RPC_APP_RESULT_INVALID_PDU                = 0x0404,
    RPC_APP_RESULT_ATTR_NOT_FOUND             = 0x040A,
    RPC_APP_RESULT_ATTR_NOT_LONG              = 0x040B,
    RPC_APP_RESULT_INSUFFICIENT_RESOURCES     = 0x0411,
    RPC_APP_RESULT_APP_ERR                    = 0x0480,
    RPC_APP_RESULT_CCCD_IMPROPERLY_CONFIGURED = 0x04FD,
    RPC_APP_RESULT_PROC_ALREADY_IN_PROGRESS   = 0x04FE
};

interface rpc_wifi_lwip{
    rpc_lwip_accept(int32 s, in binary addr, inout uint32 addrlen) -> int32
    rpc_lwip_getpeername (int32 s, out binary name, inout uint32 namelen) -> int32
    rpc_wifi_send_ssl_data(uint32 ssl_client, in binary data, uint16 len) -> int32
    rpc_func(uint8 u8, uint16 u16, uint32 u32, int8 i8, int16 i16, int32 i32, bool b) -> void
    rpc_wifi_ssl_set_rootCA(string rootCABuff @retain) -> uint32
    rpc_wifi_ssl_get_rootCA(out string rootCABuff @nullable @max_length(3092)) -> uint32
}

struct RPC_T_LOCAL_APPEARANCE
{
    uint16 local_appearance; // hello there
    uint8[2]  padding;
};

struct RPC_T_GAP_CONN_INFO
{ 
    RPC_T_GAP_CONN_STATE conn_state;             //!< Connection state
    RPC_T_GAP_ROLE       role;                   //!< Device role
    uint8[6]             remote_bd;              //!< Remote address
    uint8                remote_bd_type;         //!< Remote address type
};

enum RPC_T_GAP_WHITE_LIST_OP
{
    RPC_GAP_WHITE_LIST_OP_CLEAR = 0,    /**<  Clear white list. */
    RPC_GAP_WHITE_LIST_OP_ADD,          /**<  Add a device to the white list. */
    RPC_GAP_WHITE_LIST_OP_REMOVE        /**<  Remove a device from the white list. */
};

@group("ble_api")
interface rpc_gap{
    rpc_gap_set_param(RPC_T_GAP_PARAM_TYPE param, in binary value) -> RPC_T_GAP_CAUSE
    rpc_gap_get_param(RPC_T_GAP_PARAM_TYPE param, out binary value) -> RPC_T_GAP_CAUSE
    rpc_gap_set_pairable_mode() -> RPC_T_GAP_CAUSE
};

interface rpc_gap_bone{
    rpc_le_bond_set_param(RPC_T_LE_BOND_PARAM_TYPE param, in binary value) -> RPC_T_GAP_CAUSE
    rpc_le_bond_get_param(RPC_T_LE_BOND_PARAM_TYPE param, out binary value) -> RPC_T_GAP_CAUSE
    rpc_le_bond_pair(uint8 conn_id) -> RPC_T_GAP_CAUSE
    rpc_le_bond_get_display_key(uint8 conn_id, out uint32 key) -> RPC_T_GAP_CAUSE
    rpc_le_bond_passkey_input_confirm(uint8 conn_id, uint32 passcode, RPC_T_GAP_CFM_CAUSE cause) -> RPC_T_GAP_CAUSE
    rpc_le_bond_oob_input_confirm(uint8 conn_id, RPC_T_GAP_CFM_CAUSE cause) -> RPC_T_GAP_CAUSE
    rpc_le_bond_just_work_confirm(uint8 conn_id, RPC_T_GAP_CFM_CAUSE cause) -> RPC_T_GAP_CAUSE
    rpc_le_bond_passkey_display_confirm(uint8 conn_id, RPC_T_GAP_CFM_CAUSE cause) -> RPC_T_GAP_CAUSE
    rpc_le_bond_user_confirm(uint8 conn_id, RPC_T_GAP_CFM_CAUSE cause) -> RPC_T_GAP_CAUSE
    rpc_le_bond_cfg_local_key_distribute(uint8 init_dist, uint8 rsp_dist) -> RPC_T_GAP_CAUSE
    rpc_le_bond_clear_all_keys() -> void
    rpc_le_bond_delete_by_idx(uint8 idx) -> RPC_T_GAP_CAUSE
    rpc_le_bond_delete_by_bd(in uint8[6] bd_addr, RPC_T_GAP_REMOTE_ADDR_TYPE bd_type) -> RPC_T_GAP_CAUSE
    rpc_le_bond_get_sec_level(uint8 conn_id, out RPC_T_GAP_SEC_LEVEL sec_type) -> RPC_T_GAP_CAUSE
};

interface rpc_gap_adv {
    rpc_le_adv_set_param(RPC_T_LE_ADV_PARAM_TYPE param, in binary value) -> RPC_T_GAP_CAUSE
    rpc_le_adv_get_param(RPC_T_LE_ADV_PARAM_TYPE param, out binary value) -> RPC_T_GAP_CAUSE
    rpc_le_adv_start() -> RPC_T_GAP_CAUSE
    rpc_le_adv_stop() -> RPC_T_GAP_CAUSE;
    rpc_le_adv_update_param() -> RPC_T_GAP_CAUSE;
};

interface rpc_gap_storage {
    rpc_flash_save_local_name(in RPC_T_LOCAL_NAME p_data) -> uint32
    rpc_flash_load_local_name(out RPC_T_LOCAL_NAME p_data) -> uint32
    rpc_flash_save_local_appearance(in RPC_T_LOCAL_APPEARANCE p_data) -> uint32
    rpc_flash_load_local_appearance(out RPC_T_LOCAL_APPEARANCE p_data) -> uint32
    rpc_le_find_key_entry(in uint8[6] bd_addr, RPC_T_GAP_REMOTE_ADDR_TYPE bd_type) ->  @nullable RPC_T_LE_KEY_ENTRY 
    rpc_le_find_key_entry_by_idx(uint8 idx) ->  @nullable RPC_T_LE_KEY_ENTRY 
    rpc_le_get_bond_dev_num() -> uint8
    rpc_le_get_low_priority_bond() ->  @nullable RPC_T_LE_KEY_ENTRY 
    rpc_le_get_high_priority_bond() ->  @nullable RPC_T_LE_KEY_ENTRY 
    rpc_le_set_high_priority_bond(in uint8[6] bd_addr,  RPC_T_GAP_REMOTE_ADDR_TYPE bd_type) -> bool
    rpc_le_resolve_random_address(in uint8[6] unresolved_addr, inout uint8[6] resolved_addr, inout RPC_T_GAP_IDENT_ADDR_TYPE resolved_addr_type) -> bool
    rpc_le_get_cccd_data(in RPC_T_LE_KEY_ENTRY p_entry, out RPC_T_LE_CCCD p_data) -> bool
    rpc_le_gen_bond_dev(in uint8[6] bd_addr, RPC_T_GAP_REMOTE_ADDR_TYPE bd_type, RPC_T_GAP_LOCAL_ADDR_TYPE local_bd_type, in binary local_ltk, RPC_T_LE_KEY_TYPE key_type, RPC_T_LE_CCCD p_cccd) -> bool
    rpc_le_get_dev_bond_info_len() -> uint16
    rpc_le_set_dev_bond_info(in binary p_data, out bool exist) ->  @nullable RPC_T_LE_KEY_ENTRY 
    rpc_le_get_dev_bond_info(RPC_T_LE_KEY_ENTRY p_entry, out binary p_data) -> bool
};

interface rpc_wifi_lwip{
    rpc_lwip_errno() -> int32
    rpc_netconn_gethostbyname(string name, out binary addr) -> int8
    //rpc_lwip_gethostbyname(string name) -> @nullable binary 
    //rpc_lwip_gethostbyname_r(string name, out binary ret, out binary buf, out binary result @retain, out uint32 h_errnop) -> int32
    //rpc_lwip_getaddrinfo(string name, string servname, in binary hints, out binary res @retain) -> int32
    rpc_dns_gethostbyname_addrtype(string hostname @retain, out binary addr  @retain, uint32 found, in binary callback_arg @nullable @retain, uint8 dns_addrtype) -> int8
}

interface rpc_wifi_callback {
    oneway rpc_wifi_event_callback(binary event)
}
    `
	l := NewLexer(input)
	p := NewParser(l)

	program := p.ParserProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParserProgram() returned nil.")
	}

	for _, stmt := range program.Statements {
		fmt.Printf("%s\n", stmt)
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors.", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q.", msg)
	}
	t.FailNow()
}
