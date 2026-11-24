# Symmetric Binary Tree

> Striver SDE Sheet 127 · Day 19 Binary Tree III — self-mirror check: recurse on (left, right) with the child pairs crossed. Go, stdlib only, tested.

## Problem

Is a binary tree symmetric around its vertical center line — its left half the mirror image of its right?

## Approach

"Symmetric" is a statement about *two* subtrees, so the natural helper takes two nodes: a and b mirror each other iff their values match, a's left mirrors b's right, and a's right mirrors b's left. The **crossed** recursive pairs are the entire algorithm — uncross them and you've written Identical (#121) instead, which happily accepts trees whose halves are equal rather than mirrored. Kick off with `mirror(root.Left, root.Right)`; empty trees and single nodes are symmetric by the shared nil base case.

**Complexity:** O(n) time, O(h) recursion.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
