# Pinata Pinner

Small utility that adds an entire directory recursively to [IPFS](https://ipfs.io/) using the [Pinata](https://pinata.cloud) pinning service.

*This utility is community-contributed and has no affiliation with Pinata or IPFS. It's released under MIT license.*

## Installation

If you have Go 1.12+ installed in your machine, you can just run:

````sh
go get -u github.com/ItalyPaleAle/pinatapinner
````

Alternatively, you can download pre-built binaries from the [releases page](https://github.com/ItalyPaleAle/pinatapinner/releases).

## Pin a folder

Pinata Pinner can only pin folders (if you want to pin individual files, the Pinata APIs are easy to invoke with just a curl call).

You will need your Pinata API key and secret, that you can get from your [Pinata account page](https://pinata.cloud/account). Set the key and secret as environmental variables:

````sh
export PINATA_API_KEY="..."
export PINATA_SECRET_KEY="..."
````

To pin a folder called "my_site":

````sh
pinatapinner my_site
````

You can optionally give a name to your pinned folder, which will be displayed on the Pinata portal only:

````sh
pinatapinner my_site "My personal website"
````

The result will contain the IPFS hash of your folder, for example:

````text
$ pinatapinner my_site "My personal website"
{"IpfsHash":"QmakGEBp4HJZ6tkFydbyvF6bVvFThqfAwnQS6F4D7ie7hL","PinSize":"17678985","Timestamp":"2019-07-12T19:02:10.076Z"}
````

You can then browse your folder via IPFS at the address `/ipfs/QmakGEBp4HJZ6tkFydbyvF6bVvFThqfAwnQS6F4D7ie7hL`. You can also browse it via a public gateway, for example: https://gateway.ipfs.io/ipfs/QmakGEBp4HJZ6tkFydbyvF6bVvFThqfAwnQS6F4D7ie7hL 
