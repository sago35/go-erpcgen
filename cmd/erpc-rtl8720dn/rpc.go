package main

import (
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
	var result string
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_system_ack() (uint8, error) {
	if debug {
		fmt.Printf("rpc_system_ack()\n")
	}
	msg := startWriteMessage(0x00, 0x01, 0x02, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result uint8
	//err = readInt32(&result)

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
	var result bool
	//err = readInt32(&result)

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
	//err = readInt32(&result)

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
	//err = readInt32(&result)

	seq++
	return err
}

func rpc_gap_set_param() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_gap_set_param()\n")
	}
	msg := startWriteMessage(0x00, 0x03, 0x01, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_gap_get_param() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_gap_get_param()\n")
	}
	msg := startWriteMessage(0x00, 0x03, 0x02, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

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
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_bond_set_param() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_bond_set_param()\n")
	}
	msg := startWriteMessage(0x00, 0x04, 0x01, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_bond_get_param() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_bond_get_param()\n")
	}
	msg := startWriteMessage(0x00, 0x04, 0x02, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_bond_pair() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_bond_pair()\n")
	}
	msg := startWriteMessage(0x00, 0x04, 0x03, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_bond_get_display_key() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_bond_get_display_key()\n")
	}
	msg := startWriteMessage(0x00, 0x04, 0x04, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_bond_passkey_input_confirm() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_bond_passkey_input_confirm()\n")
	}
	msg := startWriteMessage(0x00, 0x04, 0x05, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_bond_oob_input_confirm() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_bond_oob_input_confirm()\n")
	}
	msg := startWriteMessage(0x00, 0x04, 0x06, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_bond_just_work_confirm() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_bond_just_work_confirm()\n")
	}
	msg := startWriteMessage(0x00, 0x04, 0x07, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_bond_passkey_display_confirm() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_bond_passkey_display_confirm()\n")
	}
	msg := startWriteMessage(0x00, 0x04, 0x08, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_bond_user_confirm() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_bond_user_confirm()\n")
	}
	msg := startWriteMessage(0x00, 0x04, 0x09, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_bond_cfg_local_key_distribute() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_bond_cfg_local_key_distribute()\n")
	}
	msg := startWriteMessage(0x00, 0x04, 0x0A, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

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
	//err = readInt32(&result)

	seq++
	return err
}

func rpc_le_bond_delete_by_idx() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_bond_delete_by_idx()\n")
	}
	msg := startWriteMessage(0x00, 0x04, 0x0C, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_bond_delete_by_bd() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_bond_delete_by_bd()\n")
	}
	msg := startWriteMessage(0x00, 0x04, 0x0D, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_bond_get_sec_level() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_bond_get_sec_level()\n")
	}
	msg := startWriteMessage(0x00, 0x04, 0x0E, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_gap_init() (bool, error) {
	if debug {
		fmt.Printf("rpc_le_gap_init()\n")
	}
	msg := startWriteMessage(0x00, 0x05, 0x01, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	var result bool
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_gap_msg_info_way() error {
	if debug {
		fmt.Printf("rpc_le_gap_msg_info_way()\n")
	}
	msg := startWriteMessage(0x00, 0x05, 0x02, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received
	//err = readInt32(&result)

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
	var result uint8
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_set_gap_param() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_set_gap_param()\n")
	}
	msg := startWriteMessage(0x00, 0x05, 0x04, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_get_gap_param() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_get_gap_param()\n")
	}
	msg := startWriteMessage(0x00, 0x05, 0x05, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_modify_white_list() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_modify_white_list()\n")
	}
	msg := startWriteMessage(0x00, 0x05, 0x06, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_gen_rand_addr() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_gen_rand_addr()\n")
	}
	msg := startWriteMessage(0x00, 0x05, 0x07, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_set_rand_addr() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_set_rand_addr()\n")
	}
	msg := startWriteMessage(0x00, 0x05, 0x08, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_cfg_local_identity_address() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_cfg_local_identity_address()\n")
	}
	msg := startWriteMessage(0x00, 0x05, 0x09, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_set_host_chann_classif() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_set_host_chann_classif()\n")
	}
	msg := startWriteMessage(0x00, 0x05, 0x0A, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_write_default_data_len() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_write_default_data_len()\n")
	}
	msg := startWriteMessage(0x00, 0x05, 0x0B, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_gap_config_cccd_not_check() error {
	if debug {
		fmt.Printf("rpc_gap_config_cccd_not_check()\n")
	}
	msg := startWriteMessage(0x00, 0x06, 0x01, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received
	//err = readInt32(&result)

	seq++
	return err
}

func rpc_gap_config_ccc_bits_count() error {
	if debug {
		fmt.Printf("rpc_gap_config_ccc_bits_count()\n")
	}
	msg := startWriteMessage(0x00, 0x06, 0x02, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received
	//err = readInt32(&result)

	seq++
	return err
}

func rpc_gap_config_max_attribute_table_count() error {
	if debug {
		fmt.Printf("rpc_gap_config_max_attribute_table_count()\n")
	}
	msg := startWriteMessage(0x00, 0x06, 0x03, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received
	//err = readInt32(&result)

	seq++
	return err
}

func rpc_gap_config_max_mtu_size() error {
	if debug {
		fmt.Printf("rpc_gap_config_max_mtu_size()\n")
	}
	msg := startWriteMessage(0x00, 0x06, 0x04, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received
	//err = readInt32(&result)

	seq++
	return err
}

func rpc_gap_config_bte_pool_size() error {
	if debug {
		fmt.Printf("rpc_gap_config_bte_pool_size()\n")
	}
	msg := startWriteMessage(0x00, 0x06, 0x05, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received
	//err = readInt32(&result)

	seq++
	return err
}

func rpc_gap_config_bt_report_buf_num() error {
	if debug {
		fmt.Printf("rpc_gap_config_bt_report_buf_num()\n")
	}
	msg := startWriteMessage(0x00, 0x06, 0x06, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received
	//err = readInt32(&result)

	seq++
	return err
}

func rpc_gap_config_le_key_storage_flag() error {
	if debug {
		fmt.Printf("rpc_gap_config_le_key_storage_flag()\n")
	}
	msg := startWriteMessage(0x00, 0x06, 0x07, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received
	//err = readInt32(&result)

	seq++
	return err
}

func rpc_gap_config_max_le_paired_device() error {
	if debug {
		fmt.Printf("rpc_gap_config_max_le_paired_device()\n")
	}
	msg := startWriteMessage(0x00, 0x06, 0x08, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received
	//err = readInt32(&result)

	seq++
	return err
}

func rpc_gap_config_max_le_link_num() error {
	if debug {
		fmt.Printf("rpc_gap_config_max_le_link_num()\n")
	}
	msg := startWriteMessage(0x00, 0x06, 0x09, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received
	//err = readInt32(&result)

	seq++
	return err
}

func rpc_le_adv_set_param() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_adv_set_param()\n")
	}
	msg := startWriteMessage(0x00, 0x07, 0x01, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_adv_get_param() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_adv_get_param()\n")
	}
	msg := startWriteMessage(0x00, 0x07, 0x02, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

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
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

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
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

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
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_scan_set_param() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_scan_set_param()\n")
	}
	msg := startWriteMessage(0x00, 0x08, 0x01, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_scan_get_param() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_scan_get_param()\n")
	}
	msg := startWriteMessage(0x00, 0x08, 0x02, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

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
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_scan_timer_start() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_scan_timer_start()\n")
	}
	msg := startWriteMessage(0x00, 0x08, 0x04, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

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
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_scan_info_filter() (bool, error) {
	if debug {
		fmt.Printf("rpc_le_scan_info_filter()\n")
	}
	msg := startWriteMessage(0x00, 0x08, 0x06, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	var result bool
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_get_conn_param() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_get_conn_param()\n")
	}
	msg := startWriteMessage(0x00, 0x09, 0x01, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_get_conn_info() (bool, error) {
	if debug {
		fmt.Printf("rpc_le_get_conn_info()\n")
	}
	msg := startWriteMessage(0x00, 0x09, 0x02, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	var result bool
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_get_conn_addr() (bool, error) {
	if debug {
		fmt.Printf("rpc_le_get_conn_addr()\n")
	}
	msg := startWriteMessage(0x00, 0x09, 0x03, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	var result bool
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_get_conn_id() (bool, error) {
	if debug {
		fmt.Printf("rpc_le_get_conn_id()\n")
	}
	msg := startWriteMessage(0x00, 0x09, 0x04, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	var result bool
	//err = readInt32(&result)

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
	var result uint8
	//err = readInt32(&result)

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
	var result uint8
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_disconnect() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_disconnect()\n")
	}
	msg := startWriteMessage(0x00, 0x09, 0x07, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_read_rssi() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_read_rssi()\n")
	}
	msg := startWriteMessage(0x00, 0x09, 0x08, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_set_data_len() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_set_data_len()\n")
	}
	msg := startWriteMessage(0x00, 0x09, 0x09, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_set_phy() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_set_phy()\n")
	}
	msg := startWriteMessage(0x00, 0x09, 0x0A, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_set_conn_param() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_set_conn_param()\n")
	}
	msg := startWriteMessage(0x00, 0x09, 0x0B, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_connect() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_connect()\n")
	}
	msg := startWriteMessage(0x00, 0x09, 0x0C, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_update_conn_param() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_le_update_conn_param()\n")
	}
	msg := startWriteMessage(0x00, 0x09, 0x0D, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_flash_save_local_name() (uint32, error) {
	if debug {
		fmt.Printf("rpc_flash_save_local_name()\n")
	}
	msg := startWriteMessage(0x00, 0x0A, 0x01, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result uint32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_flash_load_local_name() (uint32, error) {
	if debug {
		fmt.Printf("rpc_flash_load_local_name()\n")
	}
	msg := startWriteMessage(0x00, 0x0A, 0x02, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result uint32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_flash_save_local_appearance() (uint32, error) {
	if debug {
		fmt.Printf("rpc_flash_save_local_appearance()\n")
	}
	msg := startWriteMessage(0x00, 0x0A, 0x03, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result uint32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_flash_load_local_appearance() (uint32, error) {
	if debug {
		fmt.Printf("rpc_flash_load_local_appearance()\n")
	}
	msg := startWriteMessage(0x00, 0x0A, 0x04, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result uint32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_find_key_entry() (RPC_T_LE_KEY_ENTRY, error) {
	if debug {
		fmt.Printf("rpc_le_find_key_entry()\n")
	}
	msg := startWriteMessage(0x00, 0x0A, 0x05, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_LE_KEY_ENTRY
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_find_key_entry_by_idx() (RPC_T_LE_KEY_ENTRY, error) {
	if debug {
		fmt.Printf("rpc_le_find_key_entry_by_idx()\n")
	}
	msg := startWriteMessage(0x00, 0x0A, 0x06, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_LE_KEY_ENTRY
	//err = readInt32(&result)

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
	var result uint8
	//err = readInt32(&result)

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
	var result RPC_T_LE_KEY_ENTRY
	//err = readInt32(&result)

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
	var result RPC_T_LE_KEY_ENTRY
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_set_high_priority_bond() (bool, error) {
	if debug {
		fmt.Printf("rpc_le_set_high_priority_bond()\n")
	}
	msg := startWriteMessage(0x00, 0x0A, 0x0A, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	var result bool
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_resolve_random_address() (bool, error) {
	if debug {
		fmt.Printf("rpc_le_resolve_random_address()\n")
	}
	msg := startWriteMessage(0x00, 0x0A, 0x0B, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	var result bool
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_get_cccd_data() (bool, error) {
	if debug {
		fmt.Printf("rpc_le_get_cccd_data()\n")
	}
	msg := startWriteMessage(0x00, 0x0A, 0x0C, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	var result bool
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_gen_bond_dev() (bool, error) {
	if debug {
		fmt.Printf("rpc_le_gen_bond_dev()\n")
	}
	msg := startWriteMessage(0x00, 0x0A, 0x0D, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	var result bool
	//err = readInt32(&result)

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
	var result uint16
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_set_dev_bond_info() (RPC_T_LE_KEY_ENTRY, error) {
	if debug {
		fmt.Printf("rpc_le_set_dev_bond_info()\n")
	}
	msg := startWriteMessage(0x00, 0x0A, 0x0F, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_LE_KEY_ENTRY
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_le_get_dev_bond_info() (bool, error) {
	if debug {
		fmt.Printf("rpc_le_get_dev_bond_info()\n")
	}
	msg := startWriteMessage(0x00, 0x0A, 0x10, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	var result bool
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_ble_client_init() (bool, error) {
	if debug {
		fmt.Printf("rpc_ble_client_init()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x01, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	var result bool
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_ble_add_client() (uint8, error) {
	if debug {
		fmt.Printf("rpc_ble_add_client()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x02, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result uint8
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_client_init() error {
	if debug {
		fmt.Printf("rpc_client_init()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x03, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received
	//err = readInt32(&result)

	seq++
	return err
}

func rpc_client_all_primary_srv_discovery() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_client_all_primary_srv_discovery()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x04, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_client_by_uuid_srv_discovery() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_client_by_uuid_srv_discovery()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x05, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_client_by_uuid128_srv_discovery() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_client_by_uuid128_srv_discovery()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x06, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_client_relationship_discovery() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_client_relationship_discovery()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x07, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_client_all_char_discovery() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_client_all_char_discovery()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x08, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_client_by_uuid_char_discovery() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_client_by_uuid_char_discovery()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x09, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_client_by_uuid128_char_discovery() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_client_by_uuid128_char_discovery()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x0A, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_client_all_char_descriptor_discovery() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_client_all_char_descriptor_discovery()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x0B, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_client_attr_read() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_client_attr_read()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x0C, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_client_attr_read_using_uuid() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_client_attr_read_using_uuid()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x0D, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_client_attr_write() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_client_attr_write()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x0E, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_client_attr_ind_confirm() (RPC_T_GAP_CAUSE, error) {
	if debug {
		fmt.Printf("rpc_client_attr_ind_confirm()\n")
	}
	msg := startWriteMessage(0x00, 0x0B, 0x0F, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_GAP_CAUSE
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_ble_server_init() (bool, error) {
	if debug {
		fmt.Printf("rpc_ble_server_init()\n")
	}
	msg := startWriteMessage(0x00, 0x0C, 0x01, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	var result bool
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_ble_create_service() (uint8, error) {
	if debug {
		fmt.Printf("rpc_ble_create_service()\n")
	}
	msg := startWriteMessage(0x00, 0x0C, 0x02, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result uint8
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_ble_delete_service() (bool, error) {
	if debug {
		fmt.Printf("rpc_ble_delete_service()\n")
	}
	msg := startWriteMessage(0x00, 0x0C, 0x03, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	var result bool
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_ble_service_start() (uint8, error) {
	if debug {
		fmt.Printf("rpc_ble_service_start()\n")
	}
	msg := startWriteMessage(0x00, 0x0C, 0x04, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result uint8
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_ble_get_servie_handle() (uint8, error) {
	if debug {
		fmt.Printf("rpc_ble_get_servie_handle()\n")
	}
	msg := startWriteMessage(0x00, 0x0C, 0x05, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result uint8
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_ble_create_char() (uint16, error) {
	if debug {
		fmt.Printf("rpc_ble_create_char()\n")
	}
	msg := startWriteMessage(0x00, 0x0C, 0x06, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result uint16
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_ble_create_desc() (uint16, error) {
	if debug {
		fmt.Printf("rpc_ble_create_desc()\n")
	}
	msg := startWriteMessage(0x00, 0x0C, 0x07, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result uint16
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_server_send_data() (bool, error) {
	if debug {
		fmt.Printf("rpc_server_send_data()\n")
	}
	msg := startWriteMessage(0x00, 0x0C, 0x08, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	var result bool
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_ble_server_get_attr_value() ([]byte, error) {
	if debug {
		fmt.Printf("rpc_ble_server_get_attr_value()\n")
	}
	msg := startWriteMessage(0x00, 0x0C, 0x09, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return nil, err
	}

	<-received
	var result []byte
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_server_exec_write_confirm() (bool, error) {
	if debug {
		fmt.Printf("rpc_server_exec_write_confirm()\n")
	}
	msg := startWriteMessage(0x00, 0x0C, 0x0A, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	var result bool
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_server_attr_write_confirm() (bool, error) {
	if debug {
		fmt.Printf("rpc_server_attr_write_confirm()\n")
	}
	msg := startWriteMessage(0x00, 0x0C, 0x0B, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	var result bool
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_server_attr_read_confirm() (bool, error) {
	if debug {
		fmt.Printf("rpc_server_attr_read_confirm()\n")
	}
	msg := startWriteMessage(0x00, 0x0C, 0x0C, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	var result bool
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_ble_handle_gap_msg() (RPC_T_APP_RESULT, error) {
	if debug {
		fmt.Printf("rpc_ble_handle_gap_msg()\n")
	}
	msg := startWriteMessage(0x00, 0x0D, 0x01, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_APP_RESULT
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_ble_gap_callback() (RPC_T_APP_RESULT, error) {
	if debug {
		fmt.Printf("rpc_ble_gap_callback()\n")
	}
	msg := startWriteMessage(0x00, 0x0D, 0x02, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_APP_RESULT
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_ble_gattc_callback() (RPC_T_APP_RESULT, error) {
	if debug {
		fmt.Printf("rpc_ble_gattc_callback()\n")
	}
	msg := startWriteMessage(0x00, 0x0D, 0x03, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_APP_RESULT
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_ble_gatts_callback() (RPC_T_APP_RESULT, error) {
	if debug {
		fmt.Printf("rpc_ble_gatts_callback()\n")
	}
	msg := startWriteMessage(0x00, 0x0D, 0x04, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result RPC_T_APP_RESULT
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_connect() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_connect()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x01, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_connect_bssid() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_connect_bssid()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x02, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

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
	var result int32
	//err = readInt32(&result)

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
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_is_up() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_is_up()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x05, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_is_ready_to_transceive() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_is_ready_to_transceive()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x06, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_set_mac_address() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_mac_address()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x07, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_get_mac_address() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_mac_address()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x08, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

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
	var result int32
	//err = readInt32(&result)

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
	var result int32
	//err = readInt32(&result)

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
	var result int32
	//err = readInt32(&result)

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
	//err = readInt32(&result)

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
	//err = readInt32(&result)

	seq++
	return err
}

func rpc_wifi_get_associated_client_list() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_associated_client_list()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x0E, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_get_ap_bssid() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_ap_bssid()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x0F, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_get_ap_info() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_ap_info()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x10, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_set_country() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_country()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x11, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_get_sta_max_data_rate() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_sta_max_data_rate()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x12, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_get_rssi() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_rssi()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x13, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_set_channel() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_channel()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x14, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_get_channel() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_channel()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x15, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_change_channel_plan() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_change_channel_plan()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x16, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_register_multicast_address() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_register_multicast_address()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x17, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_unregister_multicast_address() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_unregister_multicast_address()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x18, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

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
	var result int32
	//err = readInt32(&result)

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
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_on() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_on()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x1B, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

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
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_set_mode() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_mode()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x1D, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

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
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_set_power_mode() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_power_mode()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x1F, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_set_tdma_param() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_tdma_param()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x20, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_set_lps_dtim() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_lps_dtim()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x21, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_get_lps_dtim() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_lps_dtim()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x22, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_set_lps_thresh() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_lps_thresh()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x23, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_set_lps_level() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_lps_level()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x24, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_set_mfp_support() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_mfp_support()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x25, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_start_ap() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_start_ap()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x26, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_start_ap_with_hidden_ssid() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_start_ap_with_hidden_ssid()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x27, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_set_pscan_chan() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_pscan_chan()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x28, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_get_setting() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_setting()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x29, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_set_network_mode() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_network_mode()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x2A, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_get_network_mode() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_network_mode()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x2B, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_set_wps_phase() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_wps_phase()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x2C, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_restart_ap() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_restart_ap()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x2D, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_config_autoreconnect() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_config_autoreconnect()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x2E, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_set_autoreconnect() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_autoreconnect()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x2F, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_get_autoreconnect() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_autoreconnect()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x30, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

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
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_add_custom_ie() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_add_custom_ie()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x32, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_update_custom_ie() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_update_custom_ie()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x33, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

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
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_set_indicate_mgnt() error {
	if debug {
		fmt.Printf("rpc_wifi_set_indicate_mgnt()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x35, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received
	//err = readInt32(&result)

	seq++
	return err
}

func rpc_wifi_get_drv_ability() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_drv_ability()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x36, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_set_channel_plan() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_channel_plan()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x37, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_get_channel_plan() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_channel_plan()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x38, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

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
	var result int32
	//err = readInt32(&result)

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
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_set_ch_deauth() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_ch_deauth()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x3B, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

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
	var result uint8
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_set_tx_pause_data() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_set_tx_pause_data()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x3D, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_get_reconnect_data() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_reconnect_data()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x3E, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

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
	var result int32
	//err = readInt32(&result)

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
	var result int32
	//err = readInt32(&result)

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
	var result bool
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_scan_get_ap_records() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_scan_get_ap_records()\n")
	}
	msg := startWriteMessage(0x00, 0x0E, 0x42, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

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
	var result uint16
	//err = readInt32(&result)

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
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcpip_adapter_sta_start() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_sta_start()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x02, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcpip_adapter_ap_start() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_ap_start()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x03, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcpip_adapter_stop() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_stop()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x04, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcpip_adapter_up() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_up()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x05, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcpip_adapter_down() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_down()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x06, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcpip_adapter_get_ip_info() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_get_ip_info()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x07, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcpip_adapter_set_ip_info() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_set_ip_info()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x08, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcpip_adapter_set_dns_info() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_set_dns_info()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x09, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcpip_adapter_get_dns_info() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_get_dns_info()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x0A, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcpip_adapter_dhcps_start() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_dhcps_start()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x0B, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcpip_adapter_dhcps_stop() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_dhcps_stop()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x0C, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcpip_adapter_dhcpc_start() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_dhcpc_start()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x0D, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcpip_adapter_dhcpc_stop() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_dhcpc_stop()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x0E, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcpip_adapter_set_hostname() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_set_hostname()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x0F, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcpip_adapter_get_hostname() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_get_hostname()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x10, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcpip_adapter_get_mac() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_get_mac()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x11, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcpip_adapter_set_mac() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_adapter_set_mac()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x12, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcpip_api_call() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_api_call()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x13, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcp_connect() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_connect()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x14, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcp_recved() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_recved()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x15, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcp_abort() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_abort()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x16, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcp_write() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_write()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x17, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcp_output() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_output()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x18, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcp_close() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_close()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x19, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcp_bind() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_bind()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x1A, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcp_new_ip_type() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_new_ip_type()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x1B, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcp_arg() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_arg()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x1C, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcp_err() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_err()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x1D, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcp_recv() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_recv()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x1E, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcp_sent() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_sent()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x1F, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcp_accept() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_accept()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x20, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcp_poll() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_poll()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x21, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcp_listen_with_backlog() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_listen_with_backlog()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x22, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_pbuf_free() (int32, error) {
	if debug {
		fmt.Printf("rpc_pbuf_free()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x23, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_ip4addr_ntoa() (string, error) {
	if debug {
		fmt.Printf("rpc_ip4addr_ntoa()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x24, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return "", err
	}

	<-received
	var result string
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_inet_chksum() (uint16, error) {
	if debug {
		fmt.Printf("rpc_inet_chksum()\n")
	}
	msg := startWriteMessage(0x00, 0x0F, 0x25, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result uint16
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_lwip_accept() (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_accept()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x01, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_lwip_bind() (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_bind()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x02, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_lwip_shutdown() (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_shutdown()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x03, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_lwip_getpeername() (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_getpeername()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x04, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_lwip_getsockname() (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_getsockname()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x05, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_lwip_getsockopt() (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_getsockopt()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x06, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_lwip_setsockopt() (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_setsockopt()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x07, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_lwip_close() (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_close()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x08, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_lwip_connect() (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_connect()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x09, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_lwip_listen() (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_listen()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x0A, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_lwip_available() (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_available()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x0B, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_lwip_recv() (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_recv()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x0C, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_lwip_read() (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_read()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x0D, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_lwip_recvfrom() (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_recvfrom()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x0E, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_lwip_send() (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_send()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x0F, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_lwip_sendmsg() (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_sendmsg()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x10, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_lwip_sendto() (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_sendto()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x11, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_lwip_socket() (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_socket()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x12, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_lwip_write() (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_write()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x13, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_lwip_writev() (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_writev()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x14, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_lwip_select() (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_select()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x15, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_lwip_ioctl() (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_ioctl()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x16, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_lwip_fcntl() (int32, error) {
	if debug {
		fmt.Printf("rpc_lwip_fcntl()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x17, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

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
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_netconn_gethostbyname() (int8, error) {
	if debug {
		fmt.Printf("rpc_netconn_gethostbyname()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x19, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int8
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_dns_gethostbyname_addrtype() (int8, error) {
	if debug {
		fmt.Printf("rpc_dns_gethostbyname_addrtype()\n")
	}
	msg := startWriteMessage(0x00, 0x10, 0x1A, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int8
	//err = readInt32(&result)

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
	var result uint32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_ssl_client_destroy() error {
	if debug {
		fmt.Printf("rpc_wifi_ssl_client_destroy()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x02, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received
	//err = readInt32(&result)

	seq++
	return err
}

func rpc_wifi_ssl_init() error {
	if debug {
		fmt.Printf("rpc_wifi_ssl_init()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x03, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received
	//err = readInt32(&result)

	seq++
	return err
}

func rpc_wifi_ssl_set_socket() error {
	if debug {
		fmt.Printf("rpc_wifi_ssl_set_socket()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x04, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received
	//err = readInt32(&result)

	seq++
	return err
}

func rpc_wifi_ssl_set_timeout() error {
	if debug {
		fmt.Printf("rpc_wifi_ssl_set_timeout()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x05, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received
	//err = readInt32(&result)

	seq++
	return err
}

func rpc_wifi_ssl_get_socket() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_ssl_get_socket()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x06, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_ssl_get_timeout() (uint32, error) {
	if debug {
		fmt.Printf("rpc_wifi_ssl_get_timeout()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x07, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result uint32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_ssl_set_rootCA() (uint32, error) {
	if debug {
		fmt.Printf("rpc_wifi_ssl_set_rootCA()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x08, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result uint32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_ssl_get_rootCA() (uint32, error) {
	if debug {
		fmt.Printf("rpc_wifi_ssl_get_rootCA()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x09, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result uint32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_ssl_set_cliCert() (uint32, error) {
	if debug {
		fmt.Printf("rpc_wifi_ssl_set_cliCert()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x0A, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result uint32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_ssl_get_cliCert() (uint32, error) {
	if debug {
		fmt.Printf("rpc_wifi_ssl_get_cliCert()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x0B, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result uint32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_ssl_set_cliKey() (uint32, error) {
	if debug {
		fmt.Printf("rpc_wifi_ssl_set_cliKey()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x0C, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result uint32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_ssl_get_cliKey() (uint32, error) {
	if debug {
		fmt.Printf("rpc_wifi_ssl_get_cliKey()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x0D, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result uint32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_ssl_set_pskIdent() (uint32, error) {
	if debug {
		fmt.Printf("rpc_wifi_ssl_set_pskIdent()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x0E, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result uint32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_ssl_get_pskIdent() (uint32, error) {
	if debug {
		fmt.Printf("rpc_wifi_ssl_get_pskIdent()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x0F, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result uint32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_ssl_set_psKey() (uint32, error) {
	if debug {
		fmt.Printf("rpc_wifi_ssl_set_psKey()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x10, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result uint32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_ssl_get_psKey() (uint32, error) {
	if debug {
		fmt.Printf("rpc_wifi_ssl_get_psKey()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x11, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result uint32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_start_ssl_client() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_start_ssl_client()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x12, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_stop_ssl_socket() error {
	if debug {
		fmt.Printf("rpc_wifi_stop_ssl_socket()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x13, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received
	//err = readInt32(&result)

	seq++
	return err
}

func rpc_wifi_data_to_read() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_data_to_read()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x14, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_send_ssl_data() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_send_ssl_data()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x15, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_get_ssl_receive() (int32, error) {
	if debug {
		fmt.Printf("rpc_wifi_get_ssl_receive()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x16, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_verify_ssl_fingerprint() (bool, error) {
	if debug {
		fmt.Printf("rpc_wifi_verify_ssl_fingerprint()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x17, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	var result bool
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_verify_ssl_dn() (bool, error) {
	if debug {
		fmt.Printf("rpc_wifi_verify_ssl_dn()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x18, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return false, err
	}

	<-received
	var result bool
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_ssl_strerror() error {
	if debug {
		fmt.Printf("rpc_wifi_ssl_strerror()\n")
	}
	msg := startWriteMessage(0x00, 0x11, 0x19, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received
	//err = readInt32(&result)

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
	var result int32
	//err = readInt32(&result)

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
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_mdns_service_add() (int32, error) {
	if debug {
		fmt.Printf("rpc_mdns_service_add()\n")
	}
	msg := startWriteMessage(0x00, 0x12, 0x03, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_mdns_service_remove() (int32, error) {
	if debug {
		fmt.Printf("rpc_mdns_service_remove()\n")
	}
	msg := startWriteMessage(0x00, 0x12, 0x04, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_mdns_service_txt_item_set() (int32, error) {
	if debug {
		fmt.Printf("rpc_mdns_service_txt_item_set()\n")
	}
	msg := startWriteMessage(0x00, 0x12, 0x05, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_mdns_service_instance_name_set() (int32, error) {
	if debug {
		fmt.Printf("rpc_mdns_service_instance_name_set()\n")
	}
	msg := startWriteMessage(0x00, 0x12, 0x06, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_mdns_instance_name_set() (int32, error) {
	if debug {
		fmt.Printf("rpc_mdns_instance_name_set()\n")
	}
	msg := startWriteMessage(0x00, 0x12, 0x07, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_mdns_hostname_set() (int32, error) {
	if debug {
		fmt.Printf("rpc_mdns_hostname_set()\n")
	}
	msg := startWriteMessage(0x00, 0x12, 0x08, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_mdns_query_a() (int32, error) {
	if debug {
		fmt.Printf("rpc_mdns_query_a()\n")
	}
	msg := startWriteMessage(0x00, 0x12, 0x09, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_mdns_query_ptr() (int32, error) {
	if debug {
		fmt.Printf("rpc_mdns_query_ptr()\n")
	}
	msg := startWriteMessage(0x00, 0x12, 0x0A, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_mdns_query_ptr_result_basic() (int32, error) {
	if debug {
		fmt.Printf("rpc_mdns_query_ptr_result_basic()\n")
	}
	msg := startWriteMessage(0x00, 0x12, 0x0B, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_mdns_query_ptr_result_txt() (int32, error) {
	if debug {
		fmt.Printf("rpc_mdns_query_ptr_result_txt()\n")
	}
	msg := startWriteMessage(0x00, 0x12, 0x0C, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_mdns_query_ptr_result_addr() (int32, error) {
	if debug {
		fmt.Printf("rpc_mdns_query_ptr_result_addr()\n")
	}
	msg := startWriteMessage(0x00, 0x12, 0x0D, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

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
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_wifi_event_callback() error {
	if debug {
		fmt.Printf("rpc_wifi_event_callback()\n")
	}
	msg := startWriteMessage(0x00, 0x13, 0x01, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received
	//err = readInt32(&result)

	seq++
	return err
}

func rpc_wifi_dns_found() error {
	if debug {
		fmt.Printf("rpc_wifi_dns_found()\n")
	}
	msg := startWriteMessage(0x00, 0x13, 0x02, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return err
	}

	<-received
	//err = readInt32(&result)

	seq++
	return err
}

func rpc_tcpip_api_call_fn() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcpip_api_call_fn()\n")
	}
	msg := startWriteMessage(0x00, 0x13, 0x03, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcp_connected_fn() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_connected_fn()\n")
	}
	msg := startWriteMessage(0x00, 0x13, 0x04, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcp_recv_fn() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_recv_fn()\n")
	}
	msg := startWriteMessage(0x00, 0x13, 0x05, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcp_accept_fn() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_accept_fn()\n")
	}
	msg := startWriteMessage(0x00, 0x13, 0x06, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcp_err_fn() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_err_fn()\n")
	}
	msg := startWriteMessage(0x00, 0x13, 0x07, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcp_sent_fn() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_sent_fn()\n")
	}
	msg := startWriteMessage(0x00, 0x13, 0x08, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}

func rpc_tcp_poll_fn() (int32, error) {
	if debug {
		fmt.Printf("rpc_tcp_poll_fn()\n")
	}
	msg := startWriteMessage(0x00, 0x13, 0x09, uint32(seq))
	err := performRequest(msg)
	if err != nil {
		return 0, err
	}

	<-received
	var result int32
	//err = readInt32(&result)

	seq++
	return result, err
}
