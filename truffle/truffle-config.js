module.exports = {
rpc: {
host:"localhost",
port:8543
},
networks: {
development: {
host: "localhost", //our network is running on localhost
port: 8543, // port where your blockchain is running
network_id: "1234",
from: "0xefd51fa9f7b385742291a58af792acd9d1890e7d", // use the account-id generated during the setup process
gas: 20000000
}
}
};