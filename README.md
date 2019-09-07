# FUDForum authentication provider for INN2

This repository contains FUDForum authentication provider for INN2.

Reason this project exists is a big need in federated (well, okay, kind of, INN2 isn't federated but peering software) forum for [Medium ISP](https://mediumfoundation.org) which works on Yggdrasil. After googling I decided that write one from scratch will be a very big deal and tried to re-use existing technologies. I've choosed NNTP (aka Usenet) server INN2 and FUDForum, which able to sync messages from and to NNTP server. After that came "the problem of authorization" - I was using htpasswd-formatted file first, but users will definetely want easy-to-use GUI (WUI, web UI, okay) to register or reset their passwords. And inn2-fudforum-auth appeared.

## Dependencies

Nothing except Go compiler. Version 1.9+ is recommended. All dependencies are vendored, so inn2-fudforum-auth will compile event in constraint environments.

Right now authentication provider is able to connect only to PostgreSQL database. MRs are welcome for other providers.

## Installation

*TBW*

## Configuration

*TBW*