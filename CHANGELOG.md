# Changelog

All notable changes to this project will be documented in this file.

### [1.4.2](https://github.com/airnity/network-manager/compare/1.4.1...1.4.2) (2023-08-18)

### Bug Fixes

- **gre.go:** improve error handling and error messages in createTunnel and deleteTunnel functions ([dc8c92d](https://github.com/airnity/network-manager/commit/dc8c92da1751d1ace0b004268d7154f78b8a3dcf))

### [1.4.1](https://github.com/airnity/network-manager/compare/1.4.0...1.4.1) (2023-08-18)

### Bug Fixes

- **gre.go:** handle error when creating tunnel in Synchronize() method ([ec63cf0](https://github.com/airnity/network-manager/commit/ec63cf0229c52f4ec98d103bcb920e2e6f29ed74))

## [1.4.0](https://github.com/airnity/network-manager/compare/1.3.1...1.4.0) (2023-08-18)

### Features

- **config.go:** add VRF field to Tunnel struct ([495a574](https://github.com/airnity/network-manager/commit/495a5742afbe6ee06d626d18be33e4200b37e382))
- **network-manager:** add support for GRE tunnels ([2c704c0](https://github.com/airnity/network-manager/commit/2c704c00b37679839e74a692b7304aa07e72246c))
- **network-manager:** add VRF support ([c85b7e8](https://github.com/airnity/network-manager/commit/c85b7e8d09c9f05f6796298a79240ebde23b8480))

### Bug Fixes

- **exec.go:** change log.Panic to log.Error and log the output of the command ([cb92515](https://github.com/airnity/network-manager/commit/cb9251566ff661cd39a7ddfeecd7aefc257b9427))

### [1.3.1](https://github.com/airnity/network-manager/compare/1.3.0...1.3.1) (2023-07-19)

### Bug Fixes

- **exec.go:** change log.Error to log.Panic for better error handling ([397b12a](https://github.com/airnity/network-manager/commit/397b12a43d9b790e30c7444b6411ca691fcaca48))

### Refactor

- **config.go:** remove unused 'State' field from NatRule struct ([688c1b5](https://github.com/airnity/network-manager/commit/688c1b51ec7b58213aff5cc29a4b2269046eb027))

## [1.3.0](https://github.com/airnity/network-manager/compare/1.2.0...1.3.0) (2023-07-18)

### Features

- **network-manager:** add support for NAT rules ([f095dee](https://github.com/airnity/network-manager/commit/f095dee082ac26abce3241373ca36c703b4a7a3a))

### Bug Fixes

- **vrf:** change logger parameter type from log.Logger to \*log.Logger in NewClient function ([44a8f47](https://github.com/airnity/network-manager/commit/44a8f478043a4029f948d93bd3531aca3a5790e6))

### Chore

- **deps:** bump golang.org/x/text from 0.3.7 to 0.3.8 ([c7c669c](https://github.com/airnity/network-manager/commit/c7c669c8332264f422d6bc096b1b83b18e2ba20b))

## [1.2.0](https://github.com/airnity/network-manager/compare/1.1.0...1.2.0) (2023-01-10)

### Features

- Add rp_filter=0 on gre tunnels ([304b080](https://github.com/airnity/network-manager/commit/304b080dac506e8c946c34ae97b0c8427c8fa1d9))

## [1.1.0](https://github.com/airnity/network-manager/compare/1.0.0...1.1.0) (2022-06-22)

### Features

- Change config file path ([6572dbd](https://github.com/airnity/network-manager/commit/6572dbda9a97b6b86ce61fc0e2199149ea1a0fca))

## 1.0.0 (2022-06-22)

### Chore

- Add pre-commit hooks ([ce2ad23](https://github.com/airnity/network-manager/commit/ce2ad2334003785f3768dcb134ec92ed6a0badee))
- First commit ([2794325](https://github.com/airnity/network-manager/commit/2794325b08dc5f79b4750a73add28be3cd571c6a))

### Continuous Integration

- Add release worklfow ([68fdcf5](https://github.com/airnity/network-manager/commit/68fdcf58f44faca6dbe9ebf6f9ddae81846883dc))
- Add workflow for master branch ([7772d1f](https://github.com/airnity/network-manager/commit/7772d1f07fa4595c092d3c131d1c5a95b1d6e403))
- Update release workflow ([fb909d2](https://github.com/airnity/network-manager/commit/fb909d25e16fa8651124b10ee8928ea7f755a38d))
