# go-utils

[![codecov](https://codecov.io/gh/yiranzai/go-utils/branch/main/graph/badge.svg)](https://codecov.io/gh/yiranzai/go-utils)
[![Go Report Card](https://goreportcard.com/badge/github.com/yiranzai/go-utils)](https://goreportcard.com/report/github.com/yiranzai/go-utils)
[![Sourcegraph](https://sourcegraph.com/github.com/yiranzai/go-utils/-/badge.svg)](https://sourcegraph.com/github.com/yiranzai/go-utils?badge)
[![Open Source Helpers](https://www.codetriage.com/yiranzai/go-utils/badges/users.svg)](https://www.codetriage.com/yiranzai/go-utils)
[![Release](https://img.shields.io/github/release/yiranzai/go-utils.svg?style=flat-square)](https://github.com/yiranzai/go-utils/releases)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fyiranzai%2Fgo-utils.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fyiranzai%2Fgo-utils?ref=badge_shield)

Golang utils for me.

## 目录

---

<!--ts-->
   * [go-utils](#go-utils)
      * [目录](#目录)
      * [Install](#install)
      * [Usage](#usage)
         * [Math](#math)
         * [Leetcode](#leetcode)
      * [License](#license)

<!-- Added by: runner, at: Sat Mar 27 11:21:47 UTC 2021 -->

<!--te-->

---

## Install

```sh
go get github.com/yiranzai/go-utils
```

## Usage

Something.

### Math

- `MinInt(vars ...int)` // get the min int vars
- `MinInt64(vars ...int64)` // get the min int64 vars
- `MaxInt(vars ...int)` // get the max int vars
- `MaxInt64(vars ...int64)` // get the max int64 vars
- `AbsInt(abs int) int` // abs int
- `AbsInt64(abs int64) int64` // abs int64

### Leetcode

- `GenerateTree(treeList []interface{})` // generate leetcode tree
- `GenerateList(list []int)` // generate leetcode list
- `LoopEqualList(t *testing.T, head *ListNode, list []int)` // test diff list and array
- `DeepEqualTree(t *testing.T, head, right *TreeNode)` // test recursion diff tree

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for the full license text.

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fyiranzai%2Fgo-utils.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fyiranzai%2Fgo-utils?ref=badge_large)
