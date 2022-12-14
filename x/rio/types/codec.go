package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateCert{}, "rio/CreateCert", nil)
	cdc.RegisterConcrete(&MsgSendCert{}, "rio/SendCert", nil)
	cdc.RegisterConcrete(&MsgCreateResume{}, "rio/CreateResume", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateCert{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSendCert{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateResume{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
