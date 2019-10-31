package tezosprotocol

const (
	// PrimK_parameter and the remaining constants are adapted from Adapted from
	// https://gitlab.com/tezos/tezos/blob/master/src%2Fproto_alpha%2Flib_protocol%2Fmichelson_v1_primitives.ml
	// They don't follow go's name conventions in order to preserve easy comparison
	// with the original ocaml code and easy updates to keep up with protocol changes
	PrimK_parameter = iota
	PrimK_storage
	PrimK_code
	PrimD_False
	PrimD_Elt
	PrimD_Left
	PrimD_None
	PrimD_Pair
	PrimD_Right
	PrimD_Some
	PrimD_True
	PrimD_Unit
	PrimI_PACK
	PrimI_UNPACK
	PrimI_BLAKE2B
	PrimI_SHA256
	PrimI_SHA512
	PrimI_ABS
	PrimI_ADD
	PrimI_AMOUNT
	PrimI_AND
	PrimI_BALANCE
	PrimI_CAR
	PrimI_CDR
	PrimI_CHAIN_ID
	PrimI_CHECK_SIGNATURE
	PrimI_COMPARE
	PrimI_CONCAT
	PrimI_CONS
	PrimI_CREATE_ACCOUNT
	PrimI_CREATE_CONTRACT
	PrimI_IMPLICIT_ACCOUNT
	PrimI_DIP
	PrimI_DROP
	PrimI_DUP
	PrimI_EDIV
	PrimI_EMPTY_BIG_MAP
	PrimI_EMPTY_MAP
	PrimI_EMPTY_SET
	PrimI_EQ
	PrimI_EXEC
	PrimI_APPLY
	PrimI_FAILWITH
	PrimI_GE
	PrimI_GET
	PrimI_GT
	PrimI_HASH_KEY
	PrimI_IF
	PrimI_IF_CONS
	PrimI_IF_LEFT
	PrimI_IF_NONE
	PrimI_INT
	PrimI_LAMBDA
	PrimI_LE
	PrimI_LEFT
	PrimI_LOOP
	PrimI_LSL
	PrimI_LSR
	PrimI_LT
	PrimI_MAP
	PrimI_MEM
	PrimI_MUL
	PrimI_NEG
	PrimI_NEQ
	PrimI_NIL
	PrimI_NONE
	PrimI_NOT
	PrimI_NOW
	PrimI_OR
	PrimI_PAIR
	PrimI_PUSH
	PrimI_RIGHT
	PrimI_SIZE
	PrimI_SOME
	PrimI_SOURCE
	PrimI_SENDER
	PrimI_SELF
	PrimI_SLICE
	PrimI_STEPS_TO_QUOTA
	PrimI_SUB
	PrimI_SWAP
	PrimI_TRANSFER_TOKENS
	PrimI_SET_DELEGATE
	PrimI_UNIT
	PrimI_UPDATE
	PrimI_XOR
	PrimI_ITER
	PrimI_LOOP_LEFT
	PrimI_ADDRESS
	PrimI_CONTRACT
	PrimI_ISNAT
	PrimI_CAST
	PrimI_RENAME
	PrimI_DIG
	PrimI_DUG
	PrimT_bool
	PrimT_contract
	PrimT_int
	PrimT_key
	PrimT_key_hash
	PrimT_lambda
	PrimT_list
	PrimT_map
	PrimT_big_map
	PrimT_nat
	PrimT_option
	PrimT_or
	PrimT_pair
	PrimT_set
	PrimT_signature
	PrimT_string
	PrimT_bytes
	PrimT_mutez
	PrimT_timestamp
	PrimT_unit
	PrimT_operation
	PrimT_address
	PrimT_chain_id
)
