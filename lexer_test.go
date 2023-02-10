package main

import (
	"testing"
)

func TestLexer(t *testing.T) {
	l := NewLexer(`
@output_dir("erpc_shim") 
@types_header("rpc_system_header.h")

program rpc
@group("system")

interface rpc_system {
    rpc_system_version() -> string
    rpc_system_ack(uint8 c) -> uint8
}

import "rpc_ble_common.erpc"

/** @brief Event type to inform app*/
enum RPC_T_SERVICE_CALLBACK_TYPE
{
    RPC_SERVICE_CALLBACK_TYPE_INDIFICATION_NOTIFICATION = 1,    /**< CCCD update event */
};

enum RPC_T_GAP_CAUSE
{
    RPC_GAP_CAUSE_ERROR_UNKNOWN     = 0xFF//!< Unknown error.
};

interface rpc_wifi_lwip{
    rpc_lwip_accept(int32 s, in binary addr, inout uint32 addrlen) -> int32
    rpc_lwip_getpeername (int32 s, out binary name, inout uint32 namelen) -> int32
    rpc_wifi_send_ssl_data(uint32 ssl_client, in binary data, uint16 len) -> int32
    rpc_func(uint8 u8, uint16 u16, uint32 u32, int8 i8, int16 i16, int32 i32, bool b) -> void
    rpc_wifi_ssl_set_rootCA(string rootCABuff @retain) -> uint32
    rpc_wifi_ssl_get_rootCA(out string rootCABuff @nullable @max_length(3092)) -> uint32
    rpc_wifi_get_mac_address(out uint8[18] mac) -> int32
}

struct RPC_T_LOCAL_APPEARANCE
{
    uint16 local_appearance;
    uint8[2]  padding;
};
    `)

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{TOK_AT, "@"},
		{TOK_IDENT, "output_dir"},
		{TOK_LPAREN, "("},
		{TOK_STRING, `erpc_shim`},
		{TOK_RPAREN, ")"},

		{TOK_AT, "@"},
		{TOK_IDENT, "types_header"},
		{TOK_LPAREN, "("},
		{TOK_STRING, `rpc_system_header.h`},
		{TOK_RPAREN, ")"},

		{TOK_PROGRAM, "program"},
		{TOK_IDENT, "rpc"},

		{TOK_AT, "@"},
		{TOK_IDENT, "group"},
		{TOK_LPAREN, "("},
		{TOK_STRING, `system`},
		{TOK_RPAREN, ")"},

		{TOK_INTERFACE, "interface"},
		{TOK_IDENT, "rpc_system"},
		{TOK_LBRACE, "{"},

		{TOK_IDENT, "rpc_system_version"},
		{TOK_LPAREN, "("},
		{TOK_RPAREN, ")"},
		{TOK_ARROW, "->"},
		{TOK_TYPE, "string"},

		{TOK_IDENT, "rpc_system_ack"},
		{TOK_LPAREN, "("},
		{TOK_TYPE, "uint8"},
		{TOK_IDENT, "c"},
		{TOK_RPAREN, ")"},
		{TOK_ARROW, "->"},
		{TOK_TYPE, "uint8"},

		{TOK_RBRACE, "}"},

		{TOK_IMPORT, "import"},
		{TOK_STRING, "rpc_ble_common.erpc"},

		{TOK_COMMENT, "/** @brief Event type to inform app*/"},
		{TOK_ENUM, "enum"},
		{TOK_IDENT, "RPC_T_SERVICE_CALLBACK_TYPE"},
		{TOK_LBRACE, "{"},
		{TOK_IDENT, "RPC_SERVICE_CALLBACK_TYPE_INDIFICATION_NOTIFICATION"},
		{TOK_ASSIGN, "="},
		{TOK_NUM, "1"},
		{TOK_COMMA, ","},
		{TOK_COMMENT, "/**< CCCD update event */"},
		{TOK_RBRACE, "}"},
		{TOK_SEMICOLON, ";"},

		{TOK_ENUM, "enum"},
		{TOK_IDENT, "RPC_T_GAP_CAUSE"},
		{TOK_LBRACE, "{"},
		{TOK_IDENT, "RPC_GAP_CAUSE_ERROR_UNKNOWN"},
		{TOK_ASSIGN, "="},
		{TOK_NUM, "0xFF"},
		{TOK_COMMENT, "//!< Unknown error."},
		{TOK_RBRACE, "}"},
		{TOK_SEMICOLON, ";"},

		{TOK_INTERFACE, "interface"},
		{TOK_IDENT, "rpc_wifi_lwip"},
		{TOK_LBRACE, "{"},

		{TOK_IDENT, "rpc_lwip_accept"},
		{TOK_LPAREN, "("},
		{TOK_TYPE, "int32"},
		{TOK_IDENT, "s"},
		{TOK_COMMA, ","},
		{TOK_TYPE, "in"},
		{TOK_TYPE, "binary"},
		{TOK_IDENT, "addr"},
		{TOK_COMMA, ","},
		{TOK_TYPE, "inout"},
		{TOK_TYPE, "uint32"},
		{TOK_IDENT, "addrlen"},
		{TOK_RPAREN, ")"},
		{TOK_ARROW, "->"},
		{TOK_TYPE, "int32"},

		{TOK_IDENT, "rpc_lwip_getpeername"},
		{TOK_LPAREN, "("},
		{TOK_TYPE, "int32"},
		{TOK_IDENT, "s"},
		{TOK_COMMA, ","},
		{TOK_TYPE, "out"},
		{TOK_TYPE, "binary"},
		{TOK_IDENT, "name"},
		{TOK_COMMA, ","},
		{TOK_TYPE, "inout"},
		{TOK_TYPE, "uint32"},
		{TOK_IDENT, "namelen"},
		{TOK_RPAREN, ")"},
		{TOK_ARROW, "->"},
		{TOK_TYPE, "int32"},

		{TOK_IDENT, "rpc_wifi_send_ssl_data"},
		{TOK_LPAREN, "("},
		{TOK_TYPE, "uint32"},
		{TOK_IDENT, "ssl_client"},
		{TOK_COMMA, ","},
		{TOK_TYPE, "in"},
		{TOK_TYPE, "binary"},
		{TOK_IDENT, "data"},
		{TOK_COMMA, ","},
		{TOK_TYPE, "uint16"},
		{TOK_IDENT, "len"},
		{TOK_RPAREN, ")"},
		{TOK_ARROW, "->"},
		{TOK_TYPE, "int32"},

		{TOK_IDENT, "rpc_func"},
		{TOK_LPAREN, "("},
		{TOK_TYPE, "uint8"},
		{TOK_IDENT, "u8"},
		{TOK_COMMA, ","},
		{TOK_TYPE, "uint16"},
		{TOK_IDENT, "u16"},
		{TOK_COMMA, ","},
		{TOK_TYPE, "uint32"},
		{TOK_IDENT, "u32"},
		{TOK_COMMA, ","},
		{TOK_TYPE, "int8"},
		{TOK_IDENT, "i8"},
		{TOK_COMMA, ","},
		{TOK_TYPE, "int16"},
		{TOK_IDENT, "i16"},
		{TOK_COMMA, ","},
		{TOK_TYPE, "int32"},
		{TOK_IDENT, "i32"},
		{TOK_COMMA, ","},
		{TOK_TYPE, "bool"},
		{TOK_IDENT, "b"},
		{TOK_RPAREN, ")"},
		{TOK_ARROW, "->"},
		{TOK_TYPE, "void"},

		{TOK_IDENT, "rpc_wifi_ssl_set_rootCA"},
		{TOK_LPAREN, "("},
		{TOK_TYPE, "string"},
		{TOK_IDENT, "rootCABuff"},
		{TOK_AT, "@"},
		{TOK_IDENT, "retain"},
		{TOK_RPAREN, ")"},
		{TOK_ARROW, "->"},
		{TOK_TYPE, "uint32"},

		{TOK_IDENT, "rpc_wifi_ssl_get_rootCA"},
		{TOK_LPAREN, "("},
		{TOK_TYPE, "out"},
		{TOK_TYPE, "string"},
		{TOK_IDENT, "rootCABuff"},
		{TOK_AT, "@"},
		{TOK_IDENT, "nullable"},
		{TOK_AT, "@"},
		{TOK_IDENT, "max_length"},
		{TOK_LPAREN, "("},
		{TOK_NUM, "3092"},
		{TOK_RPAREN, ")"},
		{TOK_RPAREN, ")"},
		{TOK_ARROW, "->"},
		{TOK_TYPE, "uint32"},

		{TOK_IDENT, "rpc_wifi_get_mac_address"},
		{TOK_LPAREN, "("},
		{TOK_TYPE, "out"},
		{TOK_TYPE, "uint8"},
		{TOK_LBRACKET, "["},
		{TOK_NUM, "18"},
		{TOK_RBRACKET, "]"},
		{TOK_IDENT, "mac"},
		{TOK_RPAREN, ")"},
		{TOK_ARROW, "->"},
		{TOK_TYPE, "int32"},

		{TOK_RBRACE, "}"},

		{TOK_STRUCT, "struct"},
		{TOK_IDENT, "RPC_T_LOCAL_APPEARANCE"},
		{TOK_LBRACE, "{"},
		{TOK_TYPE, "uint16"},
		{TOK_IDENT, "local_appearance"},
		{TOK_SEMICOLON, ";"},
		{TOK_TYPE, "uint8"},
		{TOK_LBRACKET, "["},
		{TOK_NUM, "2"},
		{TOK_RBRACKET, "]"},
		{TOK_IDENT, "padding"},
		{TOK_SEMICOLON, ";"},
		{TOK_RBRACE, "}"},
	}

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Errorf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Errorf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}

}
