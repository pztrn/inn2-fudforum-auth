# FUDForum authentication provider for INN2

This repository contains FUDForum authentication provider for INN2.

Reason this project exists is a big need in federated (well, okay, kind of, INN2 isn't federated but peering software) forum for [Medium ISP](https://mediumfoundation.org) which works on Yggdrasil. After googling I decided that write one from scratch will be a very big deal and tried to re-use existing technologies. I've choosed NNTP (aka Usenet) server INN2 and FUDForum, which able to sync messages from and to NNTP server. After that came "the problem of authorization" - I was using htpasswd-formatted file first, but users will definetely want easy-to-use GUI (WUI, web UI, okay) to register or reset their passwords. And inn2-fudforum-auth appeared.

## Dependencies

Nothing except Go compiler. Version 1.9+ is recommended. All dependencies are vendored, so inn2-fudforum-auth will compile event in constraint environments.

Right now authentication provider is able to connect only to PostgreSQL database. MRs are welcome for other providers.

## Installation

Right now you can build it yourself by using ``go get``:

```bash
go get -u -v develop.pztrn.name/pztrn/inn2-fudforum-auth
```

Binary will be placed in ``$GOPATH/bin``. Use [this configuration example](/inn2-fudforum-auth.dist.yaml) as example and tune it.

## Configuration

### Provider

See [this configuration example](/inn2-fudforum-auth.dist.yaml), it has comments for each section.

Don't forget to define default group and other groups your INN2 is using in ``readers.conf``!

### INN2

INN2 authentication uses ``auth`` and ``access`` blocks, first for authentication and second for authorization. We should define them both for each users group. Example for ``admin`` group from configuration example:

```text
auth admin {
    hosts: *
    auth: /usr/local/bin/inn2-fudforum-auth -config /etc/news/inn2-fudforum-auth.yaml
}

access admin {
    users: "*@admin"
    newsgroups: *
}
```
