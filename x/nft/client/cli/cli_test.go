package cli_test

import (
	"fmt"
	"testing"

	"github.com/cosmos/gogoproto/proto"
	"github.com/stretchr/testify/suite"
	"github.com/tidwall/gjson"

	"github.com/cometbft/cometbft/crypto"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/CudoVentures/cudos-node/simapp"
	"github.com/CudoVentures/cudos-node/testutil"
	nftcli "github.com/CudoVentures/cudos-node/x/nft/client/cli"
	nfttestutil "github.com/CudoVentures/cudos-node/x/nft/client/testutil"
	nftKeeper "github.com/CudoVentures/cudos-node/x/nft/keeper"
	nfttypes "github.com/CudoVentures/cudos-node/x/nft/types"
)

type IntegrationTestSuite struct {
	suite.Suite
	keeper  nftKeeper.Keeper
	cfg     network.Config
	network *network.Network
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	cfg := simapp.NewConfig(s.T().TempDir())
	cfg.NumValidators = 1

	s.cfg = cfg
	network, err := network.New(s.T(), s.T().TempDir(), cfg)
	s.Require().NoError(err)
	s.network = network

	_, err = s.network.WaitForHeight(3) // The network is fully initialized after 3 blocks
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) TestNft() {
	val := s.network.Validators[0]

	// ---------------------------------------------------------------------------

	from := val.Address
	tokenName := "Test Token"
	tokenURI := "uri"
	tokenData := "data"
	denomName := "name"
	denomSymbol := "symbol"
	denom := "denom"
	schema := "schema"

	gasArgs := []string{
		fmt.Sprintf("--%s=%s", flags.FlagGas, flags.GasFlagAuto),
		fmt.Sprintf("--%s=1.3", flags.FlagGasAdjustment),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(1))).String()),
	}

	//------test GetCmdIssueDenom()-------------
	args := []string{
		fmt.Sprintf("--%s=%s", nftcli.FlagDenomName, denomName),
		fmt.Sprintf("--%s=%s", nftcli.FlagDenomSymbol, denomSymbol),
		fmt.Sprintf("--%s=%s", nftcli.FlagSchema, schema),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
	}

	args = append(args, gasArgs...)

	expectedCode := uint32(0)

	bz, err := nfttestutil.IssueDenomExec(val.ClientCtx, from.String(), denom, args...)
	s.Require().NoError(err)
	testutil.WaitForBlock()
	txResp, err := testutil.QueryJustBroadcastedTx(val.ClientCtx, bz)
	s.Require().NoError(err)
	s.Require().Equal(expectedCode, txResp.Code)
	denomID := gjson.Get(txResp.RawLog, "0.events.#(type==\"issue_denom\").attributes.#(key==\"denom_id\").value").String()

	//------test GetCmdQueryDenom()-------------
	respType := proto.Message(&nfttypes.Denom{})
	bz, err = nfttestutil.QueryDenomExec(val.ClientCtx, denomID)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType))
	denomItem := respType.(*nfttypes.Denom)
	s.Require().Equal(denomName, denomItem.Name)
	s.Require().Equal(denomSymbol, denomItem.Symbol)
	s.Require().Equal(schema, denomItem.Schema)

	//------test GetCmdQueryDenoms()-------------
	respType = proto.Message(&nfttypes.QueryDenomsResponse{})
	bz, err = nfttestutil.QueryDenomsExec(val.ClientCtx)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType))
	denomsResp := respType.(*nfttypes.QueryDenomsResponse)
	s.Require().Equal(1, len(denomsResp.Denoms))
	s.Require().Equal(denomID, denomsResp.Denoms[0].Id)

	//------test GetCmdMintNFT()-------------
	args = []string{
		fmt.Sprintf("--%s=%s", nftcli.FlagTokenData, tokenData),
		fmt.Sprintf("--%s=%s", nftcli.FlagRecipient, from.String()),
		fmt.Sprintf("--%s=%s", nftcli.FlagTokenURI, tokenURI),
		fmt.Sprintf("--%s=%s", nftcli.FlagTokenName, tokenName),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
	}
	args = append(args, gasArgs...)

	bz, err = nfttestutil.MintNFTExec(val.ClientCtx, from.String(), denomID, args...)
	s.Require().NoError(err)
	testutil.WaitForBlock()
	txResp, err = testutil.QueryJustBroadcastedTx(val.ClientCtx, bz)
	s.Require().NoError(err)
	s.Require().Equal(expectedCode, txResp.Code)
	tokenID := gjson.Get(txResp.RawLog, "0.events.#(type==\"mint_nft\").attributes.#(key==\"token_id\").value").String()

	//------test GetCmdQuerySupply()-------------
	respType = proto.Message(&nfttypes.QuerySupplyResponse{})
	bz, err = nfttestutil.QuerySupplyExec(val.ClientCtx, denomID)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType))
	supplyResp := respType.(*nfttypes.QuerySupplyResponse)
	s.Require().Equal(uint64(1), supplyResp.Amount)

	//------test GetCmdQueryNFT()-------------
	respType = proto.Message(&nfttypes.BaseNFT{})
	bz, err = nfttestutil.QueryNFTExec(val.ClientCtx, denomID, tokenID)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType))
	nftItem := respType.(*nfttypes.BaseNFT)
	s.Require().Equal(tokenID, nftItem.Id)
	s.Require().Equal(tokenName, nftItem.Name)
	s.Require().Equal(tokenURI, nftItem.URI)
	s.Require().Equal(tokenData, nftItem.Data)
	s.Require().Equal(from.String(), nftItem.Owner)

	//------test GetCmdQueryOwner()-------------
	respType = proto.Message(&nfttypes.QueryOwnerResponse{})
	bz, err = nfttestutil.QueryOwnerExec(val.ClientCtx, from.String())
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType))
	ownerResp := respType.(*nfttypes.QueryOwnerResponse)
	s.Require().Equal(from.String(), ownerResp.Owner.Address)
	s.Require().Equal(denom, ownerResp.Owner.IDCollections[0].DenomId)
	s.Require().Equal(tokenID, ownerResp.Owner.IDCollections[0].TokenIds[0])

	//------test GetCmdQueryCollection()-------------
	respType = proto.Message(&nfttypes.QueryCollectionResponse{})
	bz, err = nfttestutil.QueryCollectionExec(val.ClientCtx, denomID)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType))
	collectionItem := respType.(*nfttypes.QueryCollectionResponse)
	s.Require().Equal(1, len(collectionItem.Collection.NFTs))

	//------test GetCmdEditNFT()-------------
	newTokenData := "newdata"
	newTokenURI := "newuri"
	newTokenName := "new Test Token"
	args = []string{
		fmt.Sprintf("--%s=%s", nftcli.FlagTokenData, newTokenData),
		fmt.Sprintf("--%s=%s", nftcli.FlagTokenURI, newTokenURI),
		fmt.Sprintf("--%s=%s", nftcli.FlagTokenName, newTokenName),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
	}

	args = append(args, gasArgs...)

	bz, err = nfttestutil.EditNFTExec(val.ClientCtx, from.String(), denomID, tokenID, args...)
	s.Require().NoError(err)
	testutil.WaitForBlock()
	txResp, err = testutil.QueryJustBroadcastedTx(val.ClientCtx, bz)
	s.Require().NoError(err)
	s.Require().Equal(expectedCode, txResp.Code)

	respType = proto.Message(&nfttypes.BaseNFT{})
	bz, err = nfttestutil.QueryNFTExec(val.ClientCtx, denomID, tokenID)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType))
	newNftItem := respType.(*nfttypes.BaseNFT)
	s.Require().Equal(newTokenName, newNftItem.Name)
	s.Require().Equal(newTokenURI, newNftItem.URI)
	s.Require().Equal(newTokenData, newNftItem.Data)

	//------test GetCmdTransferNFT()-------------
	recipient := sdk.AccAddress(crypto.AddressHash([]byte("dgsbl")))

	args = []string{
		fmt.Sprintf("--%s=%s", nftcli.FlagTokenData, newTokenData),
		fmt.Sprintf("--%s=%s", nftcli.FlagTokenURI, newTokenURI),
		fmt.Sprintf("--%s=%s", nftcli.FlagTokenName, newTokenName),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
	}

	args = append(args, gasArgs...)

	bz, err = nfttestutil.TransferNFTExec(val.ClientCtx, from.String(), recipient.String(), denomID, tokenID, args...)
	s.Require().NoError(err)
	testutil.WaitForBlock()
	txResp, err = testutil.QueryJustBroadcastedTx(val.ClientCtx, bz)
	s.Require().NoError(err)
	s.Require().Equal(expectedCode, txResp.Code)

	respType = proto.Message(&nfttypes.BaseNFT{})
	bz, err = nfttestutil.QueryNFTExec(val.ClientCtx, denomID, tokenID)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType))
	nftItem = respType.(*nfttypes.BaseNFT)
	s.Require().Equal(tokenID, nftItem.Id)
	s.Require().Equal(newTokenName, nftItem.Name)
	s.Require().Equal(newTokenURI, nftItem.URI)
	s.Require().Equal(newTokenData, nftItem.Data)
	s.Require().Equal(recipient.String(), nftItem.Owner)

	//------test GetCmdApproveNFT() GetCmdQueryApproveNFT() -------------

	approvedAddress := sdk.AccAddress(crypto.AddressHash([]byte("dgsbl")))

	// mint new NFT
	args = []string{
		fmt.Sprintf("--%s=%s", nftcli.FlagTokenData, tokenData),
		fmt.Sprintf("--%s=%s", nftcli.FlagRecipient, from.String()),
		fmt.Sprintf("--%s=%s", nftcli.FlagTokenURI, tokenURI),
		fmt.Sprintf("--%s=%s", nftcli.FlagTokenName, tokenName),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
	}

	args = append(args, gasArgs...)

	bz, err = nfttestutil.MintNFTExec(val.ClientCtx, from.String(), denomID, args...)
	s.Require().NoError(err)
	testutil.WaitForBlock()
	txResp, err = testutil.QueryJustBroadcastedTx(val.ClientCtx, bz)
	s.Require().NoError(err)
	s.Require().Equal(expectedCode, txResp.Code)
	tokenID2 := gjson.Get(txResp.RawLog, "0.events.#(type==\"mint_nft\").attributes.#(key==\"token_id\").value").String()

	args = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
	}

	args = append(args, gasArgs...)

	bz, err = nfttestutil.ApproveNFTExec(val.ClientCtx, from.String(), approvedAddress.String(), denomID, tokenID2, args...)
	s.Require().NoError(err)
	testutil.WaitForBlock()
	txResp, err = testutil.QueryJustBroadcastedTx(val.ClientCtx, bz)
	s.Require().NoError(err)
	s.Require().Equal(expectedCode, txResp.Code)

	respType = proto.Message(&nfttypes.BaseNFT{})
	bz, err = nfttestutil.QueryNFTExec(val.ClientCtx, denomID, tokenID2)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType))
	nftItem = respType.(*nfttypes.BaseNFT)
	s.Require().Equal(s.keeper.IsApprovedAddress(nftItem, approvedAddress.String()), true)

	respType = proto.Message(&nfttypes.QueryApprovalsNFTResponse{})
	bz, err = nfttestutil.QueryIsApprovedNFT(val.ClientCtx, denomID, tokenID2)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType))
	isApprovedNft := respType.(*nfttypes.QueryApprovalsNFTResponse)
	s.Assert().Contains(isApprovedNft.ApprovedAddresses, approvedAddress.String())

	//------test GetCmdApproveAll  GetCmdQueryApproveAll-------------

	args = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
	}

	args = append(args, gasArgs...)

	// respType = proto.Message(&sdk.TxResponse{})

	bz, err = nfttestutil.ApproveAll(val.ClientCtx, from.String(), approvedAddress.String(), "true", args...)
	s.Require().NoError(err)
	testutil.WaitForBlock()
	txResp, err = testutil.QueryJustBroadcastedTx(val.ClientCtx, bz)
	s.Require().NoError(err)
	s.Require().Equal(expectedCode, txResp.Code)

	respType = proto.Message(&nfttypes.QueryApprovalsIsApprovedForAllResponse{})
	bz, err = nfttestutil.QueryIsApprovedAll(val.ClientCtx, from.String(), approvedAddress.String())
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType))
	isApprovedAllResponse := respType.(*nfttypes.QueryApprovalsIsApprovedForAllResponse)
	s.Require().Equal(isApprovedAllResponse.IsApproved, true)

	//------test GetCmdBurnNFT()-------------
	args = []string{
		fmt.Sprintf("--%s=%s", nftcli.FlagTokenData, newTokenData),
		fmt.Sprintf("--%s=%s", nftcli.FlagRecipient, from.String()),
		fmt.Sprintf("--%s=%s", nftcli.FlagTokenURI, newTokenURI),
		fmt.Sprintf("--%s=%s", nftcli.FlagTokenName, newTokenName),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
	}

	args = append(args, gasArgs...)

	bz, err = nfttestutil.MintNFTExec(val.ClientCtx, from.String(), denomID, args...)
	s.Require().NoError(err)
	testutil.WaitForBlock()
	txResp, err = testutil.QueryJustBroadcastedTx(val.ClientCtx, bz)
	s.Require().NoError(err)
	s.Require().Equal(expectedCode, txResp.Code)
	newTokenId := gjson.Get(txResp.RawLog, "0.events.#(type==\"mint_nft\").attributes.#(key==\"token_id\").value").String()

	respType = proto.Message(&nfttypes.QuerySupplyResponse{})
	bz, err = nfttestutil.QuerySupplyExec(val.ClientCtx, denomID)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType))
	supplyResp = respType.(*nfttypes.QuerySupplyResponse)
	s.Require().Equal(uint64(3), supplyResp.Amount)

	args = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
	}

	args = append(args, gasArgs...)

	bz, err = nfttestutil.BurnNFTExec(val.ClientCtx, from.String(), denomID, newTokenId, args...)
	s.Require().NoError(err)
	testutil.WaitForBlock()
	txResp, err = testutil.QueryJustBroadcastedTx(val.ClientCtx, bz)
	s.Require().NoError(err)
	s.Require().Equal(expectedCode, txResp.Code)

	respType = proto.Message(&nfttypes.QuerySupplyResponse{})
	bz, err = nfttestutil.QuerySupplyExec(val.ClientCtx, denomID)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType))
	supplyResp = respType.(*nfttypes.QuerySupplyResponse)
	s.Require().Equal(uint64(2), supplyResp.Amount)

	//------test GetCmdTransferDenom()-------------
	denomRecipient := sdk.AccAddress(crypto.AddressHash([]byte("dgsbl")))

	args = []string{
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	bz, err = nfttestutil.TransferDenomExec(val.ClientCtx, from.String(), denomRecipient.String(), denomID, args...)
	s.Require().NoError(err)
	testutil.WaitForBlock()
	txResp, err = testutil.QueryJustBroadcastedTx(val.ClientCtx, bz)
	s.Require().NoError(err)
	s.Require().Equal(expectedCode, txResp.Code)

	respType = proto.Message(&nfttypes.Denom{})
	bz, err = nfttestutil.QueryDenomExec(val.ClientCtx, denomID)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType))
	denomItem2 := respType.(*nfttypes.Denom)
	s.Require().Equal(denomRecipient.String(), denomItem2.Creator)
	s.Require().Equal(denomName, denomItem2.Name)
	s.Require().Equal(schema, denomItem2.Schema)
	s.Require().Equal(denomSymbol, denomItem2.Symbol)
}
