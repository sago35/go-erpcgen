package main

import (
	"encoding/binary"
	"fmt"
)

var (
	debug = true
)

func rpc_system_version() (string, error) {
	if debug {
		fmt.Printf("rpc_system_version()\n")
	}
	msg := startWriteMessage(0x00, 0x01, 0x01, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return "", err
	}

	<-received
	widx := 8
	var result string
	result = string(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_system_ack(c uint8) (uint8, error) {
	if debug {
		fmt.Printf("rpc_system_ack()\n")
	}
	msg := startWriteMessage(0x00, 0x01, 0x02, uint32(seq))

	// c : 1 byte main.GoArgument{In:true, Out:false, Name:"c", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(c>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result uint8
	result = uint8(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_ble_init() (bool, error) {
	if debug {
		fmt.Printf("rpc_ble_init()\n")
	}
	msg := startWriteMessage(0x00, 0x02, 0x01, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	widx := 8
	var result bool
	result = binary.LittleEndian.Uint32(payload[widx:]) == 1

	seq++
	return result, err
}

func rpc_ble_start() error {
	if debug {
		fmt.Printf("rpc_ble_start()\n")
	}
	msg := startWriteMessage(0x00, 0x02, 0x02, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received

	seq++
	return err
}

func rpc_ble_deinit() error {
	if debug {
		fmt.Printf("rpc_ble_deinit()\n")
	}
	msg := startWriteMessage(0x00, 0x02, 0x03, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received

	seq++
	return err
}

func rpc_gap_set_param(param RPC_T_GAP_PARAM_TYPE, value []byte) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_gap_set_param()\n")
	}
	msg := startWriteMessage(0x00, 0x03, 0x01, uint32(seq))

	// param : 4 byte main.GoArgument{In:true, Out:false, Name:"param", Typ:[]string{"RPC_T_GAP_PARAM_TYPE"}, Nullable:false}
	msg = append(msg, byte(param>>0))
	msg = append(msg, byte(param>>8))
	msg = append(msg, byte(param>>16))
	msg = append(msg, byte(param>>24))
	// value : 0 byte main.GoArgument{In:true, Out:false, Name:"value", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, value...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_gap_get_param(param RPC_T_GAP_PARAM_TYPE, value []byte) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_gap_get_param()\n")
	}
	msg := startWriteMessage(0x00, 0x03, 0x02, uint32(seq))

	// param : 4 byte main.GoArgument{In:true, Out:false, Name:"param", Typ:[]string{"RPC_T_GAP_PARAM_TYPE"}, Nullable:false}
	msg = append(msg, byte(param>>0))
	msg = append(msg, byte(param>>8))
	msg = append(msg, byte(param>>16))
	msg = append(msg, byte(param>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// value : 0 byte main.GoArgument{In:false, Out:true, Name:"value", Typ:[]string{"[]byte"}, Nullable:false}
	value_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if value_length > 0 {
		copy(value, payload[widx:widx+int(value_length)])
		widx += int(value_length)
	}

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_gap_set_pairable_mode() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_gap_set_pairable_mode()\n")
	}
	msg := startWriteMessage(0x00, 0x03, 0x03, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_bond_set_param(param RPC_T_LE_BOND_PARAM_TYPE, value []byte) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_bond_set_param()\n")
	}
	msg := startWriteMessage(0x00, 0x04, 0x01, uint32(seq))

	// param : 4 byte main.GoArgument{In:true, Out:false, Name:"param", Typ:[]string{"RPC_T_LE_BOND_PARAM_TYPE"}, Nullable:false}
	msg = append(msg, byte(param>>0))
	msg = append(msg, byte(param>>8))
	msg = append(msg, byte(param>>16))
	msg = append(msg, byte(param>>24))
	// value : 0 byte main.GoArgument{In:true, Out:false, Name:"value", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, value...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_bond_get_param(param RPC_T_LE_BOND_PARAM_TYPE, value []byte) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_bond_get_param()\n")
	}
	msg := startWriteMessage(0x00, 0x04, 0x02, uint32(seq))

	// param : 4 byte main.GoArgument{In:true, Out:false, Name:"param", Typ:[]string{"RPC_T_LE_BOND_PARAM_TYPE"}, Nullable:false}
	msg = append(msg, byte(param>>0))
	msg = append(msg, byte(param>>8))
	msg = append(msg, byte(param>>16))
	msg = append(msg, byte(param>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// value : 0 byte main.GoArgument{In:false, Out:true, Name:"value", Typ:[]string{"[]byte"}, Nullable:false}
	value_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if value_length > 0 {
		copy(value, payload[widx:widx+int(value_length)])
		widx += int(value_length)
	}

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_bond_pair(conn_id uint8) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_bond_pair()\n")
	}
	msg := startWriteMessage(0x00, 0x04, 0x03, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_bond_get_display_key(conn_id uint8, key uint32) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_bond_get_display_key()\n")
	}
	msg := startWriteMessage(0x00, 0x04, 0x04, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// key : 4 byte main.GoArgument{In:false, Out:true, Name:"key", Typ:[]string{"uint32"}, Nullable:false}
// not impl

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_bond_passkey_input_confirm(conn_id uint8, passcode uint32, cause RPC_T_GAP_CFM_CAUSE) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_bond_passkey_input_confirm()\n")
	}
	msg := startWriteMessage(0x00, 0x04, 0x05, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))
	// passcode : 4 byte main.GoArgument{In:true, Out:false, Name:"passcode", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(passcode>>0))
	msg = append(msg, byte(passcode>>8))
	msg = append(msg, byte(passcode>>16))
	msg = append(msg, byte(passcode>>24))
	// cause : 4 byte main.GoArgument{In:true, Out:false, Name:"cause", Typ:[]string{"RPC_T_GAP_CFM_CAUSE"}, Nullable:false}
	msg = append(msg, byte(cause>>0))
	msg = append(msg, byte(cause>>8))
	msg = append(msg, byte(cause>>16))
	msg = append(msg, byte(cause>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_bond_oob_input_confirm(conn_id uint8, cause RPC_T_GAP_CFM_CAUSE) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_bond_oob_input_confirm()\n")
	}
	msg := startWriteMessage(0x00, 0x04, 0x06, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))
	// cause : 4 byte main.GoArgument{In:true, Out:false, Name:"cause", Typ:[]string{"RPC_T_GAP_CFM_CAUSE"}, Nullable:false}
	msg = append(msg, byte(cause>>0))
	msg = append(msg, byte(cause>>8))
	msg = append(msg, byte(cause>>16))
	msg = append(msg, byte(cause>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_bond_just_work_confirm(conn_id uint8, cause RPC_T_GAP_CFM_CAUSE) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_bond_just_work_confirm()\n")
	}
	msg := startWriteMessage(0x00, 0x04, 0x07, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))
	// cause : 4 byte main.GoArgument{In:true, Out:false, Name:"cause", Typ:[]string{"RPC_T_GAP_CFM_CAUSE"}, Nullable:false}
	msg = append(msg, byte(cause>>0))
	msg = append(msg, byte(cause>>8))
	msg = append(msg, byte(cause>>16))
	msg = append(msg, byte(cause>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_bond_passkey_display_confirm(conn_id uint8, cause RPC_T_GAP_CFM_CAUSE) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_bond_passkey_display_confirm()\n")
	}
	msg := startWriteMessage(0x00, 0x04, 0x08, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))
	// cause : 4 byte main.GoArgument{In:true, Out:false, Name:"cause", Typ:[]string{"RPC_T_GAP_CFM_CAUSE"}, Nullable:false}
	msg = append(msg, byte(cause>>0))
	msg = append(msg, byte(cause>>8))
	msg = append(msg, byte(cause>>16))
	msg = append(msg, byte(cause>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_bond_user_confirm(conn_id uint8, cause RPC_T_GAP_CFM_CAUSE) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_bond_user_confirm()\n")
	}
	msg := startWriteMessage(0x00, 0x04, 0x09, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))
	// cause : 4 byte main.GoArgument{In:true, Out:false, Name:"cause", Typ:[]string{"RPC_T_GAP_CFM_CAUSE"}, Nullable:false}
	msg = append(msg, byte(cause>>0))
	msg = append(msg, byte(cause>>8))
	msg = append(msg, byte(cause>>16))
	msg = append(msg, byte(cause>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_bond_cfg_local_key_distribute(init_dist uint8, rsp_dist uint8) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_bond_cfg_local_key_distribute()\n")
	}
	msg := startWriteMessage(0x00, 0x04, 0x0A, uint32(seq))

	// init_dist : 1 byte main.GoArgument{In:true, Out:false, Name:"init_dist", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(init_dist>>0))
	// rsp_dist : 1 byte main.GoArgument{In:true, Out:false, Name:"rsp_dist", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(rsp_dist>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_bond_clear_all_keys() error {
	if debug {
		fmt.Printf("rpc_le_bond_clear_all_keys()\n")
	}
	msg := startWriteMessage(0x00, 0x04, 0x0B, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received

	seq++
	return err
}

func rpc_le_bond_delete_by_idx(idx uint8) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_bond_delete_by_idx()\n")
	}
	msg := startWriteMessage(0x00, 0x04, 0x0C, uint32(seq))

	// idx : 1 byte main.GoArgument{In:true, Out:false, Name:"idx", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(idx>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_bond_delete_by_bd(bd_addr uint8, bd_type RPC_T_GAP_REMOTE_ADDR_TYPE) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_bond_delete_by_bd()\n")
	}
	msg := startWriteMessage(0x00, 0x04, 0x0D, uint32(seq))

	// bd_addr : 1 byte main.GoArgument{In:true, Out:false, Name:"bd_addr", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(bd_addr>>0))
	// bd_type : 4 byte main.GoArgument{In:true, Out:false, Name:"bd_type", Typ:[]string{"RPC_T_GAP_REMOTE_ADDR_TYPE"}, Nullable:false}
	msg = append(msg, byte(bd_type>>0))
	msg = append(msg, byte(bd_type>>8))
	msg = append(msg, byte(bd_type>>16))
	msg = append(msg, byte(bd_type>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_bond_get_sec_level(conn_id uint8, sec_type RPC_T_GAP_SEC_LEVEL) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_bond_get_sec_level()\n")
	}
	msg := startWriteMessage(0x00, 0x04, 0x0E, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// sec_type : 4 byte main.GoArgument{In:false, Out:true, Name:"sec_type", Typ:[]string{"RPC_T_GAP_SEC_LEVEL"}, Nullable:false}
// not impl

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_gap_init(link_num uint8) (bool, error) {
	if debug {
		fmt.Printf("rpc_le_gap_init()\n")
	}
	msg := startWriteMessage(0x00, 0x05, 0x01, uint32(seq))

	// link_num : 1 byte main.GoArgument{In:true, Out:false, Name:"link_num", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(link_num>>0))

	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	widx := 8

	var result bool
	result = binary.LittleEndian.Uint32(payload[widx:]) == 1

	seq++
	return result, err
}

func rpc_le_gap_msg_info_way(use_msg bool) error {
	if debug {
		fmt.Printf("rpc_le_gap_msg_info_way()\n")
	}
	msg := startWriteMessage(0x00, 0x05, 0x02, uint32(seq))

	// use_msg : 1 byte main.GoArgument{In:true, Out:false, Name:"use_msg", Typ:[]string{"bool"}, Nullable:false}
	if use_msg {
		msg = append(msg, 1)
	} else {
		msg = append(msg, 0)
	}

	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received


	seq++
	return err
}

func rpc_le_get_max_link_num() (uint8, error) {
	if debug {
		fmt.Printf("rpc_le_get_max_link_num()\n")
	}
	msg := startWriteMessage(0x00, 0x05, 0x03, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result uint8
	result = uint8(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_set_gap_param(param RPC_T_GAP_LE_PARAM_TYPE, value []byte) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_set_gap_param()\n")
	}
	msg := startWriteMessage(0x00, 0x05, 0x04, uint32(seq))

	// param : 4 byte main.GoArgument{In:true, Out:false, Name:"param", Typ:[]string{"RPC_T_GAP_LE_PARAM_TYPE"}, Nullable:false}
	msg = append(msg, byte(param>>0))
	msg = append(msg, byte(param>>8))
	msg = append(msg, byte(param>>16))
	msg = append(msg, byte(param>>24))
	// value : 0 byte main.GoArgument{In:true, Out:false, Name:"value", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, value...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_get_gap_param(param RPC_T_GAP_LE_PARAM_TYPE, value []byte) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_get_gap_param()\n")
	}
	msg := startWriteMessage(0x00, 0x05, 0x05, uint32(seq))

	// param : 4 byte main.GoArgument{In:true, Out:false, Name:"param", Typ:[]string{"RPC_T_GAP_LE_PARAM_TYPE"}, Nullable:false}
	msg = append(msg, byte(param>>0))
	msg = append(msg, byte(param>>8))
	msg = append(msg, byte(param>>16))
	msg = append(msg, byte(param>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// value : 0 byte main.GoArgument{In:false, Out:true, Name:"value", Typ:[]string{"[]byte"}, Nullable:false}
	value_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if value_length > 0 {
		copy(value, payload[widx:widx+int(value_length)])
		widx += int(value_length)
	}

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_modify_white_list(operation RPC_T_GAP_WHITE_LIST_OP, bd_addr uint8, bd_type RPC_T_GAP_REMOTE_ADDR_TYPE) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_modify_white_list()\n")
	}
	msg := startWriteMessage(0x00, 0x05, 0x06, uint32(seq))

	// operation : 4 byte main.GoArgument{In:true, Out:false, Name:"operation", Typ:[]string{"RPC_T_GAP_WHITE_LIST_OP"}, Nullable:false}
	msg = append(msg, byte(operation>>0))
	msg = append(msg, byte(operation>>8))
	msg = append(msg, byte(operation>>16))
	msg = append(msg, byte(operation>>24))
	// bd_addr : 1 byte main.GoArgument{In:true, Out:false, Name:"bd_addr", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(bd_addr>>0))
	// bd_type : 4 byte main.GoArgument{In:true, Out:false, Name:"bd_type", Typ:[]string{"RPC_T_GAP_REMOTE_ADDR_TYPE"}, Nullable:false}
	msg = append(msg, byte(bd_type>>0))
	msg = append(msg, byte(bd_type>>8))
	msg = append(msg, byte(bd_type>>16))
	msg = append(msg, byte(bd_type>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_gen_rand_addr(rand_addr_type RPC_T_GAP_RAND_ADDR_TYPE, random_bd uint8) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_gen_rand_addr()\n")
	}
	msg := startWriteMessage(0x00, 0x05, 0x07, uint32(seq))

	// rand_addr_type : 4 byte main.GoArgument{In:true, Out:false, Name:"rand_addr_type", Typ:[]string{"RPC_T_GAP_RAND_ADDR_TYPE"}, Nullable:false}
	msg = append(msg, byte(rand_addr_type>>0))
	msg = append(msg, byte(rand_addr_type>>8))
	msg = append(msg, byte(rand_addr_type>>16))
	msg = append(msg, byte(rand_addr_type>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// random_bd : 1 byte main.GoArgument{In:false, Out:true, Name:"random_bd", Typ:[]string{"uint8"}, Nullable:false}
// not impl

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_set_rand_addr(random_bd uint8) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_set_rand_addr()\n")
	}
	msg := startWriteMessage(0x00, 0x05, 0x08, uint32(seq))

	// random_bd : 1 byte main.GoArgument{In:true, Out:false, Name:"random_bd", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(random_bd>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_cfg_local_identity_address(addr uint8, ident_addr_type RPC_T_GAP_IDENT_ADDR_TYPE) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_cfg_local_identity_address()\n")
	}
	msg := startWriteMessage(0x00, 0x05, 0x09, uint32(seq))

	// addr : 1 byte main.GoArgument{In:true, Out:false, Name:"addr", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(addr>>0))
	// ident_addr_type : 4 byte main.GoArgument{In:true, Out:false, Name:"ident_addr_type", Typ:[]string{"RPC_T_GAP_IDENT_ADDR_TYPE"}, Nullable:false}
	msg = append(msg, byte(ident_addr_type>>0))
	msg = append(msg, byte(ident_addr_type>>8))
	msg = append(msg, byte(ident_addr_type>>16))
	msg = append(msg, byte(ident_addr_type>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_set_host_chann_classif(p_channel_map uint8) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_set_host_chann_classif()\n")
	}
	msg := startWriteMessage(0x00, 0x05, 0x0A, uint32(seq))

	// p_channel_map : 1 byte main.GoArgument{In:true, Out:false, Name:"p_channel_map", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(p_channel_map>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_write_default_data_len(tx_octets uint16, tx_time uint16) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_write_default_data_len()\n")
	}
	msg := startWriteMessage(0x00, 0x05, 0x0B, uint32(seq))

	// tx_octets : 2 byte main.GoArgument{In:true, Out:false, Name:"tx_octets", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(tx_octets>>0))
	msg = append(msg, byte(tx_octets>>8))
	// tx_time : 2 byte main.GoArgument{In:true, Out:false, Name:"tx_time", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(tx_time>>0))
	msg = append(msg, byte(tx_time>>8))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_gap_config_cccd_not_check(cccd_not_check_flag RPC_T_GAP_CONFIG_GATT_CCCD_NOT_CHECK) error {
	if debug {
		fmt.Printf("rpc_gap_config_cccd_not_check()\n")
	}
	msg := startWriteMessage(0x00, 0x06, 0x01, uint32(seq))

	// cccd_not_check_flag : 4 byte main.GoArgument{In:true, Out:false, Name:"cccd_not_check_flag", Typ:[]string{"RPC_T_GAP_CONFIG_GATT_CCCD_NOT_CHECK"}, Nullable:false}
	msg = append(msg, byte(cccd_not_check_flag>>0))
	msg = append(msg, byte(cccd_not_check_flag>>8))
	msg = append(msg, byte(cccd_not_check_flag>>16))
	msg = append(msg, byte(cccd_not_check_flag>>24))

	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received


	seq++
	return err
}

func rpc_gap_config_ccc_bits_count(gatt_server_ccc_bits_count uint8, gatt_storage_ccc_bits_count uint8) error {
	if debug {
		fmt.Printf("rpc_gap_config_ccc_bits_count()\n")
	}
	msg := startWriteMessage(0x00, 0x06, 0x02, uint32(seq))

	// gatt_server_ccc_bits_count : 1 byte main.GoArgument{In:true, Out:false, Name:"gatt_server_ccc_bits_count", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(gatt_server_ccc_bits_count>>0))
	// gatt_storage_ccc_bits_count : 1 byte main.GoArgument{In:true, Out:false, Name:"gatt_storage_ccc_bits_count", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(gatt_storage_ccc_bits_count>>0))

	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received


	seq++
	return err
}

func rpc_gap_config_max_attribute_table_count(gatt_max_attribute_table_count uint8) error {
	if debug {
		fmt.Printf("rpc_gap_config_max_attribute_table_count()\n")
	}
	msg := startWriteMessage(0x00, 0x06, 0x03, uint32(seq))

	// gatt_max_attribute_table_count : 1 byte main.GoArgument{In:true, Out:false, Name:"gatt_max_attribute_table_count", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(gatt_max_attribute_table_count>>0))

	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received


	seq++
	return err
}

func rpc_gap_config_max_mtu_size(att_max_mtu_size uint16) error {
	if debug {
		fmt.Printf("rpc_gap_config_max_mtu_size()\n")
	}
	msg := startWriteMessage(0x00, 0x06, 0x04, uint32(seq))

	// att_max_mtu_size : 2 byte main.GoArgument{In:true, Out:false, Name:"att_max_mtu_size", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(att_max_mtu_size>>0))
	msg = append(msg, byte(att_max_mtu_size>>8))

	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received


	seq++
	return err
}

func rpc_gap_config_bte_pool_size(bte_pool_size uint8) error {
	if debug {
		fmt.Printf("rpc_gap_config_bte_pool_size()\n")
	}
	msg := startWriteMessage(0x00, 0x06, 0x05, uint32(seq))

	// bte_pool_size : 1 byte main.GoArgument{In:true, Out:false, Name:"bte_pool_size", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(bte_pool_size>>0))

	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received


	seq++
	return err
}

func rpc_gap_config_bt_report_buf_num(bt_report_buf_num uint8) error {
	if debug {
		fmt.Printf("rpc_gap_config_bt_report_buf_num()\n")
	}
	msg := startWriteMessage(0x00, 0x06, 0x06, uint32(seq))

	// bt_report_buf_num : 1 byte main.GoArgument{In:true, Out:false, Name:"bt_report_buf_num", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(bt_report_buf_num>>0))

	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received


	seq++
	return err
}

func rpc_gap_config_le_key_storage_flag(le_key_storage_flag uint16) error {
	if debug {
		fmt.Printf("rpc_gap_config_le_key_storage_flag()\n")
	}
	msg := startWriteMessage(0x00, 0x06, 0x07, uint32(seq))

	// le_key_storage_flag : 2 byte main.GoArgument{In:true, Out:false, Name:"le_key_storage_flag", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(le_key_storage_flag>>0))
	msg = append(msg, byte(le_key_storage_flag>>8))

	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received


	seq++
	return err
}

func rpc_gap_config_max_le_paired_device(max_le_paired_device uint8) error {
	if debug {
		fmt.Printf("rpc_gap_config_max_le_paired_device()\n")
	}
	msg := startWriteMessage(0x00, 0x06, 0x08, uint32(seq))

	// max_le_paired_device : 1 byte main.GoArgument{In:true, Out:false, Name:"max_le_paired_device", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(max_le_paired_device>>0))

	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received


	seq++
	return err
}

func rpc_gap_config_max_le_link_num(le_link_num uint8) error {
	if debug {
		fmt.Printf("rpc_gap_config_max_le_link_num()\n")
	}
	msg := startWriteMessage(0x00, 0x06, 0x09, uint32(seq))

	// le_link_num : 1 byte main.GoArgument{In:true, Out:false, Name:"le_link_num", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(le_link_num>>0))

	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received


	seq++
	return err
}

func rpc_le_adv_set_param(param RPC_T_LE_ADV_PARAM_TYPE, value []byte) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_adv_set_param()\n")
	}
	msg := startWriteMessage(0x00, 0x07, 0x01, uint32(seq))

	// param : 4 byte main.GoArgument{In:true, Out:false, Name:"param", Typ:[]string{"RPC_T_LE_ADV_PARAM_TYPE"}, Nullable:false}
	msg = append(msg, byte(param>>0))
	msg = append(msg, byte(param>>8))
	msg = append(msg, byte(param>>16))
	msg = append(msg, byte(param>>24))
	// value : 0 byte main.GoArgument{In:true, Out:false, Name:"value", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, value...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_adv_get_param(param RPC_T_LE_ADV_PARAM_TYPE, value []byte) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_adv_get_param()\n")
	}
	msg := startWriteMessage(0x00, 0x07, 0x02, uint32(seq))

	// param : 4 byte main.GoArgument{In:true, Out:false, Name:"param", Typ:[]string{"RPC_T_LE_ADV_PARAM_TYPE"}, Nullable:false}
	msg = append(msg, byte(param>>0))
	msg = append(msg, byte(param>>8))
	msg = append(msg, byte(param>>16))
	msg = append(msg, byte(param>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// value : 0 byte main.GoArgument{In:false, Out:true, Name:"value", Typ:[]string{"[]byte"}, Nullable:false}
	value_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if value_length > 0 {
		copy(value, payload[widx:widx+int(value_length)])
		widx += int(value_length)
	}

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_adv_start() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_adv_start()\n")
	}
	msg := startWriteMessage(0x00, 0x07, 0x03, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_adv_stop() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_adv_stop()\n")
	}
	msg := startWriteMessage(0x00, 0x07, 0x04, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_adv_update_param() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_adv_update_param()\n")
	}
	msg := startWriteMessage(0x00, 0x07, 0x05, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_scan_set_param(param RPC_T_LE_SCAN_PARAM_TYPE, value []byte) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_scan_set_param()\n")
	}
	msg := startWriteMessage(0x00, 0x08, 0x01, uint32(seq))

	// param : 4 byte main.GoArgument{In:true, Out:false, Name:"param", Typ:[]string{"RPC_T_LE_SCAN_PARAM_TYPE"}, Nullable:false}
	msg = append(msg, byte(param>>0))
	msg = append(msg, byte(param>>8))
	msg = append(msg, byte(param>>16))
	msg = append(msg, byte(param>>24))
	// value : 0 byte main.GoArgument{In:true, Out:false, Name:"value", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, value...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_scan_get_param(param RPC_T_LE_SCAN_PARAM_TYPE, value []byte) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_scan_get_param()\n")
	}
	msg := startWriteMessage(0x00, 0x08, 0x02, uint32(seq))

	// param : 4 byte main.GoArgument{In:true, Out:false, Name:"param", Typ:[]string{"RPC_T_LE_SCAN_PARAM_TYPE"}, Nullable:false}
	msg = append(msg, byte(param>>0))
	msg = append(msg, byte(param>>8))
	msg = append(msg, byte(param>>16))
	msg = append(msg, byte(param>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// value : 0 byte main.GoArgument{In:false, Out:true, Name:"value", Typ:[]string{"[]byte"}, Nullable:false}
	value_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if value_length > 0 {
		copy(value, payload[widx:widx+int(value_length)])
		widx += int(value_length)
	}

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_scan_start() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_scan_start()\n")
	}
	msg := startWriteMessage(0x00, 0x08, 0x03, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_scan_timer_start(tick uint32) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_scan_timer_start()\n")
	}
	msg := startWriteMessage(0x00, 0x08, 0x04, uint32(seq))

	// tick : 4 byte main.GoArgument{In:true, Out:false, Name:"tick", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(tick>>0))
	msg = append(msg, byte(tick>>8))
	msg = append(msg, byte(tick>>16))
	msg = append(msg, byte(tick>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_scan_stop() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_scan_stop()\n")
	}
	msg := startWriteMessage(0x00, 0x08, 0x05, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_scan_info_filter(enable bool, offset uint8, len uint8, p_filter uint8) (bool, error) {
	if debug {
		fmt.Printf("rpc_le_scan_info_filter()\n")
	}
	msg := startWriteMessage(0x00, 0x08, 0x06, uint32(seq))

	// enable : 1 byte main.GoArgument{In:true, Out:false, Name:"enable", Typ:[]string{"bool"}, Nullable:false}
	if enable {
		msg = append(msg, 1)
	} else {
		msg = append(msg, 0)
	}
	// offset : 1 byte main.GoArgument{In:true, Out:false, Name:"offset", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(offset>>0))
	// len : 1 byte main.GoArgument{In:true, Out:false, Name:"len", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(len>>0))
	// p_filter : 1 byte main.GoArgument{In:true, Out:false, Name:"p_filter", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(p_filter>>0))

	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	widx := 8

	var result bool
	result = binary.LittleEndian.Uint32(payload[widx:]) == 1

	seq++
	return result, err
}

func rpc_le_get_conn_param(param RPC_T_LE_CONN_PARAM_TYPE, value []byte, conn_id uint8) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_get_conn_param()\n")
	}
	msg := startWriteMessage(0x00, 0x09, 0x01, uint32(seq))

	// param : 4 byte main.GoArgument{In:true, Out:false, Name:"param", Typ:[]string{"RPC_T_LE_CONN_PARAM_TYPE"}, Nullable:false}
	msg = append(msg, byte(param>>0))
	msg = append(msg, byte(param>>8))
	msg = append(msg, byte(param>>16))
	msg = append(msg, byte(param>>24))
	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// value : 0 byte main.GoArgument{In:false, Out:true, Name:"value", Typ:[]string{"[]byte"}, Nullable:false}
	value_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if value_length > 0 {
		copy(value, payload[widx:widx+int(value_length)])
		widx += int(value_length)
	}

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_get_conn_info(conn_id uint8, p_conn_info RPC_T_GAP_CONN_INFO) (bool, error) {
	if debug {
		fmt.Printf("rpc_le_get_conn_info()\n")
	}
	msg := startWriteMessage(0x00, 0x09, 0x02, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))

	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	widx := 8
	// p_conn_info : 4 byte main.GoArgument{In:false, Out:true, Name:"p_conn_info", Typ:[]string{"RPC_T_GAP_CONN_INFO"}, Nullable:false}
// not impl

	var result bool
	result = binary.LittleEndian.Uint32(payload[widx:]) == 1

	seq++
	return result, err
}

func rpc_le_get_conn_addr(conn_id uint8, bd_addr uint8, bd_type uint8) (bool, error) {
	if debug {
		fmt.Printf("rpc_le_get_conn_addr()\n")
	}
	msg := startWriteMessage(0x00, 0x09, 0x03, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))

	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	widx := 8
	// bd_addr : 1 byte main.GoArgument{In:false, Out:true, Name:"bd_addr", Typ:[]string{"uint8"}, Nullable:false}
// not impl
	// bd_type : 1 byte main.GoArgument{In:false, Out:true, Name:"bd_type", Typ:[]string{"uint8"}, Nullable:false}
// not impl

	var result bool
	result = binary.LittleEndian.Uint32(payload[widx:]) == 1

	seq++
	return result, err
}

func rpc_le_get_conn_id(bd_addr uint8, bd_type uint8, p_conn_id uint8) (bool, error) {
	if debug {
		fmt.Printf("rpc_le_get_conn_id()\n")
	}
	msg := startWriteMessage(0x00, 0x09, 0x04, uint32(seq))

	// bd_addr : 1 byte main.GoArgument{In:true, Out:false, Name:"bd_addr", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(bd_addr>>0))
	// bd_type : 1 byte main.GoArgument{In:true, Out:false, Name:"bd_type", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(bd_type>>0))

	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	widx := 8
	// p_conn_id : 1 byte main.GoArgument{In:false, Out:true, Name:"p_conn_id", Typ:[]string{"uint8"}, Nullable:false}
// not impl

	var result bool
	result = binary.LittleEndian.Uint32(payload[widx:]) == 1

	seq++
	return result, err
}

func rpc_le_get_active_link_num() (uint8, error) {
	if debug {
		fmt.Printf("rpc_le_get_active_link_num()\n")
	}
	msg := startWriteMessage(0x00, 0x09, 0x05, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result uint8
	result = uint8(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_get_idle_link_num() (uint8, error) {
	if debug {
		fmt.Printf("rpc_le_get_idle_link_num()\n")
	}
	msg := startWriteMessage(0x00, 0x09, 0x06, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result uint8
	result = uint8(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_disconnect(conn_id uint8) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_disconnect()\n")
	}
	msg := startWriteMessage(0x00, 0x09, 0x07, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_read_rssi(conn_id uint8) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_read_rssi()\n")
	}
	msg := startWriteMessage(0x00, 0x09, 0x08, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_set_data_len(conn_id uint8, tx_octets uint16, tx_time uint16) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_set_data_len()\n")
	}
	msg := startWriteMessage(0x00, 0x09, 0x09, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))
	// tx_octets : 2 byte main.GoArgument{In:true, Out:false, Name:"tx_octets", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(tx_octets>>0))
	msg = append(msg, byte(tx_octets>>8))
	// tx_time : 2 byte main.GoArgument{In:true, Out:false, Name:"tx_time", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(tx_time>>0))
	msg = append(msg, byte(tx_time>>8))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_set_phy(conn_id uint8, all_phys uint8, tx_phys uint8, rx_phys uint8, phy_options RPC_T_GAP_PHYS_OPTIONS) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_set_phy()\n")
	}
	msg := startWriteMessage(0x00, 0x09, 0x0A, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))
	// all_phys : 1 byte main.GoArgument{In:true, Out:false, Name:"all_phys", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(all_phys>>0))
	// tx_phys : 1 byte main.GoArgument{In:true, Out:false, Name:"tx_phys", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(tx_phys>>0))
	// rx_phys : 1 byte main.GoArgument{In:true, Out:false, Name:"rx_phys", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(rx_phys>>0))
	// phy_options : 4 byte main.GoArgument{In:true, Out:false, Name:"phy_options", Typ:[]string{"RPC_T_GAP_PHYS_OPTIONS"}, Nullable:false}
	msg = append(msg, byte(phy_options>>0))
	msg = append(msg, byte(phy_options>>8))
	msg = append(msg, byte(phy_options>>16))
	msg = append(msg, byte(phy_options>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_set_conn_param(conn_type RPC_T_GAP_CONN_PARAM_TYPE, p_conn_param RPC_T_GAP_LE_CONN_REQ_PARAM) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_set_conn_param()\n")
	}
	msg := startWriteMessage(0x00, 0x09, 0x0B, uint32(seq))

	// conn_type : 4 byte main.GoArgument{In:true, Out:false, Name:"conn_type", Typ:[]string{"RPC_T_GAP_CONN_PARAM_TYPE"}, Nullable:false}
	msg = append(msg, byte(conn_type>>0))
	msg = append(msg, byte(conn_type>>8))
	msg = append(msg, byte(conn_type>>16))
	msg = append(msg, byte(conn_type>>24))
	// p_conn_param : 4 byte main.GoArgument{In:true, Out:false, Name:"p_conn_param", Typ:[]string{"RPC_T_GAP_LE_CONN_REQ_PARAM"}, Nullable:false}
	msg = append(msg, byte(p_conn_param>>0))
	msg = append(msg, byte(p_conn_param>>8))
	msg = append(msg, byte(p_conn_param>>16))
	msg = append(msg, byte(p_conn_param>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_connect(init_phys uint8, remote_bd uint8, remote_bd_type RPC_T_GAP_REMOTE_ADDR_TYPE, local_bd_type RPC_T_GAP_LOCAL_ADDR_TYPE, scan_timeout uint16) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_connect()\n")
	}
	msg := startWriteMessage(0x00, 0x09, 0x0C, uint32(seq))

	// init_phys : 1 byte main.GoArgument{In:true, Out:false, Name:"init_phys", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(init_phys>>0))
	// remote_bd : 1 byte main.GoArgument{In:true, Out:false, Name:"remote_bd", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(remote_bd>>0))
	// remote_bd_type : 4 byte main.GoArgument{In:true, Out:false, Name:"remote_bd_type", Typ:[]string{"RPC_T_GAP_REMOTE_ADDR_TYPE"}, Nullable:false}
	msg = append(msg, byte(remote_bd_type>>0))
	msg = append(msg, byte(remote_bd_type>>8))
	msg = append(msg, byte(remote_bd_type>>16))
	msg = append(msg, byte(remote_bd_type>>24))
	// local_bd_type : 4 byte main.GoArgument{In:true, Out:false, Name:"local_bd_type", Typ:[]string{"RPC_T_GAP_LOCAL_ADDR_TYPE"}, Nullable:false}
	msg = append(msg, byte(local_bd_type>>0))
	msg = append(msg, byte(local_bd_type>>8))
	msg = append(msg, byte(local_bd_type>>16))
	msg = append(msg, byte(local_bd_type>>24))
	// scan_timeout : 2 byte main.GoArgument{In:true, Out:false, Name:"scan_timeout", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(scan_timeout>>0))
	msg = append(msg, byte(scan_timeout>>8))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_update_conn_param(conn_id uint8, conn_interval_min uint16, conn_interval_max uint16, conn_latency uint16, supervision_timeout uint16, ce_length_min uint16, ce_length_max uint16) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_update_conn_param()\n")
	}
	msg := startWriteMessage(0x00, 0x09, 0x0D, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))
	// conn_interval_min : 2 byte main.GoArgument{In:true, Out:false, Name:"conn_interval_min", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(conn_interval_min>>0))
	msg = append(msg, byte(conn_interval_min>>8))
	// conn_interval_max : 2 byte main.GoArgument{In:true, Out:false, Name:"conn_interval_max", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(conn_interval_max>>0))
	msg = append(msg, byte(conn_interval_max>>8))
	// conn_latency : 2 byte main.GoArgument{In:true, Out:false, Name:"conn_latency", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(conn_latency>>0))
	msg = append(msg, byte(conn_latency>>8))
	// supervision_timeout : 2 byte main.GoArgument{In:true, Out:false, Name:"supervision_timeout", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(supervision_timeout>>0))
	msg = append(msg, byte(supervision_timeout>>8))
	// ce_length_min : 2 byte main.GoArgument{In:true, Out:false, Name:"ce_length_min", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(ce_length_min>>0))
	msg = append(msg, byte(ce_length_min>>8))
	// ce_length_max : 2 byte main.GoArgument{In:true, Out:false, Name:"ce_length_max", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(ce_length_max>>0))
	msg = append(msg, byte(ce_length_max>>8))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_flash_save_local_name(p_data RPC_T_LOCAL_NAME) (uint32, error) {
	if debug {
		fmt.Printf("rpc_flash_save_local_name()\n")
	}
	msg := startWriteMessage(0x00, 0x0A, 0x01, uint32(seq))

	// p_data : 4 byte main.GoArgument{In:true, Out:false, Name:"p_data", Typ:[]string{"RPC_T_LOCAL_NAME"}, Nullable:false}
	msg = append(msg, byte(p_data>>0))
	msg = append(msg, byte(p_data>>8))
	msg = append(msg, byte(p_data>>16))
	msg = append(msg, byte(p_data>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result uint32
	result = uint32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_flash_load_local_name(p_data RPC_T_LOCAL_NAME) (uint32, error) {
	if debug {
		fmt.Printf("rpc_flash_load_local_name()\n")
	}
	msg := startWriteMessage(0x00, 0x0A, 0x02, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// p_data : 4 byte main.GoArgument{In:false, Out:true, Name:"p_data", Typ:[]string{"RPC_T_LOCAL_NAME"}, Nullable:false}
// not impl

	var result uint32
	result = uint32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_flash_save_local_appearance(p_data RPC_T_LOCAL_APPEARANCE) (uint32, error) {
	if debug {
		fmt.Printf("rpc_flash_save_local_appearance()\n")
	}
	msg := startWriteMessage(0x00, 0x0A, 0x03, uint32(seq))

	// p_data : 4 byte main.GoArgument{In:true, Out:false, Name:"p_data", Typ:[]string{"RPC_T_LOCAL_APPEARANCE"}, Nullable:false}
	msg = append(msg, byte(p_data>>0))
	msg = append(msg, byte(p_data>>8))
	msg = append(msg, byte(p_data>>16))
	msg = append(msg, byte(p_data>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result uint32
	result = uint32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_flash_load_local_appearance(p_data RPC_T_LOCAL_APPEARANCE) (uint32, error) {
	if debug {
		fmt.Printf("rpc_flash_load_local_appearance()\n")
	}
	msg := startWriteMessage(0x00, 0x0A, 0x04, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// p_data : 4 byte main.GoArgument{In:false, Out:true, Name:"p_data", Typ:[]string{"RPC_T_LOCAL_APPEARANCE"}, Nullable:false}
// not impl

	var result uint32
	result = uint32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_find_key_entry(bd_addr uint8, bd_type RPC_T_GAP_REMOTE_ADDR_TYPE) (RPC_T_LE_KEY_ENTRY, error) {
	if debug {
		fmt.Printf("rpc_le_find_key_entry()\n")
	}
	msg := startWriteMessage(0x00, 0x0A, 0x05, uint32(seq))

	// bd_addr : 1 byte main.GoArgument{In:true, Out:false, Name:"bd_addr", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(bd_addr>>0))
	// bd_type : 4 byte main.GoArgument{In:true, Out:false, Name:"bd_type", Typ:[]string{"RPC_T_GAP_REMOTE_ADDR_TYPE"}, Nullable:false}
	msg = append(msg, byte(bd_type>>0))
	msg = append(msg, byte(bd_type>>8))
	msg = append(msg, byte(bd_type>>16))
	msg = append(msg, byte(bd_type>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_LE_KEY_ENTRY
	result = RPC_T_LE_KEY_ENTRY(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_find_key_entry_by_idx(idx uint8) (RPC_T_LE_KEY_ENTRY, error) {
	if debug {
		fmt.Printf("rpc_le_find_key_entry_by_idx()\n")
	}
	msg := startWriteMessage(0x00, 0x0A, 0x06, uint32(seq))

	// idx : 1 byte main.GoArgument{In:true, Out:false, Name:"idx", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(idx>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_LE_KEY_ENTRY
	result = RPC_T_LE_KEY_ENTRY(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_get_bond_dev_num() (uint8, error) {
	if debug {
		fmt.Printf("rpc_le_get_bond_dev_num()\n")
	}
	msg := startWriteMessage(0x00, 0x0A, 0x07, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result uint8
	result = uint8(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_get_low_priority_bond() (RPC_T_LE_KEY_ENTRY, error) {
	if debug {
		fmt.Printf("rpc_le_get_low_priority_bond()\n")
	}
	msg := startWriteMessage(0x00, 0x0A, 0x08, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result RPC_T_LE_KEY_ENTRY
	result = RPC_T_LE_KEY_ENTRY(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_get_high_priority_bond() (RPC_T_LE_KEY_ENTRY, error) {
	if debug {
		fmt.Printf("rpc_le_get_high_priority_bond()\n")
	}
	msg := startWriteMessage(0x00, 0x0A, 0x09, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result RPC_T_LE_KEY_ENTRY
	result = RPC_T_LE_KEY_ENTRY(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_set_high_priority_bond(bd_addr uint8, bd_type RPC_T_GAP_REMOTE_ADDR_TYPE) (bool, error) {
	if debug {
		fmt.Printf("rpc_le_set_high_priority_bond()\n")
	}
	msg := startWriteMessage(0x00, 0x0A, 0x0A, uint32(seq))

	// bd_addr : 1 byte main.GoArgument{In:true, Out:false, Name:"bd_addr", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(bd_addr>>0))
	// bd_type : 4 byte main.GoArgument{In:true, Out:false, Name:"bd_type", Typ:[]string{"RPC_T_GAP_REMOTE_ADDR_TYPE"}, Nullable:false}
	msg = append(msg, byte(bd_type>>0))
	msg = append(msg, byte(bd_type>>8))
	msg = append(msg, byte(bd_type>>16))
	msg = append(msg, byte(bd_type>>24))

	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	widx := 8

	var result bool
	result = binary.LittleEndian.Uint32(payload[widx:]) == 1

	seq++
	return result, err
}

func rpc_le_resolve_random_address(unresolved_addr uint8, resolved_addr uint8, resolved_addr_type RPC_T_GAP_IDENT_ADDR_TYPE) (bool, error) {
	if debug {
		fmt.Printf("rpc_le_resolve_random_address()\n")
	}
	msg := startWriteMessage(0x00, 0x0A, 0x0B, uint32(seq))

	// unresolved_addr : 1 byte main.GoArgument{In:true, Out:false, Name:"unresolved_addr", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(unresolved_addr>>0))
	// resolved_addr : 1 byte main.GoArgument{In:true, Out:true, Name:"resolved_addr", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(resolved_addr>>0))
	// resolved_addr_type : 4 byte main.GoArgument{In:true, Out:true, Name:"resolved_addr_type", Typ:[]string{"RPC_T_GAP_IDENT_ADDR_TYPE"}, Nullable:false}
	msg = append(msg, byte(resolved_addr_type>>0))
	msg = append(msg, byte(resolved_addr_type>>8))
	msg = append(msg, byte(resolved_addr_type>>16))
	msg = append(msg, byte(resolved_addr_type>>24))

	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	widx := 8
	// resolved_addr : 1 byte main.GoArgument{In:true, Out:true, Name:"resolved_addr", Typ:[]string{"uint8"}, Nullable:false}
// not impl
	// resolved_addr_type : 4 byte main.GoArgument{In:true, Out:true, Name:"resolved_addr_type", Typ:[]string{"RPC_T_GAP_IDENT_ADDR_TYPE"}, Nullable:false}
// not impl

	var result bool
	result = binary.LittleEndian.Uint32(payload[widx:]) == 1

	seq++
	return result, err
}

func rpc_le_get_cccd_data(p_entry RPC_T_LE_KEY_ENTRY, p_data RPC_T_LE_CCCD) (bool, error) {
	if debug {
		fmt.Printf("rpc_le_get_cccd_data()\n")
	}
	msg := startWriteMessage(0x00, 0x0A, 0x0C, uint32(seq))

	// p_entry : 4 byte main.GoArgument{In:true, Out:false, Name:"p_entry", Typ:[]string{"RPC_T_LE_KEY_ENTRY"}, Nullable:false}
	msg = append(msg, byte(p_entry>>0))
	msg = append(msg, byte(p_entry>>8))
	msg = append(msg, byte(p_entry>>16))
	msg = append(msg, byte(p_entry>>24))

	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	widx := 8
	// p_data : 4 byte main.GoArgument{In:false, Out:true, Name:"p_data", Typ:[]string{"RPC_T_LE_CCCD"}, Nullable:false}
// not impl

	var result bool
	result = binary.LittleEndian.Uint32(payload[widx:]) == 1

	seq++
	return result, err
}

func rpc_le_gen_bond_dev(bd_addr uint8, bd_type RPC_T_GAP_REMOTE_ADDR_TYPE, local_bd_type RPC_T_GAP_LOCAL_ADDR_TYPE, local_ltk []byte, key_type RPC_T_LE_KEY_TYPE, p_cccd RPC_T_LE_CCCD) (bool, error) {
	if debug {
		fmt.Printf("rpc_le_gen_bond_dev()\n")
	}
	msg := startWriteMessage(0x00, 0x0A, 0x0D, uint32(seq))

	// bd_addr : 1 byte main.GoArgument{In:true, Out:false, Name:"bd_addr", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(bd_addr>>0))
	// bd_type : 4 byte main.GoArgument{In:true, Out:false, Name:"bd_type", Typ:[]string{"RPC_T_GAP_REMOTE_ADDR_TYPE"}, Nullable:false}
	msg = append(msg, byte(bd_type>>0))
	msg = append(msg, byte(bd_type>>8))
	msg = append(msg, byte(bd_type>>16))
	msg = append(msg, byte(bd_type>>24))
	// local_bd_type : 4 byte main.GoArgument{In:true, Out:false, Name:"local_bd_type", Typ:[]string{"RPC_T_GAP_LOCAL_ADDR_TYPE"}, Nullable:false}
	msg = append(msg, byte(local_bd_type>>0))
	msg = append(msg, byte(local_bd_type>>8))
	msg = append(msg, byte(local_bd_type>>16))
	msg = append(msg, byte(local_bd_type>>24))
	// local_ltk : 0 byte main.GoArgument{In:true, Out:false, Name:"local_ltk", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, local_ltk...)
	// key_type : 4 byte main.GoArgument{In:true, Out:false, Name:"key_type", Typ:[]string{"RPC_T_LE_KEY_TYPE"}, Nullable:false}
	msg = append(msg, byte(key_type>>0))
	msg = append(msg, byte(key_type>>8))
	msg = append(msg, byte(key_type>>16))
	msg = append(msg, byte(key_type>>24))
	// p_cccd : 4 byte main.GoArgument{In:true, Out:false, Name:"p_cccd", Typ:[]string{"RPC_T_LE_CCCD"}, Nullable:false}
	msg = append(msg, byte(p_cccd>>0))
	msg = append(msg, byte(p_cccd>>8))
	msg = append(msg, byte(p_cccd>>16))
	msg = append(msg, byte(p_cccd>>24))

	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	widx := 8

	var result bool
	result = binary.LittleEndian.Uint32(payload[widx:]) == 1

	seq++
	return result, err
}

func rpc_le_get_dev_bond_info_len() (uint16, error) {
	if debug {
		fmt.Printf("rpc_le_get_dev_bond_info_len()\n")
	}
	msg := startWriteMessage(0x00, 0x0A, 0x0E, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result uint16
	result = uint16(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_set_dev_bond_info(p_data []byte, exist bool) (RPC_T_LE_KEY_ENTRY, error) {
	if debug {
		fmt.Printf("rpc_le_set_dev_bond_info()\n")
	}
	msg := startWriteMessage(0x00, 0x0A, 0x0F, uint32(seq))

	// p_data : 0 byte main.GoArgument{In:true, Out:false, Name:"p_data", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, p_data...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// exist : 1 byte main.GoArgument{In:false, Out:true, Name:"exist", Typ:[]string{"bool"}, Nullable:false}
// not impl

	var result RPC_T_LE_KEY_ENTRY
	result = RPC_T_LE_KEY_ENTRY(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_le_get_dev_bond_info(p_entry RPC_T_LE_KEY_ENTRY, p_data []byte) (bool, error) {
	if debug {
		fmt.Printf("rpc_le_get_dev_bond_info()\n")
	}
	msg := startWriteMessage(0x00, 0x0A, 0x10, uint32(seq))

	// p_entry : 4 byte main.GoArgument{In:true, Out:false, Name:"p_entry", Typ:[]string{"RPC_T_LE_KEY_ENTRY"}, Nullable:false}
	msg = append(msg, byte(p_entry>>0))
	msg = append(msg, byte(p_entry>>8))
	msg = append(msg, byte(p_entry>>16))
	msg = append(msg, byte(p_entry>>24))

	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	widx := 8
	// p_data : 0 byte main.GoArgument{In:false, Out:true, Name:"p_data", Typ:[]string{"[]byte"}, Nullable:false}
	p_data_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if p_data_length > 0 {
		copy(p_data, payload[widx:widx+int(p_data_length)])
		widx += int(p_data_length)
	}

	var result bool
	result = binary.LittleEndian.Uint32(payload[widx:]) == 1

	seq++
	return result, err
}

func rpc_ble_client_init(num uint8) (bool, error) {
	if debug {
		fmt.Printf("rpc_ble_client_init()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x01, uint32(seq))

	// num : 1 byte main.GoArgument{In:true, Out:false, Name:"num", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(num>>0))

	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	widx := 8

	var result bool
	result = binary.LittleEndian.Uint32(payload[widx:]) == 1

	seq++
	return result, err
}

func rpc_ble_add_client(app_id uint8, link_num uint8) (uint8, error) {
	if debug {
		fmt.Printf("rpc_ble_add_client()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x02, uint32(seq))

	// app_id : 1 byte main.GoArgument{In:true, Out:false, Name:"app_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(app_id>>0))
	// link_num : 1 byte main.GoArgument{In:true, Out:false, Name:"link_num", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(link_num>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result uint8
	result = uint8(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_client_init(client_num uint8) error {
	if debug {
		fmt.Printf("rpc_client_init()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x03, uint32(seq))

	// client_num : 1 byte main.GoArgument{In:true, Out:false, Name:"client_num", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(client_num>>0))

	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received


	seq++
	return err
}

func rpc_client_all_primary_srv_discovery(conn_id uint8, client_id uint8) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_client_all_primary_srv_discovery()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x04, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))
	// client_id : 1 byte main.GoArgument{In:true, Out:false, Name:"client_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(client_id>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_client_by_uuid_srv_discovery(conn_id uint8, client_id uint8, uuid16 uint16) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_client_by_uuid_srv_discovery()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x05, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))
	// client_id : 1 byte main.GoArgument{In:true, Out:false, Name:"client_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(client_id>>0))
	// uuid16 : 2 byte main.GoArgument{In:true, Out:false, Name:"uuid16", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(uuid16>>0))
	msg = append(msg, byte(uuid16>>8))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_client_by_uuid128_srv_discovery(conn_id uint8, client_id uint8, p_uuid128 uint8) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_client_by_uuid128_srv_discovery()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x06, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))
	// client_id : 1 byte main.GoArgument{In:true, Out:false, Name:"client_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(client_id>>0))
	// p_uuid128 : 1 byte main.GoArgument{In:true, Out:false, Name:"p_uuid128", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(p_uuid128>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_client_relationship_discovery(conn_id uint8, client_id uint8, start_handle uint16, end_handle uint16) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_client_relationship_discovery()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x07, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))
	// client_id : 1 byte main.GoArgument{In:true, Out:false, Name:"client_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(client_id>>0))
	// start_handle : 2 byte main.GoArgument{In:true, Out:false, Name:"start_handle", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(start_handle>>0))
	msg = append(msg, byte(start_handle>>8))
	// end_handle : 2 byte main.GoArgument{In:true, Out:false, Name:"end_handle", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(end_handle>>0))
	msg = append(msg, byte(end_handle>>8))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_client_all_char_discovery(conn_id uint8, client_id uint8, start_handle uint16, end_handle uint16) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_client_all_char_discovery()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x08, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))
	// client_id : 1 byte main.GoArgument{In:true, Out:false, Name:"client_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(client_id>>0))
	// start_handle : 2 byte main.GoArgument{In:true, Out:false, Name:"start_handle", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(start_handle>>0))
	msg = append(msg, byte(start_handle>>8))
	// end_handle : 2 byte main.GoArgument{In:true, Out:false, Name:"end_handle", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(end_handle>>0))
	msg = append(msg, byte(end_handle>>8))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_client_by_uuid_char_discovery(conn_id uint8, client_id uint8, start_handle uint16, end_handle uint16, uuid16 uint16) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_client_by_uuid_char_discovery()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x09, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))
	// client_id : 1 byte main.GoArgument{In:true, Out:false, Name:"client_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(client_id>>0))
	// start_handle : 2 byte main.GoArgument{In:true, Out:false, Name:"start_handle", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(start_handle>>0))
	msg = append(msg, byte(start_handle>>8))
	// end_handle : 2 byte main.GoArgument{In:true, Out:false, Name:"end_handle", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(end_handle>>0))
	msg = append(msg, byte(end_handle>>8))
	// uuid16 : 2 byte main.GoArgument{In:true, Out:false, Name:"uuid16", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(uuid16>>0))
	msg = append(msg, byte(uuid16>>8))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_client_by_uuid128_char_discovery(conn_id uint8, client_id uint8, start_handle uint16, end_handle uint16, p_uuid128 uint8) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_client_by_uuid128_char_discovery()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x0A, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))
	// client_id : 1 byte main.GoArgument{In:true, Out:false, Name:"client_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(client_id>>0))
	// start_handle : 2 byte main.GoArgument{In:true, Out:false, Name:"start_handle", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(start_handle>>0))
	msg = append(msg, byte(start_handle>>8))
	// end_handle : 2 byte main.GoArgument{In:true, Out:false, Name:"end_handle", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(end_handle>>0))
	msg = append(msg, byte(end_handle>>8))
	// p_uuid128 : 1 byte main.GoArgument{In:true, Out:false, Name:"p_uuid128", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(p_uuid128>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_client_all_char_descriptor_discovery(conn_id uint8, client_id uint8, start_handle uint16, end_handle uint16) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_client_all_char_descriptor_discovery()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x0B, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))
	// client_id : 1 byte main.GoArgument{In:true, Out:false, Name:"client_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(client_id>>0))
	// start_handle : 2 byte main.GoArgument{In:true, Out:false, Name:"start_handle", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(start_handle>>0))
	msg = append(msg, byte(start_handle>>8))
	// end_handle : 2 byte main.GoArgument{In:true, Out:false, Name:"end_handle", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(end_handle>>0))
	msg = append(msg, byte(end_handle>>8))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_client_attr_read(conn_id uint8, client_id uint8, handle uint16) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_client_attr_read()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x0C, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))
	// client_id : 1 byte main.GoArgument{In:true, Out:false, Name:"client_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(client_id>>0))
	// handle : 2 byte main.GoArgument{In:true, Out:false, Name:"handle", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(handle>>0))
	msg = append(msg, byte(handle>>8))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_client_attr_read_using_uuid(conn_id uint8, client_id uint8, start_handle uint16, end_handle uint16, uuid16 uint16, p_uuid128 uint8) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_client_attr_read_using_uuid()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x0D, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))
	// client_id : 1 byte main.GoArgument{In:true, Out:false, Name:"client_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(client_id>>0))
	// start_handle : 2 byte main.GoArgument{In:true, Out:false, Name:"start_handle", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(start_handle>>0))
	msg = append(msg, byte(start_handle>>8))
	// end_handle : 2 byte main.GoArgument{In:true, Out:false, Name:"end_handle", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(end_handle>>0))
	msg = append(msg, byte(end_handle>>8))
	// uuid16 : 2 byte main.GoArgument{In:true, Out:false, Name:"uuid16", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(uuid16>>0))
	msg = append(msg, byte(uuid16>>8))
	// p_uuid128 : 1 byte main.GoArgument{In:true, Out:false, Name:"p_uuid128", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(p_uuid128>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_client_attr_write(conn_id uint8, client_id uint8, write_type RPC_T_GATT_WRITE_TYPE, handle uint16, data []byte) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_client_attr_write()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x0E, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))
	// client_id : 1 byte main.GoArgument{In:true, Out:false, Name:"client_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(client_id>>0))
	// write_type : 4 byte main.GoArgument{In:true, Out:false, Name:"write_type", Typ:[]string{"RPC_T_GATT_WRITE_TYPE"}, Nullable:false}
	msg = append(msg, byte(write_type>>0))
	msg = append(msg, byte(write_type>>8))
	msg = append(msg, byte(write_type>>16))
	msg = append(msg, byte(write_type>>24))
	// handle : 2 byte main.GoArgument{In:true, Out:false, Name:"handle", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(handle>>0))
	msg = append(msg, byte(handle>>8))
	// data : 0 byte main.GoArgument{In:true, Out:false, Name:"data", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, data...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_client_attr_ind_confirm(conn_id uint8) (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_client_attr_ind_confirm()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x0F, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_GAP_CAUSE
	result = RPC_T_GAP_CAUSE(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_ble_server_init(num uint8) (bool, error) {
	if debug {
		fmt.Printf("rpc_ble_server_init()\n")
	}
	msg := startWriteMessage(0x00, 0x0C, 0x01, uint32(seq))

	// num : 1 byte main.GoArgument{In:true, Out:false, Name:"num", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(num>>0))

	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	widx := 8

	var result bool
	result = binary.LittleEndian.Uint32(payload[widx:]) == 1

	seq++
	return result, err
}

func rpc_ble_create_service(uuid uint8, uuid_length uint8, is_primary bool) (uint8, error) {
	if debug {
		fmt.Printf("rpc_ble_create_service()\n")
	}
	msg := startWriteMessage(0x00, 0x0C, 0x02, uint32(seq))

	// uuid : 1 byte main.GoArgument{In:true, Out:false, Name:"uuid", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(uuid>>0))
	// uuid_length : 1 byte main.GoArgument{In:true, Out:false, Name:"uuid_length", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(uuid_length>>0))
	// is_primary : 1 byte main.GoArgument{In:true, Out:false, Name:"is_primary", Typ:[]string{"bool"}, Nullable:false}
	if is_primary {
		msg = append(msg, 1)
	} else {
		msg = append(msg, 0)
	}

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result uint8
	result = uint8(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_ble_delete_service(app_id uint8) (bool, error) {
	if debug {
		fmt.Printf("rpc_ble_delete_service()\n")
	}
	msg := startWriteMessage(0x00, 0x0C, 0x03, uint32(seq))

	// app_id : 1 byte main.GoArgument{In:true, Out:false, Name:"app_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(app_id>>0))

	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	widx := 8

	var result bool
	result = binary.LittleEndian.Uint32(payload[widx:]) == 1

	seq++
	return result, err
}

func rpc_ble_service_start(app_id uint8) (uint8, error) {
	if debug {
		fmt.Printf("rpc_ble_service_start()\n")
	}
	msg := startWriteMessage(0x00, 0x0C, 0x04, uint32(seq))

	// app_id : 1 byte main.GoArgument{In:true, Out:false, Name:"app_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(app_id>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result uint8
	result = uint8(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_ble_get_servie_handle(app_id uint8) (uint8, error) {
	if debug {
		fmt.Printf("rpc_ble_get_servie_handle()\n")
	}
	msg := startWriteMessage(0x00, 0x0C, 0x05, uint32(seq))

	// app_id : 1 byte main.GoArgument{In:true, Out:false, Name:"app_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(app_id>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result uint8
	result = uint8(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_ble_create_char(app_id uint8, uuid uint8, uuid_length uint8, properties uint8, permissions uint32) (uint16, error) {
	if debug {
		fmt.Printf("rpc_ble_create_char()\n")
	}
	msg := startWriteMessage(0x00, 0x0C, 0x06, uint32(seq))

	// app_id : 1 byte main.GoArgument{In:true, Out:false, Name:"app_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(app_id>>0))
	// uuid : 1 byte main.GoArgument{In:true, Out:false, Name:"uuid", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(uuid>>0))
	// uuid_length : 1 byte main.GoArgument{In:true, Out:false, Name:"uuid_length", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(uuid_length>>0))
	// properties : 1 byte main.GoArgument{In:true, Out:false, Name:"properties", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(properties>>0))
	// permissions : 4 byte main.GoArgument{In:true, Out:false, Name:"permissions", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(permissions>>0))
	msg = append(msg, byte(permissions>>8))
	msg = append(msg, byte(permissions>>16))
	msg = append(msg, byte(permissions>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result uint16
	result = uint16(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_ble_create_desc(app_id uint8, char_handle uint16, uuid uint8, uuid_length uint8, flags uint8, permissions uint32, value_length uint16, p_value []byte) (uint16, error) {
	if debug {
		fmt.Printf("rpc_ble_create_desc()\n")
	}
	msg := startWriteMessage(0x00, 0x0C, 0x07, uint32(seq))

	// app_id : 1 byte main.GoArgument{In:true, Out:false, Name:"app_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(app_id>>0))
	// char_handle : 2 byte main.GoArgument{In:true, Out:false, Name:"char_handle", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(char_handle>>0))
	msg = append(msg, byte(char_handle>>8))
	// uuid : 1 byte main.GoArgument{In:true, Out:false, Name:"uuid", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(uuid>>0))
	// uuid_length : 1 byte main.GoArgument{In:true, Out:false, Name:"uuid_length", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(uuid_length>>0))
	// flags : 1 byte main.GoArgument{In:true, Out:false, Name:"flags", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(flags>>0))
	// permissions : 4 byte main.GoArgument{In:true, Out:false, Name:"permissions", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(permissions>>0))
	msg = append(msg, byte(permissions>>8))
	msg = append(msg, byte(permissions>>16))
	msg = append(msg, byte(permissions>>24))
	// value_length : 2 byte main.GoArgument{In:true, Out:false, Name:"value_length", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(value_length>>0))
	msg = append(msg, byte(value_length>>8))
	// p_value : 0 byte main.GoArgument{In:true, Out:false, Name:"p_value", Typ:[]string{"[]byte"}, Nullable:true}
	msg = append(msg, p_value...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result uint16
	result = uint16(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_server_send_data(conn_id uint8, service_id uint8, attrib_index uint16, data []byte, pdu_type RPC_T_GATT_PDU_TYPE) (bool, error) {
	if debug {
		fmt.Printf("rpc_server_send_data()\n")
	}
	msg := startWriteMessage(0x00, 0x0C, 0x08, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))
	// service_id : 1 byte main.GoArgument{In:true, Out:false, Name:"service_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(service_id>>0))
	// attrib_index : 2 byte main.GoArgument{In:true, Out:false, Name:"attrib_index", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(attrib_index>>0))
	msg = append(msg, byte(attrib_index>>8))
	// data : 0 byte main.GoArgument{In:true, Out:false, Name:"data", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, data...)
	// pdu_type : 4 byte main.GoArgument{In:true, Out:false, Name:"pdu_type", Typ:[]string{"RPC_T_GATT_PDU_TYPE"}, Nullable:false}
	msg = append(msg, byte(pdu_type>>0))
	msg = append(msg, byte(pdu_type>>8))
	msg = append(msg, byte(pdu_type>>16))
	msg = append(msg, byte(pdu_type>>24))

	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	widx := 8

	var result bool
	result = binary.LittleEndian.Uint32(payload[widx:]) == 1

	seq++
	return result, err
}

func rpc_ble_server_get_attr_value(app_id uint8, attr_handle uint16) ([]byte, error) {
	if debug {
		fmt.Printf("rpc_ble_server_get_attr_value()\n")
	}
	msg := startWriteMessage(0x00, 0x0C, 0x09, uint32(seq))

	// app_id : 1 byte main.GoArgument{In:true, Out:false, Name:"app_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(app_id>>0))
	// attr_handle : 2 byte main.GoArgument{In:true, Out:false, Name:"attr_handle", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(attr_handle>>0))
	msg = append(msg, byte(attr_handle>>8))

	err := performRequest(msg)
	if err != nil {
		return nil, err
	}

	<-received
	widx := 8

	var result []byte
	result_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	copy(result, payload[widx:result_length])

	seq++
	return result, err
}

func rpc_server_exec_write_confirm(conn_id uint8, cause uint16, handle uint16) (bool, error) {
	if debug {
		fmt.Printf("rpc_server_exec_write_confirm()\n")
	}
	msg := startWriteMessage(0x00, 0x0C, 0x0A, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))
	// cause : 2 byte main.GoArgument{In:true, Out:false, Name:"cause", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(cause>>0))
	msg = append(msg, byte(cause>>8))
	// handle : 2 byte main.GoArgument{In:true, Out:false, Name:"handle", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(handle>>0))
	msg = append(msg, byte(handle>>8))

	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	widx := 8

	var result bool
	result = binary.LittleEndian.Uint32(payload[widx:]) == 1

	seq++
	return result, err
}

func rpc_server_attr_write_confirm(conn_id uint8, service_id uint8, attrib_index uint16, cause RPC_T_APP_RESULT) (bool, error) {
	if debug {
		fmt.Printf("rpc_server_attr_write_confirm()\n")
	}
	msg := startWriteMessage(0x00, 0x0C, 0x0B, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))
	// service_id : 1 byte main.GoArgument{In:true, Out:false, Name:"service_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(service_id>>0))
	// attrib_index : 2 byte main.GoArgument{In:true, Out:false, Name:"attrib_index", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(attrib_index>>0))
	msg = append(msg, byte(attrib_index>>8))
	// cause : 4 byte main.GoArgument{In:true, Out:false, Name:"cause", Typ:[]string{"RPC_T_APP_RESULT"}, Nullable:false}
	msg = append(msg, byte(cause>>0))
	msg = append(msg, byte(cause>>8))
	msg = append(msg, byte(cause>>16))
	msg = append(msg, byte(cause>>24))

	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	widx := 8

	var result bool
	result = binary.LittleEndian.Uint32(payload[widx:]) == 1

	seq++
	return result, err
}

func rpc_server_attr_read_confirm(conn_id uint8, service_id uint8, attrib_index uint16, data []byte, cause RPC_T_APP_RESULT) (bool, error) {
	if debug {
		fmt.Printf("rpc_server_attr_read_confirm()\n")
	}
	msg := startWriteMessage(0x00, 0x0C, 0x0C, uint32(seq))

	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))
	// service_id : 1 byte main.GoArgument{In:true, Out:false, Name:"service_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(service_id>>0))
	// attrib_index : 2 byte main.GoArgument{In:true, Out:false, Name:"attrib_index", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(attrib_index>>0))
	msg = append(msg, byte(attrib_index>>8))
	// data : 0 byte main.GoArgument{In:true, Out:false, Name:"data", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, data...)
	// cause : 4 byte main.GoArgument{In:true, Out:false, Name:"cause", Typ:[]string{"RPC_T_APP_RESULT"}, Nullable:false}
	msg = append(msg, byte(cause>>0))
	msg = append(msg, byte(cause>>8))
	msg = append(msg, byte(cause>>16))
	msg = append(msg, byte(cause>>24))

	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	widx := 8

	var result bool
	result = binary.LittleEndian.Uint32(payload[widx:]) == 1

	seq++
	return result, err
}

func rpc_ble_handle_gap_msg(gap_msg []byte) (RPC_T_APP_RESULT, error) {
	if debug {
		fmt.Printf("rpc_ble_handle_gap_msg()\n")
	}
	msg := startWriteMessage(0x00, 0x0D, 0x01, uint32(seq))

	// gap_msg : 0 byte main.GoArgument{In:true, Out:false, Name:"gap_msg", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, gap_msg...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_APP_RESULT
	result = RPC_T_APP_RESULT(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_ble_gap_callback(cb_type uint8, cb_data []byte) (RPC_T_APP_RESULT, error) {
	if debug {
		fmt.Printf("rpc_ble_gap_callback()\n")
	}
	msg := startWriteMessage(0x00, 0x0D, 0x02, uint32(seq))

	// cb_type : 1 byte main.GoArgument{In:true, Out:false, Name:"cb_type", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(cb_type>>0))
	// cb_data : 0 byte main.GoArgument{In:true, Out:false, Name:"cb_data", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, cb_data...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_APP_RESULT
	result = RPC_T_APP_RESULT(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_ble_gattc_callback(gatt_if uint8, conn_id uint8, cb_data []byte, extra_data []byte) (RPC_T_APP_RESULT, error) {
	if debug {
		fmt.Printf("rpc_ble_gattc_callback()\n")
	}
	msg := startWriteMessage(0x00, 0x0D, 0x03, uint32(seq))

	// gatt_if : 1 byte main.GoArgument{In:true, Out:false, Name:"gatt_if", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(gatt_if>>0))
	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))
	// cb_data : 0 byte main.GoArgument{In:true, Out:false, Name:"cb_data", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, cb_data...)
	// extra_data : 0 byte main.GoArgument{In:true, Out:false, Name:"extra_data", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, extra_data...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result RPC_T_APP_RESULT
	result = RPC_T_APP_RESULT(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_ble_gatts_callback(gatt_if uint8, conn_id uint8, attrib_index uint16, event RPC_T_SERVICE_CALLBACK_TYPE, property uint16, read_cb_data []byte, write_cb_data []byte, app_cb_data []byte) (RPC_T_APP_RESULT, error) {
	if debug {
		fmt.Printf("rpc_ble_gatts_callback()\n")
	}
	msg := startWriteMessage(0x00, 0x0D, 0x04, uint32(seq))

	// gatt_if : 1 byte main.GoArgument{In:true, Out:false, Name:"gatt_if", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(gatt_if>>0))
	// conn_id : 1 byte main.GoArgument{In:true, Out:false, Name:"conn_id", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(conn_id>>0))
	// attrib_index : 2 byte main.GoArgument{In:true, Out:false, Name:"attrib_index", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(attrib_index>>0))
	msg = append(msg, byte(attrib_index>>8))
	// event : 4 byte main.GoArgument{In:true, Out:false, Name:"event", Typ:[]string{"RPC_T_SERVICE_CALLBACK_TYPE"}, Nullable:false}
	msg = append(msg, byte(event>>0))
	msg = append(msg, byte(event>>8))
	msg = append(msg, byte(event>>16))
	msg = append(msg, byte(event>>24))
	// property : 2 byte main.GoArgument{In:true, Out:false, Name:"property", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(property>>0))
	msg = append(msg, byte(property>>8))
	// write_cb_data : 0 byte main.GoArgument{In:true, Out:false, Name:"write_cb_data", Typ:[]string{"[]byte"}, Nullable:true}
	msg = append(msg, write_cb_data...)
	// app_cb_data : 0 byte main.GoArgument{In:true, Out:false, Name:"app_cb_data", Typ:[]string{"[]byte"}, Nullable:true}
	msg = append(msg, app_cb_data...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// read_cb_data : 0 byte main.GoArgument{In:false, Out:true, Name:"read_cb_data", Typ:[]string{"[]byte"}, Nullable:true}
	read_cb_data_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 1
	if read_cb_data_length == 1 {
		read_cb_data_length = binary.LittleEndian.Uint32(payload[widx:])
		widx += 4
	}
	if read_cb_data_length > 0 {
		copy(read_cb_data, payload[widx:widx+int(read_cb_data_length)])
		widx += int(read_cb_data_length)
	}

	var result RPC_T_APP_RESULT
	result = RPC_T_APP_RESULT(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_connect(ssid string, password string, security_type uint32, key_id int32, semaphore uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_connect()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x01, uint32(seq))

	// ssid : 0 byte main.GoArgument{In:true, Out:false, Name:"ssid", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(ssid)), byte(len(ssid)>>8), byte(len(ssid)>>16), byte(len(ssid)>>24))
	msg = append(msg, []byte(ssid)...)
	// password : 0 byte main.GoArgument{In:true, Out:false, Name:"password", Typ:[]string{"string"}, Nullable:true}
	if len(password) == 0 {
		msg = append(msg, 1)
	} else {
		msg = append(msg, 0)
		msg = append(msg, byte(len(password)), byte(len(password)>>8), byte(len(password)>>16), byte(len(password)>>24))
		msg = append(msg, []byte(password)...)
	}
	// security_type : 4 byte main.GoArgument{In:true, Out:false, Name:"security_type", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(security_type>>0))
	msg = append(msg, byte(security_type>>8))
	msg = append(msg, byte(security_type>>16))
	msg = append(msg, byte(security_type>>24))
	// key_id : 4 byte main.GoArgument{In:true, Out:false, Name:"key_id", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(key_id>>0))
	msg = append(msg, byte(key_id>>8))
	msg = append(msg, byte(key_id>>16))
	msg = append(msg, byte(key_id>>24))
	// semaphore : 4 byte main.GoArgument{In:true, Out:false, Name:"semaphore", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(semaphore>>0))
	msg = append(msg, byte(semaphore>>8))
	msg = append(msg, byte(semaphore>>16))
	msg = append(msg, byte(semaphore>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_connect_bssid(bssid []byte, ssid string, password string, security_type uint32, key_id int32, semaphore uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_connect_bssid()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x02, uint32(seq))

	// bssid : 0 byte main.GoArgument{In:true, Out:false, Name:"bssid", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, bssid...)
	// ssid : 0 byte main.GoArgument{In:true, Out:false, Name:"ssid", Typ:[]string{"string"}, Nullable:true}
	if len(ssid) == 0 {
		msg = append(msg, 1)
	} else {
		msg = append(msg, 0)
		msg = append(msg, byte(len(ssid)), byte(len(ssid)>>8), byte(len(ssid)>>16), byte(len(ssid)>>24))
		msg = append(msg, []byte(ssid)...)
	}
	// password : 0 byte main.GoArgument{In:true, Out:false, Name:"password", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(password)), byte(len(password)>>8), byte(len(password)>>16), byte(len(password)>>24))
	msg = append(msg, []byte(password)...)
	// security_type : 4 byte main.GoArgument{In:true, Out:false, Name:"security_type", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(security_type>>0))
	msg = append(msg, byte(security_type>>8))
	msg = append(msg, byte(security_type>>16))
	msg = append(msg, byte(security_type>>24))
	// key_id : 4 byte main.GoArgument{In:true, Out:false, Name:"key_id", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(key_id>>0))
	msg = append(msg, byte(key_id>>8))
	msg = append(msg, byte(key_id>>16))
	msg = append(msg, byte(key_id>>24))
	// semaphore : 4 byte main.GoArgument{In:true, Out:false, Name:"semaphore", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(semaphore>>0))
	msg = append(msg, byte(semaphore>>8))
	msg = append(msg, byte(semaphore>>16))
	msg = append(msg, byte(semaphore>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_disconnect() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_disconnect()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x03, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_is_connected_to_ap() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_is_connected_to_ap()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x04, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_is_up(itf uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_is_up()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x05, uint32(seq))

	// itf : 4 byte main.GoArgument{In:true, Out:false, Name:"itf", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(itf>>0))
	msg = append(msg, byte(itf>>8))
	msg = append(msg, byte(itf>>16))
	msg = append(msg, byte(itf>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_is_ready_to_transceive(itf uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_is_ready_to_transceive()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x06, uint32(seq))

	// itf : 4 byte main.GoArgument{In:true, Out:false, Name:"itf", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(itf>>0))
	msg = append(msg, byte(itf>>8))
	msg = append(msg, byte(itf>>16))
	msg = append(msg, byte(itf>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_set_mac_address(mac []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_mac_address()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x07, uint32(seq))

	// mac : 0 byte main.GoArgument{In:true, Out:false, Name:"mac", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, mac...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_get_mac_address(mac uint8) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_mac_address()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x08, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// mac : 1 byte main.GoArgument{In:false, Out:true, Name:"mac", Typ:[]string{"uint8"}, Nullable:false}
// not impl

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_enable_powersave() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_enable_powersave()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x09, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_resume_powersave() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_resume_powersave()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x0A, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_disable_powersave() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_disable_powersave()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x0B, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_btcoex_set_bt_on() error {
	if debug {
		fmt.Printf("rpc_wifi_btcoex_set_bt_on()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x0C, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received

	seq++
	return err
}

func rpc_wifi_btcoex_set_bt_off() error {
	if debug {
		fmt.Printf("rpc_wifi_btcoex_set_bt_off()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x0D, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received

	seq++
	return err
}

func rpc_wifi_get_associated_client_list(client_list_buffer []byte, buffer_length uint16) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_associated_client_list()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x0E, uint32(seq))

	// buffer_length : 2 byte main.GoArgument{In:true, Out:false, Name:"buffer_length", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(buffer_length>>0))
	msg = append(msg, byte(buffer_length>>8))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// client_list_buffer : 0 byte main.GoArgument{In:false, Out:true, Name:"client_list_buffer", Typ:[]string{"[]byte"}, Nullable:false}
	client_list_buffer_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if client_list_buffer_length > 0 {
		copy(client_list_buffer, payload[widx:widx+int(client_list_buffer_length)])
		widx += int(client_list_buffer_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_get_ap_bssid(bssid uint8) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_ap_bssid()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x0F, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// bssid : 1 byte main.GoArgument{In:false, Out:true, Name:"bssid", Typ:[]string{"uint8"}, Nullable:false}
// not impl

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_get_ap_info(ap_info []byte, security uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_ap_info()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x10, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// ap_info : 0 byte main.GoArgument{In:false, Out:true, Name:"ap_info", Typ:[]string{"[]byte"}, Nullable:false}
	ap_info_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if ap_info_length > 0 {
		copy(ap_info, payload[widx:widx+int(ap_info_length)])
		widx += int(ap_info_length)
	}
	// security : 4 byte main.GoArgument{In:false, Out:true, Name:"security", Typ:[]string{"uint32"}, Nullable:false}
// not impl

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_set_country(country_code uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_country()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x11, uint32(seq))

	// country_code : 4 byte main.GoArgument{In:true, Out:false, Name:"country_code", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(country_code>>0))
	msg = append(msg, byte(country_code>>8))
	msg = append(msg, byte(country_code>>16))
	msg = append(msg, byte(country_code>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_get_sta_max_data_rate(inidata_rate uint8) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_sta_max_data_rate()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x12, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// inidata_rate : 1 byte main.GoArgument{In:false, Out:true, Name:"inidata_rate", Typ:[]string{"uint8"}, Nullable:false}
// not impl

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_get_rssi(pRSSI int32) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_rssi()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x13, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// pRSSI : 4 byte main.GoArgument{In:false, Out:true, Name:"pRSSI", Typ:[]string{"int32"}, Nullable:false}
// not impl

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_set_channel(channel int32) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_channel()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x14, uint32(seq))

	// channel : 4 byte main.GoArgument{In:true, Out:false, Name:"channel", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(channel>>0))
	msg = append(msg, byte(channel>>8))
	msg = append(msg, byte(channel>>16))
	msg = append(msg, byte(channel>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_get_channel(channel int32) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_channel()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x15, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// channel : 4 byte main.GoArgument{In:false, Out:true, Name:"channel", Typ:[]string{"int32"}, Nullable:false}
// not impl

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_change_channel_plan(channel_plan uint8) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_change_channel_plan()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x16, uint32(seq))

	// channel_plan : 1 byte main.GoArgument{In:true, Out:false, Name:"channel_plan", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(channel_plan>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_register_multicast_address(mac uint8) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_register_multicast_address()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x17, uint32(seq))

	// mac : 1 byte main.GoArgument{In:true, Out:false, Name:"mac", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(mac>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_unregister_multicast_address(mac uint8) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_unregister_multicast_address()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x18, uint32(seq))

	// mac : 1 byte main.GoArgument{In:true, Out:false, Name:"mac", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(mac>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_rf_on() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_rf_on()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x19, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_rf_off() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_rf_off()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x1A, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_on(mode uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_on()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x1B, uint32(seq))

	// mode : 4 byte main.GoArgument{In:true, Out:false, Name:"mode", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(mode>>0))
	msg = append(msg, byte(mode>>8))
	msg = append(msg, byte(mode>>16))
	msg = append(msg, byte(mode>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_off() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_off()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x1C, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_set_mode(mode uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_mode()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x1D, uint32(seq))

	// mode : 4 byte main.GoArgument{In:true, Out:false, Name:"mode", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(mode>>0))
	msg = append(msg, byte(mode>>8))
	msg = append(msg, byte(mode>>16))
	msg = append(msg, byte(mode>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_off_fastly() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_off_fastly()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x1E, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_set_power_mode(ips_mode uint8, lps_mode uint8) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_power_mode()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x1F, uint32(seq))

	// ips_mode : 1 byte main.GoArgument{In:true, Out:false, Name:"ips_mode", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(ips_mode>>0))
	// lps_mode : 1 byte main.GoArgument{In:true, Out:false, Name:"lps_mode", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(lps_mode>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_set_tdma_param(slot_period uint8, rfon_period_len_1 uint8, rfon_period_len_2 uint8, rfon_period_len_3 uint8) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_tdma_param()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x20, uint32(seq))

	// slot_period : 1 byte main.GoArgument{In:true, Out:false, Name:"slot_period", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(slot_period>>0))
	// rfon_period_len_1 : 1 byte main.GoArgument{In:true, Out:false, Name:"rfon_period_len_1", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(rfon_period_len_1>>0))
	// rfon_period_len_2 : 1 byte main.GoArgument{In:true, Out:false, Name:"rfon_period_len_2", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(rfon_period_len_2>>0))
	// rfon_period_len_3 : 1 byte main.GoArgument{In:true, Out:false, Name:"rfon_period_len_3", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(rfon_period_len_3>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_set_lps_dtim(dtim uint8) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_lps_dtim()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x21, uint32(seq))

	// dtim : 1 byte main.GoArgument{In:true, Out:false, Name:"dtim", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(dtim>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_get_lps_dtim(dtim uint8) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_lps_dtim()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x22, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// dtim : 1 byte main.GoArgument{In:false, Out:true, Name:"dtim", Typ:[]string{"uint8"}, Nullable:false}
// not impl

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_set_lps_thresh(mode uint8) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_lps_thresh()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x23, uint32(seq))

	// mode : 1 byte main.GoArgument{In:true, Out:false, Name:"mode", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(mode>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_set_lps_level(lps_level uint8) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_lps_level()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x24, uint32(seq))

	// lps_level : 1 byte main.GoArgument{In:true, Out:false, Name:"lps_level", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(lps_level>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_set_mfp_support(value uint8) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_mfp_support()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x25, uint32(seq))

	// value : 1 byte main.GoArgument{In:true, Out:false, Name:"value", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(value>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_start_ap(ssid string, password string, security_type uint32, channel int32) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_start_ap()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x26, uint32(seq))

	// ssid : 0 byte main.GoArgument{In:true, Out:false, Name:"ssid", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(ssid)), byte(len(ssid)>>8), byte(len(ssid)>>16), byte(len(ssid)>>24))
	msg = append(msg, []byte(ssid)...)
	// password : 0 byte main.GoArgument{In:true, Out:false, Name:"password", Typ:[]string{"string"}, Nullable:true}
	if len(password) == 0 {
		msg = append(msg, 1)
	} else {
		msg = append(msg, 0)
		msg = append(msg, byte(len(password)), byte(len(password)>>8), byte(len(password)>>16), byte(len(password)>>24))
		msg = append(msg, []byte(password)...)
	}
	// security_type : 4 byte main.GoArgument{In:true, Out:false, Name:"security_type", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(security_type>>0))
	msg = append(msg, byte(security_type>>8))
	msg = append(msg, byte(security_type>>16))
	msg = append(msg, byte(security_type>>24))
	// channel : 4 byte main.GoArgument{In:true, Out:false, Name:"channel", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(channel>>0))
	msg = append(msg, byte(channel>>8))
	msg = append(msg, byte(channel>>16))
	msg = append(msg, byte(channel>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_start_ap_with_hidden_ssid(ssid string, password string, security_type uint32, channel int32) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_start_ap_with_hidden_ssid()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x27, uint32(seq))

	// ssid : 0 byte main.GoArgument{In:true, Out:false, Name:"ssid", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(ssid)), byte(len(ssid)>>8), byte(len(ssid)>>16), byte(len(ssid)>>24))
	msg = append(msg, []byte(ssid)...)
	// password : 0 byte main.GoArgument{In:true, Out:false, Name:"password", Typ:[]string{"string"}, Nullable:true}
	if len(password) == 0 {
		msg = append(msg, 1)
	} else {
		msg = append(msg, 0)
		msg = append(msg, byte(len(password)), byte(len(password)>>8), byte(len(password)>>16), byte(len(password)>>24))
		msg = append(msg, []byte(password)...)
	}
	// security_type : 4 byte main.GoArgument{In:true, Out:false, Name:"security_type", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(security_type>>0))
	msg = append(msg, byte(security_type>>8))
	msg = append(msg, byte(security_type>>16))
	msg = append(msg, byte(security_type>>24))
	// channel : 4 byte main.GoArgument{In:true, Out:false, Name:"channel", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(channel>>0))
	msg = append(msg, byte(channel>>8))
	msg = append(msg, byte(channel>>16))
	msg = append(msg, byte(channel>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_set_pscan_chan(channel_list []byte, pscan_config uint8) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_pscan_chan()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x28, uint32(seq))

	// channel_list : 0 byte main.GoArgument{In:true, Out:false, Name:"channel_list", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, channel_list...)
	// pscan_config : 1 byte main.GoArgument{In:true, Out:false, Name:"pscan_config", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(pscan_config>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_get_setting(ifname string, pSetting []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_setting()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x29, uint32(seq))

	// ifname : 0 byte main.GoArgument{In:true, Out:false, Name:"ifname", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(ifname)), byte(len(ifname)>>8), byte(len(ifname)>>16), byte(len(ifname)>>24))
	msg = append(msg, []byte(ifname)...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// pSetting : 0 byte main.GoArgument{In:false, Out:true, Name:"pSetting", Typ:[]string{"[]byte"}, Nullable:false}
	pSetting_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if pSetting_length > 0 {
		copy(pSetting, payload[widx:widx+int(pSetting_length)])
		widx += int(pSetting_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_set_network_mode(mode uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_network_mode()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x2A, uint32(seq))

	// mode : 4 byte main.GoArgument{In:true, Out:false, Name:"mode", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(mode>>0))
	msg = append(msg, byte(mode>>8))
	msg = append(msg, byte(mode>>16))
	msg = append(msg, byte(mode>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_get_network_mode(pmode uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_network_mode()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x2B, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// pmode : 4 byte main.GoArgument{In:false, Out:true, Name:"pmode", Typ:[]string{"uint32"}, Nullable:false}
// not impl

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_set_wps_phase(is_trigger_wps uint8) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_wps_phase()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x2C, uint32(seq))

	// is_trigger_wps : 1 byte main.GoArgument{In:true, Out:false, Name:"is_trigger_wps", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(is_trigger_wps>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_restart_ap(ssid []byte, password []byte, security_type uint32, channel int32) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_restart_ap()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x2D, uint32(seq))

	// ssid : 0 byte main.GoArgument{In:true, Out:false, Name:"ssid", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, ssid...)
	// password : 0 byte main.GoArgument{In:true, Out:false, Name:"password", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, password...)
	// security_type : 4 byte main.GoArgument{In:true, Out:false, Name:"security_type", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(security_type>>0))
	msg = append(msg, byte(security_type>>8))
	msg = append(msg, byte(security_type>>16))
	msg = append(msg, byte(security_type>>24))
	// channel : 4 byte main.GoArgument{In:true, Out:false, Name:"channel", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(channel>>0))
	msg = append(msg, byte(channel>>8))
	msg = append(msg, byte(channel>>16))
	msg = append(msg, byte(channel>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_config_autoreconnect(mode uint8, retry_times uint8, timeout uint16) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_config_autoreconnect()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x2E, uint32(seq))

	// mode : 1 byte main.GoArgument{In:true, Out:false, Name:"mode", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(mode>>0))
	// retry_times : 1 byte main.GoArgument{In:true, Out:false, Name:"retry_times", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(retry_times>>0))
	// timeout : 2 byte main.GoArgument{In:true, Out:false, Name:"timeout", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(timeout>>0))
	msg = append(msg, byte(timeout>>8))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_set_autoreconnect(mode uint8) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_autoreconnect()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x2F, uint32(seq))

	// mode : 1 byte main.GoArgument{In:true, Out:false, Name:"mode", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(mode>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_get_autoreconnect(mode uint8) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_autoreconnect()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x30, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// mode : 1 byte main.GoArgument{In:false, Out:true, Name:"mode", Typ:[]string{"uint8"}, Nullable:false}
// not impl

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_get_last_error() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_last_error()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x31, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_add_custom_ie(cus_ie []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_add_custom_ie()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x32, uint32(seq))

	// cus_ie : 0 byte main.GoArgument{In:true, Out:false, Name:"cus_ie", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, cus_ie...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_update_custom_ie(cus_ie []byte, ie_index int32) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_update_custom_ie()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x33, uint32(seq))

	// cus_ie : 0 byte main.GoArgument{In:true, Out:false, Name:"cus_ie", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, cus_ie...)
	// ie_index : 4 byte main.GoArgument{In:true, Out:false, Name:"ie_index", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(ie_index>>0))
	msg = append(msg, byte(ie_index>>8))
	msg = append(msg, byte(ie_index>>16))
	msg = append(msg, byte(ie_index>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_del_custom_ie() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_del_custom_ie()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x34, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_set_indicate_mgnt(enable int32) error {
	if debug {
		fmt.Printf("rpc_wifi_set_indicate_mgnt()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x35, uint32(seq))

	// enable : 4 byte main.GoArgument{In:true, Out:false, Name:"enable", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(enable>>0))
	msg = append(msg, byte(enable>>8))
	msg = append(msg, byte(enable>>16))
	msg = append(msg, byte(enable>>24))

	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received


	seq++
	return err
}

func rpc_wifi_get_drv_ability(ability uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_drv_ability()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x36, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// ability : 4 byte main.GoArgument{In:false, Out:true, Name:"ability", Typ:[]string{"uint32"}, Nullable:false}
// not impl

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_set_channel_plan(channel_plan uint8) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_channel_plan()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x37, uint32(seq))

	// channel_plan : 1 byte main.GoArgument{In:true, Out:false, Name:"channel_plan", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(channel_plan>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_get_channel_plan(channel_plan uint8) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_channel_plan()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x38, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// channel_plan : 1 byte main.GoArgument{In:false, Out:true, Name:"channel_plan", Typ:[]string{"uint8"}, Nullable:false}
// not impl

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_enable_forwarding() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_enable_forwarding()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x39, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_disable_forwarding() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_disable_forwarding()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x3A, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_set_ch_deauth(enable uint8) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_ch_deauth()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x3B, uint32(seq))

	// enable : 1 byte main.GoArgument{In:true, Out:false, Name:"enable", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(enable>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_get_band_type() (uint8, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_band_type()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x3C, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result uint8
	result = uint8(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_set_tx_pause_data(NewState uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_tx_pause_data()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x3D, uint32(seq))

	// NewState : 4 byte main.GoArgument{In:true, Out:false, Name:"NewState", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(NewState>>0))
	msg = append(msg, byte(NewState>>8))
	msg = append(msg, byte(NewState>>16))
	msg = append(msg, byte(NewState>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_get_reconnect_data(wifi_info []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_reconnect_data()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x3E, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// wifi_info : 0 byte main.GoArgument{In:false, Out:true, Name:"wifi_info", Typ:[]string{"[]byte"}, Nullable:false}
	wifi_info_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if wifi_info_length > 0 {
		copy(wifi_info, payload[widx:widx+int(wifi_info_length)])
		widx += int(wifi_info_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_clear_reconnect_data() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_clear_reconnect_data()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x3F, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_scan_start() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_scan_start()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x40, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_is_scaning() (bool, error) {
	if debug {
		fmt.Printf("rpc_wifi_is_scaning()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x41, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	widx := 8
	var result bool
	result = binary.LittleEndian.Uint32(payload[widx:]) == 1

	seq++
	return result, err
}

func rpc_wifi_scan_get_ap_records(number uint16, _scanResult []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_scan_get_ap_records()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x42, uint32(seq))

	// number : 2 byte main.GoArgument{In:true, Out:false, Name:"number", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(number>>0))
	msg = append(msg, byte(number>>8))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// _scanResult : 0 byte main.GoArgument{In:false, Out:true, Name:"_scanResult", Typ:[]string{"[]byte"}, Nullable:false}
	_scanResult_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if _scanResult_length > 0 {
		copy(_scanResult, payload[widx:widx+int(_scanResult_length)])
		widx += int(_scanResult_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_scan_get_ap_num() (uint16, error) {
	if debug {
		fmt.Printf("rpc_wifi_scan_get_ap_num()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x43, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result uint16
	result = uint16(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcpip_adapter_init() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_init()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x01, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcpip_adapter_sta_start(mac []byte, ip_info []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_sta_start()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x02, uint32(seq))

	// mac : 0 byte main.GoArgument{In:true, Out:false, Name:"mac", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, mac...)
	// ip_info : 0 byte main.GoArgument{In:true, Out:false, Name:"ip_info", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, ip_info...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcpip_adapter_ap_start(mac []byte, ip_info []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_ap_start()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x03, uint32(seq))

	// mac : 0 byte main.GoArgument{In:true, Out:false, Name:"mac", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, mac...)
	// ip_info : 0 byte main.GoArgument{In:true, Out:false, Name:"ip_info", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, ip_info...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcpip_adapter_stop(tcpip_if uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_stop()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x04, uint32(seq))

	// tcpip_if : 4 byte main.GoArgument{In:true, Out:false, Name:"tcpip_if", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(tcpip_if>>0))
	msg = append(msg, byte(tcpip_if>>8))
	msg = append(msg, byte(tcpip_if>>16))
	msg = append(msg, byte(tcpip_if>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcpip_adapter_up(tcpip_if uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_up()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x05, uint32(seq))

	// tcpip_if : 4 byte main.GoArgument{In:true, Out:false, Name:"tcpip_if", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(tcpip_if>>0))
	msg = append(msg, byte(tcpip_if>>8))
	msg = append(msg, byte(tcpip_if>>16))
	msg = append(msg, byte(tcpip_if>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcpip_adapter_down(tcpip_if uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_down()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x06, uint32(seq))

	// tcpip_if : 4 byte main.GoArgument{In:true, Out:false, Name:"tcpip_if", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(tcpip_if>>0))
	msg = append(msg, byte(tcpip_if>>8))
	msg = append(msg, byte(tcpip_if>>16))
	msg = append(msg, byte(tcpip_if>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcpip_adapter_get_ip_info(tcpip_if uint32, ip_info []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_get_ip_info()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x07, uint32(seq))

	// tcpip_if : 4 byte main.GoArgument{In:true, Out:false, Name:"tcpip_if", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(tcpip_if>>0))
	msg = append(msg, byte(tcpip_if>>8))
	msg = append(msg, byte(tcpip_if>>16))
	msg = append(msg, byte(tcpip_if>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// ip_info : 0 byte main.GoArgument{In:false, Out:true, Name:"ip_info", Typ:[]string{"[]byte"}, Nullable:false}
	ip_info_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if ip_info_length > 0 {
		copy(ip_info, payload[widx:widx+int(ip_info_length)])
		widx += int(ip_info_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcpip_adapter_set_ip_info(tcpip_if uint32, ip_info []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_set_ip_info()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x08, uint32(seq))

	// tcpip_if : 4 byte main.GoArgument{In:true, Out:false, Name:"tcpip_if", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(tcpip_if>>0))
	msg = append(msg, byte(tcpip_if>>8))
	msg = append(msg, byte(tcpip_if>>16))
	msg = append(msg, byte(tcpip_if>>24))
	// ip_info : 0 byte main.GoArgument{In:true, Out:false, Name:"ip_info", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, ip_info...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcpip_adapter_set_dns_info(tcpip_if uint32, dns_type uint32, dns []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_set_dns_info()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x09, uint32(seq))

	// tcpip_if : 4 byte main.GoArgument{In:true, Out:false, Name:"tcpip_if", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(tcpip_if>>0))
	msg = append(msg, byte(tcpip_if>>8))
	msg = append(msg, byte(tcpip_if>>16))
	msg = append(msg, byte(tcpip_if>>24))
	// dns_type : 4 byte main.GoArgument{In:true, Out:false, Name:"dns_type", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(dns_type>>0))
	msg = append(msg, byte(dns_type>>8))
	msg = append(msg, byte(dns_type>>16))
	msg = append(msg, byte(dns_type>>24))
	// dns : 0 byte main.GoArgument{In:true, Out:false, Name:"dns", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, dns...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcpip_adapter_get_dns_info(tcpip_if uint32, dns_type uint32, dns []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_get_dns_info()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x0A, uint32(seq))

	// tcpip_if : 4 byte main.GoArgument{In:true, Out:false, Name:"tcpip_if", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(tcpip_if>>0))
	msg = append(msg, byte(tcpip_if>>8))
	msg = append(msg, byte(tcpip_if>>16))
	msg = append(msg, byte(tcpip_if>>24))
	// dns_type : 4 byte main.GoArgument{In:true, Out:false, Name:"dns_type", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(dns_type>>0))
	msg = append(msg, byte(dns_type>>8))
	msg = append(msg, byte(dns_type>>16))
	msg = append(msg, byte(dns_type>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// dns : 0 byte main.GoArgument{In:false, Out:true, Name:"dns", Typ:[]string{"[]byte"}, Nullable:false}
	dns_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if dns_length > 0 {
		copy(dns, payload[widx:widx+int(dns_length)])
		widx += int(dns_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcpip_adapter_dhcps_start(tcpip_if uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_dhcps_start()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x0B, uint32(seq))

	// tcpip_if : 4 byte main.GoArgument{In:true, Out:false, Name:"tcpip_if", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(tcpip_if>>0))
	msg = append(msg, byte(tcpip_if>>8))
	msg = append(msg, byte(tcpip_if>>16))
	msg = append(msg, byte(tcpip_if>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcpip_adapter_dhcps_stop(tcpip_if uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_dhcps_stop()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x0C, uint32(seq))

	// tcpip_if : 4 byte main.GoArgument{In:true, Out:false, Name:"tcpip_if", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(tcpip_if>>0))
	msg = append(msg, byte(tcpip_if>>8))
	msg = append(msg, byte(tcpip_if>>16))
	msg = append(msg, byte(tcpip_if>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcpip_adapter_dhcpc_start(tcpip_if uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_dhcpc_start()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x0D, uint32(seq))

	// tcpip_if : 4 byte main.GoArgument{In:true, Out:false, Name:"tcpip_if", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(tcpip_if>>0))
	msg = append(msg, byte(tcpip_if>>8))
	msg = append(msg, byte(tcpip_if>>16))
	msg = append(msg, byte(tcpip_if>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcpip_adapter_dhcpc_stop(tcpip_if uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_dhcpc_stop()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x0E, uint32(seq))

	// tcpip_if : 4 byte main.GoArgument{In:true, Out:false, Name:"tcpip_if", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(tcpip_if>>0))
	msg = append(msg, byte(tcpip_if>>8))
	msg = append(msg, byte(tcpip_if>>16))
	msg = append(msg, byte(tcpip_if>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcpip_adapter_set_hostname(tcpip_if uint32, hostname string) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_set_hostname()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x0F, uint32(seq))

	// tcpip_if : 4 byte main.GoArgument{In:true, Out:false, Name:"tcpip_if", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(tcpip_if>>0))
	msg = append(msg, byte(tcpip_if>>8))
	msg = append(msg, byte(tcpip_if>>16))
	msg = append(msg, byte(tcpip_if>>24))
	// hostname : 0 byte main.GoArgument{In:true, Out:false, Name:"hostname", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(hostname)), byte(len(hostname)>>8), byte(len(hostname)>>16), byte(len(hostname)>>24))
	msg = append(msg, []byte(hostname)...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcpip_adapter_get_hostname(tcpip_if uint32, hostname string) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_get_hostname()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x10, uint32(seq))

	// tcpip_if : 4 byte main.GoArgument{In:true, Out:false, Name:"tcpip_if", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(tcpip_if>>0))
	msg = append(msg, byte(tcpip_if>>8))
	msg = append(msg, byte(tcpip_if>>16))
	msg = append(msg, byte(tcpip_if>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// hostname : 0 byte main.GoArgument{In:false, Out:true, Name:"hostname", Typ:[]string{"string"}, Nullable:false}
	hostname_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if hostname_length > 0 {
		hostname = string(payload[widx:widx+int(hostname_length)])
		widx += int(hostname_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcpip_adapter_get_mac(tcpip_if uint32, mac []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_get_mac()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x11, uint32(seq))

	// tcpip_if : 4 byte main.GoArgument{In:true, Out:false, Name:"tcpip_if", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(tcpip_if>>0))
	msg = append(msg, byte(tcpip_if>>8))
	msg = append(msg, byte(tcpip_if>>16))
	msg = append(msg, byte(tcpip_if>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// mac : 0 byte main.GoArgument{In:false, Out:true, Name:"mac", Typ:[]string{"[]byte"}, Nullable:false}
	mac_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if mac_length > 0 {
		copy(mac, payload[widx:widx+int(mac_length)])
		widx += int(mac_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcpip_adapter_set_mac(tcpip_if uint32, mac []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_set_mac()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x12, uint32(seq))

	// tcpip_if : 4 byte main.GoArgument{In:true, Out:false, Name:"tcpip_if", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(tcpip_if>>0))
	msg = append(msg, byte(tcpip_if>>8))
	msg = append(msg, byte(tcpip_if>>16))
	msg = append(msg, byte(tcpip_if>>24))
	// mac : 0 byte main.GoArgument{In:true, Out:false, Name:"mac", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, mac...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcpip_api_call(fn []byte, call []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_api_call()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x13, uint32(seq))

	// fn : 0 byte main.GoArgument{In:true, Out:false, Name:"fn", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, fn...)
	// call : 0 byte main.GoArgument{In:true, Out:false, Name:"call", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, call...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcp_connect(pcb_in []byte, pcb_out []byte, ipaddr []byte, port uint16, connected []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_connect()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x14, uint32(seq))

	// pcb_in : 0 byte main.GoArgument{In:true, Out:false, Name:"pcb_in", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, pcb_in...)
	// ipaddr : 0 byte main.GoArgument{In:true, Out:false, Name:"ipaddr", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, ipaddr...)
	// port : 2 byte main.GoArgument{In:true, Out:false, Name:"port", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(port>>0))
	msg = append(msg, byte(port>>8))
	// connected : 0 byte main.GoArgument{In:true, Out:false, Name:"connected", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, connected...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// pcb_out : 0 byte main.GoArgument{In:false, Out:true, Name:"pcb_out", Typ:[]string{"[]byte"}, Nullable:false}
	pcb_out_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if pcb_out_length > 0 {
		copy(pcb_out, payload[widx:widx+int(pcb_out_length)])
		widx += int(pcb_out_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcp_recved(pcb_in []byte, pcb_out []byte, len uint16) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_recved()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x15, uint32(seq))

	// pcb_in : 0 byte main.GoArgument{In:true, Out:false, Name:"pcb_in", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, pcb_in...)
	// len : 2 byte main.GoArgument{In:true, Out:false, Name:"len", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(len>>0))
	msg = append(msg, byte(len>>8))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// pcb_out : 0 byte main.GoArgument{In:false, Out:true, Name:"pcb_out", Typ:[]string{"[]byte"}, Nullable:false}
	pcb_out_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if pcb_out_length > 0 {
		copy(pcb_out, payload[widx:widx+int(pcb_out_length)])
		widx += int(pcb_out_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcp_abort(pcb_in []byte, pcb_out []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_abort()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x16, uint32(seq))

	// pcb_in : 0 byte main.GoArgument{In:true, Out:false, Name:"pcb_in", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, pcb_in...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// pcb_out : 0 byte main.GoArgument{In:false, Out:true, Name:"pcb_out", Typ:[]string{"[]byte"}, Nullable:false}
	pcb_out_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if pcb_out_length > 0 {
		copy(pcb_out, payload[widx:widx+int(pcb_out_length)])
		widx += int(pcb_out_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcp_write(pcb_in []byte, pcb_out []byte, data []byte, apiflags uint8) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_write()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x17, uint32(seq))

	// pcb_in : 0 byte main.GoArgument{In:true, Out:false, Name:"pcb_in", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, pcb_in...)
	// data : 0 byte main.GoArgument{In:true, Out:false, Name:"data", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, data...)
	// apiflags : 1 byte main.GoArgument{In:true, Out:false, Name:"apiflags", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(apiflags>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// pcb_out : 0 byte main.GoArgument{In:false, Out:true, Name:"pcb_out", Typ:[]string{"[]byte"}, Nullable:false}
	pcb_out_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if pcb_out_length > 0 {
		copy(pcb_out, payload[widx:widx+int(pcb_out_length)])
		widx += int(pcb_out_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcp_output(pcb_in []byte, pcb_out []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_output()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x18, uint32(seq))

	// pcb_in : 0 byte main.GoArgument{In:true, Out:false, Name:"pcb_in", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, pcb_in...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// pcb_out : 0 byte main.GoArgument{In:false, Out:true, Name:"pcb_out", Typ:[]string{"[]byte"}, Nullable:false}
	pcb_out_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if pcb_out_length > 0 {
		copy(pcb_out, payload[widx:widx+int(pcb_out_length)])
		widx += int(pcb_out_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcp_close(pcb_in []byte, pcb_out []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_close()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x19, uint32(seq))

	// pcb_in : 0 byte main.GoArgument{In:true, Out:false, Name:"pcb_in", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, pcb_in...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// pcb_out : 0 byte main.GoArgument{In:false, Out:true, Name:"pcb_out", Typ:[]string{"[]byte"}, Nullable:false}
	pcb_out_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if pcb_out_length > 0 {
		copy(pcb_out, payload[widx:widx+int(pcb_out_length)])
		widx += int(pcb_out_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcp_bind(pcb_in []byte, pcb_out []byte, ipaddr []byte, port uint16) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_bind()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x1A, uint32(seq))

	// pcb_in : 0 byte main.GoArgument{In:true, Out:false, Name:"pcb_in", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, pcb_in...)
	// ipaddr : 0 byte main.GoArgument{In:true, Out:false, Name:"ipaddr", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, ipaddr...)
	// port : 2 byte main.GoArgument{In:true, Out:false, Name:"port", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(port>>0))
	msg = append(msg, byte(port>>8))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// pcb_out : 0 byte main.GoArgument{In:false, Out:true, Name:"pcb_out", Typ:[]string{"[]byte"}, Nullable:false}
	pcb_out_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if pcb_out_length > 0 {
		copy(pcb_out, payload[widx:widx+int(pcb_out_length)])
		widx += int(pcb_out_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcp_new_ip_type(ip_type uint8, pcb_out []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_new_ip_type()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x1B, uint32(seq))

	// ip_type : 1 byte main.GoArgument{In:true, Out:false, Name:"ip_type", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(ip_type>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// pcb_out : 0 byte main.GoArgument{In:false, Out:true, Name:"pcb_out", Typ:[]string{"[]byte"}, Nullable:false}
	pcb_out_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if pcb_out_length > 0 {
		copy(pcb_out, payload[widx:widx+int(pcb_out_length)])
		widx += int(pcb_out_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcp_arg(pcb_in []byte, pcb_out []byte, func_arg []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_arg()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x1C, uint32(seq))

	// pcb_in : 0 byte main.GoArgument{In:true, Out:false, Name:"pcb_in", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, pcb_in...)
	// func_arg : 0 byte main.GoArgument{In:true, Out:false, Name:"func_arg", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, func_arg...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// pcb_out : 0 byte main.GoArgument{In:false, Out:true, Name:"pcb_out", Typ:[]string{"[]byte"}, Nullable:false}
	pcb_out_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if pcb_out_length > 0 {
		copy(pcb_out, payload[widx:widx+int(pcb_out_length)])
		widx += int(pcb_out_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcp_err(pcb_in []byte, pcb_out []byte, func_err []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_err()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x1D, uint32(seq))

	// pcb_in : 0 byte main.GoArgument{In:true, Out:false, Name:"pcb_in", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, pcb_in...)
	// func_err : 0 byte main.GoArgument{In:true, Out:false, Name:"func_err", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, func_err...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// pcb_out : 0 byte main.GoArgument{In:false, Out:true, Name:"pcb_out", Typ:[]string{"[]byte"}, Nullable:false}
	pcb_out_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if pcb_out_length > 0 {
		copy(pcb_out, payload[widx:widx+int(pcb_out_length)])
		widx += int(pcb_out_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcp_recv(pcb_in []byte, pcb_out []byte, func_recv []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_recv()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x1E, uint32(seq))

	// pcb_in : 0 byte main.GoArgument{In:true, Out:false, Name:"pcb_in", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, pcb_in...)
	// func_recv : 0 byte main.GoArgument{In:true, Out:false, Name:"func_recv", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, func_recv...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// pcb_out : 0 byte main.GoArgument{In:false, Out:true, Name:"pcb_out", Typ:[]string{"[]byte"}, Nullable:false}
	pcb_out_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if pcb_out_length > 0 {
		copy(pcb_out, payload[widx:widx+int(pcb_out_length)])
		widx += int(pcb_out_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcp_sent(pcb_in []byte, pcb_out []byte, func_sent []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_sent()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x1F, uint32(seq))

	// pcb_in : 0 byte main.GoArgument{In:true, Out:false, Name:"pcb_in", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, pcb_in...)
	// func_sent : 0 byte main.GoArgument{In:true, Out:false, Name:"func_sent", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, func_sent...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// pcb_out : 0 byte main.GoArgument{In:false, Out:true, Name:"pcb_out", Typ:[]string{"[]byte"}, Nullable:false}
	pcb_out_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if pcb_out_length > 0 {
		copy(pcb_out, payload[widx:widx+int(pcb_out_length)])
		widx += int(pcb_out_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcp_accept(pcb_in []byte, pcb_out []byte, func_accept []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_accept()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x20, uint32(seq))

	// pcb_in : 0 byte main.GoArgument{In:true, Out:false, Name:"pcb_in", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, pcb_in...)
	// func_accept : 0 byte main.GoArgument{In:true, Out:false, Name:"func_accept", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, func_accept...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// pcb_out : 0 byte main.GoArgument{In:false, Out:true, Name:"pcb_out", Typ:[]string{"[]byte"}, Nullable:false}
	pcb_out_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if pcb_out_length > 0 {
		copy(pcb_out, payload[widx:widx+int(pcb_out_length)])
		widx += int(pcb_out_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcp_poll(pcb_in []byte, pcb_out []byte, func_poll []byte, interval uint8) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_poll()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x21, uint32(seq))

	// pcb_in : 0 byte main.GoArgument{In:true, Out:false, Name:"pcb_in", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, pcb_in...)
	// func_poll : 0 byte main.GoArgument{In:true, Out:false, Name:"func_poll", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, func_poll...)
	// interval : 1 byte main.GoArgument{In:true, Out:false, Name:"interval", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(interval>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// pcb_out : 0 byte main.GoArgument{In:false, Out:true, Name:"pcb_out", Typ:[]string{"[]byte"}, Nullable:false}
	pcb_out_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if pcb_out_length > 0 {
		copy(pcb_out, payload[widx:widx+int(pcb_out_length)])
		widx += int(pcb_out_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcp_listen_with_backlog(pcb_in []byte, pcb_out []byte, backlog uint8) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_listen_with_backlog()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x22, uint32(seq))

	// pcb_in : 0 byte main.GoArgument{In:true, Out:false, Name:"pcb_in", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, pcb_in...)
	// backlog : 1 byte main.GoArgument{In:true, Out:false, Name:"backlog", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(backlog>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// pcb_out : 0 byte main.GoArgument{In:false, Out:true, Name:"pcb_out", Typ:[]string{"[]byte"}, Nullable:false}
	pcb_out_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if pcb_out_length > 0 {
		copy(pcb_out, payload[widx:widx+int(pcb_out_length)])
		widx += int(pcb_out_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_pbuf_free(p []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_pbuf_free()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x23, uint32(seq))

	// p : 0 byte main.GoArgument{In:true, Out:false, Name:"p", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, p...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_ip4addr_ntoa(ip4_addr_in []byte) (string, error) {
	if debug {
		fmt.Printf("rpc_ip4addr_ntoa()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x24, uint32(seq))

	// ip4_addr_in : 0 byte main.GoArgument{In:true, Out:false, Name:"ip4_addr_in", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, ip4_addr_in...)

	err := performRequest(msg)
	if err != nil {
		return "", err
	}

	<-received
	widx := 8

	var result string
	result = string(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_inet_chksum(dataptr_in []byte) (uint16, error) {
	if debug {
		fmt.Printf("rpc_inet_chksum()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x25, uint32(seq))

	// dataptr_in : 0 byte main.GoArgument{In:true, Out:false, Name:"dataptr_in", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, dataptr_in...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result uint16
	result = uint16(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_lwip_accept(s int32, addr []byte, addrlen uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_accept()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x01, uint32(seq))

	// s : 4 byte main.GoArgument{In:true, Out:false, Name:"s", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(s>>0))
	msg = append(msg, byte(s>>8))
	msg = append(msg, byte(s>>16))
	msg = append(msg, byte(s>>24))
	// addr : 0 byte main.GoArgument{In:true, Out:false, Name:"addr", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, addr...)
	// addrlen : 4 byte main.GoArgument{In:true, Out:true, Name:"addrlen", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(addrlen>>0))
	msg = append(msg, byte(addrlen>>8))
	msg = append(msg, byte(addrlen>>16))
	msg = append(msg, byte(addrlen>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// addrlen : 4 byte main.GoArgument{In:true, Out:true, Name:"addrlen", Typ:[]string{"uint32"}, Nullable:false}
// not impl

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_lwip_bind(s int32, name []byte, namelen uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_bind()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x02, uint32(seq))

	// s : 4 byte main.GoArgument{In:true, Out:false, Name:"s", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(s>>0))
	msg = append(msg, byte(s>>8))
	msg = append(msg, byte(s>>16))
	msg = append(msg, byte(s>>24))
	// name : 0 byte main.GoArgument{In:true, Out:false, Name:"name", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, name...)
	// namelen : 4 byte main.GoArgument{In:true, Out:false, Name:"namelen", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(namelen>>0))
	msg = append(msg, byte(namelen>>8))
	msg = append(msg, byte(namelen>>16))
	msg = append(msg, byte(namelen>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_lwip_shutdown(s int32, how int32) (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_shutdown()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x03, uint32(seq))

	// s : 4 byte main.GoArgument{In:true, Out:false, Name:"s", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(s>>0))
	msg = append(msg, byte(s>>8))
	msg = append(msg, byte(s>>16))
	msg = append(msg, byte(s>>24))
	// how : 4 byte main.GoArgument{In:true, Out:false, Name:"how", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(how>>0))
	msg = append(msg, byte(how>>8))
	msg = append(msg, byte(how>>16))
	msg = append(msg, byte(how>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_lwip_getpeername(s int32, name []byte, namelen uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_getpeername()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x04, uint32(seq))

	// s : 4 byte main.GoArgument{In:true, Out:false, Name:"s", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(s>>0))
	msg = append(msg, byte(s>>8))
	msg = append(msg, byte(s>>16))
	msg = append(msg, byte(s>>24))
	// namelen : 4 byte main.GoArgument{In:true, Out:true, Name:"namelen", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(namelen>>0))
	msg = append(msg, byte(namelen>>8))
	msg = append(msg, byte(namelen>>16))
	msg = append(msg, byte(namelen>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// name : 0 byte main.GoArgument{In:false, Out:true, Name:"name", Typ:[]string{"[]byte"}, Nullable:false}
	name_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if name_length > 0 {
		copy(name, payload[widx:widx+int(name_length)])
		widx += int(name_length)
	}
	// namelen : 4 byte main.GoArgument{In:true, Out:true, Name:"namelen", Typ:[]string{"uint32"}, Nullable:false}
// not impl

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_lwip_getsockname(s int32, name []byte, namelen uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_getsockname()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x05, uint32(seq))

	// s : 4 byte main.GoArgument{In:true, Out:false, Name:"s", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(s>>0))
	msg = append(msg, byte(s>>8))
	msg = append(msg, byte(s>>16))
	msg = append(msg, byte(s>>24))
	// namelen : 4 byte main.GoArgument{In:true, Out:true, Name:"namelen", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(namelen>>0))
	msg = append(msg, byte(namelen>>8))
	msg = append(msg, byte(namelen>>16))
	msg = append(msg, byte(namelen>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// name : 0 byte main.GoArgument{In:false, Out:true, Name:"name", Typ:[]string{"[]byte"}, Nullable:false}
	name_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if name_length > 0 {
		copy(name, payload[widx:widx+int(name_length)])
		widx += int(name_length)
	}
	// namelen : 4 byte main.GoArgument{In:true, Out:true, Name:"namelen", Typ:[]string{"uint32"}, Nullable:false}
// not impl

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_lwip_getsockopt(s int32, level int32, optname int32, in_optval []byte, out_optval []byte, optlen uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_getsockopt()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x06, uint32(seq))

	// s : 4 byte main.GoArgument{In:true, Out:false, Name:"s", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(s>>0))
	msg = append(msg, byte(s>>8))
	msg = append(msg, byte(s>>16))
	msg = append(msg, byte(s>>24))
	// level : 4 byte main.GoArgument{In:true, Out:false, Name:"level", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(level>>0))
	msg = append(msg, byte(level>>8))
	msg = append(msg, byte(level>>16))
	msg = append(msg, byte(level>>24))
	// optname : 4 byte main.GoArgument{In:true, Out:false, Name:"optname", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(optname>>0))
	msg = append(msg, byte(optname>>8))
	msg = append(msg, byte(optname>>16))
	msg = append(msg, byte(optname>>24))
	// in_optval : 0 byte main.GoArgument{In:true, Out:false, Name:"in_optval", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, in_optval...)
	// optlen : 4 byte main.GoArgument{In:true, Out:true, Name:"optlen", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(optlen>>0))
	msg = append(msg, byte(optlen>>8))
	msg = append(msg, byte(optlen>>16))
	msg = append(msg, byte(optlen>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// out_optval : 0 byte main.GoArgument{In:false, Out:true, Name:"out_optval", Typ:[]string{"[]byte"}, Nullable:false}
	out_optval_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if out_optval_length > 0 {
		copy(out_optval, payload[widx:widx+int(out_optval_length)])
		widx += int(out_optval_length)
	}
	// optlen : 4 byte main.GoArgument{In:true, Out:true, Name:"optlen", Typ:[]string{"uint32"}, Nullable:false}
// not impl

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_lwip_setsockopt(s int32, level int32, optname int32, optval []byte, optlen uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_setsockopt()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x07, uint32(seq))

	// s : 4 byte main.GoArgument{In:true, Out:false, Name:"s", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(s>>0))
	msg = append(msg, byte(s>>8))
	msg = append(msg, byte(s>>16))
	msg = append(msg, byte(s>>24))
	// level : 4 byte main.GoArgument{In:true, Out:false, Name:"level", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(level>>0))
	msg = append(msg, byte(level>>8))
	msg = append(msg, byte(level>>16))
	msg = append(msg, byte(level>>24))
	// optname : 4 byte main.GoArgument{In:true, Out:false, Name:"optname", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(optname>>0))
	msg = append(msg, byte(optname>>8))
	msg = append(msg, byte(optname>>16))
	msg = append(msg, byte(optname>>24))
	// optval : 0 byte main.GoArgument{In:true, Out:false, Name:"optval", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, optval...)
	// optlen : 4 byte main.GoArgument{In:true, Out:false, Name:"optlen", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(optlen>>0))
	msg = append(msg, byte(optlen>>8))
	msg = append(msg, byte(optlen>>16))
	msg = append(msg, byte(optlen>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_lwip_close(s int32) (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_close()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x08, uint32(seq))

	// s : 4 byte main.GoArgument{In:true, Out:false, Name:"s", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(s>>0))
	msg = append(msg, byte(s>>8))
	msg = append(msg, byte(s>>16))
	msg = append(msg, byte(s>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_lwip_connect(s int32, name []byte, namelen uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_connect()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x09, uint32(seq))

	// s : 4 byte main.GoArgument{In:true, Out:false, Name:"s", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(s>>0))
	msg = append(msg, byte(s>>8))
	msg = append(msg, byte(s>>16))
	msg = append(msg, byte(s>>24))
	// name : 0 byte main.GoArgument{In:true, Out:false, Name:"name", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, name...)
	// namelen : 4 byte main.GoArgument{In:true, Out:false, Name:"namelen", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(namelen>>0))
	msg = append(msg, byte(namelen>>8))
	msg = append(msg, byte(namelen>>16))
	msg = append(msg, byte(namelen>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_lwip_listen(s int32, backlog int32) (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_listen()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x0A, uint32(seq))

	// s : 4 byte main.GoArgument{In:true, Out:false, Name:"s", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(s>>0))
	msg = append(msg, byte(s>>8))
	msg = append(msg, byte(s>>16))
	msg = append(msg, byte(s>>24))
	// backlog : 4 byte main.GoArgument{In:true, Out:false, Name:"backlog", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(backlog>>0))
	msg = append(msg, byte(backlog>>8))
	msg = append(msg, byte(backlog>>16))
	msg = append(msg, byte(backlog>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_lwip_available(s int32) (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_available()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x0B, uint32(seq))

	// s : 4 byte main.GoArgument{In:true, Out:false, Name:"s", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(s>>0))
	msg = append(msg, byte(s>>8))
	msg = append(msg, byte(s>>16))
	msg = append(msg, byte(s>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_lwip_recv(s int32, mem []byte, len uint32, flags int32, timeout uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_recv()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x0C, uint32(seq))

	// s : 4 byte main.GoArgument{In:true, Out:false, Name:"s", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(s>>0))
	msg = append(msg, byte(s>>8))
	msg = append(msg, byte(s>>16))
	msg = append(msg, byte(s>>24))
	// len : 4 byte main.GoArgument{In:true, Out:false, Name:"len", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(len>>0))
	msg = append(msg, byte(len>>8))
	msg = append(msg, byte(len>>16))
	msg = append(msg, byte(len>>24))
	// flags : 4 byte main.GoArgument{In:true, Out:false, Name:"flags", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(flags>>0))
	msg = append(msg, byte(flags>>8))
	msg = append(msg, byte(flags>>16))
	msg = append(msg, byte(flags>>24))
	// timeout : 4 byte main.GoArgument{In:true, Out:false, Name:"timeout", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(timeout>>0))
	msg = append(msg, byte(timeout>>8))
	msg = append(msg, byte(timeout>>16))
	msg = append(msg, byte(timeout>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// mem : 0 byte main.GoArgument{In:false, Out:true, Name:"mem", Typ:[]string{"[]byte"}, Nullable:false}
	mem_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if mem_length > 0 {
		copy(mem, payload[widx:widx+int(mem_length)])
		widx += int(mem_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_lwip_read(s int32, mem []byte, len uint32, timeout uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_read()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x0D, uint32(seq))

	// s : 4 byte main.GoArgument{In:true, Out:false, Name:"s", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(s>>0))
	msg = append(msg, byte(s>>8))
	msg = append(msg, byte(s>>16))
	msg = append(msg, byte(s>>24))
	// len : 4 byte main.GoArgument{In:true, Out:false, Name:"len", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(len>>0))
	msg = append(msg, byte(len>>8))
	msg = append(msg, byte(len>>16))
	msg = append(msg, byte(len>>24))
	// timeout : 4 byte main.GoArgument{In:true, Out:false, Name:"timeout", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(timeout>>0))
	msg = append(msg, byte(timeout>>8))
	msg = append(msg, byte(timeout>>16))
	msg = append(msg, byte(timeout>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// mem : 0 byte main.GoArgument{In:false, Out:true, Name:"mem", Typ:[]string{"[]byte"}, Nullable:false}
	mem_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if mem_length > 0 {
		copy(mem, payload[widx:widx+int(mem_length)])
		widx += int(mem_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_lwip_recvfrom(s int32, mem []byte, len uint32, flags int32, from []byte, fromlen uint32, timeout uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_recvfrom()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x0E, uint32(seq))

	// s : 4 byte main.GoArgument{In:true, Out:false, Name:"s", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(s>>0))
	msg = append(msg, byte(s>>8))
	msg = append(msg, byte(s>>16))
	msg = append(msg, byte(s>>24))
	// len : 4 byte main.GoArgument{In:true, Out:false, Name:"len", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(len>>0))
	msg = append(msg, byte(len>>8))
	msg = append(msg, byte(len>>16))
	msg = append(msg, byte(len>>24))
	// flags : 4 byte main.GoArgument{In:true, Out:false, Name:"flags", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(flags>>0))
	msg = append(msg, byte(flags>>8))
	msg = append(msg, byte(flags>>16))
	msg = append(msg, byte(flags>>24))
	// fromlen : 4 byte main.GoArgument{In:true, Out:true, Name:"fromlen", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(fromlen>>0))
	msg = append(msg, byte(fromlen>>8))
	msg = append(msg, byte(fromlen>>16))
	msg = append(msg, byte(fromlen>>24))
	// timeout : 4 byte main.GoArgument{In:true, Out:false, Name:"timeout", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(timeout>>0))
	msg = append(msg, byte(timeout>>8))
	msg = append(msg, byte(timeout>>16))
	msg = append(msg, byte(timeout>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// mem : 0 byte main.GoArgument{In:false, Out:true, Name:"mem", Typ:[]string{"[]byte"}, Nullable:false}
	mem_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if mem_length > 0 {
		copy(mem, payload[widx:widx+int(mem_length)])
		widx += int(mem_length)
	}
	// from : 0 byte main.GoArgument{In:false, Out:true, Name:"from", Typ:[]string{"[]byte"}, Nullable:false}
	from_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if from_length > 0 {
		copy(from, payload[widx:widx+int(from_length)])
		widx += int(from_length)
	}
	// fromlen : 4 byte main.GoArgument{In:true, Out:true, Name:"fromlen", Typ:[]string{"uint32"}, Nullable:false}
// not impl

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_lwip_send(s int32, dataptr []byte, flags int32) (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_send()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x0F, uint32(seq))

	// s : 4 byte main.GoArgument{In:true, Out:false, Name:"s", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(s>>0))
	msg = append(msg, byte(s>>8))
	msg = append(msg, byte(s>>16))
	msg = append(msg, byte(s>>24))
	// dataptr : 0 byte main.GoArgument{In:true, Out:false, Name:"dataptr", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, dataptr...)
	// flags : 4 byte main.GoArgument{In:true, Out:false, Name:"flags", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(flags>>0))
	msg = append(msg, byte(flags>>8))
	msg = append(msg, byte(flags>>16))
	msg = append(msg, byte(flags>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_lwip_sendmsg(s int32, msg_name []byte, msg_iov []byte, msg_control []byte, msg_flags int32, flags int32) (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_sendmsg()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x10, uint32(seq))

	// s : 4 byte main.GoArgument{In:true, Out:false, Name:"s", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(s>>0))
	msg = append(msg, byte(s>>8))
	msg = append(msg, byte(s>>16))
	msg = append(msg, byte(s>>24))
	// msg_name : 0 byte main.GoArgument{In:true, Out:false, Name:"msg_name", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, msg_name...)
	// msg_iov : 0 byte main.GoArgument{In:true, Out:false, Name:"msg_iov", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, msg_iov...)
	// msg_control : 0 byte main.GoArgument{In:true, Out:false, Name:"msg_control", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, msg_control...)
	// msg_flags : 4 byte main.GoArgument{In:true, Out:false, Name:"msg_flags", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(msg_flags>>0))
	msg = append(msg, byte(msg_flags>>8))
	msg = append(msg, byte(msg_flags>>16))
	msg = append(msg, byte(msg_flags>>24))
	// flags : 4 byte main.GoArgument{In:true, Out:false, Name:"flags", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(flags>>0))
	msg = append(msg, byte(flags>>8))
	msg = append(msg, byte(flags>>16))
	msg = append(msg, byte(flags>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_lwip_sendto(s int32, dataptr []byte, flags int32, to []byte, tolen uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_sendto()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x11, uint32(seq))

	// s : 4 byte main.GoArgument{In:true, Out:false, Name:"s", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(s>>0))
	msg = append(msg, byte(s>>8))
	msg = append(msg, byte(s>>16))
	msg = append(msg, byte(s>>24))
	// dataptr : 0 byte main.GoArgument{In:true, Out:false, Name:"dataptr", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, dataptr...)
	// flags : 4 byte main.GoArgument{In:true, Out:false, Name:"flags", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(flags>>0))
	msg = append(msg, byte(flags>>8))
	msg = append(msg, byte(flags>>16))
	msg = append(msg, byte(flags>>24))
	// to : 0 byte main.GoArgument{In:true, Out:false, Name:"to", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, to...)
	// tolen : 4 byte main.GoArgument{In:true, Out:false, Name:"tolen", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(tolen>>0))
	msg = append(msg, byte(tolen>>8))
	msg = append(msg, byte(tolen>>16))
	msg = append(msg, byte(tolen>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_lwip_socket(domain int32, l_type int32, protocol int32) (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_socket()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x12, uint32(seq))

	// domain : 4 byte main.GoArgument{In:true, Out:false, Name:"domain", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(domain>>0))
	msg = append(msg, byte(domain>>8))
	msg = append(msg, byte(domain>>16))
	msg = append(msg, byte(domain>>24))
	// l_type : 4 byte main.GoArgument{In:true, Out:false, Name:"l_type", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(l_type>>0))
	msg = append(msg, byte(l_type>>8))
	msg = append(msg, byte(l_type>>16))
	msg = append(msg, byte(l_type>>24))
	// protocol : 4 byte main.GoArgument{In:true, Out:false, Name:"protocol", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(protocol>>0))
	msg = append(msg, byte(protocol>>8))
	msg = append(msg, byte(protocol>>16))
	msg = append(msg, byte(protocol>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_lwip_write(s int32, dataptr []byte, size uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_write()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x13, uint32(seq))

	// s : 4 byte main.GoArgument{In:true, Out:false, Name:"s", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(s>>0))
	msg = append(msg, byte(s>>8))
	msg = append(msg, byte(s>>16))
	msg = append(msg, byte(s>>24))
	// dataptr : 0 byte main.GoArgument{In:true, Out:false, Name:"dataptr", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, dataptr...)
	// size : 4 byte main.GoArgument{In:true, Out:false, Name:"size", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(size>>0))
	msg = append(msg, byte(size>>8))
	msg = append(msg, byte(size>>16))
	msg = append(msg, byte(size>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_lwip_writev(s int32, iov []byte, iovcnt int32) (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_writev()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x14, uint32(seq))

	// s : 4 byte main.GoArgument{In:true, Out:false, Name:"s", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(s>>0))
	msg = append(msg, byte(s>>8))
	msg = append(msg, byte(s>>16))
	msg = append(msg, byte(s>>24))
	// iov : 0 byte main.GoArgument{In:true, Out:false, Name:"iov", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, iov...)
	// iovcnt : 4 byte main.GoArgument{In:true, Out:false, Name:"iovcnt", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(iovcnt>>0))
	msg = append(msg, byte(iovcnt>>8))
	msg = append(msg, byte(iovcnt>>16))
	msg = append(msg, byte(iovcnt>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_lwip_select(maxfdp1 int32, readset []byte, writeset []byte, exceptset []byte, timeout []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_select()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x15, uint32(seq))

	// maxfdp1 : 4 byte main.GoArgument{In:true, Out:false, Name:"maxfdp1", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(maxfdp1>>0))
	msg = append(msg, byte(maxfdp1>>8))
	msg = append(msg, byte(maxfdp1>>16))
	msg = append(msg, byte(maxfdp1>>24))
	// readset : 0 byte main.GoArgument{In:true, Out:false, Name:"readset", Typ:[]string{"[]byte"}, Nullable:true}
	msg = append(msg, readset...)
	// writeset : 0 byte main.GoArgument{In:true, Out:false, Name:"writeset", Typ:[]string{"[]byte"}, Nullable:true}
	msg = append(msg, writeset...)
	// exceptset : 0 byte main.GoArgument{In:true, Out:false, Name:"exceptset", Typ:[]string{"[]byte"}, Nullable:true}
	msg = append(msg, exceptset...)
	// timeout : 0 byte main.GoArgument{In:true, Out:false, Name:"timeout", Typ:[]string{"[]byte"}, Nullable:true}
	msg = append(msg, timeout...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_lwip_ioctl(s int32, cmd uint32, in_argp []byte, out_argp []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_ioctl()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x16, uint32(seq))

	// s : 4 byte main.GoArgument{In:true, Out:false, Name:"s", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(s>>0))
	msg = append(msg, byte(s>>8))
	msg = append(msg, byte(s>>16))
	msg = append(msg, byte(s>>24))
	// cmd : 4 byte main.GoArgument{In:true, Out:false, Name:"cmd", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(cmd>>0))
	msg = append(msg, byte(cmd>>8))
	msg = append(msg, byte(cmd>>16))
	msg = append(msg, byte(cmd>>24))
	// in_argp : 0 byte main.GoArgument{In:true, Out:false, Name:"in_argp", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, in_argp...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// out_argp : 0 byte main.GoArgument{In:false, Out:true, Name:"out_argp", Typ:[]string{"[]byte"}, Nullable:false}
	out_argp_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if out_argp_length > 0 {
		copy(out_argp, payload[widx:widx+int(out_argp_length)])
		widx += int(out_argp_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_lwip_fcntl(s int32, cmd int32, val int32) (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_fcntl()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x17, uint32(seq))

	// s : 4 byte main.GoArgument{In:true, Out:false, Name:"s", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(s>>0))
	msg = append(msg, byte(s>>8))
	msg = append(msg, byte(s>>16))
	msg = append(msg, byte(s>>24))
	// cmd : 4 byte main.GoArgument{In:true, Out:false, Name:"cmd", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(cmd>>0))
	msg = append(msg, byte(cmd>>8))
	msg = append(msg, byte(cmd>>16))
	msg = append(msg, byte(cmd>>24))
	// val : 4 byte main.GoArgument{In:true, Out:false, Name:"val", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(val>>0))
	msg = append(msg, byte(val>>8))
	msg = append(msg, byte(val>>16))
	msg = append(msg, byte(val>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_lwip_errno() (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_errno()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x18, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_netconn_gethostbyname(name string, addr []byte) (int8, error) {
	if debug {
		fmt.Printf("rpc_netconn_gethostbyname()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x19, uint32(seq))

	// name : 0 byte main.GoArgument{In:true, Out:false, Name:"name", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(name)), byte(len(name)>>8), byte(len(name)>>16), byte(len(name)>>24))
	msg = append(msg, []byte(name)...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// addr : 0 byte main.GoArgument{In:false, Out:true, Name:"addr", Typ:[]string{"[]byte"}, Nullable:false}
	addr_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if addr_length > 0 {
		copy(addr, payload[widx:widx+int(addr_length)])
		widx += int(addr_length)
	}

	var result int8
	result = int8(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_dns_gethostbyname_addrtype(hostname string, addr []byte, found uint32, callback_arg []byte, dns_addrtype uint8) (int8, error) {
	if debug {
		fmt.Printf("rpc_dns_gethostbyname_addrtype()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x1A, uint32(seq))

	// hostname : 0 byte main.GoArgument{In:true, Out:false, Name:"hostname", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(hostname)), byte(len(hostname)>>8), byte(len(hostname)>>16), byte(len(hostname)>>24))
	msg = append(msg, []byte(hostname)...)
	// found : 4 byte main.GoArgument{In:true, Out:false, Name:"found", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(found>>0))
	msg = append(msg, byte(found>>8))
	msg = append(msg, byte(found>>16))
	msg = append(msg, byte(found>>24))
	// callback_arg : 0 byte main.GoArgument{In:true, Out:false, Name:"callback_arg", Typ:[]string{"[]byte"}, Nullable:true}
	msg = append(msg, callback_arg...)
	// dns_addrtype : 1 byte main.GoArgument{In:true, Out:false, Name:"dns_addrtype", Typ:[]string{"uint8"}, Nullable:false}
	msg = append(msg, byte(dns_addrtype>>0))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// addr : 0 byte main.GoArgument{In:false, Out:true, Name:"addr", Typ:[]string{"[]byte"}, Nullable:false}
	addr_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if addr_length > 0 {
		copy(addr, payload[widx:widx+int(addr_length)])
		widx += int(addr_length)
	}

	var result int8
	result = int8(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_ssl_client_create() (uint32, error) {
	if debug {
		fmt.Printf("rpc_wifi_ssl_client_create()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x01, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result uint32
	result = uint32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_ssl_client_destroy(ssl_client uint32) error {
	if debug {
		fmt.Printf("rpc_wifi_ssl_client_destroy()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x02, uint32(seq))

	// ssl_client : 4 byte main.GoArgument{In:true, Out:false, Name:"ssl_client", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(ssl_client>>0))
	msg = append(msg, byte(ssl_client>>8))
	msg = append(msg, byte(ssl_client>>16))
	msg = append(msg, byte(ssl_client>>24))

	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received


	seq++
	return err
}

func rpc_wifi_ssl_init(ssl_client uint32) error {
	if debug {
		fmt.Printf("rpc_wifi_ssl_init()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x03, uint32(seq))

	// ssl_client : 4 byte main.GoArgument{In:true, Out:false, Name:"ssl_client", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(ssl_client>>0))
	msg = append(msg, byte(ssl_client>>8))
	msg = append(msg, byte(ssl_client>>16))
	msg = append(msg, byte(ssl_client>>24))

	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received


	seq++
	return err
}

func rpc_wifi_ssl_set_socket(ssl_client uint32, socket int32) error {
	if debug {
		fmt.Printf("rpc_wifi_ssl_set_socket()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x04, uint32(seq))

	// ssl_client : 4 byte main.GoArgument{In:true, Out:false, Name:"ssl_client", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(ssl_client>>0))
	msg = append(msg, byte(ssl_client>>8))
	msg = append(msg, byte(ssl_client>>16))
	msg = append(msg, byte(ssl_client>>24))
	// socket : 4 byte main.GoArgument{In:true, Out:false, Name:"socket", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(socket>>0))
	msg = append(msg, byte(socket>>8))
	msg = append(msg, byte(socket>>16))
	msg = append(msg, byte(socket>>24))

	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received


	seq++
	return err
}

func rpc_wifi_ssl_set_timeout(ssl_client uint32, timeout uint32) error {
	if debug {
		fmt.Printf("rpc_wifi_ssl_set_timeout()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x05, uint32(seq))

	// ssl_client : 4 byte main.GoArgument{In:true, Out:false, Name:"ssl_client", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(ssl_client>>0))
	msg = append(msg, byte(ssl_client>>8))
	msg = append(msg, byte(ssl_client>>16))
	msg = append(msg, byte(ssl_client>>24))
	// timeout : 4 byte main.GoArgument{In:true, Out:false, Name:"timeout", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(timeout>>0))
	msg = append(msg, byte(timeout>>8))
	msg = append(msg, byte(timeout>>16))
	msg = append(msg, byte(timeout>>24))

	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received


	seq++
	return err
}

func rpc_wifi_ssl_get_socket(ssl_client uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_ssl_get_socket()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x06, uint32(seq))

	// ssl_client : 4 byte main.GoArgument{In:true, Out:false, Name:"ssl_client", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(ssl_client>>0))
	msg = append(msg, byte(ssl_client>>8))
	msg = append(msg, byte(ssl_client>>16))
	msg = append(msg, byte(ssl_client>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_ssl_get_timeout(ssl_client uint32) (uint32, error) {
	if debug {
		fmt.Printf("rpc_wifi_ssl_get_timeout()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x07, uint32(seq))

	// ssl_client : 4 byte main.GoArgument{In:true, Out:false, Name:"ssl_client", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(ssl_client>>0))
	msg = append(msg, byte(ssl_client>>8))
	msg = append(msg, byte(ssl_client>>16))
	msg = append(msg, byte(ssl_client>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result uint32
	result = uint32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_ssl_set_rootCA(ssl_client uint32, rootCABuff string) (uint32, error) {
	if debug {
		fmt.Printf("rpc_wifi_ssl_set_rootCA()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x08, uint32(seq))

	// ssl_client : 4 byte main.GoArgument{In:true, Out:false, Name:"ssl_client", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(ssl_client>>0))
	msg = append(msg, byte(ssl_client>>8))
	msg = append(msg, byte(ssl_client>>16))
	msg = append(msg, byte(ssl_client>>24))
	// rootCABuff : 0 byte main.GoArgument{In:true, Out:false, Name:"rootCABuff", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(rootCABuff)), byte(len(rootCABuff)>>8), byte(len(rootCABuff)>>16), byte(len(rootCABuff)>>24))
	msg = append(msg, []byte(rootCABuff)...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result uint32
	result = uint32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_ssl_get_rootCA(ssl_client uint32, rootCABuff string) (uint32, error) {
	if debug {
		fmt.Printf("rpc_wifi_ssl_get_rootCA()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x09, uint32(seq))

	// ssl_client : 4 byte main.GoArgument{In:true, Out:false, Name:"ssl_client", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(ssl_client>>0))
	msg = append(msg, byte(ssl_client>>8))
	msg = append(msg, byte(ssl_client>>16))
	msg = append(msg, byte(ssl_client>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// rootCABuff : 0 byte main.GoArgument{In:false, Out:true, Name:"rootCABuff", Typ:[]string{"string"}, Nullable:true}
	rootCABuff_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 1
	if rootCABuff_length == 1 {
		rootCABuff_length = binary.LittleEndian.Uint32(payload[widx:])
		widx += 4
	}
	if rootCABuff_length > 0 {
		rootCABuff = string(payload[widx:widx+int(rootCABuff_length)])
		widx += int(rootCABuff_length)
	}

	var result uint32
	result = uint32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_ssl_set_cliCert(ssl_client uint32, cli_cert string) (uint32, error) {
	if debug {
		fmt.Printf("rpc_wifi_ssl_set_cliCert()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x0A, uint32(seq))

	// ssl_client : 4 byte main.GoArgument{In:true, Out:false, Name:"ssl_client", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(ssl_client>>0))
	msg = append(msg, byte(ssl_client>>8))
	msg = append(msg, byte(ssl_client>>16))
	msg = append(msg, byte(ssl_client>>24))
	// cli_cert : 0 byte main.GoArgument{In:true, Out:false, Name:"cli_cert", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(cli_cert)), byte(len(cli_cert)>>8), byte(len(cli_cert)>>16), byte(len(cli_cert)>>24))
	msg = append(msg, []byte(cli_cert)...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result uint32
	result = uint32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_ssl_get_cliCert(ssl_client uint32, cli_cert string) (uint32, error) {
	if debug {
		fmt.Printf("rpc_wifi_ssl_get_cliCert()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x0B, uint32(seq))

	// ssl_client : 4 byte main.GoArgument{In:true, Out:false, Name:"ssl_client", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(ssl_client>>0))
	msg = append(msg, byte(ssl_client>>8))
	msg = append(msg, byte(ssl_client>>16))
	msg = append(msg, byte(ssl_client>>24))
	// cli_cert : 0 byte main.GoArgument{In:true, Out:false, Name:"cli_cert", Typ:[]string{"string"}, Nullable:true}
	if len(cli_cert) == 0 {
		msg = append(msg, 1)
	} else {
		msg = append(msg, 0)
		msg = append(msg, byte(len(cli_cert)), byte(len(cli_cert)>>8), byte(len(cli_cert)>>16), byte(len(cli_cert)>>24))
		msg = append(msg, []byte(cli_cert)...)
	}

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result uint32
	result = uint32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_ssl_set_cliKey(ssl_client uint32, cli_key string) (uint32, error) {
	if debug {
		fmt.Printf("rpc_wifi_ssl_set_cliKey()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x0C, uint32(seq))

	// ssl_client : 4 byte main.GoArgument{In:true, Out:false, Name:"ssl_client", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(ssl_client>>0))
	msg = append(msg, byte(ssl_client>>8))
	msg = append(msg, byte(ssl_client>>16))
	msg = append(msg, byte(ssl_client>>24))
	// cli_key : 0 byte main.GoArgument{In:true, Out:false, Name:"cli_key", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(cli_key)), byte(len(cli_key)>>8), byte(len(cli_key)>>16), byte(len(cli_key)>>24))
	msg = append(msg, []byte(cli_key)...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result uint32
	result = uint32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_ssl_get_cliKey(ssl_client uint32, cli_key string) (uint32, error) {
	if debug {
		fmt.Printf("rpc_wifi_ssl_get_cliKey()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x0D, uint32(seq))

	// ssl_client : 4 byte main.GoArgument{In:true, Out:false, Name:"ssl_client", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(ssl_client>>0))
	msg = append(msg, byte(ssl_client>>8))
	msg = append(msg, byte(ssl_client>>16))
	msg = append(msg, byte(ssl_client>>24))
	// cli_key : 0 byte main.GoArgument{In:true, Out:false, Name:"cli_key", Typ:[]string{"string"}, Nullable:true}
	if len(cli_key) == 0 {
		msg = append(msg, 1)
	} else {
		msg = append(msg, 0)
		msg = append(msg, byte(len(cli_key)), byte(len(cli_key)>>8), byte(len(cli_key)>>16), byte(len(cli_key)>>24))
		msg = append(msg, []byte(cli_key)...)
	}

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result uint32
	result = uint32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_ssl_set_pskIdent(ssl_client uint32, pskIdent string) (uint32, error) {
	if debug {
		fmt.Printf("rpc_wifi_ssl_set_pskIdent()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x0E, uint32(seq))

	// ssl_client : 4 byte main.GoArgument{In:true, Out:false, Name:"ssl_client", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(ssl_client>>0))
	msg = append(msg, byte(ssl_client>>8))
	msg = append(msg, byte(ssl_client>>16))
	msg = append(msg, byte(ssl_client>>24))
	// pskIdent : 0 byte main.GoArgument{In:true, Out:false, Name:"pskIdent", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(pskIdent)), byte(len(pskIdent)>>8), byte(len(pskIdent)>>16), byte(len(pskIdent)>>24))
	msg = append(msg, []byte(pskIdent)...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result uint32
	result = uint32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_ssl_get_pskIdent(ssl_client uint32, pskIdent string) (uint32, error) {
	if debug {
		fmt.Printf("rpc_wifi_ssl_get_pskIdent()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x0F, uint32(seq))

	// ssl_client : 4 byte main.GoArgument{In:true, Out:false, Name:"ssl_client", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(ssl_client>>0))
	msg = append(msg, byte(ssl_client>>8))
	msg = append(msg, byte(ssl_client>>16))
	msg = append(msg, byte(ssl_client>>24))
	// pskIdent : 0 byte main.GoArgument{In:true, Out:false, Name:"pskIdent", Typ:[]string{"string"}, Nullable:true}
	if len(pskIdent) == 0 {
		msg = append(msg, 1)
	} else {
		msg = append(msg, 0)
		msg = append(msg, byte(len(pskIdent)), byte(len(pskIdent)>>8), byte(len(pskIdent)>>16), byte(len(pskIdent)>>24))
		msg = append(msg, []byte(pskIdent)...)
	}

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result uint32
	result = uint32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_ssl_set_psKey(ssl_client uint32, psKey string) (uint32, error) {
	if debug {
		fmt.Printf("rpc_wifi_ssl_set_psKey()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x10, uint32(seq))

	// ssl_client : 4 byte main.GoArgument{In:true, Out:false, Name:"ssl_client", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(ssl_client>>0))
	msg = append(msg, byte(ssl_client>>8))
	msg = append(msg, byte(ssl_client>>16))
	msg = append(msg, byte(ssl_client>>24))
	// psKey : 0 byte main.GoArgument{In:true, Out:false, Name:"psKey", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(psKey)), byte(len(psKey)>>8), byte(len(psKey)>>16), byte(len(psKey)>>24))
	msg = append(msg, []byte(psKey)...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result uint32
	result = uint32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_ssl_get_psKey(ssl_client uint32, psKey string) (uint32, error) {
	if debug {
		fmt.Printf("rpc_wifi_ssl_get_psKey()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x11, uint32(seq))

	// ssl_client : 4 byte main.GoArgument{In:true, Out:false, Name:"ssl_client", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(ssl_client>>0))
	msg = append(msg, byte(ssl_client>>8))
	msg = append(msg, byte(ssl_client>>16))
	msg = append(msg, byte(ssl_client>>24))
	// psKey : 0 byte main.GoArgument{In:true, Out:false, Name:"psKey", Typ:[]string{"string"}, Nullable:true}
	if len(psKey) == 0 {
		msg = append(msg, 1)
	} else {
		msg = append(msg, 0)
		msg = append(msg, byte(len(psKey)), byte(len(psKey)>>8), byte(len(psKey)>>16), byte(len(psKey)>>24))
		msg = append(msg, []byte(psKey)...)
	}

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result uint32
	result = uint32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_start_ssl_client(ssl_client uint32, host string, port uint32, timeout int32) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_start_ssl_client()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x12, uint32(seq))

	// ssl_client : 4 byte main.GoArgument{In:true, Out:false, Name:"ssl_client", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(ssl_client>>0))
	msg = append(msg, byte(ssl_client>>8))
	msg = append(msg, byte(ssl_client>>16))
	msg = append(msg, byte(ssl_client>>24))
	// host : 0 byte main.GoArgument{In:true, Out:false, Name:"host", Typ:[]string{"string"}, Nullable:true}
	if len(host) == 0 {
		msg = append(msg, 1)
	} else {
		msg = append(msg, 0)
		msg = append(msg, byte(len(host)), byte(len(host)>>8), byte(len(host)>>16), byte(len(host)>>24))
		msg = append(msg, []byte(host)...)
	}
	// port : 4 byte main.GoArgument{In:true, Out:false, Name:"port", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(port>>0))
	msg = append(msg, byte(port>>8))
	msg = append(msg, byte(port>>16))
	msg = append(msg, byte(port>>24))
	// timeout : 4 byte main.GoArgument{In:true, Out:false, Name:"timeout", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(timeout>>0))
	msg = append(msg, byte(timeout>>8))
	msg = append(msg, byte(timeout>>16))
	msg = append(msg, byte(timeout>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_stop_ssl_socket(ssl_client uint32) error {
	if debug {
		fmt.Printf("rpc_wifi_stop_ssl_socket()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x13, uint32(seq))

	// ssl_client : 4 byte main.GoArgument{In:true, Out:false, Name:"ssl_client", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(ssl_client>>0))
	msg = append(msg, byte(ssl_client>>8))
	msg = append(msg, byte(ssl_client>>16))
	msg = append(msg, byte(ssl_client>>24))

	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received


	seq++
	return err
}

func rpc_wifi_data_to_read(ssl_client uint32) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_data_to_read()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x14, uint32(seq))

	// ssl_client : 4 byte main.GoArgument{In:true, Out:false, Name:"ssl_client", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(ssl_client>>0))
	msg = append(msg, byte(ssl_client>>8))
	msg = append(msg, byte(ssl_client>>16))
	msg = append(msg, byte(ssl_client>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_send_ssl_data(ssl_client uint32, data []byte, len uint16) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_send_ssl_data()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x15, uint32(seq))

	// ssl_client : 4 byte main.GoArgument{In:true, Out:false, Name:"ssl_client", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(ssl_client>>0))
	msg = append(msg, byte(ssl_client>>8))
	msg = append(msg, byte(ssl_client>>16))
	msg = append(msg, byte(ssl_client>>24))
	// data : 0 byte main.GoArgument{In:true, Out:false, Name:"data", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, data...)
	// len : 2 byte main.GoArgument{In:true, Out:false, Name:"len", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(len>>0))
	msg = append(msg, byte(len>>8))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_get_ssl_receive(ssl_client uint32, data []byte, length int32) (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_ssl_receive()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x16, uint32(seq))

	// ssl_client : 4 byte main.GoArgument{In:true, Out:false, Name:"ssl_client", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(ssl_client>>0))
	msg = append(msg, byte(ssl_client>>8))
	msg = append(msg, byte(ssl_client>>16))
	msg = append(msg, byte(ssl_client>>24))
	// length : 4 byte main.GoArgument{In:true, Out:false, Name:"length", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(length>>0))
	msg = append(msg, byte(length>>8))
	msg = append(msg, byte(length>>16))
	msg = append(msg, byte(length>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// data : 0 byte main.GoArgument{In:false, Out:true, Name:"data", Typ:[]string{"[]byte"}, Nullable:false}
	data_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if data_length > 0 {
		copy(data, payload[widx:widx+int(data_length)])
		widx += int(data_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_verify_ssl_fingerprint(ssl_client uint32, fp string, domain_name string) (bool, error) {
	if debug {
		fmt.Printf("rpc_wifi_verify_ssl_fingerprint()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x17, uint32(seq))

	// ssl_client : 4 byte main.GoArgument{In:true, Out:false, Name:"ssl_client", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(ssl_client>>0))
	msg = append(msg, byte(ssl_client>>8))
	msg = append(msg, byte(ssl_client>>16))
	msg = append(msg, byte(ssl_client>>24))
	// fp : 0 byte main.GoArgument{In:true, Out:false, Name:"fp", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(fp)), byte(len(fp)>>8), byte(len(fp)>>16), byte(len(fp)>>24))
	msg = append(msg, []byte(fp)...)
	// domain_name : 0 byte main.GoArgument{In:true, Out:false, Name:"domain_name", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(domain_name)), byte(len(domain_name)>>8), byte(len(domain_name)>>16), byte(len(domain_name)>>24))
	msg = append(msg, []byte(domain_name)...)

	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	widx := 8

	var result bool
	result = binary.LittleEndian.Uint32(payload[widx:]) == 1

	seq++
	return result, err
}

func rpc_wifi_verify_ssl_dn(ssl_client uint32, domain_name string) (bool, error) {
	if debug {
		fmt.Printf("rpc_wifi_verify_ssl_dn()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x18, uint32(seq))

	// ssl_client : 4 byte main.GoArgument{In:true, Out:false, Name:"ssl_client", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(ssl_client>>0))
	msg = append(msg, byte(ssl_client>>8))
	msg = append(msg, byte(ssl_client>>16))
	msg = append(msg, byte(ssl_client>>24))
	// domain_name : 0 byte main.GoArgument{In:true, Out:false, Name:"domain_name", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(domain_name)), byte(len(domain_name)>>8), byte(len(domain_name)>>16), byte(len(domain_name)>>24))
	msg = append(msg, []byte(domain_name)...)

	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	widx := 8

	var result bool
	result = binary.LittleEndian.Uint32(payload[widx:]) == 1

	seq++
	return result, err
}

func rpc_wifi_ssl_strerror(errnum int32, buffer []byte, buflen uint32) error {
	if debug {
		fmt.Printf("rpc_wifi_ssl_strerror()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x19, uint32(seq))

	// errnum : 4 byte main.GoArgument{In:true, Out:false, Name:"errnum", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(errnum>>0))
	msg = append(msg, byte(errnum>>8))
	msg = append(msg, byte(errnum>>16))
	msg = append(msg, byte(errnum>>24))
	// buflen : 4 byte main.GoArgument{In:true, Out:false, Name:"buflen", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(buflen>>0))
	msg = append(msg, byte(buflen>>8))
	msg = append(msg, byte(buflen>>16))
	msg = append(msg, byte(buflen>>24))

	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received
	widx := 8
	// buffer : 0 byte main.GoArgument{In:false, Out:true, Name:"buffer", Typ:[]string{"[]byte"}, Nullable:false}
	buffer_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if buffer_length > 0 {
		copy(buffer, payload[widx:widx+int(buffer_length)])
		widx += int(buffer_length)
	}


	seq++
	return err
}

func rpc_mdns_init() (int32, error) {
	if debug {
		fmt.Printf("rpc_mdns_init()\n")
	}
	msg := startWriteMessage(0x00, 0x12, 0x01, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_mdns_free() (int32, error) {
	if debug {
		fmt.Printf("rpc_mdns_free()\n")
	}
	msg := startWriteMessage(0x00, 0x12, 0x02, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_mdns_service_add(instance_name string, service_type string, proto string, port uint16) (int32, error) {
	if debug {
		fmt.Printf("rpc_mdns_service_add()\n")
	}
	msg := startWriteMessage(0x00, 0x12, 0x03, uint32(seq))

	// instance_name : 0 byte main.GoArgument{In:true, Out:false, Name:"instance_name", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(instance_name)), byte(len(instance_name)>>8), byte(len(instance_name)>>16), byte(len(instance_name)>>24))
	msg = append(msg, []byte(instance_name)...)
	// service_type : 0 byte main.GoArgument{In:true, Out:false, Name:"service_type", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(service_type)), byte(len(service_type)>>8), byte(len(service_type)>>16), byte(len(service_type)>>24))
	msg = append(msg, []byte(service_type)...)
	// proto : 0 byte main.GoArgument{In:true, Out:false, Name:"proto", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(proto)), byte(len(proto)>>8), byte(len(proto)>>16), byte(len(proto)>>24))
	msg = append(msg, []byte(proto)...)
	// port : 2 byte main.GoArgument{In:true, Out:false, Name:"port", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(port>>0))
	msg = append(msg, byte(port>>8))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_mdns_service_remove(service_type string, proto string) (int32, error) {
	if debug {
		fmt.Printf("rpc_mdns_service_remove()\n")
	}
	msg := startWriteMessage(0x00, 0x12, 0x04, uint32(seq))

	// service_type : 0 byte main.GoArgument{In:true, Out:false, Name:"service_type", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(service_type)), byte(len(service_type)>>8), byte(len(service_type)>>16), byte(len(service_type)>>24))
	msg = append(msg, []byte(service_type)...)
	// proto : 0 byte main.GoArgument{In:true, Out:false, Name:"proto", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(proto)), byte(len(proto)>>8), byte(len(proto)>>16), byte(len(proto)>>24))
	msg = append(msg, []byte(proto)...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_mdns_service_txt_item_set(service_type string, proto string, key string, value string) (int32, error) {
	if debug {
		fmt.Printf("rpc_mdns_service_txt_item_set()\n")
	}
	msg := startWriteMessage(0x00, 0x12, 0x05, uint32(seq))

	// service_type : 0 byte main.GoArgument{In:true, Out:false, Name:"service_type", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(service_type)), byte(len(service_type)>>8), byte(len(service_type)>>16), byte(len(service_type)>>24))
	msg = append(msg, []byte(service_type)...)
	// proto : 0 byte main.GoArgument{In:true, Out:false, Name:"proto", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(proto)), byte(len(proto)>>8), byte(len(proto)>>16), byte(len(proto)>>24))
	msg = append(msg, []byte(proto)...)
	// key : 0 byte main.GoArgument{In:true, Out:false, Name:"key", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(key)), byte(len(key)>>8), byte(len(key)>>16), byte(len(key)>>24))
	msg = append(msg, []byte(key)...)
	// value : 0 byte main.GoArgument{In:true, Out:false, Name:"value", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(value)), byte(len(value)>>8), byte(len(value)>>16), byte(len(value)>>24))
	msg = append(msg, []byte(value)...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_mdns_service_instance_name_set(service string, proto string, instance string) (int32, error) {
	if debug {
		fmt.Printf("rpc_mdns_service_instance_name_set()\n")
	}
	msg := startWriteMessage(0x00, 0x12, 0x06, uint32(seq))

	// service : 0 byte main.GoArgument{In:true, Out:false, Name:"service", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(service)), byte(len(service)>>8), byte(len(service)>>16), byte(len(service)>>24))
	msg = append(msg, []byte(service)...)
	// proto : 0 byte main.GoArgument{In:true, Out:false, Name:"proto", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(proto)), byte(len(proto)>>8), byte(len(proto)>>16), byte(len(proto)>>24))
	msg = append(msg, []byte(proto)...)
	// instance : 0 byte main.GoArgument{In:true, Out:false, Name:"instance", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(instance)), byte(len(instance)>>8), byte(len(instance)>>16), byte(len(instance)>>24))
	msg = append(msg, []byte(instance)...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_mdns_instance_name_set(instance_name string) (int32, error) {
	if debug {
		fmt.Printf("rpc_mdns_instance_name_set()\n")
	}
	msg := startWriteMessage(0x00, 0x12, 0x07, uint32(seq))

	// instance_name : 0 byte main.GoArgument{In:true, Out:false, Name:"instance_name", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(instance_name)), byte(len(instance_name)>>8), byte(len(instance_name)>>16), byte(len(instance_name)>>24))
	msg = append(msg, []byte(instance_name)...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_mdns_hostname_set(hostname string) (int32, error) {
	if debug {
		fmt.Printf("rpc_mdns_hostname_set()\n")
	}
	msg := startWriteMessage(0x00, 0x12, 0x08, uint32(seq))

	// hostname : 0 byte main.GoArgument{In:true, Out:false, Name:"hostname", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(hostname)), byte(len(hostname)>>8), byte(len(hostname)>>16), byte(len(hostname)>>24))
	msg = append(msg, []byte(hostname)...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_mdns_query_a(host_name string, timeout uint32, addr []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_mdns_query_a()\n")
	}
	msg := startWriteMessage(0x00, 0x12, 0x09, uint32(seq))

	// host_name : 0 byte main.GoArgument{In:true, Out:false, Name:"host_name", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(host_name)), byte(len(host_name)>>8), byte(len(host_name)>>16), byte(len(host_name)>>24))
	msg = append(msg, []byte(host_name)...)
	// timeout : 4 byte main.GoArgument{In:true, Out:false, Name:"timeout", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(timeout>>0))
	msg = append(msg, byte(timeout>>8))
	msg = append(msg, byte(timeout>>16))
	msg = append(msg, byte(timeout>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// addr : 0 byte main.GoArgument{In:false, Out:true, Name:"addr", Typ:[]string{"[]byte"}, Nullable:false}
	addr_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if addr_length > 0 {
		copy(addr, payload[widx:widx+int(addr_length)])
		widx += int(addr_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_mdns_query_ptr(service_type string, proto string, timeout uint32, max_results int32, result_total int32) (int32, error) {
	if debug {
		fmt.Printf("rpc_mdns_query_ptr()\n")
	}
	msg := startWriteMessage(0x00, 0x12, 0x0A, uint32(seq))

	// service_type : 0 byte main.GoArgument{In:true, Out:false, Name:"service_type", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(service_type)), byte(len(service_type)>>8), byte(len(service_type)>>16), byte(len(service_type)>>24))
	msg = append(msg, []byte(service_type)...)
	// proto : 0 byte main.GoArgument{In:true, Out:false, Name:"proto", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(proto)), byte(len(proto)>>8), byte(len(proto)>>16), byte(len(proto)>>24))
	msg = append(msg, []byte(proto)...)
	// timeout : 4 byte main.GoArgument{In:true, Out:false, Name:"timeout", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(timeout>>0))
	msg = append(msg, byte(timeout>>8))
	msg = append(msg, byte(timeout>>16))
	msg = append(msg, byte(timeout>>24))
	// max_results : 4 byte main.GoArgument{In:true, Out:false, Name:"max_results", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(max_results>>0))
	msg = append(msg, byte(max_results>>8))
	msg = append(msg, byte(max_results>>16))
	msg = append(msg, byte(max_results>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// result_total : 4 byte main.GoArgument{In:false, Out:true, Name:"result_total", Typ:[]string{"int32"}, Nullable:false}
// not impl

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_mdns_query_ptr_result_basic(result_target int32, scan_result []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_mdns_query_ptr_result_basic()\n")
	}
	msg := startWriteMessage(0x00, 0x12, 0x0B, uint32(seq))

	// result_target : 4 byte main.GoArgument{In:true, Out:false, Name:"result_target", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(result_target>>0))
	msg = append(msg, byte(result_target>>8))
	msg = append(msg, byte(result_target>>16))
	msg = append(msg, byte(result_target>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// scan_result : 0 byte main.GoArgument{In:false, Out:true, Name:"scan_result", Typ:[]string{"[]byte"}, Nullable:false}
	scan_result_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if scan_result_length > 0 {
		copy(scan_result, payload[widx:widx+int(scan_result_length)])
		widx += int(scan_result_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_mdns_query_ptr_result_txt(result_target int32, txt_target int32, txt []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_mdns_query_ptr_result_txt()\n")
	}
	msg := startWriteMessage(0x00, 0x12, 0x0C, uint32(seq))

	// result_target : 4 byte main.GoArgument{In:true, Out:false, Name:"result_target", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(result_target>>0))
	msg = append(msg, byte(result_target>>8))
	msg = append(msg, byte(result_target>>16))
	msg = append(msg, byte(result_target>>24))
	// txt_target : 4 byte main.GoArgument{In:true, Out:false, Name:"txt_target", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(txt_target>>0))
	msg = append(msg, byte(txt_target>>8))
	msg = append(msg, byte(txt_target>>16))
	msg = append(msg, byte(txt_target>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// txt : 0 byte main.GoArgument{In:false, Out:true, Name:"txt", Typ:[]string{"[]byte"}, Nullable:false}
	txt_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if txt_length > 0 {
		copy(txt, payload[widx:widx+int(txt_length)])
		widx += int(txt_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_mdns_query_ptr_result_addr(result_target int32, addr_target int32, addr []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_mdns_query_ptr_result_addr()\n")
	}
	msg := startWriteMessage(0x00, 0x12, 0x0D, uint32(seq))

	// result_target : 4 byte main.GoArgument{In:true, Out:false, Name:"result_target", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(result_target>>0))
	msg = append(msg, byte(result_target>>8))
	msg = append(msg, byte(result_target>>16))
	msg = append(msg, byte(result_target>>24))
	// addr_target : 4 byte main.GoArgument{In:true, Out:false, Name:"addr_target", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(addr_target>>0))
	msg = append(msg, byte(addr_target>>8))
	msg = append(msg, byte(addr_target>>16))
	msg = append(msg, byte(addr_target>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	// addr : 0 byte main.GoArgument{In:false, Out:true, Name:"addr", Typ:[]string{"[]byte"}, Nullable:false}
	addr_length := binary.LittleEndian.Uint32(payload[widx:])
	widx += 4
	if addr_length > 0 {
		copy(addr, payload[widx:widx+int(addr_length)])
		widx += int(addr_length)
	}

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_mdns_query_results_free() (int32, error) {
	if debug {
		fmt.Printf("rpc_mdns_query_results_free()\n")
	}
	msg := startWriteMessage(0x00, 0x12, 0x0E, uint32(seq))


	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8
	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_wifi_event_callback(event []byte) error {
	if debug {
		fmt.Printf("rpc_wifi_event_callback()\n")
	}
	msg := startWriteMessage(0x00, 0x13, 0x01, uint32(seq))

	// event : 0 byte main.GoArgument{In:true, Out:false, Name:"event", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, event...)

	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received


	seq++
	return err
}

func rpc_wifi_dns_found(hostname string, ipaddr []byte, arg []byte) error {
	if debug {
		fmt.Printf("rpc_wifi_dns_found()\n")
	}
	msg := startWriteMessage(0x00, 0x13, 0x02, uint32(seq))

	// hostname : 0 byte main.GoArgument{In:true, Out:false, Name:"hostname", Typ:[]string{"string"}, Nullable:false}
	msg = append(msg, byte(len(hostname)), byte(len(hostname)>>8), byte(len(hostname)>>16), byte(len(hostname)>>24))
	msg = append(msg, []byte(hostname)...)
	// ipaddr : 0 byte main.GoArgument{In:true, Out:false, Name:"ipaddr", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, ipaddr...)
	// arg : 0 byte main.GoArgument{In:true, Out:false, Name:"arg", Typ:[]string{"[]byte"}, Nullable:true}
	msg = append(msg, arg...)

	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received


	seq++
	return err
}

func rpc_tcpip_api_call_fn(fn uint32, call []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_api_call_fn()\n")
	}
	msg := startWriteMessage(0x00, 0x13, 0x03, uint32(seq))

	// fn : 4 byte main.GoArgument{In:true, Out:false, Name:"fn", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(fn>>0))
	msg = append(msg, byte(fn>>8))
	msg = append(msg, byte(fn>>16))
	msg = append(msg, byte(fn>>24))
	// call : 0 byte main.GoArgument{In:true, Out:false, Name:"call", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, call...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcp_connected_fn(fn uint32, arg []byte, tpcb []byte, err_val int32) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_connected_fn()\n")
	}
	msg := startWriteMessage(0x00, 0x13, 0x04, uint32(seq))

	// fn : 4 byte main.GoArgument{In:true, Out:false, Name:"fn", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(fn>>0))
	msg = append(msg, byte(fn>>8))
	msg = append(msg, byte(fn>>16))
	msg = append(msg, byte(fn>>24))
	// arg : 0 byte main.GoArgument{In:true, Out:false, Name:"arg", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, arg...)
	// tpcb : 0 byte main.GoArgument{In:true, Out:false, Name:"tpcb", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, tpcb...)
	// err_val : 4 byte main.GoArgument{In:true, Out:false, Name:"err_val", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(err_val>>0))
	msg = append(msg, byte(err_val>>8))
	msg = append(msg, byte(err_val>>16))
	msg = append(msg, byte(err_val>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcp_recv_fn(fn uint32, arg []byte, tpcb []byte, p_data []byte, p_addr []byte, err_val int32) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_recv_fn()\n")
	}
	msg := startWriteMessage(0x00, 0x13, 0x05, uint32(seq))

	// fn : 4 byte main.GoArgument{In:true, Out:false, Name:"fn", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(fn>>0))
	msg = append(msg, byte(fn>>8))
	msg = append(msg, byte(fn>>16))
	msg = append(msg, byte(fn>>24))
	// arg : 0 byte main.GoArgument{In:true, Out:false, Name:"arg", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, arg...)
	// tpcb : 0 byte main.GoArgument{In:true, Out:false, Name:"tpcb", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, tpcb...)
	// p_data : 0 byte main.GoArgument{In:true, Out:false, Name:"p_data", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, p_data...)
	// p_addr : 0 byte main.GoArgument{In:true, Out:false, Name:"p_addr", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, p_addr...)
	// err_val : 4 byte main.GoArgument{In:true, Out:false, Name:"err_val", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(err_val>>0))
	msg = append(msg, byte(err_val>>8))
	msg = append(msg, byte(err_val>>16))
	msg = append(msg, byte(err_val>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcp_accept_fn(fn uint32, arg []byte, newpcb []byte, err_val int32) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_accept_fn()\n")
	}
	msg := startWriteMessage(0x00, 0x13, 0x06, uint32(seq))

	// fn : 4 byte main.GoArgument{In:true, Out:false, Name:"fn", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(fn>>0))
	msg = append(msg, byte(fn>>8))
	msg = append(msg, byte(fn>>16))
	msg = append(msg, byte(fn>>24))
	// arg : 0 byte main.GoArgument{In:true, Out:false, Name:"arg", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, arg...)
	// newpcb : 0 byte main.GoArgument{In:true, Out:false, Name:"newpcb", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, newpcb...)
	// err_val : 4 byte main.GoArgument{In:true, Out:false, Name:"err_val", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(err_val>>0))
	msg = append(msg, byte(err_val>>8))
	msg = append(msg, byte(err_val>>16))
	msg = append(msg, byte(err_val>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcp_err_fn(fn uint32, arg []byte, err_val int32) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_err_fn()\n")
	}
	msg := startWriteMessage(0x00, 0x13, 0x07, uint32(seq))

	// fn : 4 byte main.GoArgument{In:true, Out:false, Name:"fn", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(fn>>0))
	msg = append(msg, byte(fn>>8))
	msg = append(msg, byte(fn>>16))
	msg = append(msg, byte(fn>>24))
	// arg : 0 byte main.GoArgument{In:true, Out:false, Name:"arg", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, arg...)
	// err_val : 4 byte main.GoArgument{In:true, Out:false, Name:"err_val", Typ:[]string{"int32"}, Nullable:false}
	msg = append(msg, byte(err_val>>0))
	msg = append(msg, byte(err_val>>8))
	msg = append(msg, byte(err_val>>16))
	msg = append(msg, byte(err_val>>24))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcp_sent_fn(fn uint32, arg []byte, tpcb []byte, len uint16) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_sent_fn()\n")
	}
	msg := startWriteMessage(0x00, 0x13, 0x08, uint32(seq))

	// fn : 4 byte main.GoArgument{In:true, Out:false, Name:"fn", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(fn>>0))
	msg = append(msg, byte(fn>>8))
	msg = append(msg, byte(fn>>16))
	msg = append(msg, byte(fn>>24))
	// arg : 0 byte main.GoArgument{In:true, Out:false, Name:"arg", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, arg...)
	// tpcb : 0 byte main.GoArgument{In:true, Out:false, Name:"tpcb", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, tpcb...)
	// len : 2 byte main.GoArgument{In:true, Out:false, Name:"len", Typ:[]string{"uint16"}, Nullable:false}
	msg = append(msg, byte(len>>0))
	msg = append(msg, byte(len>>8))

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

func rpc_tcp_poll_fn(fn uint32, arg []byte, tpcb []byte) (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_poll_fn()\n")
	}
	msg := startWriteMessage(0x00, 0x13, 0x09, uint32(seq))

	// fn : 4 byte main.GoArgument{In:true, Out:false, Name:"fn", Typ:[]string{"uint32"}, Nullable:false}
	msg = append(msg, byte(fn>>0))
	msg = append(msg, byte(fn>>8))
	msg = append(msg, byte(fn>>16))
	msg = append(msg, byte(fn>>24))
	// arg : 0 byte main.GoArgument{In:true, Out:false, Name:"arg", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, arg...)
	// tpcb : 0 byte main.GoArgument{In:true, Out:false, Name:"tpcb", Typ:[]string{"[]byte"}, Nullable:false}
	msg = append(msg, tpcb...)

	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	widx := 8

	var result int32
	result = int32(binary.LittleEndian.Uint32(payload[widx:]))

	seq++
	return result, err
}

