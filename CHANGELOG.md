# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/).

## [v0.49.0] (2023-12-21)
[v0.49.0]: https://github.com/CoboGlobal/cobo-go-api/compare/v0.48.0...v0.49.0
### Added
- Add New Params: Add parameter `amount` for API coin_info for Custodial Wallet. https://github.com/CoboGlobal/cobo-go-api/pull/63

## [v0.48.0] (2023-12-07)
[v0.48.0]: https://github.com/CoboGlobal/cobo-go-api/compare/v0.47.0...v0.48.0
### Added
- Add New API: Add get_max_send_amount API for MPC Wallet. https://github.com/CoboGlobal/cobo-go-api/pull/60

## [v0.47.0] (2023-11-22)
[v0.47.0]: https://github.com/CoboGlobal/cobo-go-api/compare/v0.46.0...v0.47.0
### Added
- Add API for Primer Broker to facilitate MPC wallet integration with Cobo Guard. https://github.com/CoboGlobal/cobo-go-api/pull/56

## [v0.46.0] (2023-11-08)
[v0.46.0]: https://github.com/CoboGlobal/cobo-go-api/compare/v0.45.0...v0.46.0
### Added
- Add New API: Add transactions_by_time_ex API for Custody Wallet. https://github.com/CoboGlobal/cobo-go-api/pull/54

## [v0.45.0] (2023-10-26)
[v0.45.0]: https://github.com/CoboGlobal/cobo-go-api/compare/v0.44.0...v0.45.0
### Added
- Add New API: MPC Wallet add Update Address Description API.https://github.com/CoboGlobal/cobo-go-api/pull/53
### Changed
- Add New Params: Custodial Wallet New Withdraw Request and MPC Wallet Create Transaction API add remark param. https://github.com/CoboGlobal/cobo-go-api/pull/51

## [v0.44.0] (2023-07-27)
[v0.44.0]: https://github.com/CoboGlobal/cobo-go-api/compare/v0.43.0...v0.44.0
### Changed
- Add API to get result of SignMessage. https://github.com/CoboGlobal/cobo-go-api/pull/46

## [v0.43.0] (2023-07-20)
[v0.43.0]: https://github.com/CoboGlobal/cobo-go-api/compare/v0.42.0...v0.43.0
### Changed
- Add API to list MPC TSS Node info. https://github.com/CoboGlobal/cobo-go-api/pull/44

## [v0.42.0] (2023-06-27)
[v0.42.0]: https://github.com/CoboGlobal/cobo-go-api/compare/v0.41.0...v0.42.0
### Changed
- Bug fix. https://github.com/CoboGlobal/cobo-go-api/pull/41

## [v0.41.0] (2023-06-01)
[v0.41.0]: https://github.com/CoboGlobal/cobo-go-api/compare/v0.40.0...v0.41.0
### Changed
- MPC create transaction and RBF API support param fee_amount for UTXO model. https://github.com/CoboGlobal/cobo-go-api/pull/38

## [v0.40.0] (2023-04-20)
[v0.40.0]: https://github.com/CoboGlobal/cobo-go-api/compare/v0.39.0...v0.40.0
### Changed
- Sign nonce support lower Go version. https://github.com/CoboGlobal/cobo-go-api/pull/34

## [v0.39.0] (2023-04-18)
[v0.39.0]: https://github.com/CoboGlobal/cobo-go-api/compare/v0.38.0...v0.39.0
### Changed
- Support MPC Web3 wallet SaaS api. https://github.com/CoboGlobal/cobo-go-api/pull/25

## [v0.38.0] (2023-04-17)
[v0.38.0]: https://github.com/CoboGlobal/cobo-go-api/compare/v0.37.0...v0.38.0
### Changed
- check resp timestamp and signature. https://github.com/CoboGlobal/cobo-go-api/pull/30

## [v0.37.0] (2023-04-10)
[v0.37.0]: https://github.com/CoboGlobal/cobo-go-api/compare/v0.36.0...v0.37.0
### Changed
- Update request nonce. https://github.com/CoboGlobal/cobo-go-api/pull/28

## [v0.36.0] (2023-03-22)
[v0.36.0]: https://github.com/CoboGlobal/cobo-go-api/compare/v0.35.0...v0.36.0
### Changed
- Fix error return. https://github.com/CoboGlobal/cobo-go-api/pull/26

## [v0.35.0] (2023-03-13)
[v0.35.0]: https://github.com/CoboGlobal/cobo-go-api/compare/v0.34.0...v0.35.0
### Changed
- Update MPC wallet SaaS api. https://github.com/CoboGlobal/cobo-go-api/pull/22, https://github.com/CoboGlobal/cobo-go-api/pull/23

## [v0.34.0] (2023-02-16)
[v0.34.0]: https://github.com/CoboGlobal/cobo-go-api/compare/v0.33.0...v0.34.0
### Changed
- Update MPC wallet SaaS api. https://github.com/CoboGlobal/cobo-go-api/pull/16
- Update Web3 wallet SaaS api. https://github.com/CoboGlobal/cobo-go-api/pull/20

## [v0.33.0] (2022-12-26)
[v0.33.0]: https://github.com/CoboGlobal/cobo-go-api/compare/v0.32.0...v0.33.0

### Fixed
- Fixed the MPC Wallet creation transaction bug. https://github.com/CoboGlobal/cobo-go-api/pull/16

## [v0.32.0] (2022-12-23)
[v0.32.0]: https://github.com/CoboGlobal/cobo-go-api/compare/v0.31.0...v0.32.0

### Changed
- Update MPC client method. https://github.com/CoboGlobal/cobo-go-api/pull/12, https://github.com/CoboGlobal/cobo-go-api/pull/13


## [v0.31.0] (2022-12-20)
[v0.31.0]: https://github.com/CoboGlobal/cobo-go-api/compare/v0.30.0...v0.31.0

### Added
- Support MPC wallet SaaS api. https://github.com/CoboGlobal/cobo-go-api/pull/10


## [v0.30.0] (2022-12-05)
[v0.30.0]: https://github.com/CoboGlobal/cobo-go-api/compare/v0.29.0...v0.30.0

### Added
- Support Web3 wallet SaaS api. https://github.com/CoboGlobal/cobo-go-api/pull/8, https://github.com/CoboGlobal/cobo-go-api/pull/9

## [v0.29.0] (2022-07-11)
[v0.29.0]: https://github.com/CoboGlobal/cobo-go-api/compare/v0.28.0...v0.29.0

### Added
- The receiving address list supports paging queries. https://github.com/CoboGlobal/cobo-go-api/pull/5

## [v0.28.0] (2022-06-27)
[v0.28.0]: https://github.com/CoboGlobal/cobo-go-api/compare/v0.27.0...v0.28.0

### Added 
- Transaction records can be searched using TXID. https://github.com/CoboGlobal/cobo-go-api/pull/3


## [v0.27.0] (2022-04-30)
[v0.27.0]: https://github.com/CoboGlobal/cobo-go-api/compare/v0.26.0...v0.27.0

### Changed
- Update dependencies. https://github.com/CoboGlobal/cobo-go-api/commit/b2db00b077143e7b83654c783c16998ece9ff5e6




