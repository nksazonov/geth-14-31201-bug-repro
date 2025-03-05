# Repro code for issue #31203

## Bindings

Go bindings are already included in this repo. They were produced with the following commands:

```bash
solc --abi --bin ./Dummy.sol -o build/ --overwrite

abigen --bin build/Dummy.bin \                    
          --abi build/Dummy.abi \
          --pkg main \
          --type DummyPostShanghai \
          --out dummy_post_shanghai.go
```

with `solc` version `0.8.20` for Post-Shanghai, and `0.8.19` for Pre-Shanghai, with `PreShanghai` naming for the latter.

## Reproduction

With go-ethereum version `v1.14.13`, run the following command:

```bash
go run .
```

You will see the following output:

```txt
created tx to deploy dummy (pre shanghai) contract
failed to deploy dummy (post shanghai) contract: invalid opcode: PUSH0
```

## Other findings

- If a block is included right before the creating a deployment tx for the post-shanghai contract, the deployment will succeed.
  (uncomment line 46 in `main.go` to see this in action)
- Deployment of post-shanghai contract also succeeds with geth version `v1.15.5`
