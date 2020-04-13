from web3.auto import w3
import binascii
with open("/root/go/bin/node1/keystore/UTC--2020-04-10T00-34-12.744101299Z--efd51fa9f7b385742291a58af792acd9d1890e7d") as keyfile:
	encrypted_key = keyfile.read()
	private_key = w3.eth.account.decrypt(encrypted_key, 
	                                 '1')
	print(private_key)
	print(binascii.b2a_hex(private_key))