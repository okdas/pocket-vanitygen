# pocket-vanitygen

Generates vanity/beautiful addresses for Pocket Network.

As pocket network addresses are hex addresses, you can only look for hex characters (numbers + abcdef).

It is recommended to not look for more than 5 characters per each pattern as it will take a long time to find a match.

## Installation

1. [Install golang](https://go.dev/dl/)
2. `go install github.com/okdas/pocket-vanitygen@latest`

## Usage

```
./pocket-vanitygen -patterns=420-69,000-00
Looking for:  [[420 69] [000 00]]
running on cpus:  10
Found 42038fdef512c0e5064f6e9412d6e7315e603e69 c91a5293bdd9ce06b9ff7597d58518a309aed726af9d087f03dd44de8f8fe9d68904131d8170b5c22be48f9c1a0baf7f6f44293f9b50f66407d84d77b5f2948d
Found 420eb1bd7c312de89445725d7052d267ac37e969 156e13e3ecfd55ef8d1fa528993f81a5786a5f2254f97b84eefe2aa28e1998f7d825a735b4b02893a1cbfba70868c7498c68eedc3fdbcf60d440e43f00b05ce2
Found 000512553570397f401dae9e61b22c2b0abef800 028bc44d3a2956d061a6ee0a22fb8f9f9b5a8c79e8c2978f03b24ddf8b2cc1e210d8b8196959b587d97c85589de9da782425aee9a249d76b653b979bb85f238a
```

As the output shows, you get an address and the private key for that address in one line.

You can specify as many patterns as you want, separated by commas. Pattern consists of beginning and end (dash separated, dash always must be provided e.g. 420- or -69). You can skip either beginning or end, but not both.

You can also limit amount of CPU threads to use with `-cpus` flag. By default it uses all available CPUs.
