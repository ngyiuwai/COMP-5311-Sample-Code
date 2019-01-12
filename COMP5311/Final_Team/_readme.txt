A Simple Blockchain Project for PolyU COMP5331 2018 Winter Semester



Planning to

- Replace teammate's code by my code

  (e.g.	Block will be rewrote in a more understandable way. And use a fixed length header, e.g. nonce = 5 --> nonce = 000005)
  (	Merkle Tree has some error. The tree building could be modified to a more efficent way too.)
  (	Database is slow. Maybe use some other data format, e.g. JSON/ XML but not external database library)
  (	TCP message size is limited at 8kb. Should allow messages to be partitioned and sent in muliple datagram)
- Rewrite the ugly code

- Write an introduction with graphs (data flow dragram? UML?)

in late Dec 2018/ early Jan 2019



Contribution:


- dingyuzhu	(github.com/dingyuzhu)
--- block.go

- Tony		(github.com/tony0021074)
--- database.go

- Franco	(github.com/francofong)
--- merkletree.go

- ngyiuwai	(github.com/ngyiuwai)
--- all others (UI, network, dataflow control, blockchain operations in memory)
--- i.e. UI.go, nodeMiner.go, nodeController.go, blockchain.go
--- also works on software specification


by Ng Yiu Wai (Karl), 1 Dec 2018
