package client

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/basket/types/v1"
)

// QueryBasketCmd returns a query command that retrieves a basket.
func QueryBasketCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "basket [basket-denom]",
		Short:   "Gets the info for a basket",
		Long:    "Retrieves the information for a basket given a specific basket denom",
		Example: "regen q ecocredit basket eco.uC.NCT",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			client := types.NewQueryClient(ctx)

			res, err := client.Basket(cmd.Context(), &types.QueryBasketRequest{BasketDenom: args[0]})
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

// QueryBasketsCmd returns a query that retrieves an optionally paginated list of baskets.
func QueryBasketsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "baskets",
		Short: "Retrieves all baskets",
		Long:  "Retrieves all baskets currently in state, with optional pagination",
		Example: `
regen q ecocredit baskets
regen q ecocredit baskets --limit 10 --offset 10
		`,
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			client := types.NewQueryClient(ctx)
			res, err := client.Baskets(cmd.Context(), &types.QueryBasketsRequest{Pagination: pagination})
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "baskets")
	return cmd
}

// QueryBasketBalanceCmd returns a query command that retrieves the balance of a credit batch in the basket.
func QueryBasketBalanceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "basket-balance [basket-denom] [batch-denom]",
		Short: "Retrieves the balance of a credit batch in the basket",
		Long:  "Retrieves the balance of a credit batch in the basket",
		Example: `
regen q ecocredit basket-balance eco.uC.NCT C01-001-20210101-20220101-001
		`,
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			client := types.NewQueryClient(ctx)

			res, err := client.BasketBalance(cmd.Context(), &types.QueryBasketBalanceRequest{
				BasketDenom: args[0],
				BatchDenom:  args[1],
			})
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// QueryBasketBalancesCmd returns a query command that retrieves the balance of each credit batch for the given basket denom.
func QueryBasketBalancesCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "basket-balances [basket-denom]",
		Short: "Retrieves the balance of each credit batch for the given basket denom",
		Long:  "Retrieves the balance of each credit batch for the given basket denom",
		Example: `
regen q ecocredit basket-balances eco.uC.NCT
regen q ecocredit basket-balances eco.uC.NCT --limit 10 --offset 10
		`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			pagination, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			client := types.NewQueryClient(ctx)
			res, err := client.BasketBalances(cmd.Context(), &types.QueryBasketBalancesRequest{
				BasketDenom: args[0],
				Pagination:  pagination,
			})
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "basket-balances")

	return cmd
}

// QueryBasketFeeCmd returns a query command that retrieves the basket fees.
func QueryBasketFeeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "basket-fee",
		Short: "Retrieves the basket fee",
		Long:  "Retrieves the basket fee",
		Example: `
regen q ecocredit basket-fee
		`,
		Args: cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			client := types.NewQueryClient(ctx)
			res, err := client.BasketFee(cmd.Context(), &types.QueryBasketFeeRequest{})
			if err != nil {
				return err
			}

			return ctx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
