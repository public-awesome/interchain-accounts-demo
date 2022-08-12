# Groups Demo

Placeholder section for `x/group` integration. This is a work-in-progress and will either be moved to a separate guide or replace the current demo.
This requires a fix to `x/group` in cosmos-sdk to ensure events are correctly propagated to the current context and emitted on successful execution of a group propsal.
This fix will be included in the next patch release of the sdk: `v0.46.1`

See: https://github.com/cosmos/cosmos-sdk/pull/12888

- Bootstrap two chains, configure the relayer and create an IBC connection (on top of clients that are created as well).

```bash
# hermes
make init-hermes

# go relayer
make init-golang-relayer
```

:warning: **NOTE:** When you want to use both relayers interchangeably, using both of these `make` commands will set up two seperate connections (which is not needed and can lead to confusion). In the case of using both relayers, perform:
```bash
make init-golang-rly
./network/hermes/restore-keys.sh
```

- Start the relayer.

```bash
#hermes
make start-hermes

#go relayer
make start-golang-rly
```

- Setup the CLI client configurations for test chains `test-1` and `test-2`.

```
icad config keyring-backend test --home ./data/test-1;
icad config node tcp://localhost:16657 --home ./data/test-1;
icad config --home ./data/test-1;

icad config keyring-backend test --home ./data/test-2;
icad config node tcp://localhost:26657 --home ./data/test-2;
icad config --home ./data/test-2;
```

- Create a new group using the `members.json` file provided.

```
icad tx group create-group $WALLET_1 test-metadata members.json --from $WALLET_1 --home ./data/test-1

icad q group group-info 1 --home ./data/test-1
```

- Create a new group policy using the `policy.json` file provided. A new `ThresholdDecisionPolicy` is created with a threshold of `2` and voting period of `60s`.

```
icad tx group create-group-policy $WALLET_1 1 policy-meta policy.json --home ./data/test-1

icad q group group-policies-by-group 1 --home ./data/test-1
```

- Extract the group policy address from the response JSON.
```
export GROUP_POLICY_ADDR=$(icad q group group-policies-by-group 1 --home ./data/test-1 --home ./data/test-1 -o json | jq -r '.group_policies[0].address') && echo $GROUP_POLICY_ADDR
cosmos1afk9zr2hn2jsac63h4hm60vl9z3e5u69gndzf7c99cqge3vzwjzsfwkgpd
```

- Send some funds to the group policy address

```
icad tx bank send $WALLET_1 $GROUP_POLICY_ADDR 10000000stake --home ./data/test-1
```

- Submit a new propsal to execute a `/intertx.MsgRegisterAccount` where the account owner is the group policy address.

```
icad tx group submit-proposal prop-register.json --home ./data/test-1 --from $WALLET_1

icad q group proposal 1 --home ./data/test-1
```

- Vote on the propsal using both of the group members.

```
icad tx group vote 1 $WALLET_1 --home ./data/test-1 VOTE_OPTION_YES meta --from $WALLET_1
icad tx group vote 1 $WALLET_2 --home ./data/test-1 VOTE_OPTION_YES meta --from $WALLET_2
```

- Execute the proposal once the voting period of `60s` has elapsed and votes have been tallied.

```
icad tx group exec 1 --from $WALLET_1 --home ./data/test-1 --gas 500000 -b block
```

- Submit a new proposal to execute a `/intertx.MsgSubmitTx` that includes a `MsgDelegate` to execute using the interchain account registered on chain `test-2`.

```
icad tx group submit-proposal prop-sendtx.json --home ./data/test-1 --from $WALLET_1

icad q group proposal 2 --home ./data/test-1
```

- Once again, vote on the proposal with both of the group members.

```
icad tx group vote 2 $WALLET_1 --home ./data/test-1 VOTE_OPTION_YES meta --from $WALLET_1
icad tx group vote 2 $WALLET_2 --home ./data/test-1 VOTE_OPTION_YES meta --from $WALLET_2
```

- Execute the propsal once it has been accepted.

```
icad tx group exec 2 --from $WALLET_1 --home ./data/test-1 --gas 500000 -b block
```