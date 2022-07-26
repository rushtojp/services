// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/wallet.proto

package wallet

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Wallet service

func NewWalletEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Wallet service

type WalletService interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Credit(ctx context.Context, in *CreditRequest, opts ...client.CallOption) (*CreditResponse, error)
	Debit(ctx context.Context, in *DebitRequest, opts ...client.CallOption) (*DebitResponse, error)
	Balance(ctx context.Context, in *BalanceRequest, opts ...client.CallOption) (*BalanceResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	Transactions(ctx context.Context, in *TransactionsRequest, opts ...client.CallOption) (*TransactionsResponse, error)
	Transfer(ctx context.Context, in *TransferRequest, opts ...client.CallOption) (*TransferResponse, error)
}

type walletService struct {
	c    client.Client
	name string
}

func NewWalletService(name string, c client.Client) WalletService {
	return &walletService{
		c:    c,
		name: name,
	}
}

func (c *walletService) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.name, "Wallet.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletService) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.name, "Wallet.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletService) Credit(ctx context.Context, in *CreditRequest, opts ...client.CallOption) (*CreditResponse, error) {
	req := c.c.NewRequest(c.name, "Wallet.Credit", in)
	out := new(CreditResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletService) Debit(ctx context.Context, in *DebitRequest, opts ...client.CallOption) (*DebitResponse, error) {
	req := c.c.NewRequest(c.name, "Wallet.Debit", in)
	out := new(DebitResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletService) Balance(ctx context.Context, in *BalanceRequest, opts ...client.CallOption) (*BalanceResponse, error) {
	req := c.c.NewRequest(c.name, "Wallet.Balance", in)
	out := new(BalanceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "Wallet.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletService) Transactions(ctx context.Context, in *TransactionsRequest, opts ...client.CallOption) (*TransactionsResponse, error) {
	req := c.c.NewRequest(c.name, "Wallet.Transactions", in)
	out := new(TransactionsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *walletService) Transfer(ctx context.Context, in *TransferRequest, opts ...client.CallOption) (*TransferResponse, error) {
	req := c.c.NewRequest(c.name, "Wallet.Transfer", in)
	out := new(TransferResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Wallet service

type WalletHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Credit(context.Context, *CreditRequest, *CreditResponse) error
	Debit(context.Context, *DebitRequest, *DebitResponse) error
	Balance(context.Context, *BalanceRequest, *BalanceResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
	Transactions(context.Context, *TransactionsRequest, *TransactionsResponse) error
	Transfer(context.Context, *TransferRequest, *TransferResponse) error
}

func RegisterWalletHandler(s server.Server, hdlr WalletHandler, opts ...server.HandlerOption) error {
	type wallet interface {
		Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error
		Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error
		Credit(ctx context.Context, in *CreditRequest, out *CreditResponse) error
		Debit(ctx context.Context, in *DebitRequest, out *DebitResponse) error
		Balance(ctx context.Context, in *BalanceRequest, out *BalanceResponse) error
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
		Transactions(ctx context.Context, in *TransactionsRequest, out *TransactionsResponse) error
		Transfer(ctx context.Context, in *TransferRequest, out *TransferResponse) error
	}
	type Wallet struct {
		wallet
	}
	h := &walletHandler{hdlr}
	return s.Handle(s.NewHandler(&Wallet{h}, opts...))
}

type walletHandler struct {
	WalletHandler
}

func (h *walletHandler) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.WalletHandler.Create(ctx, in, out)
}

func (h *walletHandler) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.WalletHandler.Delete(ctx, in, out)
}

func (h *walletHandler) Credit(ctx context.Context, in *CreditRequest, out *CreditResponse) error {
	return h.WalletHandler.Credit(ctx, in, out)
}

func (h *walletHandler) Debit(ctx context.Context, in *DebitRequest, out *DebitResponse) error {
	return h.WalletHandler.Debit(ctx, in, out)
}

func (h *walletHandler) Balance(ctx context.Context, in *BalanceRequest, out *BalanceResponse) error {
	return h.WalletHandler.Balance(ctx, in, out)
}

func (h *walletHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.WalletHandler.List(ctx, in, out)
}

func (h *walletHandler) Transactions(ctx context.Context, in *TransactionsRequest, out *TransactionsResponse) error {
	return h.WalletHandler.Transactions(ctx, in, out)
}

func (h *walletHandler) Transfer(ctx context.Context, in *TransferRequest, out *TransferResponse) error {
	return h.WalletHandler.Transfer(ctx, in, out)
}
