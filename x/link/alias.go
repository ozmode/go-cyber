package link

import (
	"github.com/cybercongress/go-cyber/x/link/exported"
	"github.com/cybercongress/go-cyber/x/link/internal/keeper"
	"github.com/cybercongress/go-cyber/x/link/internal/types"
)

const (
	ModuleName = types.ModuleName
	StoreKey   = types.StoreKey
	RouterKey  = types.RouterKey
)

var (
	// keeper
	NewLinkKeeper      = keeper.NewLinkKeeper
	NewIndexedKeeper   = keeper.NewIndexedKeeper
	NewCidNumberKeeper = keeper.NewCidNumberKeeper

	// types
	RegisterCodec = types.RegisterCodec
	NewMsg        = types.NewMsg
	NewLink       = types.NewLink

	ErrCyberlinkExist = types.ErrCyberlinkExist
	ErrInvalidCid 	  = types.ErrInvalidCid
	ErrZeroLinks      = types.ErrZeroLinks
)

type (
	// exported
    Keeper          = exported.KeeperI
    IndexedKeeper   = exported.IndexedKeeperI
    CidNumberKeeper = exported.CidNumberKeeperI

	// types
	Msg         = types.Msg
	Links       = types.Links
	CidLinks    = types.CidLinks
	Cid         = types.Cid
	CidNumber   = types.CidNumber
	CidsFilter  = types.CidsFilter
	Link        = types.Link
	CompactLink = types.CompactLink
	LinkFilter  = types.LinkFilter

)
