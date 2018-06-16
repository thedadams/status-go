/*
	Package benchmarks contains tests that can be used
	to run benchmarks and stress tests of our cluster components.


	Example usage:

		1. Start a Whisper node with mail server capability:
			./build/bin/statusd \
				-networkid=4 \
				-maxpeers=100 \
				-shh \
				-shh.pow=0.002 \
				-shh.mailserver \
				-shh.passwordfile=./static/keys/wnodepassword \
				-log DEBUG
		2. Generate some messages:
			go test -v -timeout=30s -run TestSendMessages ./t/benchmarks \
				-peerurl=$ENODE_ADDR \
				-msgcount=200 \
				-msgbatchsize=50
		3. Retrieve them from the mail server:
			go test -v -timeout=30s -parallel 20 \
				-run TestConcurrentMailserverPeers
				./t/benchmarks \
				-peerurl=$ENODE_ADDR \
				-msgcount=200

		The result of the last command will tell you how long it took to
		retrieve 200 messages with 20 concurrent peers (20 * 200 messages
		in total).

		The result may be affected due to limitations of the host
		on which it was called. It's recommended running mail server
		on a different machine and running the third command
		from some beefy server.
*/

package benchmarks
