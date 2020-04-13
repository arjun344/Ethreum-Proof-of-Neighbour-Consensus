echo "Making geth"

source ~/.profile

rm -rf geth
cd $HOME/go/src/github.com/ethereum/go-ethereum/cmd
go install -v ./geth

